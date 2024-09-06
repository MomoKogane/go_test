package main

import "fmt"

// 定义接口
type MyInterface interface {
	DoSomething()
}

// 定义结构体类型
type OtherStruct struct{}

// 实现接口方法
func (s OtherStruct) DoSomething() {
	fmt.Println("Do Something!\n Interface implementation")
}

// 定义一个符合MyFunc类型的函数
func funcVar(x int) int {
	return x * 2
}

func main() {
	// 定义并输出整数类型变量
	var intVar int = 10
	fmt.Println("整数类型变量:", intVar)

	// 定义并输出浮点类型变量
	var floatVar float64 = 3.14
	fmt.Println("浮点类型变量:", floatVar)

	// 定义并输出字符串类型变量
	var stringVar string = "Hello, World!"
	fmt.Println("字符串类型变量:", stringVar)

	// 定义并输出布尔类型变量
	var boolVar bool = true
	fmt.Println("布尔类型变量:", boolVar)

	// 定义并输出复杂类型变量
	var complexVar complex128 = complex(1, 2)
	fmt.Println("复杂类型变量:", complexVar)

	// 定义并输出指针类型变量
	var intVarPointer *int = &intVar
	fmt.Println("指针类型变量:", intVarPointer)

	// 定义并输出数组类型变量
	var arrayVar [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("数组类型变量:", arrayVar)

	// 定义并输出Channel类型变量
	var channelVar chan int = make(chan int) // 定义一个int类型的Channel变量
	fmt.Println("Channel类型变量:", channelVar)  // 输出Channel类型变量

	// 定义并输出函数类型变量
	fmt.Println(
		"函数类型变量:",
		funcVar(2),
	)

	// 定义并输出切片类型变量
	var sliceVar []int = []int{1, 2, 3, 4, 5}
	fmt.Println("切片类型变量:", sliceVar)

	// 定义并输出结构化类型变量
	type MyStruct struct {
		Field1 int
		Field2 string
	}
	var structVar MyStruct = MyStruct{10, "Hello"}
	fmt.Println("结构化类型变量:", structVar)

	//接口类型变量
	var interfaceVar MyInterface = OtherStruct{}
	fmt.Println("接口类型变量:", interfaceVar)
	interfaceVar.DoSomething()

	// 定义并输出Map类型变量
	var mapVar map[string]int = map[string]int{"one": 1, "two": 2}
	fmt.Println("Map类型变量:", mapVar)
}
