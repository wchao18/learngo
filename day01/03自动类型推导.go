package main

import "fmt"

func main1() {
	var sun int = 20
	var sun1 = 20 //省略类型
	fmt.Println(sun)
	fmt.Println(sun1)
}

func main2() {
	//自动类型推导
	a := 20        //int
	b := 23.656666 //float64
	c := "哈哈哈"     //string
	fmt.Println(a, b, c)
}
