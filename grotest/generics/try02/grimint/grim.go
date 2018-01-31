package grim

import (
	fmt "fmt"
)

func Yando(x T, f func(T)) {
	fmt.Printf("'Hi' from mycode.gro:grim(%T).Yando, saying: ", x)
	f(x)
}
