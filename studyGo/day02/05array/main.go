package main

import "fmt"

/*
	数组
	存放元素的容器
	必须指定存放元素的类型和容量（长度）
	数组的长度是数组类型的一部分
*/

func main() {
	// var a1 [3]bool //[true false false]
	// var a2 [4]bool //[true true false false]
	// fmt.Printf("a1:%T\na2:%T\n", a1, a2)

	// //数组的初始化
	// var b1 [3]int
	// var b2 = [4]int{1, 2}
	// fmt.Println(b1)
	// fmt.Println(b2)
	a := [...]int{1, 3, 5, 7, 9}
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == 8 {
				fmt.Println(i, j)

			}
		}
	}

}

