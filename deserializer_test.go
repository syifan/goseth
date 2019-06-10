package goseth_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/syifan/goseth"
)

var _ = Describe("Deserializer", func() {
	var (
		d Deserializer
	)

	BeforeEach(func() {
		d = MakeDeserializer()
	})

	It("should deserilize bool", func() {
		str := `{
			"value": true,
			"type": "bool"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(BeTrue())
	})

	It("should deserilize int", func() {
		str := `{
			"value": 1,
			"type": "int"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(1))
	})

	It("should deserilize int8", func() {
		str := `{
			"value": 1,
			"type": "int8"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(int8(1)))
	})

	It("should deserilize int16", func() {
		str := `{
			"value": 1,
			"type": "int16"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(int16(1)))
	})

	It("should deserilize int32", func() {
		str := `{
			"value": 1,
			"type": "int32"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(int32(1)))
	})

	It("should deserilize int64", func() {
		str := `{
			"value": 1,
			"type": "int64"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(int64(1)))
	})

	It("should deserilize uint", func() {
		str := `{
			"value": 1,
			"type": "uint"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(uint(1)))
	})

	It("should deserilize uint8", func() {
		str := `{
			"value": 1,
			"type": "uint8"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(uint8(1)))
	})

	It("should deserilize uint16", func() {
		str := `{
			"value": 1,
			"type": "uint16"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(uint16(1)))
	})

	It("should deserilize uint32", func() {
		str := `{
			"value": 1,
			"type": "uint32"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(uint32(1)))
	})

	It("should deserilize uint64", func() {
		str := `{
			"value": 1,
			"type": "uint64"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(uint64(1)))
	})

	It("should deserilize float32", func() {
		str := `{
			"value": 0.1,
			"type": "float32"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(float32(0.1)))
	})

	It("should deserilize float64", func() {
		str := `{
			"value": 0.1,
			"type": "float64"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(float64(0.1)))
	})

	It("should deserialize struct", func() {
		// str := `{
		// 	"value": {
		// 		"Field1": {
		// 			"value": false,
		// 			"type": "bool"
		// 		},
		// 		"field2": {
		// 			"value": 1,
		// 			"type": "int"
		// 		}
		// 	},
		// 	"type": "github.com/syifan/seth_test.sampleStruct1"
		// }`

		// a, err := d.Deserialize(strings.NewReader(str))

		// Expect(err).To(BeNil())
		// exp := sampleStruct1{Field1: false, field2: 1}
		// Expect(a).To(Equal(exp))
	})

	// It("should deserialize pointer", func() {
	// 	str := `{
	// 		"value": 1234
	// 		"type": "*github.com/syifan/seth_test.sampleStruct1"
	// 		"dict": {
	// 			"1234": {
	// 				value:
	// 			}
	// 		}
	// 	}`

	// 	a, err := d.Deserialize(strings.NewReader(str))

	// 	Expect(err).To(BeNil())
	// 	Expect(a).To(Equal(float64(0.1)))

	// })
})
