package main

import (
	"fmt"
	"math"
)

func main1() {
	//定义好的变量一定要去使用
	var a int //error
}

func main() {
	var a = 20
	var a = 20     //在一个作用域内,变量名不能重复定义,在重复赋值中有一个特殊案例
	fmt.Println(a) //error
}

func main1() {
	//变量的定义与使用
	//var 变量名 数据类型 = 数值
	var sun int = 123
	sun = sun + 10
	sun += 20
	fmt.Println(sun)
}

func main2() {
	//变量先声明后使用,没有赋值默认0
	var sun int
	//sun = 123
	fmt.Println(sun)
}

func main3() {
	var value float64 = 2
	//使用系统math包里面的函数
	var sum float64 = math.Pow(value, 10)
	fmt.Println(sum)
}
