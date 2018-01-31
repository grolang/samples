// +build ignore

package main

import (
	fmt "fmt"
	goother "github.com/grolang/samples/goother"
	basicfor05 "github.com/grolang/samples/grotest/basicfor05"
)

import myfloat "github.com/grolang/samples/grotest/generics/try05/myfloat"
import mystruct "github.com/grolang/samples/grotest/generics/try05/mystruct"

func init() {
	fmt.Println("'Hi' from try05.gro")
	goother.RunIt()
	basicfor05.RunIt()
	myfloat.RunIt()
	mystruct.RunIt()
}

func main() {}
