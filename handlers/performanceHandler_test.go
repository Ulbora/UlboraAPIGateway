package handlers

import (
	mgr "UlboraApiGateway/managers"
	gwmon "UlboraApiGateway/monitor"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	//"net/http"
	//"net/http/httptest"
	"testing"
	//"time"
)

var clustCidPer int64 = 69
var routePer int64
var routeURLPerID int64
var connectedForPer bool
var edbPer gwmon.GatewayPerformanceMonitor
var gwRoutesPer mgr.GatewayRoutes
var hPer Handler

func TestPer_ConnectFor(t *testing.T) {
	edbPer.DbConfig.Host = "localhost:3306"
	edbPer.DbConfig.DbUser = "admin"
	edbPer.DbConfig.DbPw = "admin"
	edbPer.DbConfig.DatabaseName = "ulbora_api_gateway"
	connectedForPer = edbPer.ConnectDb()
	if connectedForPer != true {
		t.Fail()
	}
	gwRoutesPer.GwDB.DbConfig = edbPer.DbConfig
	//gwRoutes.GwDB.DbConfig = gwRoutes.GwDB.DbConfig
	//cp.Host = "http://localhost:3010"
	testMode = true
	hPer.DbConfig = edb.DbConfig
}

// func TestErr_SetManager(t *testing.T) {
// 	SetManager(edb)
// }

func TestPer_InsertClient(t *testing.T) {
	var c mgr.Client
	c.APIKey = "12233hgdd333"
	c.ClientID = clustCidPer
	c.Enabled = true
	c.Level = "small"

	res := gwRoutesPer.GwDB.InsertClient(&c)
	if res.Success == true && res.ID != -1 {
		fmt.Print("new client Id: ")
		fmt.Println(clustCidPer)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestPer_InsertRestRoute(t *testing.T) {
	var rr mgr.RestRoute
	rr.Route = "content"
	rr.ClientID = clustCidPer

	res := gwRoutesPer.GwDB.InsertRestRoute(&rr)
	if res.Success == true && res.ID != -1 {
		routePer = res.ID
		fmt.Print("new route Id: ")
		fmt.Println(routeErr)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestPer_InsertRouteURL(t *testing.T) {
	var ru mgr.RouteURL
	ru.Name = "blue"
	ru.URL = "http://www.apigateway.com/blue/"
	ru.Active = false
	ru.RouteID = routePer
	ru.ClientID = clustCidPer

	res := gwRoutesPer.GwDB.InsertRouteURL(&ru)
	if res.Success == true && res.ID != -1 {
		routeURLPerID = res.ID
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
	ru2.RouteID = routePer
	ru2.ClientID = clustCidPer

	res2 := gwRoutesPer.GwDB.InsertRouteURL(&ru2)
	if res2.Success == true && res2.ID != -1 {
		//routeURLID33 = res2.ID
		fmt.Print("new route url Id: ")
		fmt.Println(res2.ID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestPer_HandlePerSuperMedia(t *testing.T) {
	var p gwmon.GwPerformance
	p.ClientID = clustCidPer
	p.LatencyMsTotal = 125
	p.Entered = time.Now()
	p.RestRouteID = routePer
	p.RouteURIID = routeURLPerID
	aJSON, _ := json.Marshal(p)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	h.HandlePeformanceSuper(w, r)
	fmt.Print("Media Code: ")
	fmt.Println(w.Code)
	if w.Code != http.StatusUnsupportedMediaType {
		t.Fail()
	}
}

func TestPer_HandleErrorsSuperReq(t *testing.T) {
	var p gwmon.GwPerformance
	//p.ClientID = clustCidPer
	p.LatencyMsTotal = 125
	p.Entered = time.Now()
	p.RestRouteID = routePer
	p.RouteURIID = routeURLPerID
	aJSON, _ := json.Marshal(p)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.HandlePeformanceSuper(w, r)
	fmt.Print("Req Code: ")
	fmt.Println(w.Code)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestPer_HandleErrorsSuper(t *testing.T) {
	var p gwmon.GwPerformance
	p.ClientID = clustCidPer
	p.LatencyMsTotal = 125
	p.Entered = time.Now()
	p.RestRouteID = routePer
	p.RouteURIID = routeURLPerID
	aJSON, _ := json.Marshal(p)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.HandlePeformanceSuper(w, r)
	fmt.Print("Status Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy []gwmon.GwPerformance
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK && len(bdy) != 0 {
		t.Fail()
	}
}

func TestPer_HandleErrorsMedia(t *testing.T) {
	var p gwmon.GwPerformance
	//p.ClientID = clustCidPer
	p.LatencyMsTotal = 125
	p.Entered = time.Now()
	p.RestRouteID = routePer
	p.RouteURIID = routeURLPerID
	aJSON, _ := json.Marshal(p)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	h.HandlePeformance(w, r)
	fmt.Print("Media Code: ")
	fmt.Println(w.Code)
	if w.Code != http.StatusUnsupportedMediaType {
		t.Fail()
	}
}

func TestPer_HandleErrorsReq(t *testing.T) {
	var p gwmon.GwPerformance
	//p.ClientID = clustCidPer
	p.LatencyMsTotal = 125
	p.Entered = time.Now()
	//p.RestRouteID = routePer
	p.RouteURIID = routeURLPerID
	aJSON, _ := json.Marshal(p)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.HandlePeformance(w, r)
	fmt.Print("Req Code: ")
	fmt.Println(w.Code)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestPer_HandleErrors(t *testing.T) {
	var p gwmon.GwPerformance
	//p.ClientID = clustCidPer
	p.LatencyMsTotal = 125
	p.Entered = time.Now()
	p.RestRouteID = routePer
	p.RouteURIID = routeURLPerID
	aJSON, _ := json.Marshal(p)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "69")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.HandlePeformance(w, r)
	fmt.Print("Status Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy []gwmon.GwPerformance
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK && len(bdy) != 0 {
		t.Fail()
	}
}

func TestPer_DeleteClient(t *testing.T) {
	var c mgr.Client
	c.ClientID = clustCidPer
	res := gwRoutesPer.GwDB.DeleteClient(&c)
	if res.Success != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestPer_TestCloseDb(t *testing.T) {
	success := gwRoutesPer.GwDB.CloseDb()
	if success != true {
		t.Fail()
	}
}
