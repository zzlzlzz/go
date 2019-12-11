package main

import "fmt"

//流程控制——跳出for循环
func main() {
	// //当i=5时跳出for循环
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break //当i=5时跳出for循环
	// 	}
	// 	fmt.Println(i)
	// }
	// fmt.Println("over")

	//当i=5时，跳过此次for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("over")

}
