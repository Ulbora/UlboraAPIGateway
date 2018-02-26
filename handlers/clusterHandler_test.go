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
func Test_DeleteClientForCache(t *testing.T) {
	var c mgr.Client
	c.ClientID = clustCid
	res := gwRoutes.GwDB.DeleteClient(&c)
	if res.Success != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}
