package main

import "fmt"

/*
bool
 */
func booltest1() {
	var a bool
	fmt.Println(a)

	a = true
	fmt.Println(a)

	b := true
	fmt.Println(b)
}

func main() {
	booltest1()
}
