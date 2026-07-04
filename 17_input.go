package main

import "fmt"

func main() {
	var v int
	fmt.Println("请输入一个整型：")
	fmt.Scanf("%d", &v)
	//fmt.Scan(&v)
	fmt.Println("v = ", v)
}
