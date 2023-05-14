package repository

import (
	"errors"
	"gorm.io/gorm"
	"grpc-todolist/user/internal/service"
	"grpc-todolist/user/pkg/util/pwd"
)

type User struct {
	UserId   uint   `gorm:"primarykey"`
	UserName string `gorm:"unique"`
	NickName string
	Password string
}

const (
	PasswordCost = 12 // 密码加密难度
)

func (user *User) CheckUserExist(req *service.UserRequest) bool {
	if err := DB.Where("user_name=?", req.UserName).First(&user).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (user *User) ShowUserInfo(req *service.UserRequest) (err error) {
	if exist := user.CheckUserExist(req); exist {
		return nil
	}

	return errors.New("user not exist")
}

func (*User) CreateUser(req *service.UserRequest) (err error) {
	var count int64
	DB.Where("user_name=?", req.UserName).Count(&count)
	if count != 0 {
		return errors.New("username exists")
	}

	user := User{
		UserName: req.UserName,
		NickName: req.NickName,
	}

	// password digest
	_ = user.SetPassword(req.Password)

	err = DB.Create(&user).Error
	return err
}

func (user *User) SetPassword(password string) (err error) {
	hashPwd, err := pwd.HashPwd(password, PasswordCost)
	if err != nil {
		return err
	}

	user.Password = hashPwd
	return nil
}

func BuildUser(item User) *service.UserModel {
	userModel := service.UserModel{
		UserID:   uint64(item.UserId),
		UserName: item.UserName,
		NickName: item.NickName,
	}

	return &userModel
}
