package tools

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token, err := JwtGenerateToken(2, "123", "123", 2)
	assert.Nil(t, err)
	fmt.Println(token)
}

func TestJwtParseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IjEyMyIsInBhc3N3b3JkIjoiMTIzIiwiZXhwIjoxNjA5NTI4MjY2LCJqdGkiOiIyIiwiaWF0IjoxNjA5NTIxMDY2fQ.p04UPRfIXGEJWW3GwZNftq_uEDV4DCjWjGwgqPVupEk"
	parseToken, err := JwtParseToken(token)
	assert.Nil(t, err)
	fmt.Print(parseToken)
}