package handlers

import (
	"io/ioutil"
	"strconv"
	//cb "UlboraApiGateway/circuitbreaker"
	mgr "UlboraApiGateway/managers"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var tghcid int64 = 46
var tghRR int64

//var rrID int64

//var routeErr int64
//var routeURLErrID int64
var connectedForTgh bool
var gwTgh mgr.GatewayDB
var gwTH Handler

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
	testMode = true
	gwTH.DbConfig = gwTgh.DbConfig
	//hrr.DbConfig = gwRR.DbConfig
}

func TestGwHandler_HandleGwRoutePostUrl(t *testing.T) {
	var c challenge
	c.Answer = "test"
	c.Key = "test"

	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	gwTH.HandleGwRoute(w, r)

	fmt.Print("Code for no route: ")
	fmt.Println(w.Code)
	if w.Code != 0 {
		t.Fail()
	}
}

func TestGwHandler_HandleClientInsert(t *testing.T) {
	var c mgr.Client
	c.ClientID = tghcid
	c.APIKey = "123456"
	c.Enabled = true
	c.Level = "small"
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	gwTH.HandleClientPost(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.Success != true {
		t.Fail()
	}
}

func TestGwHandler_HandleRestRouteInsert(t *testing.T) {
	var rr mgr.RestRoute
	rr.ClientID = tghcid
	rr.Route = "test"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "46")
	r.Header.Set("u-api-key", "123456")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	gwTH.HandleRestRouteSuperPost(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.Success != true {
		t.Fail()
	} else {
		tghRR = bdy.ID
	}
}

func TestGwHandler_HandleRouteUrlInsert(t *testing.T) {
	var rr mgr.RouteURL
	rr.ClientID = tghcid
	rr.Name = "blue"
	rr.RouteID = tghRR
	rr.URL = "test/url"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "46")
	r.Header.Set("u-api-key", "123456")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	gwTH.HandleRouteURLSuperPost(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.Success != true {
		t.Fail()
	} else {
		rruuID = bdy.ID
	}
}

func TestGwHandler_HandleGwRoutePost(t *testing.T) {
	var c challenge
	c.Answer = "test"
	c.Key = "test"

	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/np?route=test&rname=blue", bytes.NewBuffer(aJSON))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("u-client-id", "46")
	r.Header.Set("u-api-key", "123456")

	w := httptest.NewRecorder()

	gwTH.HandleGwRoute(w, r)

	fmt.Print("Code for post: ")
	fmt.Println(w.Code)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestGwHandler_HandleGwRoutePut(t *testing.T) {
	var c challenge
	c.Answer = "test"
	c.Key = "test"

	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("PUT", "/np?route=test&rname=blue", bytes.NewBuffer(aJSON))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("u-client-id", "46")
	r.Header.Set("u-api-key", "123456")

	w := httptest.NewRecorder()

	gwTH.HandleGwRoute(w, r)

	fmt.Print("Code for put: ")
	fmt.Println(w.Code)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestGwHandler_HandleGwRoutePatch(t *testing.T) {
	var c challenge
	c.Answer = "test"
	c.Key = "test"

	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("PATCH", "/np?route=test&rname=blue", bytes.NewBuffer(aJSON))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("u-client-id", "46")
	r.Header.Set("u-api-key", "123456")

	w := httptest.NewRecorder()

	gwTH.HandleGwRoute(w, r)

	fmt.Print("Code patch: ")
	fmt.Println(w.Code)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestGwHandler_HandleGwRouteGet(t *testing.T) {
	var c challenge
	c.Answer = "test"
	c.Key = "test"

	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("GET", "/np?route=test&rname=blue", bytes.NewBuffer(aJSON))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("u-client-id", "46")
	r.Header.Set("u-api-key", "123456")

	w := httptest.NewRecorder()

	gwTH.HandleGwRoute(w, r)

	fmt.Print("Code get: ")
	fmt.Println(w.Code)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestGwHandler_HandleGwRouteDel(t *testing.T) {
	var c challenge
	c.Answer = "test"
	c.Key = "test"

	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("DELETE", "/np?route=test&rname=blue", bytes.NewBuffer(aJSON))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("u-client-id", "46")
	r.Header.Set("u-api-key", "123456")

	w := httptest.NewRecorder()

	gwTH.HandleGwRoute(w, r)

	fmt.Print("Code for delete: ")
	fmt.Println(w.Code)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestGwHandler_HandleGwRouteOpt(t *testing.T) {
	var c challenge
	c.Answer = "test"
	c.Key = "test"

	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("OPTIONS", "/np?route=test&rname=blue", bytes.NewBuffer(aJSON))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("u-client-id", "46")
	r.Header.Set("u-api-key", "123456")

	w := httptest.NewRecorder()

	gwTH.HandleGwRoute(w, r)

	fmt.Print("Code options: ")
	fmt.Println(w.Code)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestGwHandler_HandleDel(t *testing.T) {
	//var routeBkStr string = strconv.FormatInt(routeBk, 10)
	var CIDStr string = strconv.FormatInt(tghcid, 10)
	r, _ := http.NewRequest("DELETE", "/test?clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "24")
	r.Header.Set("clientId", "24")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	gwTH.HandleClientDelete(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy BreakerResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp delete: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK {
		t.Fail()
	}
}
