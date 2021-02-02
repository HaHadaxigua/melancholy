package test

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/internal/basic/tools"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJwtToken(t *testing.T){
	token, err := tools.tools.JwtGenerateToken(2, "123", "123", 2)
	assert.Nil(t, err)
	fmt.Println(token)

	parseToken, err := tools.JwtParseToken(token)
	assert.Nil(t, err)
	fmt.Print(parseToken)
}
