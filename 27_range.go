package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	str := "abc"
	for i := 0; i < len(str); i++ {
		fmt.Printf("1-str[%d]=%c\n ", i, str[i])
	}

	for i, data := range str {
		fmt.Printf("2-str[%d]=%c\n ", i, data)
	}

	for i := range str {
		fmt.Printf("3-str[%d]=%c\n ", i, str[i])
	}

	for i, _ := range str {
		fmt.Printf("4-str[%d]=%c\n ", i, str[i])
	}

	for _, s := range str {
		fmt.Printf("5-%c\n ", s)
	}
}
