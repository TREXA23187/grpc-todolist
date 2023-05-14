package pwd

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// HashPwd hash加密密码
func HashPwd(pwd string, cost int) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), cost)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(hash), nil
}

// CheckPwd 验证密码 hash之后的密码 输入的密码
func CheckPwd(hashPwd string, pwd string) bool {
	byteHash := []byte(hashPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, []byte(pwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
