package main

import "fmt"

//单行注释
//go定义变量
func test1() {
	var a int = 1
	a = 10
	a = a + 2
	fmt.Println(a)

	var b = 10
	fmt.Println(b)

	c := 20
	fmt.Println(c)
}

/*
	多行注释
	自动推导类型
*/
func test2() {
	PI := 3.14159
	r := 2.5
	//计算周长
	s := PI * r * 2

	fmt.Println("周长：", s)
}

//多重赋值
func test3() {
	a, b, c, d := 10, 'a', "wing", 2.34

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

//交换值
func test4() {
	a, b := 10, 23
	fmt.Println(a)
	fmt.Println(b)

	a, b = b, a
	fmt.Println(a)
	fmt.Println(b)
}

//匿名变量
func test5() {
	_, b, _, d := 1, 2, 3, 4
	//fmt.Println(_)
	fmt.Println(b)
	//fmt.Println(_)
	fmt.Println(d)
}

func main() {
	test1()
	//test2()
	//test3()
	//test4()
	//test5()
}
