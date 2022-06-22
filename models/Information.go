package models

import "time"

type Information struct {
	Id         int       `json:"id" orm:"column(id)"`
	User_id    string    `json:"user_id" orm:"column(user_id)"`
	Active_name   string    `json:"active_name" orm:"column(active_name)"`
	Active_place       string    `json:"active_place" orm:"column(active_place)"`
	Time_begin        string    `json:"time_begin" orm:"column(time_begin)"`
	Time_end        string    `json:"time_end" orm:"column(time_end)"`
	Active_nature     string    `json:"active_nature" orm:"column(active_nature)"`
	Special_resources      string    `json:"special_resources" orm:"column(special_resources)"`
	Form      string    `json:"form" orm:"column(form)"`
	Created_at time.Time `gorm:"column:create_at;default:null" json:"create_at"`
	Updated_at time.Time `gorm:"column:updated_at;default:null" json:"updated_at"`
}

func (Information) InformationName() string{
	return "information"
}