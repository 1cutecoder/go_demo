package main

import "fmt"

func main() {
	var a int
	fmt.Println("a = ", a)
	a = 10
	fmt.Println("a = ", a)

	var b int = 10
	fmt.Println("b = ", b)
	//自动推导
	c := 100
	fmt.Printf("c type is %T\n ", c)
	fmt.Printf("c = %d\n ", c)
	fmt.Printf("a = %d,c = %d\n ", a, c)
	//交换两个数
	i := 10
	j := 20
	i, j = j, i
	fmt.Printf("i = %d,j = %d\n ", i, j)

}
