package main

import "fmt"

func main() {
	// i := 11
	// if i < 100 {
	// 	fmt.Println(i)
	// 	i = i + 11
	// }
	// fmt.Println(i)

	// for a := 0; a < i; a++ {
	// 	fmt.Println(a)
	// }

	// var a = 0

	// for ; a+1 < i; a++ {
	// 	fmt.Println(a)
	// }

	// fmt.Println(true)

	// for a < 20 {
	// 	fmt.Println(a)
	// 	a++
	// }

	//九九乘法表
	for n := 1; n < 10; n++ {
		for a := 1; a < n+1; a++ {
			fmt.Printf("%d*%d=%d\t", n, a, a*n)
		}
		fmt.Println()
	}

	// s := "hello"
	// fmt.Println(len(s))
	// for _, v := range s {
	// 	fmt.Printf("%c\n", v)
	// }

}
