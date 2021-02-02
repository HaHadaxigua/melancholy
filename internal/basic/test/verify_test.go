/******
** @date : 1/16/2021 1:16 AM
** @author : zrx
** @description:
******/
package test

import (
	tools2 "github.com/HaHadaxigua/melancholy/internal/basic/tools"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckEmail(t *testing.T) {
	emailOk := "123@qq.com"
	assert.Equal(t, true, tools2.CheckEmail(emailOk))
	emailFalse := "122123.com"
	assert.Equal(t, false, tools2.CheckEmail(emailFalse))
}

func TestCheckUsername(t *testing.T) {
	userNameOk := "ax21"
	assert.Equal(t, true, tools2.CheckUsername(userNameOk))
	userNameFailed := "@!"
	assert.Equal(t, false, tools2.CheckUsername(userNameFailed))
}
