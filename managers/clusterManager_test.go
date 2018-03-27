package managers

import (
	"strconv"
	//"time"
	//"strconv"
	//env "UlboraApiGateway/environment"
	ch "UlboraApiGateway/cache"
	cb "UlboraApiGateway/circuitbreaker"
	"fmt"
	"testing"
)

var gwRoutes GatewayRoutes
var clustCbDB cb.CircuitBreaker
var connectedForCache bool

var cpc ch.CProxy
var clustCid int64
var routeClust int64
var routeClustURLID int64

func TestClusterGatewayRoutes_ConnectForCache(t *testing.T) {
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
	cpc.Host = "http://localhost:3010"
	gwRoutes.GwCacheHost = cpc.Host
	clustCbDB.DbConfig = gwRoutes.GwDB.DbConfig
	clustCbDB.CacheHost = cpc.Host
}

func TestClusterGatewayRoutes_InsertClientForCache(t *testing.T) {
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

func TestClusterGatewayRoutes_InsertRestRoute(t *testing.T) {
	var rr RestRoute
	rr.Route = "content"
	rr.ClientID = clustCid

	res := gatewayDB3.InsertRestRoute(&rr)
	if res.Success == true && res.ID != -1 {
		routeClust = res.ID
		fmt.Print("new route Id: ")
		fmt.Println(routeClust)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestClusterGatewayRoutes_InsertRouteURL(t *testing.T) {
	var ru RouteURL
	ru.Name = "blue"
	ru.URL = "http://www.apigateway.com/blue/"
	ru.Active = false
	ru.RouteID = routeClust
	ru.ClientID = clustCid

	res := gatewayDB3.InsertRouteURL(&ru)
	if res.Success == true && res.ID != -1 {
		//routeURLID3 = res.ID
		fmt.Print("new route url Id: ")
		fmt.Println(res.ID)
		routeClustURLID = res.ID
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}

	var ru2 RouteURL
	ru2.Name = "sideb"
	ru2.URL = "http://www.apigateway.com/blue/"
	ru2.Active = false
	ru2.RouteID = routeClust
	ru2.ClientID = clustCid

	res2 := gatewayDB3.InsertRouteURL(&ru2)
	if res2.Success == true && res2.ID != -1 {
		//routeURLID33 = res2.ID
		fmt.Print("new route url Id: ")
		fmt.Println(res2.ID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestClusterGatewayRoutes_InsertBreaker(t *testing.T) {
	var b cb.Breaker
	b.ClientID = clustCid
	b.FailureThreshold = 2
	b.HealthCheckTimeSeconds = 120
	b.FailoverRouteName = "blue"
	b.OpenFailCode = 500
	b.RestRouteID = routeClust
	b.RouteURIID = routeClustURLID
	suc, err := clustCbDB.InsertBreaker(&b)
	if suc != true || err != nil {
		t.Fail()
	}
}

// func TestGatewayRoutes_SetGatewayRouteStatus2(t *testing.T) {

// 	//cp2.Host = "http://localhost:3010"
// 	gwRoutes.ClientID = clustCid
// 	gwRoutes.Route = "testroute"
// 	//gwRoutes.APIKey = "12345"
// 	gwRoutes.GwCacheHost = "http://localhost1:3010"

// 	res := gwRoutes.SetGatewayRouteStatus()
// 	if res != false {
// 		t.Fail()
// 	}
// }

// func TestGatewayRoutes_GetGatewayRouteStatus(t *testing.T) {

// 	//cp2.Host = "http://localhost:3010"
// 	gwRoutes.ClientID = clustCid
// 	gwRoutes.Route = "testroute"
// 	//gwRoutes.APIKey = "12345"
// 	gwRoutes.GwCacheHost = env.GetCacheHost() // "http://localhost:3010"

// 	res := gwRoutes.GetGatewayRouteStatus()
// 	fmt.Println(res)
// 	if res.Success != true && res.RouteModified != true {
// 		t.Fail()
// 	}
// }

// func TestGatewayRoutes_GetGatewayRouteStatus2(t *testing.T) {

// 	gwRoutes.ClientID = clustCid
// 	gwRoutes.Route = "testroute"
// 	//gwRoutes.APIKey = "12345"
// 	gwRoutes.GwCacheHost = "http://localhost2:3010"
// 	res := gwRoutes.GetGatewayRouteStatus()
// 	fmt.Println(res)
// 	if res.Success == true || res.RouteModified == true {
// 		t.Fail()
// 	}
// }

func TestClusterGatewayRoutes_GetClusterGwRoutes(t *testing.T) {
	gwRoutes.ClientID = clustCid
	gwRoutes.Route = "content"
	gwRoutes.APIKey = "12233hgdd333"
	//gwRoutes.GwCacheHost = "http://localhost2:3010"
	res := gwRoutes.GetClusterGwRoutes()
	fmt.Print("found routes: ")
	fmt.Println(res)
	if len(*res) != 2 {
		t.Fail()
	}
}

func TestClusterGatewayRoutes_ClearClusterGwRoutes(t *testing.T) {
	//time.Sleep(2000 * time.Millisecond)
	gwRoutes.ClientID = clustCid
	gwRoutes.Route = "content"
	res := gwRoutes.ClearClusterGwRoutes()
	if res != true {
		t.Fail()
	}
}

func TestClusterGatewayRoutes_GetCache(t *testing.T) {
	var cid = strconv.FormatInt(clustCid, 10)
	var key = cid + ":cluster:" + "content"
	rtn := cpc.Get(key)
	fmt.Print("routes after del: ")
	fmt.Println(rtn)
	if rtn.Success == true {
		t.Fail()
	}
}

func TestClusterGatewayRoutes_TripClusterGwRoutes1(t *testing.T) {
	gwRoutes.ClientID = clustCid
	//gwRoutes.Route = "content"
	//gwRoutes.APIKey = "12233hgdd333"
	//gwRoutes.GwCacheHost = "http://localhost2:3010"
	var b cb.Breaker
	b.ClientID = clustCid
	b.FailureThreshold = 2
	b.HealthCheckTimeSeconds = 120
	b.FailoverRouteName = "blue"
	b.OpenFailCode = 500
	b.RestRouteID = routeClust
	b.RouteURIID = routeClustURLID
	res := gwRoutes.TripClusterGwRoutes(&b)
	fmt.Print("found routes: ")
	fmt.Println(res)
	if res != true {
		t.Fail()
	}
}

func TestClusterGatewayRoutes_TripClusterGwRoutes2(t *testing.T) {
	gwRoutes.ClientID = clustCid
	//gwRoutes.Route = "content"
	//gwRoutes.APIKey = "12233hgdd333"
	//gwRoutes.GwCacheHost = "http://localhost2:3010"
	var b cb.Breaker
	b.ClientID = clustCid
	b.FailureThreshold = 2
	b.HealthCheckTimeSeconds = 120
	b.FailoverRouteName = "blue"
	b.OpenFailCode = 500
	b.RestRouteID = routeClust
	b.RouteURIID = routeClustURLID
	res := gwRoutes.TripClusterGwRoutes(&b)
	fmt.Print("found routes: ")
	fmt.Println(res)
	if res != true {
		t.Fail()
	}
}

func TestClusterGatewayRoutes_TripClusterGwRoutes3(t *testing.T) {
	gwRoutes.ClientID = clustCid
	//gwRoutes.Route = "content"
	//gwRoutes.APIKey = "12233hgdd333"
	//gwRoutes.GwCacheHost = "http://localhost2:3010"
	var b cb.Breaker
	b.ClientID = clustCid
	b.FailureThreshold = 2
	b.HealthCheckTimeSeconds = 120
	b.FailoverRouteName = "blue"
	b.OpenFailCode = 500
	b.RestRouteID = routeClust
	b.RouteURIID = routeClustURLID
	res := gwRoutes.TripClusterGwRoutes(&b)
	fmt.Print("found routes: ")
	fmt.Println(res)
	if res != true {
		t.Fail()
	}
}

func TestClusterGatewayRoutes_GetBreakerStatus(t *testing.T) {

	res := clustCbDB.GetStatus(clustCid, routeClustURLID)
	fmt.Print("routes status: ")
	fmt.Println(res)
	if res.Open != true {
		t.Fail()
	}
}

// func TestClusterGatewayRoutes_GetClusterGwRoutes2(t *testing.T) {
// 	time.Sleep(2000 * time.Millisecond)
// 	gwRoutes.ClientID = clustCid
// 	gwRoutes.Route = "content"
// 	//gwRoutes.GwCacheHost = "http://localhost2:3010"
// 	res := gwRoutes.GetClusterGwRoutes()
// 	fmt.Print("found routes: ")
// 	fmt.Println(res)
// 	if len(*res) != 2 {
// 		t.Fail()
// 	}
// }

// func TestGatewayRoutes_DeleteGatewayRouteStatus(t *testing.T) {

// 	gwRoutes.ClientID = clustCid
// 	gwRoutes.Route = "testroute"
// 	gwRoutes.APIKey = "12345"
// 	gwRoutes.GwCacheHost = "http://localhost:3010"
// 	res := gwRoutes.DeleteGatewayRouteStatus()
// 	fmt.Println(res)
// 	if res.Success == true {
// 		t.Fail()
// 	}
// }

// func TestGatewayRoutes_DeleteGatewayRouteStatus2(t *testing.T) {

// 	gwRoutes.ClientID = clustCid
// 	gwRoutes.Route = "testroute"
// 	gwRoutes.APIKey = "12233hgdd333"
// 	gwRoutes.GwCacheHost = "http://localhost:3010"
// 	res := gwRoutes.DeleteGatewayRouteStatus()
// 	fmt.Println(res)
// 	if res.Success != true {
// 		t.Fail()
// 	}
// }

func TestClusterGatewayRoutes_DeleteClientForCache(t *testing.T) {
	var c Client
	c.ClientID = clustCid
	res := gatewayDB.DeleteClient(&c)
	if res.Success != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestClusterGatewayRoutes_TestCloseDb2(t *testing.T) {
	success := gatewayDB3.CloseDb()
	if success != true {
		t.Fail()
	}
}
