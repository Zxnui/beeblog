package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type ReplyController struct {
	BaseController
}

func (this *ReplyController) Post() {

}

func (this *ReplyController) Get() {

}

func (this *ReplyController) Add() {
	tid := this.Input().Get("tid")

	name := this.Input().Get("nickname")
	if len(name) == 0 {
		name = "匿名"
	}
	err := models.AddReply(tid, name, this.Input().Get("content"))

	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic/view/"+tid, 302)
}

func (this *ReplyController) Delete() {
	if !checkAccount(this.Ctx) {
		return
	}
	tid := this.Input().Get("tid")
	err := models.DeleteReply(this.Input().Get("rid"))
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic/view/"+tid, 302)
}
