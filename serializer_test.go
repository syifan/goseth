package goseth_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/syifan/goseth"
)

var idExp = `[0-9]+@[a-zA-Z_0-9\./]+`

var _ = Describe("Serializer Impl", func() {
	var (
		s  Serializer
		sb strings.Builder
	)

	BeforeEach(func() {
		s = MakeSerializer()
		sb = strings.Builder{}
	})

	It("should serialize int", func() {
		a := int(6)
		exp := `{"root":"` + idExp + `","dict":{"` + idExp + `":{"v":6,"t":"int","k":2}}`

		err := s.Serialize(&a, &sb)

		Expect(err).To(BeNil())
		Expect(sb.String()).To(MatchRegexp(exp))
	})

	It("should serialize int32", func() {
		a := int32(6)
		exp := `{"root":"` + idExp + `","dict":{"` + idExp + `":{"v":6,"t":"int32","k":5}}}`

		err := s.Serialize(&a, &sb)

		Expect(err).To(BeNil())
		Expect(sb.String()).To(MatchRegexp(exp))
	})

	It("should serialize float32", func() {
		a := float32(0.02)
		exp := `{"root":"` + idExp + `","dict":{"` + idExp + `":{"v":[0-9\.]+,"t":"float32","k":13}}}`

		err := s.Serialize(&a, &sb)

		Expect(err).To(BeNil())
		Expect(sb.String()).To(MatchRegexp(exp))
	})

	It("should serialize bool", func() {
		a := true
		exp := `{"root":"` + idExp + `","dict":{"` + idExp + `":{"v":true,"t":"bool","k":1}}}`

		err := s.Serialize(&a, &sb)

		Expect(err).To(BeNil())
		Expect(sb.String()).To(MatchRegexp(exp))
	})

	It("should serialize simple struct", func() {
		a := sampleStruct1{}
		exp := `{"root":"` + idExp + `","dict":{"` + idExp + `":{"v":{"Field1":"` + idExp + `","field2":"` + idExp + `"},"t":"github.com/syifan/goseth_test.sampleStruct1","k":25},"` + idExp + `":{"v":false,"t":"bool","k":1},"` + idExp + `":{"v":0,"t":"int","k":2}}`

		err := s.Serialize(&a, &sb)

		Expect(err).To(BeNil())
		Expect(sb.String()).To(MatchRegexp(exp))
	})

	// It("should serialize pointer", func() {
	// 	a := &sampleStruct1{}
	// 	re := `{"value":[0-9]+,"type":"\*goseth_test.sampleStruct1","dict":{"[0-9]+":{"value":{"Field1":{"value":false,"type":"bool"},"field2":{"value":0,"type":"int"}},"type":"goseth_test.sampleStruct1"}}}`

	// 	err := s.Serialize(a, &sb)

	// 	Expect(err).To(BeNil())
	// 	Expect(sb.String()).To(MatchRegexp(re))
	// })

	// It("should serialize recursive data", func() {
	// 	s1 := sampleStruct2{}
	// 	s2 := sampleStruct2{}
	// 	s1.another = &s2
	// 	s2.another = &s1
	// 	re := `{"value":{"another":{"value":[0-9]+,"type":"\*goseth_test.sampleStruct2"}},"type":"goseth_test.sampleStruct2","dict":{"[0-9]+":{"value":{"another":{"value":[0-9]+,"type":"\*goseth_test.sampleStruct2"}},"type":"goseth_test.sampleStruct2"},"[0-9]+":{"value":{"another":{"value":[0-9]+,"type":"\*goseth_test.sampleStruct2"}},"type":"goseth_test.sampleStruct2"}}}`

	// 	err := s.Serialize(s1, &sb)

	// 	Expect(err).To(BeNil())
	// 	Expect(sb.String()).To(MatchRegexp(re))
	// })
})
