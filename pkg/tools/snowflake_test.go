package tools

import (
	"fmt"
	"testing"
)

func TestSnowflakeId(t *testing.T) {
	id, err := SnowflakeId()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(id)
}
