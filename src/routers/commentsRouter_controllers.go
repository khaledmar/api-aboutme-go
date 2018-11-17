package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["gitlab.com/portfoliocurriculo/cvbackend/controllers:ResumeController"] = append(beego.GlobalControllerRouter["gitlab.com/portfoliocurriculo/cvbackend/controllers:ResumeController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:resumeID`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
