package main

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/internal/basic/tools"
)

// 用于生成持久的token

func main() {
	token, err := tools.JwtGenerateToken(2, "admin@admin.com", "123456", 30*24)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(token)

	claims, err := tools.JwtParseToken(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(claims.Id)
}
