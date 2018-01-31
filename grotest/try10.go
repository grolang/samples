// +build ignore

package main

import (
	fmt "fmt"
)

import yourstruct "github.com/grolang/samples/grotest/generics/try10/yourstruct"

func init() {
	fmt.Println("'Hi' from try10.gro")
	yourstruct.RunIt()
}

func main() {}
