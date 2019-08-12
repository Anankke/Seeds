package utils

import "testing"
// you should pass a environment variable named VERIFYKEY
func TestEnvConfig(t *testing.T) {
	c := GetConfig()
	verifyKey := c.Get("verifyKey");
	if  verifyKey=="Hello" {
		t.Error("can`t load env VERIFYKEY")
	}

}
