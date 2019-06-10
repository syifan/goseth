package seth_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSerialization(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Seth Suite")
}

type sampleStruct1 struct {
	Field1 bool
	field2 int
}

type sampleStruct2 struct {
	another *sampleStruct2
}

func init() {
	reg := goseth.GetTypeRegistry()
	reg.Register((*sampleStruct1)(nil))
	reg.Register((*sampleStruct2)(nil))
}
