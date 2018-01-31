// +build ignore

package main

import (
	fmt "fmt"
	grub "github.com/grolang/samples/grotest/grub"
)

import grimint "github.com/grolang/samples/grotest/generics/try02/grimint"

func init() {
	fmt.Println("'Hi' from mycode.gro")
	grimint.Yando(789, grub.Hooch)
}

func main() {}
