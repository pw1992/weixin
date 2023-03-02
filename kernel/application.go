package kernel

import (
	"github.com/pw1992/weixin"
	"github.com/pw1992/weixin/kernel/cache"
	"github.com/pw1992/weixin/kernel/config"
	"github.com/pw1992/weixin/kernel/contracts"
	"github.com/pw1992/weixin/kernel/serror"
	"github.com/pw1992/weixin/kernel/slog"
)

type Application struct {
	Config     *config.Config
	HttpClient *weixin.HttpClient
	Log        *slog.Log
	Error      *serror.Error
	Cache      contracts.Cacher
}

func NewApplication() *Application {
	app := &Application{
		Config:     config.NewConfig(),
		HttpClient: weixin.NewHttpClient(),
		Log:        slog.New(),
		Cache:      nil,
	}

	//设置换粗适配器
	cacheAdapter := app.Config.GetString("cacheAdapter")
	switch cacheAdapter {
	case "redis":
		app.Cache = cache.NewRedis()
	case "file":
		app.Cache = cache.NewFile()
	default:
		app.Cache = cache.NewFile()
	}

	return app
}
