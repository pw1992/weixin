package auth

import (
	"testing"
)

func TestNewAccessToken(t *testing.T) {
	newToken := NewAccessToken()
	token := newToken.GetToken()
	if len(token) == 0 {
		t.Errorf("获取access_token 失败 %s", token)
	}
}
