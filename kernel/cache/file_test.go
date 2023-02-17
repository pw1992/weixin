package cache

import (
	"testing"
	"time"
)

func TestFile_Set(t *testing.T) {
	file := NewFile()
	ok := file.Set("name1", "pw", time.Now().Second())
	if !ok {
		t.Error("fail")
	}
}

func TestFile_Get(t *testing.T) {
	file := NewFile()
	file.Set("name12", "张三", time.Now().Second())

	val := file.Get("name12", 111)

	if _, ok := val.(string); !ok {
		t.Errorf("断言失败: 期望值string 原值%T", val)
	}

	//map
	testV1 := map[string]string{
		"f1": "v1",
		"f2": "v2",
	}
	file.Set("test1", testV1, 11)
	cache1 := file.Get("test1")

	if _, ok := cache1.(map[string]interface{}); !ok {
		t.Errorf("断言失败: 期望值%T 原值%T", testV1, cache1)
	}

	//切片
	testV12 := []string{"1", "2", "3", "v1", "v2"}
	file.Set("test2", testV12, 11)
	cache2 := file.Get("test2")
	if _, ok := cache2.([]interface{}); !ok {
		t.Errorf("断言失败: 期望值%T 原值%T", testV12, cache2)
	}
}

func TestCache_Delete(t *testing.T) {
	cache := NewFile()
	cache.Set("test-delete", "test-delete-val", 100)

	cache.Delete("test-delete")

	get := cache.Get("test-delete")
	if get != nil {
		t.Errorf("失败  预期:%v  实际:%v", nil, get)
	}
}
