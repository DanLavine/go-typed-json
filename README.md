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

and when encoding any of the JSONTYPE's Values, we convert them to a string. This way we can preserve the object's
raw data and trust the common JSON specification for how to treat string values when sending them. Then
each Decode operation needs to parse the string value into its raw data type.


| JSONTYPE     | Value |
|:--           | :-- |
| INT          | 1 |
| INT8         | 2 |
| INT16        | 3 |
| INT32        | 4 |
| INT64        | 5 |
| UINT         | 6 |
| UINT8        | 7 |
| UINT16       | 8 |
| UINT32       | 9 |
| UINT64       | 10 |
| FLOAT32      | 11 |
| FLOAT64      | 12 |
| STRING       | 13 |
| BOOL         | 14 |
| COMPLEX64    | 15 |
| COMPLEX128   | 16 |
| []INT        | 101 |
| []INT8       | 102 |
| []INT16      | 103 |
| []INT32      | 104 |
| []INT64      | 105 |
| []UINT       | 106 |
| []UINT8      | 107 |
| []UINT16     | 108 |
| []UINT32     | 109 |
| []UINT64     | 110 |
| []FLOAT32    | 111 |
| []FLOAT64    | 112 |
| []STRING     | 113 |
| []BOOL       | 114 |
| []COMPLEX64  | 115 |
| []COMPLEX128 | 116 |

NOTE that there is no 0 JSONTYPE. This is because objects initialization state is 0 and we don't have a way of 
knowing unset vs JSONTYPE 0.

## Adding your own data types

There are a few ways to add your own encoding and decoding rules for custom types. The first and easiest way
is to use the [Global Codec](#global-codec) if all possible Marhalers and Unmarshalers are going to use the same
set of rules. Otherwise, using a [Custom Codec](#custom-coded) for individual `TypedJson` values will be easiest. This package does
not provide specific rules for tagging `any or interface{}` types on how they should be encoded/decoded since its main
goal is to accept arbitrary structures of unknow data.

NOTE:
If you want to define your own data types, I suggest starting at a much larger index `1024`, just in case there are
more data types added later, we don't conflict with each other

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