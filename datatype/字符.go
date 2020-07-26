package main

import "fmt"

func tunetest1()  {
	var a byte = '0'
	var b byte = 'a'

	//打印字符型对应的ASCII值
	fmt.Println(a)
	fmt.Println(b)

	//打印字符
	fmt.Printf("%c, %c", a, b)
}

func main() {
	tunetest1()
}
