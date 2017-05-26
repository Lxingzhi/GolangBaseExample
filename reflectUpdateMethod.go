package main //通过反射来调用方法

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	u := User{2, "fdf", 23}
	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello") //得到方法

	args := []reflect.Value{reflect.ValueOf("jone")} //根据参数加slice
	name := mv.Call(args)                            //有返回值
	fmt.Println("fanhuide name ", name)
}

func (u User) Hello(name string) string { //有一个参数

	fmt.Println("hello", name, ",my name is", u.Name)
	return name
}
