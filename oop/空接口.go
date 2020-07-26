package main

import (
	"fmt"
)

/*
空接口  可以接收任意类型数据
*/

func kongtest1() {
	var i interface{}
	i = 10
	fmt.Println(i)         //10
	fmt.Printf("%p\n", &i) //0xc00002c1f0

	i = "hello world"
	fmt.Println(i)         //hello world
	fmt.Printf("%p\n", &i) //0xc00002c1f0

	var arr = [3]int{1, 2, 3}
	i = arr
	fmt.Println(i)         //[1 2 3]
	fmt.Printf("%p\n", &i) //0xc00002c1f0
}

/*
空接口切片，切片里面可以放任意类型数据
*/
func kongtest2() {
	var i []interface{}
	i = append(i, 1, 2, "hello", "我试试", [3]int{1, 2, 3})

	for idx, v := range i {
		fmt.Println(idx, v)
	}
}

func main() {
	//kongtest1()
	kongtest2()
}
