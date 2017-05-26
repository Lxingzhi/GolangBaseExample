//常量练习
package main

import (
	"fmt"
	"strconv"
)

const sss = "12344"
const (
	a = len(sss)
	b
	c
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
LABEL:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println("continue" + strconv.Itoa(i))
			goto LABEL
		}
	}
}
