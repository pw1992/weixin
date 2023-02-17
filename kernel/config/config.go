package config

import (
	"fmt"
	"github.com/pw1992/weixin/kernel/serror"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	*viper.Viper
}

func NewConfig(args ...string) *Config {
	abspath, _ := filepath.Abs("%$HOME%")
	index := strings.Index(abspath, "weixin")

	fmt.Println("%$HOME%", abspath[:index+6])
	c := &Config{viper.New()}
	c.Viper.SetConfigType("yaml")
	c.Viper.SetConfigName("weixin")
	if len(args) > 0 {
		c.Viper.AddConfigPath(args[0])
	} else {
		abspath, _ := filepath.Abs("")
		index := strings.Index(abspath, "weixin")
		path := abspath[:index+6] + "/.config"

		if _, err := os.Stat(path); os.IsNotExist(err) {
			serror.NewError("配置文件不存在：", 500, err)
		}
		c.Viper.AddConfigPath(path)
	}
	err := c.Viper.ReadInConfig()
	if err != nil {
		serror.NewError("配置文件初始化失败:"+err.Error(), 500, err).Throw()
	}
	return c
}

func (c *Config) GetString(key string) string {
	return c.Viper.GetString(key)
}
func (c *Config) GetInt(key string) int {
	return c.Viper.GetInt(key)
}
func (c *Config) GetBool(key string) bool {
	return c.Viper.GetBool(key)
}
func (c *Config) GetStringMap(key string) map[string]interface{} {
	return c.Viper.GetStringMap(key)
}
