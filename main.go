//  Golang的一些规则：
//	大写字母开头的变量是可导出的，也就是其它包可以读取的，是公有变量；小写字母开头的就是不可导出的，是私有变量。
//	大写字母开头的函数也是一样，相当于class中的带public关键词的公有函数；小写字母开头的就是有private关键词的私有函数。

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

	//  2.2.4  array.
	//  Method 1, declare => assign values
	var arr01 [5]float64
	arr01[1] = 100
	//  Method 2. declare & assign simultaneously.
	//  定义了一个长度为8，并给前5个数赋初值，其他值默认为0。
	arr02 := [8]int{0, 1, 2, 3, 4}
	arr03 := [...]float64{11, 12, 13, 16, 23, 26, 28, 30} //  省略元素长度，Go会自动根据元素个数来计算长度。

	//  二维数组
	arr04 := [...][2]float64{[2]float64{1, 3}, [2]float64{2, 4}}
	//  上式可以简化为：
	//  arr04 := [...][2]float64{{1,3}, {2,4}}

	//  2.2.5  slice "动态数组"
	var s01, s02 []float64
	s01 = arr03[0:3]
	s02 = arr03[4:6]
	//  Test for slice, slice's cap & array.
	//  对array切片后得到slice，slice对应的cap是它的起点到其original array结尾的长度。
	var arr05 [10]int
	s03 := arr05[2:4]   // s03.cap为从arr05的index 2到9的长度，8.
	s04 := arr05[2:4:7] //  指定cap到index=7但不包含7 => cap = 5

	fmt.Println("a0 =", a0)
	fmt.Println("a1 + a2 + a3 =", a1+a2+a3)
	fmt.Println("b1 + b2 + b3 =", b1+b2+b3)
	fmt.Println("a4 =", a4)
	fmt.Println("Is Go good?", isGogood, "\n", "Is Go bad?", isGobad)
	fmt.Println("An example of multiplication of complex nums: ", c0*c1)
	fmt.Println("Again:", str)
	fmt.Println("%s\n", message)
	fmt.Println("Test iota:", a5, a6, a7, a8, a9)
	fmt.Println("Show the array arr02: ", arr02)
	fmt.Println("Show the array arr03: ", arr03)
	fmt.Println("Show the 2D array arr04: ", arr04)
	fmt.Println("Show the slice s01 & s02: ", s01, s02)
	fmt.Println("Show the cap of s03: ", cap(s03))
	fmt.Println("Show the cap of s04: ", cap(s04))
}
