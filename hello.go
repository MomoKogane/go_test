package main //声明程序属于main包，
//每个程序都需要一个main包且该报中要使用一个main函数作为程序的入口点

import "fmt" //引入fmt包，该包提供了格式化输出的功能，如Println()函数，用于输出字符串到标准输出

// main函数是程序的入口点
func main() {
	//在Go语言中，{符号不能单独出现，
	// 必须与关键字一起使用，表示一个代码块的开始
	// 打印输出"Hello, World!"
	fmt.Println("Hello, World!")         //调用Println()函数输出字符串到标准输出
	fmt.Println("除非多行代码卸载同一行\n不然不需要分号;") //调用Println()函数输出字符串到标准输出
}
