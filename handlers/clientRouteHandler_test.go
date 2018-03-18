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

var clientHandid int64 = 24

//var routeErr int64
//var routeURLErrID int64
var connectedForCh bool
var gwCh mgr.GatewayDB

//var gwRoutesErr mgr.GatewayRoutes
var hch Handler

func TestCliRt_ConnectForErr(t *testing.T) {
	gwCh.DbConfig.Host = "localhost:3306"
	gwCh.DbConfig.DbUser = "admin"
	gwCh.DbConfig.DbPw = "admin"
	gwCh.DbConfig.DatabaseName = "ulbora_api_gateway"
	connectedForCh = gwCh.ConnectDb()
	if connectedForCh != true {
		t.Fail()
	}
	//gwRoutesErr.GwDB.DbConfig = edb.DbConfig
	//gwRoutes.GwDB.DbConfig = gwRoutes.GwDB.DbConfig
	//cp.Host = "http://localhost:3010"
	testMode = true
	hch.DbConfig = gwCh.DbConfig
}

func TestCliRt_HandleClientInsertMedia(t *testing.T) {
	var c mgr.Client
	c.ClientID = clientHandid
	c.APIKey = "123456"
	c.Enabled = true
	c.Level = "small"
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "24")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hch.HandleClientPost(w, r)
	fmt.Print("Media Code: ")
	fmt.Println(w.Code)
	if w.Code != http.StatusUnsupportedMediaType {
		t.Fail()
	}
}

func TestCliRt_HandleClientInsertReq(t *testing.T) {
	var c mgr.Client
	//c.ClientID = clientHandid
	c.APIKey = "123456"
	c.Enabled = true
	c.Level = "small"
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "24")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hch.HandleClientPost(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestCliRt_HandleClientInsertMethod(t *testing.T) {
	var c mgr.Client
	c.ClientID = clientHandid
	c.APIKey = "123456"
	c.Enabled = true
	c.Level = "small"
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "24")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hch.HandleClientPost(w, r)
	fmt.Print(" Code: ")
	fmt.Println(w.Code)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestCliRt_HandleClientInsertAuth(t *testing.T) {
	testMode = false
	var c mgr.Client
	c.ClientID = clientHandid
	c.APIKey = "123456"
	c.Enabled = true
	c.Level = "small"
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "24")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hch.HandleClientPost(w, r)
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

func TestCliRt_HandleClientInsert(t *testing.T) {
	var c mgr.Client
	c.ClientID = clientHandid
	c.APIKey = "123456"
	c.Enabled = true
	c.Level = "small"
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "24")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hch.HandleClientPost(w, r)
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

func TestCliRt_HandleClientUpdateMedia(t *testing.T) {
	var c mgr.Client
	c.ClientID = clientHandid
	c.APIKey = "123456"
	c.Enabled = false
	c.Level = "small"
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "24")
	r.Header.Set("u-api-key", "12233hgdd3335")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hch.HandleClientPut(w, r)
	fmt.Print("Update Code: ")
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

func TestCliRt_HandleClientUpdateReq(t *testing.T) {
	var c mgr.Client
	//c.ClientID = clientHandid
	c.APIKey = "123456"
	c.Enabled = false
	c.Level = "small"
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "24")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hch.HandleClientPut(w, r)
	fmt.Print("Update Code: ")
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

func TestCliRt_HandleClientUpdateMethod(t *testing.T) {
	var c mgr.Client
	c.ClientID = clientHandid
	c.APIKey = "123456"
	c.Enabled = false
	c.Level = "small"
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "24")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hch.HandleClientPut(w, r)
	fmt.Print("Update Code: ")
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

func TestCliRt_HandleClientUpdateAuth(t *testing.T) {
	testMode = false
	var c mgr.Client
	c.ClientID = clientHandid
	c.APIKey = "123456"
	c.Enabled = false
	c.Level = "small"
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "24")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hch.HandleClientPut(w, r)
	fmt.Print("Update Code: ")
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

func TestCliRt_HandleClientUpdate(t *testing.T) {
	var c mgr.Client
	c.ClientID = clientHandid
	c.APIKey = "123456"
	c.Enabled = false
	c.Level = "small"
	aJSON, _ := json.Marshal(c)
	r, _ := http.NewRequest("PUT", "/test", bytes.NewBuffer(aJSON))
	r.Header.Set("u-client-id", "24")
	r.Header.Set("u-api-key", "12233hgdd3335")
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	hch.HandleClientPut(w, r)
	fmt.Print("Update Code: ")
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

func TestCliRt_HandleGetReq(t *testing.T) {
	//var routeBkStr string = strconv.FormatInt(routeBk, 10)
	//var CIDStr string = strconv.FormatInt(clientHandid, 10)
	r, _ := http.NewRequest("GET", "/test?clientId=e", nil)
	r.Header.Set("u-client-id", "24")
	r.Header.Set("clientId", "24")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBk.HandleClientGet(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.Client
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get: ")
	fmt.Println(bdy)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestCliRt_HandleGetMethod(t *testing.T) {
	//var routeBkStr string = strconv.FormatInt(routeBk, 10)
	var CIDStr string = strconv.FormatInt(clientHandid, 10)
	r, _ := http.NewRequest("DELETE", "/test?clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "24")
	r.Header.Set("clientId", "24")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBk.HandleClientGet(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.Client
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get: ")
	fmt.Println(bdy)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestCliRt_HandleGetAuth(t *testing.T) {
	//var routeBkStr string = strconv.FormatInt(routeBk, 10)
	testMode = false
	var CIDStr string = strconv.FormatInt(clientHandid, 10)
	r, _ := http.NewRequest("GET", "/test?clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "24")
	r.Header.Set("clientId", "24")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBk.HandleClientGet(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.Client
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get: ")
	fmt.Println(bdy)
	if w.Code != http.StatusUnauthorized {
		t.Fail()
	}
	testMode = true
}

func TestCliRt_HandleGet(t *testing.T) {
	//var routeBkStr string = strconv.FormatInt(routeBk, 10)
	var CIDStr string = strconv.FormatInt(clientHandid, 10)
	r, _ := http.NewRequest("GET", "/test?clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "24")
	r.Header.Set("clientId", "24")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBk.HandleClientGet(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.Client
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.ClientID != clientHandid || bdy.Enabled != false {
		t.Fail()
	}
}

func TestCliRt_HandleGetListMethod(t *testing.T) {
	//var routeBkStr string = strconv.FormatInt(routeBk, 10)
	//var CIDStr string = strconv.FormatInt(clientHandid, 10)
	r, _ := http.NewRequest("POST", "/test", nil)
	r.Header.Set("u-client-id", "24")
	r.Header.Set("clientId", "24")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBk.HandleClientList(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy []mgr.Client
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get List: ")
	fmt.Println(bdy)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestCliRt_HandleGetListAuth(t *testing.T) {
	testMode = false
	//var routeBkStr string = strconv.FormatInt(routeBk, 10)
	//var CIDStr string = strconv.FormatInt(clientHandid, 10)
	r, _ := http.NewRequest("GET", "/test", nil)
	r.Header.Set("u-client-id", "24")
	r.Header.Set("clientId", "24")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBk.HandleClientList(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy []mgr.Client
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get List: ")
	fmt.Println(bdy)
	if w.Code != http.StatusUnauthorized {
		t.Fail()
	}
	testMode = true
}

func TestCliRt_HandleGetList(t *testing.T) {
	//var routeBkStr string = strconv.FormatInt(routeBk, 10)
	//var CIDStr string = strconv.FormatInt(clientHandid, 10)
	r, _ := http.NewRequest("GET", "/test", nil)
	r.Header.Set("u-client-id", "24")
	r.Header.Set("clientId", "24")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBk.HandleClientList(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy []mgr.Client
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get List: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || len(bdy) == 0 {
		t.Fail()
	}
}

func TestCliRt_HandleDelReq(t *testing.T) {
	//var routeBkStr string = strconv.FormatInt(routeBk, 10)
	//var CIDStr string = strconv.FormatInt(clientHandid, 10)
	r, _ := http.NewRequest("DELETE", "/test?clientId=w", nil)
	r.Header.Set("u-client-id", "24")
	r.Header.Set("clientId", "24")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBk.HandleClientDelete(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy BreakerResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp delete: ")
	fmt.Println(bdy)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestCliRt_HandleDelMethod(t *testing.T) {
	//var routeBkStr string = strconv.FormatInt(routeBk, 10)
	var CIDStr string = strconv.FormatInt(clientHandid, 10)
	r, _ := http.NewRequest("GET", "/test?clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "24")
	r.Header.Set("clientId", "24")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBk.HandleClientDelete(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy BreakerResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp delete: ")
	fmt.Println(bdy)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestCliRt_HandleDelAuth(t *testing.T) {
	testMode = false
	//var routeBkStr string = strconv.FormatInt(routeBk, 10)
	var CIDStr string = strconv.FormatInt(clientHandid, 10)
	r, _ := http.NewRequest("DELETE", "/test?clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "24")
	r.Header.Set("clientId", "24")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBk.HandleClientDelete(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy BreakerResponse
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp delete: ")
	fmt.Println(bdy)
	if w.Code != http.StatusUnauthorized {
		t.Fail()
	}
	testMode = true
}

func TestCliRt_HandleDel(t *testing.T) {
	//var routeBkStr string = strconv.FormatInt(routeBk, 10)
	var CIDStr string = strconv.FormatInt(clientHandid, 10)
	r, _ := http.NewRequest("DELETE", "/test?clientId="+CIDStr, nil)
	r.Header.Set("u-client-id", "24")
	r.Header.Set("clientId", "24")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hBk.HandleClientDelete(w, r)
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

func TestCliRt_TestCloseDb(t *testing.T) {
	success := gwCh.DbConfig.CloseDb()
	if success != true {
		t.Fail()
	}
}
