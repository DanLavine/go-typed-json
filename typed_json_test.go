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

	t.Run("It can encode the value properly", func(t *testing.T) {
		tInt := &gotypedjson.TypedJson{Type: gotypedjson.INT, Value: 4}

		data, err := json.Marshal(tInt)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(string(data)).To(Equal(rawData))
	})

	t.Run("It can decode the value properly", func(t *testing.T) {
		tInt := &gotypedjson.TypedJson{}

		err := json.Unmarshal([]byte(rawData), tInt)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(tInt.Type).To(Equal(gotypedjson.INT))
		g.Expect(tInt.Value.(int)).To(Equal(4))
	})
}

func Test_Int8(t *testing.T) {
	g := NewGomegaWithT(t)

	rawData := `{"Type":1,"Value":"4"}`

	t.Run("It can encode the value properly", func(t *testing.T) {
		tInt8 := &gotypedjson.TypedJson{Type: gotypedjson.INT8, Value: int8(4)}

		data, err := json.Marshal(tInt8)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(string(data)).To(Equal(rawData))
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

	t.Run("It can encode the value properly", func(t *testing.T) {
		tInt := &gotypedjson.TypedJson{Type: gotypedjson.INT16, Value: int16(4)}

		data, err := json.Marshal(tInt)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(string(data)).To(Equal(rawData))
	})

	t.Run("It can decode the value properly", func(t *testing.T) {
		tInt := &gotypedjson.TypedJson{}

		err := json.Unmarshal([]byte(rawData), tInt)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(tInt.Type).To(Equal(gotypedjson.INT16))
		g.Expect(tInt.Value.(int16)).To(Equal(int16(4)))
	})
}
