// BubbleSort project 冒泡排序
package main

import (
	"fmt"
)

func main() {
	//数组定义是大括号，
	a := [5]int{5, 2, 6, 3, 9}

	fmt.Println(a)
	lenth := len(a)
	for i := 0; i < lenth; i++ {
		for j := i + 1; j < lenth; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)
}
