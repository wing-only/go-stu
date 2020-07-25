package main

import "fmt"

/*
在函数外部定义结构体，作用域是全局的
*/

type Student struct {
	id   int
	name string
	sex  string
	age  int
}

/*
通过结构体定义变量
 */
func structtest1() {
	var s Student
	s.id = 1001
	s.name = "wing"
	s.sex = "男"
	s.age = 23

	fmt.Println(s)		//{1001 wing 男 23}


	var s2 Student = Student{
		id:   1002,
		name: "张飞",
		sex:  "男",
		age:  40,
	}
	fmt.Println(s2)		//{1002 张飞 男 40}
}

/*

 */
func structtest2() {

}

func structtest3() {

}

func structtest4() {

}

func structtest5() {

}


func main() {
	structtest1()
}
