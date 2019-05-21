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
		var a int
		str := `{
			"value": 1, 
			"t": "int"
		}`

		err := d.Deserialize(a, strings.NewReader(str))

		Expect(err).To(BeNil())
		Expect(a).To(Equal(1))
	})
})
