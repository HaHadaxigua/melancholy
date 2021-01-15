/******
** @date : 1/16/2021 1:16 AM
** @author : zrx
** @description:
******/
package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckEmail(t *testing.T) {
	emailOk := "123@qq.com"
	assert.Equal(t, true, CheckEmail(emailOk))
	emailFalse := "122123.com"
	assert.Equal(t, false, CheckEmail(emailFalse))
}

func TestCheckUsername(t *testing.T) {
	userNameOk := "ax21"
	assert.Equal(t, true, CheckUsername(userNameOk))
	userNameFailed := "@!"
	assert.Equal(t, false, CheckUsername(userNameFailed))
}
