package seth

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

// MakeDeserializer creates the default deserializer
func MakeDeserializer() Deserializer {
	return deserializerImpl{}
}

type deserializerImpl struct{}

func (d deserializerImpl) Deserialize(
	item interface{},
	r io.Reader,
) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	var m map[string]*json.RawMessage
	err = json.Unmarshal(data, &m)

	fmt.Println(string(data), err, m["t"], m["value"])

	return nil
}
