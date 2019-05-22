package seth

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strconv"
)

// MakeDeserializer creates the default deserializer
func MakeDeserializer() Deserializer {
	return deserializerImpl{}
}

type deserializerImpl struct{}

func (d deserializerImpl) Deserialize(
	r io.Reader,
) (interface{}, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var m map[string]*json.RawMessage
	err = json.Unmarshal(data, &m)

	return d.parseValue(m)
}

func (d deserializerImpl) parseValue(
	m map[string]*json.RawMessage,
) (interface{}, error) {
	typeNameBytes, err := m["t"].MarshalJSON()
	if err != nil {
		return nil, err
	}
	var typeName string
	json.Unmarshal(typeNameBytes, &typeName)

	switch typeName {
	case "int":
		return d.parseIntValue(m["value"])
	}
	return nil, nil
}

func (d deserializerImpl) parseIntValue(m *json.RawMessage) (int, error) {
	valueBytes, err := m.MarshalJSON()
	if err != nil {
		return 0, err
	}
	valueInt, err := strconv.ParseInt(string(valueBytes), 10, 32)
	return int(valueInt), nil
}
