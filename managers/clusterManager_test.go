package managers

import (
	env "UlboraApiGateway/environment"
	"fmt"
	"testing"
)

var gwRoutes GatewayRoutes

//var cp2 ch.CProxy
var clustCid int64

func TestGatewayRoutes_SetGatewayRouteStatus(t *testing.T) {
	clustCid = 8
	gwRoutes.ClientID = clustCid
	gwRoutes.Route = "testroute"
	gwRoutes.APIKey = "12345"
	gwRoutes.GwCacheHost = env.GetCacheHost() // "http://localhost:3010"

	res := gwRoutes.SetGatewayRouteStatus()
	if res != true {
		t.Fail()
	}
}

func TestGatewayRoutes_SetGatewayRouteStatus2(t *testing.T) {
	clustCid = 8
	//cp2.Host = "http://localhost:3010"
	gwRoutes.ClientID = clustCid
	gwRoutes.Route = "testroute"
	gwRoutes.APIKey = "12345"
	gwRoutes.GwCacheHost = "http://localhost1:3010"

	res := gwRoutes.SetGatewayRouteStatus()
	if res != false {
		t.Fail()
	}
}

func TestGatewayRoutes_GetGatewayRouteStatus(t *testing.T) {
	clustCid = 8
	//cp2.Host = "http://localhost:3010"
	gwRoutes.ClientID = clustCid
	gwRoutes.Route = "testroute"
	gwRoutes.APIKey = "12345"
	gwRoutes.GwCacheHost = env.GetCacheHost() // "http://localhost:3010"

	res := gwRoutes.GetGatewayRouteStatus()
	fmt.Println(res)
	if res.Success != true && res.RouteModified != true {
		t.Fail()
	}
}

func TestGatewayRoutes_GetGatewayRouteStatus2(t *testing.T) {
	clustCid = 8
	gwRoutes.ClientID = clustCid
	gwRoutes.Route = "testroute"
	gwRoutes.APIKey = "12345"
	gwRoutes.GwCacheHost = "http://localhost2:3010"
	res := gwRoutes.GetGatewayRouteStatus()
	fmt.Println(res)
	if res.Success == true || res.RouteModified == true {
		t.Fail()
	}
}
