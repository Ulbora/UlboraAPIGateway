package cache

import (
	"fmt"
	"testing"
)

func TestCProxy_SetMarshal(t *testing.T) {
	var cp CProxy
	cp.Host = "http://localhost:3010"
	//var i Item
	//i.Key = "test111222"
	//i.Value = "ddddd"
	res := cp.Set(nil)
	fmt.Print("Resp: ")
	fmt.Println(res)
	if res.Success == true {
		t.Fail()
	}
}

func TestCProxy_SetReq(t *testing.T) {
	var cp CProxy
	//cp.Host = "http://localhost:3010"
	var i Item
	i.Key = "test111222"
	i.Value = "ddddd"
	res := cp.Set(&i)
	fmt.Print("Resp: ")
	fmt.Println(res)
	if res.Success == true {
		t.Fail()
	}
}

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

func TestCProxy_GetUrl(t *testing.T) {
	var cp CProxy
	//cp.Host = "http://localhost:3010"
	var key = "test111222"

	res := cp.Get(key)
	fmt.Print("Resp: ")
	fmt.Println(res)
	if res.Success == true || res.Value == "ddddd" {
		t.Fail()
	}
}

func TestCProxy_Get(t *testing.T) {
	var cp CProxy
	cp.Host = "http://localhost:3010"
	var key = "test111222"

	res := cp.Get(key)
	fmt.Print("Resp: ")
	fmt.Println(res)
	if res.Success != true || res.Value != "ddddd" {
		t.Fail()
	}
}

func TestCProxy_DeleteUrl(t *testing.T) {
	var cp CProxy
	cp.Host = "2345678"
	var key = "test111222"
	res := cp.Delete(key)
	fmt.Print("Resp: ")
	fmt.Println(res)
	if res.Success == true {
		t.Fail()
	}
}

func TestCProxy_DeleteReq(t *testing.T) {
	var cp CProxy
	cp.Host = "http://localhost:3010"
	var key string // = "test111222"
	res := cp.Delete(key)
	fmt.Print("Resp: ")
	fmt.Println(res)
	if res.Success == true {
		t.Fail()
	}
}

func TestCProxy_Delete(t *testing.T) {
	var cp CProxy
	cp.Host = "http://localhost:3010"
	var key = "test111222"
	res := cp.Delete(key)
	fmt.Print("Resp: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestCProxy_Get2(t *testing.T) {
	var cp CProxy
	cp.Host = "http://localhost:3010"
	var key = "test111222"

	res := cp.Get(key)
	fmt.Print("Resp: ")
	fmt.Println(res)
	if res.Success != false || res.Value != "" {
		t.Fail()
	}
}
