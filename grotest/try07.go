// +build ignore

package main

import (
	fmt "fmt"
)

import intpair "github.com/grolang/samples/grotest/generics/try07/intpair"

func init() {
	pr := intpair.List{3, 4}
	fmt.Printf("Hello, %v", pr)
}

func main() {}
