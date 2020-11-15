package tools

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/HaHadaxigua/melancholy/pkg/msg"
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
