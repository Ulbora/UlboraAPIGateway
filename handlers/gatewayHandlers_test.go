package handlers

import (
	cb "UlboraApiGateway/circuitbreaker"
	mgr "UlboraApiGateway/managers"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var tghcid int64 = 46

//var rrID int64

//var routeErr int64
//var routeURLErrID int64
var connectedForTgh bool
var gwTgh mgr.GatewayDB

func TestGwHandler_Connect(t *testing.T) {
	gwTgh.DbConfig.Host = "localhost:3306"
	gwTgh.DbConfig.DbUser = "admin"
	gwTgh.DbConfig.DbPw = "admin"
	gwTgh.DbConfig.DatabaseName = "ulbora_api_gateway"
	connectedForTgh = gwTgh.ConnectDb()
	if connectedForTgh != true {
		t.Fail()
	}
	//gwRoutesErr.GwDB.DbConfig = edb.DbConfig
	//gwRoutes.GwDB.DbConfig = gwRoutes.GwDB.DbConfig
	//cp.Host = "http://localhost:3010"
	//testMode = true
	//hrr.DbConfig = gwRR.DbConfig
}

func TestGwHandler_HandleGwRoutePost(t *testing.T) {
	//var p passParams
	h := new(Handler)
	var cbr cb.CircuitBreaker
	cbr.DbConfig = gwTgp.DbConfig
	h.CbDB = cbr
	// b := new(cb.Breaker)
	// gwr := new(mgr.GatewayRoutes)
	// rts := new(mgr.GatewayRouteURL)
	// rts.URL = "http://challenge.myapigateway.com"
	// fpath := "rs/challenge"
	// var q = make(url.Values, 0)
	// q.Set("p1", "param1")
	// p.code = &q
	var c challange
	c.Answer = "test"
	c.Key = "test"

	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("Content-Type", "application/json")
	//p.r = r
	w := httptest.NewRecorder()

	h.HandleGwRoute(w, r)
	//p.w = w

	//["p1"] = ["param1"]

	//rtn := doPostPutPatch(&p)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	if w.Code != 0 {
		t.Fail()
	}
}
