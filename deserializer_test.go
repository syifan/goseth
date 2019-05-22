package seth_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/syifan/seth"
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
			"t": "bool"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(BeTrue())
	})
	It("should deserilize int", func() {
		str := `{
			"value": 1,
			"t": "int"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(1))
	})

	It("should deserilize int8", func() {
		str := `{
			"value": 1,
			"t": "int8"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(int8(1)))
	})

	It("should deserilize int16", func() {
		str := `{
			"value": 1,
			"t": "int16"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(int16(1)))
	})

	It("should deserilize int32", func() {
		str := `{
			"value": 1,
			"t": "int32"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(int32(1)))
	})

	It("should deserilize int64", func() {
		str := `{
			"value": 1,
			"t": "int64"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(int64(1)))
	})

	It("should deserilize uint", func() {
		str := `{
			"value": 1,
			"t": "uint"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(uint(1)))
	})

	It("should deserilize uint8", func() {
		str := `{
			"value": 1,
			"t": "uint8"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(uint8(1)))
	})

	It("should deserilize uint16", func() {
		str := `{
			"value": 1,
			"t": "uint16"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(uint16(1)))
	})

	It("should deserilize uint32", func() {
		str := `{
			"value": 1,
			"t": "uint32"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(uint32(1)))
	})

	It("should deserilize uint64", func() {
		str := `{
			"value": 1,
			"t": "uint64"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(uint64(1)))
	})

	It("should deserilize float32", func() {
		str := `{
			"value": 0.1,
			"t": "float32"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(float32(0.1)))
	})

	It("should deserilize float64", func() {
		str := `{
			"value": 0.1,
			"t": "float64"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(float64(0.1)))
	})
})
