package main

import "fmt"

func test1() {
	fmt.Println("hello world1")
}
func test2() {
	//可以在程序中直接调用panic 调用后程序会终止执行
	panic("hello world2")
}
func test3() {
	fmt.Println("hello world3")
}

func main() {
	test1()
	test2()		//函数中报错
	test3()		//不会执行
}
