package kernel

import (
	"github.com/pw1992/weixin"
	"github.com/pw1992/weixin/kernel/cache"
	"github.com/pw1992/weixin/kernel/config"
	"github.com/pw1992/weixin/kernel/contracts"
	"github.com/pw1992/weixin/kernel/serror"
	"github.com/pw1992/weixin/kernel/slog"
	"github.com/spf13/viper"
	"os"
)

type Application struct {
	Config     *config.Config
	HttpClient *weixin.HttpClient
	Log        *slog.Log
	Error      *serror.Error
	Cache      contracts.Cacher
}

func NewApplication() *Application {
	return &Application{
		Config:     config.NewConfig(viper.New()),
		HttpClient: weixin.NewHttpClient(),
		Log:        slog.New(),
		Cache:      cache.NewFile(os.TempDir()),
	}
}
