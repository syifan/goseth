package seth

import (
	"encoding/json"
	"errors"
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
	case "bool":
		return d.parseBoolValue(m["value"])
	case "int":
		return d.parseIntValue(m["value"])
	case "int8":
		return d.parseInt8Value(m["value"])
	case "int16":
		return d.parseInt16Value(m["value"])
	case "int32":
		return d.parseInt32Value(m["value"])
	case "int64":
		return d.parseInt64Value(m["value"])
	case "uint":
		return d.parseUintValue(m["value"])
	case "uint8":
		return d.parseUint8Value(m["value"])
	case "uint16":
		return d.parseUint16Value(m["value"])
	case "uint32":
		return d.parseUint32Value(m["value"])
	case "uint64":
		return d.parseUint64Value(m["value"])

	default:
		return nil, errors.New("type " + typeName + "is not supported")
	}
}

func (d deserializerImpl) parseBoolValue(m *json.RawMessage) (bool, error) {
	valueString := rawMsgMustConvertToString(m)
	switch valueString {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, errors.New("bool value can only be true or false")
	}
}

func (d deserializerImpl) parseIntValue(m *json.RawMessage) (int, error) {
	valueString := rawMsgMustConvertToString(m)
	value := stringMustConvertToInt(valueString)
	return int(value), nil
}

func (d deserializerImpl) parseInt8Value(m *json.RawMessage) (int8, error) {
	valueString := rawMsgMustConvertToString(m)
	value := stringMustConvertToInt(valueString)
	return int8(value), nil
}

func (d deserializerImpl) parseInt16Value(m *json.RawMessage) (int16, error) {
	valueString := rawMsgMustConvertToString(m)
	value := stringMustConvertToInt(valueString)
	return int16(value), nil
}

func (d deserializerImpl) parseInt32Value(m *json.RawMessage) (int32, error) {
	valueString := rawMsgMustConvertToString(m)
	value := stringMustConvertToInt(valueString)
	return int32(value), nil
}

func (d deserializerImpl) parseInt64Value(m *json.RawMessage) (int64, error) {
	valueString := rawMsgMustConvertToString(m)
	value := stringMustConvertToInt(valueString)
	return int64(value), nil
}

func (d deserializerImpl) parseUintValue(m *json.RawMessage) (uint, error) {
	valueString := rawMsgMustConvertToString(m)
	value := stringMustConvertToInt(valueString)
	return uint(value), nil
}

func (d deserializerImpl) parseUint8Value(m *json.RawMessage) (uint8, error) {
	valueString := rawMsgMustConvertToString(m)
	value := stringMustConvertToInt(valueString)
	return uint8(value), nil
}

func (d deserializerImpl) parseUint16Value(m *json.RawMessage) (uint16, error) {
	valueString := rawMsgMustConvertToString(m)
	value := stringMustConvertToInt(valueString)
	return uint16(value), nil
}

func (d deserializerImpl) parseUint32Value(m *json.RawMessage) (uint32, error) {
	valueString := rawMsgMustConvertToString(m)
	value := stringMustConvertToInt(valueString)
	return uint32(value), nil
}

func (d deserializerImpl) parseUint64Value(m *json.RawMessage) (uint64, error) {
	valueString := rawMsgMustConvertToString(m)
	value := stringMustConvertToInt(valueString)
	return uint64(value), nil
}

func rawMsgMustConvertToString(m *json.RawMessage) string {
	valueBytes, err := m.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(valueBytes)
}

func stringMustConvertToInt(s string) int64 {
	valueInt, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return valueInt
}

func stringMustConvertToUint(s string) uint64 {
	valueInt, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return valueInt
}
