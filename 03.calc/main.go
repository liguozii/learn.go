package main

import "fmt"

func main() {
	var a, b = 3, 10
	fmt.Println("a + b = ", a+b)
	fmt.Println("a - b = ", a-b)
	fmt.Println("a * b = ", a*b) // 注意类型导致溢出
	fmt.Println("a / b = ", a/b) // 注意不能除以0 否者会panic退出程序，后续会有异常检测
	fmt.Println("a % b = ", a%b)
}
