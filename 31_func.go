package main

import "fmt"

func Test(args ...int) {
	for _, n := range args { //遍历参数列表
		fmt.Printf("%d ", n)
	}
}

func main() {
	//函数调用，可传0到多个参数
	Test()
	Test(1)
	fmt.Println("\n=======================")

	Test(4, 5, 6, 7)
}
