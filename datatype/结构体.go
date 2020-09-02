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

	fmt.Println(s) //{1001 wing 男 23}

	var s2 Student = Student{
		id:   1002,
		name: "张飞",
		sex:  "男",
		age:  40,
	}
	fmt.Println(s2) //{1002 张飞 男 40}
}

/*
结构体变量比较
*/
func structtest2() {
	var s1 Student
	s1.id = 1001
	s1.name = "wing"
	s1.sex = "男"
	s1.age = 23

	fmt.Println(s1) //{1001 wing 男 23}

	var s2 Student = Student{
		id:   1001,
		name: "wing",
		sex:  "男",
		age:  23,
	}

	fmt.Println(s2) //{1001 wing 男 23}

	//比较两个结构体是否相等
	fmt.Println(s1 == s2)   //true
	fmt.Println(&s1 == &s2) //false
}

/*
结构体数组和切片
*/
func structtest3() {
	var arr [5]Student = [5]Student{{id: 1001, name: "张三", sex: "男", age: 23}, {id: 1001, name: "张三", sex: "男", age: 23}}

	fmt.Println(arr) //[{1001 张三 男 23} {1001 张三 男 23} {0   0} {0   0} {0   0}]

	s := []Student{{id: 1001, name: "张三", sex: "男", age: 23}, {id: 1001, name: "张三", sex: "男", age: 23}}
	fmt.Println(s)		//[{1001 张三 男 23} {1001 张三 男 23}]
}

/*
结构体作为函数参数
结构体作为函数参数  值传递
 */
func structtest4() {
	var s1 Student
	s1.id = 1001
	s1.name = "wing"
	s1.sex = "男"
	s1.age = 23

	fmt.Println("函数调用前：", s1) //{1001 wing 男 23}

	structtest5(s1)

	fmt.Println("函数调用后：", s1) //{1001 wing 男 23}
}

func structtest5(s1 Student) {
	s1.id = 1002
	s1.name = "张飞"
	fmt.Println("函数形参", s1) //{1002 张飞 男 23}
}

func main() {
	//structtest1()
	//structtest2()
	//structtest3()
	structtest4()
}
