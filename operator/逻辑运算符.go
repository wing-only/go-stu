package main

import "fmt"

// &&
func logictest1()  {
	var a bool = false
	var b = true
	fmt.Println(a && b)
}

// ||
func logictest2()  {
	var a bool = false
	var b = true
	fmt.Println(a || b)

}
// !
func logictest3()  {
	var a bool = false
	var b = true
	fmt.Println(!a)
	fmt.Println(!b)
}

func main() {
	logictest1()
	logictest2()
	logictest3()
}
