package main

import "fmt"

//定义函数类型
type functype func(a int, b int) int

func functypetest1(a int, b int) (c int) {
	c = a + b
	return
}

func main() {
	var f functype			//函数类型变量
	f = functypetest1
	a := f(2, 3)
	fmt.Println(a)

}
