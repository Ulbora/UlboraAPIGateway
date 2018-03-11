package handlers

import (
	//"net/http/httptest"
	//"bytes"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestHandler_parseQueryString(t *testing.T) {
	var q = make(url.Values, 0)
	q.Set("p1", "param1")
	q.Set("p2", "param2")
	rtn := parseQueryString(q)
	fmt.Print("querystring: ")
	fmt.Println(rtn)
	if rtn != "?p1=param1&p2=param2" {
		t.Fail()
	}
}

func TestHandler_getCacheHost(t *testing.T) {
	rtn := getCacheHost()
	if rtn != "http://localhost:3010" {
		t.Fail()
	}
}

func TestHandler_buildHeaders(t *testing.T) {
	pr, _ := http.NewRequest("POST", "/test", nil)
	pr.Header.Set("Content-Type", "application/json")
	sr, _ := http.NewRequest("POST", "/test", nil)
	buildHeaders(pr, sr)
	h := sr.Header
	var key string
	var val string
	for hn, v := range h {
		key = hn
		val = v[0]
	}
	fmt.Print("key: ")
	fmt.Println(key)
	fmt.Print("val: ")
	fmt.Println(val)
	if key != "Content-Type" || val != "application/json" {
		t.Fail()
	}
}

func TestHandler_buildRespHeaders(t *testing.T) {
	pw := new(http.Response)
	sw := new(http.ResponseWriter)
	fmt.Print("pw: ")
	fmt.Println(pw)
	fmt.Print("sw: ")
	fmt.Println(sw)
	buildRespHeaders(pw, *sw)
	if pw == nil || sw == nil {
		t.Fail()
	}
}

func TestHandler_getAuth(t *testing.T) {
	r, _ := http.NewRequest("POST", "/test", nil)
	res := getAuth(r)
	if res.ValidationURL != "http://localhost:3000/rs/token/validate" {
		t.Fail()
	}
}

func TestHandler_getHeaders(t *testing.T) {
	r, _ := http.NewRequest("POST", "/test", nil)
	r.Header.Set("clientId", "12345")
	res := getHeaders(r)
	fmt.Print("res: ")
	fmt.Println(res.clientID)
	if res.clientID != 12345 {
		t.Fail()
	}
}

func TestHandler_paramsOKs(t *testing.T) {
	p := new(passParams)
	rtn := paramsOK(p)
	if rtn != false {
		t.Fail()
	}
}
