// +build ignore

package main

import (
	fmt "fmt"
)

import myint "github.com/grolang/samples/grotest/generics/try04/myint"

func init() {
	fmt.Println("'Hi' from try04.gro")
	myint.RunIt()
}

func main() {}
