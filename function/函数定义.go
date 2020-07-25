package main

import "fmt"

//普通函数
func functest1(a int, b int) {
	sum := a + b
	fmt.Println(sum)
}

//带返回值函数
func functest2(a int, b int) int {
	return a + b
}

//可变参数
func functest3(args ...int) {
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}

func main() {
	//functest1(1, 2)
	//fmt.Println(functest2(1, 2))
	functest3(1, 2, 3, 'a', 'b')
}
