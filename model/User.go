package model

import (
	"encoding/base64"
	"gin-blog/utils/messages"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

// CheckUser 用户校验
func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return messages.ERROR_USERNAME_USED
	}
	return messages.SUCCSE
}

func VerifyUserExists(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return messages.SUCCSE
	}
	return messages.ERROR_USER_NOT_EXIST
}

// CreateUser 创建用户
func CreateUser(data *User) int {
	data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return messages.ERROR
	}
	return messages.SUCCSE
}

// GetUsers 获取用户信息
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// EditUserInfo 用户信息编辑
func EditUserInfo(id int, data *User) int {
	var user User
	var maps = make(map[string]any)
	maps["username"] = data.Username
	maps["role"] = data.Role
	if data.Password != "" {
		maps["password"] = ScryptPw(data.Password)
	}
	err = db.Model(&user).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return messages.ERROR
	}
	return messages.SUCCSE
}

// DeleteUser 删除用户
func DeleteUser(name string) int {
	var user User
	err = db.Where("username = ?", name).Delete(&user).Error
	if err != nil {
		return messages.ERROR
	}
	return messages.SUCCSE
}

// ScryptPw 密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}

	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

// CheckLogin 登录验证
func CheckLogin(username string, password string) int {
	var user User

	db.Where("username = ?", username).First(&user)

	if user.ID == 0 {
		return messages.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password) != user.Password {
		return messages.ERROR_PASSWORD_WRONG
	}
	if user.Role != 0 {
		return messages.ERROR_USER_NO_RIGHT
	}
	return messages.SUCCSE
}
