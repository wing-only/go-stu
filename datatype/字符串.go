package main

import "fmt"

/*
string

 */
func strtest1() {
	var a string = "我爱学习"
	var b string = "我爱学习golang"

	//字符串相加，将两个字符串连起来
	c := a + b
	fmt.Println(c)

}

func strtest2() {
	var str = "中国r"

	//在go语言中一个汉字算作3个字符，为了和Linux同一处理
	num := len(str)
	fmt.Println(num) //7

}

func strtest3() {
	var a string = "hello\nworld"
	fmt.Println(a)
}

func main() {
	//strtest1()
	//strtest2()
	strtest3()
}
