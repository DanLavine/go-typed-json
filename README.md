# go-typed-json
---------------
[godoc](https://pkg.go.dev/github.com/DanLavine/go-typed-json)


Go Typed Json provides a simple stricture to pass arbitrary types of data through JSON. By default,
golang will treat all interface numbers as float64, but sometimes you might want to specify that
they should be parsed as int64 or uint32 instead. This package provides a simple class to define
how arbitrary data fields should be encoded and decoded over JSON.

To do this in a language agnostic way, the model is defined as:
```
type TypedJson struct {
	// Type defines how to encode/decode the value
	Type JSONTYPE `json:"Type"`

	// Values with the type associated for Type
	Value any `json:"Value"`
}
```

When encoding any of the JSONTYPE's Values, we convert them to a string. This way we can preserve the object's
raw data and trust the common JSON specification for how to treat string values when sending them. Then
each Decode operation needs to parse the string value into its raw data type.

| JSONTYPE         | Value               | Details                                                                                            |
|:--               | :--                 | :--                                                                                                |
| INT              | "_int"              | signed 8 bytes                                                                                     |
| INT8             | "_int8"             | signed 1 byte                                                                                      |
| INT16            | "_int16"            | signed 2 bytes                                                                                     |
| INT32            | "_int32"            | signed 4 bytes                                                                                     |
| INT64            | "_int64"            | signed 8 bytes                                                                                     |
| UINT             | "_uint"             | unsigned 8 bytes                                                                                   |
| UINT8            | "_uint8"            | unsigned 1 byte                                                                                    |
| UINT16           | "_uint16"           | unsigned 2 bytes                                                                                   |
| UINT32           | "_uint32"           | unsigned 4 bytes                                                                                   |
| UINT64           | "_uint64"           | unsigned 8 bytes                                                                                   |
| FLOAT32          | "_float32"          | (IEEE 754 32-bit floating-point numbers)                                                           |
| FLOAT64          | "_float64"          | (IEEE 754 64-bit floating-point numbers)                                                           |
| STRING           | "_string"           | utf8 encoded string                                                                                |
| BOOL             | "_bool"             | `true` or `false`                                                                                  |
| DATETIME         | "_datetime"         | RFC3339 string encoded datetime                                                                    |
| DURATION         | "_duration"         | string format https://pkg.go.dev/time#ParseDuration                                                |
| COMPLEX64        | "_complex64"        | string of complex numbers https://go.dev/play/p/hyTDeE84G5y                                        |
| COMPLEX128       | "_complex128"       | string of complex numbers https://go.dev/play/p/JohwkAq58BE                                        |
| INT_SLICE        | "_int_slice"        | list of `,` seperated strings of type: signed 8 bytes                                              |
| INT8_SLICE       | "_int8_slice"       | list of `,` seperated strings of type: signed 1 byte                                               |
| INT16_SLICE      | "_int16_slice"      | list of `,` seperated strings of type: signed 2 bytes                                              |
| INT32_SLICE      | "_int32_slice"      | list of `,` seperated strings of type: signed 4 bytes                                              |
| INT64_SLICE      | "_int64_slice"      | list of `,` seperated strings of type: signed 8 bytes                                              |
| UINT_SLICE       | "_uint_slice"       | list of `,` seperated strings of type: unsigned 8 bytes                                            |
| UINT8_SLICE      | "_uint8_slice"      | list of `,` seperated strings of type: unsigned 1 byte                                             |
| UINT16_SLICE     | "_uint16_slice"     | list of `,` seperated strings of type: unsigned 2 bytes                                            |
| UINT32_SLICE     | "_uint32_slice"     | list of `,` seperated strings of type: unsigned 4 bytes                                            |
| UINT64_SLICE     | "_uint64_slice"     | list of `,` seperated strings of type: unsigned 8 bytes                                            |
| FLOAT32_SLICE    | "_float32_slice"    | list of `,` seperated strings of type: (IEEE 754 32-bit floating-point numbers)                    |
| FLOAT64_SLICE    | "_float64_slice"    | list of `,` seperated strings of type: (IEEE 754 64-bit floating-point numbers)                    |
| STRING_SLICE     | "_string_slice"     | list of `,` seperated strings of type: base64 encoded utf8 string                                         |
| BOOL_SLICE       | "_bool_slice"       | list of `,` seperated strings of type: `true` or `false`                                           |
| DATETIME_SLICE   | "_datetime_slice"   | list of `,` seperated strings of type: RFC3339 string encoded datetime                             |
| DURATION_SLICE   | "_duration_slice"   | list of `,` seperated strings of type: string format https://pkg.go.dev/time#ParseDuration         |
| COMPLEX64_SLICE  | "_complex64_slice"  | list of `,` seperated strings of type: string of complex numbers https://go.dev/play/p/9Dz12hWk8yp |
| COMPLEX128_SLICE | "_complex128_slice" | list of `,` seperated strings of type: string of complex numbers https://go.dev/play/p/I0CpIXk5O32 |

## Adding your own data types

There are a few ways to add your own encoding and decoding rules for custom types. The first and easiest way
is to use the [Global Codec](#global-codec) if all possible Marhalers and Unmarshalers are going to use the same
set of rules. Otherwise, using a [Custom Codec](#custom-coded) for individual `TypedJson` values will be easiest. This package does
not provide specific rules for tagging `any or interface{}` types on how they should be encoded/decoded since its main
goal is to accept arbitrary structures of unknow data.

NOTE:
Go-Typed-JSON reserves all key words begining with an `_`. This allows the packages to add any addition built in data
types, all prefiexed with a `_`.

#### Global Codec

To override the global coded that is used by anyone that imports the same package, you can update the shared var
```
var GlobalCodec CustomCodec = nil
```

#### Custom Coded

When using a known data types and structures that have specific rules for the model you are working with, you can instantiate objects
with the `NewTypedJson` function. This allows for custome encoders to be attacked to `TypedJson` structures
```
func NewTypedJson(jsonType JSONTYPE, value any, customCodec CustomCodec) *TypedJson {
	...
}
```