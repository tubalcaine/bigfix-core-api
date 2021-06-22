package main

import (
	"fmt"

	"github.com/tubalcaine/bigfix-core-api/bigfix"
)

// This is not meant to be used in the package. It is just a sample main
// that makes use of some of the bigfix module features as an example.

func main() {
	fmt.Println("Demo of bigfix module")
	fmt.Println(bigfix.Stub())
}
