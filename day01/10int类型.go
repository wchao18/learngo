package main

import "fmt"

func main01() {
	//int有符号、uint无符号
	//分类:int8、int16、int32、int64;uint8、uint16、uint32、uint64
	//int 会根据操作系统的不同在内存中占用的字节也是不同的
	var a int
	var b uint
	fmt.Println(a, b)
}

//有符号溢出测试
func main() {
	var a int = 10
	fmt.Printf("%T\n", a)
	//假如出现了溢出 正号会变成负号
	b := a - 100000000000
	fmt.Println(b)
}

//无符号溢出测试
func main03() {
	var a uint = 10
	fmt.Printf("%T\n", a)
	//出现了溢出,从最小值变成最大值
	b := a - 20
	fmt.Println(b)
}
