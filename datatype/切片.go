package main

import (
	"fmt"
	"strconv"
)

//切片定义
func slicetest1() {
	//定义切片
	var s []int
	s = []int{1, 2}
	fmt.Println(s)

	//定义切片 并赋值
	var s1 = []int{1, 2, 3}
	fmt.Println(s1)

	//make函数定义切片
	s2 := make([]int, 4)
	s2[0] = 1
	s2[1] = 2
	fmt.Println(s2)
}

/*
切片的 操作
len()		//长度
cap()		//容量
append()	//添加数据

容量大于等于长度
容量扩展：如果整体数据没有超过1024字节 每次扩展为上一次的倍数  超过1024 每次扩展上一次的1/4
*/
func slicetest2() {
	s1 := make([]int, 5)
	fmt.Println("切片：", s1)      //[0 0 0 0 0]
	fmt.Println("长度：", len(s1)) //5
	fmt.Println("容量：", cap(s1)) //5

	s1 = append(s1, 1, 2, 3, 4)
	fmt.Println("切片：", s1)      //[0 0 0 0 0 1 2 3 4]
	fmt.Println("长度：", len(s1)) //9
	fmt.Println("容量：", cap(s1)) //10
}

/*
切片截取
切片名[起始下标:]
切片名[:结束位置] 不包含结束位置
切片名[起始位置:结束位置]
切片名[起始位置:结束位置:容量]
切片名[:]获取切片中所有元素
*/
func slicetest3() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(s)             //[0 1 2 3 4 5 6 7]
	fmt.Println("长度：", len(s)) //8
	fmt.Println("容量：", cap(s)) //8

	s1 := s[2:]
	fmt.Println(s1)             //[2 3 4 5 6 7]
	fmt.Println("长度：", len(s1)) //6
	fmt.Println("容量：", cap(s1)) //6

	s1 = s[:2]
	fmt.Println(s1)             //[0 1]
	fmt.Println("长度：", len(s1)) //2
	fmt.Println("容量：", cap(s1)) //8

	s1 = s[2:5]
	fmt.Println(s1)             //[2 3 4]
	fmt.Println("长度：", len(s1)) //3
	fmt.Println("容量：", cap(s1)) //6

	s1 = s[4:5:7]
	fmt.Println(s1)             //[4]
	fmt.Println("长度：", len(s1)) //1
	fmt.Println("容量：", cap(s1)) //3

	s1 = s[:]
	fmt.Println(s1)             //[0 1 2 3 4 5 6 7]
	fmt.Println("长度：", len(s1)) //8
	fmt.Println("容量：", cap(s1)) //8
}

/*
切片拷贝
copy(dest, source)
copy(dest, source[1:3])
使用拷贝操作后s1 s2是两个独立空间 不会相互影响
*/
func slicetest4() {
	s1 := []int{1, 2, 3, 4}
	s2 := make([]int, 3)

	copy(s2, s1)

	fmt.Println(s1)        //[1 2 3 4]
	fmt.Printf("%p\n", s1) //0xc000010360
	fmt.Println(s2)        //[1 2 3]
	fmt.Printf("%p\n", s2) //0xc000010380

	s2 = make([]int, 5)
	copy(s2, s1)

	fmt.Println(s1)        //[1 2 3 4]
	fmt.Printf("%p\n", s1) //0xc000010360
	fmt.Println(s2)        //[1 2 3 4 0]
	fmt.Printf("%p\n", s2) //0xc00000c2d0
}

/*
切片作为函数参数
切片作为函数参数是地址传递 形参可以改变实参的值
*/
func slicetest5() {
	s1 := []int{1, 2, 3, 4}

	slicetest6(s1)

	fmt.Println("after:", s1) //[9 10 3 4]
}
func slicetest6(s []int) {
	s[0] = 9
	s[1] = 10
	fmt.Println("形参修改：", s) //[9 10 3 4]
}

type Person struct {
	id    uint64
	name  string
	sex   *int
	hobby *string
}

func slicetest7() {
	p1 := Person{}
	p1.id = 1001
	p1.name = "张三"
	sex1 := 1
	p1.sex = &sex1
	hobby1 := "swim"
	p1.hobby = &hobby1

	p2 := Person{}
	p2.id = 1002
	p2.name = "lisi"
	sex2 := 2
	p2.sex = &sex2
	hobby2 := "jump"
	p2.hobby = &hobby2

	a := []Person{p1, p2}
	fmt.Println(a)
	for _, p := range a {
		fmt.Println(p)
	}

	fmt.Println("-----------------------")

	for i, p := range a {
		p.name = p.name + "-" + strconv.Itoa(i)
		newHobby := *p.hobby + "-" + strconv.Itoa(i)
		p.hobby = &newHobby
	}
	fmt.Println(a)
	for _, p := range a {
		fmt.Println(p)
	}

}

func slicetest8() {
	p1 := Person{}
	p1.id = 1001
	p1.name = "张三"
	sex1 := 1
	p1.sex = &sex1
	hobby1 := "swim"
	p1.hobby = &hobby1

	p2 := Person{}
	p2.id = 1002
	p2.name = "lisi"
	sex2 := 2
	p2.sex = &sex2
	hobby2 := "jump"
	p2.hobby = &hobby2

	a := []*Person{&p1, &p2}
	fmt.Println(a)
	for _, p := range a {
		fmt.Println(*p)
	}

	fmt.Println("-----------------------")

	for i, p := range a {
		p.name = p.name + "-" + strconv.Itoa(i)
		newHobby := *p.hobby + "-" + strconv.Itoa(i)
		p.hobby = &newHobby
	}
	fmt.Println(a)
	for _, p := range a {
		fmt.Println(*p)
	}

}
func main() {
	//slicetest1()
	//slicetest2()
	//slicetest3()
	//slicetest4()
	//slicetest5()
	slicetest7()
	//slicetest8()
}
