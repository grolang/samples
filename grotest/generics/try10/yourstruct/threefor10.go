package threefor10

import (
	fmt "fmt"
)

import myess "github.com/grolang/samples/grotest/generics/try10/yourstruct/generics/threefor10/threefor10/myess"

func RunIt() {
	var s S
	var t T
	fmt.Printf("'Hi' from try10.gro:threefor10(%T, %T).RunIt\n", s, t)
	myess.DoIt()
}
