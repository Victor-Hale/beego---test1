package main

import (
	"aceniubi/global"
	_ "aceniubi/global"
	_ "aceniubi/routers"
	"errors"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error
func init(){
	//host:=beego.AppConfig.String("host")
	//port:=beego.AppConfig.String("port")
	//databaseName:=beego.AppConfig.String("databaseName")
	//userName:=beego.AppConfig.String("userName")
	//password:=beego.AppConfig.String("password")
	//orm.RegisterDataBase("default",
	//	"mysql",
	//	userName+":"+password+"@tcp("+host+":"+port+")/"+databaseName+"?charset=utf8&parseTime=true&loc=Local")
	//err :=orm.RegisterDataBase("default", "mysql", "third:123456@tcp(139.196.106.241)/third?charset=utf8")
	//if err != nil {
	//	return
	//}
	//orm.Debug = true
	dsn := "third:123456@tcp(139.196.106.241)/third?charset=utf8mb4&parseTime=True&loc=Local"
	global.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		 errors.New("err，数据库连接失败")
	}

}
func main() {
	beego.BConfig.WebConfig.AutoRender = false
	beego.Run()
}

