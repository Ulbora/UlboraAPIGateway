package managers

import (
	"fmt"
	"testing"
)

var gatewayDB3 GatewayDB
var connected3 bool
var clientID3 int64
var insertID3 int64
var routeID3 int64

var routeURLID3 int64
var routeURLID33 int64

func TestRouteURL_ConnectDb2(t *testing.T) {
	clientID3 = 7
	gatewayDB3.DbConfig.Host = "localhost:3306"
	gatewayDB3.DbConfig.DbUser = "admin"
	gatewayDB3.DbConfig.DbPw = "admin"
	gatewayDB3.DbConfig.DatabaseName = "ulbora_api_gateway"
	connected3 = gatewayDB3.ConnectDb()
	if connected3 != true {
		t.Fail()
	}
}

func TestRouteURL_InsertClient2(t *testing.T) {
	var c Client
	c.APIKey = "12233hgdd333"
	c.ClientID = clientID3
	c.Enabled = true
	c.Level = "small"

	res := gatewayDB3.InsertClient(&c)
	if res.Success == true && res.ID != -1 {
		insertID3 = clientID3
		fmt.Print("new client Id: ")
		fmt.Println(insertID3)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestRouteURL_InsertRestRoute(t *testing.T) {
	var rr RestRoute
	rr.Route = "content"
	rr.ClientID = clientID3

	res := gatewayDB3.InsertRestRoute(&rr)
	if res.Success == true && res.ID != -1 {
		routeID3 = res.ID
		fmt.Print("new route Id: ")
		fmt.Println(routeID3)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestRouteURL_InsertRouteURL(t *testing.T) {
	var ru RouteURL
	ru.Name = "blue"
	ru.URL = "http://www.apigateway.com/blue/"
	ru.Active = false
	ru.RouteID = routeID3
	ru.ClientID = clientID3

	res := gatewayDB3.InsertRouteURL(&ru)
	if res.Success == true && res.ID != -1 {
		routeURLID3 = res.ID
		fmt.Print("new route url Id: ")
		fmt.Println(routeURLID3)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}

	var ru2 RouteURL
	ru2.Name = "sideb"
	ru2.URL = "http://www.apigateway.com/blue/"
	ru2.Active = false
	ru2.RouteID = routeID3
	ru2.ClientID = clientID3

	res2 := gatewayDB3.InsertRouteURL(&ru2)
	if res2.Success == true && res2.ID != -1 {
		routeURLID33 = res2.ID
		fmt.Print("new route url Id: ")
		fmt.Println(routeURLID33)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestRouteURL_UpdateRouteURL(t *testing.T) {
	var ru RouteURL
	ru.ID = routeURLID33
	ru.Name = "green"
	ru.URL = "http://www.apigateway.com/green/"
	ru.RouteID = routeID3
	ru.ClientID = clientID3

	gatewayDB.GwCacheHost = "http://localhost:3010"
	res := gatewayDB.UpdateRouteURL(&ru)
	if res.Success != true {
		fmt.Println("database update failed")
		t.Fail()
	}
}

func TestRouteURL_ActivateRouteURL(t *testing.T) {
	var ru RouteURL
	ru.ID = routeURLID33
	ru.RouteID = routeID3
	ru.ClientID = clientID3

	res := gatewayDB.ActivateRouteURL(&ru)
	if res.Success != true {
		fmt.Println("database update failed")
		t.Fail()
	}
}

func TestRouteURL_ActivateRouteURL2(t *testing.T) {
	var ru RouteURL
	ru.ID = routeURLID3
	ru.RouteID = routeID3
	ru.ClientID = clientID3

	res := gatewayDB.ActivateRouteURL(&ru)
	if res.Success != true {
		fmt.Println("database update failed")
		t.Fail()
	}
}

func TestRouteURL_ActivateRouteURL3(t *testing.T) {
	var ru RouteURL
	ru.ID = routeURLID33
	ru.RouteID = routeID3
	ru.ClientID = clientID3

	res := gatewayDB.ActivateRouteURL(&ru)
	if res.Success != true {
		fmt.Println("database update failed")
		t.Fail()
	}
}

func TestRouteURL_GetRouteURL(t *testing.T) {
	var ru RouteURL
	ru.ID = routeURLID33
	ru.RouteID = routeID3
	ru.ClientID = clientID3
	res := gatewayDB.GetRouteURL(&ru)
	fmt.Println("")
	fmt.Print("found route URL: ")
	fmt.Println(res)
	if res.Active != true {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestRouteURL_GetRouteURLList(t *testing.T) {
	var ru RouteURL
	ru.RouteID = routeID3
	ru.ClientID = clientID3
	res := gatewayDB.GetRouteURLList(&ru)
	fmt.Println("")
	fmt.Print("found route URL list: ")
	fmt.Println(res)
	if len(*res) != 2 {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestRouteURL_DeleteActiveRouteURL(t *testing.T) {
	var ru RouteURL
	ru.ID = routeURLID33
	ru.RouteID = routeID3
	ru.ClientID = clientID3
	res := gatewayDB.DeleteRouteURL(&ru)
	if res.Success != false {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestRouteURL_DeleteNonActiveRouteURL(t *testing.T) {
	var ru RouteURL
	ru.ID = routeURLID3
	ru.RouteID = routeID3
	ru.ClientID = clientID3
	res := gatewayDB.DeleteRouteURL(&ru)
	if res.Success != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestRouteURL_DeleteClient(t *testing.T) {
	var c Client
	c.ClientID = clientID3
	res := gatewayDB.DeleteClient(&c)
	if res.Success != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestRouteURL_TestCloseDb(t *testing.T) {
	success := gatewayDB3.CloseDb()
	if success != true {
		t.Fail()
	}
}
