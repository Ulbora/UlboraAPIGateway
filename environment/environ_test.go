package environment

import (
	"fmt"
	"testing"
)

func TestGetCacheHost(t *testing.T) {
	res := GetCacheHost()
	fmt.Print("res:")
	fmt.Println(res)
	if res != "http://localhost:3010" {
		t.Fail()
	}
}
