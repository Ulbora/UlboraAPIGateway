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

var rruHandid int64 = 38
var rruID int64
var rruuID int64
var rruuID2 int64

//var routeErr int64
//var routeURLErrID int64
var connectedForRRu bool
var gwRRu mgr.GatewayDB

//var gwRoutesErr mgr.GatewayRoutes
var hrru Handler

func TestRRtu1_Connect(t *testing.T) {
	gwRRu.DbConfig.Host = "localhost:3306"
	gwRRu.DbConfig.DbUser = "admin"
	gwRRu.DbConfig.DbPw = "admin"
	gwRRu.DbConfig.DatabaseName = "ulbora_api_gateway"
	connectedForRRu = gwRRu.ConnectDb()
	if connectedForRRu != true {
		t.Fail()
	}
	//gwRoutesErr.GwDB.DbConfig = edb.DbConfig
	//gwRoutes.GwDB.DbConfig = gwRoutes.GwDB.DbConfig
	//cp.Host = "http://localhost:3010"
	testMode = true
	hrru.DbConfig = gwRRu.DbConfig
}

func TestRRtu1_HandleClientInsert(t *testing.T) {
	var c mgr.Client
	c.ClientID = rruHandid
	c.APIKey = "123456"
	c.Enabled = true
	c.Level = "small"
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleClientPost(w, r)
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

func TestRRtu1_HandleRestRouteInsert(t *testing.T) {
	var rr mgr.RestRoute
	rr.ClientID = rruHandid
	rr.Route = "test"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRestRouteSuperPost(w, r)
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
		rruID = bdy.ID
	}
}

func TestRRtu1_HandleRouteUrlInsertMedia(t *testing.T) {
	var rr mgr.RouteURL
	rr.ClientID = rruHandid
	rr.Name = "blue"
	rr.RouteID = rruID
	rr.URL = "test/url"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLSuperPost(w, r)
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

func TestRRtu1_HandleRouteUrlInsertReq(t *testing.T) {
	var rr mgr.RouteURL
	rr.ClientID = rruHandid
	rr.Name = "blue"
	//rr.RouteID = rruID
	rr.URL = "test/url"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLSuperPost(w, r)
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

func TestRRtu1_HandleRouteUrlInsertMethod(t *testing.T) {
	var rr mgr.RouteURL
	rr.ClientID = rruHandid
	rr.Name = "blue"
	rr.RouteID = rruID
	rr.URL = "test/url"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLSuperPost(w, r)
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

func TestRRtu1_HandleRouteUrlInsertAuth(t *testing.T) {
	testMode = false
	var rr mgr.RouteURL
	rr.ClientID = rruHandid
	rr.Name = "blue"
	rr.RouteID = rruID
	rr.URL = "test/url"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLSuperPost(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusUnauthorized {
		t.Fail()
	}
	testMode = true
}

func TestRRtu1_HandleRouteUrlInsert(t *testing.T) {
	var rr mgr.RouteURL
	rr.ClientID = rruHandid
	rr.Name = "blue"
	rr.RouteID = rruID
	rr.URL = "test/url"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLSuperPost(w, r)
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
		rruuID = bdy.ID
	}
}

func TestRRtu1_HandleRouteUrlMedia(t *testing.T) {
	var rr mgr.RouteURL
	rr.ID = rruuID
	rr.ClientID = rruHandid
	rr.Name = "green"
	rr.RouteID = rruID
	rr.URL = "test/url"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLSuperPut(w, r)
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

func TestRRtu1_HandleRouteUrlUpdateReq(t *testing.T) {
	var rr mgr.RouteURL
	//rr.ID = rruuID
	rr.ClientID = rruHandid
	rr.Name = "green"
	rr.RouteID = rruID
	rr.URL = "test/url"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLSuperPut(w, r)
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

func TestRRtu1_HandleRouteUrlUpdateMethod(t *testing.T) {
	var rr mgr.RouteURL
	rr.ID = rruuID
	rr.ClientID = rruHandid
	rr.Name = "green"
	rr.RouteID = rruID
	rr.URL = "test/url"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLSuperPut(w, r)
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

func TestRRtu1_HandleRouteUrlUpdateAuth(t *testing.T) {
	testMode = false
	var rr mgr.RouteURL
	rr.ID = rruuID
	rr.ClientID = rruHandid
	rr.Name = "green"
	rr.RouteID = rruID
	rr.URL = "test/url"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLSuperPut(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusUnauthorized {
		t.Fail()
	}
	testMode = true
}

func TestRRtu1_HandleRouteUrlUpdate(t *testing.T) {
	var rr mgr.RouteURL
	rr.ID = rruuID
	rr.ClientID = rruHandid
	rr.Name = "green"
	rr.RouteID = rruID
	rr.URL = "test/url"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLSuperPut(w, r)
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

func TestRRtu1_HandleGetReq(t *testing.T) {
	var idStr string = strconv.FormatInt(rruuID, 10)
	//var routeIDStr string = strconv.FormatInt(rruID, 10)
	var CIDStr string = strconv.FormatInt(rruHandid, 10)
	r, _ := http.NewRequest("GET", "/test?id="+idStr+"&routeId=g&clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "38")
	r.Header.Set("clientId", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLSuperGet(w, r)
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

func TestRRtu1_HandleGetMethod(t *testing.T) {
	var idStr string = strconv.FormatInt(rruuID, 10)
	var routeIDStr string = strconv.FormatInt(rruID, 10)
	var CIDStr string = strconv.FormatInt(rruHandid, 10)
	r, _ := http.NewRequest("DELETE", "/test?id="+idStr+"&routeId="+routeIDStr+"&clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "38")
	r.Header.Set("clientId", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLSuperGet(w, r)
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

func TestRRtu1_HandleGetAuth(t *testing.T) {
	testMode = false
	var idStr string = strconv.FormatInt(rruuID, 10)
	var routeIDStr string = strconv.FormatInt(rruID, 10)
	var CIDStr string = strconv.FormatInt(rruHandid, 10)
	r, _ := http.NewRequest("GET", "/test?id="+idStr+"&routeId="+routeIDStr+"&clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "38")
	r.Header.Set("clientId", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLSuperGet(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.RouteURL
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get URL: ")
	fmt.Println(bdy)
	if w.Code != http.StatusUnauthorized {
		t.Fail()
	}
	testMode = true
}

func TestRRtu1_HandleGet(t *testing.T) {
	var idStr string = strconv.FormatInt(rruuID, 10)
	var routeIDStr string = strconv.FormatInt(rruID, 10)
	var CIDStr string = strconv.FormatInt(rruHandid, 10)
	r, _ := http.NewRequest("GET", "/test?id="+idStr+"&routeId="+routeIDStr+"&clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "38")
	r.Header.Set("clientId", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLSuperGet(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.RouteURL
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get URL: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.ClientID != rruHandid || bdy.ID != rruuID || bdy.Active != false {
		t.Fail()
	}
}

func TestRRtu1_HandleRouteUrlActivateMedia(t *testing.T) {
	var rr mgr.RouteURL
	rr.ID = rruuID
	rr.RouteID = rruID
	rr.ClientID = rruHandid
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLActivateSuper(w, r)
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

func TestRRtu1_HandleRouteUrlActivateReq(t *testing.T) {
	var rr mgr.RouteURL
	//rr.ID = rruuID
	rr.RouteID = rruID
	rr.ClientID = rruHandid
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLActivateSuper(w, r)
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

func TestRRtu1_HandleRouteUrlActivateMethod(t *testing.T) {
	var rr mgr.RouteURL
	rr.ID = rruuID
	rr.RouteID = rruID
	rr.ClientID = rruHandid
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLActivateSuper(w, r)
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

func TestRRtu1_HandleRouteUrlActivateAuth(t *testing.T) {
	testMode = false
	var rr mgr.RouteURL
	rr.ID = rruuID
	rr.RouteID = rruID
	rr.ClientID = rruHandid
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLActivateSuper(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp: ")
	fmt.Println(bdy)
	if w.Code != http.StatusUnauthorized {
		t.Fail()
	}
	testMode = true
}

func TestRRtu1_HandleRouteUrlActivate(t *testing.T) {
	var rr mgr.RouteURL
	rr.ID = rruuID
	rr.RouteID = rruID
	rr.ClientID = rruHandid
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLActivateSuper(w, r)
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

func TestRRtu_HandleGet3(t *testing.T) {
	var idStr string = strconv.FormatInt(rruuID, 10)
	var routeIDStr string = strconv.FormatInt(rruID, 10)
	var CIDStr string = strconv.FormatInt(rruHandid, 10)
	r, _ := http.NewRequest("GET", "/test?id="+idStr+"&routeId="+routeIDStr+"&clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "38")
	r.Header.Set("clientId", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLSuperGet(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.RouteURL
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get URL: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.ClientID != rruHandid || bdy.ID != rruuID || bdy.Active != true {
		t.Fail()
	}
}

func TestRRtu1_HandleGetListReq(t *testing.T) {
	//var routeIDStr string = strconv.FormatInt(rruID, 10)
	var CIDStr string = strconv.FormatInt(rruHandid, 10)
	r, _ := http.NewRequest("GET", "/test?routeId=e&clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "38")
	r.Header.Set("clientId", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLSuperList(w, r)
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

func TestRRtu1_HandleGetListMethod(t *testing.T) {
	var routeIDStr string = strconv.FormatInt(rruID, 10)
	var CIDStr string = strconv.FormatInt(rruHandid, 10)
	r, _ := http.NewRequest("DELETE", "/test?routeId="+routeIDStr+"&clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "38")
	r.Header.Set("clientId", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLSuperList(w, r)
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

func TestRRtu1_HandleGetListAuth(t *testing.T) {
	testMode = false
	var routeIDStr string = strconv.FormatInt(rruID, 10)
	var CIDStr string = strconv.FormatInt(rruHandid, 10)
	r, _ := http.NewRequest("GET", "/test?routeId="+routeIDStr+"&clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "38")
	r.Header.Set("clientId", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLSuperList(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy []mgr.RouteURL
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get URL List: ")
	fmt.Println(bdy)
	if w.Code != http.StatusUnauthorized {
		t.Fail()
	}
	testMode = true
}

func TestRRtu1_HandleGetList(t *testing.T) {
	var routeIDStr string = strconv.FormatInt(rruID, 10)
	var CIDStr string = strconv.FormatInt(rruHandid, 10)
	r, _ := http.NewRequest("GET", "/test?routeId="+routeIDStr+"&clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "38")
	r.Header.Set("clientId", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLSuperList(w, r)
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

func TestRRtu1_HandleDeleteAuth(t *testing.T) {
	testMode = false
	var idStr string = strconv.FormatInt(rruuID, 10)
	var routeIDStr string = strconv.FormatInt(rruID, 10)
	var CIDStr string = strconv.FormatInt(rruHandid, 10)
	r, _ := http.NewRequest("DELETE", "/test?id="+idStr+"&routeId="+routeIDStr+"&clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "38")
	r.Header.Set("clientId", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLSuperDelete(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.GatewayResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Delete URL: ")
	fmt.Println(bdy)
	if w.Code != http.StatusUnauthorized {
		t.Fail()
	}
	testMode = true
}

func TestRRtu1_HandleDelete(t *testing.T) {
	var idStr string = strconv.FormatInt(rruuID, 10)
	var routeIDStr string = strconv.FormatInt(rruID, 10)
	var CIDStr string = strconv.FormatInt(rruHandid, 10)
	r, _ := http.NewRequest("DELETE", "/test?id="+idStr+"&routeId="+routeIDStr+"&clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "38")
	r.Header.Set("clientId", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLSuperDelete(w, r)
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

func TestRRtu1_HandleRouteUrlInsert2(t *testing.T) {
	var rr mgr.RouteURL
	rr.ClientID = rruHandid
	rr.Name = "blue"
	rr.RouteID = rruID
	rr.URL = "test/url"
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLSuperPost(w, r)
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
		rruuID2 = bdy.ID
	}
}

func TestRRtu1_HandleRouteUrlActivate2(t *testing.T) {
	var rr mgr.RouteURL
	rr.ID = rruuID2
	rr.RouteID = rruID
	rr.ClientID = rruHandid
	aJSON, _ := json.Marshal(rr)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLActivateSuper(w, r)
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

func TestRRtu1_HandleDelete2Req(t *testing.T) {
	var idStr string = strconv.FormatInt(rruuID, 10)
	//var routeIDStr string = strconv.FormatInt(rruID, 10)
	var CIDStr string = strconv.FormatInt(rruHandid, 10)
	r, _ := http.NewRequest("DELETE", "/test?id="+idStr+"&routeId=w&clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "38")
	r.Header.Set("clientId", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLSuperDelete(w, r)
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

func TestRRtu1_HandleDelete2Method(t *testing.T) {
	var idStr string = strconv.FormatInt(rruuID, 10)
	var routeIDStr string = strconv.FormatInt(rruID, 10)
	var CIDStr string = strconv.FormatInt(rruHandid, 10)
	r, _ := http.NewRequest("GET", "/test?id="+idStr+"&routeId="+routeIDStr+"&clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "38")
	r.Header.Set("clientId", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLSuperDelete(w, r)
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

func TestRRtu1_HandleDelete2(t *testing.T) {
	var idStr string = strconv.FormatInt(rruuID, 10)
	var routeIDStr string = strconv.FormatInt(rruID, 10)
	var CIDStr string = strconv.FormatInt(rruHandid, 10)
	r, _ := http.NewRequest("DELETE", "/test?id="+idStr+"&routeId="+routeIDStr+"&clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "38")
	r.Header.Set("clientId", "38")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hrru.HandleRouteURLSuperDelete(w, r)
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

func TestRRtu1_HandleDel(t *testing.T) {
	//var routeBkStr string = strconv.FormatInt(routeBk, 10)
	var CIDStr string = strconv.FormatInt(rruHandid, 10)
	r, _ := http.NewRequest("DELETE", "/test?clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "24")
	r.Header.Set("clientId", "24")
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

func TestRRtu1_TestCloseDb(t *testing.T) {
	success := gwRRu.DbConfig.CloseDb()
	if success != true {
		t.Fail()
	}
}
