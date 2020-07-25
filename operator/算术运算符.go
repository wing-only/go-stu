package main

import "fmt"

//除法
func pxprtest1() {
	a := 10
	b := 0
	c := a / b
	fmt.Println(c)
}

//取余
func exprtest2() {
	a := 10
	b := 2
	c := a % b
	fmt.Println(c)
}

//自增、自减
func exprtest3() {
	a := 10
	a++
	fmt.Println(a)

	a--
	fmt.Println(a)
}
func main() {
	//pxprtest1()
	//exprtest2()
	exprtest3()
}
