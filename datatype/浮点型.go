package main

import (
	"fmt"
	"math/big"
)

/*
float32
float64
*/
func floattest1() {
	var a float32
	var b float64

	//float32默认小数位保留7位有效数据，会在末尾 +1 操作
	a = 3.141593400000000

	//float64默认小数位保留15位有效位
	b = 3.1415936273242342324367464375

	fmt.Println(a)
	fmt.Println(b)
}

func floattest2() {
	a := 0.1
	b := 0.2
	c := a * b
	fmt.Println(c)
}
func floattest3() {
	a := big.NewFloat(0.1)
	b := big.NewFloat(0.2)
	c := a.Mul(a, b)
	fmt.Println(c.String())

}

func main() {
	//floattest1()
	//floattest2()
	floattest3()
}
