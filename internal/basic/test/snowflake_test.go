package test

import (
	"fmt"
	"github.com/HaHadaxigua/melancholy/internal/basic/tools"
	"testing"
)

func TestSnowflakeId(t *testing.T) {
	id, err := tools.tools.SnowflakeId()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(id)
}
