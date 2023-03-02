package cache

import (
	"testing"
)

func TestName(t *testing.T) {
	newRedis := NewRedis()
	newRedis.Set("name1", "zhangsan1", 100)
	get := newRedis.Get("name1")
	t.Logf("获取的redis值:%s", get)

}
