package main

import "fmt"

func main01() {
	//多重赋值变量个数和值的个数要一一对应
	var a, b, c = 10, 2.3, "哈哈哈" //不同变量类型
	var d, e, f = 30, 40, 50     //相同变量类型
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
}

//ps:特殊用法,可以重复定义变量
func main02() {
	a, b := 10, "哈哈"
	fmt.Println(a, b)
	//如果左边定义了多个变量，则多个变量的组合中，只要有一个不同，如下，就能正常编译：
	a, c, b := 10, 2.3, "test"
	//匿名变量,除去匿名变量,其它变量只要出现在其它声明中，编译错误
	//_, a, b, c := 10, 20, "哈哈1", 2.3 //error
	_, a, b, c, d := 10, 20, "哈哈1", 2.3, 10 //correct
	fmt.Println(a, b, c, d)
}

//数据交换方式1
func main03() {
	a, b := 20, 30
	tmp := a
	a = b
	b = tmp
	fmt.Println(a, b)
}

//交换数据方式2
func main04() {
	a, b := 20, 30
	a, b = b, a
	fmt.Println(a, b)
}
