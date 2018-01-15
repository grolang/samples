// +build ignore

package main

import (
	assert "github.com/grolang/gro/assert"
	dyn "github.com/grolang/gro/ops"
)

import ops "github.com/grolang/gro/ops"

func init() {
	assert.AssertTrue(1 + 3 == 4)
	{
		assert.AssertTrue(3 + 4 == 7)
		assert.AssertTrue(9 - 2 == 7)
		assert.AssertTrue("string: abcdefg" == ops.Sprf("%T: %[1]s", "abcdefg"))
	}
	{
		assert.AssertTrue(dyn.IsEqual(dyn.Plus(4, 5), 9))
		assert.AssertTrue(dyn.IsEqual((dyn.And(true, func() interface{} {
			return false
		})), false))
		assert.AssertTrue(dyn.IsEqual(dyn.Negate(true), false))
		assert.AssertTrue(dyn.IsEqual(ops.Sprf(dyn.MakeText("%T: %[1]d"), dyn.Plus(4, 5)), dyn.MakeText("ops.Int: 9")))
		assert.AssertTrue(dyn.IsEqual(dyn.MakeText("utf88.Text: abcdefg"), ops.Sprf(dyn.MakeText("%T: %[1]s"), dyn.MakeText("abcdefg"))))
		assert.AssertTrue(dyn.IsEqual(dyn.MakeText("utf88.Codepoint: 'a'"), ops.Sprf(dyn.MakeText("%T: %[1]s"), dyn.Runex("a"))))
	}
}

func main() {}

type any = interface{}
