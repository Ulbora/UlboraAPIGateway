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

var clustCidBks int64 = 69
var routeBks int64
var routeURLBksID int64
var connectedForBks bool
var edbBks cb.CircuitBreaker
var gwRoutesBks mgr.GatewayRoutes
var hBks Handler
var bksID int64

func TestBks_ConnectFor(t *testing.T) {
	edbBks.DbConfig.Host = "localhost:3306"
	edbBks.DbConfig.DbUser = "admin"
	edbBks.DbConfig.DbPw = "admin"
	edbBks.DbConfig.DatabaseName = "ulbora_api_gateway"
	connectedForBks = edbBks.ConnectDb()
	if connectedForBks != true {
		t.Fail()
	}
	edbBks.CacheHost = getCacheHost()
	fmt.Print("cache host: ")
	fmt.Println(edbBks.CacheHost)
	gwRoutesBks.GwDB.DbConfig = edbBks.DbConfig
	//gwRoutes.GwDB.DbConfig = gwRoutes.GwDB.DbConfig
	//cp.Host = "http://localhost:3010"
	testMode = true
	hBks.DbConfig = edbBks.DbConfig
}

// func TestErr_SetManager(t *testing.T) {
// 	SetManager(edb)
// }

func TestBks_InsertClient(t *testing.T) {
	var c mgr.Client
	c.APIKey = "12233hgdd333"
	c.ClientID = clustCidBks
	c.Enabled = true
	c.Level = "small"

	res := gwRoutesBks.GwDB.InsertClient(&c)
	if res.Success == true && res.ID != -1 {
		fmt.Print("new client Id: ")
		fmt.Println(clustCidBks)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestBks_InsertRestRoute(t *testing.T) {
	var rr mgr.RestRoute
	rr.Route = "content"
	rr.ClientID = clustCidBks

	res := gwRoutesBks.GwDB.InsertRestRoute(&rr)
	if res.Success == true && res.ID != -1 {
		routeBks = res.ID
		fmt.Print("new route Id: ")
		fmt.Println(routeErr)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestBks_InsertRouteURL(t *testing.T) {
	var ru mgr.RouteURL
	ru.Name = "blue"
	ru.URL = "http://www.apigateway.com/blue/"
	ru.Active = false
	ru.RouteID = routeBks
	ru.ClientID = clustCidBks

	res := gwRoutesBks.GwDB.InsertRouteURL(&ru)
	if res.Success == true && res.ID != -1 {
		routeURLBksID = res.ID
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
	ru2.RouteID = routeBks
	ru2.ClientID = clustCidBks

	res2 := gwRoutesBks.GwDB.InsertRouteURL(&ru2)
	if res2.Success == true && res2.ID != -1 {
		//routeURLID33 = res2.ID
		fmt.Print("new route url Id: ")
		fmt.Println(res2.ID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestBks_HandleSuperMedia(t *testing.T) {
	var c cb.Breaker
	c.ClientID = clustCidPer
	c.FailoverRouteName = "test"
	c.FailureCount = 1
	c.FailureThreshold = 3
	c.HealthCheckTimeSeconds = 120
	c.RestRouteID = routeBks
	c.RouteURIID = routeURLBksID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperPost(w, r)
	fmt.Print("Media Code: ")
	fmt.Println(w.Code)
	if w.Code != http.StatusUnsupportedMediaType {
		t.Fail()
	}
}

func TestBks_HandleSuperReq(t *testing.T) {
	var c cb.Breaker
	//c.ClientID = clustCidPer
	c.FailoverRouteName = "test"
	c.FailureCount = 1
	c.FailureThreshold = 3
	c.HealthCheckTimeSeconds = 120
	c.RestRouteID = routeBks
	c.RouteURIID = routeURLBksID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperPost(w, r)
	fmt.Print("Req Code: ")
	fmt.Println(w.Code)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestBks_HandleSuperMethod(t *testing.T) {
	var c cb.Breaker
	c.ClientID = clustCidPer
	c.FailoverRouteName = "test"
	//c.FailureCount = 1
	c.FailureThreshold = 3
	c.HealthCheckTimeSeconds = 120
	c.RestRouteID = routeBks
	c.RouteURIID = routeURLBksID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperPost(w, r)
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

func TestBks_HandleSuperAuth(t *testing.T) {
	testMode = false
	var c cb.Breaker
	c.ClientID = clustCidPer
	c.FailoverRouteName = "test"
	//c.FailureCount = 1
	c.FailureThreshold = 3
	c.HealthCheckTimeSeconds = 120
	c.RestRouteID = routeBks
	c.RouteURIID = routeURLBksID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperPost(w, r)
	fmt.Print("Status Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy BreakerResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusUnauthorized {
		t.Fail()
	}
	testMode = true
}

func TestBks_HandleSuper(t *testing.T) {
	var c cb.Breaker
	c.ClientID = clustCidPer
	c.FailoverRouteName = "test"
	//c.FailureCount = 1
	c.FailureThreshold = 3
	c.HealthCheckTimeSeconds = 120
	c.RestRouteID = routeBks
	c.RouteURIID = routeURLBksID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperPost(w, r)
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

func TestBks_HandleSuperGetReq(t *testing.T) {
	//var routeBksStr string = strconv.FormatInt(routeBks, 10)
	var routeURLBksIDStr string = strconv.FormatInt(routeURLBksID, 10)
	r, _ := http.NewRequest("GET", "/test?clientId=69&routeId=r&urlId="+routeURLBksIDStr, nil)
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperGet(w, r)
	fmt.Print("Media Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy cb.Breaker
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestBks_HandleSuperGetReq2(t *testing.T) {
	var routeBksStr string = strconv.FormatInt(routeBks, 10)
	var routeURLBksIDStr string = strconv.FormatInt(routeURLBksID, 10)
	r, _ := http.NewRequest("GET", "/test?clientId=ddd&routeId="+routeBksStr+"&urlId="+routeURLBksIDStr, nil)
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "e")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperGet(w, r)
	fmt.Print("Media Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy cb.Breaker
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp bad clientId: ")
	fmt.Println(bdy)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestBks_HandleSuperGetReq3(t *testing.T) {
	var routeBksStr string = strconv.FormatInt(routeBks, 10)
	//var routeURLBksIDStr string = strconv.FormatInt(routeURLBksID, 10)
	r, _ := http.NewRequest("GET", "/test?clientId=69&routeId="+routeBksStr+"&urlId=ee", nil)
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperGet(w, r)
	fmt.Print("Media Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy cb.Breaker
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestBks_HandleSuperGetMethod(t *testing.T) {
	var routeBksStr string = strconv.FormatInt(routeBks, 10)
	var routeURLBksIDStr string = strconv.FormatInt(routeURLBksID, 10)
	r, _ := http.NewRequest("DELETE", "/test?clientId=69&routeId="+routeBksStr+"&urlId="+routeURLBksIDStr, nil)
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperGet(w, r)
	fmt.Print("Media Code: ")
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

func TestBks_HandleSuperGetAuth(t *testing.T) {
	testMode = false
	var routeBksStr string = strconv.FormatInt(routeBks, 10)
	var routeURLBksIDStr string = strconv.FormatInt(routeURLBksID, 10)
	r, _ := http.NewRequest("GET", "/test?clientId=69&routeId="+routeBksStr+"&urlId="+routeURLBksIDStr, nil)
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperGet(w, r)
	fmt.Print("Media Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy cb.Breaker
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusUnauthorized {
		t.Fail()
	}
	testMode = true
}

func TestBks_HandleSuperGet(t *testing.T) {
	var routeBksStr string = strconv.FormatInt(routeBks, 10)
	var routeURLBksIDStr string = strconv.FormatInt(routeURLBksID, 10)
	r, _ := http.NewRequest("GET", "/test?clientId=69&routeId="+routeBksStr+"&urlId="+routeURLBksIDStr, nil)
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperGet(w, r)
	fmt.Print("Media Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy cb.Breaker
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.ID == 0 {
		t.Fail()
	} else {
		bksID = bdy.ID
	}
}

func TestBks_HandleSuperPutReq(t *testing.T) {
	var c cb.Breaker
	//c.ID = bksID
	c.ClientID = clustCidPer
	c.FailoverRouteName = "test2"
	c.FailureCount = 1
	c.FailureThreshold = 3
	c.HealthCheckTimeSeconds = 120
	c.RestRouteID = routeBks
	c.RouteURIID = routeURLBksID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperPut(w, r)
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

func TestBks_HandleSuperPutMedia(t *testing.T) {
	var c cb.Breaker
	c.ID = bksID
	c.ClientID = clustCidPer
	c.FailoverRouteName = "test2"
	//c.FailureCount = 4
	c.FailureThreshold = 1
	c.HealthCheckTimeSeconds = 120
	c.RestRouteID = routeBks
	c.RouteURIID = routeURLBksID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperPut(w, r)
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

func TestBks_HandleSuperPutMethos(t *testing.T) {
	var c cb.Breaker
	c.ID = bksID
	c.ClientID = clustCidPer
	c.FailoverRouteName = "test2"
	//c.FailureCount = 4
	c.FailureThreshold = 1
	c.HealthCheckTimeSeconds = 120
	c.RestRouteID = routeBks
	c.RouteURIID = routeURLBksID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperPut(w, r)
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

func TestBks_HandleSuperPutAuth(t *testing.T) {
	testMode = false
	var c cb.Breaker
	c.ID = bksID
	c.ClientID = clustCidPer
	c.FailoverRouteName = "test2"
	//c.FailureCount = 4
	c.FailureThreshold = 1
	c.HealthCheckTimeSeconds = 120
	c.RestRouteID = routeBks
	c.RouteURIID = routeURLBksID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperPut(w, r)
	fmt.Print("Status Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy BreakerResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusUnauthorized {
		t.Fail()
	}
	testMode = true
}

func TestBks_HandleSuperPut(t *testing.T) {
	var c cb.Breaker
	c.ID = bksID
	c.ClientID = clustCidPer
	c.FailoverRouteName = "test2"
	//c.FailureCount = 4
	c.FailureThreshold = 1
	c.HealthCheckTimeSeconds = 120
	c.RestRouteID = routeBks
	c.RouteURIID = routeURLBksID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperPut(w, r)
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

func TestBks_HandleSuperGet4(t *testing.T) {
	var routeBksStr string = strconv.FormatInt(routeBks, 10)
	var routeURLBksIDStr string = strconv.FormatInt(routeURLBksID, 10)
	r, _ := http.NewRequest("GET", "/test?clientId=69&routeId="+routeBksStr+"&urlId="+routeURLBksIDStr, nil)
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperGet(w, r)
	fmt.Print("Media Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy cb.Breaker
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp get: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.ID == 0 || bdy.FailoverRouteName != "test2" {
		t.Fail()
	}
}

func TestBks_CircuitBreakerTrip(t *testing.T) {
	bk := new(cb.Breaker)
	bk.ClientID = clustCidPer
	bk.RouteURIID = routeURLBksID
	edbBks.Trip(bk)
	edbBks.Trip(bk)
}

func TestBks_HandleSuperStatusReq(t *testing.T) {
	//var routeBksStr string = strconv.FormatInt(routeBks, 10)
	var routeURLBksIDStr string = strconv.FormatInt(routeURLBksID, 10)
	r, _ := http.NewRequest("GET", "/test?clientId=f&urlId="+routeURLBksIDStr, nil)
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBks.HandleBreakerStatusSuper(w, r)
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

func TestBks_HandleSuperStatusReq2(t *testing.T) {
	//var routeBksStr string = strconv.FormatInt(routeBks, 10)
	//var routeURLBksIDStr string = strconv.FormatInt(routeURLBksID, 10)
	r, _ := http.NewRequest("GET", "/test?clientId=69&urlId=rr", nil)
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBks.HandleBreakerStatusSuper(w, r)
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

func TestBks_HandleSuperStatusMethod(t *testing.T) {
	//var routeBksStr string = strconv.FormatInt(routeBks, 10)
	var routeURLBksIDStr string = strconv.FormatInt(routeURLBksID, 10)
	r, _ := http.NewRequest("POST", "/test?clientId=69&urlId="+routeURLBksIDStr, nil)
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBks.HandleBreakerStatusSuper(w, r)
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

func TestBks_HandleSuperStatusAuth(t *testing.T) {
	//var routeBksStr string = strconv.FormatInt(routeBks, 10)
	testMode = false
	var routeURLBksIDStr string = strconv.FormatInt(routeURLBksID, 10)
	r, _ := http.NewRequest("GET", "/test?clientId=69&urlId="+routeURLBksIDStr, nil)
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBks.HandleBreakerStatusSuper(w, r)
	fmt.Print("Status Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy cb.Status
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Status Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusUnauthorized {
		t.Fail()
	}
	testMode = true
}

func TestBks_HandleSuperStatus(t *testing.T) {
	//var routeBksStr string = strconv.FormatInt(routeBks, 10)
	var routeURLBksIDStr string = strconv.FormatInt(routeURLBksID, 10)
	r, _ := http.NewRequest("GET", "/test?clientId=69&urlId="+routeURLBksIDStr, nil)
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBks.HandleBreakerStatusSuper(w, r)
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

func TestBks_HandleSuperResetReq(t *testing.T) {
	var c cb.Breaker
	c.ClientID = clustCidPer
	//c.RouteURIID = routeURLBksID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperReset(w, r)
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

func TestBks_HandleSuperResetMedia(t *testing.T) {
	var c cb.Breaker
	c.ClientID = clustCidPer
	c.RouteURIID = routeURLBksID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperReset(w, r)
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

func TestBks_HandleSuperResetMethod(t *testing.T) {
	var c cb.Breaker
	c.ClientID = clustCidPer
	c.RouteURIID = routeURLBksID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperReset(w, r)
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

func TestBks_HandleSuperResetAuth(t *testing.T) {
	testMode = false
	var c cb.Breaker
	c.ClientID = clustCidPer
	c.RouteURIID = routeURLBksID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperReset(w, r)
	fmt.Print("Status Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy BreakerResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusUnauthorized {
		t.Fail()
	}
	testMode = true
}

func TestBks_HandleSuperReset(t *testing.T) {
	var c cb.Breaker
	c.ClientID = clustCidPer
	c.RouteURIID = routeURLBksID
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperReset(w, r)
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

func TestBks_HandleSuperStatus2(t *testing.T) {
	//var routeBksStr string = strconv.FormatInt(routeBks, 10)
	var routeURLBksIDStr string = strconv.FormatInt(routeURLBksID, 10)
	r, _ := http.NewRequest("GET", "/test?clientId=69&urlId="+routeURLBksIDStr, nil)
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBks.HandleBreakerStatusSuper(w, r)
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

func TestBks_HandleSuperDelReq(t *testing.T) {
	var routeBksStr string = strconv.FormatInt(routeBks, 10)
	var routeURLBksIDStr string = strconv.FormatInt(routeURLBksID, 10)
	r, _ := http.NewRequest("DELETE", "/test?clientId=e&routeId="+routeBksStr+"&urlId="+routeURLBksIDStr, nil)
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperDelete(w, r)
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

func TestBks_HandleSuperDelReq2(t *testing.T) {
	//var routeBksStr string = strconv.FormatInt(routeBks, 10)
	var routeURLBksIDStr string = strconv.FormatInt(routeURLBksID, 10)
	r, _ := http.NewRequest("DELETE", "/test?clientId=69&routeId=ee&urlId="+routeURLBksIDStr, nil)
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperDelete(w, r)
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

func TestBks_HandleSuperDelReq3(t *testing.T) {
	var routeBksStr string = strconv.FormatInt(routeBks, 10)
	//var routeURLBksIDStr string = strconv.FormatInt(routeURLBksID, 10)
	r, _ := http.NewRequest("DELETE", "/test?clientId=69&routeId="+routeBksStr+"&urlId=dd", nil)
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperDelete(w, r)
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

func TestBks_HandleSuperDelMethod(t *testing.T) {
	var routeBksStr string = strconv.FormatInt(routeBks, 10)
	var routeURLBksIDStr string = strconv.FormatInt(routeURLBksID, 10)
	r, _ := http.NewRequest("GET", "/test?clientId=69&routeId="+routeBksStr+"&urlId="+routeURLBksIDStr, nil)
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperDelete(w, r)
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

func TestBks_HandleSuperDelAuth(t *testing.T) {
	testMode = false
	var routeBksStr string = strconv.FormatInt(routeBks, 10)
	var routeURLBksIDStr string = strconv.FormatInt(routeURLBksID, 10)
	r, _ := http.NewRequest("DELETE", "/test?clientId=69&routeId="+routeBksStr+"&urlId="+routeURLBksIDStr, nil)
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperDelete(w, r)
	fmt.Print("Media Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy BreakerResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp delete: ")
	fmt.Println(bdy)
	if w.Code != http.StatusUnauthorized {
		t.Fail()
	}
	testMode = true
}

func TestBks_HandleSuperDel(t *testing.T) {
	var routeBksStr string = strconv.FormatInt(routeBks, 10)
	var routeURLBksIDStr string = strconv.FormatInt(routeURLBksID, 10)
	r, _ := http.NewRequest("DELETE", "/test?clientId=69&routeId="+routeBksStr+"&urlId="+routeURLBksIDStr, nil)
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperDelete(w, r)
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

func TestBks_HandleSuperGet5(t *testing.T) {
	var routeBksStr string = strconv.FormatInt(routeBks, 10)
	var routeURLBksIDStr string = strconv.FormatInt(routeURLBksID, 10)
	r, _ := http.NewRequest("GET", "/test?clientId=69&routeId="+routeBksStr+"&urlId="+routeURLBksIDStr, nil)
	r.Header.Set("u-client-id", "69")
	r.Header.Set("clientId", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBks.HandleBreakerSuperGet(w, r)
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

func TestBks_DeleteClient(t *testing.T) {
	var c mgr.Client
	c.ClientID = clustCidBks
	res := gwRoutesBks.GwDB.DeleteClient(&c)
	if res.Success != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestBks_TestCloseDb(t *testing.T) {
	success := gwRoutesBks.GwDB.CloseDb()
	if success != true {
		t.Fail()
	}
}
