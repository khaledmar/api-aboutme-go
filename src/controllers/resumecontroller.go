package controllers

import (
	"github.com/astaxie/beego"
	"gitlab.com/portfoliocurriculo/cvbackend/models"
)

// Operations about object
type ResumeController struct {
	beego.Controller
}

// @Title Get
// @Description find object by resumeID
// @Param	resumeID		path 	string	true		"the resumeID you want to get"
// @Success 200 {object} models.Resume
// @Failure 403 :resumeID is empty
// @router /:resumeID [get]
func (o *ResumeController) Get() {
	resumeID := o.Ctx.Input.Param(":resumeID")
	if resumeID != "" {
		ob, err := models.GetOneResume(resumeID)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}
