package main

import "fmt"

func main() {
	// var n = 1

	// //繁琐的写法
	// if n == 1 {
	// 	fmt.Println("大拇指")
	// } else if n == 2 {
	// 	fmt.Println("食指")
	// } else if n == 3 {
	// 	fmt.Println("中指")
	// } else if n == 4 {
	// 	fmt.Println("无名指")
	// } else if n == 5 {
	// 	fmt.Println("小拇指")
	// } else {
	// 	fmt.Println("无效的输入")
	// }

	//使用switch
	switch a := 4; a {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入")
	}

}
