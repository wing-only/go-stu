package main

import "fmt"

func addrtest1() {
	var a int = 10
	fmt.Println(&a)		//0xc00000e0a8

	p := &a
	fmt.Println(p)		//0xc00000e0a8
	fmt.Println(*p)		//10
}

func main() {
	addrtest1()
}
