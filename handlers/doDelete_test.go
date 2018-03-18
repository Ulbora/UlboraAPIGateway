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

var tgdcid int64 = 46

//var rrID int64

//var routeErr int64
//var routeURLErrID int64
var connectedForTgd bool
var gwTgd mgr.GatewayDB

//var gwRoutesErr mgr.GatewayRoutes
//var hrr Handler

func TestGatewayDel_Connect(t *testing.T) {
	gwTgd.DbConfig.Host = "localhost:3306"
	gwTgd.DbConfig.DbUser = "admin"
	gwTgd.DbConfig.DbPw = "admin"
	gwTgd.DbConfig.DatabaseName = "ulbora_api_gateway"
	connectedForTgd = gwTgd.ConnectDb()
	if connectedForTgd != true {
		t.Fail()
	}
	//gwRoutesErr.GwDB.DbConfig = edb.DbConfig
	//gwRoutes.GwDB.DbConfig = gwRoutes.GwDB.DbConfig
	//cp.Host = "http://localhost:3010"
	//testMode = true
	//hrr.DbConfig = gwRR.DbConfig
}

func TestGatewayDel_doDelMethod(t *testing.T) {
	var p passParams
	p.h = new(Handler)
	var cbr cb.CircuitBreaker
	cbr.DbConfig = gwTgd.DbConfig
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
	r, _ := http.NewRequest("PUT", "/test", nil)
	//r.Header.Set("Content-Type", "application/json")
	p.r = r
	w := httptest.NewRecorder()
	p.w = w

	//["p1"] = ["param1"]

	rtn := doDelete(&p)
	fmt.Print("doGet Res: ")
	fmt.Println(rtn)
	if rtn.rtnCode != http.StatusNotFound {
		t.Fail()
	}
}

func TestGatewayDel_doDelReq(t *testing.T) {
	var p passParams
	p.h = new(Handler)
	var cbr cb.CircuitBreaker
	cbr.DbConfig = gwTgd.DbConfig
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
	r, _ := http.NewRequest("DELETE", "/test", nil)
	//r.Header.Set("Content-Type", "application/json")
	p.r = r
	w := httptest.NewRecorder()
	p.w = w

	//["p1"] = ["param1"]

	rtn := doDelete(&p)
	fmt.Print("doDel bad req Res: ")
	fmt.Println(rtn)
	if rtn.rtnCode != http.StatusBadRequest {
		t.Fail()
	}
}

func TestGatewayDel_doDelReq2(t *testing.T) {
	var p passParams
	p.h = new(Handler)
	var cbr cb.CircuitBreaker
	cbr.DbConfig = gwTgd.DbConfig
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
	r, _ := http.NewRequest("fff", "", nil)
	//r.Header.Set("Content-Type", "application/json")
	p.r = r
	w := httptest.NewRecorder()
	p.w = w

	//["p1"] = ["param1"]

	rtn := doDelete(&p)
	fmt.Print("doDel bad req Res2: ")
	fmt.Println(rtn)
	if rtn.rtnCode != http.StatusBadRequest {
		t.Fail()
	}
}
func TestGatewayDel_doDel(t *testing.T) {
	var p passParams
	p.h = new(Handler)
	var cbr cb.CircuitBreaker
	cbr.DbConfig = gwTgd.DbConfig
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

	rtn := doDelete(&p)
	fmt.Print("doGet Res: ")
	fmt.Println(rtn)
	if rtn.rtnCode != http.StatusNotFound {
		t.Fail()
	}
}
