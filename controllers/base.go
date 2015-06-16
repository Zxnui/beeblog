package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

type BaseController struct {
	beego.Controller
	i18n.Locale
}

func init() {
	i18n.SetMessage("en-US", "conf/local_en-US.ini")
	i18n.SetMessage("zh-CN", "conf/local_zh-CN.ini")

	beego.AddFuncMap("i18n", i18n.Tr)
}

func (this *BaseController) Prepare() {
	lang := this.GetString("lang")
	if lang == "zh-CN" {
		this.Lang = lang
	} else {
		this.Lang = "en-US"
	}

	this.Data["Lang"] = this.Lang
}
