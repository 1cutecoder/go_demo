package main

import "fmt"

func myfunc(tmp ...int) {
	for _, data := range tmp { //遍历参数列表
		fmt.Printf("myfunc data = %d ", data)
	}
}
func myfunc2(tmp ...int) {
	for _, data := range tmp { //遍历参数列表
		fmt.Printf("myfunc2 data = %d ", data)
	}
}
func Test1(args ...int) {
	for _, n := range args { //遍历参数列表
		fmt.Printf("%d ", n)
	}
	fmt.Println("\n")

	myfunc2(args[:2]...)
	fmt.Println("\n=======================")
	myfunc2(args[2:]...)
}

func main() {
	//函数调用，可传0到多个参数
	//Test1()
	//Test1(1)
	fmt.Println("\n=======================")

	Test1(4, 5, 6, 7)
}
