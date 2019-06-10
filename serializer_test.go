package goseth_test

import (
	"regexp"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/syifan/goseth"
)

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
		exp := `{"value":6,"type":"int"}`

		err := s.Serialize(a, &sb)

		Expect(err).To(BeNil())
		Expect(sb.String()).To(Equal(exp))
	})

	It("should serialize int32", func() {
		a := int32(6)
		exp := `{"value":6,"type":"int32"}`

		err := s.Serialize(a, &sb)

		Expect(err).To(BeNil())
		Expect(sb.String()).To(Equal(exp))
	})

	It("should serialize float32", func() {
		a := float32(0.02)
		regex := regexp.MustCompile(
			`{"value":[0-9e\.]+,"type":"float32"}`)

		err := s.Serialize(a, &sb)

		Expect(err).To(BeNil())
		Expect(regex.Match([]byte(sb.String()))).To(BeTrue())
	})

	It("should serialize float64", func() {
		a := float64(0.0123456789)
		regex := regexp.MustCompile(
			`{"value":[0-9e\.]+,"type":"float64"}`)

		err := s.Serialize(a, &sb)

		Expect(err).To(BeNil())
		Expect(regex.Match([]byte(sb.String()))).To(BeTrue())
	})

	It("should serialize bool", func() {
		a := true
		exp := `{"value":true,"type":"bool"}`

		err := s.Serialize(a, &sb)

		Expect(err).To(BeNil())
		Expect(sb.String()).To(Equal(exp))
	})

	It("should serialize simple struct", func() {
		a := sampleStruct1{}
		exp := `{"value":{"Field1":{"value":false,"type":"bool"},"field2":{"value":0,"type":"int"}},"type":"goseth_test.sampleStruct1"}`

		err := s.Serialize(a, &sb)

		Expect(err).To(BeNil())
		Expect(sb.String()).To(Equal(exp))
	})

	It("should serialize pointer", func() {
		a := &sampleStruct1{}
		re := `{"value":[0-9]+,"type":"\*goseth_test.sampleStruct1","dict":{"[0-9]+":{"value":{"Field1":{"value":false,"type":"bool"},"field2":{"value":0,"type":"int"}},"type":"goseth_test.sampleStruct1"}}}`

		err := s.Serialize(a, &sb)

		Expect(err).To(BeNil())
		Expect(sb.String()).To(MatchRegexp(re))
	})

	It("should serialize recursive data", func() {
		s1 := sampleStruct2{}
		s2 := sampleStruct2{}
		s1.another = &s2
		s2.another = &s1
		re := `{"value":{"another":{"value":[0-9]+,"type":"\*goseth_test.sampleStruct2"}},"type":"goseth_test.sampleStruct2","dict":{"[0-9]+":{"value":{"another":{"value":[0-9]+,"type":"\*goseth_test.sampleStruct2"}},"type":"goseth_test.sampleStruct2"},"[0-9]+":{"value":{"another":{"value":[0-9]+,"type":"\*goseth_test.sampleStruct2"}},"type":"goseth_test.sampleStruct2"}}}`

		err := s.Serialize(s1, &sb)

		Expect(err).To(BeNil())
		Expect(sb.String()).To(MatchRegexp(re))
	})
})
