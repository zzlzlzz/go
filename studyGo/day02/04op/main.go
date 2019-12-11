package main

import "fmt"

func main() {
	var (
		a = 5
		b = 10
	)
	//算术运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ //a=a+1
	b-- //b=b-1
	fmt.Println(a, b)

	//关系运算符
	fmt.Println(a == b) //go语言是强类型，相同类型变量才可以进行比较
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a > b)
	fmt.Println(a <= b)
	fmt.Println(a < b)

	//逻辑运算符
	fmt.Println(a > 1 && b < 4)
	fmt.Println(a > 1 || b < 4)

	//取反，原来为真就为假，原来为假就为真
	isMarried := false
	fmt.Println(!isMarried)
	fmt.Println(isMarried)

	//位运算：针对二进制数
	//5的二进制表示：101
	//2的二进制表示：10

	//&:按位与(对应位置均为一为1)
	fmt.Println(5 & 2) //000
	//|:按位或(两位有1个为1就为1)
	fmt.Println(5 | 2) //111
	//^:按位异或(梁伟文为不一样为1)
	fmt.Println(5 ^ 2) //111
	//<<:将二进制位左移指定位数
	fmt.Println(5 << 1)  //1010=>10
	fmt.Println(1 << 10) //10000000000=>1024
	//>>:将二进制位右移指定位数
	fmt.Println(5 >> 1)  //10=>2
	var m = int8(1)      //只能存8位
	fmt.Println(m << 10) //10000000000，只存8位，即为0

	//192.168.1.1
	//权限 文件操作

	var x int
	x += 1
	fmt.Println(x)
}
