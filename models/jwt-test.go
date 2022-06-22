package models

import "time"

type Jwt struct {
	Id         int       `json:"id" orm:"column(id)"`
	Username    string    `json:"account" orm:"column(username)"`
	Password   string    `json:"password" orm:"column(password)"`
	Created_at time.Time `gorm:"column:create_at;default:null" json:"create_at"`
	Updated_at time.Time `gorm:"column:updated_at;default:null" json:"updated_at"`
}
func (Jwt) TableName() string{
	return "wettest"
}

type Password struct {
	Password   string    `json:"password" orm:"column(password)"`
}
func (Password) TableName() string{
	return "wettest"
}


type Username struct {
	Username   string    `json:"username" orm:"column(username)"`
}
func (Username) TableName() string{
	return "wettest"
}