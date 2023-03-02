package kernel

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

func DD(args ...interface{}) {
	format := ""
	for i := 0; i < len(args); i++ {
		format = format + "----\nType: %s\n%v\n----\n\n"
	}
	var inputs []interface{}
	for _, i := range args {
		if i == nil {
			inputs = append(inputs, "nil", "")
			continue
		}
		val := reflect.ValueOf(i)
		switch val.Kind() {
		case reflect.Map, reflect.Array, reflect.Interface, reflect.Slice:
			inputs = append(inputs, val.Type(), JsonEncode(i))
		default:
			inputs = append(inputs, val.Type(), i)
		}
	}
	fmt.Printf(format, inputs...)
	os.Exit(0)
}

func JsonEncode(i interface{}) (jsonstr string) {
	jsonbytes, err := json.Marshal(i)
	if err != nil {
		return "{}"
	}
	return string(jsonbytes)
}

func GetRootPath(activePath string) string {
	abs, err := filepath.Abs(activePath)
	if err != nil {
		panic(err)
	}
	index := strings.Index(abs, "weixin")
	rootpath := abs[0 : index+6]
	return filepath.Join(rootpath, activePath)
}
