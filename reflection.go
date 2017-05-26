package main //反射

import (
	"fmt"
	"reflect" //反射
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (user User) Hello() {
	fmt.Println("hello")
}

func main() {
	u := User{1, "ok", 12}
	Info(u)
}

func Info(o interface{}) {
	if k := t.Kind(); k != reflect.Struct { //判断是不是struct类型
		fmt.Println("类型不对")
		return
	}

	t := reflect.TypeOf(o) //获得类型
	fmt.Println("Type", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")              //打印字段
	for i := 0; i < t.NumField(); i++ { //所有的字段
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s:%v = %v\n", f.Name, f.Type, val) //字段对应的值
	}
	for i := 0; i < t.NumMethod(); i++ { //获取方法的信息
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)

	}
}
