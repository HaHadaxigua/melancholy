package test

import (
	"fmt"
	"testing"
)

var (
	name string
)

func TestReadInput(t *testing.T) {
	fmt.Println("input name!")
	fmt.Scanln(&name)
	fmt.Println(name)
}
