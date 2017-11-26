package circuitbreaker

import (
	db "UlboraApiGateway/database"
	"fmt"
	"testing"
	"time"
)

var gatewayDB CircuitBreaker
var dbt db.DbConfig

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

	dbt.Host = "localhost:3306"
	dbt.DbUser = "admin"
	dbt.DbPw = "admin"
	dbt.DatabaseName = "ulbora_api_gateway"
	connected2 = dbt.ConnectDb()
	if connected1 != true || connected2 != true {
		t.Fail()
	}

}

func TestCircuitBreaker_InsertClient(t *testing.T) {
	var a []interface{}
	a = append(a, clientID, "12233hgdd333", true, "small")
	success, _ := dbt.InsertClient(a...)
	if success == true {
		fmt.Println("inserted record")
	}
}

func TestCircuitBreaker_InsertRestRoute(t *testing.T) {
	var a []interface{}
	a = append(a, "content", clientID)
	success, insID := dbt.InsertRestRoute(a...)
	if success == true {
		fmt.Println("inserted record")
	}
	routeID = insID
}

func TestCircuitBreaker_InsertRouteURL(t *testing.T) {

	var a []interface{}
	a = append(a, "blue", "http://www.apigateway.com/blue/", false, routeID, clientID)
	success, insID := dbt.InsertRouteURL(a...)
	if success == true {
		fmt.Println("inserted record")
	}
	routeURLID = insID
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
	b.HealthCheckTimeSeconds = 2
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

var thebreaker *Breaker

func TestCircuitBreaker_GetRouteBreaker2(t *testing.T) {
	var b Breaker
	b.ClientID = clientID
	b.RestRouteID = routeID
	b.RouteURIID = routeURLID
	res := gatewayDB.GetBreaker(&b)
	thebreaker = res
	fmt.Println("")
	fmt.Print("found breaker: ")
	fmt.Println(res)
	if res.FailureThreshold != 3 {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestCircuitBreaker_GetStatus(t *testing.T) {
	res := gatewayDB.GetStatus(clientID, routeURLID)
	fmt.Println("")
	fmt.Print("found breaker status: ")
	fmt.Println(res)
	if res.Warning != false {
		fmt.Println("status failed")
		t.Fail()
	}
}

func TestCircuitBreaker_Trip(t *testing.T) {
	gatewayDB.Trip(thebreaker)
	res := gatewayDB.GetStatus(clientID, routeURLID)
	fmt.Println(res)
	if res.Open == true || res.Warning != true || res.PartialOpen != true || res.FailoverRouteName != "" {
		fmt.Println("circuit breaker should be partially open")
		t.Fail()
	}
}

func TestCircuitBreaker_Trip2(t *testing.T) {
	gatewayDB.Trip(thebreaker)
	res := gatewayDB.GetStatus(clientID, routeURLID)
	fmt.Println(res)
	if res.Open == true || res.Warning != true || res.PartialOpen != true || res.FailoverRouteName != "" {
		fmt.Println("circuit breaker should be partially open")
		t.Fail()
	}
}

func TestCircuitBreaker_Trip3(t *testing.T) {
	gatewayDB.Trip(thebreaker)
	res := gatewayDB.GetStatus(clientID, routeURLID)
	fmt.Println(res)
	if res.Open != true || res.Warning != true || res.PartialOpen == true || res.FailoverRouteName != "green" {
		fmt.Println("circuit breaker should be open")
		t.Fail()
	}
	time.Sleep(time.Second * 3)
}

func TestCircuitBreaker_GetStatus2(t *testing.T) {
	res := gatewayDB.GetStatus(clientID, routeURLID)
	fmt.Println("")
	fmt.Print("found breaker status: ")
	fmt.Println(res)
	if res.Open == true || res.Warning != true || res.PartialOpen != true || res.FailoverRouteName != "" {
		fmt.Println("circuit breaker should be partially open")
		t.Fail()
	}
}

func TestCircuitBreaker_Trip4(t *testing.T) {
	gatewayDB.Trip(thebreaker)
	res := gatewayDB.GetStatus(clientID, routeURLID)
	fmt.Println(res)
	if res.Open != true || res.Warning != true || res.PartialOpen == true || res.FailoverRouteName != "green" {
		fmt.Println("circuit breaker should be open")
		t.Fail()
	}
}

var badBreaker *Breaker

func TestCircuitBreaker_GetRouteBreakerBad(t *testing.T) {
	var b Breaker
	b.ClientID = clientID
	b.RestRouteID = 44444
	b.RouteURIID = 77777
	res := gatewayDB.GetBreaker(&b)
	badBreaker = res
	fmt.Println("")
	fmt.Print("found bad breaker: ")
	fmt.Println(res)
	bid = res.ID
	if res.ID != 0 {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestCircuitBreaker_TripBad(t *testing.T) {
	gatewayDB.Trip(badBreaker)
	res := gatewayDB.GetStatus(clientID, 77777)
	fmt.Println(res)
	if res.Open == true || res.Warning == true || res.PartialOpen == true || res.FailoverRouteName != "" {
		fmt.Println("bad circuit breaker should do nothing")
		t.Fail()
	}
}

func TestCircuitBreaker_Reset(t *testing.T) {
	gatewayDB.Reset(clientID, routeURLID)
	res := gatewayDB.GetStatus(clientID, routeURLID)
	fmt.Println("status of bad circuit breaker: ")
	fmt.Println(res)
	if res.Open == true || res.Warning == true || res.PartialOpen == true || res.FailoverRouteName != "" {
		fmt.Println("circuit breaker should be closed")
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
	var a []interface{}
	a = append(a, clientID)
	success := dbt.DeleteClient(a...)
	if success == true {
		fmt.Println("deleted record")
	}
}

func TestCircuitBreaker_TestCloseDb(t *testing.T) {
	success := gatewayDB.CloseDb()
	success2 := dbt.CloseDb()
	if success != true || success2 != true {
		//if success != true {
		t.Fail()
	}
}
