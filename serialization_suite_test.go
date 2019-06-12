package goseth_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/syifan/goseth"
)

func TestSerialization(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Seth Suite")
}

type sampleInterface interface{}

type sampleStruct1 struct {
	Field1 bool
	field2 int
}

type sampleStruct2 struct {
	another *sampleStruct2
}

type sampelStruct3 struct {
	s sampleStruct1
}

func init() {
	reg := goseth.GetTypeRegistry()
	reg.Register((*sampleStruct1)(nil))
	reg.Register((*sampleStruct2)(nil))
}
