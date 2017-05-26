package main

import (
	"fmt"
)

type USB interface { //定义接口
	Name() string //有返回值
	Connecter     //嵌入接口
}

type Connecter interface {
	Connect() //没有返回值
}

type PhoneConnecter struct {
	name string
}

func main() {
	var a USB
	a = PhoneConnecter{"phone"}
	a.Connect()
	Disconnect(a)
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("connecter:", pc.name)
}

func Disconnect(usb interface{}) { //参数是空接口时，用switch
	/*	if pc, ok := usb.(PhoneConnecter); ok { //判断类型
		fmt.Println("Disconnect", pc.name)
		return
	}*/

	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnect", v.name)
	default:

		fmt.Println("未知的设备")
	}

}
