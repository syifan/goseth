package goseth_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/syifan/goseth"
)

var idExp = `[0-9]+@[a-zA-Z_0-9\./\[\]]+`

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
		a.Field1 = true
		exp := `{"root":"` + idExp + `","dict":{"` + idExp + `":{"v":{"Field1":"` + idExp + `","field2":"` + idExp + `"},"t":"github.com/syifan/goseth_test.sampleStruct1","k":25},"` + idExp + `":{"v":true,"t":"bool","k":1},"` + idExp + `":{"v":0,"t":"int","k":2}}`

		err := s.Serialize(&a, &sb)

		Expect(err).To(BeNil())
		Expect(sb.String()).To(MatchRegexp(exp))
	})

	It("should serialize recursive data", func() {
		s1 := sampleStruct2{}
		s2 := sampleStruct2{}
		s1.another = &s2
		s2.another = &s1
		re := `{"root":"` + idExp + `","dict":{"` + idExp + `":{"v":{"another":"` + idExp + `"},"t":"github.com/syifan/goseth_test.sampleStruct2","k":25},"` + idExp + `":{"v":{"another":"` + idExp + `"},"t":"github.com/syifan/goseth_test.sampleStruct2","k":25}}}`

		err := s.Serialize(&s1, &sb)

		Expect(err).To(BeNil())
		Expect(sb.String()).To(MatchRegexp(re))
	})

	It("should serialize slice", func() {
		d := []int{1, 2, 3}
		re := `{"root":"` + idExp + `","dict":{"` + idExp + `":["` + idExp + `","` + idExp + `","` + idExp + `"],"` + idExp + `":{"v":1,"t":"int","k":2},"` + idExp + `":{"v":2,"t":"int","k":2},"` + idExp + `":{"v":3,"t":"int","k":2}}}`

		err := s.Serialize(&d, &sb)

		Expect(err).To(BeNil())
		Expect(sb.String()).To(MatchRegexp(re))
	})

	It("should serialize pointer to interface", func() {
		s1 := &sampleStruct1{}
		var s2 sampleInterface
		s2 = s1
		re := `{"root":"` + idExp + `","dict":{"` + idExp + `":{"v":{"Field1":"` + idExp + `","field2":"` + idExp + `"},"t":"github.com/syifan/goseth_test.sampleStruct1","k":25},"` + idExp + `":{"v":false,"t":"bool","k":1},"` + idExp + `":{"v":0,"t":"int","k":2}}}`

		err := s.Serialize(&s2, &sb)

		Expect(err).To(BeNil())
		Expect(sb.String()).To(MatchRegexp(re))
	})

	It("should serialize nested structs", func() {
		b := sampelStruct3{s: sampleStruct1{Field1: true}}
		re := `{"root":"` + idExp + `","dict":{"` + idExp + `":{"v":{"s":"` + idExp + `"},"t":"github.com/syifan/goseth_test.sampelStruct3","k":25},"` + idExp + `":{"v":{"Field1":"` + idExp + `","field2":"` + idExp + `"},"t":"github.com/syifan/goseth_test.sampleStruct1","k":25},"` + idExp + `":{"v":true,"t":"bool","k":1},"` + idExp + `":{"v":0,"t":"int","k":2}}}`

		err := s.Serialize(&b, &sb)

		Expect(err).To(BeNil())
		Expect(sb.String()).To(MatchRegexp(re))
	})

})
