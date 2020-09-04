package main

import "fmt"

/*
	var 指针 *数据类型
 */

func pointtest1() {
	var p *int
	fmt.Println(p)			//<nil>
	//fmt.Println(*p)		//报错

	p = new(int)

	fmt.Println(p)			//0xc0000a0098
	fmt.Println(*p)			//0
}

func pointtest2() {
	var a int = 123
	var p *int
	p = &a
	fmt.Println(a)		//123
	fmt.Println(p)		//0xc00000e0a8
	fmt.Println(*p)		//123

	*p = 234
	fmt.Println(a)		//234
	fmt.Println(p)		//0xc00000e0a8
	fmt.Println(*p)		//234
}

func pointtest3() {
	var p *int
	var a int = 23
	var b int = 24
	p = &a

	fmt.Println(a)		//23
	fmt.Println(b)		//24
	fmt.Println(p)		//0xc00000e0a8
	fmt.Println(*p)		//23

	p = &b

	fmt.Println(a)		//23
	fmt.Println(b)		//24
	fmt.Println(p)		//0xc00000e0c0
	fmt.Println(*p)		//24
}

/*
指针作为函数参数 1
指针作为函数参数是地址传递
*/
func pointtest4() {
	var p *int
	var a int = 23
	p = &a

	fmt.Println("函数调用前：", a, p)		//23 0xc00000e0a8

	pointtest5(p)

	fmt.Println("函数调用后：", a, p)		//24 0xc00000e0a8
}
func pointtest5(p *int) {
	var b int = 24
	*p = b
	fmt.Println("函数调用：", p)			//0xc00000e0a8
}

/*
指针作为函数参数 2
 */
func pointtest6()  {
	var p *int
	var a int = 23
	p = &a

	fmt.Println("函数调用前：", a, p)		//23 0xc00000e0a8

	pointtest7(p)

	fmt.Println("函数调用后：", a, p)		//23 0xc00000e0a8  函数里面改变地址后就释放了，这里还是原来的地址
}
func pointtest7(p *int)  {
	var b int = 24
	p = &b
	fmt.Println("函数调用：", p)			//0xc00000e0c8   改变了地址
}

func pointtest8()  {
	var a = 10
	var b = &a
	var c = &b
	var d = &c

	fmt.Println(a)
	fmt.Printf("%T", a)
	fmt.Println(b)
	fmt.Printf("%T", b)
	fmt.Println(c)
	fmt.Printf("%T", c)
	fmt.Println(d)
	fmt.Printf("%T", d)
}
func main() {
	//pointtest1()
	//pointtest2()
	//pointtest3()
	//pointtest4()
	//pointtest6()
	pointtest8()
}
