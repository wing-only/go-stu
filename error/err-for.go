package main

import "fmt"

/*
for循环里 处理异常， 需要把循环体放到函数里面
*/
func errfortest1() {
	for i := 0; i < 5; i++ {
		errfortest2()
	}
}

func errfortest2() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("异常了。catch")
		}

		fmt.Println("最终执行。 finally")
	}()

	a := 2
	b := 0
	c := a / b
	fmt.Println(c)
}

func main() {
	errfortest1()
}
