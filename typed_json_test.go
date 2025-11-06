package gotypedjson_test

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"time"

	gotypedjson "github.com/DanLavine/go-typed-json"
	. "github.com/onsi/gomega"
)

func Test_NewTypedJson(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("It panics if a codec encode value is missing", func(t *testing.T) {
		codec := gotypedjson.CustomCodec{
			gotypedjson.BOOL: gotypedjson.Codec{
				Encode: func(val any) (string, error) { return "", nil },
			},
		}

		g.Expect(func() {
			gotypedjson.NewTypedJson(gotypedjson.INT, int(2), codec)
		}).To(Panic())
	})

	t.Run("It panics if a codec decode value is missing", func(t *testing.T) {
		codec := gotypedjson.CustomCodec{
			gotypedjson.BOOL: gotypedjson.Codec{
				Decode: func(s string) (any, error) { return nil, nil },
			},
		}

		g.Expect(func() {
			gotypedjson.NewTypedJson(gotypedjson.INT, int(2), codec)
		}).To(Panic())
	})
}

func Test_NewTypedJsonDecoder(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("It panics if a codec encode value is missing", func(t *testing.T) {
		codec := gotypedjson.CustomCodec{
			gotypedjson.BOOL: gotypedjson.Codec{
				Encode: func(val any) (string, error) { return "", nil },
			},
		}

		g.Expect(func() {
			gotypedjson.NewTypedJsonDecoder(codec)
		}).To(Panic())
	})

	t.Run("It panics if a codec decode value is missing", func(t *testing.T) {
		codec := gotypedjson.CustomCodec{
			gotypedjson.BOOL: gotypedjson.Codec{
				Decode: func(s string) (any, error) { return nil, nil },
			},
		}

		g.Expect(func() {
			gotypedjson.NewTypedJsonDecoder(codec)
		}).To(Panic())
	})
}

func Test_Int(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":"_int","Value":"4"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tInt := &gotypedjson.TypedJson{Type: gotypedjson.INT, Value: "nope"}
			data, err := json.Marshal(tInt)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to an int"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode the value properly", func(t *testing.T) {
			tInt := &gotypedjson.TypedJson{Type: gotypedjson.INT, Value: 4}

			data, err := json.Marshal(tInt)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawData))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an incorrect value", func(t *testing.T) {
			tInt := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_int","Value":"nope"}`), tInt)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'nope' to an int"))
		})

		t.Run("It can decode the value properly", func(t *testing.T) {
			tInt := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawData), tInt)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tInt.Type).To(Equal(gotypedjson.INT))
			g.Expect(tInt.Value.(int)).To(Equal(4))
		})
	})
}

func Test_Int8(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":"_int8","Value":"4"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tInt8 := &gotypedjson.TypedJson{Type: gotypedjson.INT8, Value: "nope"}
			data, err := json.Marshal(tInt8)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to an int8"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode the value properly", func(t *testing.T) {
			tInt8 := &gotypedjson.TypedJson{Type: gotypedjson.INT8, Value: int8(4)}

			data, err := json.Marshal(tInt8)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawData))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an incorrect value", func(t *testing.T) {
			tInt8 := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_int8","Value":"nope"}`), tInt8)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'nope' to an int8"))
		})

		t.Run("It can decode the value properly", func(t *testing.T) {
			tInt8 := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawData), tInt8)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tInt8.Type).To(Equal(gotypedjson.INT8))
			g.Expect(tInt8.Value.(int8)).To(Equal(int8(4)))
		})
	})
}

func Test_Int16(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":"_int16","Value":"4"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tInt16 := &gotypedjson.TypedJson{Type: gotypedjson.INT16, Value: "nope"}
			data, err := json.Marshal(tInt16)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to an int16"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode the value properly", func(t *testing.T) {
			tInt16 := &gotypedjson.TypedJson{Type: gotypedjson.INT16, Value: int16(4)}

			data, err := json.Marshal(tInt16)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawData))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an incorrect value", func(t *testing.T) {
			tInt16 := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_int16","Value":"nope"}`), tInt16)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'nope' to an int16"))
		})

		t.Run("It can decode the value properly", func(t *testing.T) {
			tInt16 := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawData), tInt16)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tInt16.Type).To(Equal(gotypedjson.INT16))
			g.Expect(tInt16.Value.(int16)).To(Equal(int16(4)))
		})
	})
}

func Test_Int32(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":"_int32","Value":"4"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tInt32 := &gotypedjson.TypedJson{Type: gotypedjson.INT32, Value: "nope"}
			data, err := json.Marshal(tInt32)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to an int32"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode the value properly", func(t *testing.T) {
			tInt32 := &gotypedjson.TypedJson{Type: gotypedjson.INT32, Value: int32(4)}

			data, err := json.Marshal(tInt32)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawData))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an incorrect value", func(t *testing.T) {
			tInt32 := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_int32","Value":"nope"}`), tInt32)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'nope' to an int32"))
		})

		t.Run("It can decode the value properly", func(t *testing.T) {
			tInt32 := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawData), tInt32)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tInt32.Type).To(Equal(gotypedjson.INT32))
			g.Expect(tInt32.Value.(int32)).To(Equal(int32(4)))
		})
	})
}

func Test_Int64(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":"_int64","Value":"4"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tInt64 := &gotypedjson.TypedJson{Type: gotypedjson.INT64, Value: "nope"}
			data, err := json.Marshal(tInt64)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to an int64"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode the value properly", func(t *testing.T) {
			tInt64 := &gotypedjson.TypedJson{Type: gotypedjson.INT64, Value: int64(4)}

			data, err := json.Marshal(tInt64)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawData))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an incorrect value", func(t *testing.T) {
			tInt64 := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_int64","Value":"nope"}`), tInt64)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'nope' to an int64"))
		})

		t.Run("It can decode the value properly", func(t *testing.T) {
			tInt16 := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawData), tInt16)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tInt16.Type).To(Equal(gotypedjson.INT64))
			g.Expect(tInt16.Value.(int64)).To(Equal(int64(4)))
		})
	})
}

func Test_Uint(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":"_uint","Value":"4"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tUint := &gotypedjson.TypedJson{Type: gotypedjson.UINT, Value: "nope"}
			data, err := json.Marshal(tUint)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to an uint"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode the value properly", func(t *testing.T) {
			tUint := &gotypedjson.TypedJson{Type: gotypedjson.UINT, Value: uint(4)}

			data, err := json.Marshal(tUint)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawData))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an incorrect value", func(t *testing.T) {
			tUint := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_uint","Value":"nope"}`), tUint)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'nope' to a uint"))
		})

		t.Run("It can decode the value properly", func(t *testing.T) {
			tInt := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawData), tInt)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tInt.Type).To(Equal(gotypedjson.UINT))
			g.Expect(tInt.Value.(uint)).To(Equal(uint(4)))
		})
	})
}

func Test_Uint8(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":"_uint8","Value":"4"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tUint8 := &gotypedjson.TypedJson{Type: gotypedjson.UINT8, Value: "nope"}
			data, err := json.Marshal(tUint8)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to an uint8"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode the value properly", func(t *testing.T) {
			tUint8 := &gotypedjson.TypedJson{Type: gotypedjson.UINT8, Value: uint8(4)}

			data, err := json.Marshal(tUint8)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawData))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an incorrect value", func(t *testing.T) {
			tUint8 := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_uint8","Value":"nope"}`), tUint8)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'nope' to a uint8"))
		})

		t.Run("It can decode the value properly", func(t *testing.T) {
			tUint8 := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawData), tUint8)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tUint8.Type).To(Equal(gotypedjson.UINT8))
			g.Expect(tUint8.Value.(uint8)).To(Equal(uint8(4)))
		})
	})
}

func Test_Uint16(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":"_uint16","Value":"4"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tUint16 := &gotypedjson.TypedJson{Type: gotypedjson.UINT16, Value: "nope"}
			data, err := json.Marshal(tUint16)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to an uint16"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode the value properly", func(t *testing.T) {
			tUint16 := &gotypedjson.TypedJson{Type: gotypedjson.UINT16, Value: uint16(4)}

			data, err := json.Marshal(tUint16)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawData))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an incorrect value", func(t *testing.T) {
			tUint16 := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_uint16","Value":"nope"}`), tUint16)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'nope' to a uint16"))
		})

		t.Run("It can decode the value properly", func(t *testing.T) {
			tUint16 := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawData), tUint16)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tUint16.Type).To(Equal(gotypedjson.UINT16))
			g.Expect(tUint16.Value.(uint16)).To(Equal(uint16(4)))
		})
	})
}

func Test_Uint32(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":"_uint32","Value":"4"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tUint32 := &gotypedjson.TypedJson{Type: gotypedjson.UINT32, Value: "nope"}
			data, err := json.Marshal(tUint32)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to an uint32"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode the value properly", func(t *testing.T) {
			tUint32 := &gotypedjson.TypedJson{Type: gotypedjson.UINT32, Value: uint32(4)}

			data, err := json.Marshal(tUint32)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawData))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an incorrect value", func(t *testing.T) {
			tUint32 := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_uint32","Value":"nope"}`), tUint32)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'nope' to a uint32"))
		})

		t.Run("It can decode the value properly", func(t *testing.T) {
			tUint32 := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawData), tUint32)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tUint32.Type).To(Equal(gotypedjson.UINT32))
			g.Expect(tUint32.Value.(uint32)).To(Equal(uint32(4)))
		})
	})
}

func Test_Uint64(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":"_uint64","Value":"4"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tUint64 := &gotypedjson.TypedJson{Type: gotypedjson.UINT64, Value: "nope"}
			data, err := json.Marshal(tUint64)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to an uint64"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode the value properly", func(t *testing.T) {
			tUint64 := &gotypedjson.TypedJson{Type: gotypedjson.UINT64, Value: uint64(4)}

			data, err := json.Marshal(tUint64)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawData))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an incorrect value", func(t *testing.T) {
			tUint64 := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_uint64","Value":"nope"}`), tUint64)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'nope' to a uint64"))
		})

		t.Run("It can decode the value properly", func(t *testing.T) {
			tUint64 := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawData), tUint64)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tUint64.Type).To(Equal(gotypedjson.UINT64))
			g.Expect(tUint64.Value.(uint64)).To(Equal(uint64(4)))
		})
	})
}

func Test_Float32(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":"_float32","Value":"4E+00"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tFloat32 := &gotypedjson.TypedJson{Type: gotypedjson.FLOAT32, Value: "nope"}
			data, err := json.Marshal(tFloat32)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to a float32"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode the value properly", func(t *testing.T) {
			tFloat32 := &gotypedjson.TypedJson{Type: gotypedjson.FLOAT32, Value: float32(4.0)}

			data, err := json.Marshal(tFloat32)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawData))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an incorrect value", func(t *testing.T) {
			tFloat32 := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_float32","Value":"nope"}`), tFloat32)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'nope' to a float32"))
		})

		t.Run("It can decode the value properly", func(t *testing.T) {
			tFloat32 := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawData), tFloat32)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tFloat32.Type).To(Equal(gotypedjson.FLOAT32))
			g.Expect(tFloat32.Value.(float32)).To(Equal(float32(4.0)))
		})
	})
}

func Test_Float64(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":"_float64","Value":"4E+00"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tFloat64 := &gotypedjson.TypedJson{Type: gotypedjson.FLOAT64, Value: "nope"}
			data, err := json.Marshal(tFloat64)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to a float64"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode the value properly", func(t *testing.T) {
			tFloat64 := &gotypedjson.TypedJson{Type: gotypedjson.FLOAT64, Value: float64(4.0)}

			data, err := json.Marshal(tFloat64)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawData))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an incorrect value", func(t *testing.T) {
			tFloat64 := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_float64","Value":"nope"}`), tFloat64)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'nope' to a float64"))
		})

		t.Run("It can decode the value properly", func(t *testing.T) {
			tFloat64 := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawData), tFloat64)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tFloat64.Type).To(Equal(gotypedjson.FLOAT64))
			g.Expect(tFloat64.Value.(float64)).To(Equal(float64(4.0)))
		})
	})
}

func Test_String(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":"_string","Value":"proper string"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tString := &gotypedjson.TypedJson{Type: gotypedjson.STRING, Value: 3}
			data, err := json.Marshal(tString)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast '3' to a string"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode the value properly", func(t *testing.T) {
			tString := &gotypedjson.TypedJson{Type: gotypedjson.STRING, Value: "proper string"}

			data, err := json.Marshal(tString)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawData))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an incorrect value", func(t *testing.T) {
			tString := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_string","Value":3}`), tString)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: cannot unmarshal number into Go struct field .Value of type string"))
		})

		t.Run("It can decode the value properly", func(t *testing.T) {
			tString := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawData), tString)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tString.Type).To(Equal(gotypedjson.STRING))
			g.Expect(tString.Value.(string)).To(Equal("proper string"))
		})
	})
}

func Test_Bool(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":"_bool","Value":"true"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tBool := &gotypedjson.TypedJson{Type: gotypedjson.BOOL, Value: "nope"}
			data, err := json.Marshal(tBool)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to a bool"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode the value properly", func(t *testing.T) {
			tBool := &gotypedjson.TypedJson{Type: gotypedjson.BOOL, Value: true}

			data, err := json.Marshal(tBool)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawData))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an incorrect value", func(t *testing.T) {
			tBool := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_bool","Value":"nope"}`), tBool)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'nope' to a bool"))
		})

		t.Run("It can decode the value properly", func(t *testing.T) {
			tBool := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawData), tBool)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tBool.Type).To(Equal(gotypedjson.BOOL))
			g.Expect(tBool.Value.(bool)).To(BeTrue())
		})
	})
}

func Test_DateTime(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tDateTime := &gotypedjson.TypedJson{Type: gotypedjson.DATETIME, Value: "nope"}
			data, err := json.Marshal(tDateTime)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to a datetime"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode the value properly", func(t *testing.T) {
			currentTime := time.Now()
			tDateTime := &gotypedjson.TypedJson{Type: gotypedjson.DATETIME, Value: currentTime}

			data, err := json.Marshal(tDateTime)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(fmt.Sprintf(`{"Type":"_datetime","Value":"%s"}`, currentTime.Format(time.RFC3339))))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an incorrect value", func(t *testing.T) {
			tDateTime := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_datetime","Value":"nope"}`), tDateTime)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'nope' to a datetime"))
		})

		t.Run("It can decode the value properly", func(t *testing.T) {
			tDateTime := &gotypedjson.TypedJson{}
			testTime := time.Now()

			err := json.Unmarshal([]byte(fmt.Sprintf(`{"Type":"_datetime","Value":"%s"}`, testTime.Format(time.RFC3339))), tDateTime)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tDateTime.Type).To(Equal(gotypedjson.DATETIME))
			g.Expect(tDateTime.Value.(time.Time).Format(time.RFC3339)).To(Equal(testTime.Format(time.RFC3339)))
		})
	})
}

func Test_Duration(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tDuration := &gotypedjson.TypedJson{Type: gotypedjson.TIME_DURATION, Value: "nope"}

			data, err := json.Marshal(tDuration)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to a time duration"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode the value properly", func(t *testing.T) {
			tDuration := &gotypedjson.TypedJson{Type: gotypedjson.TIME_DURATION, Value: time.Duration(4 * time.Millisecond)}

			data, err := json.Marshal(tDuration)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(`{"Type":"_duration","Value":"4ms"}`))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an incorrect value", func(t *testing.T) {
			tDuration := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_duration","Value":"nope"}`), tDuration)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'nope' to a time duration"))
		})

		t.Run("It can decode the value properly", func(t *testing.T) {
			tDuration := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_duration","Value":"4ms"}`), tDuration)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tDuration.Type).To(Equal(gotypedjson.TIME_DURATION))
			g.Expect(tDuration.Value.(time.Duration)).To(Equal(4 * time.Millisecond))
		})
	})
}

func Test_Complex64(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":"_complex64","Value":"(1E+00-3E+00i)"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tComplex64 := &gotypedjson.TypedJson{Type: gotypedjson.COMPLEX64, Value: "nope"}
			data, err := json.Marshal(tComplex64)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to a complex64"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode the value properly", func(t *testing.T) {
			tComplex64 := &gotypedjson.TypedJson{Type: gotypedjson.COMPLEX64, Value: complex64(complex(float32(1.0), float32(-3.0)))}

			data, err := json.Marshal(tComplex64)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawData))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an incorrect value", func(t *testing.T) {
			tComplex64 := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_complex64","Value":"nope"}`), tComplex64)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'nope' to a complex64"))
		})

		t.Run("It can decode the value properly", func(t *testing.T) {
			tComplex64 := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawData), tComplex64)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tComplex64.Type).To(Equal(gotypedjson.COMPLEX64))
			g.Expect(tComplex64.Value.(complex64)).To(Equal(complex64(complex(1, -3))))
		})
	})
}

func Test_Complex128(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":"_complex128","Value":"(1E+00-3E+00i)"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tComplex128 := &gotypedjson.TypedJson{Type: gotypedjson.COMPLEX128, Value: "nope"}
			data, err := json.Marshal(tComplex128)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to a complex128"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode the value properly", func(t *testing.T) {
			tComplex128 := &gotypedjson.TypedJson{Type: gotypedjson.COMPLEX128, Value: complex(1.0, -3.0)}

			data, err := json.Marshal(tComplex128)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawData))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an incorrect value", func(t *testing.T) {
			tComplex128 := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_complex128","Value":"nope"}`), tComplex128)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'nope' to a complex128"))
		})

		t.Run("It can decode the value properly", func(t *testing.T) {
			tComplex128 := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawData), tComplex128)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tComplex128.Type).To(Equal(gotypedjson.COMPLEX128))
			g.Expect(tComplex128.Value.(complex128)).To(Equal(complex(1, -3)))
		})
	})
}

func Test_Int_Slice(t *testing.T) {
	g := NewGomegaWithT(t)

	rawDataEmpty := `{"Type":"_int_array","Value":""}`
	rawDataSingle := `{"Type":"_int_array","Value":"1"}`
	rawDataMulti := `{"Type":"_int_array","Value":"1,2"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tInt := &gotypedjson.TypedJson{Type: gotypedjson.INT_SLICE, Value: "nope"}
			data, err := json.Marshal(tInt)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to a []int"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode an empty slice", func(t *testing.T) {
			tInt := &gotypedjson.TypedJson{Type: gotypedjson.INT_SLICE, Value: nil}

			data, err := json.Marshal(tInt)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataEmpty))
		})

		t.Run("It can encode a single value properly", func(t *testing.T) {
			tInt := &gotypedjson.TypedJson{Type: gotypedjson.INT_SLICE, Value: []int{1}}

			data, err := json.Marshal(tInt)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataSingle))
		})

		t.Run("It can encode multiple value properly", func(t *testing.T) {
			tInt := &gotypedjson.TypedJson{Type: gotypedjson.INT_SLICE, Value: []int{1, 2}}

			data, err := json.Marshal(tInt)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataMulti))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an invalid value", func(t *testing.T) {
			tInt := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_int_array","Value":"hello"}`), tInt)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'hello' to an int"))
		})

		t.Run("It can decode an empty slice", func(t *testing.T) {
			tInt := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataEmpty), tInt)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tInt.Type).To(Equal(gotypedjson.INT_SLICE))
			g.Expect(tInt.Value.([]int)).To(Equal([]int{}))
		})

		t.Run("It can decode a single value properly", func(t *testing.T) {
			tInt := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataSingle), tInt)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tInt.Type).To(Equal(gotypedjson.INT_SLICE))
			g.Expect(tInt.Value.([]int)).To(Equal([]int{1}))
		})

		t.Run("It can decode multiple values properly", func(t *testing.T) {
			tInt := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataMulti), tInt)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tInt.Type).To(Equal(gotypedjson.INT_SLICE))
			g.Expect(tInt.Value.([]int)).To(Equal([]int{1, 2}))
		})
	})
}

func Test_Int8_Slice(t *testing.T) {
	g := NewGomegaWithT(t)

	rawDataEmpty := `{"Type":"_int8_array","Value":""}`
	rawDataSingle := `{"Type":"_int8_array","Value":"1"}`
	rawDataMulti := `{"Type":"_int8_array","Value":"1,2"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tInt8S := &gotypedjson.TypedJson{Type: gotypedjson.INT8_SLICE, Value: "nope"}
			data, err := json.Marshal(tInt8S)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to a []int8"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode an empty slice", func(t *testing.T) {
			tInt8S := &gotypedjson.TypedJson{Type: gotypedjson.INT8_SLICE, Value: nil}

			data, err := json.Marshal(tInt8S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataEmpty))
		})

		t.Run("It can encode a single value properly", func(t *testing.T) {
			tInt8S := &gotypedjson.TypedJson{Type: gotypedjson.INT8_SLICE, Value: []int8{1}}

			data, err := json.Marshal(tInt8S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataSingle))
		})

		t.Run("It can encode multiple value properly", func(t *testing.T) {
			tInt8S := &gotypedjson.TypedJson{Type: gotypedjson.INT8_SLICE, Value: []int8{1, 2}}

			data, err := json.Marshal(tInt8S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataMulti))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an invalid value", func(t *testing.T) {
			tInt8S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_int8_array","Value":"hello"}`), tInt8S)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'hello' to an int8"))
		})

		t.Run("It can decode an empty slice", func(t *testing.T) {
			tInt8S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataEmpty), tInt8S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tInt8S.Type).To(Equal(gotypedjson.INT8_SLICE))
			g.Expect(tInt8S.Value.([]int8)).To(Equal([]int8{}))
		})

		t.Run("It can decode a single value properly", func(t *testing.T) {
			tInt8S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataSingle), tInt8S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tInt8S.Type).To(Equal(gotypedjson.INT8_SLICE))
			g.Expect(tInt8S.Value.([]int8)).To(Equal([]int8{1}))
		})

		t.Run("It can decode multiple values properly", func(t *testing.T) {
			tInt8S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataMulti), tInt8S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tInt8S.Type).To(Equal(gotypedjson.INT8_SLICE))
			g.Expect(tInt8S.Value.([]int8)).To(Equal([]int8{1, 2}))
		})
	})
}

func Test_Int16_Slice(t *testing.T) {
	g := NewGomegaWithT(t)

	rawDataEmpty := `{"Type":"_int16_array","Value":""}`
	rawDataSingle := `{"Type":"_int16_array","Value":"1"}`
	rawDataMulti := `{"Type":"_int16_array","Value":"1,2"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tInt16S := &gotypedjson.TypedJson{Type: gotypedjson.INT16_SLICE, Value: "nope"}
			data, err := json.Marshal(tInt16S)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to a []int16"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode an empty slice", func(t *testing.T) {
			tInt16S := &gotypedjson.TypedJson{Type: gotypedjson.INT16_SLICE, Value: nil}

			data, err := json.Marshal(tInt16S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataEmpty))
		})

		t.Run("It can encode a single value properly", func(t *testing.T) {
			tInt16S := &gotypedjson.TypedJson{Type: gotypedjson.INT16_SLICE, Value: []int16{1}}

			data, err := json.Marshal(tInt16S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataSingle))
		})

		t.Run("It can encode multiple value properly", func(t *testing.T) {
			tInt16S := &gotypedjson.TypedJson{Type: gotypedjson.INT16_SLICE, Value: []int16{1, 2}}

			data, err := json.Marshal(tInt16S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataMulti))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an invalid value", func(t *testing.T) {
			tInt16S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_int16_array","Value":"hello"}`), tInt16S)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'hello' to an int16"))
		})

		t.Run("It can decode an empty slice", func(t *testing.T) {
			tInt16S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataEmpty), tInt16S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tInt16S.Type).To(Equal(gotypedjson.INT16_SLICE))
			g.Expect(tInt16S.Value.([]int16)).To(Equal([]int16{}))
		})

		t.Run("It can decode a single value properly", func(t *testing.T) {
			tInt16S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataSingle), tInt16S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tInt16S.Type).To(Equal(gotypedjson.INT16_SLICE))
			g.Expect(tInt16S.Value.([]int16)).To(Equal([]int16{1}))
		})

		t.Run("It can decode multiple values properly", func(t *testing.T) {
			tInt16S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataMulti), tInt16S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tInt16S.Type).To(Equal(gotypedjson.INT16_SLICE))
			g.Expect(tInt16S.Value.([]int16)).To(Equal([]int16{1, 2}))
		})
	})
}

func Test_Int32_Slice(t *testing.T) {
	g := NewGomegaWithT(t)

	rawDataEmpty := `{"Type":"_int32_array","Value":""}`
	rawDataSingle := `{"Type":"_int32_array","Value":"1"}`
	rawDataMulti := `{"Type":"_int32_array","Value":"1,2"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tInt32S := &gotypedjson.TypedJson{Type: gotypedjson.INT32_SLICE, Value: "nope"}
			data, err := json.Marshal(tInt32S)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to a []int32"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode an empty slice", func(t *testing.T) {
			tInt32S := &gotypedjson.TypedJson{Type: gotypedjson.INT32_SLICE, Value: nil}

			data, err := json.Marshal(tInt32S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataEmpty))
		})

		t.Run("It can encode a single value properly", func(t *testing.T) {
			tInt32S := &gotypedjson.TypedJson{Type: gotypedjson.INT32_SLICE, Value: []int32{1}}

			data, err := json.Marshal(tInt32S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataSingle))
		})

		t.Run("It can encode multiple value properly", func(t *testing.T) {
			tInt32S := &gotypedjson.TypedJson{Type: gotypedjson.INT32_SLICE, Value: []int32{1, 2}}

			data, err := json.Marshal(tInt32S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataMulti))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an invalid value", func(t *testing.T) {
			tInt32S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_int32_array","Value":"hello"}`), tInt32S)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'hello' to an int32"))
		})

		t.Run("It can decode an empty slice", func(t *testing.T) {
			tInt32S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataEmpty), tInt32S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tInt32S.Type).To(Equal(gotypedjson.INT32_SLICE))
			g.Expect(tInt32S.Value.([]int32)).To(Equal([]int32{}))
		})

		t.Run("It can decode a single value properly", func(t *testing.T) {
			tInt32S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataSingle), tInt32S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tInt32S.Type).To(Equal(gotypedjson.INT32_SLICE))
			g.Expect(tInt32S.Value.([]int32)).To(Equal([]int32{1}))
		})

		t.Run("It can decode multiple values properly", func(t *testing.T) {
			tInt32S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataMulti), tInt32S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tInt32S.Type).To(Equal(gotypedjson.INT32_SLICE))
			g.Expect(tInt32S.Value.([]int32)).To(Equal([]int32{1, 2}))
		})
	})
}

func Test_Int64_Slice(t *testing.T) {
	g := NewGomegaWithT(t)

	rawDataEmpty := `{"Type":"_int64_array","Value":""}`
	rawDataSingle := `{"Type":"_int64_array","Value":"1"}`
	rawDataMulti := `{"Type":"_int64_array","Value":"1,2"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tInt64S := &gotypedjson.TypedJson{Type: gotypedjson.INT64_SLICE, Value: "nope"}
			data, err := json.Marshal(tInt64S)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to a []int64"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode an empty slice", func(t *testing.T) {
			tInt64S := &gotypedjson.TypedJson{Type: gotypedjson.INT64_SLICE, Value: nil}

			data, err := json.Marshal(tInt64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataEmpty))
		})

		t.Run("It can encode a single value properly", func(t *testing.T) {
			tInt64S := &gotypedjson.TypedJson{Type: gotypedjson.INT64_SLICE, Value: []int64{1}}

			data, err := json.Marshal(tInt64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataSingle))
		})

		t.Run("It can encode multiple value properly", func(t *testing.T) {
			tInt64S := &gotypedjson.TypedJson{Type: gotypedjson.INT64_SLICE, Value: []int64{1, 2}}

			data, err := json.Marshal(tInt64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataMulti))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an invalid value", func(t *testing.T) {
			tInt64S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_int64_array","Value":"hello"}`), tInt64S)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'hello' to an int64"))
		})

		t.Run("It can decode an empty slice", func(t *testing.T) {
			tInt64S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataEmpty), tInt64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tInt64S.Type).To(Equal(gotypedjson.INT64_SLICE))
			g.Expect(tInt64S.Value.([]int64)).To(Equal([]int64{}))
		})

		t.Run("It can decode a single value properly", func(t *testing.T) {
			tInt64S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataSingle), tInt64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tInt64S.Type).To(Equal(gotypedjson.INT64_SLICE))
			g.Expect(tInt64S.Value.([]int64)).To(Equal([]int64{1}))
		})

		t.Run("It can decode multiple values properly", func(t *testing.T) {
			tInt64S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataMulti), tInt64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tInt64S.Type).To(Equal(gotypedjson.INT64_SLICE))
			g.Expect(tInt64S.Value.([]int64)).To(Equal([]int64{1, 2}))
		})
	})
}

func Test_Uint_Slice(t *testing.T) {
	g := NewGomegaWithT(t)

	rawDataEmpty := `{"Type":"_uint_array","Value":""}`
	rawDataSingle := `{"Type":"_uint_array","Value":"1"}`
	rawDataMulti := `{"Type":"_uint_array","Value":"1,2"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tUint := &gotypedjson.TypedJson{Type: gotypedjson.UINT_SLICE, Value: "nope"}
			data, err := json.Marshal(tUint)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to a []uint"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode an empty slice", func(t *testing.T) {
			tUintS := &gotypedjson.TypedJson{Type: gotypedjson.UINT_SLICE, Value: nil}

			data, err := json.Marshal(tUintS)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataEmpty))
		})

		t.Run("It can encode a single value properly", func(t *testing.T) {
			tUint := &gotypedjson.TypedJson{Type: gotypedjson.UINT_SLICE, Value: []uint{1}}

			data, err := json.Marshal(tUint)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataSingle))
		})

		t.Run("It can encode multiple value properly", func(t *testing.T) {
			tUint := &gotypedjson.TypedJson{Type: gotypedjson.UINT_SLICE, Value: []uint{1, 2}}

			data, err := json.Marshal(tUint)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataMulti))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an invalid value", func(t *testing.T) {
			tUint := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_uint_array","Value":"hello"}`), tUint)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'hello' to a uint"))
		})

		t.Run("It can decode an empty slice", func(t *testing.T) {
			tUintS := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataEmpty), tUintS)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tUintS.Type).To(Equal(gotypedjson.UINT_SLICE))
			g.Expect(tUintS.Value.([]uint)).To(Equal([]uint{}))
		})

		t.Run("It can decode a single value properly", func(t *testing.T) {
			tUint := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataSingle), tUint)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tUint.Type).To(Equal(gotypedjson.UINT_SLICE))
			g.Expect(tUint.Value.([]uint)).To(Equal([]uint{1}))
		})

		t.Run("It can decode multiple values properly", func(t *testing.T) {
			tUint := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataMulti), tUint)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tUint.Type).To(Equal(gotypedjson.UINT_SLICE))
			g.Expect(tUint.Value.([]uint)).To(Equal([]uint{1, 2}))
		})
	})
}

func Test_Uint8_Slice(t *testing.T) {
	g := NewGomegaWithT(t)

	rawDataEmpty := `{"Type":"_uint8_array","Value":""}`
	rawDataSingle := `{"Type":"_uint8_array","Value":"1"}`
	rawDataMulti := `{"Type":"_uint8_array","Value":"1,2"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tUint8S := &gotypedjson.TypedJson{Type: gotypedjson.UINT8_SLICE, Value: "nope"}
			data, err := json.Marshal(tUint8S)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to a []uint8"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode an empty slice", func(t *testing.T) {
			tUint8S := &gotypedjson.TypedJson{Type: gotypedjson.UINT8_SLICE, Value: nil}

			data, err := json.Marshal(tUint8S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataEmpty))
		})

		t.Run("It can encode a single value properly", func(t *testing.T) {
			tUint8S := &gotypedjson.TypedJson{Type: gotypedjson.UINT8_SLICE, Value: []uint8{1}}

			data, err := json.Marshal(tUint8S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataSingle))
		})

		t.Run("It can encode multiple value properly", func(t *testing.T) {
			tUint8S := &gotypedjson.TypedJson{Type: gotypedjson.UINT8_SLICE, Value: []uint8{1, 2}}

			data, err := json.Marshal(tUint8S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataMulti))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an invalid value", func(t *testing.T) {
			tUint8S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_uint8_array","Value":"hello"}`), tUint8S)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'hello' to a uint8"))
		})

		t.Run("It can decode an empty slice", func(t *testing.T) {
			tUint8S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataEmpty), tUint8S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tUint8S.Type).To(Equal(gotypedjson.UINT8_SLICE))
			g.Expect(tUint8S.Value.([]uint8)).To(Equal([]uint8{}))
		})

		t.Run("It can decode a single value properly", func(t *testing.T) {
			tUint8S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataSingle), tUint8S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tUint8S.Type).To(Equal(gotypedjson.UINT8_SLICE))
			g.Expect(tUint8S.Value.([]uint8)).To(Equal([]uint8{1}))
		})

		t.Run("It can decode multiple values properly", func(t *testing.T) {
			tUint8S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataMulti), tUint8S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tUint8S.Type).To(Equal(gotypedjson.UINT8_SLICE))
			g.Expect(tUint8S.Value.([]uint8)).To(Equal([]uint8{1, 2}))
		})
	})
}

func Test_Uint16_Slice(t *testing.T) {
	g := NewGomegaWithT(t)

	rawDataEmpty := `{"Type":"_uint16_array","Value":""}`
	rawDataSingle := `{"Type":"_uint16_array","Value":"1"}`
	rawDataMulti := `{"Type":"_uint16_array","Value":"1,2"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tUint16S := &gotypedjson.TypedJson{Type: gotypedjson.UINT16_SLICE, Value: "nope"}
			data, err := json.Marshal(tUint16S)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to a []uint16"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode an empty slice", func(t *testing.T) {
			tUint8S := &gotypedjson.TypedJson{Type: gotypedjson.UINT16_SLICE, Value: nil}

			data, err := json.Marshal(tUint8S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataEmpty))
		})

		t.Run("It can encode a single value properly", func(t *testing.T) {
			tUint16S := &gotypedjson.TypedJson{Type: gotypedjson.UINT16_SLICE, Value: []uint16{1}}

			data, err := json.Marshal(tUint16S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataSingle))
		})

		t.Run("It can encode multiple value properly", func(t *testing.T) {
			tUint16S := &gotypedjson.TypedJson{Type: gotypedjson.UINT16_SLICE, Value: []uint16{1, 2}}

			data, err := json.Marshal(tUint16S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataMulti))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an invalid value", func(t *testing.T) {
			tUint16S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_uint16_array","Value":"hello"}`), tUint16S)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'hello' to a uint16"))
		})

		t.Run("It can decode an empty slice", func(t *testing.T) {
			tUint16S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataEmpty), tUint16S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tUint16S.Type).To(Equal(gotypedjson.UINT16_SLICE))
			g.Expect(tUint16S.Value.([]uint16)).To(Equal([]uint16{}))
		})

		t.Run("It can decode a single value properly", func(t *testing.T) {
			tUint16S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataSingle), tUint16S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tUint16S.Type).To(Equal(gotypedjson.UINT16_SLICE))
			g.Expect(tUint16S.Value.([]uint16)).To(Equal([]uint16{1}))
		})

		t.Run("It can decode multiple values properly", func(t *testing.T) {
			tUint16S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataMulti), tUint16S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tUint16S.Type).To(Equal(gotypedjson.UINT16_SLICE))
			g.Expect(tUint16S.Value.([]uint16)).To(Equal([]uint16{1, 2}))
		})
	})
}

func Test_Uint32_Slice(t *testing.T) {
	g := NewGomegaWithT(t)

	rawDataEmpty := `{"Type":"_uint32_array","Value":""}`
	rawDataSingle := `{"Type":"_uint32_array","Value":"1"}`
	rawDataMulti := `{"Type":"_uint32_array","Value":"1,2"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tUint32S := &gotypedjson.TypedJson{Type: gotypedjson.UINT32_SLICE, Value: "nope"}
			data, err := json.Marshal(tUint32S)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to a []uint32"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode an empty slice", func(t *testing.T) {
			tUint32S := &gotypedjson.TypedJson{Type: gotypedjson.UINT32_SLICE, Value: nil}

			data, err := json.Marshal(tUint32S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataEmpty))
		})

		t.Run("It can encode a single value properly", func(t *testing.T) {
			tUint32S := &gotypedjson.TypedJson{Type: gotypedjson.UINT32_SLICE, Value: []uint32{1}}

			data, err := json.Marshal(tUint32S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataSingle))
		})

		t.Run("It can encode multiple value properly", func(t *testing.T) {
			tUint32S := &gotypedjson.TypedJson{Type: gotypedjson.UINT32_SLICE, Value: []uint32{1, 2}}

			data, err := json.Marshal(tUint32S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataMulti))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an invalid value", func(t *testing.T) {
			tUint32S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_uint32_array","Value":"hello"}`), tUint32S)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'hello' to a uint32"))
		})

		t.Run("It can decode an empty slice", func(t *testing.T) {
			tUint32S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataEmpty), tUint32S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tUint32S.Type).To(Equal(gotypedjson.UINT32_SLICE))
			g.Expect(tUint32S.Value.([]uint32)).To(Equal([]uint32{}))
		})

		t.Run("It can decode a single value properly", func(t *testing.T) {
			tUint32S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataSingle), tUint32S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tUint32S.Type).To(Equal(gotypedjson.UINT32_SLICE))
			g.Expect(tUint32S.Value.([]uint32)).To(Equal([]uint32{1}))
		})

		t.Run("It can decode multiple values properly", func(t *testing.T) {
			tUint32S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataMulti), tUint32S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tUint32S.Type).To(Equal(gotypedjson.UINT32_SLICE))
			g.Expect(tUint32S.Value.([]uint32)).To(Equal([]uint32{1, 2}))
		})
	})
}

func Test_Uint64_Slice(t *testing.T) {
	g := NewGomegaWithT(t)

	rawDataEmpty := `{"Type":"_uint64_array","Value":""}`
	rawDataSingle := `{"Type":"_uint64_array","Value":"1"}`
	rawDataMulti := `{"Type":"_uint64_array","Value":"1,2"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tUint64S := &gotypedjson.TypedJson{Type: gotypedjson.UINT64_SLICE, Value: "nope"}
			data, err := json.Marshal(tUint64S)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to a []uint64"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode an empty slice", func(t *testing.T) {
			tUint64S := &gotypedjson.TypedJson{Type: gotypedjson.UINT64_SLICE, Value: nil}

			data, err := json.Marshal(tUint64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataEmpty))
		})

		t.Run("It can encode a single value properly", func(t *testing.T) {
			tUint64S := &gotypedjson.TypedJson{Type: gotypedjson.UINT64_SLICE, Value: []uint64{1}}

			data, err := json.Marshal(tUint64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataSingle))
		})

		t.Run("It can encode multiple value properly", func(t *testing.T) {
			tUint64S := &gotypedjson.TypedJson{Type: gotypedjson.UINT64_SLICE, Value: []uint64{1, 2}}

			data, err := json.Marshal(tUint64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataMulti))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an invalid value", func(t *testing.T) {
			tUint64S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_uint64_array","Value":"hello"}`), tUint64S)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'hello' to a uint64"))
		})

		t.Run("It can decode an empty slice", func(t *testing.T) {
			tUint64S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataEmpty), tUint64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tUint64S.Type).To(Equal(gotypedjson.UINT64_SLICE))
			g.Expect(tUint64S.Value.([]uint64)).To(Equal([]uint64{}))
		})

		t.Run("It can decode a single value properly", func(t *testing.T) {
			tUint64S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataSingle), tUint64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tUint64S.Type).To(Equal(gotypedjson.UINT64_SLICE))
			g.Expect(tUint64S.Value.([]uint64)).To(Equal([]uint64{1}))
		})

		t.Run("It can decode multiple values properly", func(t *testing.T) {
			tUint64S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataMulti), tUint64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tUint64S.Type).To(Equal(gotypedjson.UINT64_SLICE))
			g.Expect(tUint64S.Value.([]uint64)).To(Equal([]uint64{1, 2}))
		})
	})
}

func Test_Float32_Slice(t *testing.T) {
	g := NewGomegaWithT(t)

	rawDataEmpty := `{"Type":"_float32_array","Value":""}`
	rawDataSingle := `{"Type":"_float32_array","Value":"1E+00"}`
	rawDataMulti := `{"Type":"_float32_array","Value":"1E+00,2E+00"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tFloat32S := &gotypedjson.TypedJson{Type: gotypedjson.FLOAT32_SLICE, Value: "nope"}
			data, err := json.Marshal(tFloat32S)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to a []float32"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode an empty slice", func(t *testing.T) {
			tFloat32S := &gotypedjson.TypedJson{Type: gotypedjson.FLOAT32_SLICE, Value: nil}

			data, err := json.Marshal(tFloat32S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataEmpty))
		})

		t.Run("It can encode a single value properly", func(t *testing.T) {
			tFloat32S := &gotypedjson.TypedJson{Type: gotypedjson.FLOAT32_SLICE, Value: []float32{1}}

			data, err := json.Marshal(tFloat32S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataSingle))
		})

		t.Run("It can encode multiple value properly", func(t *testing.T) {
			tFloat32S := &gotypedjson.TypedJson{Type: gotypedjson.FLOAT32_SLICE, Value: []float32{1, 2}}

			data, err := json.Marshal(tFloat32S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataMulti))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an invalid value", func(t *testing.T) {
			tFloat32S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_float32_array","Value":"hello"}`), tFloat32S)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'hello' to a float32"))
		})

		t.Run("It can decode an empty slice", func(t *testing.T) {
			tFloat32S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataEmpty), tFloat32S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tFloat32S.Type).To(Equal(gotypedjson.FLOAT32_SLICE))
			g.Expect(tFloat32S.Value.([]float32)).To(Equal([]float32{}))
		})

		t.Run("It can decode a single value properly", func(t *testing.T) {
			tFloat32S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataSingle), tFloat32S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tFloat32S.Type).To(Equal(gotypedjson.FLOAT32_SLICE))
			g.Expect(tFloat32S.Value.([]float32)).To(Equal([]float32{1}))
		})

		t.Run("It can decode multiple values properly", func(t *testing.T) {
			tFloat32S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataMulti), tFloat32S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tFloat32S.Type).To(Equal(gotypedjson.FLOAT32_SLICE))
			g.Expect(tFloat32S.Value.([]float32)).To(Equal([]float32{1, 2}))
		})
	})
}

func Test_Float64_Slice(t *testing.T) {
	g := NewGomegaWithT(t)

	rawDataEmpty := `{"Type":"_float64_array","Value":""}`
	rawDataSingle := `{"Type":"_float64_array","Value":"1E+00"}`
	rawDataMulti := `{"Type":"_float64_array","Value":"1E+00,2E+00"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tFloat64S := &gotypedjson.TypedJson{Type: gotypedjson.FLOAT64_SLICE, Value: "nope"}
			data, err := json.Marshal(tFloat64S)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to a []float64"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode an empty slice", func(t *testing.T) {
			tFloat64S := &gotypedjson.TypedJson{Type: gotypedjson.FLOAT64_SLICE, Value: nil}

			data, err := json.Marshal(tFloat64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataEmpty))
		})

		t.Run("It can encode a single value properly", func(t *testing.T) {
			tFloat64S := &gotypedjson.TypedJson{Type: gotypedjson.FLOAT64_SLICE, Value: []float64{1}}

			data, err := json.Marshal(tFloat64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataSingle))
		})

		t.Run("It can encode multiple value properly", func(t *testing.T) {
			tFloat64S := &gotypedjson.TypedJson{Type: gotypedjson.FLOAT64_SLICE, Value: []float64{1, 2}}

			data, err := json.Marshal(tFloat64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataMulti))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an invalid value", func(t *testing.T) {
			tFloat64S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_float64_array","Value":"hello"}`), tFloat64S)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'hello' to a float64"))
		})

		t.Run("It can decode an empty slice", func(t *testing.T) {
			tFloat64S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataEmpty), tFloat64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tFloat64S.Type).To(Equal(gotypedjson.FLOAT64_SLICE))
			g.Expect(tFloat64S.Value.([]float64)).To(Equal([]float64{}))
		})

		t.Run("It can decode a single value properly", func(t *testing.T) {
			tFloat64S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataSingle), tFloat64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tFloat64S.Type).To(Equal(gotypedjson.FLOAT64_SLICE))
			g.Expect(tFloat64S.Value.([]float64)).To(Equal([]float64{1}))
		})

		t.Run("It can decode multiple values properly", func(t *testing.T) {
			tFloat64S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataMulti), tFloat64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tFloat64S.Type).To(Equal(gotypedjson.FLOAT64_SLICE))
			g.Expect(tFloat64S.Value.([]float64)).To(Equal([]float64{1, 2}))
		})
	})
}

func Test_String_Slice(t *testing.T) {
	g := NewGomegaWithT(t)

	rawDataEmpty := `{"Type":"_string_array","Value":""}`
	rawDataSingle := `{"Type":"_string_array","Value":"aGVsbG8="}`
	rawDataMulti := `{"Type":"_string_array","Value":"aGVsbG8=,d29ybGQ="}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			stringS := &gotypedjson.TypedJson{Type: gotypedjson.STRING_SLICE, Value: 1}
			data, err := json.Marshal(stringS)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast '1' to a []string"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode an empty slice", func(t *testing.T) {
			tStringS := &gotypedjson.TypedJson{Type: gotypedjson.STRING_SLICE, Value: nil}

			data, err := json.Marshal(tStringS)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataEmpty))
		})

		t.Run("It can encode a single value properly", func(t *testing.T) {
			stringS := &gotypedjson.TypedJson{Type: gotypedjson.STRING_SLICE, Value: []string{"hello"}}

			data, err := json.Marshal(stringS)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataSingle))
		})

		t.Run("It can encode multiple value properly", func(t *testing.T) {
			stringS := &gotypedjson.TypedJson{Type: gotypedjson.STRING_SLICE, Value: []string{"hello", "world"}}

			data, err := json.Marshal(stringS)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataMulti))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode a non-base64 value", func(t *testing.T) {
			tFloat64S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_string_array","Value":"/>??a"}`), tFloat64S)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("string '/>??a' is not an expected base64"))
		})

		t.Run("It can decode an empty slice", func(t *testing.T) {
			tStringS := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataEmpty), tStringS)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tStringS.Type).To(Equal(gotypedjson.STRING_SLICE))
			g.Expect(tStringS.Value.([]string)).To(Equal([]string{}))
		})

		t.Run("It can decode a single value properly", func(t *testing.T) {
			tFloat64S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataSingle), tFloat64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tFloat64S.Type).To(Equal(gotypedjson.STRING_SLICE))
			g.Expect(tFloat64S.Value.([]string)).To(Equal([]string{"hello"}))
		})

		t.Run("It can decode multiple values properly", func(t *testing.T) {
			tFloat64S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataMulti), tFloat64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tFloat64S.Type).To(Equal(gotypedjson.STRING_SLICE))
			g.Expect(tFloat64S.Value.([]string)).To(Equal([]string{"hello", "world"}))
		})
	})
}

func Test_Bool_Slice(t *testing.T) {
	g := NewGomegaWithT(t)

	rawDataEmpty := `{"Type":"_bool_array","Value":""}`
	rawDataSingle := `{"Type":"_bool_array","Value":"true"}`
	rawDataMulti := `{"Type":"_bool_array","Value":"true,false"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tBoolS := &gotypedjson.TypedJson{Type: gotypedjson.BOOL_SLICE, Value: "nope"}
			data, err := json.Marshal(tBoolS)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to a []bool"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode an empty slice", func(t *testing.T) {
			tBoolS := &gotypedjson.TypedJson{Type: gotypedjson.BOOL_SLICE, Value: nil}

			data, err := json.Marshal(tBoolS)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataEmpty))
		})

		t.Run("It can encode a single value properly", func(t *testing.T) {
			tBoolS := &gotypedjson.TypedJson{Type: gotypedjson.BOOL_SLICE, Value: []bool{true}}

			data, err := json.Marshal(tBoolS)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataSingle))
		})

		t.Run("It can encode multiple value properly", func(t *testing.T) {
			tBoolS := &gotypedjson.TypedJson{Type: gotypedjson.BOOL_SLICE, Value: []bool{true, false}}

			data, err := json.Marshal(tBoolS)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataMulti))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an invalid value", func(t *testing.T) {
			tBoolS := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_bool_array","Value":"hello"}`), tBoolS)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'hello' to a bool"))
		})

		t.Run("It can decode an empty slice", func(t *testing.T) {
			tBoolS := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataEmpty), tBoolS)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tBoolS.Type).To(Equal(gotypedjson.BOOL_SLICE))
			g.Expect(tBoolS.Value.([]bool)).To(Equal([]bool{}))
		})

		t.Run("It can decode a single value properly", func(t *testing.T) {
			tBoolS := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataSingle), tBoolS)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tBoolS.Type).To(Equal(gotypedjson.BOOL_SLICE))
			g.Expect(tBoolS.Value.([]bool)).To(Equal([]bool{true}))
		})

		t.Run("It can decode multiple values properly", func(t *testing.T) {
			tBoolS := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataMulti), tBoolS)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tBoolS.Type).To(Equal(gotypedjson.BOOL_SLICE))
			g.Expect(tBoolS.Value.([]bool)).To(Equal([]bool{true, false}))
		})
	})
}

func Test_DateTime_Slice(t *testing.T) {
	g := NewGomegaWithT(t)

	timeOne := time.Now()
	timeTwo := time.Now()

	rawDataEmpty := `{"Type":"_datetime_array","Value":""}`
	rawDataSingle := fmt.Sprintf(`{"Type":"_datetime_array","Value":"%s"}`, timeOne.Format(time.RFC3339))
	rawDataMulti := fmt.Sprintf(`{"Type":"_datetime_array","Value":"%s,%s"}`, timeOne.Format(time.RFC3339), timeTwo.Format(time.RFC3339))

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tDateTimeS := &gotypedjson.TypedJson{Type: gotypedjson.DATETIME_SLICE, Value: "nope"}

			data, err := json.Marshal(tDateTimeS)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to a []datetime"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode an empty slice", func(t *testing.T) {
			tDateTimeS := &gotypedjson.TypedJson{Type: gotypedjson.DATETIME_SLICE, Value: nil}

			data, err := json.Marshal(tDateTimeS)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataEmpty))
		})

		t.Run("It can encode a single value properly", func(t *testing.T) {
			tDateTimeS := &gotypedjson.TypedJson{Type: gotypedjson.DATETIME_SLICE, Value: []time.Time{timeOne}}

			data, err := json.Marshal(tDateTimeS)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataSingle))
		})

		t.Run("It can encode multiple value properly", func(t *testing.T) {
			tDateTimeS := &gotypedjson.TypedJson{Type: gotypedjson.DATETIME_SLICE, Value: []time.Time{timeOne, timeTwo}}

			data, err := json.Marshal(tDateTimeS)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataMulti))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an invalid value", func(t *testing.T) {
			tDateTimeS := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_datetime_array","Value":"hello"}`), tDateTimeS)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'hello' to a datetime"))
		})

		t.Run("It can decode an empty slice", func(t *testing.T) {
			tDateTimeS := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataEmpty), tDateTimeS)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tDateTimeS.Type).To(Equal(gotypedjson.DATETIME_SLICE))
			g.Expect(tDateTimeS.Value.([]time.Time)).To(Equal([]time.Time{}))
		})

		t.Run("It can decode a single value properly", func(t *testing.T) {
			tDateTimeS := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataSingle), tDateTimeS)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tDateTimeS.Type).To(Equal(gotypedjson.DATETIME_SLICE))
			g.Expect(len(tDateTimeS.Value.(([]time.Time)))).To(Equal(1))
			g.Expect(tDateTimeS.Value.([]time.Time)[0].Format(time.RFC3339)).To(Equal(timeOne.Format(time.RFC3339)))
		})

		t.Run("It can decode multiple values properly", func(t *testing.T) {
			tDateTimeS := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataMulti), tDateTimeS)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tDateTimeS.Type).To(Equal(gotypedjson.DATETIME_SLICE))
			g.Expect(len(tDateTimeS.Value.(([]time.Time)))).To(Equal(2))
			g.Expect(tDateTimeS.Value.([]time.Time)[0].Format(time.RFC3339)).To(Equal(timeOne.Format(time.RFC3339)))
			g.Expect(tDateTimeS.Value.([]time.Time)[1].Format(time.RFC3339)).To(Equal(timeTwo.Format(time.RFC3339)))
		})
	})
}

func Test_Duration_Slice(t *testing.T) {
	g := NewGomegaWithT(t)

	rawDataEmpty := `{"Type":"_duration_array","Value":""}`
	rawDataSingle := `{"Type":"_duration_array","Value":"2s"}`
	rawDataMulti := `{"Type":"_duration_array","Value":"2s,3ms"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tDurationS := &gotypedjson.TypedJson{Type: gotypedjson.TIME_DURATION_SLICE, Value: "nope"}

			data, err := json.Marshal(tDurationS)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to a []duration"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode an empty slice", func(t *testing.T) {
			tDurationS := &gotypedjson.TypedJson{Type: gotypedjson.TIME_DURATION_SLICE, Value: nil}

			data, err := json.Marshal(tDurationS)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataEmpty))
		})

		t.Run("It can encode a single value properly", func(t *testing.T) {
			tDurationS := &gotypedjson.TypedJson{Type: gotypedjson.TIME_DURATION_SLICE, Value: []time.Duration{2 * time.Second}}

			data, err := json.Marshal(tDurationS)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataSingle))
		})

		t.Run("It can encode multiple value properly", func(t *testing.T) {
			tDurationS := &gotypedjson.TypedJson{Type: gotypedjson.TIME_DURATION_SLICE, Value: []time.Duration{2 * time.Second, 3 * time.Millisecond}}

			data, err := json.Marshal(tDurationS)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataMulti))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an invalid value", func(t *testing.T) {
			tDurationS := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_duration_array","Value":"hello"}`), tDurationS)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'hello' to a duration"))
		})

		t.Run("It can decode an empty slice", func(t *testing.T) {
			tDurationS := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataEmpty), tDurationS)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tDurationS.Type).To(Equal(gotypedjson.TIME_DURATION_SLICE))
			g.Expect(tDurationS.Value.([]time.Duration)).To(Equal([]time.Duration{}))
		})

		t.Run("It can decode a single value properly", func(t *testing.T) {
			tDurationS := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataSingle), tDurationS)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tDurationS.Type).To(Equal(gotypedjson.TIME_DURATION_SLICE))
			g.Expect(tDurationS.Value.(([]time.Duration))).To(Equal([]time.Duration{2 * time.Second}))
		})

		t.Run("It can decode multiple values properly", func(t *testing.T) {
			tDurationS := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataMulti), tDurationS)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tDurationS.Type).To(Equal(gotypedjson.TIME_DURATION_SLICE))
			g.Expect(tDurationS.Value.(([]time.Duration))).To(Equal([]time.Duration{2 * time.Second, 3 * time.Millisecond}))
		})
	})
}

func Test_Complex64_Slice(t *testing.T) {
	g := NewGomegaWithT(t)

	rawDataEmpty := `{"Type":"_complex64_array","Value":""}`
	rawDataSingle := `{"Type":"_complex64_array","Value":"(1E+00-3E+00i)"}`
	rawDataMulti := `{"Type":"_complex64_array","Value":"(1E+00-3E+00i),(2E+00-5E+00i)"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tComplex64S := &gotypedjson.TypedJson{Type: gotypedjson.COMPLEX64_SLICE, Value: "nope"}
			data, err := json.Marshal(tComplex64S)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to a []complex64"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode an empty slice", func(t *testing.T) {
			tComplex64S := &gotypedjson.TypedJson{Type: gotypedjson.COMPLEX64_SLICE, Value: nil}

			data, err := json.Marshal(tComplex64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataEmpty))
		})

		t.Run("It can encode a single value properly", func(t *testing.T) {
			tComplex64S := &gotypedjson.TypedJson{Type: gotypedjson.COMPLEX64_SLICE, Value: []complex64{complex64(complex(float32(1), float32(-3)))}}

			data, err := json.Marshal(tComplex64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataSingle))
		})

		t.Run("It can encode multiple value properly", func(t *testing.T) {
			tComplex64S := &gotypedjson.TypedJson{Type: gotypedjson.COMPLEX64_SLICE, Value: []complex64{complex64(complex(float32(1), float32(-3))), complex64(complex(float32(2), float32(-5)))}}

			data, err := json.Marshal(tComplex64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataMulti))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an invalid value", func(t *testing.T) {
			tComplex64S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_complex64_array","Value":"hello"}`), tComplex64S)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'hello' to a complex64"))
		})

		t.Run("It can decode an empty slice", func(t *testing.T) {
			tComplex64S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataEmpty), tComplex64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tComplex64S.Type).To(Equal(gotypedjson.COMPLEX64_SLICE))
			g.Expect(tComplex64S.Value.([]complex64)).To(Equal([]complex64{}))
		})

		t.Run("It can decode a single value properly", func(t *testing.T) {
			tComplex64S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataSingle), tComplex64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tComplex64S.Type).To(Equal(gotypedjson.COMPLEX64_SLICE))
			g.Expect(tComplex64S.Value.([]complex64)).To(Equal([]complex64{complex64(complex(float32(1), float32(-3)))}))
		})

		t.Run("It can decode multiple values properly", func(t *testing.T) {
			tComplex64S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataMulti), tComplex64S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tComplex64S.Type).To(Equal(gotypedjson.COMPLEX64_SLICE))
			g.Expect(tComplex64S.Value.([]complex64)).To(Equal([]complex64{complex64(complex(float32(1), float32(-3))), complex64(complex(float32(2), float32(-5)))}))
		})
	})
}

func Test_Complex128_Slice(t *testing.T) {
	g := NewGomegaWithT(t)

	rawDataEmpty := `{"Type":"_complex128_array","Value":""}`
	rawDataSingle := `{"Type":"_complex128_array","Value":"(1E+00-3E+00i)"}`
	rawDataMulti := `{"Type":"_complex128_array","Value":"(1E+00-3E+00i),(2E+00-5E+00i)"}`

	t.Run("Encoding", func(t *testing.T) {
		t.Run("It returns an error if the type can not be cast", func(t *testing.T) {
			tComplex128S := &gotypedjson.TypedJson{Type: gotypedjson.COMPLEX128_SLICE, Value: "nope"}
			data, err := json.Marshal(tComplex128S)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("json: error calling MarshalJSON for type *gotypedjson.TypedJson: failed to cast 'nope' to a []complex128"))
			g.Expect(data).To(BeNil())
		})

		t.Run("It can encode an empty slice", func(t *testing.T) {
			tComplex128S := &gotypedjson.TypedJson{Type: gotypedjson.COMPLEX128_SLICE, Value: nil}

			data, err := json.Marshal(tComplex128S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataEmpty))
		})

		t.Run("It can encode a single value properly", func(t *testing.T) {
			tComplex128S := &gotypedjson.TypedJson{Type: gotypedjson.COMPLEX128_SLICE, Value: []complex128{complex(float64(1), float64(-3))}}

			data, err := json.Marshal(tComplex128S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataSingle))
		})

		t.Run("It can encode multiple value properly", func(t *testing.T) {
			tComplex128S := &gotypedjson.TypedJson{Type: gotypedjson.COMPLEX128_SLICE, Value: []complex128{complex(float64(1), float64(-3)), complex(float64(2), float64(-5))}}

			data, err := json.Marshal(tComplex128S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(rawDataMulti))
		})
	})

	t.Run("Decoding", func(t *testing.T) {
		t.Run("It fails to decode an invalid value", func(t *testing.T) {
			tComplex128S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(`{"Type":"_complex128_array","Value":"hello"}`), tComplex128S)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert 'hello' to a complex128"))
		})

		t.Run("It can decode an empty slice", func(t *testing.T) {
			tComplex128S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataEmpty), tComplex128S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tComplex128S.Type).To(Equal(gotypedjson.COMPLEX128_SLICE))
			g.Expect(tComplex128S.Value.([]complex128)).To(Equal([]complex128{}))
		})

		t.Run("It can decode a single value properly", func(t *testing.T) {
			tComplex128S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataSingle), tComplex128S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tComplex128S.Type).To(Equal(gotypedjson.COMPLEX128_SLICE))
			g.Expect(tComplex128S.Value.([]complex128)).To(Equal([]complex128{complex(float64(1), float64(-3))}))
		})

		t.Run("It can decode multiple values properly", func(t *testing.T) {
			tComplex128S := &gotypedjson.TypedJson{}

			err := json.Unmarshal([]byte(rawDataMulti), tComplex128S)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tComplex128S.Type).To(Equal(gotypedjson.COMPLEX128_SLICE))
			g.Expect(tComplex128S.Value.([]complex128)).To(Equal([]complex128{complex(float64(1), float64(-3)), complex(float64(2), float64(-5))}))
		})
	})
}

func Test_Codec(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("It errors if the encoder does not know they type", func(t *testing.T) {
		tCustom := gotypedjson.NewTypedJson(gotypedjson.JSONTYPE("test"), 0, nil)

		_, err := tCustom.MarshalJSON()
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("unknow type 'test' to encode"))
	})

	t.Run("It errors if the decoder does not know they type", func(t *testing.T) {
		tCustom := gotypedjson.NewTypedJsonDecoder(nil)

		err := json.Unmarshal([]byte(`{"Type":"Test","Value":"10"}`), tCustom)
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("unknown type 'Test' to decode"))
	})

	t.Run("Describe codec set on the typedJSON", func(t *testing.T) {
		codec := gotypedjson.CustomCodec{
			gotypedjson.INT: {
				Encode: func(val any) (string, error) {
					return fmt.Sprintf("%d", val.(int)+5), nil
				},
				Decode: func(s string) (any, error) {
					val, err := strconv.ParseInt(s, 10, 0)
					if err != nil {
						return nil, fmt.Errorf("failed to convert %s, to an int", s)
					}

					return int(val) - 5, nil
				},
			},
			gotypedjson.JSONTYPE("error"): {
				Encode: func(val any) (string, error) {
					return "", fmt.Errorf("error encoding")
				},
				Decode: func(s string) (any, error) {
					return "", fmt.Errorf("error decoding")
				},
			},
		}

		t.Run("It can use the custom encoder", func(t *testing.T) {
			tCustom := gotypedjson.NewTypedJson(gotypedjson.INT, 5, codec)

			data, err := tCustom.MarshalJSON()
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(`{"Type":"_int","Value":"10"}`))
		})

		t.Run("It forwards custom encoder errors", func(t *testing.T) {
			tCustom := gotypedjson.NewTypedJson(gotypedjson.JSONTYPE("error"), 5, codec)

			_, err := tCustom.MarshalJSON()
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("error encoding"))
		})

		t.Run("It can use the custom decoder", func(t *testing.T) {
			tCustom := gotypedjson.NewTypedJson(gotypedjson.INT, 0, codec)

			err := json.Unmarshal([]byte(`{"Type":"_int","Value":"10"}`), tCustom)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tCustom.Type).To(Equal(gotypedjson.INT))
			g.Expect(tCustom.Value.(int)).To(Equal(5))
		})

		t.Run("It forwards custom decoder errors", func(t *testing.T) {
			tCustom := gotypedjson.NewTypedJson(gotypedjson.JSONTYPE("error"), 0, codec)

			err := json.Unmarshal([]byte(`{"Type":"error","Value":"10"}`), tCustom)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("error decoding"))
		})
	})

	t.Run("Describe when the global codec is set", func(t *testing.T) {
		gotypedjson.GlobalCodec = gotypedjson.CustomCodec{
			gotypedjson.INT: {
				Encode: func(val any) (string, error) {
					return fmt.Sprintf("%d", val.(int)+5), nil
				},
				Decode: func(s string) (any, error) {
					val, err := strconv.ParseInt(s, 10, 0)
					if err != nil {
						return nil, fmt.Errorf("failed to convert %s, to an int", s)
					}

					return int(val) - 5, nil
				},
			},
			gotypedjson.JSONTYPE("error"): {
				Encode: func(val any) (string, error) {
					return "", fmt.Errorf("error encoding")
				},
				Decode: func(s string) (any, error) {
					return "", fmt.Errorf("error decoding")
				},
			},
		}

		defer func() {
			gotypedjson.GlobalCodec = nil
		}()

		t.Run("It can use the global encoder", func(t *testing.T) {
			tCustom := gotypedjson.NewTypedJson(gotypedjson.INT, 5, nil)

			data, err := tCustom.MarshalJSON()
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(`{"Type":"_int","Value":"10"}`))
		})

		t.Run("It forwards global encoder errors", func(t *testing.T) {
			tCustom := gotypedjson.NewTypedJson(gotypedjson.JSONTYPE("error"), 5, nil)

			_, err := tCustom.MarshalJSON()
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("error encoding"))
		})

		t.Run("It can use the global decoder", func(t *testing.T) {
			tCustom := gotypedjson.NewTypedJson(gotypedjson.INT, 0, nil)

			err := json.Unmarshal([]byte(`{"Type":"_int","Value":"10"}`), tCustom)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tCustom.Type).To(Equal(gotypedjson.INT))
			g.Expect(tCustom.Value.(int)).To(Equal(5))
		})

		t.Run("It forwards global decoder errors", func(t *testing.T) {
			tCustom := gotypedjson.NewTypedJson(gotypedjson.JSONTYPE("error"), 0, nil)

			err := json.Unmarshal([]byte(`{"Type":"error","Value":"10"}`), tCustom)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("error decoding"))
		})
	})
}

func Test_OmitEmptyPreserved(t *testing.T) {
	g := NewGomegaWithT(t)

	testStruct := struct {
		One *gotypedjson.TypedJson `json:"one,omitempty"`
		Two *gotypedjson.TypedJson `json:"two,omitempty"`
	}{}

	rawJSON := `{"one":{"Type":"_int","Value":"4"}}`

	g.Expect(json.Unmarshal([]byte(rawJSON), &testStruct)).ToNot(HaveOccurred())
	g.Expect(testStruct.One.Type).To(Equal(gotypedjson.INT))
	g.Expect(testStruct.One.Value.(int)).To(Equal(4))

	data, err := json.Marshal(testStruct)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(string(data)).To(Equal(rawJSON))
}
