package goseth

import (
	"fmt"
	"io"
	"reflect"
	"strings"
)

type serializeItem struct {
	id    uint64
	depth int
	item  reflect.Value
}

type serializer struct {
	root     any
	maxDepth int

	dict       []serializeItem
	nextDictID uint64
}

func (s *serializer) SetRoot(item any) {
	s.root = item
}

func (s *serializer) SetMaxDepth(depth int) {
	s.maxDepth = depth
}

func (s *serializer) Serialize(writer io.Writer) error {
	s.dict = make([]serializeItem, 0)
	s.nextDictID = 0

	s.addToDict(reflect.ValueOf(s.root), 0)

	fmt.Fprintf(writer, `{"r":"0","dict":`)
	s.serializeDict(writer)
	fmt.Fprintf(writer, `}`)

	return nil
}

func (s *serializer) addToDict(v reflect.Value, depth int) uint64 {
	v = s.strip(v)
	s.dict = append(s.dict, serializeItem{
		id:    s.nextDictID,
		depth: depth,
		item:  v,
	})

	id := s.nextDictID

	s.nextDictID++

	return id
}

func (s *serializer) serializeDict(writer io.Writer) {
	fmt.Fprintf(writer, `{`)

	count := 0
	for len(s.dict) > 0 {
		item := s.dict[0]
		s.dict = s.dict[1:]

		k := item.id
		v := item.item
		v = s.strip(v)

		if count > 0 {
			fmt.Fprintf(writer, `,`)
		}
		count++

		if s.isZero(v) {
			fmt.Fprintf(writer, `"%d":{"k":0,"t":"0","v":null}`, k)
		} else {
			fmt.Fprintf(writer, `"%d":{"k":%d,"t":"%s"`,
				k, v.Kind(), s.typeString(v))
			if s.maxDepth < 0 || item.depth < s.maxDepth {
				fmt.Fprint(writer, `,"v":`)
				s.serializeValue(writer, v, item.depth)
			}
			fmt.Fprintf(writer, `}`)
		}
	}

	fmt.Fprintf(writer, `}`)
}

func (s *serializer) isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Invalid:
		return true
	case reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	}
	return false
}

func (s *serializer) strip(v reflect.Value) reflect.Value {
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		return s.strip(v.Elem())
	default:
		return v
	}
}

func (s *serializer) serializeValue(
	writer io.Writer,
	v reflect.Value,
	depth int,
) {
	switch v.Kind() {
	case reflect.Bool:
		fmt.Fprintf(writer, `%t`, v.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(writer, `%d`, v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fmt.Fprintf(writer, `%d`, v.Uint())
	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(writer, `%f`, v.Float())
	case reflect.String:
		fmt.Fprintf(writer, `"%s"`, v.String())
	case reflect.Slice:
		s.serializeSlice(writer, v, depth)
	case reflect.Map:
		s.serializeMap(writer, v, depth)
	case reflect.Struct:
		s.serializeStruct(writer, v, depth)
	default:
		panic(fmt.Sprintf("kind %d not supported", v.Kind()))
	}
}

func (s *serializer) serializeMap(
	writer io.Writer,
	v reflect.Value,
	depth int,
) {
	fmt.Fprintf(writer, `{`)

	for i, key := range v.MapKeys() {
		if i > 0 {
			fmt.Fprint(writer, `,`)
		}

		keyID := s.addToDict(key, depth+1)
		valueID := s.addToDict(v.MapIndex(key), depth+1)
		fmt.Fprintf(writer, `"%d":"%d"`, keyID, valueID)
	}

	fmt.Fprintf(writer, `}`)
}

func (s *serializer) serializeSlice(
	writer io.Writer,
	v reflect.Value,
	depth int,
) {
	fmt.Fprintf(writer, `[`)
	for i := 0; i < v.Len(); i++ {
		if i > 0 {
			fmt.Fprint(writer, `,`)
		}

		id := s.addToDict(reflect.ValueOf(v.Index(i).Interface()), depth+1)
		fmt.Fprintf(writer, `"%d"`, id)
	}
	fmt.Fprintf(writer, `]`)
}

func (s *serializer) serializeStruct(
	writer io.Writer,
	v reflect.Value,
	depth int,
) {
	fmt.Fprintf(writer, `{`)

	for i := 0; i < v.NumField(); i++ {
		if i > 0 {
			fmt.Fprint(writer, `,`)
		}

		field := v.Field(i)
		fieldID := s.addToDict(field, depth+1)

		fmt.Fprintf(writer, `"%s":"%d"`, v.Type().Field(i).Name, fieldID)
	}

	fmt.Fprintf(writer, `}`)
}

func (s *serializer) typeString(value reflect.Value) string {
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
