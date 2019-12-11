package main

//声明变量
var (
	s1 = "1"
	s2 = 99
	s3 = true
)

const (
	pi       = 3.14
	notFound = 404
	n1
	n2
	n3 = iota
)

const (
	a1 = iota //0
	a2 = iota //1
	_  = iota //2
	a3 = iota //3
)

const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota
	c4
)

const (
	d1, d2 = iota + 1, iota + 2
	d3, d4 = iota + 1, iota + 2
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota) //1向左移动10*1位，即二进制10000000000，等于十进制的1024
	MB = 1 << (10 * iota) //1向左移动10*2位，即二进制2^20，等于十进制的1024*1024
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	// fmt.Println("人生苦短，Let's go!")
	// fmt.Println(s3)
	// fmt.Print(s2)
	// fmt.Printf(s1)
	// fmt.Println()
	// s4 := "Python is a good language."
	// fmt.Println(s4)
	// fmt.Println(n1)
	// fmt.Println(n2)
	// fmt.Println(n3)
	// fmt.Println(a1)
	// fmt.Println(a2)
	// fmt.Println(a3)
	// fmt.Println("d1:", d1)
	// fmt.Println("d2:", d2)
	// fmt.Println("d3:", d3)
	// fmt.Println("d4:", d4)
}
