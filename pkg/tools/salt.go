package tools

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/HaHadaxigua/melancholy/pkg/msg"
	"golang.org/x/crypto/bcrypt"
)

//GenerateSalt 生成随机的32位salt
func GenerateSalt() (string, error) {
	//生成随机盐
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		return "", msg.GenerateSaltErr
	}
	return hex.EncodeToString(salt), nil
	//生成密文
	//dk := pbkdf2.Key([]byte("mimashi1323"), salt, 1, 32, sha256.New)
}

//EncryptPassword 加密密码
func EncryptPassword(pwd, salt string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd+salt), bcrypt.DefaultCost)
	if err != nil {
		e := msg.EncryptPasswordErr
		e.Data = err.Error()
		return "", e
	}
	return string(hash), nil
}

//VerifyPassword 验证密码是否正确
func VerifyPassword(encodePwd, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodePwd), []byte(password)) //验证（对比）
	if err != nil {
		return false
	}
	return true
}
