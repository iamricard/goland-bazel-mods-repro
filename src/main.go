package main

import (
	"fmt"

	"github.com/iamricard/goland-bazel-mods-repro/proto/foopb"
)

func main() {
	p := &foopb.Bar{
		Baz: "a string",
	}
	fmt.Println(p)
}
