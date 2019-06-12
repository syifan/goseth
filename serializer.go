package goseth

import (
	"errors"
	"fmt"
	"io"
	"reflect"
)

type rootObj struct {
	root *elem
	dict map[uint64]dictEntry
}

type elem struct {
	value    interface{}
	itemType string
	itemKind int
}

type dictEntry struct {
	ptr   uint64
	value reflect.Value
}

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
	ptr := value.Pointer()
	elem := value.Elem()
	id := s.itemID(ptr, elem)
	s.addToDict(id, elem)
	err := s.serializeToWriter(writer, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *serializerImpl) itemID(
	ptr uintptr,
	value reflect.Value,
) string {
	id := fmt.Sprintf("%d@%s.%s", ptr, value.Type().PkgPath(), value.Type())
	return id
}

func (s *serializerImpl) addToDict(
	id string,
	value reflect.Value,
) {
	if _, ok := s.dict[id]; ok {
		return
	}
	s.dict[id] = value
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
	value := s.dict[id]
	switch value.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(writer, `{"v":%d,"t":"%s","k":%d,"p":"%s"}`,
			value.Int(), value.Type(), value.Kind(), id)
	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(writer, `{"v":%f,"t":"%s","k":%d,"p":"%s"}`,
			value.Float(), value.Type(), value.Kind(), id)
	case reflect.Bool:
		fmt.Fprintf(writer, `{"v":%t,"t":"%s","k":%d,"p":"%s"}`,
			value.Bool(), value.Type(), value.Kind(), id)
	case reflect.Ptr:
		itemID := s.itemID(value.Pointer(), value.Elem())
		s.serializeItem(writer, itemID)
	case reflect.Struct:
		fmt.Fprintf(writer, `{"v":{`)
		for i := 0; i < value.NumField(); i++ {
			f := value.Field(i)
			var itemID string
			if f.Kind() == reflect.Ptr {
				fPtr := f.Pointer()
				itemID = s.itemID(fPtr, f.Elem())
			} else {
				fPtr := f.Addr().Pointer()
				itemID = s.itemID(fPtr, f)
			}
			s.addToDict(itemID, f)
			s.addToDictToSerialize(itemID)
			if i > 0 {
				fmt.Fprint(writer, ",")
			}
			fieldName := value.Type().Field(i).Name
			fmt.Fprintf(writer, `"%s":"%s"`, fieldName, itemID)
			fmt.Println("processing field ", i, f, itemID, fieldName, f.Type(), f.Kind())
		}
		fmt.Fprintf(writer, `},"t":"%s","k":%d,"p":"%s"}`,
			value.Type(), value.Kind(), id)

	default:
		return errors.New(
			"type kind " + value.Kind().String() + " is not supported.")
	}
	return nil
}
