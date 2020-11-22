package main

import "fmt"

//方式一
func main1() {
	//多重赋值变量个数和值的个数要一一对应
	var a, b, c = 10, 2.3, "哈哈哈" //不同变量类型
	var d, e, f = 30, 40, 50     //相同变量类型
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
}

//特殊用法
func main() {
	a, b := 10, "哈哈"
	//a, b := 10, "哈哈"//error
	//在多重赋值的时候有新的变量可以用自动推导
	//a, b, c := 10, "哈哈", 2.3 //correct
	//匿名变量
	_, a, b, c := 10, 20, "哈哈1", 2.3
	fmt.Println(a, b, c)
}

//方式二
func main2() {
	a, b, c := 10, 2.3, "哈哈哈"
	fmt.Println(a, b, c)
}

//数据交换方式1
func main3() {
	a, b := 20, 30
	tmp := a
	a = b
	b = tmp
	fmt.Println(a, b)
}

//交换数据方式2
func main4() {
	a, b := 20, 30
	a, b = b, a
	fmt.Println(a, b)
}
