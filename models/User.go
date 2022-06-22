package models

import (
	"time"
)

type User struct {
	Id         int       `json:"id" orm:"column(id)"`
	Account    string    `json:"account" orm:"column(account)"`
	Password   string    `json:"password" orm:"column(password)"`
	Name       string    `json:"name" orm:"column(name)"`
	Age        string    `json:"age" orm:"column(age)"`
	Sex        string    `json:"sex" orm:"column(sex)"`
	Number     string    `json:"number" orm:"column(number)"`
	State      string    `json:"state" orm:"column(state)"`
	Created_at time.Time `gorm:"column:create_at;default:null" json:"create_at"`
	Updated_at time.Time `gorm:"column:updated_at;default:null" json:"updated_at"`
}
//func init() {
//	// 需要在init中注册定义的model
//	orm.RegisterModel(new(User))
//	orm.Debug = true // 是否开启调试模式 调试模式下会打印出sql语句
//}
//
//func QuerUser(user *User) bool {
//	var userData []User
//	// 获取 QueryBuilder 对象. 需要指定数据库驱动参数。
//	// 第二个返回值是错误对象，在这里略过
//	qb, _ := orm.NewQueryBuilder("mysql")
//	// 构建查询对象
//	qb.Select("*").From("user").
//		OrderBy("id").Desc().
//		Limit(10).Offset(0)
//	// 导出 SQL 语句
//	sql := qb.String()
//	// 执行 SQL 语句
//	o := orm.NewOrm()
//	_, _ = o.Raw(sql).QueryRows(&userData)
//	fmt.Println(userData)
//	return true
//}
//
//func ChaXun() *User{
//	// 创建orm对象
//	o := orm.NewOrm()
//	// 获取 QuerySeter 对象，并设置表名orders
//	qs := o.QueryTable("user")
//
//	// 定义保存查询结果的变量
//	var users []User
//
//	// 使用QuerySeter 对象构造查询条件，并执行查询。
//	num, err := qs.Filter("id", "20").  // 设置查询条件
//		Limit(10). // 限制返回行数
//		All(&users, "id", "name") // All 执行查询，并且返回结果，这里指定返回id和username字段，结果保存在users变量
//	// 上面代码的等价sql: SELECT T0.`id`, T0.`username` FROM `users` T0 WHERE T0.`city` = 'shenzhen' AND T0.`init_time` > '2019-06-28 22:00:00' LIMIT 10
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("结果行数:", num)
//	fmt.Println("结果行数:", users)
//	return nil
//}
func (User) TableName() string{
	return "user"
}