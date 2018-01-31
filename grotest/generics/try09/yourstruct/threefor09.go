package threefor09

import (
	fmt "fmt"
)

import myque "github.com/grolang/samples/grotest/generics/try09/yourstruct/generics/threefor09/threefor09/myque"
import myess "github.com/grolang/samples/grotest/generics/try09/yourstruct/generics/threefor09/threefor09/myess"

func RunIt() {
	var q Q
	var r R
	var s S
	var t T
	fmt.Printf("'Hi' from try09.gro:threefor09(%T, %T, %T, %T).RunIt\n", q, r, s, t)
	myque.DoIt()
	myess.DoIt()
}
