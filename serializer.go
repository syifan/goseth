package goseth

import (
	"errors"
	"fmt"
	"io"
	"reflect"
)

// MakeSerializer creates a default serializer
func MakeSerializer() Serializer {
	return serializerImpl{}
}

type dictEntry struct {
	ptr   uintptr
	value reflect.Value
}

type serializerImpl struct {
	dictToSerialize []dictEntry
	serializedPtr   map[uintptr]bool
}

func (s serializerImpl) Serialize(
	item interface{},
	writer io.Writer,
) error {
	s.dictToSerialize = nil
	s.serializedPtr = make(map[uintptr]bool)

	value := reflect.ValueOf(item)
	return s.serializeItem(value, writer, true)
}

func (s *serializerImpl) serializeItem(
	value reflect.Value,
	writer io.Writer,
	isRoot bool,
) error {
	fmt.Fprintf(writer, `{"value":`)
	switch value.Kind() {
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		fmt.Fprintf(writer, "%d", value.Int())
	case
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		fmt.Fprintf(writer, "%d", value.Uint())
	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(writer, "%.15f", value.Float())
	case reflect.Bool:
		fmt.Fprintf(writer, "%t", value.Bool())
	case reflect.Struct:
		err := s.serializeStruct(value, writer)
		if err != nil {
			return err
		}
	case reflect.Ptr:
		err := s.serializePtr(value, writer)
		if err != nil {
			return err
		}
	default:
		return errors.New(
			"type " + value.Kind().String() + " is not supported")
	}

	fmt.Fprintf(writer, `,"type":"%s"`, value.Type())

	if isRoot {
		s.serializeDict(writer)
	}

	fmt.Fprintf(writer, "}")

	return nil
}

func (s *serializerImpl) serializeStruct(
	value reflect.Value,
	writer io.Writer,
) error {
	fmt.Fprint(writer, "{")

	for i := 0; i < value.NumField(); i++ {
		if i > 0 {
			fmt.Fprint(writer, ",")
		}
		fmt.Fprintf(writer, `"%s":`, value.Type().Field(i).Name)
		s.serializeItem(value.Field(i), writer, false)
	}

	fmt.Fprint(writer, "}")

	return nil
}

func (s *serializerImpl) serializePtr(
	value reflect.Value,
	writer io.Writer,
) error {
	s.toSerializeInDict(value)
	fmt.Fprint(writer, value.Pointer())
	return nil
}

func (s *serializerImpl) toSerializeInDict(value reflect.Value) {
	if _, ok := s.serializedPtr[value.Pointer()]; ok {
		return
	}

	s.dictToSerialize = append(s.dictToSerialize,
		dictEntry{value.Pointer(), value.Elem()},
	)
	s.serializedPtr[value.Pointer()] = true
}

func (s *serializerImpl) serializeDict(w io.Writer) {
	if len(s.dictToSerialize) == 0 {
		return
	}

	fmt.Fprintf(w, `,"dict":{`)
	count := 0
	for len(s.dictToSerialize) > 0 {
		v := s.dictToSerialize[0]
		s.dictToSerialize = s.dictToSerialize[1:]
		if count > 0 {
			fmt.Fprint(w, `,`)
		}
		fmt.Fprintf(w, `"%d":`, v.ptr)
		s.serializeItem(v.value, w, false)
		count++
	}

	fmt.Fprintf(w, "}")
}
