package main

import "fmt"

type person1 struct {
	name string
	age int
	sex string
}

type rep struct {
	person1
	hobby string
}

type dep struct {
	person1
	worktime int
}

func (p person1) sayhello(){
	fmt.Printf("大家好，我是%s,我是%s生，我今年%d岁了。", p.name, p.sex, p.age)
}

func (r rep) sayhi() {
	r.sayhello()
	fmt.Println("我的爱好是", r.hobby)
}

func (d dep) sayhi()  {
	d.sayhello()
	fmt.Println("我的工作年限是", d.worktime)
}

func main() {
	var r = rep{
		person1: person1{
			name: "张飞",
			age:  23,
			sex:  "男",
		},
		hobby:   "拍电影",
	}

	r.sayhi()

	var d = dep{
		person1:  person1{
			name: "刘备",
			age:  30,
			sex:  "男",
		},
		worktime: 3,
	}
	d.sayhi()
}
