// +build ignore

package main

import (
	ringint "github.com/grolang/samples/container/generics/ring_run/ringint"
)

func init() {
	{
		var (
			r0 *ringint.Ring
			r1 ringint.Ring
		)
		ringint.Verify(r0, 0, 0)
		ringint.Verify(&r1, 1, 0)
		r1.Link(r0)
		ringint.Verify(r0, 0, 0)
		ringint.Verify(&r1, 1, 0)
		r1.Link(r0)
		ringint.Verify(r0, 0, 0)
		ringint.Verify(&r1, 1, 0)
		r1.Unlink(0)
		ringint.Verify(&r1, 1, 0)
	}
}

func main() {}
