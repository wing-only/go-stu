package main

import "fmt"

func comparetest1() {
	a := 10
	b := 20

	fmt.Println(a == b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a != b)
}

func main() {
	comparetest1()
}
