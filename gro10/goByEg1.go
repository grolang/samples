// +build ignore
// +build ignore
package main //包 is shorthand Kanji for `package`; 正 is short for `main`

import (
	_errors "errors" //入 is short for `import`
	fmt "fmt"
	math "math"
	time "time"
)

const _s string = "constant" //久 short for `const`

func main() { //功 short for `func`

	//Go by Example: Hello world
	//形 is short for `fmt.` which is automatically imported when used
	fmt.Println("你好,世界")

	//Go by Example: Values
	fmt.Println("go" + "lang") //Canonical "spaceless" Qu style: no spaces within statements and expressions
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false) //真 for `true`; 假 for `false`
	fmt.Println(true || false)
	fmt.Println(!true)

	//Go by Example: Variables
	{ //optionally, begin standalone block with `做`
		var _a string = "initial"
		fmt.Println(_a) //变 for `variable`; 串 for `string`
		var _b, _c int = 1, 2
		fmt.Println(_b, _c) //整 for `int`
		var _d = true
		fmt.Println(_d)
		var _e int
		fmt.Println(_e)
		_f := "short"
		fmt.Println(_f)
		//canonical Qu style: join stmts on one line with semicolons if they're written
		//as separate lines between blank lines in corresponding Go style
	}

	//Go by Example: Constants
	{
		fmt.Println(_s)
		const _n = 500000000
		const _d = 3e20 / _n
		fmt.Println(_d)
		fmt.Println(int64(_d))    //整64 for `int64`
		fmt.Println(math.Sin(_n)) //数 for package `math`
	}

	//Go by Example: For
	{
		_i := 1
		for _i <= 3 { //为 for `for`
			fmt.Println(_i)
			_i = _i + 1
		}
		for _j := 7; _j <= 9; _j++ {
			fmt.Println(_j)
		}
		for {
			fmt.Println("loop")
			break //破 for `break`
		}
	}

	//Go by Example: If/Else
	{
		if 7%2 == 0 { //如 for `if`
			fmt.Println("7 is even")
		} else { //否 for `else`
			fmt.Println("7 is odd")
		}
		if 8%4 == 0 {
			fmt.Println("8 is divisible by 4")
		}
		if _num := 9; _num < 0 {
			fmt.Println(_num, "is negative")
		} else if _num < 10 {
			fmt.Println(_num, "has 1 digit")
		} else {
			fmt.Println(_num, "has multiple digits")
		}
	}

	//Go by Example: Switch
	{
		_i := 2
		fmt.Print("write ", _i, " as ")
		switch _i { //考 for `switch`
		case 1:
			fmt.Println("one") //事 for `case`
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
		}

		switch time.Now().Weekday() { //时 for package `time`
		case time.Saturday, time.Sunday:
			fmt.Println("it's the weekend")
		default:
			fmt.Println("it's a weekday") //别 for `default`
		}

		_t := time.Now()
		switch {
		case _t.Hour() < 12:
			fmt.Println("it's before noon")
		default:
			fmt.Println("it's after noon")
		}
	}

	//Go by Example: Arrays
	{
		var _a [5]int
		fmt.Println("emp:", _a)
		_a[4] = 100
		fmt.Println("set:", _a)
		fmt.Println("get:", _a[4])
		fmt.Println("len:", len(_a)) //度 for `len`
		_b := [5]int{1, 2, 3, 4, 5}
		fmt.Println("dcl:", _b)

		var _twoD [2][3]int //BUG: semi needed here to compile
		for _i := 0; _i < 2; _i++ {
			for _j := 0; _j < 3; _j++ {
				_twoD[_i][_j] = _i + _j
			}
		}
		fmt.Println("2d: ", _twoD)
	}

	//Go by Example: Slices
	{
		_s := make([]string, 3)
		fmt.Println("emp:", _s) //造 for `make`
		_s[0] = "a"
		_s[1] = "b"
		_s[2] = "c"
		fmt.Println("set:", _s)
		fmt.Println("get:", _s[2])
		fmt.Println("len:", len(_s))
		_s = append(_s, "d")
		_s = append(_s, "e", "f")
		fmt.Println("apd:", _s) //加 for `append`
		_c := make([]string, len(_s))
		copy(_c, _s)
		fmt.Println("cpy:", _c) //副 for `copy`
		_l := _s[2:5]
		fmt.Println("sl1:", _l)
		_l = _s[:5]
		fmt.Println("sl2:", _l)
		_l = _s[2:]
		fmt.Println("sl3:", _l)
		_t := []string{"g", "h", "i"}
		fmt.Println("dcl:", _t)

		_twoD := make([][]int, 3)
		for _i := 0; _i < 3; _i++ {
			_innerLen := _i + 1
			_twoD[_i] = make([]int, _innerLen)
			for _j := 0; _j < _innerLen; _j++ {
				_twoD[_i][_j] = _i + _j
			}
		}
		fmt.Println("2d: ", _twoD)
	}

	//Go by Examples: Maps
	{
		_m := make(map[string]int) //图 for `map`
		_m["k1"] = 7
		_m["k2"] = 13
		fmt.Println("map:", _m)
		_v1 := _m["k1"]
		fmt.Println("v1: ", _v1)
		fmt.Println("len:", len(_m))

		delete(_m, "k2") //形 for `delete`
		fmt.Println("map:", _m)

		_, _prs := _m["k2"]
		fmt.Println("prs:", _prs)
		_n := map[string]int{"foo": 1, "bar": 2}
		fmt.Println("map:", _n)
	}

	//Go by Examples: Range
	{
		_nums := []int{2, 3, 4}
		_sum := 0
		for _, _num := range _nums { //围 for `range`
			_sum += _num
		}
		fmt.Println("sum:", _sum)

		for _i, _num := range _nums {
			if _num == 3 {
				fmt.Println("index:", _i)
			}
		}

		_kvs := map[string]string{"a": "apple", "b": "banana"}
		for _k, _v := range _kvs {
			fmt.Printf("%s -> %s\n", _k, _v)
		}

		for _i, _c := range "go" {
			fmt.Println(_i, _c)
		}
	}

	//Go by Example: Functions
	{
		_res := _plus(1, 2)
		fmt.Println("1+2 =", _res)
		_res = _plusPlus(1, 2, 3)
		fmt.Println("1+2+3 =", _res)
	}

	//standalone functions for above example
	//回 for `return`

	//Go by Example: Multiple Return Values
	{
		_a, _b := _vals()
		fmt.Println(_a)
		fmt.Println(_b)
		_, _c := _vals()
		fmt.Println(_c)
	}

	//Go by Example: Variadic Functions
	{
		_sum(1, 2)
		_sum(1, 2, 3)
		_nums := []int{1, 2, 3, 4}
		_sum(_nums...)
	}

	//Go by Example: Closures
	{
		_nextInt := _intSeq()
		fmt.Println(_nextInt())
		fmt.Println(_nextInt())
		fmt.Println(_nextInt())
		_newInts := _intSeq()
		fmt.Println(_newInts())
	}

	//Go by Example: Recursion
	fmt.Println(_fact(7))

	//Go by Example: Pointers
	{
		_i := 1
		fmt.Println("initial:", _i)
		_zeroval(_i)
		fmt.Println("zeroval:", _i)
		_zeroptr(&_i)
		fmt.Println("zeroptr:", _i)
		fmt.Println("pointer:", &_i)
	}

	//Go by Example: Structs
	{
		fmt.Println(_person{"Bob", 20})
		fmt.Println(_person{_name: "Alice", _age: 30})
		fmt.Println(_person{_name: "Fred"})
		fmt.Println(&_person{_name: "Ann", _age: 40})
		_s := _person{_name: "Sean", _age: 50}
		fmt.Println(_s._name)
		_sp := &_s
		fmt.Println(_sp._age)
		_sp._age = 51
		fmt.Println(_sp._age)
	}

	//种 is `type`; 构 is `struct`

	//Go by Example: Methods
	{
		_r := _rect{_width: 10, _height: 5}
		fmt.Println("area: ", _r._area())
		fmt.Println("perim:", _r._perim())
		_rp := &_r
		fmt.Println("area: ", _rp._area())
		fmt.Println("perim:", _rp._perim())
	}

	//space in `r rect` can't be eliminated for now

	//Go by Example: Interfaces
	{
		_r := _rectle{_width: 3, _height: 4}
		_c := _circle{_radius: 5}
		_measure(_r)
		_measure(_c)
	}

	//面 for `interface`
	//漂64 for `float64` //BUG: semi needed here to compile

	//Go by Example: Errors
	{
		for _, _i := range []int{7, 42} {
			if _r, _e := _f1(_i); _e != nil { //空 for `nil`
				fmt.Println("f1 failed:", _e)
			} else {
				fmt.Println("f1 worked:", _r)
			}
		}
		for _, _i := range []int{7, 42} {
			if _r, _e := _f2(_i); _e != nil {
				fmt.Println("f2 failed:", _e)
			} else {
				fmt.Println("f2 worked:", _r)
			}
		}
		_, _e := _f2(42)
		if _ae, _ok := _e.(*_argError); _ok {
			fmt.Println(_ae._arg)
			fmt.Println(_ae._prob)
		}
	}
}

func _plus(_a int, _b int) int     { return _a + _b }
func _plusPlus(_a, _b, _c int) int { return _a + _b + _c }

func _vals() (int, int) { return 3, 7 }

func _sum(_nums ...int) {
	fmt.Print(_nums, " ")
	_total := 0
	for _, _num := range _nums {
		_total += _num
	}
	fmt.Println(_total)
}

func _intSeq() func() int {
	_i := 0
	return func() int { _i += 1; return _i }
}

func _fact(_n int) int {
	if _n == 0 {
		return 1
	}
	return _n * _fact(_n-1)
}

func _zeroval(_ival int)  { _ival = 0 }
func _zeroptr(_iptr *int) { *_iptr = 0 }

type _person struct {
	_name string
	_age  int
}

type _rect struct{ _width, _height int }

func (_r *_rect) _area() int { return _r._width * _r._height }
func (_r _rect) _perim() int {
	return 2*_r._width + 2*_r._height
}

type _geometry interface {
	_area() float64
	_perim() float64
}

type _rectle struct{ _width, _height float64 }
type _circle struct{ _radius float64 }

func (_r _rectle) _area() float64  { return _r._width * _r._height }
func (_r _rectle) _perim() float64 { return 2*_r._width + 2*_r._height }

func (_c _circle) _area() float64  { return math.Pi * _c._radius * _c._radius }
func (_c _circle) _perim() float64 { return 2 * math.Pi * _c._radius }

func _measure(_g _geometry) {
	fmt.Println(_g)
	fmt.Println(_g._area())
	fmt.Println(_g._perim())
}

func _f1(_arg int) (int, error) {
	if _arg == 42 {
		return -1, _errors.New("can't work with 42")
	}
	return _arg + 3, nil
}

type _argError struct {
	_arg  int
	_prob string
}

func (_e *_argError) Error() string {
	return fmt.Sprintf("%d - %s", _e._arg, _e._prob)
}

func _f2(_arg int) (int, error) {
	if _arg == 42 {
		return -1, &_argError{_arg, "can't work with it"}
	}
	return _arg + 3, nil
}
