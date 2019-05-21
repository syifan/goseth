package seth

import (
	"io"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

func MakeDeserializer() Deserializer {
	return deserializerImpl{}
}

type valueJson struct {
	value interface{}
	t string
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
