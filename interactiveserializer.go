package goseth

import (
	"io"
	"reflect"
)

// InteractiveSerializer can gradually serialize the whole data tree.
type InteractiveSerializer interface {
	Serializer
	Reset()
}

// NewInteractiveSerializer returns a new InteractiveSerializer
func NewInteractiveSerializer() InteractiveSerializer {
	s := &interactiveSerializerImpl{
		serializer: MakeSerializer().(*serializerImpl),
	}
	return s
}

type interactiveSerializerImpl struct {
	serializer *serializerImpl
}

func (s *interactiveSerializerImpl) Serialize(
	item interface{},
	writer io.Writer,
) error {
	s.serializer.serializedPtr = make(map[string]bool)
	s.serializer.dictToSerialize = nil
	value := reflect.ValueOf(item)
	id := s.serializer.addToDict(value)
	err := s.serializer.serializeToWriter(writer, id, 0)
	if err != nil {
		return err
	}
	return nil
}

func (s *interactiveSerializerImpl) Reset() {
	s.serializer = MakeSerializer().(*serializerImpl)
}
