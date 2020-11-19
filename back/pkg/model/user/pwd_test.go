package user

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestGeneratePassword(t *testing.T){

	fmt.Println("====模拟注册====")
	u0 := User{}
	u0.Password = "pwd" //模拟注册是传递的密码
	hash, err := bcrypt.GenerateFromPassword([]byte(u0.Password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		fmt.Println(err)
	}
	encodePWD := string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	fmt.Println(encodePWD)

	fmt.Println("====模拟登录====")
	u1:=User{}
	u1.Password=encodePWD //模拟从数据库中读取到的 经过bcrypt.GenerateFromPassword处理的密码值
	loginPwd:="pwd" //用户登录时输入的密码
	// 密码验证
	err = bcrypt.CompareHashAndPassword([]byte(u1.Password), []byte(loginPwd)) //验证（对比）
	if err != nil {
		fmt.Println("pwd wrong")
	} else {
		fmt.Println("pwd ok")
	}
}
