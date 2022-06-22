package routers

import (
	"aceniubi/controllers"
	"aceniubi/jwt"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"net/http"
	"strings"
)

func init() {
	/*
	  配置跨域
	*/
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		//AllowOrigins:      []string{"https://192.168.0.102"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	var ApiAuthFilter = func(ctx *context.Context) {
			token := ctx.Request.Header.Get("Authorization")
	     	kv := strings.Split(token, " ")
	     	if len(kv) != 2 || kv[0] != "Bearer" {
				panic("AuthString无效")
		    }
		    tokenString := kv[1]
	     	fmt.Printf(tokenString)
			user,err:=jwt.ValidateToken(tokenString)
			fmt.Println(err)
		if err!=nil {
			http.Error(ctx.ResponseWriter, "Token verification not pass", http.StatusUnauthorized)
		} else {
			fmt.Println(user)
		}
		}
    beego.InsertFilter("/token/*",beego.BeforeRouter,ApiAuthFilter)
	beego.Router("/", &controllers.MainController{})
	//beego.Router("/chaxun", &controllers.UserchaController{})  //查询mysql
	//beego.Router("/api/cheak",&controllers.JwtController{},"post:CheakToken") //测试api：token
	//beego.Router("/api/mdd",&controllers.JwtController{},"post:MD5test") //测试api：MD5

	/*
	   name:操作数据库
	   author:wzh
	*/
	mysql:=beego.NewNamespace("/api",
		beego.NSNamespace("/mysql",
	beego.NSRouter("/look",&controllers.MySqlController{}),  //查询mysql
	beego.NSRouter("/update",&controllers.MySqlController{},"post:Update"), //更新
	beego.NSRouter("/delete",&controllers.MySqlController{},"post:Delete")))  //删除
	/*
		name:登陆注册
		author:wzh
	*/
	login:=beego.NewNamespace("/api",
		beego.NSNamespace("/user",
			beego.NSRouter("login",&controllers.JwtController{},"post:Login"), //添加 注册账户
			beego.NSRouter("registered",&controllers.MySqlController{},"post:Registered"))) //登陆
	/*
		name测试api
		author:wzh
	*/
	test:=beego.NewNamespace("/token",
		beego.NSNamespace("/api",
			beego.NSRouter("/hastoken",&controllers.JwtController{},"post:Hasasolt"), //测试api：has256加solt
			beego.NSRouter("/haschack",&controllers.JwtController{},"post:Haschack"), //测试api：验证has256加solt
			beego.NSRouter("/chacktoken",&controllers.JwtController{},"post:Chacktoken"))) //测试api：验证has256加solt
	beego.AddNamespace(login)
	beego.AddNamespace(test)
	beego.AddNamespace(mysql)
}
