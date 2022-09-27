package main

import (
	"fmt"
	"reflect"
)

func main() {
	var money interface{} = 10.0

	switch newmoney := money.(type) {
	case int:
		fmt.Println("money是 int", newmoney+1.0)
	case int64:
		fmt.Println("money是 int64")
	case float64:
		fmt.Println("money是 float64")
	case float32:
		fmt.Println("money是 float32")
	default:
		fmt.Println("money是未知类型")
	}
	fmt.Println(reflect.TypeOf(money))
	money = 3
	fmt.Println(reflect.TypeOf(money))
	fmt.Println("结束", money)
}
