package controllers

import (
	"aceniubi/global"
	_ "aceniubi/global"
	"aceniubi/jwt"
	_ "aceniubi/jwt"
	"aceniubi/models"
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

////查询
//type UserchaController struct {
//	JsonController
//}
//func (c*UserchaController) Get(){
//	user:= models.ChaXun()
//	c.ApiJsonReturn("请求成功", 200, user)
//}
type MySqlController struct {
	JsonController
}
//添加 注册账户  has256加密
func (c *MySqlController) Registered() {
	Username :=c.GetString("username")
	Password :=c.GetString("password")
	abs := models.Jwt{}
	//srcByte := []byte(Password)
	//sha256New := sha256.New()
	//sha256Bytes := sha256New.Sum(srcByte)
	//Password = hex.EncodeToString(sha256Bytes)
	Password, err :=jwt.HashAndSalt(Password)
	usernamecount :=models.Username{}
	var count int64
	global.Db.Model(usernamecount).Where("username",Username).Count(&count)
    if count==0{
	if err!=nil{
		panic(err)
	}else {
	result := global.Db.Model(abs).Create(map[string]interface{}{
		"username": Username,
		"password": Password,
	})
	if result!=nil{
		c.ApiJsonReturn("Registered successfully",200,"ok")
	} else{
		c.ApiJsonReturn("Registration failed",100,false)
	}
	}
    }else{
		c.ApiJsonReturn("he account has been registered. Please apply for a new account",100,false)
	}
}

//查询
func (c *MySqlController) Get() {
	user := []models.User{}
	// user := []utils.User{}
	global.Db.Find(&user)
    if user!=nil{
    c.ApiJsonReturn("请求成功",200,user)
    } else{
		c.ApiJsonReturn("请求失败",100,false)
	}
}

//更新
func (c *MySqlController) Update() {
    ID :=c.GetString("id")
	Active_name :=c.GetString("active_name")
	Active_place :=c.GetString("active_place")
	Time_begin :=c.GetString("time_begin")
	Time_end :=c.GetString("time_end")
	Active_nature :=c.GetString("active_nature")
	Special_resources :=c.GetString("special_resources")
	Form :=c.GetString("form")

	 information := []models.Information{}
	global.Db.Model(&information).Where("user_id",ID).Updates(map[string]interface{}{
		"active_name": Active_name,
		"active_place": Active_place,
		"time_begin": Time_begin,
		"time_end": Time_end,
		"active_nature": Active_nature,
		"special_resources": Special_resources,
		"form": Form})
	//information.Active_name = Active_name
	//information.Active_place = Active_place
	//information.Time_begin = Time_begin
	//information.Time_end = Time_end
	//information.Active_nature = Active_nature
	//information.Special_resources = Special_resources
	//information.Form = Form
	if information!=nil{
		c.ApiJsonReturn("请求成功",200,information)
	} else{
		c.ApiJsonReturn("请求失败",100,false)
	}
}

//删除
func (c *MySqlController) Delete() {
	ID :=c.GetString("id")
	information := []models.Information{}
	global.Db.Where("user_id",ID).Delete(&information)
	c.Ctx.WriteString("删除数据成功")
	if information!=nil{
		c.ApiJsonReturn("请求成功",200,information)
	} else{
		c.ApiJsonReturn("请求失败",100,false)
	}
}

type JwtController struct {
	JsonController
}

type User struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
}
//获取token(登陆)
func (c *JwtController) Login(){
	Username :=c.GetString("username")
	Passwordget :=c.GetString("password")
	//srcByte := []byte(Passwordget)
	//sha256New := sha256.New()
	//sha256Bytes := sha256New.Sum(srcByte)
	//Passwordget = hex.EncodeToString(sha256Bytes)

	Password := models.Password{}
	global.Db.Where("username",Username).First(&Password)
	Passwordsql := Password.Password  //数据库密码加密后的hash256加盐
	user := User{Username, Passwordget}
	token,err := jwt.GenerateToken((*jwt.User)(&user),0) //获取token
	//date,_:=jwt.ValidateToken(token)  //解密出token的password
	date :=jwt.ComparePasswords(Passwordsql,Passwordget)
     if date ==false{
		 c.ApiJsonReturn("login fail！please check your password or username！",100,false)
	 } else{
	 	if err!=nil{
			panic(err)
		}else {
			c.ApiJsonReturn("login success",200,token)
		}
	}
	//if err != nil {
	//	c.ApiJsonReturn("token get fail",100,false)
	//}else {
	//	if date.Password != Passwordget {
	//		c.ApiJsonReturn("login fail！please check your password or username！",100,false)
	//	}else {
	//		c.ApiJsonReturn("login success",200,token)
	//	}
	}

//验证token
//func (c *JwtController) CheakToken(){
//	date,err :=jwt.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjEyMzQiLCJwYXNzd29yZCI6IjMxMzIzMzM0ZTNiMGM0NDI5OGZjMWMxNDlhZmJmNGM4OTk2ZmI5MjQyN2FlNDFlNDY0OWI5MzRjYTQ5NTk5MWI3ODUyYjg1NSIsImV4cCI6MTY1NDg1NTIzOSwiaWF0IjoxNjU0ODU0NjM5LCJpc3MiOiIxMjM0In0.paoH1bL-IYCDZOt4BAa4dEPhJFFzW3-Xo63Hzl8L3PQ")
//	if err != nil {
//		fmt.Println(err)
//	}else {
//		//获取jwt
//		c.ApiJsonReturn("token",100,date)
//	}
//}

//func (c*JwtController) CheakToken(){
//	date:=jwt.GetHeaderTokenValue("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjEyMzQiLCJwYXNzd29yZCI6IjMxMzIzMzM0ZTNiMGM0NDI5OGZjMWMxNDlhZmJmNGM4OTk2ZmI5MjQyN2FlNDFlNDY0OWI5MzRjYTQ5NTk5MWI3ODUyYjg1NSIsImV4cCI6MTY1NDg2MzkzOCwiaWF0IjoxNjU0ODYzMzM4LCJpc3MiOiIxMjM0In0.NtNxYcz22Sw4qd6i_XOfjeYjBNWF3UDuBxyYZ4lQBNA")
//		c.ApiJsonReturn("请求成功",200,date)
//}
// 生成32位MD5
//func (c *JwtController) MD5test(){
//       a:="123456"
//        date:= jwt.MD5(a)
//        c.ApiJsonReturn("md5加密成功",200,date)
//
//}

//测试has256加密加盐
func (c*JwtController) Hasasolt(){
	a:="12344"
	date, err :=jwt.HashAndSalt(a)
	if err!=nil{
		c.ApiJsonReturn("has加密失败",100,err)
	}else{
	c.ApiJsonReturn("has密码",200,date)
	}
}
//验证hash加密加盐
func (c*JwtController) Haschack(){
	a:="12344"
	hsapaaword :="$2a$04$RbvdDhz6CSVQsfuTFzVDnOGwt/ZBpkmqG7AAKS07ChzrnG3evIMIK"
	date :=jwt.ComparePasswords(hsapaaword,a)
		c.ApiJsonReturn("has密码",200,date)
}
//检测token
func (c*JwtController) Chacktoken(){
	fmt.Println(jwt.Check)
	//return jwt.Check
	//token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjEyMzQ1NiIsInBhc3N3b3JkIjoiMTIzNDU2IiwiZXhwIjoxNjU1MDEwNzIyLCJpYXQiOjE2NTUwMTAxMjIsImlzcyI6IjEyMzQ1NiJ9.CCCKve0986jLqTk_v7RkVtYjxXZXY49u6_MZT7Yf34w"
	//user,err:=jwt.ValidateToken(token)
	//if err !=nil{
	//	panic(err)
	//} else {
	//	c.ApiJsonReturn("token获取成功",200,user)
	//}
}