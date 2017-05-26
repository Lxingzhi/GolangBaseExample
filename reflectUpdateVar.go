package main //通过反射来修改属性

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
	Set(&u)
	fmt.Println(u)
}
func Set(o interface{}) {
	v := reflect.ValueOf(o)                            //取值
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() { //判断类型是否正确并判断是可以被修改的
		fmt.Println("不能修改")
		return
	} else {
		v = v.Elem()
	}
	f := v.FieldByName("Name")
	if !f.IsValid() { //是否找到这个字段
		fmt.Println("没找到")
		return
	}
	if f.Kind() == reflect.String { //根据字段修改内容 类型多 用switch
		f.SetString("reflectupdate")
	}
}
