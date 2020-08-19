package main

import (
	"fmt"
	"time"
)

func timetest1() {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(now)
}
func timetest2() {
	now := time.Now()

	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Println("时：", hour)
	fmt.Println("分：", minute)
	fmt.Println("秒：", second)
}

func main() {
	//timetest1()
	timetest2()
}
