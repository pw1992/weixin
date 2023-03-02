package kernel

import "testing"

func TestGetRootPath(t *testing.T) {
	path := GetRootPath(".log")
	t.Log(path)
}
