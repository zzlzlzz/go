package main

import "fmt"

//append()为切片追加元素
func main() {
	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("s1:%v len(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1))

	//s1[3] = "广州" //错误写法，导致编译错误，索引越界
	// fmt.Println(s1)

	//调用append（）函数必须使用原来的切片变量接受返回值
	s1 = append(s1, "广州") // append加元素,原来的底层数组放不下的时候,Go底层就会把底层数组换一个
	fmt.Printf("s1:%v len(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "杭州", "成都")
	fmt.Printf("s1:%v len(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1))
	ss := []string{"武汉", "西安", "苏州"}
	s1 = append(s1, ss...) //...表示拆开
	fmt.Printf("s1:%v len(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1))

}
