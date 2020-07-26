package main

import "fmt"

type stu struct {
	name string
	age int
	sex string
}

//方法 值传递
func (s stu) printinfo()  {
	s.name = "flank"
	s.age = 99
	s.sex = "未知"

	fmt.Println(s.name)
	fmt.Println(s.age)
	fmt.Println(s.sex)
}

//方法 地址传递
func (s *stu) editinfo(name string, age int, sex string){
	s.name = name
	s.age = age
	s.sex = sex
}

func main() {
	s := stu{
		name: "wing",
		age:  23,
		sex:  "男",
	}

	s.printinfo()		//flank 99 未知
	fmt.Println(s)		//{wing 23 男}

	s.editinfo("李雷", 24, "女")
	fmt.Println(s)		//{李雷 24 女}

}
