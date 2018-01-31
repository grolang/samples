// +build ignore

package main

import (
	fmt "fmt"
)

import yourstruct "github.com/grolang/samples/grotest/generics/try06/yourstruct"

func init() {
	fmt.Println("'Hi' from try06.gro")
	yourstruct.RunIt()
}

func main() {}
