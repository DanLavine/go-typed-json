package gotypedjson_test

import (
	"encoding/json"
	"testing"

	gotypedjson "github.com/DanLavine/go-typed-json"
	. "github.com/onsi/gomega"
)

func Test_Int(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":0,"Value":"4"}`

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

			err := json.Unmarshal([]byte(`{"Type":0,"Value":"nope"}`), tInt)
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

	rawData := `{"Type":1,"Value":"4"}`

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

	t.Run("It can decode the value properly", func(t *testing.T) {
		tInt := &gotypedjson.TypedJson{}

		err := json.Unmarshal([]byte(rawData), tInt)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(tInt.Type).To(Equal(gotypedjson.INT8))
		g.Expect(tInt.Value.(int8)).To(Equal(int8(4)))
	})
}

func Test_Int16(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":2,"Value":"4"}`

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

	t.Run("It can decode the value properly", func(t *testing.T) {
		tInt := &gotypedjson.TypedJson{}

		err := json.Unmarshal([]byte(rawData), tInt)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(tInt.Type).To(Equal(gotypedjson.INT16))
		g.Expect(tInt.Value.(int16)).To(Equal(int16(4)))
	})
}

func Test_Int32(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":3,"Value":"4"}`

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

	t.Run("It can decode the value properly", func(t *testing.T) {
		tInt := &gotypedjson.TypedJson{}

		err := json.Unmarshal([]byte(rawData), tInt)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(tInt.Type).To(Equal(gotypedjson.INT32))
		g.Expect(tInt.Value.(int32)).To(Equal(int32(4)))
	})
}

func Test_Int64(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":4,"Value":"4"}`

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

	t.Run("It can decode the value properly", func(t *testing.T) {
		tInt := &gotypedjson.TypedJson{}

		err := json.Unmarshal([]byte(rawData), tInt)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(tInt.Type).To(Equal(gotypedjson.INT64))
		g.Expect(tInt.Value.(int64)).To(Equal(int64(4)))
	})
}

func Test_Uint(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":5,"Value":"4"}`

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

	t.Run("It can decode the value properly", func(t *testing.T) {
		tInt := &gotypedjson.TypedJson{}

		err := json.Unmarshal([]byte(rawData), tInt)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(tInt.Type).To(Equal(gotypedjson.UINT))
		g.Expect(tInt.Value.(uint)).To(Equal(uint(4)))
	})
}

func Test_Uint8(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":6,"Value":"4"}`

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

	t.Run("It can decode the value properly", func(t *testing.T) {
		tUint8 := &gotypedjson.TypedJson{}

		err := json.Unmarshal([]byte(rawData), tUint8)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(tUint8.Type).To(Equal(gotypedjson.UINT8))
		g.Expect(tUint8.Value.(uint8)).To(Equal(uint8(4)))
	})
}

func Test_Uint16(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":7,"Value":"4"}`

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

	t.Run("It can decode the value properly", func(t *testing.T) {
		tUint16 := &gotypedjson.TypedJson{}

		err := json.Unmarshal([]byte(rawData), tUint16)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(tUint16.Type).To(Equal(gotypedjson.UINT16))
		g.Expect(tUint16.Value.(uint16)).To(Equal(uint16(4)))
	})
}

func Test_Uint32(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":8,"Value":"4"}`

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

	t.Run("It can decode the value properly", func(t *testing.T) {
		tUint32 := &gotypedjson.TypedJson{}

		err := json.Unmarshal([]byte(rawData), tUint32)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(tUint32.Type).To(Equal(gotypedjson.UINT32))
		g.Expect(tUint32.Value.(uint32)).To(Equal(uint32(4)))
	})
}

func Test_Uint64(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":9,"Value":"4"}`

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

	t.Run("It can decode the value properly", func(t *testing.T) {
		tUint64 := &gotypedjson.TypedJson{}

		err := json.Unmarshal([]byte(rawData), tUint64)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(tUint64.Type).To(Equal(gotypedjson.UINT64))
		g.Expect(tUint64.Value.(uint64)).To(Equal(uint64(4)))
	})
}
