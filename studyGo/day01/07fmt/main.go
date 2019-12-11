package main

import "fmt"

func main() {
	n := 100
	s := "hello"
	a := 1
	fmt.Printf("%T\n", n)
	fmt.Printf("%v\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("字符串：%s\n", s)
	fmt.Printf("字符串：%v\n", s)
	fmt.Printf("字符串：%#v\n", s)
	_ = a
}
