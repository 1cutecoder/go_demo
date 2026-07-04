package main

import "fmt"

func main() {
	var a int = 3
	if a == 3 { //条件表达式没有括号
		fmt.Println("a==3")
	}

	//支持一个初始化表达式, 初始化字句和条件表达式直接需要用分号分隔
	if b := 3; b == 3 {
		fmt.Println("b==3")
	}

	if a := 5; a == 5 {
		fmt.Println("a==5")
	} else { //左大括号必须和条件语句或else在同一行
		fmt.Println("a!=5")
	}

	if a := 10; a > 3 {
		fmt.Println("a>3")
	} else if a < 3 {
		fmt.Println("a<3")
	} else if a == 3 {
		fmt.Println("a==3")
	} else {
		fmt.Println("error")
	}
	s := "屌丝"
	if s == "王思聪" {
		fmt.Println("左手一个跑车")
	}
}
