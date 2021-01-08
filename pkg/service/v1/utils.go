/******
** @date : 1/9/2021 1:14 AM
** @author : zrx
******/
package v1

import (
	"github.com/HaHadaxigua/melancholy/pkg/consts"
	"regexp"
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
