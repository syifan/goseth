package seth_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/syifan/seth"
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
	reg := seth.GetTypeRegistry()
	reg.Register((*sampleStruct1)(nil))
	reg.Register((*sampleStruct2)(nil))
}
