package gwerrors

import (
	mgr "UlboraApiGateway/managers"
	"fmt"
	"testing"
	"time"
)

var gatewayDB GatewayErrorMonitor
var gatewayDB2 mgr.GatewayDB
var connected1 bool
var connected2 bool
var clientID int64
var insertID int64

var routeID int64

var routeURLID int64

func TestGatewayErrorMonitor_ConnectDb(t *testing.T) {
	clientID = 46677777777
	gatewayDB.DbConfig.Host = "localhost:3306"
	gatewayDB.DbConfig.DbUser = "admin"
	gatewayDB.DbConfig.DbPw = "admin"
	gatewayDB.DbConfig.DatabaseName = "ulbora_api_gateway"
	connected1 = gatewayDB.ConnectDb()

	gatewayDB2.DbConfig.Host = "localhost:3306"
	gatewayDB2.DbConfig.DbUser = "admin"
	gatewayDB2.DbConfig.DbPw = "admin"
	gatewayDB2.DbConfig.DatabaseName = "ulbora_api_gateway"
	connected2 = gatewayDB2.ConnectDb()
	if connected1 != true || connected2 != true {
		t.Fail()
	}
}

func TestGatewayErrorMonitor_InsertClient(t *testing.T) {
	var c mgr.Client
	c.APIKey = "12233hgdd333"
	c.ClientID = clientID
	c.Enabled = true
	c.Level = "small"

	res := gatewayDB2.InsertClient(&c)
	if res.Success == true && res.ID != -1 {
		fmt.Print("new client Id: ")
		fmt.Println(clientID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestGatewayErrorMonitor_InsertRestRoute(t *testing.T) {
	var rr mgr.RestRoute
	rr.Route = "content"
	rr.ClientID = clientID

	res := gatewayDB2.InsertRestRoute(&rr)
	if res.Success == true && res.ID != -1 {
		routeID = res.ID
		fmt.Print("new route Id: ")
		fmt.Println(routeID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestGatewayErrorMonitor_InsertRouteURL(t *testing.T) {
	var ru mgr.RouteURL
	ru.Name = "blue"
	ru.URL = "http://www.apigateway.com/blue/"
	ru.Active = false
	ru.RouteID = routeID
	ru.ClientID = clientID

	res := gatewayDB2.InsertRouteURL(&ru)
	if res.Success == true && res.ID != -1 {
		routeURLID = res.ID
		fmt.Print("new route url Id: ")
		fmt.Println(routeURLID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestGatewayErrorMonitor_InsertRouteError(t *testing.T) {
	var e GwError
	e.ClientID = clientID
	e.Code = 500
	e.Entered = time.Now().Add(time.Hour * -2400)
	e.Message = "internal error"
	e.RestRouteID = routeID
	e.RouteURIID = routeURLID
	suc, err := gatewayDB.InsertRouteError(&e)
	if suc != true || err != nil {
		t.Fail()
	}
}

func TestGatewayErrorMonitor_SaveRouteError(t *testing.T) {
	// var e GwError
	// e.ClientID = clientID
	// e.Code = 500
	// e.Entered = time.Now().Add(time.Hour * -2400)
	// e.Message = "internal error"
	// e.RestRouteID = routeID
	// e.RouteURIID = routeURLID
	// suc, err := gatewayDB.InsertRouteError(&e)
	suc, err := gatewayDB.SaveRouteError(clientID, 400, "something went wrong", routeID, routeURLID)
	if suc != true || err != nil {
		t.Fail()
	}
}

func TestGatewayErrorMonitor_GetRouteError(t *testing.T) {
	var e GwError
	e.ClientID = clientID
	e.RestRouteID = routeID
	e.RouteURIID = routeURLID
	res := gatewayDB.GetRouteError(&e)
	fmt.Println("")
	fmt.Print("found gw error list: ")
	fmt.Println(res)
	if len(*res) == 0 {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestGatewayErrorMonitor_DeleteRouteError(t *testing.T) {
	res := gatewayDB.DeleteRouteError()
	if res != false {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestGatewayErrorMonitor_DeleteClient(t *testing.T) {
	var c mgr.Client
	c.ClientID = clientID
	res := gatewayDB2.DeleteClient(&c)
	if res.Success != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestGatewayErrorMonitor_TestCloseDb(t *testing.T) {
	success := gatewayDB.CloseDb()
	success2 := gatewayDB2.CloseDb()
	if success != true || success2 != true {
		t.Fail()
	}
}
