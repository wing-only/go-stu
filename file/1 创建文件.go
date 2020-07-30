package main

import (
	"fmt"
	"io"
	"os"
)

/*
创建文件
 */
func filecreatetest1()  {
	file, err := os.Create("E:/a.txt")
	if err != nil{
		fmt.Println("创建文件失败")
		return
	}
	defer file.Close()

	name := file.Name()
	fmt.Println(name)
}

/*
写入文件
 */
func filecreatetest2()  {
	//如果文件存在，会覆盖原来的内容
	file, err := os.Create("E:/a.txt")
	if err != nil{
		fmt.Println("创建文件失败")
		return
	}
	defer file.Close()

	//写入文件1
	n, err := file.WriteString("fuck you bitch \n")
	fmt.Println(n)
	n1, err := file.WriteString("let's go \n")
	fmt.Println(n1)

	//写入文件2
	b := []byte{'h', 'e', 'l', 'l', 'o'}
	n, err = file.Write(b)

	str := " world \n"
	bytes := []byte(str)
	file.Write(bytes)

	//写人文件3
	//会覆盖原始文件内容
	file.WriteAt(b, 1)		//off 偏移量

}

/*
打开文件
 */
func filecreatetest3()  {
	//以只读方式打开
	//file, err := os.Open("E:/a.txt")

	//打开文件（文件名，打开模式，打开权限）
	file, err := os.OpenFile("E:/a.txt", os.O_RDWR, 6)
	if err != nil{
		fmt.Println("文件打开失败")
		return
	}
	defer file.Close()

	b := []byte("你愁啥")
	file.WriteAt(b, 1)

	//获取文件
	seek, err := file.Seek(0, io.SeekEnd)
	fmt.Println(seek)

	b1 := []byte("抽你咋地")
	file.WriteAt(b1, seek)
}
/*

 */
func filecreatetest4()  {

}
/*

 */
func filecreatetest5()  {

}
/*

 */
func filecreatetest6()  {

}
/*

 */
func filecreatetest7()  {

}
/*

 */
func filecreatetest8()  {

}
func main() {
	//filecreatetest1()
	//filecreatetest2()
	filecreatetest3()

}
