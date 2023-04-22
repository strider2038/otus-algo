package main

import (
	"fmt"

	"github.com/bits-and-blooms/bloom/v3"
)

func main() {
	filter := bloom.NewWithEstimates(1000000, 0.01)
	filter.AddString("test")
	fmt.Println(filter.TestString("test"))
	fmt.Println(filter.TestString("empty"))
}
