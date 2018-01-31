// +build ignore

package main

import (
	reflect "reflect"
	groo "github.com/grolang/gro/ops"
	assert "github.com/grolang/gro/assert"
	time "time"
)

import ops "github.com/grolang/gro/ops"

var v int

const c = 2

type t int

func init() {
	{
		assert.AssertTrue(groo.IsEqual(groo.Mod(reflect.TypeOf(7), groo.MakeText("v")), groo.MakeText("int")))
		assert.AssertTrue(groo.IsEqual(groo.Mod(reflect.TypeOf(groo.Identity(7)), groo.MakeText("v")), groo.MakeText("int64")))
		assert.AssertTrue(groo.IsEqual(groo.Mod(reflect.TypeOf(groo.Negate(7)), groo.MakeText("v")), groo.MakeText("int64")))
		var a = 0x100000000
		assert.AssertTrue(groo.IsEqual(groo.Mod(reflect.TypeOf(groo.Identity(a)), groo.MakeText("v")), groo.MakeText("int64")))
		assert.AssertTrue(groo.IsEqual(groo.Mod(reflect.TypeOf(groo.Mult(a, a)), groo.MakeText("v")), groo.MakeText("ops.BigInt")))
		assert.AssertTrue(groo.IsEqual(groo.Mod(reflect.TypeOf(groo.Divide(7, 9)), groo.MakeText("v")), groo.MakeText("ops.BigRat")))
		assert.AssertTrue(groo.IsEqual(groo.Mod(reflect.TypeOf(1.1e250), groo.MakeText("v")), groo.MakeText("float64")))
		assert.AssertTrue(groo.IsEqual(groo.Mult(1.1e250, 1.1e250), inf))
	}
	assert.AssertTrue(groo.IsEqual(groo.Plus(1, 3), 4))
	assert.AssertTrue(groo.IsEqual(groo.Plus(3, 4), 7))
	assert.AssertTrue(groo.IsEqual(groo.Minus(9, 2), 7))
	assert.AssertTrue(groo.IsEqual(groo.Plus(4, 5), 9))
}

func retFloat() any {
	return 1.234
}

//unary reflect-operator
func init() {
	assert.AssertTrue(groo.IsEqual(retFloat(), 1.234))
	assert.AssertTrue(groo.IsEqual(groo.Divide(37, 23), groo.Divide(37, (groo.Plus(20, 3)))))
	assert.AssertTrue(groo.IsEqual(groo.Negate(0x80), groo.Minus(groo.Negate(0x40), 0x40)))
	assert.AssertTrue(groo.IsEqual(0xff, 0xff))
	assert.AssertTrue(groo.IsEqual(123456789, groo.Plus(groo.Plus(groo.Mult(123, 1000000), groo.Mult(456, 1000)), 789)))
	assert.AssertTrue(groo.IsEqual(groo.Mod(7, 4), 3))
	assert.AssertTrue(groo.IsEqual(groo.Plus(1.2, 3.4i), groo.Plus(1.2, 3.4i)))
	func() {
		assert.AssertTrue(groo.IsEqual((groo.And(true, func() interface{} {
			return false
		})), false))
		assert.AssertTrue(groo.IsEqual(groo.Negate(true), false))
		assert.AssertTrue(groo.IsEqual(groo.Negate(false), true))
		assert.AssertTrue(groo.IsEqual(groo.Negate(nil), nil))
	}()
	func() {
		assert.AssertTrue(groo.IsEqual((groo.And(true, func() interface{} {
			return false
		})), false))
	}()
	func() {
		assert.AssertTrue(groo.IsEqual(groo.Plus(false, false), false))
		assert.AssertTrue(groo.IsEqual(groo.Plus(true, false), true))
		assert.AssertTrue(groo.IsEqual(groo.Plus(false, true), true))
		assert.AssertTrue(groo.IsEqual(groo.Plus(true, true), true))
		assert.AssertTrue(groo.IsEqual(groo.Minus(false, false), true))
	}()
	func() {
		assert.AssertTrue(groo.IsEqual(groo.Mult(false, false), false))
		assert.AssertTrue(groo.IsEqual(groo.Mult(true, false), false))
		assert.AssertTrue(groo.IsEqual(groo.Mult(false, true), false))
		assert.AssertTrue(groo.IsEqual(groo.Mult(true, true), true))
	}()
	func() {
		assert.AssertTrue(groo.IsEqual(groo.Divide(1, nil), nil))
		assert.AssertTrue(groo.IsEqual(groo.Divide(1, true), false))
		assert.AssertTrue(groo.IsEqual(groo.Divide(1, false), true))
	}()
	func() {
		assert.AssertTrue(groo.IsEqual(groo.Divide(false, false), false))
		assert.AssertTrue(groo.IsEqual(groo.Divide(true, false), true))
		assert.AssertTrue(groo.IsEqual(groo.Divide(false, true), false))
		assert.AssertTrue(groo.IsEqual(groo.Divide(true, true), false))
	}()
	func() {
		assert.AssertTrue(groo.IsEqual(inf, inf))
		assert.AssertTrue(groo.IsEqual(groo.Divide(1, 0), inf))
		assert.AssertTrue(groo.IsEqual(groo.Divide(1, 1), 1))
		assert.AssertTrue(groo.IsEqual(0, 0))
		assert.AssertTrue(groo.IsEqual(groo.Divide(1, inf), 0))
		assert.AssertTrue(groo.IsEqual(groo.Divide(nil, nil), nil))
		assert.AssertTrue(groo.IsEqual(groo.Divide(0, 0), nil))
		assert.AssertTrue(groo.IsEqual(groo.Divide(inf, inf), nil))
	}()
	func() {
		assert.AssertTrue(groo.IsEqual(groo.Mod(4, 0), inf))
		assert.AssertTrue(groo.IsEqual(groo.Mod(4, inf), 4))
		assert.AssertTrue(groo.IsEqual(groo.Mod(0, 0), inf))
		assert.AssertTrue(groo.IsEqual(groo.Mod(inf, inf), inf))
	}()
	func() {
		assert.AssertTrue(groo.IsEqual(groo.Plus(time.Date(2017, 12, 12, 0, 0, 0, 0, time.UTC), 7), time.Date(2017, 12, 19, 0, 0, 0, 0, time.UTC)))
		assert.AssertTrue(groo.IsNotEqual(groo.Plus(time.Date(2017, 12, 12, 0, 0, 0, 0, time.UTC), 7), time.Date(2017, 12, 20, 0, 0, 0, 0, time.UTC)))
		assert.AssertTrue(groo.IsEqual(groo.Minus(time.Date(2017, 12, 12, 0, 0, 0, 0, time.UTC), 5), time.Date(2017, 12, 7, 0, 0, 0, 0, time.UTC)))
		assert.AssertTrue(groo.IsEqual(groo.Minus(time.Date(2017, 12, 12, 0, 0, 0, 0, time.UTC), 5), time.Date(2017, 12, 07, 0, 0, 0, 0, time.UTC)))
		assert.AssertTrue(groo.IsEqual(groo.Plus(7, time.Date(2017, 12, 12, 0, 0, 0, 0, time.UTC)), time.Date(2017, 12, 19, 0, 0, 0, 0, time.UTC)))
		assert.AssertTrue(groo.IsGreaterThan(time.Date(2017, 12, 12, 0, 0, 0, 0, time.UTC), time.Date(2003, 8, 29, 0, 0, 0, 0, time.UTC)))
	}()
	func() {
		assert.AssertTrue(groo.IsEqual(groo.NewPair(0, 5), groo.NewPair(0, 5)))
		assert.AssertTrue(groo.IsNotEqual(groo.NewPair(0, 5), groo.NewPair(2, 13)))
		assert.AssertTrue(groo.IsNotEqual(groo.NewPair(0, 5), groo.NewPair(0, 13)))
		assert.AssertTrue(groo.IsNotEqual(groo.NewPair(0, 5), groo.NewPair(2, 5)))
		assert.AssertTrue(groo.IsEqual(groo.NewPair(0, 20), groo.NewPair(0, 20)))
		assert.AssertTrue(groo.IsEqual(groo.NewPair(12, groo.Inf), groo.NewPair(12, inf)))
		assert.AssertTrue(groo.IsEqual(groo.NewPair(0, groo.Inf), groo.NewPair(0, inf)))
	}()
	assert.AssertTrue(groo.IsEqual(groo.Plus(groo.MakeText(""), groo.Mod(groo.Runex("\\t"), groo.MakeText("v"))), groo.MakeText("\t")))
	assert.AssertTrue(groo.IsEqual(groo.Plus(groo.MakeText(""), groo.Mod(groo.Runex("\\t"), groo.MakeText("v"))), groo.MakeText(`	`)))
	assert.AssertTrue(groo.IsEqual(groo.MakeText(`	`), groo.Mod(groo.Runex("\\t"), groo.MakeText("v"))))
	assert.AssertTrue(groo.IsEqual(ops.Sprf(groo.MakeText("Hey, you!")), groo.MakeText("Hey, you!")))
	assert.AssertTrue(groo.IsEqual(groo.Plus(groo.Plus(groo.Plus(groo.MakeText("Hello, "), groo.Mod(groo.MakeText("to"), groo.MakeText("s "))), groo.Mod(groo.MakeText("the"), groo.MakeText("s "))), groo.Mod(groo.MakeText("world"), groo.MakeText("s!"))), groo.MakeText("Hello, to the world!")))
	assert.AssertTrue(groo.IsEqual(groo.Plus(groo.MakeText("there are "), groo.Mod(6, groo.MakeText("d green bottles"))), groo.MakeText("there are 6 green bottles")))
	assert.AssertTrue(groo.IsEqual(groo.Mod(123, groo.MakeText("5d has type %[1]T.")), groo.MakeText("  123 has type int64.")))
	assert.AssertTrue(groo.IsEqual(groo.Plus(groo.Mod(123, groo.MakeText("5d has type %[1]T, and ")), groo.Mod(true, groo.MakeText("t is true."))), groo.MakeText("  123 has type int64, and true is true.")))
	func() {
		assert.AssertTrue(groo.IsEqual(groo.Mod(groo.Runex("a"), groo.MakeText("v")), groo.MakeText("a")))
		assert.AssertTrue(groo.IsEqual(groo.Mod(groo.Runex("ab"), groo.MakeText("v")), groo.MakeText("[ab]")))
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.RightShift(groo.MakeText("a"), groo.MakeText("a|b"))), groo.MakeText("v")), groo.MakeText("{{{ 0 1 a}}}")))
	}()
	func() {
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.Mult(groo.NewPair(0, groo.Inf), groo.Runex("a"))), groo.MakeText("v")), groo.MakeText(`a*`)))
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.Mult(groo.Runex("a"), groo.NewPair(0, groo.Inf))), groo.MakeText("v")), groo.MakeText(`a*?`)))
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.Mult(groo.NewPair(0, groo.Inf), groo.MakeText("a"))), groo.MakeText("v")), groo.MakeText(`(?:\Qa\E)*`)))
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.Mult(groo.MakeText("a"), groo.NewPair(0, groo.Inf))), groo.MakeText("v")), groo.MakeText(`(?:\Qa\E)*?`)))
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.Mult(groo.NewPair(0, groo.Inf), groo.Runex("a-z"))), groo.MakeText("v")), groo.MakeText(`[a-z]*`)))
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.Mult(groo.NewPair(0, groo.Inf), groo.MakeText("abc"))), groo.MakeText("v")), groo.MakeText(`(?:\Qabc\E)*`)))
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.Mult(groo.NewPair(0, groo.Inf), (groo.Alt(groo.Runex("a"), groo.Runex("b"))))), groo.MakeText("v")), groo.MakeText(`[ab]*`)))
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.Mult(groo.NewPair(0, groo.Inf), (groo.Alt(groo.Runex("a"), groo.MakeText("bc"))))), groo.MakeText("v")), groo.MakeText(`(?:a|(?:\Qbc\E))*`)))
	}()
	func() {
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.Mult(groo.NewPair(1, groo.Inf), groo.Runex("a"))), groo.MakeText("v")), groo.MakeText(`a+`)))
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.Mult(groo.Runex("a"), groo.NewPair(1, groo.Inf))), groo.MakeText("v")), groo.MakeText(`a+?`)))
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.Mult(groo.NewPair(0, 1), groo.Runex("a"))), groo.MakeText("v")), groo.MakeText(`a?`)))
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.Mult(3, groo.Runex("a"))), groo.MakeText("v")), groo.MakeText(`a{3}`)))
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.Mult(groo.Runex("a"), 3)), groo.MakeText("v")), groo.MakeText(`a{3}?`)))
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.Mult(groo.NewPair(3, 5), groo.Runex("a"))), groo.MakeText("v")), groo.MakeText(`a{3,5}`)))
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.Mult(groo.NewPair(3, groo.Inf), groo.Runex("a"))), groo.MakeText("v")), groo.MakeText(`a{3,}`)))
		assert.AssertTrue(groo.IsEqual(groo.Mod(groo.Not(groo.Runex("a")), groo.MakeText("v")), groo.MakeText(`[^a]`)))
		assert.AssertTrue(groo.IsEqual(groo.Mod(groo.Not(groo.Runex("2a-c")), groo.MakeText("v")), groo.MakeText(`[^2a-c]`)))
		assert.AssertTrue(groo.IsEqual(groo.Mod(groo.Not(groo.Runex("^2a-c")), groo.MakeText("v")), groo.MakeText(`[2a-c]`)))
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.Alt(groo.Runex("a"), groo.Runex("b"))), groo.MakeText("v")), groo.MakeText(`[ab]`)))
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.Alt(groo.Runex("a"), groo.MakeText("fg"))), groo.MakeText("v")), groo.MakeText(`a|(?:\Qfg\E)`)))
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.Alt(groo.Not(groo.Runex("a-e")), groo.Not(groo.Runex("g-j")))), groo.MakeText("v")), groo.MakeText(`[^a-eg-j]`)))
	}()
	p := ops.Spr
	func() {
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.Seq(groo.Runex("a"), groo.Runex("b"))), groo.MakeText("v")), groo.MakeText(`ab`)))
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.Seq(groo.Runex("a-e"), groo.MakeText("fgh"))), groo.MakeText("v")), groo.MakeText(`[a-e]fgh`)))
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.Seq((groo.Alt(groo.Runex("f"), groo.Runex("g"))), (groo.Alt(groo.Runex("J"), groo.Runex("k"))))), groo.MakeText("v")), groo.MakeText(`[fg][Jk]`)))
		assert.AssertTrue(groo.IsEqual(p(groo.Seq((groo.Alt(groo.Runex("f"), groo.Runex("g"))), (groo.Alt(groo.Runex("J"), groo.Runex("k"))))), groo.MakeText(`[fg][Jk]`)))
	}()
	func() {
		assert.AssertTrue(groo.IsEqual(groo.Mod(groo.Identity([]int{1, 2, 3}), groo.MakeText("T")), groo.MakeText("ops.Slice")))
		assert.AssertTrue(groo.IsEqual(groo.Mod([]int{1, 2, 3}, groo.MakeText("T")), groo.MakeText("ops.Slice")))
		assert.AssertTrue(groo.IsEqual(groo.MakeText("ops.Slice"), groo.Mod([]int{1, 2, 3}, groo.MakeText("T"))))
		assert.AssertTrue(groo.IsEqual(groo.Mod([]int{1, 2, 3}, groo.MakeText("v")), groo.MakeText("{1, 2, 3}")))
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.Plus([]int{1, 2, 3}, []any{groo.MakeText("4"), groo.MakeText("5")})), groo.MakeText("v")), groo.MakeText(`{1, 2, 3, 4, 5}`)))
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.LeftShift([]int{1, 2, 3}, groo.MakeText("456"))), groo.MakeText("v")), groo.MakeText(`{1, 2, 3, 456}`)))
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.LeftShift([]int{1, 2, 3}, []any{groo.MakeText("4"), groo.MakeText("5")})), groo.MakeText("v")), groo.MakeText(`{1, 2, 3, {4, 5}}`)))
	}()
	func() {
		assert.AssertTrue(groo.IsEqual(groo.LeftShift(groo.InitMap(groo.NewPair(11, groo.MakeText("Hey")), groo.NewPair(12, groo.MakeText("Wow")), groo.NewPair(13, groo.MakeText("Dude"))), groo.NewPair(14, groo.MakeText("Man"))), ops.InitMap(groo.NewPair(11, groo.MakeText("Hey")), groo.NewPair(12, groo.MakeText("Wow")), groo.NewPair(13, groo.MakeText("Dude")), groo.NewPair(14, groo.MakeText("Man")))))
		assert.AssertTrue(groo.IsEqual(groo.LeftShift(groo.InitMap(groo.NewPair(11, groo.MakeText("Hey")), groo.NewPair(12, groo.MakeText("Wow")), groo.NewPair(13, groo.MakeText("Dude"))), groo.NewPair(14, groo.MakeText("Man"))), groo.InitMap(groo.NewPair(11, groo.MakeText("Hey")), groo.NewPair(12, groo.MakeText("Wow")), groo.NewPair(13, groo.MakeText("Dude")), groo.NewPair(14, groo.MakeText("Man")))))
		assert.AssertTrue(groo.IsEqual(groo.Mod(groo.LeftShift(groo.InitMap(groo.NewPair(11, groo.MakeText("Hey")), groo.NewPair(12, groo.MakeText("Wow")), groo.NewPair(13, groo.MakeText("Dude"))), groo.NewPair(14, groo.MakeText("Man"))), groo.MakeText("v")), groo.MakeText("{11: Hey, 12: Wow, 13: Dude, 14: Man}")))
		assert.AssertTrue(groo.IsEqual(groo.Plus(groo.InitMap(groo.NewPair(11, groo.MakeText("Hey")), groo.NewPair(12, groo.MakeText("Wow")), groo.NewPair(13, groo.MakeText("Dude"))), groo.InitMap(groo.NewPair(14, groo.MakeText("Man")), groo.NewPair(12, groo.MakeText("NewWow")), groo.NewPair(8, groo.MakeText("OldMan")))), groo.InitMap(groo.NewPair(11, groo.MakeText("Hey")), groo.NewPair(13, groo.MakeText("Dude")), groo.NewPair(14, groo.MakeText("Man")), groo.NewPair(12, groo.MakeText("NewWow")), groo.NewPair(8, groo.MakeText("OldMan")))))
	}()
	func() {
		assert.AssertTrue(groo.IsEqual(groo.Mod(7, groo.MakeText("v %[1]T")), groo.MakeText("7 int64")))
		assert.AssertTrue(groo.IsEqual(groo.Mod(groo.Reflect(7), groo.MakeText("v %[1]T")), groo.MakeText("7 reflect.Value")))
		assert.AssertTrue(groo.IsEqual(groo.Mod(groo.Reflect(groo.Reflect(7)), groo.MakeText("v %[1]T")), groo.MakeText("7 int64")))
	}()
	{
		a := groo.Identity([]any{11, 12, 13})
		groo.LeftShiftAssign(&a, 14)
		assert.AssertTrue(groo.IsEqual(groo.Mod(a, groo.MakeText("v")), groo.MakeText("{11, 12, 13, 14}")))
	}
	{
		a := groo.Identity([]any{11, 12, 13})
		a = groo.LeftShift(a, 14)
		assert.AssertTrue(groo.IsEqual(groo.Mod(a, groo.MakeText("v")), groo.MakeText("{11, 12, 13, 14}")))
	}
	{
		a := groo.Identity([]any{groo.MakeText("Hey"), groo.MakeText("Wow"), groo.MakeText("Pal"), groo.MakeText("Mate"), groo.MakeText("Dude")})
		assert.AssertTrue(groo.IsEqual(groo.Mod((*groo.GetIndex(&a, 2)), groo.MakeText("v")), groo.MakeText("Pal")))
		assert.AssertTrue(groo.IsEqual(groo.Mod((*groo.GetIndex(&a, 2, 4)), groo.MakeText("v")), groo.MakeText("{Pal, Mate}")))
	}
	{
		a := groo.MakeText("Hey there!")
		assert.AssertTrue(groo.IsEqual(groo.Mod((*groo.GetIndex(&a, 4)), groo.MakeText("v")), groo.MakeText("t")))
		assert.AssertTrue(groo.IsEqual(groo.Mod((*groo.GetIndex(&a, 4)), groo.MakeText("v %[1]T")), groo.MakeText("t utf88.Codepoint")))
		assert.AssertTrue(groo.IsEqual(groo.Mod((*groo.GetIndex(&a, 4, 7)), groo.MakeText("v")), groo.MakeText("the")))
	}
	{
		x := (groo.RightShift(groo.MakeText("a"), groo.MakeText("a|b")))
		assert.AssertTrue(groo.IsEqual(groo.Mod((*groo.GetIndex(&(*groo.GetIndex(&x, 0)), 0)).(ops.RegexMatch).Match, groo.MakeText("v")), groo.MakeText("a")))
	}
	func() {
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.Mult(groo.MakeText("abc"), 3)), groo.MakeText("v %[1]T")), groo.MakeText(`(?:\Qabc\E){3}? ops.Regex`)))
	}()
	func() {
		assert.AssertTrue(groo.IsEqual(func(groo_it ...interface{}) interface{} {
			h := groo.MakeText("H")
			I := groo.MakeText("i")
			return groo.Plus(h, I)
		}(), groo.MakeText("Hi")))
		assert.AssertTrue(groo.IsEqual(groo.Mod(func(groo_it ...interface{}) interface{} {
			return groo.Mult(groo_it[0], groo_it[1])
		}(3, 4), groo.MakeText("d")), groo.MakeText("12")))
	}()
	func() {
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.RightShift(groo.MakeText("abc"), groo.Runex("a"))), groo.MakeText("v")), groo.MakeText("{{{ 0 1 a}}}")))
		assert.AssertTrue(groo.IsEqual(groo.Mod((groo.RightShift(groo.MakeText("bcd"), groo.Runex("a"))), groo.MakeText("v")), groo.MakeText("{}")))
	}()
	{
		{
			letterA := func(groo_it ...interface{}) interface{} {
				return groo.Runex("a")
			}
			assert.AssertTrue(groo.IsEqual(groo.RightShift(groo.MakeText("abc"), letterA), groo.Runex("a")))
			assert.AssertTrue(groo.IsEqual(groo.RightShift(groo.MakeText("bcd"), letterA), nil))
		}
	}
	{
		{
			letterA := func(groo_it ...interface{}) interface{} {
				return func(groo_it ...interface{}) interface{} {
					return groo.Runex("a")
				}
			}
			assert.AssertTrue(groo.IsEqual(groo.RightShift(groo.MakeText("abc"), letterA), groo.Runex("a")))
			assert.AssertTrue(groo.IsEqual(groo.RightShift(groo.MakeText("defg"), letterA), nil))
		}
	}
	{
		var expr func(...any) any
		paren := func(groo_it ...interface{}) interface{} {
			return groo.LeftSeq(groo.RightSeq(groo.MakeText("("), func(groo_it ...interface{}) interface{} {
				return expr
			}), groo.Runex(")"))
		} //call &> or <& on a func returning a func returning a Parser
		expr = func(groo_it ...interface{}) interface{} {
			return groo.Alt(groo.Runex("a"), paren())
		} //call | on a Parser
		func() {
			assert.AssertTrue(groo.IsEqual(groo.RightShift(groo.MakeText("a"), expr()), groo.Runex("a")))
			assert.AssertTrue(groo.IsEqual(groo.RightShift(groo.MakeText("(a)"), expr()), groo.Runex("a")))
			assert.AssertTrue(groo.IsEqual(groo.RightShift(groo.MakeText("((a))"), expr()), groo.Runex("a")))
		}()
	}
	{
		var expr func(...any) any
		paren := func(groo_it ...interface{}) interface{} {
			return groo.LeftSeq(groo.RightSeq(groo.MakeText("("), expr), groo.Runex(")"))
		} //call &> or <& on a func returning a Parser
		expr = func(groo_it ...interface{}) interface{} {
			return groo.Alt(groo.Runex("a"), paren())
		} //call | on a Parser
		func() {
			assert.AssertTrue(groo.IsEqual(groo.RightShift(groo.MakeText("a"), expr()), groo.Runex("a")))
			assert.AssertTrue(groo.IsEqual(groo.RightShift(groo.MakeText("(a)"), expr()), groo.Runex("a")))
			assert.AssertTrue(groo.IsEqual(groo.RightShift(groo.MakeText("((a))"), expr()), groo.Runex("a")))
		}()
	}
	{
		var expr func(...any) any
		paren := func(groo_it ...interface{}) interface{} {
			return groo.LeftSeq(groo.RightSeq(groo.MakeText("("), func(groo_it ...interface{}) interface{} {
				return expr
			}), groo.Runex(")"))
		}
		expr = func(groo_it ...interface{}) interface{} {
			return groo.Alt(groo.Runex("a"), paren)
		} //call | on func returning a Parser
		func() {
			assert.AssertTrue(groo.IsEqual(groo.RightShift(groo.MakeText("a"), expr()), groo.Runex("a")))
			assert.AssertTrue(groo.IsEqual(groo.RightShift(groo.MakeText("(a)"), expr()), groo.Runex("a")))
			assert.AssertTrue(groo.IsEqual(groo.RightShift(groo.MakeText("((a))"), expr()), groo.Runex("a")))
		}()
	}
	{
		var expr func(...any) any
		expr = func(groo_it ...interface{}) interface{} {
			return groo.Alt(groo.Runex("a"), (groo.LeftSeq(groo.RightSeq(groo.MakeText("("), func(groo_it ...interface{}) interface{} {
				return expr
			}), groo.Runex(")"))))
		}
		func() {
			assert.AssertTrue(groo.IsEqual(groo.RightShift(groo.MakeText("a"), expr()), groo.Runex("a")))
			assert.AssertTrue(groo.IsEqual(groo.RightShift(groo.MakeText("(a)"), expr()), groo.Runex("a")))
			assert.AssertTrue(groo.IsEqual(groo.RightShift(groo.MakeText("((a))"), expr()), groo.Runex("a")))
		}()
	}
	{
		var expr func(...any) any
		expr = func(groo_it ...interface{}) interface{} {
			return groo.Alt(groo.Runex("a"), (groo.LeftSeq(groo.RightSeq(groo.MakeText("("), expr), groo.Runex(")"))))
		} //call &> or <& on a func returning a Parser
		func() {
			assert.AssertTrue(groo.IsEqual(groo.RightShift(groo.MakeText("a"), expr()), groo.Runex("a")))
			assert.AssertTrue(groo.IsEqual(groo.RightShift(groo.MakeText("(a)"), expr()), groo.Runex("a")))
			assert.AssertTrue(groo.IsEqual(groo.RightShift(groo.MakeText("((a))"), expr()), groo.Runex("a")))
		}()
	}
	{
		var expr func(...any) any
		expr = func(groo_it ...interface{}) interface{} {
			return groo.Alt(groo.Runex("a"), (groo.LeftSeq(groo.RightSeq(groo.MakeText("("), func(groo_it ...interface{}) interface{} {
				return expr
			}), groo.Runex(")"))))
		}
		func() {
			assert.AssertTrue(groo.IsEqual(groo.RightShift(groo.MakeText("a"), expr), groo.Runex("a")))
		}()
	}
	{
		var expr func(...any) any
		expr = func(groo_it ...interface{}) interface{} {
			return groo.Alt(groo.Runex("a"), groo.Mult(groo.NewPair(2, 2), (groo.LeftSeq(groo.RightSeq(groo.MakeText("("), func(groo_it ...interface{}) interface{} {
				return expr
			}), groo.Runex(")")))))
		} //call * on a Parser
		func() {
			assert.AssertTrue(groo.IsEqual(groo.RightShift(groo.MakeText("aa"), expr()), groo.Runex("a")))
			assert.AssertTrue(groo.IsEqual(groo.Mod(groo.RightShift(groo.MakeText("(a)(a)"), expr()), groo.MakeText("v")), groo.MakeText("{a, a}")))
			assert.AssertTrue(groo.IsEqual(groo.Mod(groo.RightShift(groo.MakeText("((a)(a))((a)(a))"), expr()), groo.MakeText("v")), groo.MakeText("{{a, a}, {a, a}}")))
		}()
	}
	{
		var expr func(...any) any
		expr = func(groo_it ...interface{}) interface{} {
			return groo.Alt(groo.Runex("a"), groo.Mult(groo.NewPair(2, 2), func(groo_it ...interface{}) interface{} {
				return groo.LeftSeq(groo.RightSeq(groo.MakeText("("), func(groo_it ...interface{}) interface{} {
					return expr
				}), groo.Runex(")"))
			}))
		} //call * on func returning Parser
		func() {
			assert.AssertTrue(groo.IsEqual(groo.RightShift(groo.MakeText("aa"), expr()), groo.Runex("a")))
		}()
	}
	{
		{
			manyA := groo.Mult(groo.NewPair(3, 3), func(groo_it ...interface{}) interface{} {
				return groo.Runex("a")
			}) //call * on func returning codepoint
			func() {
				assert.AssertTrue(groo.IsEqual(groo.Mod(groo.RightShift(groo.MakeText("aaab"), manyA), groo.MakeText("v")), groo.MakeText("{a, a, a}")))
			}()
		}
	}
}

func main() {}

type (
	any = interface{}
	void = struct{}
)

var inf = groo.Inf

func init() {
	groo.UseUtf88 = true
}
