package main

import "fmt"

func main() {
	//变量
	var a int
	var b int = 10
	var c int = a + b

	fmt.Println(a) //0
	fmt.Println(b) //10
	fmt.Println(c) //10

	//var PI float64  = 3.14159
	//var r float32 = 2.5

	//多重赋值
	a, b, c, d := 10, 20, 30, 40
	fmt.Println(a, b, c, d)

}
