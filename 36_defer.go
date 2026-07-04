package main

import "fmt"

func main() {
	defer fmt.Println("1.this is a defer") //main结束前调用
	defer fmt.Println("2.this is a defer") //main结束前调用
	fmt.Println("this is a test")

	/*
		运行结果：
		this is a test
		this is a defer
	*/
}
