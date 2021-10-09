package main

import "fmt"

func main() {

	//  2.1 The very first "Hello, World" go program.
	fmt.Println("Hello, World!")

	//  2.2.1 variable.
	//  定义空变量 => 赋值.
	var a0 int
	a0 = 3
	//  定义，赋值simultaneously.
	var _ int = 4
	//  同时定义多个变量.
	var a1, a2, a3 int = 5, 6, 7
	//  另一种定义多个变量的方法：分组声明.
	//var (
	//	a1 = 5
	//	a2 = 6
	//	a3 = 7
	//)
	//  可以省略变量类型，compiler会根据等号右边的数据类型自动分配。
	var b1, b2, b3 = a1, a2, a3
	//  也可以省略关键字var，直接用:=，像Python和MATLAB那样。
	//  Notice: var搭配= 和 := 作用相等。
	_, a4 := 1988, 8

	//  boolean类型.
	var isGogood bool //  The default value of a bool variable is false.
	isGogood = true
	isGobad := false

	//  complex类型.
	var c0 complex128 = 2 + 2i
	c1 := 2 - 2i

	//  string.
	//  Notice that string is immutable.
	var hello string = "Hello"
	world := "World!"
	str := hello + " " + world

	//  One way for changing the characters in string.
	tmp := []byte(str) // transfer string to byte array.
	tmp[5] = '_'
	str = string(tmp) //  transfer byte array to string.

	//  多行字符串.
	//  ``括起的字符串为Raw字符串，原模原样print。
	message := ` This is Golang, 
				 next line.`

	//  2.2.2 constant.
	const Pi float64 = 3.1415926
	const MaxThread = 10

	//  2.2.3  iota.
	//  默认值为0，iota在同一行值相同，不同行默认加1.
	//  每遇到一个const关键字，iota就会重置为0.
	const (
		a5, a6 = iota, iota
		a7     = iota
		a8     = iota
	)
	const a9 = iota

	fmt.Println("a0 =", a0)
	fmt.Println("a1 + a2 + a3 =", a1+a2+a3)
	fmt.Println("b1 + b2 + b3 =", b1+b2+b3)
	fmt.Println("a4 =", a4)
	fmt.Println("Is Go good?", isGogood, "\n", "Is Go bad?", isGobad)
	fmt.Println("An example of multiplication of complex nums: ", c0*c1)
	fmt.Println("Again:", str)
	fmt.Println("%s\n", message)
	fmt.Println("Test iota:", a5, a6, a7, a8, a9)
}
