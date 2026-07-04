package main

import "fmt"

func main() {
	switch num := 1; num {
	case 1:
		fmt.Println("按下的是1")
	case 2:
		fmt.Println("按下的是2")
	default:
		fmt.Println("按下的是xxx")
	}
}
