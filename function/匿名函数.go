package main

import "fmt"

//匿名函数
func nonametest1() {
	f := func(a int, b int) {
		fmt.Println(a + b)
	}

	f(1, 3)
}

// 匿名自调函数
func nonametest2(a int, b int) {
	func(a int, b int) {
		fmt.Println(a + b)
	}(a, b)
}
func main() {
	//nonametest1()
	nonametest2(3, 4)
}
