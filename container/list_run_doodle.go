// +build ignore

//line :8
package main

import (
//line :9
	listint "github.com/grolang/samples/container/generics/list_run_doodle/listint"
//line :10
	"fmt"
)

func init() {
	{
		gribbly.WhoAmI('?')
		l := listint.New()
		e4 := l.PushBack(4)
		e1 := l.PushFront(1)
		l.InsertBefore(3, e4)
		l.InsertAfter(2, e1)
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value)
		}
	}
}

func main() {}