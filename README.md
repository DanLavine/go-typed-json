# go-typed-json
---------------

Go Typed Json provides a simple stricture to pass arbitrary types of data through JSON. By default
golang will treat all interface numbers as float64, but sometimes you might want to specify that
they should be parsed as int64 or uint32 instead. This package provides a simple class to define
how interface types should be encoded and decoded over json. 