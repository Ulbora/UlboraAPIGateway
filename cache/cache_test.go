package cache

import (
	"fmt"
	"testing"
)

func TestCProxy_Set(t *testing.T) {
	var cp CProxy
	cp.Host = "http://localhost:3010"
	var i Item
	i.Key = "test111222"
	i.Value = "ddddd"
	res := cp.Set(&i)
	fmt.Print("Resp: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}
