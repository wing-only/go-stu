package main

import "fmt"

//简单for循环
func fortest1() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

//外部变量
func fortest2() {
	i := 1
	for ; ; {
		if i > 10 {
			break
		}
		fmt.Println(i)
		i++
	}
}

//死循环
func fortest3() {
	for {
		fmt.Println("hello")
	}
}

func main() {
	//fortest1()
	//fortest2()
	fortest3()
}
