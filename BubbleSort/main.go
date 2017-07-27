// BubbleSort project 冒泡排序
package main

import (
	"fmt"
)

func main() {
	a := [8]int{1, 4, 5, 2, 3, 66, 34, 74}
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
