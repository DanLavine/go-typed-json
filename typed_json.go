package gotypedjson

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
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

	INT_SLICE        JSONTYPE = 101
	INT8_SLICE       JSONTYPE = 102
	INT16_SLICE      JSONTYPE = 103
	INT32_SLICE      JSONTYPE = 104
	INT64_SLICE      JSONTYPE = 105
	UINT_SLICE       JSONTYPE = 106
	UINT8_SLICE      JSONTYPE = 107
	UINT16_SLICE     JSONTYPE = 108
	UINT32_SLICE     JSONTYPE = 109
	UINT64_SLICE     JSONTYPE = 110
	FLOAT32_SLICE    JSONTYPE = 111
	FLOAT64_SLICE    JSONTYPE = 112
	STRING_SLICE     JSONTYPE = 113
	BOOL_SLICE       JSONTYPE = 114
	COMPLEX64_SLICE  JSONTYPE = 115
	COMPLEX128_SLICE JSONTYPE = 116
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
	case INT_SLICE:
		tmp := []int{}
		if temp.Value != "" {
			for _, value := range strings.Split(temp.Value, ",") {
				decodedValue, err := base64.StdEncoding.DecodeString(value)
				if err != nil {
					return fmt.Errorf("string '%s' is not an expected base64", value)
				}

				val, err := strconv.ParseInt(string(decodedValue), 10, 0)
				if err != nil {
					return fmt.Errorf("failed to convert '%s' to an int", string(decodedValue))
				}

				tmp = append(tmp, int(val))
			}

			typedJson.Value = tmp
		}

		typedJson.Value = tmp

	case INT8_SLICE:
		tmp := []int8{}
		if temp.Value != "" {
			for _, value := range strings.Split(temp.Value, ",") {
				decodedValue, err := base64.StdEncoding.DecodeString(value)
				if err != nil {
					return fmt.Errorf("string '%s' is not an expected base64", value)
				}

				val, err := strconv.ParseInt(string(decodedValue), 10, 8)
				if err != nil {
					return fmt.Errorf("failed to convert '%s' to an int8", string(decodedValue))
				}

				tmp = append(tmp, int8(val))
			}
		}

		typedJson.Value = tmp
	case INT16_SLICE:
		tmp := []int16{}
		if temp.Value != "" {
			for _, value := range strings.Split(temp.Value, ",") {
				decodedValue, err := base64.StdEncoding.DecodeString(value)
				if err != nil {
					return fmt.Errorf("string '%s' is not an expected base64", value)
				}

				val, err := strconv.ParseInt(string(decodedValue), 10, 16)
				if err != nil {
					return fmt.Errorf("failed to convert '%s' to an int16", string(decodedValue))
				}

				tmp = append(tmp, int16(val))
			}
		}

		typedJson.Value = tmp
	case INT32_SLICE:
		tmp := []int32{}
		if temp.Value != "" {
			for _, value := range strings.Split(temp.Value, ",") {
				decodedValue, err := base64.StdEncoding.DecodeString(value)
				if err != nil {
					return fmt.Errorf("string '%s' is not an expected base64", value)
				}

				val, err := strconv.ParseInt(string(decodedValue), 10, 32)
				if err != nil {
					return fmt.Errorf("failed to convert '%s' to an int32", string(decodedValue))
				}

				tmp = append(tmp, int32(val))
			}
		}

		typedJson.Value = tmp
	case INT64_SLICE:
		tmp := []int64{}
		if temp.Value != "" {
			for _, value := range strings.Split(temp.Value, ",") {
				decodedValue, err := base64.StdEncoding.DecodeString(value)
				if err != nil {
					return fmt.Errorf("string '%s' is not an expected base64", value)
				}

				val, err := strconv.ParseInt(string(decodedValue), 10, 8)
				if err != nil {
					return fmt.Errorf("failed to convert '%s' to an int64", string(decodedValue))
				}

				tmp = append(tmp, int64(val))
			}
		}

		typedJson.Value = tmp
	case UINT_SLICE:
		tmp := []uint{}
		if temp.Value != "" {
			for _, value := range strings.Split(temp.Value, ",") {
				decodedValue, err := base64.StdEncoding.DecodeString(value)
				if err != nil {
					return fmt.Errorf("string '%s' is not an expected base64", value)
				}

				val, err := strconv.ParseUint(string(decodedValue), 10, 0)
				if err != nil {
					return fmt.Errorf("failed to convert '%s' to a uint", string(decodedValue))
				}

				tmp = append(tmp, uint(val))
			}
		}

		typedJson.Value = tmp
	case UINT8_SLICE:
		tmp := []uint8{}
		if temp.Value != "" {
			for _, value := range strings.Split(temp.Value, ",") {
				decodedValue, err := base64.StdEncoding.DecodeString(value)
				if err != nil {
					return fmt.Errorf("string '%s' is not an expected base64", value)
				}

				val, err := strconv.ParseUint(string(decodedValue), 10, 8)
				if err != nil {
					return fmt.Errorf("failed to convert '%s' to a uint8", string(decodedValue))
				}

				tmp = append(tmp, uint8(val))
			}
		}

		typedJson.Value = tmp
	case UINT16_SLICE:
		tmp := []uint16{}
		if temp.Value != "" {
			for _, value := range strings.Split(temp.Value, ",") {
				decodedValue, err := base64.StdEncoding.DecodeString(value)
				if err != nil {
					return fmt.Errorf("string '%s' is not an expected base64", value)
				}

				val, err := strconv.ParseUint(string(decodedValue), 10, 16)
				if err != nil {
					return fmt.Errorf("failed to convert '%s' to a uint16", string(decodedValue))
				}

				tmp = append(tmp, uint16(val))
			}
		}

		typedJson.Value = tmp
	case UINT32_SLICE:
		tmp := []uint32{}
		if temp.Value != "" {
			for _, value := range strings.Split(temp.Value, ",") {
				decodedValue, err := base64.StdEncoding.DecodeString(value)
				if err != nil {
					return fmt.Errorf("string '%s' is not an expected base64", value)
				}

				val, err := strconv.ParseUint(string(decodedValue), 10, 32)
				if err != nil {
					return fmt.Errorf("failed to convert '%s' to a uint32", string(decodedValue))
				}

				tmp = append(tmp, uint32(val))
			}
		}

		typedJson.Value = tmp
	case UINT64_SLICE:
		tmp := []uint64{}
		if temp.Value != "" {
			for _, value := range strings.Split(temp.Value, ",") {
				decodedValue, err := base64.StdEncoding.DecodeString(value)
				if err != nil {
					return fmt.Errorf("string '%s' is not an expected base64", value)
				}

				val, err := strconv.ParseUint(string(decodedValue), 10, 64)
				if err != nil {
					return fmt.Errorf("failed to convert '%s' to a uint64", string(decodedValue))
				}

				tmp = append(tmp, uint64(val))
			}
		}

		typedJson.Value = tmp
	case FLOAT32_SLICE:
		tmp := []float32{}
		if temp.Value != "" {
			for _, value := range strings.Split(temp.Value, ",") {
				decodedValue, err := base64.StdEncoding.DecodeString(value)
				if err != nil {
					return fmt.Errorf("string '%s' is not an expected base64", value)
				}

				val, err := strconv.ParseFloat(string(decodedValue), 32)
				if err != nil {
					return fmt.Errorf("failed to convert '%s' to a float32", string(decodedValue))
				}

				tmp = append(tmp, float32(val))
			}
		}

		typedJson.Value = tmp
	case FLOAT64_SLICE:
		tmp := []float64{}
		if temp.Value != "" {
			for _, value := range strings.Split(temp.Value, ",") {
				decodedValue, err := base64.StdEncoding.DecodeString(value)
				if err != nil {
					return fmt.Errorf("string '%s' is not an expected base64", value)
				}

				val, err := strconv.ParseFloat(string(decodedValue), 64)
				if err != nil {
					return fmt.Errorf("failed to convert '%s' to a float64", string(decodedValue))
				}

				tmp = append(tmp, float64(val))
			}
		}

		typedJson.Value = tmp
	case STRING_SLICE:
		tmp := []string{}
		if temp.Value != "" {
			for _, value := range strings.Split(temp.Value, ",") {
				decodedValue, err := base64.StdEncoding.DecodeString(value)
				if err != nil {
					return fmt.Errorf("string '%s' is not an expected base64", value)
				}

				tmp = append(tmp, string(decodedValue))
			}
		}

		typedJson.Value = tmp
	case BOOL_SLICE:
		tmp := []bool{}
		if temp.Value != "" {
			for _, value := range strings.Split(temp.Value, ",") {
				decodedValue, err := base64.StdEncoding.DecodeString(value)
				if err != nil {
					return fmt.Errorf("string '%s' is not an expected base64", value)
				}

				val, err := strconv.ParseBool(string(decodedValue))
				if err != nil {
					return fmt.Errorf("failed to convert '%s' to a bool", string(decodedValue))
				}

				tmp = append(tmp, val)
			}
		}

		typedJson.Value = tmp
	case COMPLEX64_SLICE:
		tmp := []complex64{}
		if temp.Value != "" {
			for _, value := range strings.Split(temp.Value, ",") {
				decodedValue, err := base64.StdEncoding.DecodeString(value)
				if err != nil {
					return fmt.Errorf("string '%s' is not an expected base64", value)
				}

				val, err := strconv.ParseComplex(string(decodedValue), 64)
				if err != nil {
					return fmt.Errorf("failed to convert '%s' to a complex64", string(decodedValue))
				}

				tmp = append(tmp, complex64(val))
			}
		}

		typedJson.Value = tmp
	case COMPLEX128_SLICE:
		tmp := []complex128{}
		if temp.Value != "" {
			for _, value := range strings.Split(temp.Value, ",") {
				decodedValue, err := base64.StdEncoding.DecodeString(value)
				if err != nil {
					return fmt.Errorf("string '%s' is not an expected base64", value)
				}

				val, err := strconv.ParseComplex(string(decodedValue), 128)
				if err != nil {
					return fmt.Errorf("failed to convert '%s' to a complex128", string(decodedValue))
				}

				tmp = append(tmp, val)
			}
		}

		typedJson.Value = tmp
	default:
		return fmt.Errorf("unknown type '%d' to decode", temp.Type)
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
	case INT_SLICE:
		if typedJson.Value == nil {
			temp.Value = ""
		} else {
			if values, ok := typedJson.Value.([]int); ok {
				for index, value := range values {
					str := base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(int64(value), 10)))
					if index == 0 {
						temp.Value = str
					} else {
						temp.Value += "," + str
					}
				}
			} else {
				return nil, fmt.Errorf("failed to cast '%v' to a []int", typedJson.Value)
			}
		}
	case INT8_SLICE:
		if typedJson.Value == nil {
			temp.Value = ""
		} else {
			if values, ok := typedJson.Value.([]int8); ok {
				for index, value := range values {
					str := base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(int64(value), 10)))
					if index == 0 {
						temp.Value = str
					} else {
						temp.Value += "," + str
					}
				}
			} else {
				return nil, fmt.Errorf("failed to cast '%v' to a []int8", typedJson.Value)
			}
		}
	case INT16_SLICE:
		if typedJson.Value == nil {
			temp.Value = ""
		} else {
			if values, ok := typedJson.Value.([]int16); ok {
				for index, value := range values {
					str := base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(int64(value), 10)))
					if index == 0 {
						temp.Value = str
					} else {
						temp.Value += "," + str
					}
				}
			} else {
				return nil, fmt.Errorf("failed to cast '%v' to a []int16", typedJson.Value)
			}
		}
	case INT32_SLICE:
		if typedJson.Value == nil {
			temp.Value = ""
		} else {
			if values, ok := typedJson.Value.([]int32); ok {
				for index, value := range values {
					str := base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(int64(value), 10)))
					if index == 0 {
						temp.Value = str
					} else {
						temp.Value += "," + str
					}
				}
			} else {
				return nil, fmt.Errorf("failed to cast '%v' to a []int32", typedJson.Value)
			}
		}
	case INT64_SLICE:
		if typedJson.Value == nil {
			temp.Value = ""
		} else {
			if values, ok := typedJson.Value.([]int64); ok {
				for index, value := range values {
					str := base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(value, 10)))
					if index == 0 {
						temp.Value = str
					} else {
						temp.Value += "," + str
					}
				}
			} else {
				return nil, fmt.Errorf("failed to cast '%v' to a []int64", typedJson.Value)
			}
		}
	case UINT_SLICE:
		if typedJson.Value == nil {
			temp.Value = ""
		} else {
			if values, ok := typedJson.Value.([]uint); ok {
				for index, value := range values {
					str := base64.StdEncoding.EncodeToString([]byte(strconv.FormatUint(uint64(value), 10)))
					if index == 0 {
						temp.Value = str
					} else {
						temp.Value += "," + str
					}
				}
			} else {
				return nil, fmt.Errorf("failed to cast '%v' to a []uint", typedJson.Value)
			}
		}
	case UINT8_SLICE:
		if typedJson.Value == nil {
			temp.Value = ""
		} else {
			if values, ok := typedJson.Value.([]uint8); ok {
				for index, value := range values {
					str := base64.StdEncoding.EncodeToString([]byte(strconv.FormatUint(uint64(value), 10)))
					if index == 0 {
						temp.Value = str
					} else {
						temp.Value += "," + str
					}
				}
			} else {
				return nil, fmt.Errorf("failed to cast '%v' to a []uint8", typedJson.Value)
			}
		}
	case UINT16_SLICE:
		if typedJson.Value == nil {
			temp.Value = ""
		} else {
			if values, ok := typedJson.Value.([]uint16); ok {
				for index, value := range values {
					str := base64.StdEncoding.EncodeToString([]byte(strconv.FormatUint(uint64(value), 10)))
					if index == 0 {
						temp.Value = str
					} else {
						temp.Value += "," + str
					}
				}
			} else {
				return nil, fmt.Errorf("failed to cast '%v' to a []uint16", typedJson.Value)
			}
		}
	case UINT32_SLICE:
		if typedJson.Value == nil {
			temp.Value = ""
		} else {
			if values, ok := typedJson.Value.([]uint32); ok {
				for index, value := range values {
					str := base64.StdEncoding.EncodeToString([]byte(strconv.FormatUint(uint64(value), 10)))
					if index == 0 {
						temp.Value = str
					} else {
						temp.Value += "," + str
					}
				}
			} else {
				return nil, fmt.Errorf("failed to cast '%v' to a []uint32", typedJson.Value)
			}
		}
	case UINT64_SLICE:
		if typedJson.Value == nil {
			temp.Value = ""
		} else {
			if values, ok := typedJson.Value.([]uint64); ok {
				for index, value := range values {
					str := base64.StdEncoding.EncodeToString([]byte(strconv.FormatUint(uint64(value), 10)))
					if index == 0 {
						temp.Value = str
					} else {
						temp.Value += "," + str
					}
				}
			} else {
				return nil, fmt.Errorf("failed to cast '%v' to a []uint64", typedJson.Value)
			}
		}
	case FLOAT32_SLICE:
		if typedJson.Value == nil {
			temp.Value = ""
		} else {
			if values, ok := typedJson.Value.([]float32); ok {
				for index, value := range values {
					str := base64.StdEncoding.EncodeToString([]byte(strconv.FormatFloat(float64(value), 'E', -1, 32)))
					if index == 0 {
						temp.Value = str
					} else {
						temp.Value += "," + str
					}
				}
			} else {
				return nil, fmt.Errorf("failed to cast '%v' to a []float32", typedJson.Value)
			}
		}
	case FLOAT64_SLICE:
		if typedJson.Value == nil {
			temp.Value = ""
		} else {
			if values, ok := typedJson.Value.([]float64); ok {
				for index, value := range values {
					str := base64.StdEncoding.EncodeToString([]byte(strconv.FormatFloat(value, 'E', -1, 64)))
					if index == 0 {
						temp.Value = str
					} else {
						temp.Value += "," + str
					}
				}
			} else {
				return nil, fmt.Errorf("failed to cast '%v' to a []float64", typedJson.Value)
			}
		}
	case STRING_SLICE:
		if typedJson.Value == nil {
			temp.Value = ""
		} else {
			if values, ok := typedJson.Value.([]string); ok {
				for index, value := range values {
					str := base64.StdEncoding.EncodeToString([]byte(value))
					if index == 0 {
						temp.Value = str
					} else {
						temp.Value += "," + str
					}
				}
			} else {
				return nil, fmt.Errorf("failed to cast '%v' to a []string", typedJson.Value)
			}
		}
	case BOOL_SLICE:
		if typedJson.Value == nil {
			temp.Value = ""
		} else {
			if values, ok := typedJson.Value.([]bool); ok {
				for index, value := range values {
					str := base64.StdEncoding.EncodeToString([]byte(strconv.FormatBool(value)))
					if index == 0 {
						temp.Value = str
					} else {
						temp.Value += "," + str
					}
				}
			} else {
				return nil, fmt.Errorf("failed to cast '%v' to a []bool", typedJson.Value)
			}
		}
	case COMPLEX64_SLICE:
		if typedJson.Value == nil {
			temp.Value = ""
		} else {
			if values, ok := typedJson.Value.([]complex64); ok {
				for index, value := range values {
					str := base64.StdEncoding.EncodeToString([]byte(strconv.FormatComplex(complex128(value), 'E', -1, 64)))
					if index == 0 {
						temp.Value = str
					} else {
						temp.Value += "," + str
					}
				}
			} else {
				return nil, fmt.Errorf("failed to cast '%v' to a []complex64", typedJson.Value)
			}
		}
	case COMPLEX128_SLICE:
		if typedJson.Value == nil {
			temp.Value = ""
		} else {
			if values, ok := typedJson.Value.([]complex128); ok {
				for index, value := range values {
					str := base64.StdEncoding.EncodeToString([]byte(strconv.FormatComplex(value, 'E', -1, 64)))
					if index == 0 {
						temp.Value = str
					} else {
						temp.Value += "," + str
					}
				}
			} else {
				return nil, fmt.Errorf("failed to cast '%v' to a []complex128", typedJson.Value)
			}
		}
	default:
		return nil, fmt.Errorf("unknow type '%d' to encode", temp.Type)
	}

	return json.Marshal(temp)
}
