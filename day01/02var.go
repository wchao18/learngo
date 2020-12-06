package main

import (
	"fmt"
	"math"
)

func main01() {
	//定义好的变量一定要去使用
	var a int //error
	a = 20
}

func main02() {
	var a = 20
	var a = 20     //在一个作用域内,变量名不能重复定义,在多重赋值中有一个特殊案例
	fmt.Println(a) //error
}

func main03() {
	//变量的定义与使用
	//var 变量名 数据类型 = 数值
	var sum int = 10
	sum = sum + 20
	sum += 30
	fmt.Println(sum)
}

func main04() {
	//变量先声明后使用,没有赋值默认0
	var sum int
	//sum = 123
	fmt.Println(sum)
}

func main05() {
	var value float64 = 2
	//使用系统math包里面的函数
	var sum = math.Pow(value, 10)
	fmt.Println(sum)
}
