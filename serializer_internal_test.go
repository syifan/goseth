package goseth

import (
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Serializer", func() {
	var (
		s  Serializer
		sb strings.Builder
	)

	BeforeEach(func() {
		s = NewSerializer()
		sb = strings.Builder{}
	})

	Context("when serializing basic values", func() {
		It("should serialize boolean", func() {
			item := true

			s.SetRoot(item)

			Expect(s.Serialize(&sb)).To(Succeed())
			Expect(sb.String()).To(Equal(
				`{"r":"0","dict":{"0":{"k":1,"t":"bool","v":true}}}`,
			))
		})

		It("should serialize int", func() {
			item := int(1)

			s.SetRoot(item)

			Expect(s.Serialize(&sb)).To(Succeed())
			Expect(sb.String()).To(Equal(
				`{"r":"0","dict":{"0":{"k":2,"t":"int","v":1}}}`,
			))
		})

		It("should serialize uint", func() {
			item := uint(1)

			s.SetRoot(item)

			Expect(s.Serialize(&sb)).To(Succeed())
			Expect(sb.String()).To(Equal(
				`{"r":"0","dict":{"0":{"k":7,"t":"uint","v":1}}}`,
			))
		})

		It("should serialize float", func() {
			item := float64(1)

			s.SetRoot(item)

			Expect(s.Serialize(&sb)).To(Succeed())
			Expect(sb.String()).To(MatchRegexp(
				`{"r":"0","dict":{"0":{"k":14,"t":"float64","v":[0-9\.]+}}}`,
			))
		})

		It("should serialize string", func() {
			item := "abcde"

			s.SetRoot(item)

			Expect(s.Serialize(&sb)).To(Succeed())
			Expect(sb.String()).To(MatchRegexp(
				`{"r":"0","dict":{"0":{"k":24,"t":"string","v":"abcde"}}}`,
			))
		})

		It("should serialize nil pointer", func() {
			var a *int
			a = nil

			s.SetRoot(a)

			Expect(s.Serialize(&sb)).To(Succeed())

			Expect(sb.String()).To(Equal(
				`{"r":"0","dict":{"0":{"k":0,"t":"0","v":null}}}`,
			))
		})
	})

	Context("when serializing containers", func() {
		It("should serialize slice", func() {
			item := []int{1, 2, 3}

			s.SetRoot(item)

			Expect(s.Serialize(&sb)).To(Succeed())
			Expect(sb.String()).To(Equal(
				`{"r":"0","dict":{` +
					`"0":{"k":23,"t":"[]int","v":["1","2","3"]},` +
					`"1":{"k":2,"t":"int","v":1},` +
					`"2":{"k":2,"t":"int","v":2},` +
					`"3":{"k":2,"t":"int","v":3}}}`,
			))
		})

		It("should serialize map", func() {
			item := map[string]int{"a": 1, "b": 2}

			s.SetRoot(item)

			Expect(s.Serialize(&sb)).To(Succeed())

			Expect(sb.String()).To(MatchRegexp(
				`{"r":"0","dict":{` +
					`"0":{"k":21,"t":"map\[string\]int","v":{"1":"2","3":"4"}},` +
					`"1":{"k":24,"t":"string","v":"[a|b]"},` +
					`"2":{"k":2,"t":"int","v":[1|2]},` +
					`"3":{"k":24,"t":"string","v":"[a|b]"},` +
					`"4":{"k":2,"t":"int","v":[1|2]}}}`,
			))
		})

		It("should serialize struct", func() {
			type Foo struct {
				A int
				b string
			}

			item := Foo{A: 1, b: "abcde"}

			s.SetRoot(item)

			Expect(s.Serialize(&sb)).To(Succeed())

			Expect(sb.String()).To(MatchRegexp(
				`{"r":"0","dict":{` +
					`"0":{"k":25,"t":"github.com/syifan/goseth.Foo","v":{"A":"1","b":"2"}},` +
					`"1":{"k":2,"t":"int","v":1},` +
					`"2":{"k":24,"t":"string","v":"abcde"}}}`,
			))
		})
	})

	Context("when serializing indirect values", func() {
		It("should serialize pointer", func() {
			type Foo struct {
				A int
				b string
			}

			item := &Foo{A: 1, b: "abcde"}

			s.SetRoot(item)

			Expect(s.Serialize(&sb)).To(Succeed())

			Expect(sb.String()).To(MatchRegexp(
				`{"r":"0","dict":{` +
					`"0":{"k":25,"t":"github.com/syifan/goseth.Foo","v":{"A":"1","b":"2"}},` +
					`"1":{"k":2,"t":"int","v":1},` +
					`"2":{"k":24,"t":"string","v":"abcde"}}}`,
			))
		})
	})

	Context("when serializing with max depth", func() {
		It("should serialize limited depth", func() {
			type Foo struct {
				Nest *Foo
				A    int
			}

			item := Foo{A: 1}

			s.SetRoot(item)
			s.SetMaxDepth(0)

			Expect(s.Serialize(&sb)).To(Succeed())
			Expect(sb.String()).To(Equal(
				`{"r":"0","dict":{` +
					`"0":{"k":25,"t":"github.com/syifan/goseth.Foo"}}}`))
		})

		It("should serialize limited depth", func() {
			type Foo struct {
				A    int
				Nest *Foo
			}

			item := Foo{A: 1, Nest: &Foo{A: 2, Nest: &Foo{A: 3}}}

			s.SetRoot(item)
			s.SetMaxDepth(2)

			Expect(s.Serialize(&sb)).To(Succeed())
			Expect(sb.String()).To(Equal(
				`{"r":"0","dict":{` +
					`"0":{"k":25,"t":"github.com/syifan/goseth.Foo","v":{"A":"1","Nest":"2"}},` +
					`"1":{"k":2,"t":"int","v":1},` +
					`"2":{"k":25,"t":"github.com/syifan/goseth.Foo","v":{"A":"3","Nest":"4"}},` +
					`"3":{"k":2,"t":"int"},` +
					`"4":{"k":25,"t":"github.com/syifan/goseth.Foo"}` +
					`}}`,
			))
		})
	})

	Context("when using entry point", func() {
		It("should serialize entry point based on struct", func() {
			type Foo struct {
				A int
				b string
			}

			item := Foo{A: 1, b: "abcde"}

			s.SetRoot(item)
			s.SetEntryPoint([]string{"b"})

			Expect(s.Serialize(&sb)).To(Succeed())

			Expect(sb.String()).To(MatchRegexp(
				`{"r":"0","dict":{` +
					`"0":{"k":24,"t":"string","v":"abcde"}}}`,
			))
		})

		It("should serialize entry point based on slice", func() {
			type Foo struct {
				A []int
				b string
			}

			item := Foo{A: []int{1, 2, 3}, b: "abcde"}

			s.SetRoot(item)
			s.SetEntryPoint([]string{"A", "2"})

			Expect(s.Serialize(&sb)).To(Succeed())

			Expect(sb.String()).To(MatchRegexp(
				`{"r":"0","dict":{` +
					`"0":{"k":2,"t":"int","v":3}}}`,
			))
		})

	})
})
