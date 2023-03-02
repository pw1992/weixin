package cache

import (
	"github.com/gomodule/redigo/redis"
	"github.com/mna/redisc"
	"github.com/pw1992/weixin/kernel/config"
	"github.com/pw1992/weixin/kernel/serror"
	"time"
)

type Redis struct {
	Conn redis.Conn
}

func NewRedis() *Redis {
	newConfig := config.NewConfig()
	stringMap := newConfig.GetStringMap("redis")
	clusters := make([]string, 0)
	for _, node := range stringMap {
		node := node.(map[string]interface{})
		clusters = append(clusters, node["host"].(string)+":"+node["port"].(string))
	}

	cluster := redisc.Cluster{
		//StartupNodes: []string{"127.0.0.1:6381", "127.0.0.1:6382", "127.0.0.1:6383"},
		StartupNodes: clusters,
		DialOptions:  []redis.DialOption{redis.DialConnectTimeout(5 * time.Second)},
		CreatePool: func(addr string, opts ...redis.DialOption) (*redis.Pool, error) {
			return &redis.Pool{
				MaxIdle:     5,
				MaxActive:   10,
				IdleTimeout: time.Minute,
				Dial: func() (redis.Conn, error) {
					return redis.Dial("tcp", addr, opts...)
				},
				TestOnBorrow: func(c redis.Conn, t time.Time) error {
					_, err := c.Do("PING")
					return err
				},
			}, nil
		},
	}
	//defer cluster.Close()
	// initialize its mapping
	if err := cluster.Refresh(); err != nil {
		serror.NewError("Refresh failed:"+err.Error(), 500)
	}
	// grab a connection from the pool
	conn := cluster.Get()
	//defer conn.Close()
	return &Redis{
		Conn: conn,
	}
}

func (r *Redis) Get(key string, defaults ...interface{}) interface{} {
	s, err := redis.String(r.Conn.Do("GET", key))
	if err != nil {
		serror.NewError("获取失败:"+err.Error(), 500, key)
	}
	if len(s) == 0 && len(defaults) != 0 {
		return defaults[0]
	}
	return s
}

func (r *Redis) Set(key string, value interface{}, ttl int) bool {
	v, ok := value.([]uint8)
	if ok {
		value = string(v)
	}

	_, err := redis.String(r.Conn.Do("SET", key, value))
	if err != nil {
		serror.NewError("设置缓存失败:"+err.Error(), 500, key)
		return false
	}

	if ttl > 0 {
		_, err := redis.String(r.Conn.Do("EXPIRE", key, ttl))
		if err != nil {
			serror.NewError("设置缓存TTL失败:"+err.Error(), 500, key)
			return false
		}
	}
	return true
}

func (r *Redis) Delete(key string) bool {
	//TODO implement me
	panic("implement me")
}

func (r *Redis) Clear() bool {
	//TODO implement me
	panic("implement me")
}
