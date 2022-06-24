// Package goseth defines a serializer for Go data structures.
package goseth

import "io"

// Serializer can serialize anything into a json format.
type Serializer interface {
	// SetRoot sets the item to be serialized.
	SetRoot(item any)

	// SetMaxDepth sets the maximum depth of the serialization. If the depth
	// is exceeded, the serializer will stop serializing the item.
	// SetMaxDepth(-1) means no limit.
	SetMaxDepth(depth int)

	// Serialize serializes the item into the writer.
	Serialize(writer io.Writer) error
}

// NewSerializer returns a new default serializer.
func NewSerializer() Serializer {
	return &serializer{maxDepth: -1}
}
