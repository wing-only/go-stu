package main

import "fmt"

/*
%c 打印字符类型数据
%d 打印整型类型数据
%f 打印浮点型数据 fmt.printf("%f")  默认保留六位小数%.2f 保留两位小数 对第三位四舍五入
%s 打印字符串类型
%p 打印变量地址（指针类型）
%t 打印布尔类型数据
%T 打印数据类型  a:=3.14 float64 a:=10 int a:='a'int32

%o 打印八进制整型数据
%x %X 打印十六进制整型数据  a-f A-F
%%  打印一个%

转义字符
'\n' 表示换行
'\\' 表示\
'\0' 表示字符串结束标志 ASCII值为0

 */
//fmt.Printf() 格式化输出
func test1() {
	var a bool = true
	fmt.Println(a)

	//%t 占位符 表示输出一个布尔类型数据
	fmt.Printf("%t\n", a)

	//%s 占位符 表示输出一个字符串类型
	var b string = "我喜欢学习golang"
	fmt.Printf("我是谁，%s\n", b)

	//%c 占位符， 表示输出一个字符类型
	var c = 'c'
	fmt.Println("c的ASCII码是:", c)
	fmt.Printf("输出字符：%c\n", c)

}

//输入 fmt.Scan()
func test2()  {
	var a int
	fmt.Scan(&a)
	fmt.Println(a)

	//%p 占位符 表示输出一个数据对应的内存地址 &a
	fmt.Printf("%p\n", &a)

	var str string
	fmt.Scan(&str)
	fmt.Println(str)

	//接收两个数据，空格或者回车作为接收结束
	var s1,s2 string
	fmt.Scan(&s1, &s2)
	fmt.Println(s1)
	fmt.Println(s2)

}

//格式输入 fmt.Scanf()
func test3()  {
	var name string
	var passwd int

	fmt.Println("请输入用户名")
	fmt.Scanf("%s", &name)
	fmt.Println("请输入密码（数字）")
	fmt.Scanf("%d", &passwd)
	fmt.Printf("用户名是%s，密码是%d", name, passwd)
}

func test4()  {
	sprintf := fmt.Sprintf("hello, %s", "wing")
	fmt.Println(sprintf)
}

func main() {
	//test1()
	//test2()
	//test3()
	test4()
}
