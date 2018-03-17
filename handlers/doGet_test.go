package handlers

import (
	cb "UlboraApiGateway/circuitbreaker"
	mgr "UlboraApiGateway/managers"
	//"bytes"
	//"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var tggcid int64 = 46

//var rrID int64

//var routeErr int64
//var routeURLErrID int64
var connectedForTgg bool
var gwTgg mgr.GatewayDB

//var gwRoutesErr mgr.GatewayRoutes
//var hrr Handler

func TestGatewayGet_Connect(t *testing.T) {
	gwTgg.DbConfig.Host = "localhost:3306"
	gwTgg.DbConfig.DbUser = "admin"
	gwTgg.DbConfig.DbPw = "admin"
	gwTgg.DbConfig.DatabaseName = "ulbora_api_gateway"
	connectedForTgg = gwTgg.ConnectDb()
	if connectedForTgg != true {
		t.Fail()
	}
	//gwRoutesErr.GwDB.DbConfig = edb.DbConfig
	//gwRoutes.GwDB.DbConfig = gwRoutes.GwDB.DbConfig
	//cp.Host = "http://localhost:3010"
	//testMode = true
	//hrr.DbConfig = gwRR.DbConfig
}

func TestGatewayGet_doGetNotFound(t *testing.T) {
	var p passParams
	p.h = new(Handler)
	var cbr cb.CircuitBreaker
	cbr.DbConfig = gwTgp.DbConfig
	p.h.CbDB = cbr
	p.b = new(cb.Breaker)
	p.gwr = new(mgr.GatewayRoutes)
	p.rts = new(mgr.GatewayRouteURL)
	p.rts.URL = "http://challenge.myapigateway.com"
	p.fpath = "rs/challenge"
	var q = make(url.Values, 0)
	q.Set("p1", "param1")
	p.code = &q

	// aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("GET", "/test", nil)
	//r.Header.Set("Content-Type", "application/json")
	p.r = r
	w := httptest.NewRecorder()
	p.w = w

	//["p1"] = ["param1"]

	rtn := doGet(&p)
	fmt.Print("doGet Res: ")
	fmt.Println(rtn)
	if rtn.rtnCode != http.StatusNotFound {
		t.Fail()
	}
}

func TestGatewayGet_doGetMethod(t *testing.T) {
	var p passParams
	p.h = new(Handler)
	var cbr cb.CircuitBreaker
	cbr.DbConfig = gwTgp.DbConfig
	p.h.CbDB = cbr
	p.b = new(cb.Breaker)
	p.gwr = new(mgr.GatewayRoutes)
	p.rts = new(mgr.GatewayRouteURL)
	p.rts.URL = "http://challenge.myapigateway.com"
	p.fpath = "rs/challenge/en_us"
	var q = make(url.Values, 0)
	q.Set("p1", "param1")
	p.code = &q

	// aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("DELETE", "/test", nil)
	//r.Header.Set("Content-Type", "application/json")
	p.r = r
	w := httptest.NewRecorder()
	p.w = w

	//["p1"] = ["param1"]

	rtn := doGet(&p)
	fmt.Print("doGet Res: ")
	fmt.Println(rtn)
	if rtn.rtnCode != http.StatusNotFound {
		t.Fail()
	}
}

func TestGatewayGet_doGetReq(t *testing.T) {
	var p passParams
	p.h = new(Handler)
	var cbr cb.CircuitBreaker
	cbr.DbConfig = gwTgp.DbConfig
	p.h.CbDB = cbr
	p.b = new(cb.Breaker)
	p.gwr = new(mgr.GatewayRoutes)
	p.rts = new(mgr.GatewayRouteURL)
	//p.rts.URL = "http://challenge.myapigateway.com"
	p.fpath = "rs/challenge/en_us"
	var q = make(url.Values, 0)
	q.Set("p1", "param1")
	p.code = &q

	// aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("GET", "/test", nil)
	//r.Header.Set("Content-Type", "application/json")
	p.r = r
	w := httptest.NewRecorder()
	p.w = w

	//["p1"] = ["param1"]

	rtn := doGet(&p)
	fmt.Print("doGet Res: ")
	fmt.Println(rtn)
	if rtn.rtnCode != http.StatusBadRequest {
		t.Fail()
	}
}

func TestGatewayGet_doGet(t *testing.T) {
	var p passParams
	p.h = new(Handler)
	var cbr cb.CircuitBreaker
	cbr.DbConfig = gwTgp.DbConfig
	p.h.CbDB = cbr
	p.b = new(cb.Breaker)
	p.gwr = new(mgr.GatewayRoutes)
	p.rts = new(mgr.GatewayRouteURL)
	p.rts.URL = "http://challenge.myapigateway.com"
	p.fpath = "rs/challenge/en_us"
	var q = make(url.Values, 0)
	q.Set("p1", "param1")
	p.code = &q

	// aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("GET", "/test", nil)
	//r.Header.Set("Content-Type", "application/json")
	p.r = r
	w := httptest.NewRecorder()
	p.w = w

	//["p1"] = ["param1"]

	rtn := doGet(&p)
	fmt.Print("doGet Res: ")
	fmt.Println(rtn)
	if rtn.rtnCode != http.StatusOK {
		t.Fail()
	}
}
