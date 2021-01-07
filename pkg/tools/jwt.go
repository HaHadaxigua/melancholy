package tools

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/pkg/conf"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"time"
)

var jwtSecret = []byte(conf.C.Application.AppSecret)

// custom claims
type userStdClaims struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func JwtGenerateToken(userID int, email, password string, d time.Duration) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(d * time.Hour) // 设置两小时后过期

	uClaims := userStdClaims{
		email,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  nowTime.Unix(),
			Issuer:    conf.C.Application.AppIss,
			Id:        fmt.Sprintf("%d", userID),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, uClaims)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		logrus.WithError(err).Fatal("config is wrong, can not generate jwt")
	}
	return token, err
}

func JwtParseToken(token string) (*userStdClaims, error) {
	if token == "" {
		return nil, fmt.Errorf("no token is found in Authorization Bearer")
	}

	tokenClaims, err := jwt.ParseWithClaims(token, &userStdClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*userStdClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
