// @APIVersion 1.0.0
// @Title Portfolio Curriculo
// @Description Api present resume
// @Contact kaim7said@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
// @Schemes http,https
// @Host knowmoreaboutme.tk:8081
package routers

import (
	"gitlab.com/portfoliocurriculo/cvbackend/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/cv",
		beego.NSNamespace("/resume",
			beego.NSInclude(
				&controllers.ResumeController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
