package main

import (
	_ "gitlab.com/portfoliocurriculo/cvbackend/routers"
	"github.com/astaxie/beego"
)

type PingController struct {
	beego.Controller
}

func (this *PingController) Get() {
	this.Ctx.WriteString("hello world")
}


	
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
		
	beego.Router("/ping", &PingController{})
	
	beego.Run()
}
