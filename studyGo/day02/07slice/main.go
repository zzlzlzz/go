package main

import "fmt"

func main() {
	//切片的定义
	var a []int    //定义一个类型为整型的切片
	var b []string //定义一个类型为字符串的切片
	a = []int{1, 2, 3}
	b = []string{"北京", "上海", "新加坡"}
	fmt.Println(a)
	fmt.Println(b)

	//长度和容量
	fmt.Printf("len(a):%d cap(a):%d\n", len(a), cap(a))
	fmt.Printf("len(b):%d cap(b):%d\n", len(b), cap(b))

	//由数组得到切片
	n := []int{12, 4, 45, 9, 6, 75, 67}
	n1 := n[0:4] //基于数组切割，左包含右不包含
	n2 := n[1:7] //基于数组切割，左包含右不包含
	n3 := n[4:]  //基于数组切割，左包含右不包含
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Printf("len(n1):%d cap(n1):%d\n", len(n1), cap(n1))
	fmt.Printf("len(n2):%d cap(n2):%d\n", len(n2), cap(n2))
	fmt.Printf("len(n3):%d cap(n3):%d\n", len(n3), cap(n3))
	//切片指向底层数组
	//切片的长度就是它元素的个数
	//切片的容量是底层数组从切片的第一个元素到最后一个元素的数量

	//切片的切片
	n4 := n2[2:]
	fmt.Printf("len(n4):%d cap(n4):%d\n", len(n4), cap(n4))
	fmt.Println(n2)
	fmt.Println(n4)
	n[len(n)-1] = 100
	fmt.Println(n2)
	fmt.Println(n4)

}
