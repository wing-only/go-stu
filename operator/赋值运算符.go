package main

import "fmt"

func fuzhitest1() {
	a := 10
	a = a + 5
	fmt.Println(a)

	a += 5
	fmt.Println(a)

	a -= 5
	fmt.Println(a)

	a *= 5
	fmt.Println(a)

	a /= 5
	fmt.Println(a)

	a %= 5
	fmt.Println(a)

	b := a + 5*a
	fmt.Println(b)
}

func main() {
	fuzhitest1()
}
