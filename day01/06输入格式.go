package main

import "fmt"

func main01() {

	var a int
	//通过键盘为变量赋值,&是一个运算符,取地址运算符
	fmt.Println(&a)
	fmt.Scan(&a)
	fmt.Println(a)
}

func main02() {
	var a int
	var b string
	//空格或者回车表示一个输入的结束
	//结束输入用回车
	fmt.Scan(&a, &b)
	fmt.Println(a, b)
}

func main03() {
	var a int
	var b string
	//ps:scanf这里不能用回车,结束输入用回车
	fmt.Scanf("%d%s", &a, &b)
	fmt.Println(a, b)
}

//case 输入三个数字求和和平均值
func main04() {
	var a, b, c int
	fmt.Scanf("%d%d%d", &a, &b, &c)
	fmt.Printf("a=%d,b=%d,c=%d", a, b, c)
	fmt.Println()
	sum := a + b + c
	avg := sum / 3
	fmt.Println("总和:", sum)
	fmt.Println("平均:", avg)
}
