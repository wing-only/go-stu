package main

import "fmt"

/*
 var arr [10]int
*/
func arrytest1() {
	var arr [10]int
	arr[0] = 1
	arr[1] = 2

	fmt.Println(arr)

	//遍历数组
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	//遍历数组
	for i, v := range arr {
		fmt.Printf("arr[%d]=%d\n", i, v)
	}
}

/*
 var arr [10]int= [10]int{1, 2, 3, 4}
*/
func arrytest2() {
	var arr = [10]int{1, 2, 3, 4}
	fmt.Println(arr)
}

/*
 自动推导类型，可以根据元素个数创建数组
 arr := [...]int{1, 2, 3}
*/
func arrytest3() {
	arr := [...]int{1, 2, 3}

	for _, v := range arr {
		fmt.Println(v)
	}
}

/*
数组作为函数参数传递
数组作为函数参数是值传递的
*/
func arrytest4() {
	arr := [...]string{"wing", "张飞", "刘备"}

	arrytest5(arr)
	fmt.Println(arr)				//[wing 张飞 刘备]
}
func arrytest5(arr [3]string) {
	arr[0] = "曹操"
	fmt.Println("形参：", arr)	//[曹操 张飞 刘备]
}
func main() {
	//arrytest1()
	//arrytest2()
	//arrytest3()
	arrytest4()
}
