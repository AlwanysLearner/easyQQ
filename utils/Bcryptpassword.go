package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// BcryptEncode 加密密码
func BcryptEncode(str string) string {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("调用bcrypt,加密异常，HASH编码失败:%w", err)
	}
	return string(bcryptPassword)
}

func VerifyPassword(sourcePwd, bcryptPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(bcryptPwd), []byte(sourcePwd))
	if err != nil {
		return false
	}
	return true
}
