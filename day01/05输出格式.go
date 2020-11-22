package main

import "fmt"

func main() {
	//format
	a := 10
	b := 3.1455
	c := "哈哈"
	//带格式的输出
	//%d 占位符输出一个整型数据
	//%f 占位符,浮点数不是精确的数据,默认保留6位小数
	//%.2f 保留两位小数,第3位小数会四舍五入
	//%s 字符串
	//\n表示转义字符,换行
	fmt.Printf("%d %f %s", a, b, c)
	fmt.Println()
	fmt.Printf("%d,%.2f,%s", a, b, c)
}
