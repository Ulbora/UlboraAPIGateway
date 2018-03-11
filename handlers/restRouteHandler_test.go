package handlers

import (
	mgr "UlboraApiGateway/managers"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	//"time"
)

var rrNsHandid int64 = 36
var rrNsID int64

//var routeErr int64
//var routeURLErrID int64
var connectedForRRNs bool
var gwRRNs mgr.GatewayDB

//var gwRoutesErr mgr.GatewayRoutes
var hrrNs Handler

func TestRRtNs_Connect(t *testing.T) {
	gwRRNs.DbConfig.Host = "localhost:3306"
	gwRRNs.DbConfig.DbUser = "admin"
	gwRRNs.DbConfig.DbPw = "admin"
	gwRRNs.DbConfig.DatabaseName = "ulbora_api_gateway"
	connectedForRRNs = gwRRNs.ConnectDb()
	if connectedForRRNs != true {
		t.Fail()
	}
	//gwRoutesErr.GwDB.DbConfig = edb.DbConfig
	//gwRoutes.GwDB.DbConfig = gwRoutes.GwDB.DbConfig
	//cp.Host = "http://localhost:3010"
	testMode = true
	hrrNs.DbConfig = gwRRNs.DbConfig
}

func TestRRtNs_HandleClientInsert(t *testing.T) {
	var c mgr.Client
	c.ClientID = rrNsHandid
	c.APIKey = "123456"
	c.Enabled = true
	c.Level = "small"
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "36")
	r.Header.Set("clientId", "36")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrrNs.HandleClientPost(w, r)
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

func TestRRtNs_HandleRestRouteInsertMedia(t *testing.T) {
	var rr mgr.RestRoute
	//rr.ClientID = rrNsHandid
	rr.Route = "test"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "36")
	r.Header.Set("clientId", "36")
	r.Header.Set("u-api-key", "12233hgdd3335")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrrNs.HandleRestRoutePost(w, r)
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

func TestRRtNs_HandleRestRouteInsertMethod(t *testing.T) {
	var rr mgr.RestRoute
	//rr.ClientID = rrNsHandid
	rr.Route = "test"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("GET", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "36")
	r.Header.Set("clientId", "36")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrrNs.HandleRestRoutePost(w, r)
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

func TestRRtNs_HandleRestRouteInsertReq(t *testing.T) {
	var rr mgr.RestRoute
	//rr.Route = "test"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "36")
	//r.Header.Set("clientId", "36")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrrNs.HandleRestRoutePost(w, r)
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

func TestRRtNs_HandleRestRouteInsert(t *testing.T) {
	var rr mgr.RestRoute
	//rr.ClientID = rrHandid
	rr.Route = "test"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "36")
	r.Header.Set("clientId", "36")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrrNs.HandleRestRoutePost(w, r)
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
		rrNsID = bdy.ID
	}
}

func TestRRtNs_HandleRestRouteUpdateMedia(t *testing.T) {
	var rr mgr.RestRoute
	//rr.ClientID = rrHandid
	rr.Route = "test"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "36")
	r.Header.Set("clientId", "36")
	r.Header.Set("u-api-key", "12233hgdd3335")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrrNs.HandleRestRoutePut(w, r)
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

func TestRRtNs_HandleRestRouteUpdateMethod(t *testing.T) {
	var rr mgr.RestRoute
	//rr.ClientID = rrHandid
	rr.Route = "test"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("GET", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "36")
	r.Header.Set("clientId", "36")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrrNs.HandleRestRoutePut(w, r)
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

func TestRRtNs_HandleRestRouteUpdateReq(t *testing.T) {
	var rr mgr.RestRoute
	//rr.ClientID = rrHandid
	//rr.Route = "test"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "36")
	r.Header.Set("clientId", "36")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrrNs.HandleRestRoutePut(w, r)
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

func TestRRtNs_HandleRestRouteUpdate(t *testing.T) {
	var rr mgr.RestRoute
	rr.ID = rrNsID
	rr.Route = "test2"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "36")
	r.Header.Set("clientId", "36")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrrNs.HandleRestRoutePut(w, r)
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

func TestRRtNs_HandleGetMethod(t *testing.T) {
	var idStr string = strconv.FormatInt(rrNsID, 10)
	r, _ := http.NewRequest("POST", "/test?id="+idStr, nil)
	r.Header.Set("u-client-id", "36")
	r.Header.Set("clientId", "36")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrrNs.HandleRestRouteGet(w, r)
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

func TestRRtNs_HandleGetReq(t *testing.T) {
	//var idStr string = strconv.FormatInt(rrID, 10)
	r, _ := http.NewRequest("GET", "/test?id=w", nil)
	r.Header.Set("u-client-id", "36")
	r.Header.Set("clientId", "36")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrrNs.HandleRestRouteGet(w, r)
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

func TestRRtNs_HandleGet(t *testing.T) {
	var idStr string = strconv.FormatInt(rrNsID, 10)
	r, _ := http.NewRequest("GET", "/test?id="+idStr, nil)
	r.Header.Set("u-client-id", "36")
	r.Header.Set("clientId", "36")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrrNs.HandleRestRouteGet(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.RestRoute
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.ClientID != rrNsHandid || bdy.ID != rrNsID {
		t.Fail()
	}
}

func TestRRtNs_HandleGetListMethod(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "/test", nil)
	r.Header.Set("u-client-id", "36")
	r.Header.Set("clientId", "36")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrrNs.HandleRestRouteList(w, r)
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

func TestRRtNs_HandleGetList(t *testing.T) {
	r, _ := http.NewRequest("GET", "/test", nil)
	r.Header.Set("u-client-id", "36")
	r.Header.Set("clientId", "36")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrrNs.HandleRestRouteList(w, r)
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

func TestRRtNs_HandleDeleteMethod(t *testing.T) {
	var idStr string = strconv.FormatInt(rrNsID, 10)
	r, _ := http.NewRequest("POST", "/test?id="+idStr, nil)
	r.Header.Set("u-client-id", "36")
	r.Header.Set("clientId", "36")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrrNs.HandleRestRouteDelete(w, r)
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

func TestRRtNs_HandleDeleteReq(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "/test?id=e", nil)
	r.Header.Set("u-client-id", "36")
	r.Header.Set("clientId", "36")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrrNs.HandleRestRouteDelete(w, r)
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

func TestRRtNs_HandleDelete(t *testing.T) {
	var idStr string = strconv.FormatInt(rrNsID, 10)
	r, _ := http.NewRequest("DELETE", "/test?id="+idStr, nil)
	r.Header.Set("u-client-id", "36")
	r.Header.Set("clientId", "36")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrrNs.HandleRestRouteDelete(w, r)
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

func TestRRtNs_HandleGet2(t *testing.T) {
	var idStr string = strconv.FormatInt(rrNsID, 10)
	r, _ := http.NewRequest("GET", "/test?id="+idStr, nil)
	r.Header.Set("u-client-id", "36")
	r.Header.Set("clientId", "36")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrrNs.HandleRestRouteGet(w, r)
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

func TestRRtNs_HandleDel(t *testing.T) {
	//var routeBkStr string = strconv.FormatInt(routeBk, 10)
	var CIDStr string = strconv.FormatInt(rrNsHandid, 10)
	r, _ := http.NewRequest("DELETE", "/test?clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "36")
	r.Header.Set("clientId", "36")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrrNs.HandleClientDelete(w, r)
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

func TestRRtNs_TestCloseDb(t *testing.T) {
	success := gwRR.DbConfig.CloseDb()
	if success != true {
		t.Fail()
	}
}
