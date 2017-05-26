package main

import (
	"fmt"
)

type TZ int

type A struct {
	name string
}

func main() {
	a := A{}
	a.Print()
	fmt.Println(a.name)

	var tz TZ
	tz.Increase(100)
	fmt.Println(tz)
}
func (a *A) Print() {
	a.name = "122"
	fmt.Println("A")
}

func (tz *TZ) Increase(num int) {
	*tz += TZ(num)
}
