package config

import (
	"github.com/pw1992/weixin/kernel/serror"
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

func NewConfig(viper2 *viper.Viper) *Config {
	c := &Config{viper2}
	return c
	c.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	c.SetConfigName("weixin") // name of config file (without extension)
	viper.AddConfigPath("$HOME/.weixin")
	err := viper.ReadInConfig()
	if err != nil {
		serror.NewError("get config fail:"+err.Error(), 500, err).Throw()
	}
	return c
}

func (c *Config) GetString(key string) string {
	m := make(map[string]string, 0)
	m["corpid"] = "ww440a797023c36add"
	m["corpsecret"] = "EOIFLOZMptopAISSIDx5QT2VYbYCyq9hXhiooiZiOhw"

	return m[key]
	return c.GetString(key)
}
func (c *Config) GetInt(key string) int {
	return c.GetInt(key)
}
func (c *Config) GetBool(key string) bool {
	return c.GetBool(key)
}
func (c *Config) GetStringMap(key string) map[string]interface{} {
	return c.GetStringMap(key)
}
