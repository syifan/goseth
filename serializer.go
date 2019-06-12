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
	dict            map[uintptr]reflect.Value
	serializedPtr   map[uintptr]bool
	dictToSerialize []uintptr
	maxLayer        int
}

// MakeSerializer creates a default serializer
func MakeSerializer() Serializer {
	return &serializerImpl{
		dict: make(map[uintptr]reflect.Value),
	}
}

func (s *serializerImpl) Serialize(
	item interface{},
	writer io.Writer,
) error {
	s.serializedPtr = make(map[uintptr]bool)
	s.dictToSerialize = nil
	value := reflect.ValueOf(item)
	ptr := value.Pointer()
	elem := value.Elem()
	s.addToDict(ptr, elem)
	err := s.serializeToWriter(writer, ptr)
	if err != nil {
		return err
	}
	return nil
}

func (s *serializerImpl) addToDict(
	ptr uintptr,
	value reflect.Value,
) {
	if _, ok := s.dict[ptr]; ok {
		return
	}
	s.dict[ptr] = value
}

func (s *serializerImpl) serializeToWriter(
	writer io.Writer,
	rootPtr uintptr,
) error {
	fmt.Fprintf(writer, `{"root":%d,"dict":`, rootPtr)
	s.addToDictToSerialize(rootPtr)
	err := s.serializeDict(writer)
	if err != nil {
		return err
	}
	fmt.Fprintf(writer, `}`)
	return nil
}

func (s *serializerImpl) addToDictToSerialize(ptr uintptr) {
	if _, ok := s.serializedPtr[ptr]; ok {
		return
	}

	s.serializedPtr[ptr] = true
	s.dictToSerialize = append(s.dictToSerialize, ptr)
}

func (s *serializerImpl) serializeDict(
	writer io.Writer,
) error {
	fmt.Fprintf(writer, "{")

	var i uint64
	for len(s.dictToSerialize) > 0 {
		ptr := s.dictToSerialize[0]
		s.dictToSerialize = s.dictToSerialize[1:]

		if i > 0 {
			fmt.Fprintf(writer, `,`)
		}
		fmt.Fprintf(writer, `"%d":`, ptr)

		err := s.serializeItem(writer, ptr)
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
	ptr uintptr,
) error {
	value := s.dict[ptr]
	switch value.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(writer, `{"v":%d,"t":"%s","k":%d,"p":%d}`,
			value.Int(), value.Type(), value.Kind(), ptr)
	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(writer, `{"v":%f,"t":"%s","k":%d,"p":%d}`,
			value.Float(), value.Type(), value.Kind(), ptr)
	case reflect.Bool:
		fmt.Fprintf(writer, `{"v":%t,"t":"%s","k":%d,"p":%d}`,
			value.Bool(), value.Type(), value.Kind(), ptr)
	case reflect.Ptr:
		s.serializeItem(writer, value.Pointer())
	case reflect.Struct:
		fmt.Fprintf(writer, `{"v":{`)
		for i := 0; i < value.NumField(); i++ {
			f := value.Field(i)
			var fPtr uintptr
			if f.Kind() == reflect.Ptr {
				fPtr = f.Pointer()
			} else {
				fPtr = f.Addr().Pointer()
			}
			s.addToDict(fPtr, f)
			s.addToDictToSerialize(fPtr)
			if i > 0 {
				fmt.Fprint(writer, ",")
			}
			fieldName := value.Type().Field(i).Name
			fmt.Fprintf(writer, `"%s":%d`, fieldName, fPtr)
			fmt.Println("processing field ", i, f, fPtr, fieldName, f.Type(), f.Kind())
		}
		fmt.Fprintf(writer, `},"t":"%s","k":%d,"p":%d}`,
			value.Type(), value.Kind(), ptr)

	default:
		return errors.New(
			"type kind " + value.Kind().String() + " is not supported.")
	}
	return nil
}
