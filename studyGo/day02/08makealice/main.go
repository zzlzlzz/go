package main

import "fmt"

//make函数创造切片

func main() {

	// s1 := make([]int, 5, 10)
	// fmt.Printf("len(s1)=%v cap(s1)=%v\n", len(s1), cap(s1))

	// s2 := make([]int, 5)
	// fmt.Printf("len(s2)=%v cap(s2)=%v\n", len(s2), cap(s2))

	// s3 := make([]int, 0, 10)
	// fmt.Printf("len(s3)=%v cap(s3)=%v\n", len(s3), cap(s3))

	//切片的赋值
	s4 := []int{1, 3, 5}
	s5 := s4 //s3和s4都指向同一个底层数组
	fmt.Println(s4, s5)
	s4[0] = 2
	fmt.Println(s4, s5)

	//切片的遍历
	//1.索引遍历
	for i := 0; i < len(s4); i++ {
		fmt.Println(i, s4[i])

	}
	//2.for range 循环
	for i, v := range s4 {
		fmt.Println(i, v)

	}

}
