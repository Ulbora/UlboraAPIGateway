package circuitbreaker

import (
	mgr "UlboraApiGateway/managers"
	"fmt"
	"testing"
	"time"
)

var gatewayDBexf CircuitBreaker
var gatewayDB2exf mgr.GatewayDB
var connected1exf bool
var connected2exf bool
var clientIDexf int64
var insertIDexf int64

var routeIDexf int64

var routeURLIDexf int64

// these tests should try and fail to use an external cashe server with the circuit breaker
// after fail, the circuit breaker should default to the internal cache
func TestCircuitBreakerexf_ConnectDb(t *testing.T) {
	clientIDexf = 477888777567
	//useing external cache server---------------
	gatewayDBexf.CacheHost = "http://localhost:3110"
	//useing external cache server---------------
	gatewayDBexf.DbConfig.Host = "localhost:3306"
	gatewayDBexf.DbConfig.DbUser = "admin"
	gatewayDBexf.DbConfig.DbPw = "admin"
	gatewayDBexf.DbConfig.DatabaseName = "ulbora_api_gateway"
	connected1exf = gatewayDBexf.ConnectDb()

	gatewayDB2exf.DbConfig.Host = "localhost:3306"
	gatewayDB2exf.DbConfig.DbUser = "admin"
	gatewayDB2exf.DbConfig.DbPw = "admin"
	gatewayDB2exf.DbConfig.DatabaseName = "ulbora_api_gateway"
	connected2exf = gatewayDB2exf.ConnectDb()
	if connected1exf != true || connected2exf != true {
		t.Fail()
	}
}

func TestCircuitBreakerexf_InsertClient(t *testing.T) {
	var c mgr.Client
	c.APIKey = "12233hgdd333"
	c.ClientID = clientIDexf
	c.Enabled = true
	c.Level = "small"

	res := gatewayDB2exf.InsertClient(&c)
	if res.Success == true && res.ID != -1 {
		fmt.Print("new client Id: ")
		fmt.Println(clientIDexf)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestCircuitBreakerexf_InsertRestRoute(t *testing.T) {
	var rr mgr.RestRoute
	rr.Route = "content"
	rr.ClientID = clientIDexf

	res := gatewayDB2exf.InsertRestRoute(&rr)
	if res.Success == true && res.ID != -1 {
		routeIDexf = res.ID
		fmt.Print("new route Id: ")
		fmt.Println(routeIDexf)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestCircuitBreakerexf_InsertRouteURL(t *testing.T) {
	var ru mgr.RouteURL
	ru.Name = "blue"
	ru.URL = "http://www.apigateway.com/blue/"
	ru.Active = false
	ru.RouteID = routeIDexf
	ru.ClientID = clientIDexf

	res := gatewayDB2exf.InsertRouteURL(&ru)
	if res.Success == true && res.ID != -1 {
		routeURLIDexf = res.ID
		fmt.Print("new route url Id: ")
		fmt.Println(routeURLIDexf)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestCircuitBreakerexf_InsertBreaker(t *testing.T) {
	var b Breaker
	b.ClientID = clientIDexf
	b.FailureThreshold = 5
	b.HealthCheckTimeSeconds = 120
	b.FailoverRouteName = "blue"
	b.OpenFailCode = 500
	b.RestRouteID = routeIDexf
	b.RouteURIID = routeURLIDexf
	suc, err := gatewayDBexf.InsertBreaker(&b)
	if suc != true || err != nil {
		t.Fail()
	}
}

var bidexf int64

func TestCircuitBreakerexf_GetRouteBreaker(t *testing.T) {
	var b Breaker
	b.ClientID = clientIDexf
	b.RestRouteID = routeIDexf
	b.RouteURIID = routeURLIDexf
	res := gatewayDBexf.GetBreaker(&b)
	fmt.Println("")
	fmt.Print("found breaker: ")
	fmt.Println(res)
	bidexf = res.ID
	if res.FailureThreshold != 5 {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestCircuitBreakerexf_UpdateRouteBreaker(t *testing.T) {
	var b Breaker
	b.ClientID = clientIDexf
	b.ID = bidexf
	b.FailureThreshold = 3
	b.HealthCheckTimeSeconds = 2
	b.FailoverRouteName = "green"
	b.OpenFailCode = 400
	b.RestRouteID = routeIDexf
	b.RouteURIID = routeURLIDexf

	suc, err := gatewayDBexf.UpdateBreaker(&b)
	if suc != true {
		fmt.Println(err)
		t.Fail()
	}
}

var thebreakerexf *Breaker

func TestCircuitBreakerexf_GetRouteBreaker2(t *testing.T) {
	var b Breaker
	b.ClientID = clientIDexf
	b.RestRouteID = routeIDexf
	b.RouteURIID = routeURLIDexf
	res := gatewayDBexf.GetBreaker(&b)
	thebreakerexf = res
	fmt.Println("")
	fmt.Print("found breaker: ")
	fmt.Println(res)
	if res.FailureThreshold != 3 {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestCircuitBreakerexf_GetStatus(t *testing.T) {
	res := gatewayDBexf.GetStatus(clientIDexf, routeURLIDexf)
	fmt.Println("")
	fmt.Print("found breaker status: ")
	fmt.Println(res)
	if res.Warning != false {
		fmt.Println("status failed")
		t.Fail()
	}
}

func TestCircuitBreakerexf_Trip(t *testing.T) {
	gatewayDBexf.Trip(thebreakerexf)
	res := gatewayDBexf.GetStatus(clientIDexf, routeURLIDexf)
	fmt.Println(res)
	if res.Open == true || res.Warning != true || res.PartialOpen != true || res.FailoverRouteName != "" {
		fmt.Println("circuit breaker should be partially open")
		t.Fail()
	}
}

func TestCircuitBreakerexf_Trip2(t *testing.T) {
	gatewayDBexf.Trip(thebreakerexf)
	res := gatewayDBexf.GetStatus(clientIDexf, routeURLIDexf)
	fmt.Println(res)
	if res.Open == true || res.Warning != true || res.PartialOpen != true || res.FailoverRouteName != "" {
		fmt.Println("circuit breaker should be partially open")
		t.Fail()
	}
}

func TestCircuitBreakerexf_Trip3(t *testing.T) {
	gatewayDBexf.Trip(thebreakerexf)
	res := gatewayDBexf.GetStatus(clientIDexf, routeURLIDexf)
	fmt.Println(res)
	if res.Open != true || res.Warning != true || res.PartialOpen == true || res.FailoverRouteName != "green" {
		fmt.Println("circuit breaker should be open")
		t.Fail()
	}
	time.Sleep(time.Second * 3)
}

func TestCircuitBreakerexf_GetStatus2(t *testing.T) {
	res := gatewayDBexf.GetStatus(clientIDexf, routeURLIDexf)
	fmt.Println("")
	fmt.Print("found breaker status: ")
	fmt.Println(res)
	if res.Open == true || res.Warning != true || res.PartialOpen != true || res.FailoverRouteName != "" {
		fmt.Println("circuit breaker should be partially open")
		t.Fail()
	}
}

func TestCircuitBreakerexf_Trip4(t *testing.T) {
	gatewayDBexf.Trip(thebreakerexf)
	res := gatewayDBexf.GetStatus(clientIDexf, routeURLIDexf)
	fmt.Println(res)
	if res.Open != true || res.Warning != true || res.PartialOpen == true || res.FailoverRouteName != "green" {
		fmt.Println("circuit breaker should be open")
		t.Fail()
	}
}

func TestCircuitBreakerexf_Reset(t *testing.T) {
	gatewayDBexf.Reset(clientIDexf, routeURLIDexf)
	res := gatewayDBexf.GetStatus(clientIDexf, routeURLIDexf)
	fmt.Println(res)
	if res.Open == true || res.Warning == true || res.PartialOpen == true || res.FailoverRouteName != "" {
		fmt.Println("circuit breaker should be closed")
		t.Fail()
	}
}

func TestCircuitBreakerexf_DeleteBreaker(t *testing.T) {
	var b Breaker
	b.ClientID = clientIDexf
	b.RestRouteID = routeIDexf
	b.RouteURIID = routeURLIDexf
	res := gatewayDBexf.DeleteBreaker(&b)
	if res != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestCircuitBreakerexf_DeleteClient(t *testing.T) {
	var c mgr.Client
	c.ClientID = clientIDexf
	res := gatewayDB2exf.DeleteClient(&c)
	if res.Success != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestCircuitBreakerexf_TestCloseDb(t *testing.T) {
	success := gatewayDBexf.CloseDb()
	success2 := gatewayDB2exf.CloseDb()
	if success != true || success2 != true {
		t.Fail()
	}
}
