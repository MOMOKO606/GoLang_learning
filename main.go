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
	//  可以省略变量类型，compiler会根据等号右边的数据类型自动分配。
	var b1, b2, b3 = a1, a2, a3
	//  也可以省略关键字var，直接用:=，像Python和MATLAB那样。
	_, a4 := 1988, 8
	//  boolean类型.
	var isGogood bool //  The default value of a bool variable is false.
	isGogood = true
	isGobad := false

	//  2.2.2 constant.
	const Pi float64 = 3.1415926
	const MaxThread = 10

	fmt.Println("a0 =", a0)
	fmt.Println("a1 + a2 + a3 =", a1+a2+a3)
	fmt.Println("b1 + b2 + b3 =", b1+b2+b3)
	fmt.Println("a4 =", a4)
	fmt.Println("Is Go good?", isGogood, "\n", "Is Go bad?", isGobad)
}
