package main

import "fmt"

type person struct {
	id   int
	name string
}

type student struct {
	person			//通过匿名字段实现继承
	age int
	sex string
}

/*
通过匿名字段实现继承
同名字段 采用就进原则 使用子类信息
 */
func nonametest1() {
	var stu student = student{
		person: person{id: 1001, name: "wing"},
		age:    23,
		sex:    "男",
	}

	fmt.Println(stu)
}

//---------------------------- 多重继承 -----------------------------------

type testa struct {
	id int
	name string
}

type testb struct {
	testa
	age int
}

type testc struct {
	testa
	testb
	id int
	sex string
}

func nonametest2()  {
	var t testc = testc{
		testa: testa{
			id:   1001,
			name: "wing",
		},
		testb: testb{
			testa: testa{
				name: "张三",
			},
			age:   23,
		},
		id: 1003,
		sex:   "男",
	}

	fmt.Println(t)		//{{1001 wing} {{0 张三} 23} 1003 男}

	fmt.Println(t.id)	//1003
	fmt.Println(t.name)	//wing
	fmt.Println(t.age)	//23
	fmt.Println(t.sex)	//男
}
func main() {
	//nonametest1()
	nonametest2()
}
