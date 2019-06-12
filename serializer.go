package goseth

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
)

type serializerImpl struct {
	dict            map[string]reflect.Value
	serializedPtr   map[string]bool
	dictToSerialize []string
	maxLayer        int
}

// MakeSerializer creates a default serializer
func MakeSerializer() Serializer {
	return &serializerImpl{
		dict: make(map[string]reflect.Value),
	}
}

func (s *serializerImpl) Serialize(
	item interface{},
	writer io.Writer,
) error {
	s.serializedPtr = make(map[string]bool)
	s.dictToSerialize = nil
	value := reflect.ValueOf(item)
	id := s.addToDict(value)
	err := s.serializeToWriter(writer, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *serializerImpl) typeString(value reflect.Value) string {
	name := value.Type().String()
	pktPath := value.Type().PkgPath()

	if pktPath == "" {
		return name
	}

	tokens := strings.Split(pktPath, "/")
	tokens = tokens[0 : len(tokens)-1]
	pktPath = strings.Join(tokens, "/")

	return fmt.Sprintf("%s/%s", pktPath, name)
}

func (s *serializerImpl) itemID(
	ptr uintptr,
	value reflect.Value,
) string {
	if s.isZero(value) {
		return "0"
	}
	id := fmt.Sprintf("%d@%s", ptr, s.typeString(value))
	return id
}

func (s *serializerImpl) isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Invalid:
		return true
	case reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	}
	return false
}

func (s *serializerImpl) addToDict(
	value reflect.Value,
) string {
	v := s.strip(value)

	var ptr uintptr
	if !v.CanAddr() {
		ptr = 0
	} else {
		ptr = v.UnsafeAddr()
	}
	id := s.itemID(ptr, v)

	if _, ok := s.dict[id]; ok {
		return id
	}
	s.dict[id] = v
	return id
}

func (s *serializerImpl) strip(v reflect.Value) reflect.Value {
	switch v.Kind() {
	case reflect.Ptr,
		reflect.Interface:
		return s.strip(v.Elem())
	default:
		return v
	}
}

func (s *serializerImpl) serializeToWriter(
	writer io.Writer,
	id string,
) error {
	fmt.Fprintf(writer, `{"root":"%s","dict":`, id)
	s.addToDictToSerialize(id)
	err := s.serializeDict(writer)
	if err != nil {
		return err
	}
	fmt.Fprintf(writer, `}`)
	return nil
}

func (s *serializerImpl) addToDictToSerialize(id string) {
	if _, ok := s.serializedPtr[id]; ok {
		return
	}

	s.serializedPtr[id] = true
	s.dictToSerialize = append(s.dictToSerialize, id)
}

func (s *serializerImpl) serializeDict(
	writer io.Writer,
) error {
	fmt.Fprintf(writer, "{")

	var i uint64
	for len(s.dictToSerialize) > 0 {
		itemID := s.dictToSerialize[0]
		s.dictToSerialize = s.dictToSerialize[1:]

		if i > 0 {
			fmt.Fprintf(writer, `,`)
		}
		fmt.Fprintf(writer, `"%s":`, itemID)

		err := s.serializeItem(writer, itemID)
		if err != nil {
			return err
		}
		i++
	}

	fmt.Fprintf(writer, "}")
	return nil
}

func (s *serializerImpl) serializeItem(
	writer io.Writer,
	id string,
) error {
	if id == "0" {
		fmt.Fprintf(writer, `{"v":0}`)
		return nil
	}

	value := s.dict[id]
	switch value.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(writer, `{"v":%d,"t":"%s","k":%d}`,
			value.Int(), s.typeString(value), value.Kind())
	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(writer, `{"v":%f,"t":"%s","k":%d}`,
			value.Float(), s.typeString(value), value.Kind())
	case reflect.Bool:
		fmt.Fprintf(writer, `{"v":%t,"t":"%s","k":%d}`,
			value.Bool(), s.typeString(value), value.Kind())
	case reflect.Slice:
		fmt.Fprintf(writer, "[")
		for i := 0; i < value.Len(); i++ {
			f := value.Index(i)
			fID := s.addToDict(f)
			s.addToDictToSerialize(fID)
			if i > 0 {
				fmt.Fprint(writer, ",")
			}
			fmt.Fprintf(writer, `"%s"`, fID)
		}
		fmt.Fprintf(writer, "]")
	case reflect.Struct:
		fmt.Fprintf(writer, `{"v":{`)
		for i := 0; i < value.NumField(); i++ {
			f := value.Field(i)
			fID := s.addToDict(f)
			s.addToDictToSerialize(fID)
			if i > 0 {
				fmt.Fprint(writer, ",")
			}
			fieldName := value.Type().Field(i).Name
			fmt.Fprintf(writer, `"%s":"%s"`, fieldName, fID)
		}
		fmt.Fprintf(writer, `},"t":"%s","k":%d}`,
			s.typeString(value), value.Kind())

	default:
		return errors.New(
			"type kind " + value.Kind().String() + " is not supported.")
	}
	return nil
}
