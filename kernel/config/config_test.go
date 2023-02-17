package config

import (
	"fmt"
	"testing"
)

func TestConfig_GetString(t *testing.T) {
	config := NewConfig()
	getString := config.GetString("corpid")
	fmt.Println(getString)
}
