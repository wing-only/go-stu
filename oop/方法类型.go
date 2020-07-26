package main

import "fmt"

type person5 struct {
	name string
	age  int
	sex  string
}
type student5 struct {
	person5
	score int
}
type funcdemo func()

func (p *person5) PrintInfo() {
	fmt.Printf("大家好，我是%s，我今年%d岁，我是%s生\n", p.name, p.age, p.sex)
}

//方法重写  在一个对象中不能出现相同的方法名  方法的接收者 带* 和不带* 表示一个相同的对象
func (s *student5) PrintInfo(name string) {
	s.name = name
	fmt.Printf("大家好，我是%s，我今年%d岁，我是%s生,我的分数是%d分\n", s.name, s.age, s.sex, s.score)
}

func hello() {
	fmt.Println("hello")
}

func main() {
	s := student5{person5{"张三", 11, "男"}, 19}
	//代码区
	fmt.Println(s.PrintInfo)         //0x4a00f0
	fmt.Println(s.person5.PrintInfo) //0x4a0150
	fmt.Println(hello)               //0x49fce0


	var f1 funcdemo
	f1 = s.person5.PrintInfo
	fmt.Printf("%T\n", f1) //main.funcdemo


	f2 := s.PrintInfo
	fmt.Printf("%T\n", f2) //func(string)


	f3 := hello
	fmt.Printf("%T\n", f3) //func()
	f3()                   //hello
	a := 10
	fmt.Printf("%p\n", &a) //0xc00000e0e0


	//方法类型(函数类型) 变量
	f4 := s.PrintInfo
	f4("李二狗")
	fmt.Println(s)         //{{李二狗 11 男} 19}
	fmt.Printf("%T\n", f4) //func(string)

	//f5:=PrintInfo		//err

}
