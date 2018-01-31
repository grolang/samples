package apex

import (
	fmt "fmt"
)

func PrtVal(x T) {
	fmt.Printf("The passed-in type is %T\n", x)
}

func PrtNil() {
	var x T
	fmt.Printf("The passed-in type is %T\n", x)
}
