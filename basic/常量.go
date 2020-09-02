package main

import "fmt"


/*
常量
格式 const 常量名 数据类型=值  一般常量大写字母
常量的值在程序运行过程中其值不能发生改变
常在在内存中的数据区进行存储   变量是在内存中的栈区存储
系统为每一个应用程序分配了1M的内存空间存储变量和函数的信息
*/
func consttest1() {
	const NAME string = "wingx"
	fmt.Println(NAME)
}

/*
iota 枚举
 */
func iotatest1() {
	const a = iota			//0
	const b, c = iota, iota	//0,0
	const d = iota			//0
	const e = 20			//20
	const f = iota			//0
	const g = iota			//0

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}

func iotatest2() {
	const (
		a    = iota       //0
		b, c = iota, iota //1,1
		d    = iota       //2
		e    = 20         //20
		f    = iota       //4
		g    = iota       //5
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}

func main() {
	//consttest1()
	iotatest1()
	//iotatest2()
}
