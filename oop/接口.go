package main

import "fmt"

/*
定义接口
接口中的方法必须有具体的实现
 */
type humaner interface {
	say()		//定义空方法
}

type per struct {
	name string
	age int
	sex string
}

type stud struct {
	per
	score int
}

type teac struct {
	per
	subject string
}

//实现接口say方法
func (s stud) say(){
	fmt.Println(s)		//{{小明 23 男} 81}
	fmt.Printf("我是%s,我是%s生，我今年%d岁了，我的成绩是%d分。\n", s.name, s.sex, s.age, s.score)
}
//实现接口say方法
func (t *teac) say()  {
	fmt.Println(t)		//&{{李红 33 女} golang}
	fmt.Printf("我是%s,我是%s生，我今年%d岁了，我教的学科是%s。\n", t.name, t.sex, t.age, t.subject)
}

//多态实现
//多态是将接口类型作为函数参数  多态实现接口的统一处理
func sayHello(h humaner)  {
	h.say()
}
func main() {
	stu := stud{
		per:   per{
			name: "小明",
			age:  23,
			sex:  "男",
		},
		score: 81,
	}

	tea := teac{
		per:     per{
			name: "李红",
			age:  33,
			sex:  "女",
		},
		subject: "golang",
	}

	var h humaner
	h = stu			//多态
	h.say()
	sayHello(h)		//多态


	h = &tea		//多态
	h.say()
	sayHello(h)		//多态
}
