package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("是哦啊是哦大事哦！")
	var a = 64
	b := strconv.Itoa(a)
	fmt.Println(b)
	a, _ = strconv.Atoi(b)
	fmt.Println(a)
}
