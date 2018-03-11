package handlers

import (
	"strconv"
	//gwerr "UlboraApiGateway/gwerrors"
	mgr "UlboraApiGateway/managers"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	//"time"
)

var rrHandid int64 = 26
var rrID int64

//var routeErr int64
//var routeURLErrID int64
var connectedForRR bool
var gwRR mgr.GatewayDB

//var gwRoutesErr mgr.GatewayRoutes
var hrr Handler

func TestRRt_Connect(t *testing.T) {
	gwRR.DbConfig.Host = "localhost:3306"
	gwRR.DbConfig.DbUser = "admin"
	gwRR.DbConfig.DbPw = "admin"
	gwRR.DbConfig.DatabaseName = "ulbora_api_gateway"
	connectedForRR = gwRR.ConnectDb()
	if connectedForRR != true {
		t.Fail()
	}
	//gwRoutesErr.GwDB.DbConfig = edb.DbConfig
	//gwRoutes.GwDB.DbConfig = gwRoutes.GwDB.DbConfig
	//cp.Host = "http://localhost:3010"
	testMode = true
	hrr.DbConfig = gwRR.DbConfig
}

func TestRRt_HandleClientInsert(t *testing.T) {
	var c mgr.Client
	c.ClientID = rrHandid
	c.APIKey = "123456"
	c.Enabled = true
	c.Level = "small"
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "26")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrr.HandleClientPost(w, r)
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

func TestRRt_HandleRestRouteInsertMedia(t *testing.T) {
	var rr mgr.RestRoute
	rr.ClientID = rrHandid
	rr.Route = "test"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "26")
	r.Header.Set("u-api-key", "12233hgdd3335")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrr.HandleRestRouteSuperPost(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusUnsupportedMediaType {
		t.Fail()
	}
}

func TestRRt_HandleRestRouteInsertMethod(t *testing.T) {
	var rr mgr.RestRoute
	rr.ClientID = rrHandid
	rr.Route = "test"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("GET", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "26")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrr.HandleRestRouteSuperPost(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestRRt_HandleRestRouteInsertReq(t *testing.T) {
	var rr mgr.RestRoute
	//rr.ClientID = rrHandid
	rr.Route = "test"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "26")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrr.HandleRestRouteSuperPost(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestRRt_HandleRestRouteInsert(t *testing.T) {
	var rr mgr.RestRoute
	rr.ClientID = rrHandid
	rr.Route = "test"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "26")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrr.HandleRestRouteSuperPost(w, r)
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
		rrID = bdy.ID
	}
}

func TestRRt_HandleRestRouteUpdateMedia(t *testing.T) {
	var rr mgr.RestRoute
	rr.ClientID = rrHandid
	rr.Route = "test"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "26")
	r.Header.Set("u-api-key", "12233hgdd3335")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrr.HandleRestRouteSuperPut(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusUnsupportedMediaType {
		t.Fail()
	}
}

func TestRRt_HandleRestRouteUpdateMethod(t *testing.T) {
	var rr mgr.RestRoute
	rr.ClientID = rrHandid
	rr.Route = "test"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("GET", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "26")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrr.HandleRestRouteSuperPut(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestRRt_HandleRestRouteUpdateReq(t *testing.T) {
	var rr mgr.RestRoute
	//rr.ClientID = rrHandid
	rr.Route = "test"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "26")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrr.HandleRestRouteSuperPut(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestRRt_HandleRestRouteUpdate(t *testing.T) {
	var rr mgr.RestRoute
	rr.ID = rrID
	rr.ClientID = rrHandid
	rr.Route = "test2"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "26")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrr.HandleRestRouteSuperPut(w, r)
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

func TestRRt_HandleGetMethod(t *testing.T) {
	var idStr string = strconv.FormatInt(rrID, 10)
	var CIDStr string = strconv.FormatInt(rrHandid, 10)
	r, _ := http.NewRequest("POST", "/test?id="+idStr+"&clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "26")
	r.Header.Set("clientId", "26")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrr.HandleRestRouteSuperGet(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.RestRoute
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get: ")
	fmt.Println(bdy)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestRRt_HandleGetReq(t *testing.T) {
	var idStr string = strconv.FormatInt(rrID, 10)
	//var CIDStr string = strconv.FormatInt(rrHandid, 10)
	r, _ := http.NewRequest("GET", "/test?id="+idStr+"&clientId=r", nil)
	r.Header.Set("u-client-id", "26")
	r.Header.Set("clientId", "26")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrr.HandleRestRouteSuperGet(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.RestRoute
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get: ")
	fmt.Println(bdy)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestRRt_HandleGet(t *testing.T) {
	var idStr string = strconv.FormatInt(rrID, 10)
	var CIDStr string = strconv.FormatInt(rrHandid, 10)
	r, _ := http.NewRequest("GET", "/test?id="+idStr+"&clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "26")
	r.Header.Set("clientId", "26")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrr.HandleRestRouteSuperGet(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.RestRoute
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.ClientID != rrHandid || bdy.ID != rrID {
		t.Fail()
	}
}

func TestRRt_HandleGetListMethod(t *testing.T) {
	//var idStr string = strconv.FormatInt(rrID, 10)
	var CIDStr string = strconv.FormatInt(rrHandid, 10)
	r, _ := http.NewRequest("DELETE", "/test?clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "26")
	r.Header.Set("clientId", "26")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrr.HandleRestRouteSuperList(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy []mgr.RestRoute
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get List: ")
	fmt.Println(bdy)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestRRt_HandleGetListReq(t *testing.T) {
	//var idStr string = strconv.FormatInt(rrID, 10)
	//var CIDStr string = strconv.FormatInt(rrHandid, 10)
	r, _ := http.NewRequest("GET", "/test?clientId=w", nil)
	r.Header.Set("u-client-id", "26")
	r.Header.Set("clientId", "26")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrr.HandleRestRouteSuperList(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy []mgr.RestRoute
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get List: ")
	fmt.Println(bdy)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestRRt_HandleGetList(t *testing.T) {
	//var idStr string = strconv.FormatInt(rrID, 10)
	var CIDStr string = strconv.FormatInt(rrHandid, 10)
	r, _ := http.NewRequest("GET", "/test?clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "26")
	r.Header.Set("clientId", "26")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrr.HandleRestRouteSuperList(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy []mgr.RestRoute
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get List: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || len(bdy) != 1 {
		t.Fail()
	}
}

func TestRRt_HandleDeleteMethod(t *testing.T) {
	var idStr string = strconv.FormatInt(rrID, 10)
	var CIDStr string = strconv.FormatInt(rrHandid, 10)
	r, _ := http.NewRequest("POST", "/test?id="+idStr+"&clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "26")
	r.Header.Set("clientId", "26")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrr.HandleRestRouteSuperDelete(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.RestRoute
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get: ")
	fmt.Println(bdy)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestRRt_HandleDeleteReq(t *testing.T) {
	//var idStr string = strconv.FormatInt(rrID, 10)
	var CIDStr string = strconv.FormatInt(rrHandid, 10)
	r, _ := http.NewRequest("DELETE", "/test?id=e"+"&clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "26")
	r.Header.Set("clientId", "26")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrr.HandleRestRouteSuperDelete(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.RestRoute
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get: ")
	fmt.Println(bdy)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestRRt_HandleDelete(t *testing.T) {
	var idStr string = strconv.FormatInt(rrID, 10)
	var CIDStr string = strconv.FormatInt(rrHandid, 10)
	r, _ := http.NewRequest("DELETE", "/test?id="+idStr+"&clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "26")
	r.Header.Set("clientId", "26")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrr.HandleRestRouteSuperDelete(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.Success != true {
		t.Fail()
	}
}

func TestRRt_HandleGet2(t *testing.T) {
	var idStr string = strconv.FormatInt(rrID, 10)
	var CIDStr string = strconv.FormatInt(rrHandid, 10)
	r, _ := http.NewRequest("GET", "/test?id="+idStr+"&clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "26")
	r.Header.Set("clientId", "26")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrr.HandleRestRouteSuperGet(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.RestRoute
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get after delete: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.ClientID != 0 || bdy.ID != 0 {
		t.Fail()
	}
}

func TestRRt_HandleDel(t *testing.T) {
	//var routeBkStr string = strconv.FormatInt(routeBk, 10)
	var CIDStr string = strconv.FormatInt(rrHandid, 10)
	r, _ := http.NewRequest("DELETE", "/test?clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "24")
	r.Header.Set("clientId", "24")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrr.HandleClientDelete(w, r)
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

func TestRRt_TestCloseDb(t *testing.T) {
	success := gwRR.DbConfig.CloseDb()
	if success != true {
		t.Fail()
	}
}
