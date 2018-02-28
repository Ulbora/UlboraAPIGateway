package handlers

import (
	env "UlboraApiGateway/environment"
	mgr "UlboraApiGateway/managers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var gwRoutes mgr.GatewayRoutes
var clustCid int64 = 97
var routeClust int64
var connectedForCache bool

func Test_ConnectForCache(t *testing.T) {
	gwRoutes.GwDB.DbConfig.Host = "localhost:3306"
	gwRoutes.GwDB.DbConfig.DbUser = "admin"
	gwRoutes.GwDB.DbConfig.DbPw = "admin"
	gwRoutes.GwDB.DbConfig.DatabaseName = "ulbora_api_gateway"
	connectedForCache = gwRoutes.GwDB.ConnectDb()
	if connectedForCache != true {
		t.Fail()
	}
	//gwRoutes.GwDB.DbConfig = gwRoutes.GwDB.DbConfig
	//cp.Host = "http://localhost:3010"
}

func Test_InsertClientForCache(t *testing.T) {
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
func Test_Initialize(t *testing.T) {

	gwRoutes.ClientID = clustCid
	gwRoutes.Route = "testroute"
	//gwRoutes.APIKey = "12345"
	gwRoutes.GwCacheHost = env.GetCacheHost() // "http://localhost:3010"

	res := gwRoutes.SetGatewayRouteStatus()
	if res != true {
		t.Fail()
	}
}

func Test_InsertRestRoute(t *testing.T) {
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

func Test_InsertRouteURL(t *testing.T) {
	var ru mgr.RouteURL
	ru.Name = "blue"
	ru.URL = "http://www.apigateway.com/blue/"
	ru.Active = false
	ru.RouteID = routeClust
	ru.ClientID = clustCid

	res := gwRoutes.GwDB.InsertRouteURL(&ru)
	if res.Success == true && res.ID != -1 {
		//routeURLID3 = res.ID
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

func Test_handleGetRouteStatus(t *testing.T) {
	r, _ := http.NewRequest("GET", "/test?route=testroute", nil)
	r.Header.Set("u-client-id", "97")
	//r.Header.Set("u-api-key", "12345")

	w := httptest.NewRecorder()
	HandleGetRouteStatus(w, r)
	var bdy mgr.GateStatusResponse
	b, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("code: ")
	fmt.Println(w.Code)
	fmt.Print("body: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.Success != true || bdy.RouteModified != true {
		t.Fail()
	}
}

func Test_handleGetRouteStatus2(t *testing.T) {
	r, _ := http.NewRequest("GET", "/test?route=testroute", nil)
	r.Header.Set("u-client-id", "999")
	//r.Header.Set("u-api-key", "12345")

	w := httptest.NewRecorder()
	HandleGetRouteStatus(w, r)
	var bdy mgr.GateStatusResponse
	b, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("code: ")
	fmt.Println(w.Code)
	fmt.Print("body: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.Success == true || bdy.RouteModified == true {
		t.Fail()
	}
}

func Test_handleGetRouteStatus3(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "/test?route=testroute", nil)
	r.Header.Set("u-client-id", "999")
	//r.Header.Set("u-api-key", "12345")

	w := httptest.NewRecorder()
	HandleGetRouteStatus(w, r)
	var bdy mgr.GateStatusResponse
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

func Test_handleDeleteRouteStatus(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "/test?route=testroute", nil)
	r.Header.Set("u-client-id", "97")
	r.Header.Set("u-api-key", "12233hgdd333")

	w := httptest.NewRecorder()
	HandleDeleteRouteStatus(w, r)
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

func Test_handleDeleteRouteStatus2(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "/test?route=testroute", nil)
	r.Header.Set("u-client-id", "97")
	r.Header.Set("u-api-key", "12233hgdd3335")

	w := httptest.NewRecorder()
	HandleDeleteRouteStatus(w, r)
	var bdy mgr.ClusterResponse
	b, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("code: ")
	fmt.Println(w.Code)
	fmt.Print("body: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.Success == true {
		t.Fail()
	}
}

func Test_handleGetClusterGwRoutes(t *testing.T) {
	r, _ := http.NewRequest("GET", "/test?route=content", nil)
	r.Header.Set("u-client-id", "97")
	r.Header.Set("u-api-key", "12233hgdd333")

	w := httptest.NewRecorder()
	HandleGetClusterGwRoutes(w, r)
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

func Test_DeleteClientForCache(t *testing.T) {
	var c mgr.Client
	c.ClientID = clustCid
	res := gwRoutes.GwDB.DeleteClient(&c)
	if res.Success != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func Test_TestCloseDb2(t *testing.T) {
	success := gwRoutes.GwDB.CloseDb()
	if success != true {
		t.Fail()
	}
}
