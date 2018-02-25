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

func Test_Initialize(t *testing.T) {
	var clustCid int64 = 99
	var gwRoutes mgr.GatewayRoutes
	gwRoutes.ClientID = clustCid
	gwRoutes.Route = "testroute"
	gwRoutes.APIKey = "12345"
	gwRoutes.GwCacheHost = env.GetCacheHost() // "http://localhost:3010"

	res := gwRoutes.SetGatewayRouteStatus()
	if res != true {
		t.Fail()
	}
}

func Test_handleGetRouteStatus(t *testing.T) {
	r, _ := http.NewRequest("GET", "/test?route=testroute", nil)
	r.Header.Set("u-client-id", "99")
	r.Header.Set("u-api-key", "12345")

	w := httptest.NewRecorder()
	HandleGetRouteStatus(w, r)
	var bdy mgr.GateStatusResponse
	b, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("code: ")
	fmt.Println(w.Code)
	fmt.Print("body: ")
	fmt.Println(w.Body)
	if w.Code != http.StatusOK || bdy.Success != true || bdy.RouteModified != true {
		t.Fail()
	}
}

func Test_handleGetRouteStatus2(t *testing.T) {
	r, _ := http.NewRequest("GET", "/test?route=testroute", nil)
	r.Header.Set("u-client-id", "999")
	r.Header.Set("u-api-key", "12345")

	w := httptest.NewRecorder()
	HandleGetRouteStatus(w, r)
	var bdy mgr.GateStatusResponse
	b, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("code: ")
	fmt.Println(w.Code)
	fmt.Print("body: ")
	fmt.Println(w.Body)
	if w.Code != http.StatusOK || bdy.Success == true || bdy.RouteModified == true {
		t.Fail()
	}
}
