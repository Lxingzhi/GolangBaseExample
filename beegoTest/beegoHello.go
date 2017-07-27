package main

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {

	this.Ctx.WriteString("Hello word")
}

func main() {
	beego.Router("/", &HomeController{}) //注册
	beego.Run()
}
