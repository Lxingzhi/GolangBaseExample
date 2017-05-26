package main

import (
	"fmt"
)

type nimingstring struct {
	string
	int
}

type human struct {
	Sex int
}
type teacher struct {
	human
	Name string
	Age  int
}

type student struct {
	human
	Name string
	Age  int
}

func main() {
	//初始化
	t := teacher{Name: "as", Age: 2, human: human{Sex: 1}}
	st := student{Name: "aasasasas", Age: 22, human: human{Sex: 0}}

	a := struct {
		Name string
		Age  int
	}{
		Name: "hif",
		Age:  10,
	}

	fmt.Println(a)

	s := nimingstring{"sad", 1}
	fmt.Println(s)

	fmt.Println(t)
	fmt.Println(st)
}
