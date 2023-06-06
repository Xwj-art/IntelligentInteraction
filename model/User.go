package model

import (
	"Ai/errmsg"
	"Ai/utils"
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"type:int(10);primary_key;auto_increment" json:"id"`
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=1,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(30);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,eq=2" label:"角色码"`
}

// 对数据库进行操作

// 首先判断是否存在

func CheckUser(name string) int {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_NAME_USED
	}
	return errmsg.SUCCESS
}

func CreateUser(data *User) int {
	data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERR
	}
	return errmsg.SUCCESS
}

func GetUserInfo(id int) (int, User) {
	var UserInfo User
	err := db.Where("id = ?", id).First(&UserInfo).Error
	if err != nil {
		return errmsg.ERR, UserInfo
	}
	return errmsg.SUCCESS, UserInfo
}

// 密码加密

func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{23, 43, 12, 5, 7, 3, 84, 2}

	HashPwd, err := scrypt.Key([]byte(password), salt, 1<<10, 8, 1, KeyLen)
	utils.HandleErr(err, "scrypt加密错误")
	fwd := base64.StdEncoding.EncodeToString(HashPwd)
	return fwd
}

func DeleteUser(id int) int {
	err := db.Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return errmsg.ERR
	}
	return errmsg.SUCCESS
}

func UpdateUser(id int, user *User) int {
	var maps = make(map[string]any)
	maps["username"] = user.Username
	maps["role"] = user.Role
	err := db.Model(&User{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERR
	}
	return errmsg.SUCCESS
}

// 登陆验证
func CheckLogin(username, password string) (int, uint) {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST, 0
	}
	if ScryptPw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG, 0
	}
	if user.Role != 2 {
		return errmsg.USER_ROLE_COM, 0
	}
	return errmsg.SUCCESS, user.ID
}
