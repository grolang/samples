//we can use any Kanji with 口-radical on LHS...

//can use `这` with locally-defined type to achieve spaceless program

package main

import _fmt "fmt"
import "fmt"     //we can use any Kanji with 口-radical on LHS...
import _fg "fmt" //...with imports that don't have their own dedicated Kanji
import _hij "fmt"
import _utf8 "unicode/utf8"
import _kl "unicode/utf8" //can even put in two aliases

type A int

var _n = 50

func main() {
	if false {
		_fg.Printf("Len: %d\n", len("hijk")+_n)
	} else {
		_hij.Printf("Len: %d\n", len("hi")+_n)
	}
	_fr, _ := _utf8.DecodeRune([]byte("lmnop"))
	_fmt.Printf("1st rune: %s; Len: %d\n", string(_fr), len("lmnop")+_n)
	_, _ = _kl.DecodeRune([]byte("lmnop"))
	_hij.Printf("Fifty: %d\n", _n)
	_fg.Printf("Fifty: %d\n", _n)
	fmt.Printf("Fifty: %d\n", _n)
}
