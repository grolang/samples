// +build ignore

package main

import (
	listint "github.com/grolang/samples/container/generics/list_run/listint"
	"fmt"
)

func init() {
	{
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