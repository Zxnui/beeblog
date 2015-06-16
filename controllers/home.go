package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.Data["IsHome"] = true
	this.TplNames = "home.html"

	this.Data["IsLogin"] = checkAccount(this.Ctx)
	var err error
	this.Data["Topics"], err = models.GetAllTopic(this.Input().Get("cate"), this.Input().Get("label"), true)
	if err != nil {
		beego.Error(err)
	}

	categories, err := models.GetAllCategory()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Categories"] = categories

}
