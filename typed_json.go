package gotypedjson

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// GlobalCodec is by default an unset codec that will be used to encode and decode all TypedJson
// structs. This is usefult to set if all models follow the same encoding and decoding rules.
var GlobalCodec CustomCodec = nil

// JSONTYPE is the custom defintion used when determining the encoding / decoding structures
type JSONTYPE int

const (
	INT        JSONTYPE = 1
	INT8       JSONTYPE = 2
	INT16      JSONTYPE = 3
	INT32      JSONTYPE = 4
	INT64      JSONTYPE = 5
	UINT       JSONTYPE = 6
	UINT8      JSONTYPE = 7
	UINT16     JSONTYPE = 8
	UINT32     JSONTYPE = 9
	UINT64     JSONTYPE = 10
	FLOAT32    JSONTYPE = 11
	FLOAT64    JSONTYPE = 12
	STRING     JSONTYPE = 13
	BOOL       JSONTYPE = 14
	COMPLEX64  JSONTYPE = 15
	COMPLEX128 JSONTYPE = 16
)

// Codec are used to Encode and Decode JSONTYPE data
type Codec struct {
	// Encoded the data into a string for data integrity
	Encode func(val any) (string, error)
	// Decode an encoded string back into its original valie
	Decode func(s string) (any, error)
}

// CustomCodec are used to associate specific types with their encoding and decoding functions
type CustomCodec map[JSONTYPE]Codec

// TypedJson define the specifc Type of JSON Value, dictaing how to encode and decode the value. By default, all values
// are encoded and decoded as strings to ensure data consistency when converting between types.
//
// For Encoding and Decoding the struct, we use the following codec priority:
//  1. customCodec - internal field on this struct that is optional
//  2. GlobalCodec - global codec that is used for all typed json structs
//  3. default     - the default encoding and decoding for this package
type TypedJson struct {
	// Type defines how to encode/decode the value
	Type JSONTYPE `json:"Type"`

	// Values with the type associated for Type
	Value any `json:"Value"`

	// codec for just this struct
	customCodec CustomCodec
}

//	PARAMETERS:
//	* jsonType    - The Type associated with the Value
//	* value       - Value that can be encoded and decoded consistently
//	* customCodec - (optional) codec that can be used for custom types, nil will just use the global and then default codec
//
//	RETURNS:
//	* *TypedJson - json object that can encode/decode typed json
//
// Returns an initalized TypedJson with the optional typed custom codec set. This can panic if the custom codec is
// missing an encode or decode function for any of the defined types. .
func NewTypedJson(jsonType JSONTYPE, value any, customCodec CustomCodec) *TypedJson {
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
		Value:       value,
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

	// try the custom codec types
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

	// try the global codec types
	if GlobalCodec != nil {
		if encoder, ok := GlobalCodec[temp.Type]; ok {
			val, err := encoder.Decode(temp.Value)
			if err != nil {
				return err
			}

			typedJson.Value = val

			return nil
		}
	}

	// try the default codec types
	switch temp.Type {
	case INT:
		val, err := strconv.ParseInt(temp.Value, 10, 0)
		if err != nil {
			return fmt.Errorf("failed to convert '%s' to an int", temp.Value)
		}
		typedJson.Value = int(val)
	case INT8:
		val, err := strconv.ParseInt(temp.Value, 10, 8)
		if err != nil {
			return fmt.Errorf("failed to convert '%s' to an int8", temp.Value)
		}
		typedJson.Value = int8(val)
	case INT16:
		val, err := strconv.ParseInt(temp.Value, 10, 16)
		if err != nil {
			return fmt.Errorf("failed to convert '%s' to an int16", temp.Value)
		}
		typedJson.Value = int16(val)
	case INT32:
		val, err := strconv.ParseInt(temp.Value, 10, 32)
		if err != nil {
			return fmt.Errorf("failed to convert '%s' to an int32", temp.Value)
		}
		typedJson.Value = int32(val)
	case INT64:
		val, err := strconv.ParseInt(temp.Value, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to convert '%s' to an int64", temp.Value)
		}
		typedJson.Value = int64(val)
	case UINT:
		val, err := strconv.ParseUint(temp.Value, 10, 0)
		if err != nil {
			return fmt.Errorf("failed to convert '%s' to a uint", temp.Value)
		}
		typedJson.Value = uint(val)
	case UINT8:
		val, err := strconv.ParseUint(temp.Value, 10, 8)
		if err != nil {
			return fmt.Errorf("failed to convert '%s' to a uint8", temp.Value)
		}
		typedJson.Value = uint8(val)
	case UINT16:
		val, err := strconv.ParseUint(temp.Value, 10, 16)
		if err != nil {
			return fmt.Errorf("failed to convert '%s' to a uint16", temp.Value)
		}
		typedJson.Value = uint16(val)
	case UINT32:
		val, err := strconv.ParseUint(temp.Value, 10, 32)
		if err != nil {
			return fmt.Errorf("failed to convert '%s' to a uint32", temp.Value)
		}
		typedJson.Value = uint32(val)
	case UINT64:
		val, err := strconv.ParseUint(temp.Value, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to convert '%s' to a uint64", temp.Value)
		}
		typedJson.Value = uint64(val)
	case FLOAT32:
		val, err := strconv.ParseFloat(temp.Value, 32)
		if err != nil {
			return fmt.Errorf("failed to convert '%s' to a float32", temp.Value)
		}
		typedJson.Value = float32(val)
	case FLOAT64:
		val, err := strconv.ParseFloat(temp.Value, 64)
		if err != nil {
			return fmt.Errorf("failed to convert '%s' to a float64", temp.Value)
		}
		typedJson.Value = float64(val)
	case STRING:
		typedJson.Value = string(temp.Value)
	case BOOL:
		val, err := strconv.ParseBool(temp.Value)
		if err != nil {
			return fmt.Errorf("failed to convert '%s' to a bool", temp.Value)
		}
		typedJson.Value = bool(val)
	case COMPLEX64:
		val, err := strconv.ParseComplex(temp.Value, 64)
		if err != nil {
			return fmt.Errorf("failed to convert '%s' to a complex64", temp.Value)
		}
		typedJson.Value = complex64(val)
	case COMPLEX128:
		val, err := strconv.ParseComplex(temp.Value, 128)
		if err != nil {
			return fmt.Errorf("failed to convert '%s' to a complex128", temp.Value)
		}
		typedJson.Value = val
	default:
		return fmt.Errorf("unknown type '%d' recevied for: %s", temp.Type, temp.Value)
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

	// try the global types
	if GlobalCodec != nil {
		if encoder, ok := GlobalCodec[typedJson.Type]; ok {
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
	case COMPLEX64:
		if _, ok := typedJson.Value.(complex64); !ok {
			return nil, fmt.Errorf("failed to cast '%v' to a complex64", typedJson.Value)
		}

		temp.Value = strconv.FormatComplex(complex128(typedJson.Value.(complex64)), 'E', -1, 64)
	case COMPLEX128:
		if _, ok := typedJson.Value.(complex128); !ok {
			return nil, fmt.Errorf("failed to cast '%v' to a complex128", typedJson.Value)
		}

		temp.Value = strconv.FormatComplex(typedJson.Value.(complex128), 'E', -1, 64)
	default:
		return nil, fmt.Errorf("unknow type to encode %d for value: %v", temp.Type, typedJson.Value)
	}

	return json.Marshal(temp)
}
