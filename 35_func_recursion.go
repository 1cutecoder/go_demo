package main

import "fmt"

func Test001() int {
	i := 1
	sum := 0
	for i = 1; i <= 100; i++ {
		sum += i
	}

	return sum
}

// 通过递归实现1+2+3……+100
func Test002(num int) int {
	if num == 1 {
		return 1
	}

	return num + Test002(num-1) //函数调用本身
}

// 通过递归实现1+2+3……+100
func Test003(num int) int {
	if num == 100 {
		return 100
	}

	return num + Test003(num+1) //函数调用本身
}

func main() {

	fmt.Println(Test001())    //5050
	fmt.Println(Test002(100)) //5050
	fmt.Println(Test003(1))   //5050
}
