package yoursfor06

import (
	fmt "fmt"
)

import mycomplex128 "github.com/grolang/samples/grotest/generics/yoursfor06/yoursfor06/mycomplex128"

func RunIt() {
	var s S
	var t T
	fmt.Printf("'Hi' from try06.gro:yourthree(%T, %T).RunIt\n", s, t)
	mycomplex128.DoIt()
}
