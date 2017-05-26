package main

import (
	"fmt"
)

func main() {
	//slice []内没有东西是slice切片不是数组
	/*var s1 []int
	fmt.Println(s1)
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := a[5:10] //a[5,6,7,8,9]
	fmt.Println(s2)*/
	s1 := make([]int, 3, 6) //拥有三个元素容量为6的slice
	fmt.Println(len(s1), cap(s1))
	fmt.Printf("%p\n", s1)
	s1 = append(s1, 2, 3, 4)
	fmt.Printf("%v %p\n", s1, s1)

	a, b := 1, 2
	A(a, b)

}
func A(s ...int) {
	s[0] = 3
	fmt.Println(s)
}
