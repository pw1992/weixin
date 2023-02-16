package contracts

type Cacher interface {
	//获取缓存
	Get(key string, defaults ...interface{}) interface{}
	//设置缓存
	Set(key string, value interface{}, ttl int) bool
	//清除缓存
	Delete(key string) bool
	//清除缓存
	Clear() bool
}
