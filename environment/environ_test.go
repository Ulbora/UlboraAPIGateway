package environment

import (
	"fmt"
	"os"
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

func TestGetCacheHostEnvVar(t *testing.T) {
	os.Setenv("CACHE_HOST", "123")
	res := GetCacheHost()
	fmt.Print("res:")
	fmt.Println(res)
	if res == "http://localhost:3010" {
		t.Fail()
	}
}
