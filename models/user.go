package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID        int        `gorm:"primary_key" json:"id" validate:"min=1"`
	Username  string     `gorm:"column:username" json:"username" validate:"min=3,max=20,alphanum"`
	Password  string     `gorm:"column:password" json:"password" validate:"min=3,max=20"`
	Phone     string     `gorm:"column:phone" json:"phone" validate:"len=11"`
	CreatedAt int        `json:"createdAt"`
	UpdatedAt int        `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

func (User) TableName() string {
	return TablePrefix + "user"
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("CreatedAt", time.Now().Unix())
	_ = scope.SetColumn("UpdatedAt", time.Now().Unix())
	_ = scope.SetColumn("Password", GetPassword(u.Password))
	_ = scope.SetColumn("ID", 0)
	return nil
}

func (u *User) BeforeUpdate(scope *gorm.Scope) error {
	_ = scope.SetColumn("UpdatedAt", time.Now().Unix())
	return nil
}

func (u *User) CheckAuth() bool {

	var user User

	db.Select("id").Where(User{
		Username: u.Username,
		Password: GetPassword(u.Password),
	}).First(&user)

	if user.ID > 0 {
		return true
	}
	return false
}

func (u *User) Insert() bool {

	var user User
	db.Select("id").Where(User{
		Username: u.Username,
	}).First(&user)

	//fmt.Println(user)
	if user.ID > 0 {
		return false
	}

	err := db.Create(u).Error
	if err != nil {
		return false
	}

	return true
}

func (u *User) PasswdChange(newUser *User) bool {

	var user User

	db.Where(User{
		Username: u.Username,
		Password: GetPassword(u.Password),
		Phone:    u.Phone,
	}).First(&user)

	if user.ID > 0 {
		user.Password = GetPassword(newUser.Password)
		db.Save(&user)
		return true
	}

	return false

}

func (u *User) Del() {
	db.Where("id=?", u.ID).Delete(&User{})
}

func (u *User) Info() bool {
	var where User

	if u.ID > 0 {
		where.ID = u.ID
	}
	if u.Username != "" {
		where.Username = u.Username
	}

	u.ID = 0
	db.Where(where).First(u)
	u.Password = "******"

	if u.ID > 0 {
		return true
	}
	return false
}
