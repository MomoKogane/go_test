package main

import "fmt"

// 定义一个全局变量
var globalVar string = "10" //Go 语言中变量的声明必须使用空格隔开，如 var a int = 10

func main() {
	// 定义一个局部变量
	localVar := "5"

	// 打印全局变量和局部变量的值
	fmt.Print("Global variable:", globalVar)
	fmt.Println(" Local variable:", localVar)
	fmt.Println("Now variable:", localVar+globalVar)
}
