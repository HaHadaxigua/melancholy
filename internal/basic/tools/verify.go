/******
** @date : 1/16/2021 1:16 AM
** @author : zrx
** @description:
******/
package tools

import (
	"github.com/HaHadaxigua/melancholy/internal/consts"
	"regexp"
	"strings"
)

func CheckUsername(username string) bool {
	if ok, _ := regexp.MatchString(consts.UserNamePattern, username); !ok {
		return false
	}
	return true
}

func CheckPassword(password string) bool {
	if ok, _ := regexp.MatchString(consts.PasswordPattern, password); !ok {
		return false
	}
	return true
}

func CheckEmail(email string) bool {
	return regexp.MustCompile(consts.EmailPattern).MatchString(email)
}

func VerifyFileName(filename string) bool {
	return regexp.MustCompile(consts.FileNamePattern).MatchString(filename) && !strings.Contains(filename, " ")
}
