package main

import "fmt"

func main01() {
	var sum int = 20
	var sum1 = 20 //省略类型
	fmt.Println(sum)
	fmt.Println(sum1)
}

func main02() {
	//自动类型推导
	a := 10
	b := 23.656666 //float64
	c := "哈哈"      //string
	fmt.Println(a, b, c)
}
