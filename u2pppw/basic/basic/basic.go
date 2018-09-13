package main

import (
	"fmt"
	"math"
	"math/cmplx"
)


//在函数外必须有关键字var   b := 3 这种初始化是不行的
var (
	aa = 3
	ss = "kkk"
	bb = true
)//包内部变量  不是全局变量  作用域是包内部

func variableZeroValue() {  //变量定义后是有初始值的
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	c := 3 + 4i
	fmt.Println(c)
	fmt.Println(cmplx.Abs(c))  //取模也就是绝对值
	//---------------------------------欧拉公式--------
	fmt.Println(cmplx.Pow(math.E,1i * math.Pi)+1)
	fmt.Println(cmplx.Exp(1i * math.Pi)+1)
	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi)+1)
	//------------------------------------------------

}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota   //iota 表示自增值的
		_
		python
		golang
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)

	euler()
	triangle()
	consts()
	enums()
}
