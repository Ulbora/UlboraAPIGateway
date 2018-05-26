package database

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var dbConfigdb DbConfig
var connecteddb bool
var clientIDdb int64
var insertIDdb int64
var routeIDdb int64
var routeURLIDdb int64
var routeURLID2db int64

func TestDbConfig_ConnectDb(t *testing.T) {
	//time.Sleep(time.Second * 10)
	//var dbConfig DbConfig
	dbConfigdb.Host = "localhost:3306"
	dbConfigdb.DbUser = "admin"
	dbConfigdb.DbPw = "admin"
	dbConfigdb.DatabaseName = "ulbora_api_gateway"
	connecteddb = dbConfigdb.ConnectDb()
	if connecteddb != true {
		t.Fail()
	} else {
		fmt.Println("database opened in database package")
	}
}

func TestDbConfig_ConnecTest(t *testing.T) {
	clientIDdb = 333333
	//var a []interface{}
	success := dbConfigdb.ConnectionTest()
	if success == true {
		fmt.Print("Connection test: ")
		fmt.Println(success)
	} else {
		fmt.Println("database connection test failed")
		t.Fail()
	}
}

func TestDbConfig_InsertClient(t *testing.T) {
	var a []interface{}
	a = append(a, clientIDdb, "23ddeee", true, "small")
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success, insID := dbConfigdb.InsertClient(a...)
	if success == true && insID != -1 {
		insertIDdb = clientIDdb
		fmt.Print("new Id: ")
		fmt.Println(insID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestDbConfig_UpdateClient(t *testing.T) {
	var a []interface{}
	a = append(a, "23ddeee", true, "medium", clientIDdb)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success := dbConfigdb.UpdateClient(a...)
	if success != true {
		fmt.Println("database update failed")
		t.Fail()
	}
}

func TestDbConfig_GetClient(t *testing.T) {
	a := []interface{}{clientIDdb}
	rowPtr := dbConfigdb.GetClient(a...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		fmt.Print("Get ")
		fmt.Println(foundRow)
		//fmt.Println("Get results: --------------------------")
		int64Val, err2 := strconv.ParseInt(foundRow[0], 10, 0)
		if err2 != nil {
			fmt.Print(err2)
		}
		if insertIDdb != int64Val {
			fmt.Print(insertIDdb)
			fmt.Print(" != ")
			fmt.Println(int64Val)
			t.Fail()
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestDbConfig_GetClientList(t *testing.T) {
	a := []interface{}{}
	rowsPtr := dbConfigdb.GetClientList(a...)
	if rowsPtr != nil {
		foundRows := rowsPtr.Rows
		fmt.Print("Get clients list ")
		fmt.Println(foundRows)
		//fmt.Println("GetList results: --------------------------")
		if len(foundRows) == 0 {
			t.Fail()
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestDbConfig_InsertRestRoute(t *testing.T) {
	var a []interface{}
	a = append(a, "mail", clientIDdb)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success, insID := dbConfigdb.InsertRestRoute(a...)
	if success == true && insID != -1 {
		routeIDdb = insID
		fmt.Print("new Id route: ")
		fmt.Println(insID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestDbConfig_UpdateRestRoute(t *testing.T) {
	var a []interface{}
	a = append(a, "mail2", routeIDdb, clientIDdb)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success := dbConfigdb.UpdateRestRoute(a...)
	if success != true {
		fmt.Println("database update failed")
		t.Fail()
	}
}

func TestDbConfig_GetRestRoute(t *testing.T) {
	a := []interface{}{routeIDdb, clientIDdb}
	rowPtr := dbConfigdb.GetRestRoute(a...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		fmt.Print("Get rest route")
		fmt.Println(foundRow)
		//fmt.Println("Get results: --------------------------")
		int64Val, err2 := strconv.ParseInt(foundRow[0], 10, 0)
		if err2 != nil {
			fmt.Print(err2)
		}
		if routeIDdb != int64Val {
			fmt.Print(insertIDdb)
			fmt.Print(" != ")
			fmt.Println(int64Val)
			t.Fail()
		}
		if foundRow[1] != "mail2" {
			t.Fail()
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestDbConfig_GetRestRouteList(t *testing.T) {
	a := []interface{}{clientIDdb}
	rowsPtr := dbConfigdb.GetRestRouteList(a...)
	if rowsPtr != nil {
		foundRows := rowsPtr.Rows
		fmt.Print("Get rest route list ")
		fmt.Println(foundRows)
		//fmt.Println("GetList results: --------------------------")
		if len(foundRows) == 0 {
			t.Fail()
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
}
func TestDbConfig_InsertRouteURL(t *testing.T) {
	var a []interface{}
	a = append(a, "blue", "http://www.apigateway.com/blue/", true, routeIDdb, clientIDdb)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success, insID := dbConfigdb.InsertRouteURL(a...)
	if success == true && insID != -1 {
		routeURLIDdb = insID
		fmt.Print("new URL Id: ")
		fmt.Println(insID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}

	var a2 []interface{}
	a2 = append(a2, "bside", "http://www.apigateway.com/green/", true, routeIDdb, clientIDdb)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success2, insID2 := dbConfigdb.InsertRouteURL(a2...)
	if success2 == true && insID2 != -1 {
		routeURLID2db = insID2
		fmt.Print("new URL Id: ")
		fmt.Println(insID2)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestDbConfig_UpdateRouteURL(t *testing.T) {
	var a []interface{}
	a = append(a, "green", "http://www.apigateway.com/green/", routeURLID2db, routeIDdb, clientIDdb)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success := dbConfigdb.UpdateRouteURL(a...)
	if success != true {
		fmt.Println("database update failed")
		t.Fail()
	}
}

func TestDbConfig_ActivateRouteURL(t *testing.T) {
	var a []interface{}
	a = append(a, routeURLID2db, routeIDdb, clientIDdb)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success := dbConfigdb.ActivateRouteURL(a...)
	if success != true {
		fmt.Println("database update failed")
		t.Fail()
	}

	var a2 []interface{}
	a2 = append(a2, routeURLID2db, routeIDdb, clientIDdb)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success2 := dbConfigdb.DeactivateOtherRouteURLs(a2...)
	if success2 != true {
		fmt.Println("database update failed")
		t.Fail()
	}
}

func TestDbConfig_GetRouteURL(t *testing.T) {
	a := []interface{}{routeURLID2db, routeIDdb, clientIDdb}
	rowPtr := dbConfigdb.GetRouteURL(a...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		fmt.Print("Get ")
		fmt.Println(foundRow)
		//fmt.Println("Get results: --------------------------")
		int64Val, err2 := strconv.ParseInt(foundRow[0], 10, 0)
		if err2 != nil {
			fmt.Print(err2)
		}
		if routeURLID2db != int64Val {
			fmt.Print(routeURLID2db)
			fmt.Print(" != ")
			fmt.Println(int64Val)
			t.Fail()
		}
		active, err3 := strconv.ParseBool(foundRow[3])
		if err3 != nil {
			fmt.Print(err3)
		}
		if active != true {
			t.Fail()
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestDbConfig_GetRouteURL2(t *testing.T) {
	a := []interface{}{routeURLIDdb, routeIDdb, clientIDdb}
	rowPtr := dbConfigdb.GetRouteURL(a...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		fmt.Print("Get ")
		fmt.Println(foundRow)
		//fmt.Println("Get results: --------------------------")
		int64Val, err2 := strconv.ParseInt(foundRow[0], 10, 0)
		if err2 != nil {
			fmt.Print(err2)
		}
		if routeURLIDdb != int64Val {
			fmt.Print(routeURLIDdb)
			fmt.Print(" != ")
			fmt.Println(int64Val)
			t.Fail()
		}
		active, err3 := strconv.ParseBool(foundRow[3])
		if err3 != nil {
			fmt.Print(err3)
		}
		if active == true {
			t.Fail()
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestDbConfig_GetRouteURLList(t *testing.T) {
	a := []interface{}{routeIDdb, clientIDdb}
	rowsPtr := dbConfigdb.GetRouteURLList(a...)
	if rowsPtr != nil {
		foundRows := rowsPtr.Rows
		fmt.Print("Get route urls ")
		fmt.Println(foundRows)
		//fmt.Println("GetList results: --------------------------")
		if len(foundRows) != 2 {
			t.Fail()
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestDbConfig_GetRouteNameURLList(t *testing.T) {
	a := []interface{}{"mail2", clientIDdb, "23ddeee"}
	rowsPtr := dbConfigdb.GetRouteNameURLList(a...)
	if rowsPtr != nil {
		foundRows := rowsPtr.Rows
		fmt.Print("Get route name urls ")
		fmt.Println(foundRows)
		//fmt.Println("GetList results: --------------------------")
		if len(foundRows) != 2 {
			t.Fail()
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestDbConfig_GetRouteURLs(t *testing.T) {
	a := []interface{}{clientIDdb, "mail2"}
	rowsPtr := dbConfigdb.GetRouteURLs(a...)
	if rowsPtr != nil {
		foundRows := rowsPtr.Rows
		fmt.Print("Get route urls ")
		fmt.Println(foundRows)
		//fmt.Println("GetList results: --------------------------")
		if len(foundRows) != 2 {
			t.Fail()
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestDbConfig_InsertRoutePerformance(t *testing.T) {
	a := []interface{}{100, 5000, time.Now().Add(time.Hour * -2400), routeURLIDdb, routeIDdb, clientIDdb}
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success, insID := dbConfigdb.InsertRoutePerformance(a...)
	if success == true && insID != -1 {
		fmt.Print("new Id performance: ")
		fmt.Println(insID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestDbConfig_GetRoutePerformance(t *testing.T) {
	a := []interface{}{routeURLIDdb, routeIDdb, clientIDdb}
	rowsPtr := dbConfigdb.GetRoutePerformance(a...)
	if rowsPtr != nil {
		foundRows := rowsPtr.Rows
		fmt.Print("Get route performance ")
		fmt.Println(foundRows)
		if len(foundRows) == 0 {
			t.Fail()
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestDbConfig_DeleteRoutePerformance(t *testing.T) {
	a := []interface{}{}
	success := dbConfigdb.DeleteRoutePerformance(a...)
	if success == true {
		fmt.Print("Deleted route performance: ")
		fmt.Println(routeURLIDdb)
	} else {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestDbConfig_InsertRouteError(t *testing.T) {
	a := []interface{}{404, "error call failed", time.Now().Add(time.Hour * -2400), routeURLIDdb, routeIDdb, clientIDdb}
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success, insID := dbConfigdb.InsertRouteError(a...)
	if success == true && insID != -1 {
		fmt.Print("new Id route error id: ")
		fmt.Println(insID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestDbConfig_GetRouteError(t *testing.T) {
	a := []interface{}{routeURLIDdb, routeIDdb, clientIDdb}
	rowsPtr := dbConfigdb.GetRouteError(a...)
	if rowsPtr != nil {
		foundRows := rowsPtr.Rows
		fmt.Print("Get route error ")
		fmt.Println(foundRows)
		if len(foundRows) == 0 {
			t.Fail()
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestDbConfig_DeleteRouteError(t *testing.T) {
	a := []interface{}{}
	success := dbConfigdb.DeleteRouteError(a...)
	if success == true {
		fmt.Print("Deleted route error: ")
		fmt.Println(routeURLIDdb)
	} else {
		fmt.Println("database delete failed")
		//t.Fail()
	}
}

var brkID int64

func TestDbConfig_InsertRouteBreaker(t *testing.T) {
	a := []interface{}{3, 500, "mail", 400, routeURLIDdb, routeIDdb, clientIDdb}
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success, insID := dbConfigdb.InsertRouteBreaker(a...)
	if success == true && insID != -1 {
		brkID = insID
		fmt.Print("new Id route breaker id: ")
		fmt.Println(insID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestDbConfig_UpdateRouteBreakerConfig(t *testing.T) {
	a := []interface{}{5, 400, "mailblue", 401, brkID, routeURLIDdb, routeIDdb, clientIDdb}
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success := dbConfigdb.UpdateRouteBreakerConfig(a...)
	if success != true {
		fmt.Println("database update failed")
		t.Fail()
	}
}

func TestDbConfig_UpdateRouteBreakerFail(t *testing.T) {
	a := []interface{}{1, time.Now(), brkID, routeURLIDdb, routeIDdb, clientIDdb}
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success := dbConfigdb.UpdateRouteBreakerFail(a...)
	if success != true {
		fmt.Println("database update failed")
		t.Fail()
	}
}

func TestDbConfig_GetBreaker(t *testing.T) {
	a := []interface{}{routeURLIDdb, routeIDdb, clientIDdb}
	rowPtr := dbConfigdb.GetBreaker(a...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		fmt.Print("Get route breaker ")
		fmt.Println(foundRow)
		code, _ := strconv.Atoi(foundRow[2])
		fallOvR := foundRow[3]
		failCnt, _ := strconv.Atoi(foundRow[5])
		if code != 400 && fallOvR != "mailblue" && failCnt != 1 {
			t.Fail()
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestDbConfig_DeleteBreaker(t *testing.T) {
	a := []interface{}{routeURLIDdb, routeIDdb, clientIDdb}
	success := dbConfigdb.DeleteBreaker(a...)
	if success == true {
		fmt.Print("Deleted route breaker: ")
		fmt.Println(routeURLIDdb)
	} else {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestDbConfig_DeleteRouteURL(t *testing.T) {
	a := []interface{}{routeURLIDdb, routeIDdb, clientIDdb}
	fmt.Println(a)
	success := dbConfigdb.DeleteRouteURL(a...)
	if success == true {
		fmt.Print("Deleted route url: ")
		fmt.Println(routeURLIDdb)
	} else {
		fmt.Print("Deleted failed for route url: ")
		fmt.Println(routeURLIDdb)
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestDbConfig_DeleteRestRoute(t *testing.T) {
	a := []interface{}{routeIDdb, clientIDdb}
	success := dbConfigdb.DeleteRestRoute(a...)
	if success == true {
		fmt.Print("Deleted route: ")
		fmt.Println(routeIDdb)
	} else {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestDbConfig_DeleteClient(t *testing.T) {
	a := []interface{}{clientIDdb}
	success := dbConfigdb.DeleteClient(a...)
	if success == true {
		fmt.Print("Deleted client ")
		fmt.Println(clientIDdb)
	} else {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestDbConfig_CloseDb(t *testing.T) {
	if connecteddb == true {
		rtn := dbConfigdb.CloseDb()
		if rtn != true {
			fmt.Println("database close failed")
			t.Fail()
		} else {
			fmt.Println("database close in database package")
		}
	}
}
