package gotypedjson

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type JSONTYPE int

type CustomEncoder map[JSONTYPE]Encoder

type Encoder struct {
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

	customEncoder CustomEncoder
}

func NewTypedJson(jsonType JSONTYPE, Value any, customEncoder CustomEncoder) *TypedJson {
	return &TypedJson{
		Type:          jsonType,
		Value:         Value,
		customEncoder: customEncoder,
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
			return fmt.Errorf("failed to convert %s, to a uint", temp.Value)
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
		typedJson.Value = temp.Value
	case BOOL:
		val, err := strconv.ParseBool(temp.Value)
		if err != nil {
			return fmt.Errorf("failed to convert %s, to a bool", temp.Value)
		}
		typedJson.Value = bool(val)
	default:
		if typedJson.customEncoder != nil {
			if encoder, ok := typedJson.customEncoder[temp.Type]; ok {
				val, err := encoder.Decode(temp.Value)
				if err != nil {
					return err
				}

				typedJson.Value = val
			} else {
				return fmt.Errorf("unknown type %d recevied for: %s", temp.Type, temp.Value)
			}
		} else {
			return fmt.Errorf("unknown type %d recevied for: %s", temp.Type, temp.Value)
		}
	}

	typedJson.Type = temp.Type

	return nil
}

func (typedJson *TypedJson) MarshalJSON() ([]byte, error) {
	temp := struct {
		Type  JSONTYPE `json:"Type"`
		Value string   `json:"Value"`
	}{
		Type: typedJson.Type,
	}

	switch typedJson.Type {
	case INT:
		temp.Value = strconv.FormatInt(int64(typedJson.Value.(int)), 10)
	case INT8:
		temp.Value = strconv.FormatInt(int64(typedJson.Value.(int8)), 10)
	case INT16:
		temp.Value = strconv.FormatInt(int64(typedJson.Value.(int16)), 10)
	case INT32:
		temp.Value = strconv.FormatInt(int64(typedJson.Value.(int32)), 10)
	case INT64:
		temp.Value = strconv.FormatInt(int64(typedJson.Value.(int64)), 10)
	case UINT:
		temp.Value = strconv.FormatUint(uint64(typedJson.Value.(uint)), 10)
	case UINT8:
		temp.Value = strconv.FormatUint(uint64(typedJson.Value.(uint8)), 10)
	case UINT16:
		temp.Value = strconv.FormatUint(uint64(typedJson.Value.(uint16)), 10)
	case UINT32:
		temp.Value = strconv.FormatUint(uint64(typedJson.Value.(uint32)), 10)
	case UINT64:
		temp.Value = strconv.FormatUint(uint64(typedJson.Value.(uint64)), 10)
	case FLOAT32:
		temp.Value = strconv.FormatFloat(float64(typedJson.Value.(float32)), 'E', -1, 32)
	case FLOAT64:
		temp.Value = strconv.FormatFloat(float64(typedJson.Value.(float32)), 'E', -1, 64)
	case STRING:
		temp.Value = typedJson.Value.(string)
	case BOOL:
		temp.Value = strconv.FormatBool(typedJson.Value.(bool))
	default:
		// might be a custom type
		if typedJson.customEncoder != nil {
			if encoder, ok := typedJson.customEncoder[typedJson.Type]; ok {
				assignString, err := encoder.Encode(typedJson.Value)
				if err != nil {
					return nil, err
				}

				temp.Value = assignString
			} else {
				return nil, fmt.Errorf("unknow type to encode %d for %v", temp.Type, typedJson.Value)
			}
		} else {
			return nil, fmt.Errorf("unknow type to encode %d for %v", temp.Type, typedJson.Value)
		}
	}

	return json.Marshal(temp)
}
