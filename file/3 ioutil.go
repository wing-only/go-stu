package main

import "io/ioutil"

func ioutiltest1() {
	ioutil.WriteFile("G:/aaa.txt", []byte("aabbccdd"), 666)

}

func main() {
	ioutiltest1()
}
