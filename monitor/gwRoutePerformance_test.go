package monitor

import (
	mgr "UlboraApiGateway/managers"
	"fmt"
	"testing"
	"time"
)

var gatewayDBPer GatewayPerformanceMonitor
var gatewayDB2Per mgr.GatewayDB
var connected1Per bool
var connected2Per bool
var clientIDTPer int64
var insertIDPer int64

var routeIDTPer int64

var routeURLIDTPer int64

func TestGatewayPerformanceMonitor_ConnectDb(t *testing.T) {
	clientIDTPer = 4334567
	gatewayDBPer.DbConfig.Host = "localhost:3306"
	gatewayDBPer.DbConfig.DbUser = "admin"
	gatewayDBPer.DbConfig.DbPw = "admin"
	gatewayDBPer.DbConfig.DatabaseName = "ulbora_api_gateway"
	gatewayDBPer.CallBatchSize = 2
	gatewayDBPer.CacheHost = "http://localhost:3010"
	connected1Per = gatewayDBPer.ConnectDb()

	gatewayDB2Per.DbConfig.Host = "localhost:3306"
	gatewayDB2Per.DbConfig.DbUser = "admin"
	gatewayDB2Per.DbConfig.DbPw = "admin"
	gatewayDB2Per.DbConfig.DatabaseName = "ulbora_api_gateway"
	connected2Per = gatewayDB2Per.ConnectDb()
	if connected1Per != true || connected2Per != true {
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_InsertClient(t *testing.T) {
	var c mgr.Client
	c.APIKey = "12233hgdd333"
	c.ClientID = clientIDTPer
	c.Enabled = true
	c.Level = "small"

	res := gatewayDB2Per.InsertClient(&c)
	if res.Success == true && res.ID != -1 {
		fmt.Print("new client Id: ")
		fmt.Println(clientIDTPer)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_InsertRestRoute(t *testing.T) {
	var rr mgr.RestRoute
	rr.Route = "content"
	rr.ClientID = clientIDTPer

	res := gatewayDB2Per.InsertRestRoute(&rr)
	if res.Success == true && res.ID != -1 {
		routeIDTPer = res.ID
		fmt.Print("new route Id: ")
		fmt.Println(routeIDTPer)
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
	ru.RouteID = routeIDTPer
	ru.ClientID = clientIDTPer

	res := gatewayDB2Per.InsertRouteURL(&ru)
	if res.Success == true && res.ID != -1 {
		routeURLIDTPer = res.ID
		fmt.Print("new route url Id: ")
		fmt.Println(routeURLIDTPer)
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
	gatewayDBPer.CacheHost = "http://localhost:3010"
	var p GwPerformance
	//p.ClientID = clientID
	p.Calls = 500
	p.Entered = time.Now().Add(time.Hour * -2400)
	p.LatencyMsTotal = 10000
	p.RestRouteID = routeIDTPer
	p.RouteURIID = routeURLIDTPer
	suc, err := gatewayDBPer.InsertRoutePerformance(&p)
	if suc == true || err == nil {
		t.Fail()
	}
}

// func TestGatewayPerformanceMonitor_TestCloseDb1(t *testing.T) {
// 	success := gatewayDB.CloseDb()
// 	//success2 := gatewayDB2.CloseDb()
// 	if success != true {
// 		t.Fail()
// 	}
// }

func TestGatewayPerformanceMonitor_InsertRoutePerformanceDb(t *testing.T) {
	gatewayDBPer.CloseDb()
	var p GwPerformance
	p.ClientID = clientIDTPer
	p.Calls = 500
	p.Entered = time.Now().Add(time.Hour * -2400)
	p.LatencyMsTotal = 10000
	p.RestRouteID = routeIDTPer
	p.RouteURIID = routeURLIDTPer
	suc, err := gatewayDBPer.InsertRoutePerformance(&p)
	if suc != true || err != nil {
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_InsertRoutePerformance(t *testing.T) {
	var p GwPerformance
	p.ClientID = clientIDTPer
	p.Calls = 500
	p.Entered = time.Now().Add(time.Hour * -2400)
	p.LatencyMsTotal = 10000
	p.RestRouteID = routeIDTPer
	p.RouteURIID = routeURLIDTPer
	suc, err := gatewayDBPer.InsertRoutePerformance(&p)
	if suc != true || err != nil {
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_GetRoutePerformance(t *testing.T) {
	var p GwPerformance
	p.ClientID = clientIDTPer
	p.RestRouteID = routeIDTPer
	p.RouteURIID = routeURLIDTPer
	res := gatewayDBPer.GetRoutePerformance(&p)
	fmt.Println("")
	fmt.Print("found gw performance list: ")
	fmt.Println(res)
	if len(*res) == 0 {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_DeleteRoutePerformance(t *testing.T) {
	res := gatewayDBPer.DeleteRoutePerformance()
	if res != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_SaveRoutePerformanceCache(t *testing.T) {
	gatewayDBPer.CacheHost = ""
	res := gatewayDBPer.SaveRoutePerformance(clientIDTPer, routeIDTPer, routeURLIDTPer, 100)
	if res != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_SaveRoutePerformanceReq(t *testing.T) {
	gatewayDBPer.CacheHost = "http://localhost:3010"
	var cid int64
	var rte int64
	fmt.Println("sending bad request:")
	res := gatewayDBPer.SaveRoutePerformance(cid, rte, routeURLIDTPer, 100)
	if res != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_SaveRoutePerformance(t *testing.T) {
	res := gatewayDBPer.SaveRoutePerformance(clientIDTPer, routeIDTPer, routeURLIDTPer, 100)
	if res != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_SaveRoutePerformanceBatchSize(t *testing.T) {
	gatewayDBPer.CallBatchSize = 1000
	res := gatewayDBPer.SaveRoutePerformance(clientIDTPer, routeIDTPer, routeURLIDTPer, 100)
	if res != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_SaveRoutePerformanceCh(t *testing.T) {
	gatewayDBPer.CacheHost = "htt"
	res := gatewayDBPer.SaveRoutePerformance(clientIDTPer, routeIDTPer, routeURLIDTPer, 100)
	if res != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_SaveRoutePerformance3(t *testing.T) {
	gatewayDBPer.CallBatchSize = 0
	res := gatewayDBPer.SaveRoutePerformance(clientIDTPer, routeIDTPer, routeURLIDTPer, 100)
	if res != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_GetRoutePerformance2(t *testing.T) {
	var p GwPerformance
	p.ClientID = clientIDTPer
	p.RestRouteID = routeIDTPer
	p.RouteURIID = routeURLIDTPer
	res := gatewayDBPer.GetRoutePerformance(&p)
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
	c.ClientID = clientIDTPer
	res := gatewayDB2Per.DeleteClient(&c)
	if res.Success != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestGatewayPerformanceMonitor_TestCloseDb(t *testing.T) {
	success := gatewayDBPer.CloseDb()
	success2 := gatewayDB2Per.CloseDb()
	if success != true || success2 != true {
		t.Fail()
	}
}
