package monitor

import (
	mgr "UlboraApiGateway/managers"
	"fmt"
	"testing"
	"time"
)

var gatewayDB GatewayPerformanceMonitor
var gatewayDB2 mgr.GatewayDB
var connected1 bool
var connected2 bool
var clientIDT int64
var insertID int64

var routeIDT int64

var routeURLIDT int64

func TestGatewayPerformanceMonitor_ConnectDb(t *testing.T) {
	clientIDT = 4334567
	gatewayDB.DbConfig.Host = "localhost:3306"
	gatewayDB.DbConfig.DbUser = "admin"
	gatewayDB.DbConfig.DbPw = "admin"
	gatewayDB.DbConfig.DatabaseName = "ulbora_api_gateway"
	gatewayDB.CallBatchSize = 2
	gatewayDB.CacheHost = "http://localhost:3010"
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

func TestGatewayPerformanceMonitor_InsertClient(t *testing.T) {
	var c mgr.Client
	c.APIKey = "12233hgdd333"
	c.ClientID = clientIDT
	c.Enabled = true
	c.Level = "small"

	res := gatewayDB2.InsertClient(&c)
	if res.Success == true && res.ID != -1 {
		fmt.Print("new client Id: ")
		fmt.Println(clientIDT)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_InsertRestRoute(t *testing.T) {
	var rr mgr.RestRoute
	rr.Route = "content"
	rr.ClientID = clientIDT

	res := gatewayDB2.InsertRestRoute(&rr)
	if res.Success == true && res.ID != -1 {
		routeIDT = res.ID
		fmt.Print("new route Id: ")
		fmt.Println(routeIDT)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_InsertRouteURL(t *testing.T) {
	var ru mgr.RouteURL
	ru.Name = "blue"
	ru.URL = "http://www.apigateway.com/blue/"
	ru.Active = false
	ru.RouteID = routeIDT
	ru.ClientID = clientIDT

	res := gatewayDB2.InsertRouteURL(&ru)
	if res.Success == true && res.ID != -1 {
		routeURLIDT = res.ID
		fmt.Print("new route url Id: ")
		fmt.Println(routeURLIDT)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

// func TestGatewayPerformanceMonitor_InsertRoutePerformanceCache(t *testing.T) {
// 	var p GwPerformance
// 	p.ClientID = clientID
// 	p.Calls = 500
// 	p.Entered = time.Now().Add(time.Hour * -2400)
// 	p.LatencyMsTotal = 10000
// 	p.RestRouteID = routeID
// 	p.RouteURIID = routeURLID
// 	suc, err := gatewayDB.InsertRoutePerformance(&p)
// 	if suc == true || err == nil {
// 		t.Fail()
// 	}
// }

func TestGatewayPerformanceMonitor_InsertRoutePerformanceReq(t *testing.T) {
	gatewayDB.CacheHost = "http://localhost:3010"
	var p GwPerformance
	//p.ClientID = clientID
	p.Calls = 500
	p.Entered = time.Now().Add(time.Hour * -2400)
	p.LatencyMsTotal = 10000
	p.RestRouteID = routeIDT
	p.RouteURIID = routeURLIDT
	suc, err := gatewayDB.InsertRoutePerformance(&p)
	if suc == true || err == nil {
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_TestCloseDb1(t *testing.T) {
	success := gatewayDB.CloseDb()
	//success2 := gatewayDB2.CloseDb()
	if success != true {
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_InsertRoutePerformanceDb(t *testing.T) {
	var p GwPerformance
	p.ClientID = clientIDT
	p.Calls = 500
	p.Entered = time.Now().Add(time.Hour * -2400)
	p.LatencyMsTotal = 10000
	p.RestRouteID = routeIDT
	p.RouteURIID = routeURLIDT
	suc, err := gatewayDB.InsertRoutePerformance(&p)
	if suc != true || err != nil {
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_InsertRoutePerformance(t *testing.T) {
	var p GwPerformance
	p.ClientID = clientIDT
	p.Calls = 500
	p.Entered = time.Now().Add(time.Hour * -2400)
	p.LatencyMsTotal = 10000
	p.RestRouteID = routeIDT
	p.RouteURIID = routeURLIDT
	suc, err := gatewayDB.InsertRoutePerformance(&p)
	if suc != true || err != nil {
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_GetRoutePerformance(t *testing.T) {
	var p GwPerformance
	p.ClientID = clientIDT
	p.RestRouteID = routeIDT
	p.RouteURIID = routeURLIDT
	res := gatewayDB.GetRoutePerformance(&p)
	fmt.Println("")
	fmt.Print("found gw performance list: ")
	fmt.Println(res)
	if len(*res) == 0 {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_DeleteRoutePerformance(t *testing.T) {
	res := gatewayDB.DeleteRoutePerformance()
	if res != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_SaveRoutePerformanceCache(t *testing.T) {
	gatewayDB.CacheHost = ""
	res := gatewayDB.SaveRoutePerformance(clientIDT, routeIDT, routeURLIDT, 100)
	if res != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_SaveRoutePerformanceReq(t *testing.T) {
	gatewayDB.CacheHost = "http://localhost:3010"
	var cid int64
	var rte int64
	fmt.Println("sending bad request:")
	res := gatewayDB.SaveRoutePerformance(cid, rte, routeURLIDT, 100)
	if res != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_SaveRoutePerformance(t *testing.T) {
	res := gatewayDB.SaveRoutePerformance(clientIDT, routeIDT, routeURLIDT, 100)
	if res != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_SaveRoutePerformanceBatchSize(t *testing.T) {
	gatewayDB.CallBatchSize = 1000
	res := gatewayDB.SaveRoutePerformance(clientIDT, routeIDT, routeURLIDT, 100)
	if res != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_SaveRoutePerformanceCh(t *testing.T) {
	gatewayDB.CacheHost = "htt"
	res := gatewayDB.SaveRoutePerformance(clientIDT, routeIDT, routeURLIDT, 100)
	if res != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_SaveRoutePerformance3(t *testing.T) {
	gatewayDB.CallBatchSize = 0
	res := gatewayDB.SaveRoutePerformance(clientIDT, routeIDT, routeURLIDT, 100)
	if res != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_GetRoutePerformance2(t *testing.T) {
	var p GwPerformance
	p.ClientID = clientIDT
	p.RestRouteID = routeIDT
	p.RouteURIID = routeURLIDT
	res := gatewayDB.GetRoutePerformance(&p)
	fmt.Println("")
	fmt.Print("found gw performance list: ")
	fmt.Println(res)
	if len(*res) == 0 {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_DeleteClient(t *testing.T) {
	var c mgr.Client
	c.ClientID = clientIDT
	res := gatewayDB2.DeleteClient(&c)
	if res.Success != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_TestCloseDb(t *testing.T) {
	success := gatewayDB.CloseDb()
	success2 := gatewayDB2.CloseDb()
	if success != true || success2 != true {
		t.Fail()
	}
}
