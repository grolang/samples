// +build ignore

package main

import (
	fmt "fmt"
)

import yourstruct "github.com/grolang/samples/grotest/generics/try09/yourstruct"

func init() {
	fmt.Println("'Hi' from try09.gro")
	yourstruct.RunIt()
}

func main() {}
