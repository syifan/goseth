package goseth

import "io"

// Serializer can serialize anything into a json format
type Serializer interface {
	Serialize(item interface{}, writer io.Writer) error
}

// Deserializer can parse a serialized string back to an object
type Deserializer interface {
	Deserialize(reader io.Reader) (interface{}, error)
}
