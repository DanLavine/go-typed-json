# go-typed-json
---------------
[godoc](https://pkg.go.dev/github.com/DanLavine/go-typed-json)


Go Typed Json provides a simple stricture to pass arbitrary types of data through JSON. By default
golang will treat all interface numbers as float64, but sometimes you might want to specify that
they should be parsed as int64 or uint32 instead. This package provides a simple class to define
how interface types should be encoded and decoded over JSON.

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
raw data and trust the common JSON specification for how to treat string values when sending or saving them. Then
each Decode operation needs to parse the string value into its raw data type.


| JSONTYPE | Value |
|:-- | :-- |
| INT      | 1 |
| INT8     | 2 |
| INT16    | 3 |
| INT32    | 4 |
| INT64    | 5 |
| UINT     | 6 |
| UINT8    | 7 |
| UINT16   | 8 |
| UINT32   | 9 |
| UINT64   | 10 |
| FLOAT32  | 11 |
| FLOAT64  | 12 |
| STRING   | 13 |
| BOOL     | 14 |

NOTE that there is no 0 JSONTYPE. This is because objects initialization state is 0 and we don't have a way of 
knowing unset vs JSONTYPE 0.

If you want to define your own data types, I suggest starting at a much larger index `256`, just in case there are
more data types added later, we don't conflict with each other
