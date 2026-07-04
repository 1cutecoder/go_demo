package main

import "fmt"

func Test01() (int, string) { //方式1
	return 250, "sb"
}

func Test020() (c int, str1 string) {
	c = 260
	str1 = "sb2"
	return
}

func Test02() (a int, str string) { //方式2, 给返回值命名
	a = 250
	str = "sb"
	return
}

func main() {
	v1, v2 := Test01() //函数调用
	_, v3 := Test02()  //函数调用， 第一个返回值丢弃
	v4, _ := Test02()  //函数调用， 第二个返回值丢弃
	_, v5 := Test020()
	V6, _ := Test020()
	fmt.Printf("v1 = %d, v2 = %s, v3 = %s, v4 = %d\n", v1, v2, v3, v4)
	fmt.Printf("v5 = %s, v6 = %d\n", v5, V6)
}
