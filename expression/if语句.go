package main

import "fmt"

func iftest1() {
	a := 10
	b := 20

	if a > b {
		fmt.Println("a大于b")
	} else if a == b {
		fmt.Println("a等于b")
	} else {
		fmt.Println("a小于b")
	}
}

func main() {
	iftest1()
}
