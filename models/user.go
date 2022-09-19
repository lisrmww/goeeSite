package models

import (
	"gorm.io/gorm"
	"goser/dao/mysql"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"index;not null;size(64);unique;comment:用户名"`
	Password string `json:"password" gorm:"not null;size(256);comment:密码"`
	UserPic  string `json:"userPic" gorm:"size(100);default:http://127.0.0.1:3009/static/upload/user/avatar.png;comment:用户头像"`
}

func init() {
	if !mysql.MySQL.Migrator().HasTable(&User{}) {
		err := mysql.MySQL.AutoMigrate(&User{})
		admin := &User{
			Username: "admin",
			Password: "123456",
		}
		mysql.MySQL.Create(&admin)
		if err != nil {
			return
		}
	}
}

func (User) Create() {

}

func (User) Delete() {

}

func (User) Update() {

}

func (User) Retrieve(user *User) {
	mysql.MySQL.First(&user)
}
