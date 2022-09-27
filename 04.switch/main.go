package main

import "fmt"

func main() {
	var money int = 20
	var busy bool = false
	switch money {
	case 20:
		fmt.Println("点个外卖")
		if busy {
			break
		}
		fmt.Println("再吃个零食")
		fallthrough
	case 200:
		fmt.Println("下个馆子")
	default:
		fmt.Println("容我想想")
	}
	fmt.Println("end")
}
