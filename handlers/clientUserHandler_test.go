package handlers

import (
	//"strconv"
	//gwerr "UlboraApiGateway/gwerrors"
	mgr "UlboraApiGateway/managers"
	//"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	//"time"
)

var uCid int64 = 29
var connectedForU bool

//var udb gwerr.GatewayErrorMonitor
var gwU mgr.GatewayDB
var hu Handler

func TestCU1_Connect(t *testing.T) {
	gwU.DbConfig.Host = "localhost:3306"
	gwU.DbConfig.DbUser = "admin"
	gwU.DbConfig.DbPw = "admin"
	gwU.DbConfig.DatabaseName = "ulbora_api_gateway"
	connectedForU = gwU.ConnectDb()
	if connectedForU != true {
		t.Fail()
	}
	//gwRoutesErr.GwDB.DbConfig = edb.DbConfig
	//gwRoutes.GwDB.DbConfig = gwRoutes.GwDB.DbConfig
	//cp.Host = "http://localhost:3010"
	testMode = true
	hu.DbConfig = gwU.DbConfig
}

// func TestErr_SetManager(t *testing.T) {
// 	SetManager(edb)
// }

func TestCU1_InsertClient(t *testing.T) {
	var c mgr.Client
	c.APIKey = "12233hgdd333"
	c.ClientID = uCid
	c.Enabled = true
	c.Level = "small"

	res := gwU.InsertClient(&c)
	if res.Success == true && res.ID != -1 {
		fmt.Print("new client Id: ")
		fmt.Println(uCid)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestCU1_HandleUserClientGetReq(t *testing.T) {
	//var routeBkStr string = strconv.FormatInt(routeBk, 10)
	//var CIDStr string = strconv.FormatInt(uCid, 10)
	r, _ := http.NewRequest("GET", "/test", nil)
	r.Header.Set("u-client-id", "29")
	//r.Header.Set("clientId", "29")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hu.HandleUserClient(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.Client
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.ClientID != 0 {
		t.Fail()
	}
}

func TestCU1_HandleUserClientGetMethod(t *testing.T) {
	//var routeBkStr string = strconv.FormatInt(routeBk, 10)
	//var CIDStr string = strconv.FormatInt(uCid, 10)
	r, _ := http.NewRequest("DELETE", "/test", nil)
	r.Header.Set("u-client-id", "29")
	r.Header.Set("clientId", "29")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hu.HandleUserClient(w, r)
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

func TestCU1_HandleUserClientGet(t *testing.T) {
	//var routeBkStr string = strconv.FormatInt(routeBk, 10)
	//var CIDStr string = strconv.FormatInt(uCid, 10)
	r, _ := http.NewRequest("GET", "/test", nil)
	r.Header.Set("u-client-id", "29")
	r.Header.Set("clientId", "29")
	r.Header.Set("u-api-key", "12233hgdd3335")
	w := httptest.NewRecorder()
	hu.HandleUserClient(w, r)
	fmt.Print("Code: ")
	fmt.Println(w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	var bdy mgr.Client
	json.Unmarshal([]byte(b), &bdy)
	fmt.Print("Resp Get: ")
	fmt.Println(bdy)
	if w.Code != http.StatusOK || bdy.ClientID != uCid || bdy.Enabled != true {
		t.Fail()
	}
}

func TestCU1_DeleteClient(t *testing.T) {
	var c mgr.Client
	c.ClientID = uCid
	res := gwU.DeleteClient(&c)
	if res.Success != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestCU1_TestCloseDb(t *testing.T) {
	success := gwU.CloseDb()
	if success != true {
		t.Fail()
	}
}
