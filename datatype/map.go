package main

import "fmt"

/*
map定义
key是唯一的  不允许重复
在map中值允许重复
*/
func maptest1() {
	var m map[int]string = map[int]string{1: "wing"}
	m[2] = "lily"
	fmt.Println(m) //map[1:wing 2:lily]

	m2 := make(map[string]string, 2)
	m2["name"] = "wing"
	m2["sex"] = "男"
	m2["hobby"] = "写代码"
	fmt.Println(m2) //map[hobby:写代码 name:wing sex:男]

	//遍历map
	for k, v := range m2 {
		fmt.Println(k, v)
	}
}

/*
删除map元素
 */
func maptest2() {
	m:=map[int]string{101:"刘备", 103:"关羽", 105:"张飞"}
	fmt.Println(m)			//map[101:刘备 103:关羽 105:张飞]

	delete(m, 103)
	fmt.Println(m)			//map[101:刘备 105:张飞]

	delete(m, 111)		//key不存在，不会报错
	fmt.Println(m)			//map[101:刘备 105:张飞]
}

/*
map作为函数参数
map作为函数参数是地址传递 （引用传递）
 */
func maptest3() {
	m:=make(map[string]string)
	m["name"] = "wing"
	m["hobby"] = "代码"

	maptest4(m)
	fmt.Println("调用后", m)			//map[code:golang hobby:代码 name:张飞]
}
func maptest4(m map[string]string) {
	m["name"] = "张飞"
	m["code"] = "golang"

	fmt.Println("形参中：", m)		//map[code:golang hobby:代码 name:张飞]
}
func main() {
	//maptest1()
	//maptest2()
	maptest3()
}
