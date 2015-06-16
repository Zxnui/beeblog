package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	BaseController
}

func (this *CategoryController) Get() {
	op := this.Input().Get("op")

	switch op {
	case "add":
		name := this.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 301)
		return
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}

		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return

	}
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	this.TplNames = "category.html"
	this.Data["IsCategory"] = true
	var err error
	this.Data["Categories"], err = models.GetAllCategory()
	if err != nil {
		beego.Error(err)
	}

}
