package handlers

import (
	cb "UlboraApiGateway/circuitbreaker"
	mgr "UlboraApiGateway/managers"
	"bytes"
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"

	//"time"
	//"io/ioutil"
	//"net/http"
	//"net/http/httptest"
	"testing"
	//"time"
)

var clustCidBk int64 = 79
var routeBk int64
var routeURLBkID int64
var connectedForBk bool
var edbBk cb.CircuitBreaker
var gwRoutesBk mgr.GatewayRoutes
var hBk Handler
var bkID int64

func TestBkn_ConnectFor(t *testing.T) {
	edbBk.DbConfig.Host = "localhost:3306"
	edbBk.DbConfig.DbUser = "admin"
	edbBk.DbConfig.DbPw = "admin"
	edbBk.DbConfig.DatabaseName = "ulbora_api_gateway"
	connectedForBk = edbBk.ConnectDb()
	if connectedForBk != true {
		t.Fail()
	}
	edbBk.CacheHost = getCacheHost()
	gwRoutesBk.GwDB.DbConfig = edbBk.DbConfig
	//gwRoutes.GwDB.DbConfig = gwRoutes.GwDB.DbConfig
	//cp.Host = "http://localhost:3010"
	testMode = true
	hBk.DbConfig = edbBk.DbConfig
}

func TestBkn2_InsertClient(t *testing.T) {
	var c mgr.Client
	c.APIKey = "12233hgdd333"
	c.ClientID = clustCidBk
	c.Enabled = true
	c.Level = "small"

	res := gwRoutesBk.GwDB.InsertClient(&c)
	if res.Success == true && res.ID != -1 {
		fmt.Print("new client Id: ")
		fmt.Println(clustCidBk)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestBkn2_InsertRestRoute(t *testing.T) {
	var rr mgr.RestRoute
	rr.Route = "content"
	rr.ClientID = clustCidBk

	res := gwRoutesBk.GwDB.InsertRestRoute(&rr)
	if res.Success == true && res.ID != -1 {
		routeBk = res.ID
		fmt.Print("new route Id: ")
		fmt.Println(routeErr)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestBkn2_InsertRouteURL(t *testing.T) {
	var ru mgr.RouteURL
	ru.Name = "blue"
	ru.URL = "http://www.apigateway.com/blue/"
	ru.Active = false
	ru.RouteID = routeBk
	ru.ClientID = clustCidBk

	res := gwRoutesBk.GwDB.InsertRouteURL(&ru)
	if res.Success == true && res.ID != -1 {
		routeURLBkID = res.ID
		fmt.Print("new route url Id: ")
		fmt.Println(res.ID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}

	var ru2 mgr.RouteURL
	ru2.Name = "sideb"
	ru2.URL = "http://www.apigateway.com/blue/"
	ru2.Active = false
	ru2.RouteID = routeBk
	ru2.ClientID = clustCidBk

	res2 := gwRoutesBk.GwDB.InsertRouteURL(&ru2)
	if res2.Success == true && res2.ID != -1 {
		//routeURLID33 = res2.ID
		fmt.Print("new route url Id: ")
		fmt.Println(res2.ID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestBkn2_HandleMedia(t *testing.T) {
	var c cb.Breaker
	//c.ClientID = clustCidBk
	c.FailoverRouteName = "test"
	c.FailureCount = 1
	c.FailureThreshold = 3
	c.HealthCheckTimeSeconds = 120
	c.RestRouteID = routeBk
	c.RouteURIID = routeURLBkID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "79")
	r.Header.Set("clientId", "79")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBk.HandleBreakerPost(w, r)
	fmt.Print("Media Code: ")
	fmt.Println(w.Code)
	if w.Code != http.StatusUnsupportedMediaType {
		t.Fail()
	}
}

func TestBkn2_HandleReq(t *testing.T) {
	var c cb.Breaker
	c.FailoverRouteName = "test"
	c.FailureCount = 1
	c.FailureThreshold = 3
	c.HealthCheckTimeSeconds = 120
	//c.RestRouteID = routeBk
	c.RouteURIID = routeURLBkID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "79")
	r.Header.Set("clientId", "79")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hBk.HandleBreakerPost(w, r)
	fmt.Print("Req Code: ")
	fmt.Println(w.Code)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestBkn2_HandleMethod(t *testing.T) {
	var c cb.Breaker
	c.FailoverRouteName = "test"
	c.FailureThreshold = 3
	c.HealthCheckTimeSeconds = 120
	c.RestRouteID = routeBk
	c.RouteURIID = routeURLBkID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("GET", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "79")
	r.Header.Set("clientId", "79")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hBk.HandleBreakerPost(w, r)
	fmt.Print("Status Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy BreakerResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestBkn2_Handle(t *testing.T) {
	var c cb.Breaker
	c.FailoverRouteName = "test"
	c.FailureThreshold = 3
	c.HealthCheckTimeSeconds = 120
	c.RestRouteID = routeBk
	c.RouteURIID = routeURLBkID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "79")
	r.Header.Set("clientId", "79")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hBk.HandleBreakerPost(w, r)
	fmt.Print("Status Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy BreakerResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.Success != true {
		t.Fail()
	}
}

func TestBkn2_HandleGet(t *testing.T) {
	var routeBkStr string = strconv.FormatInt(routeBk, 10)
	var routeURLBkIDStr string = strconv.FormatInt(routeURLBkID, 10)
	r, _ := http.NewRequest("GET", "/test?routeId="+routeBkStr+"&urlId="+routeURLBkIDStr, nil)
	r.Header.Set("u-client-id", "79")
	r.Header.Set("clientId", "79")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBk.HandleBreakerGet(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy cb.Breaker
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.ID == 0 {
		t.Fail()
	} else {
		bkID = bdy.ID
	}
}

func TestBkn2_HandleGetMethod(t *testing.T) {
	var routeBkStr string = strconv.FormatInt(routeBk, 10)
	var routeURLBkIDStr string = strconv.FormatInt(routeURLBkID, 10)
	r, _ := http.NewRequest("PUT", "/test?routeId="+routeBkStr+"&urlId="+routeURLBkIDStr, nil)
	r.Header.Set("u-client-id", "79")
	r.Header.Set("clientId", "79")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBk.HandleBreakerGet(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy cb.Breaker
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestBkn2_HandlePutMedia(t *testing.T) {
	var c cb.Breaker
	c.FailoverRouteName = "test2"
	c.FailureCount = 1
	c.FailureThreshold = 3
	c.HealthCheckTimeSeconds = 120
	c.RestRouteID = routeBk
	c.RouteURIID = routeURLBkID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "79")
	r.Header.Set("clientId", "79")
	r.Header.Set("u-api-key", "12233hgdd3335")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hBk.HandleBreakerPut(w, r)
	fmt.Print("Status Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy BreakerResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusUnsupportedMediaType {
		t.Fail()
	}
}

func TestBkn2_HandlePutReq(t *testing.T) {
	var c cb.Breaker
	c.FailoverRouteName = "test2"
	c.FailureCount = 1
	c.FailureThreshold = 3
	c.HealthCheckTimeSeconds = 120
	//c.RestRouteID = routeBk
	c.RouteURIID = routeURLBkID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "79")
	r.Header.Set("clientId", "79")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hBk.HandleBreakerPut(w, r)
	fmt.Print("Status Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy BreakerResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestBkn2_HandlePutMethod(t *testing.T) {
	var c cb.Breaker
	c.FailoverRouteName = "test2"
	c.FailureCount = 1
	c.FailureThreshold = 3
	c.HealthCheckTimeSeconds = 120
	c.RestRouteID = routeBk
	c.RouteURIID = routeURLBkID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("GET", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "79")
	r.Header.Set("clientId", "79")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hBk.HandleBreakerPut(w, r)
	fmt.Print("Status Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy BreakerResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestBkn2_HandlePut(t *testing.T) {
	var c cb.Breaker
	c.ID = bkID
	c.FailoverRouteName = "test2"
	c.FailureThreshold = 1
	c.HealthCheckTimeSeconds = 120
	c.RestRouteID = routeBk
	c.RouteURIID = routeURLBkID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "79")
	r.Header.Set("clientId", "79")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hBk.HandleBreakerPut(w, r)
	fmt.Print("Status Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy BreakerResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.Success != true {
		t.Fail()
	}
}

func TestBkn2_HandleGet2(t *testing.T) {
	var routeBkStr string = strconv.FormatInt(routeBk, 10)
	var routeURLBkIDStr string = strconv.FormatInt(routeURLBkID, 10)
	r, _ := http.NewRequest("GET", "/test?routeId="+routeBkStr+"&urlId="+routeURLBkIDStr, nil)
	r.Header.Set("u-client-id", "79")
	r.Header.Set("clientId", "79")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBk.HandleBreakerGet(w, r)
	fmt.Print("Media Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy cb.Breaker
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp get: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.ID == 0 || bdy.FailoverRouteName != "test2" {
		t.Fail()
	} else {
		bksID = bdy.ID
	}
}

func TestBkn2_CircuitBreakerTrip(t *testing.T) {
	bk := new(cb.Breaker)
	bk.ClientID = clustCidBk
	bk.RouteURIID = routeURLBkID
	edbBk.Trip(bk)
	edbBk.Trip(bk)
}

func TestBkn2_HandleStatus(t *testing.T) {
	//var routeBksStr string = strconv.FormatInt(routeBks, 10)
	var routeURLBkIDStr string = strconv.FormatInt(routeURLBkID, 10)
	r, _ := http.NewRequest("GET", "/test?urlId="+routeURLBkIDStr, nil)
	r.Header.Set("u-client-id", "79")
	r.Header.Set("clientId", "79")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBk.HandleBreakerStatus(w, r)
	fmt.Print("Status Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy cb.Status
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Status Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.Open == false {
		t.Fail()
	}
}

func TestBkn2_HandleStatusReq(t *testing.T) {
	//var routeBksStr string = strconv.FormatInt(routeBks, 10)
	//var routeURLBkIDStr string = strconv.FormatInt(routeURLBkID, 10)
	r, _ := http.NewRequest("GET", "/test?urlId=e", nil)
	r.Header.Set("u-client-id", "79")
	r.Header.Set("clientId", "79")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBk.HandleBreakerStatus(w, r)
	fmt.Print("Status Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy cb.Status
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Status Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestBkn2_HandleStatusMedia(t *testing.T) {
	//var routeBksStr string = strconv.FormatInt(routeBks, 10)
	var routeURLBkIDStr string = strconv.FormatInt(routeURLBkID, 10)
	r, _ := http.NewRequest("PUT", "/test?urlId="+routeURLBkIDStr, nil)
	r.Header.Set("u-client-id", "79")
	r.Header.Set("clientId", "79")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBk.HandleBreakerStatus(w, r)
	fmt.Print("Status Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy cb.Status
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Status Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestBkn2_HandleResetMedia(t *testing.T) {
	var c cb.Breaker
	c.RouteURIID = routeURLBkID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "79")
	r.Header.Set("clientId", "79")
	r.Header.Set("u-api-key", "12233hgdd3335")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hBk.HandleBreakerReset(w, r)
	fmt.Print("Status Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy BreakerResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusUnsupportedMediaType {
		t.Fail()
	}
}

func TestBkn2_HandleResetReq(t *testing.T) {
	var c cb.Breaker
	//c.RouteURIID = routeURLBkID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "79")
	r.Header.Set("clientId", "79")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hBk.HandleBreakerReset(w, r)
	fmt.Print("Status Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy BreakerResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestBkn2_HandleResetMethod(t *testing.T) {
	var c cb.Breaker
	c.RouteURIID = routeURLBkID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("GET", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "79")
	r.Header.Set("clientId", "79")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hBk.HandleBreakerReset(w, r)
	fmt.Print("Status Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy BreakerResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestBkn2_HandleReset(t *testing.T) {
	var c cb.Breaker
	c.RouteURIID = routeURLBkID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "79")
	r.Header.Set("clientId", "79")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hBk.HandleBreakerReset(w, r)
	fmt.Print("Status Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy BreakerResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.Success != true {
		t.Fail()
	}
}

func TestBkn2_HandleStatus2(t *testing.T) {
	//var routeBksStr string = strconv.FormatInt(routeBks, 10)
	var routeURLBkIDStr string = strconv.FormatInt(routeURLBkID, 10)
	r, _ := http.NewRequest("GET", "/test?urlId="+routeURLBkIDStr, nil)
	r.Header.Set("u-client-id", "79")
	r.Header.Set("clientId", "79")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBk.HandleBreakerStatus(w, r)
	fmt.Print("Status Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy cb.Status
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Status Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.Open != false {
		t.Fail()
	}
}

func TestBkn2_HandleDelReq(t *testing.T) {
	//var routeBkStr string = strconv.FormatInt(routeBk, 10)
	var routeURLBkIDStr string = strconv.FormatInt(routeURLBkID, 10)
	r, _ := http.NewRequest("DELETE", "/test?routeId=w&urlId="+routeURLBkIDStr, nil)
	r.Header.Set("u-client-id", "79")
	r.Header.Set("clientId", "79")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBk.HandleBreakerDelete(w, r)
	fmt.Print("Media Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy BreakerResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp delete: ")
	fmt.Println(bdy)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestBkn2_HandleDelMethod(t *testing.T) {
	var routeBkStr string = strconv.FormatInt(routeBk, 10)
	var routeURLBkIDStr string = strconv.FormatInt(routeURLBkID, 10)
	r, _ := http.NewRequest("GET", "/test?routeId="+routeBkStr+"&urlId="+routeURLBkIDStr, nil)
	r.Header.Set("u-client-id", "79")
	r.Header.Set("clientId", "79")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBk.HandleBreakerDelete(w, r)
	fmt.Print("Media Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy BreakerResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp delete: ")
	fmt.Println(bdy)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestBkn2_HandleDel(t *testing.T) {
	var routeBkStr string = strconv.FormatInt(routeBk, 10)
	var routeURLBkIDStr string = strconv.FormatInt(routeURLBkID, 10)
	r, _ := http.NewRequest("DELETE", "/test?routeId="+routeBkStr+"&urlId="+routeURLBkIDStr, nil)
	r.Header.Set("u-client-id", "79")
	r.Header.Set("clientId", "79")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBk.HandleBreakerDelete(w, r)
	fmt.Print("Media Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy BreakerResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp delete: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.Success != true {
		t.Fail()
	}
}

func TestBkn2_HandleGet3(t *testing.T) {
	var routeBkStr string = strconv.FormatInt(routeBk, 10)
	var routeURLBkIDStr string = strconv.FormatInt(routeURLBkID, 10)
	r, _ := http.NewRequest("GET", "/test?routeId="+routeBkStr+"&urlId="+routeURLBkIDStr, nil)
	r.Header.Set("u-client-id", "79")
	r.Header.Set("clientId", "79")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBk.HandleBreakerGet(w, r)
	fmt.Print("Media Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy cb.Breaker
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp get3: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.ID != 0 {
		t.Fail()
	}
}

func TestBkn2_DeleteClient(t *testing.T) {
	var c mgr.Client
	c.ClientID = clustCidBk
	res := gwRoutesBk.GwDB.DeleteClient(&c)
	if res.Success != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestBkn2_TestCloseDb(t *testing.T) {
	success := gwRoutesBk.GwDB.CloseDb()
	if success != true {
		t.Fail()
	}
}
