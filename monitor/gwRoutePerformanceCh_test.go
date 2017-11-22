package monitor

import (
	mgr "UlboraApiGateway/managers"
	"fmt"
	"testing"
	"time"
)

var gatewayDBc GatewayPerformanceMonitor
var gatewayDB2c mgr.GatewayDB
var connected1c bool
var connected2c bool
var clientIDc int64
var insertIDc int64

var routeIDc int64

var routeURLIDc int64

func TestGatewayPerformanceMonitorc_ConnectDb(t *testing.T) {
	clientIDc = 433456766777
	gatewayDBc.DbConfig.Host = "localhost:3306"
	gatewayDBc.DbConfig.DbUser = "admin"
	gatewayDBc.DbConfig.DbPw = "admin"
	gatewayDBc.DbConfig.DatabaseName = "ulbora_api_gateway"
	gatewayDBc.CallBatchSize = 2
	gatewayDBc.CacheHost = "http://localhost:3010"
	connected1c = gatewayDBc.ConnectDb()

	gatewayDB2c.DbConfig.Host = "localhost:3306"
	gatewayDB2c.DbConfig.DbUser = "admin"
	gatewayDB2c.DbConfig.DbPw = "admin"
	gatewayDB2c.DbConfig.DatabaseName = "ulbora_api_gateway"
	connected2c = gatewayDB2c.ConnectDb()
	if connected1c != true || connected2c != true {
		t.Fail()
	}
}

func TestGatewayPerformanceMonitorc_InsertClient(t *testing.T) {
	var c mgr.Client
	c.APIKey = "12233hgdd333"
	c.ClientID = clientIDc
	c.Enabled = true
	c.Level = "small"

	res := gatewayDB2c.InsertClient(&c)
	if res.Success == true && res.ID != -1 {
		fmt.Print("new client Id: ")
		fmt.Println(clientIDc)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitorc_InsertRestRoute(t *testing.T) {
	var rr mgr.RestRoute
	rr.Route = "content"
	rr.ClientID = clientIDc

	res := gatewayDB2c.InsertRestRoute(&rr)
	if res.Success == true && res.ID != -1 {
		routeIDc = res.ID
		fmt.Print("new route Id: ")
		fmt.Println(routeIDc)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitorc_InsertRouteURL(t *testing.T) {
	var ru mgr.RouteURL
	ru.Name = "blue"
	ru.URL = "http://www.apigateway.com/blue/"
	ru.Active = false
	ru.RouteID = routeIDc
	ru.ClientID = clientIDc

	res := gatewayDB2c.InsertRouteURL(&ru)
	if res.Success == true && res.ID != -1 {
		routeURLIDc = res.ID
		fmt.Print("new route url Id: ")
		fmt.Println(routeURLIDc)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitorc_InsertRoutePerformance(t *testing.T) {
	var p GwPerformance
	p.ClientID = clientIDc
	p.Calls = 500
	p.Entered = time.Now().Add(time.Hour * -2400)
	p.LatencyMsTotal = 10000
	p.RestRouteID = routeIDc
	p.RouteURIID = routeURLIDc
	suc, err := gatewayDBc.InsertRoutePerformance(&p)
	if suc != true || err != nil {
		t.Fail()
	}
}

func TestGatewayPerformanceMonitorc_GetRoutePerformance(t *testing.T) {
	var p GwPerformance
	p.ClientID = clientIDc
	p.RestRouteID = routeIDc
	p.RouteURIID = routeURLIDc
	res := gatewayDBc.GetRoutePerformance(&p)
	fmt.Println("")
	fmt.Print("found gw performance list: ")
	fmt.Println(res)
	if len(*res) == 0 {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitorc_DeleteRoutePerformance(t *testing.T) {
	res := gatewayDBc.DeleteRoutePerformance()
	if res != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitorc_SaveRoutePerformance(t *testing.T) {
	res := gatewayDBc.SaveRoutePerformance(clientIDc, routeIDc, routeURLIDc, 100)
	if res != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitorc_SaveRoutePerformance2(t *testing.T) {
	res := gatewayDBc.SaveRoutePerformance(clientIDc, routeIDc, routeURLIDc, 100)
	if res != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitorc_SaveRoutePerformance3(t *testing.T) {
	res := gatewayDBc.SaveRoutePerformance(clientIDc, routeIDc, routeURLIDc, 100)
	if res != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitorc_GetRoutePerformance2(t *testing.T) {
	var p GwPerformance
	p.ClientID = clientIDc
	p.RestRouteID = routeIDc
	p.RouteURIID = routeURLIDc
	res := gatewayDBc.GetRoutePerformance(&p)
	fmt.Println("")
	fmt.Print("found gw performance list: ")
	fmt.Println(res)
	if len(*res) == 0 {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitorc_DeleteClient(t *testing.T) {
	var c mgr.Client
	c.ClientID = clientIDc
	res := gatewayDB2c.DeleteClient(&c)
	if res.Success != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitorc_TestCloseDb(t *testing.T) {
	success := gatewayDBc.CloseDb()
	success2 := gatewayDB2c.CloseDb()
	if success != true || success2 != true {
		t.Fail()
	}
}
