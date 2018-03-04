package handlers

import (
	gwerr "UlboraApiGateway/gwerrors"
	mgr "UlboraApiGateway/managers"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var clustCidErr int64 = 99
var routeErr int64
var routeURLErrID int64
var connectedForErr bool
var edb gwerr.GatewayErrorMonitor
var gwRoutesErr mgr.GatewayRoutes
var h Handler

func TestErr_ConnectForErr(t *testing.T) {
	edb.DbConfig.Host = "localhost:3306"
	edb.DbConfig.DbUser = "admin"
	edb.DbConfig.DbPw = "admin"
	edb.DbConfig.DatabaseName = "ulbora_api_gateway"
	connectedForErr = edb.ConnectDb()
	if connectedForErr != true {
		t.Fail()
	}
	gwRoutesErr.GwDB.DbConfig = edb.DbConfig
	//gwRoutes.GwDB.DbConfig = gwRoutes.GwDB.DbConfig
	//cp.Host = "http://localhost:3010"
	testMode = true
	h.DbConfig = edb.DbConfig
}

// func TestErr_SetManager(t *testing.T) {
// 	SetManager(edb)
// }

func TestErr_InsertClient(t *testing.T) {
	var c mgr.Client
	c.APIKey = "12233hgdd333"
	c.ClientID = clustCidErr
	c.Enabled = true
	c.Level = "small"

	res := gwRoutesErr.GwDB.InsertClient(&c)
	if res.Success == true && res.ID != -1 {
		fmt.Print("new client Id: ")
		fmt.Println(clustCid)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestErr_InsertRestRoute(t *testing.T) {
	var rr mgr.RestRoute
	rr.Route = "content"
	rr.ClientID = clustCidErr

	res := gwRoutesErr.GwDB.InsertRestRoute(&rr)
	if res.Success == true && res.ID != -1 {
		routeErr = res.ID
		fmt.Print("new route Id: ")
		fmt.Println(routeErr)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestErr_InsertRouteURL(t *testing.T) {
	var ru mgr.RouteURL
	ru.Name = "blue"
	ru.URL = "http://www.apigateway.com/blue/"
	ru.Active = false
	ru.RouteID = routeErr
	ru.ClientID = clustCidErr

	res := gwRoutesErr.GwDB.InsertRouteURL(&ru)
	if res.Success == true && res.ID != -1 {
		routeURLErrID = res.ID
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
	ru2.RouteID = routeErr
	ru2.ClientID = clustCidErr

	res2 := gwRoutesErr.GwDB.InsertRouteURL(&ru2)
	if res2.Success == true && res2.ID != -1 {
		//routeURLID33 = res2.ID
		fmt.Print("new route url Id: ")
		fmt.Println(res2.ID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestErr_InsertRouteError(t *testing.T) {
	var e gwerr.GwError
	e.ClientID = clustCidErr
	e.Code = 500
	e.Entered = time.Now()
	e.Message = "internal error"
	e.RestRouteID = routeErr
	e.RouteURIID = routeURLErrID
	suc, err := edb.InsertRouteError(&e)
	if suc != true || err != nil {
		t.Fail()
	}
}

func TestErr_HandleErrorsSuperMedia(t *testing.T) {
	var e gwerr.GwError
	e.ClientID = clustCidErr
	e.Code = 400
	e.Entered = time.Now()
	e.Message = "test error"
	e.RestRouteID = routeErr
	e.RouteURIID = routeURLErrID
	aJSON, _ := json.Marshal(e)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "99")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	h.HandleErrorsSuper(w, r)
	fmt.Print("Media Code: ")
	fmt.Println(w.Code)
	if w.Code != http.StatusUnsupportedMediaType {
		t.Fail()
	}
}

func TestErr_HandleErrorsSuperReq(t *testing.T) {
	var e gwerr.GwError
	e.Code = 400
	e.Entered = time.Now()
	e.Message = "test error"
	e.RestRouteID = routeErr
	e.RouteURIID = routeURLErrID
	aJSON, _ := json.Marshal(e)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "99")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.HandleErrorsSuper(w, r)
	fmt.Print("Req Code: ")
	fmt.Println(w.Code)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestErr_HandleErrorsSuperMethod(t *testing.T) {
	var e gwerr.GwError
	e.ClientID = clustCidErr
	e.Code = 400
	e.Entered = time.Now()
	e.Message = "test error"
	e.RestRouteID = routeErr
	e.RouteURIID = routeURLErrID
	aJSON, _ := json.Marshal(e)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("clientId", "99")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.HandleErrorsSuper(w, r)
	fmt.Print("Status Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy []gwerr.GwError
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp sup: ")
	fmt.Println(bdy)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestErr_HandleErrorsSuper(t *testing.T) {
	var e gwerr.GwError
	e.ClientID = clustCidErr
	e.Code = 400
	e.Entered = time.Now()
	e.Message = "test error"
	e.RestRouteID = routeErr
	e.RouteURIID = routeURLErrID
	aJSON, _ := json.Marshal(e)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("clientId", "99")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.HandleErrorsSuper(w, r)
	fmt.Print("Status Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy []gwerr.GwError
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp sup: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || len(bdy) != 1 {
		t.Fail()
	}
}

func TestErr_HandleErrorsMedia(t *testing.T) {
	var e gwerr.GwError
	//e.ClientID = clustCidErr
	e.Code = 400
	e.Entered = time.Now()
	e.Message = "test error"
	e.RestRouteID = routeErr
	e.RouteURIID = routeURLErrID
	aJSON, _ := json.Marshal(e)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "99")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	h.HandleErrors(w, r)
	fmt.Print("Media Code: ")
	fmt.Println(w.Code)
	if w.Code != http.StatusUnsupportedMediaType {
		t.Fail()
	}
}

func TestErr_HandleErrorsReq(t *testing.T) {
	var e gwerr.GwError
	e.Code = 400
	e.Entered = time.Now()
	e.Message = "test error"
	//e.RestRouteID = routeErr
	e.RouteURIID = routeURLErrID
	aJSON, _ := json.Marshal(e)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "99")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.HandleErrors(w, r)
	fmt.Print("Req Code: ")
	fmt.Println(w.Code)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestErr_HandleErrorsMethod(t *testing.T) {
	var e gwerr.GwError
	//e.ClientID = clustCidErr
	e.Code = 400
	e.Entered = time.Now()
	e.Message = "test error"
	e.RestRouteID = routeErr
	e.RouteURIID = routeURLErrID
	aJSON, _ := json.Marshal(e)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("clientId", "99")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.HandleErrors(w, r)
	fmt.Print("Status Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy []gwerr.GwError
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestErr_HandleErrors(t *testing.T) {
	var e gwerr.GwError
	//e.ClientID = clustCidErr
	e.Code = 400
	e.Entered = time.Now()
	e.Message = "test error"
	e.RestRouteID = routeErr
	e.RouteURIID = routeURLErrID
	aJSON, _ := json.Marshal(e)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("clientId", "99")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.HandleErrors(w, r)
	fmt.Print("Status Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy []gwerr.GwError
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || len(bdy) != 1 {
		t.Fail()
	}
}

func TestErr_DeleteClient(t *testing.T) {
	var c mgr.Client
	c.ClientID = clustCidErr
	res := gwRoutesErr.GwDB.DeleteClient(&c)
	if res.Success != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestErr_TestCloseDb(t *testing.T) {
	success := gwRoutes.GwDB.CloseDb()
	if success != true {
		t.Fail()
	}
}
