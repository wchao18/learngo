package main

import "fmt"

func main() {
	//浮点类型分类:float32(小数位精确到7位)、float(小数位精确到15位)，ps:可以超出15位,但是数据不精确了
	var a float32 = 3.14
	var b float64 = 3.14
	fmt.Println(a, b)
	//保留20位有效数字
	fmt.Printf("%.20f", a)
	fmt.Println()
	//%f默认保留6位小数,四舍五入,数据会更精确
	fmt.Printf("%f,%f", a, b)
	fmt.Println()
	//自动类型推导默认的类型是float64
	c := 3.14
	fmt.Printf("%T", c)
}
