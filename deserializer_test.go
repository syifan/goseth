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

	It("should deserilize int", func() {
		str := `{
			"value": 1,
			"t": "int"
		}`

		a, err := d.Deserialize(strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(1))
	})
})
