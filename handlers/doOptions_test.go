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

// type challange struct {
// 	Answer string `json:"answer"`
// 	Key    string `json:"key"`
// }

var tgocid int64 = 46

//var rrID int64

//var routeErr int64
//var routeURLErrID int64
var connectedForTgo bool
var gwTgo mgr.GatewayDB

//var gwRoutesErr mgr.GatewayRoutes
//var hrr Handler

func TestGatewayOptions_Connect(t *testing.T) {
	gwTgo.DbConfig.Host = "localhost:3306"
	gwTgo.DbConfig.DbUser = "admin"
	gwTgo.DbConfig.DbPw = "admin"
	gwTgo.DbConfig.DatabaseName = "ulbora_api_gateway"
	connectedForTgo = gwTgo.ConnectDb()
	if connectedForTgo != true {
		t.Fail()
	}
	//gwRoutesErr.GwDB.DbConfig = edb.DbConfig
	//gwRoutes.GwDB.DbConfig = gwRoutes.GwDB.DbConfig
	//cp.Host = "http://localhost:3010"
	//testMode = true
	//hrr.DbConfig = gwRR.DbConfig
}

func TestGatewayOptions_doOption(t *testing.T) {
	var p passParams
	p.h = new(Handler)
	var cbr cb.CircuitBreaker
	cbr.DbConfig = gwTgo.DbConfig
	p.h.CbDB = cbr
	p.b = new(cb.Breaker)
	p.gwr = new(mgr.GatewayRoutes)
	p.rts = new(mgr.GatewayRouteURL)
	p.rts.URL = "http://challenge.myapigateway.com"
	p.fpath = "rs/challenge/en_us"
	var q = make(url.Values, 0)
	q.Set("p1", "param1")
	p.code = &q
	// var c challange
	// c.Answer = "test"
	// c.Key = "test"

	// aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("OPTIONS", "/test", nil)
	//r.Header.Set("Content-Type", "application/json")
	p.r = r
	w := httptest.NewRecorder()
	p.w = w

	//["p1"] = ["param1"]

	rtn := doOptions(&p)
	fmt.Print("doOptions Res: ")
	fmt.Println(rtn)
	if rtn.rtnCode != http.StatusOK {
		t.Fail()
	}
}
