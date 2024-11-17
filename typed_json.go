package gotypedjson

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type JSONTYPE int

type CustomCodec map[JSONTYPE]Codec

type Codec struct {
	Encode func(val any) (string, error)
	Decode func(s string) (any, error)
}

const (
	INT JSONTYPE = iota
	INT8
	INT16
	INT32
	INT64
	UINT
	UINT8
	UINT16
	UINT32
	UINT64
	FLOAT32
	FLOAT64
	STRING
	BOOL
)

type TypedJson struct {
	Type  JSONTYPE `json:"Type"`
	Value any      `json:"Value"`

	customCodec CustomCodec
}

func NewTypedJson(jsonType JSONTYPE, Value any, customCodec CustomCodec) *TypedJson {
	for key, value := range customCodec {
		if value.Encode == nil {
			panic(fmt.Sprintf("key %d has a nil encoder", key))
		}

		if value.Decode == nil {
			panic(fmt.Sprintf("key %d has a nil decoder", key))
		}
	}

	return &TypedJson{
		Type:        jsonType,
		Value:       Value,
		customCodec: customCodec,
	}
}

func (typedJson *TypedJson) UnmarshalJSON(b []byte) error {
	temp := &struct {
		Type  JSONTYPE `json:"Type"`
		Value string   `json:"Value"`
	}{}

	if err := json.Unmarshal(b, temp); err != nil {
		return err
	}

	typedJson.Type = temp.Type

	// try the custom types
	if typedJson.customCodec != nil {
		if encoder, ok := typedJson.customCodec[temp.Type]; ok {
			val, err := encoder.Decode(temp.Value)
			if err != nil {
				return err
			}

			typedJson.Value = val

			return nil
		}
	}

	// try the default types
	switch temp.Type {
	case INT:
		val, err := strconv.ParseInt(temp.Value, 10, 0)
		if err != nil {
			return fmt.Errorf("failed to convert %s, to an int", temp.Value)
		}
		typedJson.Value = int(val)
	case INT8:
		val, err := strconv.ParseInt(temp.Value, 10, 8)
		if err != nil {
			return fmt.Errorf("failed to convert %s, to an int8", temp.Value)
		}
		typedJson.Value = int8(val)
	case INT16:
		val, err := strconv.ParseInt(temp.Value, 10, 16)
		if err != nil {
			return fmt.Errorf("failed to convert %s, to an int16", temp.Value)
		}
		typedJson.Value = int16(val)
	case INT32:
		val, err := strconv.ParseInt(temp.Value, 10, 32)
		if err != nil {
			return fmt.Errorf("failed to convert %s, to an int32", temp.Value)
		}
		typedJson.Value = int32(val)
	case INT64:
		val, err := strconv.ParseInt(temp.Value, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to convert %s, to an int64", temp.Value)
		}
		typedJson.Value = int64(val)
	case UINT:
		val, err := strconv.ParseUint(temp.Value, 10, 0)
		if err != nil {
			return fmt.Errorf("failed to convert %s, to a uint", temp.Value)
		}
		typedJson.Value = uint(val)
	case UINT8:
		val, err := strconv.ParseUint(temp.Value, 10, 8)
		if err != nil {
			return fmt.Errorf("failed to convert %s, to a uint8", temp.Value)
		}
		typedJson.Value = uint8(val)
	case UINT16:
		val, err := strconv.ParseUint(temp.Value, 10, 16)
		if err != nil {
			return fmt.Errorf("failed to convert %s, to a uint16", temp.Value)
		}
		typedJson.Value = uint16(val)
	case UINT32:
		val, err := strconv.ParseUint(temp.Value, 10, 32)
		if err != nil {
			return fmt.Errorf("failed to convert %s, to a uint32", temp.Value)
		}
		typedJson.Value = uint32(val)
	case UINT64:
		val, err := strconv.ParseUint(temp.Value, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to convert %s, to a uint64", temp.Value)
		}
		typedJson.Value = uint64(val)
	case FLOAT32:
		val, err := strconv.ParseFloat(temp.Value, 32)
		if err != nil {
			return fmt.Errorf("failed to convert %s, to a float32", temp.Value)
		}
		typedJson.Value = float32(val)
	case FLOAT64:
		val, err := strconv.ParseFloat(temp.Value, 64)
		if err != nil {
			return fmt.Errorf("failed to convert %s, to a float64", temp.Value)
		}
		typedJson.Value = float64(val)
	case STRING:
		typedJson.Value = string(temp.Value)
	case BOOL:
		val, err := strconv.ParseBool(temp.Value)
		if err != nil {
			return fmt.Errorf("failed to convert %s, to a bool", temp.Value)
		}
		typedJson.Value = bool(val)
	default:
		return fmt.Errorf("unknown type %d recevied for: %s", temp.Type, temp.Value)
	}

	return nil
}

func (typedJson *TypedJson) MarshalJSON() ([]byte, error) {
	temp := struct {
		Type  JSONTYPE `json:"Type"`
		Value string   `json:"Value"`
	}{
		Type: typedJson.Type,
	}

	// might be a custom type
	if typedJson.customCodec != nil {
		if encoder, ok := typedJson.customCodec[typedJson.Type]; ok {
			assignString, err := encoder.Encode(typedJson.Value)
			if err != nil {
				return nil, err
			}

			temp.Value = assignString
			return json.Marshal(temp)
		}
	}

	// check the defualt types
	switch typedJson.Type {
	case INT:
		if _, ok := typedJson.Value.(int); !ok {
			return nil, fmt.Errorf("failed to cast '%v' to an int", typedJson.Value)
		}

		temp.Value = strconv.FormatInt(int64(typedJson.Value.(int)), 10)
	case INT8:
		if _, ok := typedJson.Value.(int8); !ok {
			return nil, fmt.Errorf("failed to cast '%v' to an int8", typedJson.Value)
		}

		temp.Value = strconv.FormatInt(int64(typedJson.Value.(int8)), 10)
	case INT16:
		if _, ok := typedJson.Value.(int16); !ok {
			return nil, fmt.Errorf("failed to cast '%v' to an int16", typedJson.Value)
		}

		temp.Value = strconv.FormatInt(int64(typedJson.Value.(int16)), 10)
	case INT32:
		if _, ok := typedJson.Value.(int32); !ok {
			return nil, fmt.Errorf("failed to cast '%v' to an int32", typedJson.Value)
		}

		temp.Value = strconv.FormatInt(int64(typedJson.Value.(int32)), 10)
	case INT64:
		if _, ok := typedJson.Value.(int64); !ok {
			return nil, fmt.Errorf("failed to cast '%v' to an int64", typedJson.Value)
		}

		temp.Value = strconv.FormatInt(int64(typedJson.Value.(int64)), 10)
	case UINT:
		if _, ok := typedJson.Value.(uint); !ok {
			return nil, fmt.Errorf("failed to cast '%v' to an uint", typedJson.Value)
		}

		temp.Value = strconv.FormatUint(uint64(typedJson.Value.(uint)), 10)
	case UINT8:
		if _, ok := typedJson.Value.(uint8); !ok {
			return nil, fmt.Errorf("failed to cast '%v' to an uint8", typedJson.Value)
		}

		temp.Value = strconv.FormatUint(uint64(typedJson.Value.(uint8)), 10)
	case UINT16:
		if _, ok := typedJson.Value.(uint16); !ok {
			return nil, fmt.Errorf("failed to cast '%v' to an uint16", typedJson.Value)
		}

		temp.Value = strconv.FormatUint(uint64(typedJson.Value.(uint16)), 10)
	case UINT32:
		if _, ok := typedJson.Value.(uint32); !ok {
			return nil, fmt.Errorf("failed to cast '%v' to an uint32", typedJson.Value)
		}

		temp.Value = strconv.FormatUint(uint64(typedJson.Value.(uint32)), 10)
	case UINT64:
		if _, ok := typedJson.Value.(uint64); !ok {
			return nil, fmt.Errorf("failed to cast '%v' to an uint64", typedJson.Value)
		}

		temp.Value = strconv.FormatUint(uint64(typedJson.Value.(uint64)), 10)
	case FLOAT32:
		if _, ok := typedJson.Value.(float32); !ok {
			return nil, fmt.Errorf("failed to cast '%v' to a float32", typedJson.Value)
		}

		temp.Value = strconv.FormatFloat(float64(typedJson.Value.(float32)), 'E', -1, 32)
	case FLOAT64:
		if _, ok := typedJson.Value.(float64); !ok {
			return nil, fmt.Errorf("failed to cast '%v' to a float64", typedJson.Value)
		}

		temp.Value = strconv.FormatFloat(float64(typedJson.Value.(float64)), 'E', -1, 64)
	case STRING:
		if _, ok := typedJson.Value.(string); !ok {
			return nil, fmt.Errorf("failed to cast '%v' to a string", typedJson.Value)
		}

		temp.Value = typedJson.Value.(string)
	case BOOL:
		if _, ok := typedJson.Value.(bool); !ok {
			return nil, fmt.Errorf("failed to cast '%v' to a bool", typedJson.Value)
		}

		temp.Value = strconv.FormatBool(typedJson.Value.(bool))
	default:
		return nil, fmt.Errorf("unknow type to encode %d for %v", temp.Type, typedJson.Value)
	}

	return json.Marshal(temp)
}
