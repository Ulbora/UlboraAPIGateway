package circuitbreaker

import (
	mgr "UlboraApiGateway/managers"
	"fmt"
	"testing"
)

var gatewayDB CircuitBreaker
var gatewayDB2 mgr.GatewayDB
var connected1 bool
var connected2 bool
var clientID int64
var insertID int64

var routeID int64

var routeURLID int64

func TestCircuitBreaker_ConnectDb(t *testing.T) {
	clientID = 433477888567
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

func TestCircuitBreaker_InsertClient(t *testing.T) {
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

func TestCircuitBreaker_InsertRestRoute(t *testing.T) {
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

func TestCircuitBreaker_InsertRouteURL(t *testing.T) {
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

func TestCircuitBreaker_InsertBreaker(t *testing.T) {
	var b Breaker
	b.ClientID = clientID
	b.FailureThreshold = 5
	b.HealthCheckTimeSeconds = 120
	b.FailoverRouteName = "blue"
	b.OpenFailCode = 500
	b.RestRouteID = routeID
	b.RouteURIID = routeURLID
	suc, err := gatewayDB.InsertBreaker(&b)
	if suc != true || err != nil {
		t.Fail()
	}
}

var bid int64

func TestCircuitBreaker_GetRouteBreaker(t *testing.T) {
	var b Breaker
	b.ClientID = clientID
	b.RestRouteID = routeID
	b.RouteURIID = routeURLID
	res := gatewayDB.GetBreaker(&b)
	fmt.Println("")
	fmt.Print("found breaker: ")
	fmt.Println(res)
	bid = res.ID
	if res.FailureThreshold != 5 {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestCircuitBreaker_UpdateRouteBreaker(t *testing.T) {
	var b Breaker
	b.ClientID = clientID
	b.ID = bid
	b.FailureThreshold = 3
	b.HealthCheckTimeSeconds = 60
	b.FailoverRouteName = "green"
	b.OpenFailCode = 400
	b.RestRouteID = routeID
	b.RouteURIID = routeURLID

	suc, err := gatewayDB.UpdateBreaker(&b)
	if suc != true {
		fmt.Println(err)
		t.Fail()
	}
}

func TestCircuitBreaker_GetRouteBreaker2(t *testing.T) {
	var b Breaker
	b.ClientID = clientID
	b.RestRouteID = routeID
	b.RouteURIID = routeURLID
	res := gatewayDB.GetBreaker(&b)
	fmt.Println("")
	fmt.Print("found breaker: ")
	fmt.Println(res)
	if res.FailureThreshold != 3 {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestCircuitBreaker_DeleteBreaker(t *testing.T) {
	var b Breaker
	b.ClientID = clientID
	b.RestRouteID = routeID
	b.RouteURIID = routeURLID
	res := gatewayDB.DeleteBreaker(&b)
	if res != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestCircuitBreaker_DeleteClient(t *testing.T) {
	var c mgr.Client
	c.ClientID = clientID
	res := gatewayDB2.DeleteClient(&c)
	if res.Success != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestCircuitBreaker_TestCloseDb(t *testing.T) {
	success := gatewayDB.CloseDb()
	success2 := gatewayDB2.CloseDb()
	if success != true || success2 != true {
		t.Fail()
	}
}
