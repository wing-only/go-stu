package main

import "fmt"

func switchtest1() {
	var w int
	fmt.Scan(&w)

	switch w {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	default:
		fmt.Println("星期天")
	}
}

func main() {
	switchtest1()
}
