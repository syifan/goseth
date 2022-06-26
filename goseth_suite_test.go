package goseth

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoseth(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goseth Suite")
}
