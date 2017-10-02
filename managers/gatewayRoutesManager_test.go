package managers

import (
	"fmt"
	"strconv"
	"testing"
)

var gatewayRoutes GatewayRoutes

var gatewayDB4 GatewayDB
var connected4 bool
var clientID4 int64
var insertID4 int64
var routeID4 int64

var routeURLID4 int64
var routeURLID44 int64

func TestGatewayRoutes_ConnectDb4(t *testing.T) {
	clientID4 = 8
	gatewayDB4.DbConfig.Host = "localhost:3306"
	gatewayDB4.DbConfig.DbUser = "admin"
	gatewayDB4.DbConfig.DbPw = "admin"
	gatewayDB4.DbConfig.DatabaseName = "ulbora_api_gateway"
	connected4 = gatewayDB4.ConnectDb()
	if connected4 != true {
		t.Fail()
	}
}

func TestGatewayRoutes_InsertClient4(t *testing.T) {
	var c Client
	c.APIKey = "12233hgdd333"
	c.ClientID = clientID4
	c.Enabled = true
	c.Level = "small"

	res := gatewayDB4.InsertClient(&c)
	if res.Success == true && res.ID != -1 {
		insertID4 = clientID4
		fmt.Print("new client Id: ")
		fmt.Println(insertID4)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestGatewayRoutes_InsertRestRoute4(t *testing.T) {
	var rr RestRoute
	rr.Route = "content"
	rr.ClientID = clientID4

	res := gatewayDB4.InsertRestRoute(&rr)
	if res.Success == true && res.ID != -1 {
		routeID4 = res.ID
		fmt.Print("new route Id: ")
		fmt.Println(routeID4)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestGatewayRoutes_InsertRouteURL4(t *testing.T) {
	var ru RouteURL
	ru.Name = "blue"
	ru.URL = "http://www.apigateway.com/blue/"
	ru.Active = false
	ru.RouteID = routeID4
	ru.ClientID = clientID4

	res := gatewayDB4.InsertRouteURL(&ru)
	if res.Success == true && res.ID != -1 {
		routeURLID4 = res.ID
		fmt.Print("new route url Id: ")
		fmt.Println(routeURLID4)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}

	var ru2 RouteURL
	ru2.Name = "green"
	ru2.URL = "http://www.apigateway.com/blue/"
	ru2.Active = true
	ru2.RouteID = routeID4
	ru2.ClientID = clientID4

	res2 := gatewayDB4.InsertRouteURL(&ru2)
	if res2.Success == true && res2.ID != -1 {
		routeURLID44 = res2.ID
		fmt.Print("new route url Id: ")
		fmt.Println(routeURLID44)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

// func TestGatewayRoutes_GetRouteURLs(t *testing.T) {
// 	var ru RouteURL
// 	ru.RouteID = routeID4
// 	ru.ClientID = clientID4
// 	res := gatewayDB.GetRouteURL(&ru)
// 	fmt.Println("")
// 	fmt.Print("found route URL: ")
// 	fmt.Println(res)
// 	if res.Active != true {
// 		fmt.Println("database read failed")
// 		t.Fail()
// 	}
// }

func TestGatewayRoutes_GetGatewayRoutes(t *testing.T) {
	gatewayRoutes.GwDB = gatewayDB4
	gatewayRoutes.ClientID = clientID4
	gatewayRoutes.APIKey = "12233hgdd333"
	gatewayRoutes.Route = "content"
	gatewayRoutes.GwCacheHost = "http://localhost:3010"
	var cid = strconv.FormatInt(gatewayRoutes.ClientID, 10) // string(gatewayRoutes.ClientID)
	fmt.Print("cid: ")
	fmt.Println(cid)
	var keyused = cid + ":" + gatewayRoutes.Route
	fmt.Print("Key Used: ")
	fmt.Println(keyused)
	res := gatewayRoutes.GetGatewayRoutes(true, "")
	fmt.Println("Route: ")
	fmt.Println(res)
	if res.Active != true && res.Name != "green" {
		fmt.Println("route not found")
		t.Fail()
	}
}

func TestGatewayRoutes_GetGatewayRoutes2(t *testing.T) {
	gatewayRoutes.GwDB = gatewayDB4
	gatewayRoutes.ClientID = clientID4
	gatewayRoutes.APIKey = "12233hgdd333"
	gatewayRoutes.Route = "content"
	gatewayRoutes.GwCacheHost = "http://localhost:3010"
	var cid = strconv.FormatInt(gatewayRoutes.ClientID, 10) // string(gatewayRoutes.ClientID)
	fmt.Print("cid: ")
	fmt.Println(cid)
	var keyused = cid + ":" + gatewayRoutes.Route
	fmt.Print("Key Used: ")
	fmt.Println(keyused)
	res := gatewayRoutes.GetGatewayRoutes(false, "blue")
	fmt.Println("Route: ")
	fmt.Println(res)
	if res.Active != false && res.Name != "blue" {
		fmt.Println("route not found")
		t.Fail()
	}
}

func TestGatewayRoutes_DeleteClient(t *testing.T) {
	var c Client
	c.ClientID = clientID4
	res := gatewayDB.DeleteClient(&c)
	if res.Success != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestGatewayRoutes_TestCloseDb(t *testing.T) {
	success := gatewayDB3.CloseDb()
	if success != true {
		t.Fail()
	}
}
