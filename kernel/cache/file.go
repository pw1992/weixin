package cache

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/pw1992/weixin/kernel/serror"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type File struct {
	Path string
}

func NewFile(path string) *File {
	return &File{Path: path}
}

// 获取缓存
func (f *File) Get(key string, defaults ...interface{}) interface{} {
	path := os.TempDir()
	if len(f.Path) > 0 {
		path, _ = filepath.Abs(f.Path)
	}
	md5 := md5.New()
	md5.Write([]byte(key))
	str := hex.EncodeToString(md5.Sum(nil))
	file := str[:1]

	if _, err := os.Stat(filepath.Join(path, file)); os.IsNotExist(err) {
		if len(defaults) > 0 {
			return defaults[0]
		} else {
			return nil
		}
	}

	data, err := ioutil.ReadFile(filepath.Join(path, file))
	if err != nil {
		serror.NewError("读取创建缓存文件失败:"+err.Error(), 500, err)
		return nil
	}

	allDataMap := make(map[string]interface{})
	err = json.Unmarshal(data, &allDataMap)

	if err != nil {
		serror.NewError("解析缓存失败:"+err.Error(), 500, err)
		return nil
	}

	v, ok := allDataMap[key]
	if !ok {
		if len(defaults) > 0 {
			return defaults[0]
		} else {
			return nil
		}
	}

	vv, ok := v.(map[string]interface{})
	if !ok {
		serror.NewError("allDataMap断言失败", 500, v)
		return nil
	}

	expired_at, ok := vv["expired_at"].(float64)
	if !ok {
		serror.NewError("过期时间断言失败", 500, vv)
		return nil
	}

	if int64(expired_at) < time.Now().Unix() {
		delete(allDataMap, key)

		marshal, err := json.Marshal(allDataMap)
		if err != nil {
			serror.NewError("allDataMap解析json失败", 500, err)
			return nil
		}

		ioutil.WriteFile(filepath.Join(path, file), marshal, os.ModePerm)
		return nil
	}
	return vv["value"]
}

// 设置缓存
func (f *File) Set(key string, value interface{}, ttl int) bool {
	path := os.TempDir()
	if len(f.Path) > 0 {
		path, _ = filepath.Abs(f.Path)
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModeDir)
	}

	md5 := md5.New()
	md5.Write([]byte(key))
	str := hex.EncodeToString(md5.Sum(nil))
	file := str[:1]

	if _, err := os.Stat(filepath.Join(path, file)); os.IsNotExist(err) {
		os.Create(filepath.Join(path, file))
	}

	openFile, err := os.OpenFile(filepath.Join(path, file), os.O_RDWR, os.ModePerm)
	defer openFile.Close()
	if err != nil {
		serror.NewError("读取创建缓存文件失败:"+err.Error(), 500, err).Throw()
	}
	stat, _ := openFile.Stat()
	allData := make([]byte, stat.Size())
	allDataMap := make(map[string]interface{})
	openFile.Read(allData)

	json.Unmarshal(allData, &allDataMap)

	m := make(map[string]interface{}, 0)
	m["key"] = key
	m["value"] = value
	m["expired_at"] = int64(ttl) + time.Now().Unix()
	allDataMap[key] = m

	data, _ := json.Marshal(allDataMap)
	if len(data) > 0 {
		ioutil.WriteFile(filepath.Join(path, file), data, os.ModePerm)
	}
	return true
}

func (f *File) Delete(key string) bool {
	path := os.TempDir()
	if len(f.Path) > 0 {
		path, _ = filepath.Abs(f.Path)
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModeDir)
	}

	md5 := md5.New()
	md5.Write([]byte(key))
	str := hex.EncodeToString(md5.Sum(nil))
	file := str[:1]

	file = filepath.Join(path, file)

	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		serror.NewError("读取创建缓存文件失败:"+err.Error(), 500, err)
		return false
	}

	allDataMap := make(map[string]interface{})
	err = json.Unmarshal(data, &allDataMap)

	if err != nil {
		serror.NewError("解析缓存失败:"+err.Error(), 500, err)
		return false
	}
	delete(allDataMap, key)

	input, _ := json.Marshal(allDataMap)

	ioutil.WriteFile(file, input, os.ModePerm)
	return true
}

func (f *File) Clear() bool {
	//TODO implement me
	panic("implement me")
}
