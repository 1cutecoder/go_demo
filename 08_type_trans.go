package main

import "fmt"

func main() {
	var flag bool
	flag = true
	fmt.Printf("flag = %t\n", flag)
	flag = false
	fmt.Printf("flag = %t\n", flag)
	var ch byte = 97
	//var a int = ch //err, cannot use ch (type byte) as type int in assignment
	var a int = int(ch)
	fmt.Printf("ch = %c\n", ch)
	fmt.Printf("a = %d\n", a)
}
