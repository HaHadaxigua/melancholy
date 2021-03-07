package tools

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/HaHadaxigua/melancholy/internal/response"
	"golang.org/x/crypto/bcrypt"
)

//GenerateSalt 生成随机的32位salt
func GenerateSalt() (string, error) {
	//生成随机盐
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		return "", response.GenerateSaltErr
	}
	return hex.EncodeToString(salt), nil
	//生成密文
	//dk := pbkdf2.Key([]byte("mimashi1323"), salt, 1, 32, sha256.New)
}

//EncryptPassword 加密密码
func EncryptPassword(pwd, salt string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd+salt), bcrypt.DefaultCost)
	if err != nil {
		e := response.EncryptPasswordErr
		e.Data = err.Error()
		return "", e
	}
	return string(hash), nil
}

//VerifyPassword 验证密码是否正确
func VerifyPassword(encodePwd, password, salt string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(encodePwd), []byte(password+salt)); err != nil {
		return false
	} //验证（对比）
	return true
}
