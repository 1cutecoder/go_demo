package main

import "fmt"

func main() {
	const a int = 10
	fmt.Println("a = ", a)
	var (
		c int
		d float64
	)
	c, d = 11, 3.14
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
	const (
		e       = iota
		f       = iota
		g       = iota
		h, j, k = iota, iota, iota
	)
	fmt.Println("e = ", e)
	fmt.Println("f = ", f)
	fmt.Println("g = ", g)
	fmt.Println("h = ", h)
	fmt.Println("j = ", j)
	fmt.Println("k = ", k)
	var t bool
	fmt.Println("t = ", t)
	t = a == 10
	fmt.Println("t= ", t)
	t = false
	fmt.Println("t= ", t)
	var f1 float32
	f1 = 3.14
	fmt.Println("f1 = ", f1)

	var str string
	str = "def"
	ch := str[0]
	fmt.Printf("str =%s, len =%d\n", str, len(str))
	fmt.Printf("str[0] = %c,ch=%c\n", str[0], ch)

	//`(反引号)括起的字符串为Raw字符串，即字符串在代码中的形式就是打印时的形式，它没有字符转义，换行也将原样输出。
	str2 := `hello
mike \n \r测试
`
	fmt.Println("str2 = ", str2)
	var ch1 byte
	ch1 = 97
	fmt.Println("ch1 = ", ch1)
	fmt.Printf("ch1 = %c\n", ch1)
	ch1 = 'a'
	fmt.Printf("%c, %d\n", ch1, ch1)
	fmt.Printf("大写:%d,小写:%d\n", 'A', 'a')
	fmt.Printf("大写:%c,小写:%c\n", 'a'-32, 'a')
	fmt.Printf("大写:%c,小写:%c\n", 'A', 'A'+32)

}
