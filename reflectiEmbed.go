package main //嵌入的反射

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}
type Manager struct {
	User
	title string
}

func main() {
	m := Manager{User: User{1, "sh", 34}, title: "title"}
	t := reflect.TypeOf(m)

	fmt.Printf("%#v \n ", t.FieldByIndex([]int{0, 1})) //传slice 取匿名字段中的字段
}
