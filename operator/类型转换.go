package main

import "fmt"

func typetest1() {
	a := 10
	b := 3.99

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	//把int转换为float
	c := float64(a) * b
	fmt.Println(c)
}

func typetest2() {
	var a int32 = 10
	var b int64 = 20

	//c := a + b		//错误，类型不同不能运算
	c := int64(a) + b
	fmt.Println(c)
}
func main() {
	//typetest1()
	typetest2()
}
