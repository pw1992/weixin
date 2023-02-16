package config

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestConfig_GetString(t *testing.T) {
	config := NewConfig(viper.New())
	getString := config.GetString("corpid")
	fmt.Println(getString)
}
