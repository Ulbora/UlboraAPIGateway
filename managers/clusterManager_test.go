package managers

import (
	"testing"
)

var gwRoutes GatewayRoutes

//var cp2 ch.CProxy
var clustCid int64

func TestGatewayRoutes_SetGatewayRouteStatus(t *testing.T) {
	clustCid = 8
	//cp2.Host = "http://localhost:3010"
	gwRoutes.ClientID = clustCid
	gwRoutes.Route = "testroute"
	gwRoutes.GwCacheHost = "http://localhost:3010"

	res := gwRoutes.SetGatewayRouteStatus()
	if res != true {
		t.Fail()
	}
}

func TestGatewayRoutes_GetGatewayRouteStatus(t *testing.T) {
	clustCid = 8
	//cp2.Host = "http://localhost:3010"
	gwRoutes.ClientID = clustCid
	gwRoutes.Route = "testroute"
	gwRoutes.GwCacheHost = "http://localhost:3010"

	res := gwRoutes.GetGatewayRouteStatus()
	if res.RouteModified != true {
		t.Fail()
	}
}
