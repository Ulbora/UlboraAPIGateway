package managers

import (
	"fmt"
	"testing"
)

var gatewayDB2 GatewayDB
var connected2 bool
var clientID2 int64
var insertID2 int64
var routeID int64

//var routeURLID21 int64
//var routeURLID22 int64

func TestRestRoute_ConnectDb2(t *testing.T) {
	clientID2 = 6
	gatewayDB2.DbConfig.Host = "localhost:3306"
	gatewayDB2.DbConfig.DbUser = "admin"
	gatewayDB2.DbConfig.DbPw = "admin"
	gatewayDB2.DbConfig.DatabaseName = "ulbora_api_gateway"
	connected2 = gatewayDB2.ConnectDb()
	if connected2 != true {
		t.Fail()
	}
}

func TestRestRoute_InsertClient2(t *testing.T) {
	var c Client
	c.APIKey = "12233hgdd333"
	c.ClientID = clientID2
	c.Enabled = true
	c.Level = "small"

	res := gatewayDB2.InsertClient(&c)
	if res.Success == true && res.ID != -1 {
		insertID2 = clientID2
		fmt.Print("new client Id: ")
		fmt.Println(insertID2)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestRestRoute_InsertRestRoute(t *testing.T) {
	var rr RestRoute
	rr.Route = "content"
	rr.ClientID = clientID2

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

func TestRestRoute_UpdateRestRoute(t *testing.T) {
	var rr RestRoute
	rr.ID = routeID
	rr.Route = "content2"
	rr.ClientID = clientID2
	gatewayDB.GwCacheHost = "http://localhost:3010"
	res := gatewayDB.UpdateRestRoute(&rr)
	if res.Success != true {
		fmt.Println("database update failed")
		t.Fail()
	}
}

func TestRestRoute_GetRestRoute(t *testing.T) {
	var rr RestRoute
	rr.ID = routeID
	rr.ClientID = clientID2
	res := gatewayDB.GetRestRoute(&rr)
	fmt.Println("")
	fmt.Print("found route: ")
	fmt.Println(res)
	if res.Route != "content2" {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestRestRoute_GetRestRouteList(t *testing.T) {
	var rr RestRoute
	rr.ClientID = clientID2
	res := gatewayDB.GetRestRouteList(&rr)
	fmt.Println("")
	fmt.Print("found client list: ")
	fmt.Println(res)
	if len(*res) != 1 {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestRestRoute_DeleteRestRouteList(t *testing.T) {
	var rr RestRoute
	rr.ID = routeID
	rr.ClientID = clientID2
	res := gatewayDB.DeleteRestRoute(&rr)
	if res.Success != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestRestRoute_DeleteClient(t *testing.T) {
	var c Client
	c.ClientID = clientID2
	res := gatewayDB.DeleteClient(&c)
	if res.Success != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestRestRoute_TestCloseDb(t *testing.T) {
	success := gatewayDB2.CloseDb()
	if success != true {
		t.Fail()
	}
}
