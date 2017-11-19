package circuitbreaker

import (
	mgr "UlboraApiGateway/managers"
	"fmt"
	"testing"
	"time"
)

var gatewayDBex CircuitBreaker
var gatewayDB2ex mgr.GatewayDB
var connected1ex bool
var connected2ex bool
var clientIDex int64
var insertIDex int64

var routeIDex int64

var routeURLIDex int64

// these tests should be successful with cache server running with the circuit breaker
func TestCircuitBreakerexg_ConnectDb(t *testing.T) {
	clientIDex = 4778887565677
	//useing external cache server---------------
	gatewayDBex.CacheHost = "http://localhost:3010"
	//useing external cache server---------------
	gatewayDBex.DbConfig.Host = "localhost:3306"
	gatewayDBex.DbConfig.DbUser = "admin"
	gatewayDBex.DbConfig.DbPw = "admin"
	gatewayDBex.DbConfig.DatabaseName = "ulbora_api_gateway"
	connected1ex = gatewayDBex.ConnectDb()

	gatewayDB2ex.DbConfig.Host = "localhost:3306"
	gatewayDB2ex.DbConfig.DbUser = "admin"
	gatewayDB2ex.DbConfig.DbPw = "admin"
	gatewayDB2ex.DbConfig.DatabaseName = "ulbora_api_gateway"
	connected2ex = gatewayDB2ex.ConnectDb()
	if connected1ex != true || connected2ex != true {
		t.Fail()
	}
}

func TestCircuitBreakerexg_InsertClient(t *testing.T) {
	var c mgr.Client
	c.APIKey = "12233hgdd333"
	c.ClientID = clientIDex
	c.Enabled = true
	c.Level = "small"

	res := gatewayDB2ex.InsertClient(&c)
	if res.Success == true && res.ID != -1 {
		fmt.Print("new client Id: ")
		fmt.Println(clientIDex)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestCircuitBreakerexg_InsertRestRoute(t *testing.T) {
	var rr mgr.RestRoute
	rr.Route = "content"
	rr.ClientID = clientIDex

	res := gatewayDB2ex.InsertRestRoute(&rr)
	if res.Success == true && res.ID != -1 {
		routeIDex = res.ID
		fmt.Print("new route Id: ")
		fmt.Println(routeIDex)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestCircuitBreakerexg_InsertRouteURL(t *testing.T) {
	var ru mgr.RouteURL
	ru.Name = "blue"
	ru.URL = "http://www.apigateway.com/blue/"
	ru.Active = false
	ru.RouteID = routeIDex
	ru.ClientID = clientIDex

	res := gatewayDB2ex.InsertRouteURL(&ru)
	if res.Success == true && res.ID != -1 {
		routeURLIDex = res.ID
		fmt.Print("new route url Id: ")
		fmt.Println(routeURLIDex)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestCircuitBreakerexg_InsertBreaker(t *testing.T) {
	var b Breaker
	b.ClientID = clientIDex
	b.FailureThreshold = 5
	b.HealthCheckTimeSeconds = 120
	b.FailoverRouteName = "blue"
	b.OpenFailCode = 500
	b.RestRouteID = routeIDex
	b.RouteURIID = routeURLIDex
	suc, err := gatewayDBex.InsertBreaker(&b)
	if suc != true || err != nil {
		t.Fail()
	}
}

var bidex int64

func TestCircuitBreakerexg_GetRouteBreaker(t *testing.T) {
	var b Breaker
	b.ClientID = clientIDex
	b.RestRouteID = routeIDex
	b.RouteURIID = routeURLIDex
	res := gatewayDBex.GetBreaker(&b)
	fmt.Println("")
	fmt.Print("found breaker: ")
	fmt.Println(res)
	bidex = res.ID
	if res.FailureThreshold != 5 {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestCircuitBreakerexg_UpdateRouteBreaker(t *testing.T) {
	var b Breaker
	b.ClientID = clientIDex
	b.ID = bidex
	b.FailureThreshold = 3
	b.HealthCheckTimeSeconds = 2
	b.FailoverRouteName = "green"
	b.OpenFailCode = 400
	b.RestRouteID = routeIDex
	b.RouteURIID = routeURLIDex

	suc, err := gatewayDBex.UpdateBreaker(&b)
	if suc != true {
		fmt.Println(err)
		t.Fail()
	}
}

var thebreakerex *Breaker

func TestCircuitBreakerexg_GetRouteBreaker2(t *testing.T) {
	var b Breaker
	b.ClientID = clientIDex
	b.RestRouteID = routeIDex
	b.RouteURIID = routeURLIDex
	res := gatewayDBex.GetBreaker(&b)
	thebreakerex = res
	fmt.Println("")
	fmt.Print("found breaker: ")
	fmt.Println(res)
	if res.FailureThreshold != 3 {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestCircuitBreakerexg_GetStatus(t *testing.T) {
	res := gatewayDBex.GetStatus(clientIDex, routeURLIDex)
	fmt.Println("")
	fmt.Print("found breaker status: ")
	fmt.Println(res)
	if res.Warning != false {
		fmt.Println("status failed")
		t.Fail()
	}
}

func TestCircuitBreakerexg_Trip(t *testing.T) {
	gatewayDBex.Trip(thebreakerex)
	res := gatewayDBex.GetStatus(clientIDex, routeURLIDex)
	fmt.Println(res)
	if res.Open == true || res.Warning != true || res.PartialOpen != true || res.FailoverRouteName != "" {
		fmt.Println("circuit breaker should be partially open")
		t.Fail()
	}
}

func TestCircuitBreakerexg_Trip2(t *testing.T) {
	gatewayDBex.Trip(thebreakerex)
	res := gatewayDBex.GetStatus(clientIDex, routeURLIDex)
	fmt.Println(res)
	if res.Open == true || res.Warning != true || res.PartialOpen != true || res.FailoverRouteName != "" {
		fmt.Println("circuit breaker should be partially open")
		t.Fail()
	}
}

func TestCircuitBreakerexg_Trip3(t *testing.T) {
	gatewayDBex.Trip(thebreakerex)
	res := gatewayDBex.GetStatus(clientIDex, routeURLIDex)
	fmt.Println(res)
	if res.Open != true || res.Warning != true || res.PartialOpen == true || res.FailoverRouteName != "green" {
		fmt.Println("circuit breaker should be open")
		t.Fail()
	}
	time.Sleep(time.Second * 3)
}

func TestCircuitBreakerexg_GetStatus2(t *testing.T) {
	res := gatewayDBex.GetStatus(clientIDex, routeURLIDex)
	fmt.Println("")
	fmt.Print("found breaker status: ")
	fmt.Println(res)
	if res.Open == true || res.Warning != true || res.PartialOpen != true || res.FailoverRouteName != "" {
		fmt.Println("circuit breaker should be partially open")
		t.Fail()
	}
}

func TestCircuitBreakerexg_Trip4(t *testing.T) {
	gatewayDBex.Trip(thebreakerex)
	res := gatewayDBex.GetStatus(clientIDex, routeURLIDex)
	fmt.Println(res)
	if res.Open != true || res.Warning != true || res.PartialOpen == true || res.FailoverRouteName != "green" {
		fmt.Println("circuit breaker should be open")
		t.Fail()
	}
}

func TestCircuitBreakerexg_Reset(t *testing.T) {
	gatewayDBex.Reset(clientIDex, routeURLIDex)
	res := gatewayDBex.GetStatus(clientIDex, routeURLIDex)
	fmt.Println(res)
	if res.Open == true || res.Warning == true || res.PartialOpen == true || res.FailoverRouteName != "" {
		fmt.Println("circuit breaker should be closed")
		t.Fail()
	}
}

func TestCircuitBreakerexg_DeleteBreaker(t *testing.T) {
	var b Breaker
	b.ClientID = clientIDex
	b.RestRouteID = routeIDex
	b.RouteURIID = routeURLIDex
	res := gatewayDBex.DeleteBreaker(&b)
	if res != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestCircuitBreakerexg_DeleteClient(t *testing.T) {
	var c mgr.Client
	c.ClientID = clientIDex
	res := gatewayDB2ex.DeleteClient(&c)
	if res.Success != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestCircuitBreakerexg_TestCloseDb(t *testing.T) {
	success := gatewayDBex.CloseDb()
	success2 := gatewayDB2ex.CloseDb()
	if success != true || success2 != true {
		t.Fail()
	}
}
