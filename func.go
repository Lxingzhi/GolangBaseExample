package main

import (
	"fmt"
)

func main() {
	A()
	B()
	C()

	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i=", i)
		defer func() {
			fmt.Println("defer_closure i=", i)
		}()
		fs[i] = func() {
			fmt.Println("closeure i=", i)
		}
	}
	for _, f := range fs {
		f()
	}

}
func A() {
	fmt.Println("func  a")
}
func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in b")
		}
	}()

	panic("func  b")
}

func C() {
	fmt.Println("func  c")
}
