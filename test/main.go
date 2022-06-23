package main

import (
	"fmt"
	"go/importer"
)

func main() {
	pkg, err := importer.Default().Import("github.com/onsi/ginkgo/v2")
	if err != nil {
		panic(err)
	}
	fmt.Println(pkg)
}
