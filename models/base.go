package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var Cfg = beego.AppConfig

func init() {
	//获取服务器数据库相关信息
	dbhost := beego.AppConfig.String("db_host")
	dbport := beego.AppConfig.String("db_port")
	dbuser := beego.AppConfig.String("db_user")
	dbpassword := beego.AppConfig.String("db_pass")
	dbname := beego.AppConfig.String("db_name")
	dbIdle, _ := beego.AppConfig.Int("db_max_idle_conn")
	dbConn, _ := beego.AppConfig.Int("db_max_open_conn")

	dburl := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", dburl, dbIdle, dbConn)

	//注册模型
	orm.RegisterModel(new(Category), new(Topic), new(Comment))
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	//新建数据库表
	orm.RunSyncdb("default", false, true)

	//附件上传文件夹处理
	os.Mkdir("attachment", os.ModePerm)
}
