package main

import "fmt"

//copy
func main() {
	// a1 := []int{1, 3, 5}
	// a2 := a1
	// var a3 = make([]int, 3, 3)
	// copy(a3, a1)
	// fmt.Println(a1, a2, a3)
	// a1[0] = 100
	// fmt.Println(a1, a2, a3)

	// //将a1中索引为3的元素删除
	// a1 = append(a1[:1], a1[2:]...)
	// fmt.Println(a1)

	x1 := [...]int{1, 3, 5, 7, 9} //数组
	a1 := x1[:]                   //切片
	fmt.Printf("%v %p\n", a1, a1) //打印a1的值及对应的内存地址
	// 1.切片不保存具体的值
	// 2.切片对应一个底层数组
	// 3.底层数组都是占用一块连续的内存
	a1 = append(a1[:1], a1[2:]...) //进行删除a1[1]操作，修改了底层数组的值！
	fmt.Printf("%v %p\n", a1, a1)  //打印删除后a1的值及对应的内存地址
	fmt.Println(x1)                //x1变了
}
