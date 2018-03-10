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

var rruHandidNs int64 = 39
var rruIDNs int64
var rruuIDNs int64
var rruuID2Ns int64

//var routeErr int64
//var routeURLErrID int64
var connectedForRRuNs bool
var gwRRuNs mgr.GatewayDB

//var gwRoutesErr mgr.GatewayRoutes
var hrruNs Handler

func TestRRtuNs1_Connect(t *testing.T) {
	gwRRuNs.DbConfig.Host = "localhost:3306"
	gwRRuNs.DbConfig.DbUser = "admin"
	gwRRuNs.DbConfig.DbPw = "admin"
	gwRRuNs.DbConfig.DatabaseName = "ulbora_api_gateway"
	connectedForRRuNs = gwRRuNs.ConnectDb()
	if connectedForRRuNs != true {
		t.Fail()
	}
	//gwRoutesErr.GwDB.DbConfig = edb.DbConfig
	//gwRoutes.GwDB.DbConfig = gwRoutes.GwDB.DbConfig
	//cp.Host = "http://localhost:3010"
	testMode = true
	hrruNs.DbConfig = gwRRuNs.DbConfig
}

func TestRRtuNs1_HandleClientInsert(t *testing.T) {
	var c mgr.Client
	c.ClientID = rruHandidNs
	c.APIKey = "123456"
	c.Enabled = true
	c.Level = "small"
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrruNs.HandleClientPost(w, r)
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

func TestRRtuNs1_HandleRestRouteInsert(t *testing.T) {
	var rr mgr.RestRoute
	rr.ClientID = rruHandidNs
	rr.Route = "test"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrruNs.HandleRestRouteSuperPost(w, r)
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
		rruIDNs = bdy.ID
	}
}

func TestRRtuNs1_HandleRouteUrlInsertMedia(t *testing.T) {
	var rr mgr.RouteURL
	//rr.ClientID = rruHandid
	rr.Name = "blue"
	rr.RouteID = rruIDNs
	rr.URL = "test/url"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLPost(w, r)
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

func TestRRtuNs1_HandleRouteUrlInsertReq(t *testing.T) {
	var rr mgr.RouteURL
	//rr.ClientID = rruHandid
	//rr.Name = "blue"
	//rr.RouteID = rruID
	rr.URL = "test/url"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLPost(w, r)
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

func TestRRtuNs1_HandleRouteUrlInsertMethod(t *testing.T) {
	var rr mgr.RouteURL
	//rr.ClientID = rruHandid
	rr.Name = "blue"
	rr.RouteID = rruIDNs
	rr.URL = "test/url"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLPost(w, r)
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

func TestRRtuNs1_HandleRouteUrlInsert(t *testing.T) {
	var rr mgr.RouteURL
	//rr.ClientID = rruHandid
	rr.Name = "blue"
	rr.RouteID = rruIDNs
	rr.URL = "test/url"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLPost(w, r)
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
		rruuIDNs = bdy.ID
	}
}

func TestRRtuNs1_HandleRouteUrlMedia(t *testing.T) {
	var rr mgr.RouteURL
	rr.ID = rruuIDNs
	//rr.ClientID = rruHandid
	rr.Name = "green"
	rr.RouteID = rruIDNs
	rr.URL = "test/url"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLPut(w, r)
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

func TestRRtuNs1_HandleRouteUrlUpdateReq(t *testing.T) {
	var rr mgr.RouteURL
	//rr.ID = rruuID
	//rr.ClientID = rruHandidNs
	rr.Name = "green"
	rr.RouteID = rruIDNs
	rr.URL = "test/url"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLPut(w, r)
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

func TestRRtuNs1_HandleRouteUrlUpdateMethod(t *testing.T) {
	var rr mgr.RouteURL
	rr.ID = rruuIDNs
	//rr.ClientID = rruHandid
	rr.Name = "green"
	rr.RouteID = rruIDNs
	rr.URL = "test/url"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLPut(w, r)
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

func TestRRtuNs1_HandleRouteUrlUpdate(t *testing.T) {
	var rr mgr.RouteURL
	rr.ID = rruuIDNs
	//rr.ClientID = rruHandid
	rr.Name = "green"
	rr.RouteID = rruIDNs
	rr.URL = "test/url"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLPut(w, r)
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

func TestRRtuNs1_HandleGetReq(t *testing.T) {
	var idStr string = strconv.FormatInt(rruuIDNs, 10)
	//var routeIDStr string = strconv.FormatInt(rruID, 10)
	//var CIDStr string = strconv.FormatInt(rruHandidNs, 10)
	r, _ := http.NewRequest("GET", "/test?id="+idStr+"&routeId=g", nil)
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLGet(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.RouteURL
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get URL: ")
	fmt.Println(bdy)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestRRtuNs1_HandleGetMethod(t *testing.T) {
	var idStr string = strconv.FormatInt(rruuIDNs, 10)
	var routeIDStr string = strconv.FormatInt(rruIDNs, 10)
	//var CIDStr string = strconv.FormatInt(rruHandid, 10)
	r, _ := http.NewRequest("DELETE", "/test?id="+idStr+"&routeId="+routeIDStr, nil)
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLGet(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.RouteURL
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get URL: ")
	fmt.Println(bdy)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestRRtuNs1_HandleGet(t *testing.T) {
	var idStr string = strconv.FormatInt(rruuIDNs, 10)
	var routeIDStr string = strconv.FormatInt(rruIDNs, 10)
	//var CIDStr string = strconv.FormatInt(rruHandid, 10)
	r, _ := http.NewRequest("GET", "/test?id="+idStr+"&routeId="+routeIDStr, nil)
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLGet(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.RouteURL
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get URL: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.ClientID != rruHandidNs || bdy.ID != rruuIDNs || bdy.Active != false {
		t.Fail()
	}
}

func TestRRtuNs1_HandleRouteUrlActivateMedia(t *testing.T) {
	var rr mgr.RouteURL
	rr.ID = rruuIDNs
	rr.RouteID = rruIDNs
	//rr.ClientID = rruHandid
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLActivate(w, r)
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

func TestRRtuNs1_HandleRouteUrlActivateReq(t *testing.T) {
	var rr mgr.RouteURL
	//rr.ID = rruuID
	rr.RouteID = rruIDNs
	//rr.ClientID = rruHandidNs
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLActivate(w, r)
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

func TestRRtuNs1_HandleRouteUrlActivateMethod(t *testing.T) {
	var rr mgr.RouteURL
	rr.ID = rruuIDNs
	rr.RouteID = rruIDNs
	//rr.ClientID = rruHandidNs
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLActivate(w, r)
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

func TestRRtuNs1_HandleRouteUrlActivate(t *testing.T) {
	var rr mgr.RouteURL
	rr.ID = rruuIDNs
	rr.RouteID = rruIDNs
	//rr.ClientID = rruHandidNs
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLActivate(w, r)
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

func TestRRtuNs1_HandleGet3(t *testing.T) {
	var idStr string = strconv.FormatInt(rruuIDNs, 10)
	var routeIDStr string = strconv.FormatInt(rruIDNs, 10)
	//var CIDStr string = strconv.FormatInt(rruHandid, 10)
	r, _ := http.NewRequest("GET", "/test?id="+idStr+"&routeId="+routeIDStr, nil)
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLGet(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.RouteURL
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get URL: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.ClientID != rruHandidNs || bdy.ID != rruuIDNs || bdy.Active != true {
		t.Fail()
	}
}

func TestRRtuNs1_HandleGetListReq(t *testing.T) {
	//var routeIDStr string = strconv.FormatInt(rruID, 10)
	//var CIDStr string = strconv.FormatInt(rruHandid, 10)
	r, _ := http.NewRequest("GET", "/test?routeId=e", nil)
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLList(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy []mgr.RouteURL
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get URL List: ")
	fmt.Println(bdy)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestRRtuNs1_HandleGetListMethod(t *testing.T) {
	var routeIDStr string = strconv.FormatInt(rruIDNs, 10)
	//var CIDStr string = strconv.FormatInt(rruHandid, 10)
	r, _ := http.NewRequest("DELETE", "/test?routeId="+routeIDStr, nil)
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLList(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy []mgr.RouteURL
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get URL List: ")
	fmt.Println(bdy)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestRRtuNs1_HandleGetList(t *testing.T) {
	var routeIDStr string = strconv.FormatInt(rruIDNs, 10)
	//var CIDStr string = strconv.FormatInt(rruHandid, 10)
	r, _ := http.NewRequest("GET", "/test?routeId="+routeIDStr, nil)
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLList(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy []mgr.RouteURL
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get URL List: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || len(bdy) != 1 || bdy[0].Active != true {
		t.Fail()
	}
}

func TestRRtuNs1_HandleDelete(t *testing.T) {
	var idStr string = strconv.FormatInt(rruuIDNs, 10)
	var routeIDStr string = strconv.FormatInt(rruIDNs, 10)
	//var CIDStr string = strconv.FormatInt(rruHandid, 10)
	r, _ := http.NewRequest("DELETE", "/test?id="+idStr+"&routeId="+routeIDStr, nil)
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLDelete(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Delete URL: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.Success != false {
		t.Fail()
	}
}

func TestRRtuNs1_HandleRouteUrlInsert2(t *testing.T) {
	var rr mgr.RouteURL
	//rr.ClientID = rruHandidNs
	rr.Name = "blue"
	rr.RouteID = rruIDNs
	rr.URL = "test/url"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLPost(w, r)
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
		rruuID2Ns = bdy.ID
	}
}

func TestRRtuNs1_HandleRouteUrlActivate2(t *testing.T) {
	var rr mgr.RouteURL
	rr.ID = rruuID2Ns
	rr.RouteID = rruIDNs
	//rr.ClientID = rruHandid
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLActivate(w, r)
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

func TestRRtuNs1_HandleDelete2Req(t *testing.T) {
	var idStr string = strconv.FormatInt(rruuIDNs, 10)
	//var routeIDStr string = strconv.FormatInt(rruID, 10)
	//var CIDStr string = strconv.FormatInt(rruHandid, 10)
	r, _ := http.NewRequest("DELETE", "/test?id="+idStr+"&routeId=w", nil)
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLDelete(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Delete URL: ")
	fmt.Println(bdy)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestRRtuNs1_HandleDelete2Method(t *testing.T) {
	var idStr string = strconv.FormatInt(rruuIDNs, 10)
	var routeIDStr string = strconv.FormatInt(rruIDNs, 10)
	//var CIDStr string = strconv.FormatInt(rruHandid, 10)
	r, _ := http.NewRequest("GET", "/test?id="+idStr+"&routeId="+routeIDStr, nil)
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLDelete(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Delete URL: ")
	fmt.Println(bdy)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestRRtuNs1_HandleDelete2(t *testing.T) {
	var idStr string = strconv.FormatInt(rruuIDNs, 10)
	var routeIDStr string = strconv.FormatInt(rruIDNs, 10)
	//var CIDStr string = strconv.FormatInt(rruHandid, 10)
	r, _ := http.NewRequest("DELETE", "/test?id="+idStr+"&routeId="+routeIDStr, nil)
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLDelete(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Delete URL: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.Success != true {
		t.Fail()
	}
}

func TestRRtuNs1_HandleDel(t *testing.T) {
	//var routeBkStr string = strconv.FormatInt(routeBk, 10)
	var CIDStr string = strconv.FormatInt(rruHandidNs, 10)
	r, _ := http.NewRequest("DELETE", "/test?clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "39")
	r.Header.Set("clientId", "39")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleClientDelete(w, r)
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

func TestRRtuNs1_TestCloseDb(t *testing.T) {
	success := gwRRu.DbConfig.CloseDb()
	if success != true {
		t.Fail()
	}
}
