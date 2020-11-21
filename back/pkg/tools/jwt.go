package tools

import (
	"github.com/HaHadaxigua/melancholy/pkg/conf"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(conf.C.Application.JwtSecret)

type Claims struct {
	Email string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}

//GenerateToken 生成token
func GenerateToken(email, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour) // 设置两小时后过期

	claims := Claims{
		email,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "admin",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}


//ParseToken 解析token
func ParseToken(token string)(*Claims, error){
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}