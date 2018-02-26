package managers

import (
	//"strconv"
	env "UlboraApiGateway/environment"
	"fmt"
	"testing"
)

var gwRoutes GatewayRoutes
var connectedForCache bool

//var cp2 ch.CProxy
var clustCid int64

func TestGatewayRoutes_ConnectForCache(t *testing.T) {
	clustCid = 88
	gwRoutes.GwDB.DbConfig.Host = "localhost:3306"
	gwRoutes.GwDB.DbConfig.DbUser = "admin"
	gwRoutes.GwDB.DbConfig.DbPw = "admin"
	gwRoutes.GwDB.DbConfig.DatabaseName = "ulbora_api_gateway"
	connectedForCache = gwRoutes.GwDB.ConnectDb()
	if connectedForCache != true {
		t.Fail()
	}
	//gwRoutes.GwDB.DbConfig = gwRoutes.GwDB.DbConfig
	cp.Host = "http://localhost:3010"
}

func TestGatewayRoutes_InsertClientForCache(t *testing.T) {
	var c Client
	c.APIKey = "12233hgdd333"
	c.ClientID = clustCid
	c.Enabled = true
	c.Level = "small"

	res := gatewayDB4.InsertClient(&c)
	if res.Success == true && res.ID != -1 {
		//insertID4 = clientID4
		fmt.Print("new client Id: ")
		fmt.Println(clustCid)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestGatewayRoutes_SetGatewayRouteStatus(t *testing.T) {
	//clustCid = 8
	gwRoutes.ClientID = clustCid
	gwRoutes.Route = "testroute"
	//gwRoutes.APIKey = "12345"
	gwRoutes.GwCacheHost = env.GetCacheHost() // "http://localhost:3010"

	res := gwRoutes.SetGatewayRouteStatus()
	if res != true {
		t.Fail()
	}
}

func TestGatewayRoutes_SetGatewayRouteStatus2(t *testing.T) {

	//cp2.Host = "http://localhost:3010"
	gwRoutes.ClientID = clustCid
	gwRoutes.Route = "testroute"
	//gwRoutes.APIKey = "12345"
	gwRoutes.GwCacheHost = "http://localhost1:3010"

	res := gwRoutes.SetGatewayRouteStatus()
	if res != false {
		t.Fail()
	}
}

func TestGatewayRoutes_GetGatewayRouteStatus(t *testing.T) {

	//cp2.Host = "http://localhost:3010"
	gwRoutes.ClientID = clustCid
	gwRoutes.Route = "testroute"
	//gwRoutes.APIKey = "12345"
	gwRoutes.GwCacheHost = env.GetCacheHost() // "http://localhost:3010"

	res := gwRoutes.GetGatewayRouteStatus()
	fmt.Println(res)
	if res.Success != true && res.RouteModified != true {
		t.Fail()
	}
}

func TestGatewayRoutes_GetGatewayRouteStatus2(t *testing.T) {

	gwRoutes.ClientID = clustCid
	gwRoutes.Route = "testroute"
	//gwRoutes.APIKey = "12345"
	gwRoutes.GwCacheHost = "http://localhost2:3010"
	res := gwRoutes.GetGatewayRouteStatus()
	fmt.Println(res)
	if res.Success == true || res.RouteModified == true {
		t.Fail()
	}
}

func TestGatewayRoutes_DeleteGatewayRouteStatus(t *testing.T) {

	gwRoutes.ClientID = clustCid
	gwRoutes.Route = "testroute"
	gwRoutes.APIKey = "12345"
	gwRoutes.GwCacheHost = "http://localhost:3010"
	res := gwRoutes.DeleteGatewayRouteStatus()
	fmt.Println(res)
	if res.Success == true {
		t.Fail()
	}
}

func TestGatewayRoutes_DeleteGatewayRouteStatus2(t *testing.T) {

	gwRoutes.ClientID = clustCid
	gwRoutes.Route = "testroute"
	gwRoutes.APIKey = "12233hgdd333"
	gwRoutes.GwCacheHost = "http://localhost:3010"
	res := gwRoutes.DeleteGatewayRouteStatus()
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayRoutes_DeleteClientForCache(t *testing.T) {
	var c Client
	c.ClientID = clustCid
	res := gatewayDB.DeleteClient(&c)
	if res.Success != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}
