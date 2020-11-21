
//在go语言中主函数main所属的包一定的main
package main

//导入系统的包,标准的输入输出包
import "fmt"

//func 函数格式,程序只有一个主函数(main)类似于java
func main() {
	//包.方法()
	fmt.Println("hello world!")
	test1()
}


func test1()  {
	//行注释 ctrl+/
	//fmt.Println("注释")
	//块注释  shift+ctrl+/
	/*fmt.Println()
	fmt.Println()
	fmt.Println()*/
}
