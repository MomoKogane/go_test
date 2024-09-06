package main

import (
	"fmt"
)

func main() {
	//定义三个变量，stockcode为整型数字，enddate为字符串，url为格式化字符串
	// %d 表示整型数字，%s 表示字符串
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d\nendDate=%s"

	//定义一个字符串变量url，其中%d表示整型数字，%s表示字符串
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	//函数Sprintf()用于格式化字符串，第一个参数为格式化字符串，第二个及以后的参数为要替换的变量
	fmt.Println(target_url)

	//var target_url2 = fmt.Printf(url, stockcode, enddate)
	//函数Printf()用于格式化输出，第一个参数为格式化字符串，第二个及以后的参数为要输出的变量
	//Printf函数返回两个包含两个值的元组，一个是写入的字节数，另一个是遇到的任何错误。
	var target2 int
	target2, _ = fmt.Printf(url, stockcode, enddate)
	fmt.Println()
	fmt.Println(target2)
}

// Output: Code=123&endDate=2020-12-31
