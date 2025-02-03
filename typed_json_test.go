package gotypedjson_test

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

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

func Test_Int(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":1,"Value":"4"}`

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

			err := json.Unmarshal([]byte(`{"Type":1,"Value":"nope"}`), tInt)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert nope, to an int"))
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

	rawData := `{"Type":2,"Value":"4"}`

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

			err := json.Unmarshal([]byte(`{"Type":2,"Value":"nope"}`), tInt8)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert nope, to an int8"))
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

	rawData := `{"Type":3,"Value":"4"}`

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

			err := json.Unmarshal([]byte(`{"Type":3,"Value":"nope"}`), tInt16)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert nope, to an int16"))
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

	rawData := `{"Type":4,"Value":"4"}`

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

			err := json.Unmarshal([]byte(`{"Type":4,"Value":"nope"}`), tInt32)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert nope, to an int32"))
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

	rawData := `{"Type":5,"Value":"4"}`

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

			err := json.Unmarshal([]byte(`{"Type":5,"Value":"nope"}`), tInt64)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert nope, to an int64"))
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

	rawData := `{"Type":6,"Value":"4"}`

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

			err := json.Unmarshal([]byte(`{"Type":6,"Value":"nope"}`), tUint)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert nope, to a uint"))
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

	rawData := `{"Type":7,"Value":"4"}`

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

			err := json.Unmarshal([]byte(`{"Type":7,"Value":"nope"}`), tUint8)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert nope, to a uint8"))
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

	rawData := `{"Type":8,"Value":"4"}`

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

			err := json.Unmarshal([]byte(`{"Type":8,"Value":"nope"}`), tUint16)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert nope, to a uint16"))
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

	rawData := `{"Type":9,"Value":"4"}`

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

			err := json.Unmarshal([]byte(`{"Type":9,"Value":"nope"}`), tUint32)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert nope, to a uint32"))
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

	rawData := `{"Type":10,"Value":"4"}`

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

			err := json.Unmarshal([]byte(`{"Type":10,"Value":"nope"}`), tUint64)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert nope, to a uint64"))
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

	rawData := `{"Type":11,"Value":"4E+00"}`

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

			err := json.Unmarshal([]byte(`{"Type":11,"Value":"nope"}`), tFloat32)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert nope, to a float32"))
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

	rawData := `{"Type":12,"Value":"4E+00"}`

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

			err := json.Unmarshal([]byte(`{"Type":12,"Value":"nope"}`), tFloat64)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert nope, to a float64"))
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

	rawData := `{"Type":13,"Value":"proper string"}`

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

			err := json.Unmarshal([]byte(`{"Type":13,"Value":3}`), tString)
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

	rawData := `{"Type":14,"Value":"true"}`

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

			err := json.Unmarshal([]byte(`{"Type":14,"Value":"nope"}`), tBool)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(Equal("failed to convert nope, to a bool"))
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

func Test_CustomCodec(t *testing.T) {
	g := NewGomegaWithT(t)

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
		}

		t.Run("It can use the custom encoder", func(t *testing.T) {
			tCustom := gotypedjson.NewTypedJson(gotypedjson.INT, 5, codec)

			data, err := tCustom.MarshalJSON()
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(`{"Type":1,"Value":"10"}`))
		})

		t.Run("It can use the custom decoder", func(t *testing.T) {
			tCustom := gotypedjson.NewTypedJson(gotypedjson.INT, 0, codec)

			err := json.Unmarshal([]byte(`{"Type":1,"Value":"10"}`), tCustom)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tCustom.Type).To(Equal(gotypedjson.INT))
			g.Expect(tCustom.Value.(int)).To(Equal(5))
		})

		t.Run("Context with GlobalCodec set", func(t *testing.T) {
			gotypedjson.GlobalCodec = gotypedjson.CustomCodec{
				gotypedjson.INT: gotypedjson.Codec{
					Encode: func(val any) (string, error) { return "", fmt.Errorf("fail") },
					Decode: func(s string) (any, error) { return nil, fmt.Errorf("fail") },
				},
			}

			defer func() {
				gotypedjson.GlobalCodec = nil
			}()

			t.Run("It uses the TypedJson codec", func(t *testing.T) {
				tCustom := gotypedjson.NewTypedJson(gotypedjson.INT, 0, codec)

				err := json.Unmarshal([]byte(`{"Type":1,"Value":"10"}`), tCustom)
				g.Expect(err).ToNot(HaveOccurred())
				g.Expect(tCustom.Type).To(Equal(gotypedjson.INT))
				g.Expect(tCustom.Value.(int)).To(Equal(5))
			})
		})
	})

	t.Run("Dexcribe when the global codec is set", func(t *testing.T) {
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
		}

		defer func() {
			gotypedjson.GlobalCodec = nil
		}()

		t.Run("It can use the global encoder", func(t *testing.T) {
			tCustom := gotypedjson.NewTypedJson(gotypedjson.INT, 5, nil)

			data, err := tCustom.MarshalJSON()
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(data)).To(Equal(`{"Type":1,"Value":"10"}`))
		})

		t.Run("It can use the global decoder", func(t *testing.T) {
			tCustom := gotypedjson.NewTypedJson(gotypedjson.INT, 0, nil)

			err := json.Unmarshal([]byte(`{"Type":1,"Value":"10"}`), tCustom)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tCustom.Type).To(Equal(gotypedjson.INT))
			g.Expect(tCustom.Value.(int)).To(Equal(5))
		})
	})
}

func Test_OmitEmptyPreserved(t *testing.T) {
	g := NewGomegaWithT(t)

	testStruct := struct {
		One *gotypedjson.TypedJson `json:"one,omitempty"`
		Two *gotypedjson.TypedJson `json:"two,omitempty"`
	}{}

	rawJSON := `{"one":{"Type":1,"Value":"4"}}`

	g.Expect(json.Unmarshal([]byte(rawJSON), &testStruct)).ToNot(HaveOccurred())
	g.Expect(testStruct.One.Type).To(Equal(gotypedjson.INT))
	g.Expect(testStruct.One.Value.(int)).To(Equal(4))

	data, err := json.Marshal(testStruct)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(string(data)).To(Equal(rawJSON))
}
