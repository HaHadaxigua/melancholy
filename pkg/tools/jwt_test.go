package tools

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJwtToken(t *testing.T){
	token, err := JwtGenerateToken(2, "123", "123", 2)
	assert.Nil(t, err)
	fmt.Println(token)

	parseToken, err := JwtParseToken(token)
	assert.Nil(t, err)
	fmt.Print(parseToken)
}
