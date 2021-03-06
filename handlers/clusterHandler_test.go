package handlers

import (
	cb "UlboraApiGateway/circuitbreaker"
	env "UlboraApiGateway/environment"
	gwerr "UlboraApiGateway/gwerrors"
	mgr "UlboraApiGateway/managers"
	gwmon "UlboraApiGateway/monitor"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var gwRoutes mgr.GatewayRoutes
var clErrDB gwerr.GatewayErrorMonitor
var clustCid int64 = 97
var routeClust int64
var routeClustURLID int64
var connectedForCache bool
var clustCbDB cb.CircuitBreaker
var clPermDB gwmon.GatewayPerformanceMonitor
var hcc Handler

func TestClus_ConnectForCache(t *testing.T) {
	gwRoutes.GwDB.DbConfig.Host = "localhost:3306"
	gwRoutes.GwDB.DbConfig.DbUser = "admin"
	gwRoutes.GwDB.DbConfig.DbPw = "admin"
	gwRoutes.GwDB.DbConfig.DatabaseName = "ulbora_api_gateway"
	clustCbDB.DbConfig = gwRoutes.GwDB.DbConfig
	clustCbDB.CacheHost = "http://localhost:3010"
	clErrDB.DbConfig = gwRoutes.GwDB.DbConfig
	clPermDB.DbConfig = gwRoutes.GwDB.DbConfig
	connectedForCache = gwRoutes.GwDB.ConnectDb()
	if connectedForCache != true {
		t.Fail()
	}
	hcc.DbConfig = gwRoutes.GwDB.DbConfig
	hcc.ErrDB = clErrDB
	hcc.MonDB = clPermDB
	hcc.MonDB.CallBatchSize = 1
	//gwRoutes.GwDB.DbConfig = gwRoutes.GwDB.DbConfig
	//cp.Host = "http://localhost:3010"
}

func TestClus_InsertClientForCache(t *testing.T) {
	var c mgr.Client
	c.APIKey = "12233hgdd333"
	c.ClientID = clustCid
	c.Enabled = true
	c.Level = "small"

	res := gwRoutes.GwDB.InsertClient(&c)
	if res.Success == true && res.ID != -1 {
		fmt.Print("new client Id: ")
		fmt.Println(clustCid)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}
func TestClus_Initialize(t *testing.T) {

	gwRoutes.ClientID = clustCid
	gwRoutes.Route = "testroute"
	//gwRoutes.APIKey = "12345"
	gwRoutes.GwCacheHost = env.GetCacheHost() // "http://localhost:3010"

	// res := gwRoutes.SetGatewayRouteStatus()
	// if res != true {
	// 	t.Fail()
	// }
}

func TestClus_InsertRestRoute(t *testing.T) {
	var rr mgr.RestRoute
	rr.Route = "content"
	rr.ClientID = clustCid

	res := gwRoutes.GwDB.InsertRestRoute(&rr)
	if res.Success == true && res.ID != -1 {
		routeClust = res.ID
		fmt.Print("new route Id: ")
		fmt.Println(routeClust)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestClus_InsertRouteURL(t *testing.T) {
	var ru mgr.RouteURL
	ru.Name = "blue"
	ru.URL = "http://www.apigateway.com/blue/"
	ru.Active = false
	ru.RouteID = routeClust
	ru.ClientID = clustCid

	res := gwRoutes.GwDB.InsertRouteURL(&ru)
	if res.Success == true && res.ID != -1 {
		routeClustURLID = res.ID
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
	ru2.RouteID = routeClust
	ru2.ClientID = clustCid

	res2 := gwRoutes.GwDB.InsertRouteURL(&ru2)
	if res2.Success == true && res2.ID != -1 {
		//routeURLID33 = res2.ID
		fmt.Print("new route url Id: ")
		fmt.Println(res2.ID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

// func TestClus_handleGetRouteStatus(t *testing.T) {
// 	r, _ := http.NewRequest("GET", "/test?route=testroute", nil)
// 	r.Header.Set("u-client-id", "97")
// 	//r.Header.Set("u-api-key", "12345")

// 	w := httptest.NewRecorder()
// 	hcc.HandleGetRouteStatus(w, r)
// 	var bdy mgr.GateStatusResponse
// 	b, _ := ioutil.ReadAll(w.Body)
// 	json.Unmarshal([]byte(b), &bdy)
// 	fmt.Print("code: ")
// 	fmt.Println(w.Code)
// 	fmt.Print("body: ")
// 	fmt.Println(bdy)
// 	if w.Code != http.StatusOK || bdy.Success != true || bdy.RouteModified != true {
// 		t.Fail()
// 	}
// }

// func TestClus_handleGetRouteStatusReq(t *testing.T) {
// 	r, _ := http.NewRequest("GET", "/test?route=testroute", nil)
// 	r.Header.Set("u-client-id", "999")
// 	//r.Header.Set("u-api-key", "12345")

// 	w := httptest.NewRecorder()
// 	hcc.HandleGetRouteStatus(w, r)
// 	var bdy mgr.GateStatusResponse
// 	b, _ := ioutil.ReadAll(w.Body)
// 	json.Unmarshal([]byte(b), &bdy)
// 	fmt.Print("code: ")
// 	fmt.Println(w.Code)
// 	fmt.Print("body: ")
// 	fmt.Println(bdy)
// 	if w.Code != http.StatusOK || bdy.Success == true || bdy.RouteModified == true {
// 		t.Fail()
// 	}
// }

// func TestClus_handleGetRouteStatus3(t *testing.T) {
// 	r, _ := http.NewRequest("DELETE", "/test?route=testroute", nil)
// 	r.Header.Set("u-client-id", "999")
// 	//r.Header.Set("u-api-key", "12345")

// 	w := httptest.NewRecorder()
// 	hcc.HandleGetRouteStatus(w, r)
// 	var bdy mgr.GateStatusResponse
// 	b, _ := ioutil.ReadAll(w.Body)
// 	json.Unmarshal([]byte(b), &bdy)
// 	fmt.Print("code: ")
// 	fmt.Println(w.Code)
// 	fmt.Print("body: ")
// 	fmt.Println(bdy)
// 	if w.Code != http.StatusNotFound {
// 		t.Fail()
// 	}
// }

// func TestClus_handleDeleteRouteStatusMethod(t *testing.T) {
// 	r, _ := http.NewRequest("GET", "/test?route=testroute", nil)
// 	r.Header.Set("u-client-id", "97")
// 	r.Header.Set("u-api-key", "12233hgdd333")

// 	w := httptest.NewRecorder()
// 	hcc.HandleDeleteRouteStatus(w, r)
// 	var bdy mgr.ClusterResponse
// 	b, _ := ioutil.ReadAll(w.Body)
// 	json.Unmarshal([]byte(b), &bdy)
// 	fmt.Print("code: ")
// 	fmt.Println(w.Code)
// 	fmt.Print("body: ")
// 	fmt.Println(bdy)
// 	if w.Code != http.StatusNotFound {
// 		t.Fail()
// 	}
// }

// func TestClus_handleDeleteRouteStatus(t *testing.T) {
// 	r, _ := http.NewRequest("DELETE", "/test?route=testroute", nil)
// 	r.Header.Set("u-client-id", "97")
// 	r.Header.Set("u-api-key", "12233hgdd333")

// 	w := httptest.NewRecorder()
// 	hcc.HandleDeleteRouteStatus(w, r)
// 	var bdy mgr.ClusterResponse
// 	b, _ := ioutil.ReadAll(w.Body)
// 	json.Unmarshal([]byte(b), &bdy)
// 	fmt.Print("code: ")
// 	fmt.Println(w.Code)
// 	fmt.Print("body: ")
// 	fmt.Println(bdy)
// 	if w.Code != http.StatusOK || bdy.Success != true {
// 		t.Fail()
// 	}
// }

// func TestClus_handleDeleteRouteStatus2(t *testing.T) {
// 	r, _ := http.NewRequest("DELETE", "/test?route=testroute", nil)
// 	r.Header.Set("u-client-id", "97")
// 	r.Header.Set("u-api-key", "12233hgdd3335")

// 	w := httptest.NewRecorder()
// 	hcc.HandleDeleteRouteStatus(w, r)
// 	var bdy mgr.ClusterResponse
// 	b, _ := ioutil.ReadAll(w.Body)
// 	json.Unmarshal([]byte(b), &bdy)
// 	fmt.Print("code: ")
// 	fmt.Println(w.Code)
// 	fmt.Print("body: ")
// 	fmt.Println(bdy)
// 	if w.Code != http.StatusOK || bdy.Success == true {
// 		t.Fail()
// 	}
// }

func TestClus_handleGetClusterGwRoutesMethod(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "/test?route=content", nil)
	r.Header.Set("u-client-id", "97")
	r.Header.Set("u-api-key", "12233hgdd333")

	w := httptest.NewRecorder()
	hcc.HandleGetClusterGwRoutes(w, r)
	var bdy = make([]mgr.GatewayRouteURL, 0)
	b, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("code: ")
	fmt.Println(w.Code)
	fmt.Print("body: ")
	fmt.Println(bdy)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestClus_handleGetClusterGwRoutesReq(t *testing.T) {
	r, _ := http.NewRequest("GET", "/test?route=content", nil)
	//r.Header.Set("u-client-id", "97")
	//r.Header.Set("u-api-key", "12233hgdd333")

	w := httptest.NewRecorder()
	hcc.HandleGetClusterGwRoutes(w, r)
	var bdy = make([]mgr.GatewayRouteURL, 0)
	b, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("code: ")
	fmt.Println(w.Code)
	fmt.Print("body: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || len(bdy) != 0 {
		t.Fail()
	}
}

func TestClus_handleGetClusterGwRoutes(t *testing.T) {
	r, _ := http.NewRequest("GET", "/test?route=content", nil)
	r.Header.Set("u-client-id", "97")
	r.Header.Set("u-api-key", "12233hgdd333")

	w := httptest.NewRecorder()
	hcc.HandleGetClusterGwRoutes(w, r)
	var bdy = make([]mgr.GatewayRouteURL, 0)
	b, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("code: ")
	fmt.Println(w.Code)
	fmt.Print("body: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || len(bdy) != 2 {
		t.Fail()
	}
}

func TestClus_handleClearClusterGwRoutesMethod(t *testing.T) {
	r, _ := http.NewRequest("GET", "/test?route=content", nil)
	r.Header.Set("u-client-id", "97")
	r.Header.Set("u-api-key", "12233hgdd333")

	w := httptest.NewRecorder()
	hcc.HandleClearClusterGwRoutes(w, r)
	var bdy = make([]mgr.GatewayRouteURL, 0)
	b, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("code: ")
	fmt.Println(w.Code)
	fmt.Print("body: ")
	fmt.Println(bdy)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestClus_handleClearClusterGwRoutesReq(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "/test?route=content1", nil)
	//r.Header.Set("u-client-id", "97")
	//r.Header.Set("u-api-key", "12233hgdd333")

	w := httptest.NewRecorder()
	hcc.HandleClearClusterGwRoutes(w, r)
	//var bdy = make([]mgr.GatewayRouteURL, 0)
	var bdy mgr.ClusterResponse
	b, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("code clear req: ")
	fmt.Println(w.Code)
	fmt.Print("body: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.Success == true {
		t.Fail()
	}
}

func TestClus_handleClearClusterGwRoutes(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "/test?route=content", nil)
	r.Header.Set("u-client-id", "97")
	r.Header.Set("u-api-key", "12233hgdd333")

	w := httptest.NewRecorder()
	hcc.HandleClearClusterGwRoutes(w, r)
	//var bdy = make([]mgr.GatewayRouteURL, 0)
	var bdy mgr.ClusterResponse
	b, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("code: ")
	fmt.Println(w.Code)
	fmt.Print("body: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.Success != true {
		t.Fail()
	}
}

func TestClus_TripClusterBreakerMedia(t *testing.T) {
	var b ClusterBreaker
	b.ClientID = clustCid
	b.FailureThreshold = 2
	b.HealthCheckTimeSeconds = 120
	b.FailoverRouteName = "blue"
	b.OpenFailCode = 500
	b.RestRouteID = routeClust
	b.RouteURIID = routeClustURLID
	b.Route = "content"
	aJSON, _ := json.Marshal(b)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "97")
	r.Header.Set("u-api-key", "12233hgdd333")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hcc.HandleTripClusterBreaker(w, r)
	//hcc.HandleTripClusterBreaker(w, r)
	//hcc.HandleTripClusterBreaker(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	by, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(by), &bdy)
	fmt.Print("Resp in trip: ")
	fmt.Println(bdy)
	if w.Code != http.StatusUnsupportedMediaType {
		t.Fail()
	}
}

func TestClus_TripClusterBreakerReq(t *testing.T) {
	var b ClusterBreaker
	b.ClientID = clustCid
	b.FailureThreshold = 2
	b.HealthCheckTimeSeconds = 120
	b.FailoverRouteName = "blue"
	b.OpenFailCode = 500
	//b.RestRouteID = routeClust
	//b.RouteURIID = routeClustURLID
	b.Route = "content"
	aJSON, _ := json.Marshal(b)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "97")
	r.Header.Set("u-api-key", "12233hgdd333")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hcc.HandleTripClusterBreaker(w, r)
	//hcc.HandleTripClusterBreaker(w, r)
	//hcc.HandleTripClusterBreaker(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	by, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(by), &bdy)
	fmt.Print("Resp in trip: ")
	fmt.Println(bdy)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestClus_TripClusterBreakerMethod(t *testing.T) {
	var b ClusterBreaker
	b.ClientID = clustCid
	b.FailureThreshold = 2
	b.HealthCheckTimeSeconds = 120
	b.FailoverRouteName = "blue"
	b.OpenFailCode = 500
	b.RestRouteID = routeClust
	b.RouteURIID = routeClustURLID
	b.Route = "content"
	aJSON, _ := json.Marshal(b)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "97")
	r.Header.Set("u-api-key", "12233hgdd333")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hcc.HandleTripClusterBreaker(w, r)
	//hcc.HandleTripClusterBreaker(w, r)
	//hcc.HandleTripClusterBreaker(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	by, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(by), &bdy)
	fmt.Print("Resp in trip: ")
	fmt.Println(bdy)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestClus_TripClusterBreaker1(t *testing.T) {
	var b ClusterBreaker
	b.ClientID = clustCid
	b.FailureThreshold = 2
	b.HealthCheckTimeSeconds = 120
	b.FailoverRouteName = "blue"
	b.OpenFailCode = 500
	b.RestRouteID = routeClust
	b.RouteURIID = routeClustURLID
	b.Route = "content"
	aJSON, _ := json.Marshal(b)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "97")
	r.Header.Set("u-api-key", "12233hgdd333")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hcc.HandleTripClusterBreaker(w, r)
	//hcc.HandleTripClusterBreaker(w, r)
	//hcc.HandleTripClusterBreaker(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	by, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(by), &bdy)
	fmt.Print("Resp in trip: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.Success != true {
		t.Fail()
	}
}

func TestClus_TripClusterBreaker2(t *testing.T) {
	var b ClusterBreaker
	b.ClientID = clustCid
	b.FailureThreshold = 2
	b.HealthCheckTimeSeconds = 120
	b.FailoverRouteName = "blue"
	b.OpenFailCode = 500
	b.RestRouteID = routeClust
	b.RouteURIID = routeClustURLID
	b.Route = "content"
	aJSON, _ := json.Marshal(b)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "97")
	r.Header.Set("u-api-key", "12233hgdd333")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hcc.HandleTripClusterBreaker(w, r)
	//hcc.HandleTripClusterBreaker(w, r)
	//hcc.HandleTripClusterBreaker(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	by, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(by), &bdy)
	fmt.Print("Resp in trip: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.Success != true {
		t.Fail()
	}
}

func TestClus_TripClusterBreaker3(t *testing.T) {
	var b ClusterBreaker
	b.ClientID = clustCid
	b.FailureThreshold = 2
	b.HealthCheckTimeSeconds = 120
	b.FailoverRouteName = "blue"
	b.OpenFailCode = 500
	b.RestRouteID = routeClust
	b.RouteURIID = routeClustURLID
	b.Route = "content"
	aJSON, _ := json.Marshal(b)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "97")
	r.Header.Set("u-api-key", "12233hgdd333")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hcc.HandleTripClusterBreaker(w, r)
	//hcc.HandleTripClusterBreaker(w, r)
	//hcc.HandleTripClusterBreaker(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	by, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(by), &bdy)
	fmt.Print("Resp in trip: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.Success != true {
		t.Fail()
	}
}

func TestClus_GetBreakerStatus(t *testing.T) {
	res := clustCbDB.GetStatus(clustCid, routeClustURLID)
	fmt.Print("routes status: ")
	fmt.Println(res)
	if res.Open != true {
		t.Fail()
	}
}

func TestClus_ResetClusterBreakerMedia(t *testing.T) {
	var b ClusterBreaker
	b.RouteURIID = routeClustURLID
	aJSON, _ := json.Marshal(b)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "97")
	r.Header.Set("u-api-key", "12233hgdd333")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hcc.HandleResetClusterBreaker(w, r)
	fmt.Print("Code in reset: ")
	fmt.Println(w.Code)
	by, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(by), &bdy)
	fmt.Print("Resp in trip: ")
	fmt.Println(bdy)
	if w.Code != http.StatusUnsupportedMediaType {
		t.Fail()
	}
}

func TestClus_ResetClusterBreakerReq(t *testing.T) {
	var b ClusterBreaker
	//b.RouteURIID = routeClustURLID
	aJSON, _ := json.Marshal(b)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "97")
	r.Header.Set("u-api-key", "12233hgdd333")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hcc.HandleResetClusterBreaker(w, r)
	fmt.Print("Code in reset: ")
	fmt.Println(w.Code)
	by, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(by), &bdy)
	fmt.Print("Resp in trip: ")
	fmt.Println(bdy)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestClus_ResetClusterBreakerMethod(t *testing.T) {
	var b ClusterBreaker
	b.RouteURIID = routeClustURLID
	aJSON, _ := json.Marshal(b)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "97")
	r.Header.Set("u-api-key", "12233hgdd333")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hcc.HandleResetClusterBreaker(w, r)
	fmt.Print("Code in reset: ")
	fmt.Println(w.Code)
	by, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(by), &bdy)
	fmt.Print("Resp in trip: ")
	fmt.Println(bdy)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestClus_ResetClusterBreaker(t *testing.T) {
	var b ClusterBreaker
	b.RouteURIID = routeClustURLID
	aJSON, _ := json.Marshal(b)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "97")
	r.Header.Set("u-api-key", "12233hgdd333")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hcc.HandleResetClusterBreaker(w, r)
	fmt.Print("Code in reset: ")
	fmt.Println(w.Code)
	by, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(by), &bdy)
	fmt.Print("Resp in trip: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.Success != true {
		t.Fail()
	}
}

func TestClus_GetBreakerStatus2(t *testing.T) {
	res := clustCbDB.GetStatus(clustCid, routeClustURLID)
	fmt.Print("routes status: ")
	fmt.Println(res)
	if res.Open != false {
		t.Fail()
	}
}

func TestClus_ClusterSaveRouteErrorMedia(t *testing.T) {
	var el ErrorLog
	el.ClientID = clustCid
	el.RouteID = routeClust
	el.RouteURIID = routeClustURLID
	el.ErrCode = 400
	el.Message = "failed in test"
	aJSON, _ := json.Marshal(el)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "97")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hcc.HandleClusterSaveRouteError(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	by, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(by), &bdy)
	fmt.Print("Resp in error log: ")
	fmt.Println(bdy)
	if w.Code != http.StatusUnsupportedMediaType {
		t.Fail()
	}
}

func TestClus_ClusterSaveRouteErrorReq(t *testing.T) {
	var el ErrorLog
	//el.ClientID = clustCid
	//el.RouteID = routeClust
	el.RouteURIID = routeClustURLID
	el.ErrCode = 400
	el.Message = "failed in test"
	aJSON, _ := json.Marshal(el)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "97")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hcc.HandleClusterSaveRouteError(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	by, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(by), &bdy)
	fmt.Print("Resp in error log: ")
	fmt.Println(bdy)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestClus_ClusterSaveRouteErrorMethod(t *testing.T) {
	var el ErrorLog
	el.ClientID = clustCid
	el.RouteID = routeClust
	el.RouteURIID = routeClustURLID
	el.ErrCode = 400
	el.Message = "failed in test"
	aJSON, _ := json.Marshal(el)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "97")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hcc.HandleClusterSaveRouteError(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	by, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(by), &bdy)
	fmt.Print("Resp in error log: ")
	fmt.Println(bdy)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestClus_ClusterSaveRouteError(t *testing.T) {
	var el ErrorLog
	el.ClientID = clustCid
	el.RouteID = routeClust
	el.RouteURIID = routeClustURLID
	el.ErrCode = 400
	el.Message = "failed in test"
	aJSON, _ := json.Marshal(el)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "97")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hcc.HandleClusterSaveRouteError(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	by, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(by), &bdy)
	fmt.Print("Resp in error log: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.Success != true {
		t.Fail()
	}
}

func TestClus_GetRouteError(t *testing.T) {
	var e gwerr.GwError
	e.ClientID = clustCid
	e.RestRouteID = routeClust
	e.RouteURIID = routeClustURLID
	res := clErrDB.GetRouteError(&e)
	fmt.Println("")
	fmt.Print("found gw error list: ")
	fmt.Println(res)
	if len(*res) == 0 || (*res)[0].Code != 400 || (*res)[0].Message != "failed in test" {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestClus_ClusterSaveRoutePerformanceMedia(t *testing.T) {
	var p PerformanceLog
	p.ClientID = clustCid
	p.RouteID = routeClust
	p.RouteURIID = routeClustURLID
	p.Latency = 100
	aJSON, _ := json.Marshal(p)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "97")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hcc.HandleClusterSaveRoutePerformance(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	by, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(by), &bdy)
	fmt.Print("Resp in perf log: ")
	fmt.Println(bdy)
	if w.Code != http.StatusUnsupportedMediaType {
		t.Fail()
	}
}

func TestClus_ClusterSaveRoutePerformanceReq(t *testing.T) {
	var p PerformanceLog
	//p.ClientID = clustCid
	//p.RouteID = routeClust
	p.RouteURIID = routeClustURLID
	p.Latency = 100
	aJSON, _ := json.Marshal(p)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "97")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hcc.HandleClusterSaveRoutePerformance(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	by, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(by), &bdy)
	fmt.Print("Resp in perf log: ")
	fmt.Println(bdy)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestClus_ClusterSaveRoutePerformanceMethod(t *testing.T) {
	var p PerformanceLog
	p.ClientID = clustCid
	p.RouteID = routeClust
	p.RouteURIID = routeClustURLID
	p.Latency = 100
	aJSON, _ := json.Marshal(p)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "97")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hcc.HandleClusterSaveRoutePerformance(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	by, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(by), &bdy)
	fmt.Print("Resp in perf log: ")
	fmt.Println(bdy)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestClus_ClusterSaveRoutePerformance(t *testing.T) {
	var p PerformanceLog
	p.ClientID = clustCid
	p.RouteID = routeClust
	p.RouteURIID = routeClustURLID
	p.Latency = 100
	aJSON, _ := json.Marshal(p)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "97")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hcc.HandleClusterSaveRoutePerformance(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	by, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(by), &bdy)
	fmt.Print("Resp in perf log: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.Success != true {
		t.Fail()
	}
}

func TestClus_ClusterSaveRoutePerformance2(t *testing.T) {
	var p PerformanceLog
	p.ClientID = clustCid
	p.RouteID = routeClust
	p.RouteURIID = routeClustURLID
	p.Latency = 100
	aJSON, _ := json.Marshal(p)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "97")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hcc.HandleClusterSaveRoutePerformance(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	by, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(by), &bdy)
	fmt.Print("Resp in perf log: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.Success != true {
		t.Fail()
	}
}

func TestClus_GetRoutePerformance(t *testing.T) {
	var p gwmon.GwPerformance
	p.ClientID = clustCid
	p.RestRouteID = routeClust
	p.RouteURIID = routeClustURLID
	res := clPermDB.GetRoutePerformance(&p)
	fmt.Println("")
	fmt.Print("found gw perform list: ")
	fmt.Println(res)
	if len(*res) == 0 || (*res)[0].Calls != 2 || (*res)[0].LatencyMsTotal != 200 {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestClus_DeleteClientForCache(t *testing.T) {
	var c mgr.Client
	c.ClientID = clustCid
	res := gwRoutes.GwDB.DeleteClient(&c)
	if res.Success != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestClus_TestCloseDb2(t *testing.T) {
	success := gwRoutes.GwDB.CloseDb()
	if success != true {
		t.Fail()
	}
}
