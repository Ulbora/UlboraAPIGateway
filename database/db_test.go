package database

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var dbConfig DbConfig
var connected bool
var clientID int64
var insertID int64
var routeID int64
var routeURLID int64
var routeURLID2 int64

func TestDbConfig_ConnectDb(t *testing.T) {
	//time.Sleep(time.Second * 10)
	//var dbConfig DbConfig
	dbConfig.Host = "localhost:3306"
	dbConfig.DbUser = "admin"
	dbConfig.DbPw = "admin"
	dbConfig.DatabaseName = "ulbora_api_gateway"
	connected = dbConfig.ConnectDb()
	if connected != true {
		t.Fail()
	} else {
		fmt.Println("database opened in database package")
	}
}

func TestDbConfig_ConnecTest(t *testing.T) {
	clientID = 333333
	var a []interface{}
	success := dbConfig.ConnectionTest(a...)
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
	a = append(a, clientID, "23ddeee", true, "small")
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success, insID := dbConfig.InsertClient(a...)
	if success == true && insID != -1 {
		insertID = clientID
		fmt.Print("new Id: ")
		fmt.Println(insID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestDbConfig_UpdateClient(t *testing.T) {
	var a []interface{}
	a = append(a, "23ddeee", true, "medium", clientID)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success := dbConfig.UpdateClient(a...)
	if success != true {
		fmt.Println("database update failed")
		t.Fail()
	}
}

func TestDbConfig_GetClient(t *testing.T) {
	a := []interface{}{clientID}
	rowPtr := dbConfig.GetClient(a...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		fmt.Print("Get ")
		fmt.Println(foundRow)
		//fmt.Println("Get results: --------------------------")
		int64Val, err2 := strconv.ParseInt(foundRow[0], 10, 0)
		if err2 != nil {
			fmt.Print(err2)
		}
		if insertID != int64Val {
			fmt.Print(insertID)
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
	rowsPtr := dbConfig.GetClientList(a...)
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
	a = append(a, "mail", clientID)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success, insID := dbConfig.InsertRestRoute(a...)
	if success == true && insID != -1 {
		routeID = insID
		fmt.Print("new Id route: ")
		fmt.Println(insID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestDbConfig_UpdateRestRoute(t *testing.T) {
	var a []interface{}
	a = append(a, "mail2", routeID, clientID)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success := dbConfig.UpdateRestRoute(a...)
	if success != true {
		fmt.Println("database update failed")
		t.Fail()
	}
}

func TestDbConfig_GetRestRoute(t *testing.T) {
	a := []interface{}{routeID, clientID}
	rowPtr := dbConfig.GetRestRoute(a...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		fmt.Print("Get rest route")
		fmt.Println(foundRow)
		//fmt.Println("Get results: --------------------------")
		int64Val, err2 := strconv.ParseInt(foundRow[0], 10, 0)
		if err2 != nil {
			fmt.Print(err2)
		}
		if routeID != int64Val {
			fmt.Print(insertID)
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
	a := []interface{}{clientID}
	rowsPtr := dbConfig.GetRestRouteList(a...)
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
	a = append(a, "blue", "http://www.apigateway.com/blue/", true, routeID, clientID)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success, insID := dbConfig.InsertRouteURL(a...)
	if success == true && insID != -1 {
		routeURLID = insID
		fmt.Print("new URL Id: ")
		fmt.Println(insID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}

	var a2 []interface{}
	a2 = append(a2, "bside", "http://www.apigateway.com/green/", true, routeID, clientID)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success2, insID2 := dbConfig.InsertRouteURL(a2...)
	if success2 == true && insID2 != -1 {
		routeURLID2 = insID2
		fmt.Print("new URL Id: ")
		fmt.Println(insID2)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestDbConfig_UpdateRouteURL(t *testing.T) {
	var a []interface{}
	a = append(a, "green", "http://www.apigateway.com/green/", routeURLID2, routeID, clientID)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success := dbConfig.UpdateRouteURL(a...)
	if success != true {
		fmt.Println("database update failed")
		t.Fail()
	}
}

func TestDbConfig_ActivateRouteURL(t *testing.T) {
	var a []interface{}
	a = append(a, routeURLID2, routeID, clientID)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success := dbConfig.ActivateRouteURL(a...)
	if success != true {
		fmt.Println("database update failed")
		t.Fail()
	}

	var a2 []interface{}
	a2 = append(a2, routeURLID2, routeID, clientID)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success2 := dbConfig.DeactivateOtherRouteURLs(a2...)
	if success2 != true {
		fmt.Println("database update failed")
		t.Fail()
	}
}

func TestDbConfig_GetRouteURL(t *testing.T) {
	a := []interface{}{routeURLID2, routeID, clientID}
	rowPtr := dbConfig.GetRouteURL(a...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		fmt.Print("Get ")
		fmt.Println(foundRow)
		//fmt.Println("Get results: --------------------------")
		int64Val, err2 := strconv.ParseInt(foundRow[0], 10, 0)
		if err2 != nil {
			fmt.Print(err2)
		}
		if routeURLID2 != int64Val {
			fmt.Print(routeURLID2)
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
	a := []interface{}{routeURLID, routeID, clientID}
	rowPtr := dbConfig.GetRouteURL(a...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		fmt.Print("Get ")
		fmt.Println(foundRow)
		//fmt.Println("Get results: --------------------------")
		int64Val, err2 := strconv.ParseInt(foundRow[0], 10, 0)
		if err2 != nil {
			fmt.Print(err2)
		}
		if routeURLID != int64Val {
			fmt.Print(routeURLID)
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
	a := []interface{}{routeID, clientID}
	rowsPtr := dbConfig.GetRouteURLList(a...)
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
	a := []interface{}{"mail2", clientID, "23ddeee"}
	rowsPtr := dbConfig.GetRouteNameURLList(a...)
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
	a := []interface{}{clientID, "mail2"}
	rowsPtr := dbConfig.GetRouteURLs(a...)
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
	a := []interface{}{100, 5000, time.Now().Add(time.Hour * -2400), routeURLID, routeID, clientID}
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success, insID := dbConfig.InsertRoutePerformance(a...)
	if success == true && insID != -1 {
		fmt.Print("new Id performance: ")
		fmt.Println(insID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestDbConfig_GetRoutePerformance(t *testing.T) {
	a := []interface{}{routeURLID, routeID, clientID}
	rowsPtr := dbConfig.GetRoutePerformance(a...)
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
	success := dbConfig.DeleteRoutePerformance(a...)
	if success == true {
		fmt.Print("Deleted route performance: ")
		fmt.Println(routeURLID)
	} else {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestDbConfig_InsertRouteError(t *testing.T) {
	a := []interface{}{404, "error call failed", time.Now().Add(time.Hour * -2400), routeURLID, routeID, clientID}
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success, insID := dbConfig.InsertRouteError(a...)
	if success == true && insID != -1 {
		fmt.Print("new Id route error id: ")
		fmt.Println(insID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestDbConfig_GetRouteError(t *testing.T) {
	a := []interface{}{routeURLID, routeID, clientID}
	rowsPtr := dbConfig.GetRouteError(a...)
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
	success := dbConfig.DeleteRouteError(a...)
	if success == true {
		fmt.Print("Deleted route error: ")
		fmt.Println(routeURLID)
	} else {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

var brkID int64

func TestDbConfig_InsertRouteBreaker(t *testing.T) {
	a := []interface{}{3, 500, "mail", 400, routeURLID, routeID, clientID}
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success, insID := dbConfig.InsertRouteBreaker(a...)
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
	a := []interface{}{5, 400, "mailblue", 401, brkID, routeURLID, routeID, clientID}
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success := dbConfig.UpdateRouteBreakerConfig(a...)
	if success != true {
		fmt.Println("database update failed")
		t.Fail()
	}
}

func TestDbConfig_UpdateRouteBreakerFail(t *testing.T) {
	a := []interface{}{1, time.Now(), brkID, routeURLID, routeID, clientID}
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success := dbConfig.UpdateRouteBreakerFail(a...)
	if success != true {
		fmt.Println("database update failed")
		t.Fail()
	}
}

func TestDbConfig_GetBreaker(t *testing.T) {
	a := []interface{}{routeURLID, routeID, clientID}
	rowPtr := dbConfig.GetBreaker(a...)
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
	a := []interface{}{routeURLID, routeID, clientID}
	success := dbConfig.DeleteBreaker(a...)
	if success == true {
		fmt.Print("Deleted route breaker: ")
		fmt.Println(routeURLID)
	} else {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestDbConfig_DeleteRouteURL(t *testing.T) {
	a := []interface{}{routeURLID, routeID, clientID}
	fmt.Println(a)
	success := dbConfig.DeleteRouteURL(a...)
	if success == true {
		fmt.Print("Deleted route url: ")
		fmt.Println(routeURLID)
	} else {
		fmt.Print("Deleted failed for route url: ")
		fmt.Println(routeURLID)
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestDbConfig_DeleteRestRoute(t *testing.T) {
	a := []interface{}{routeID, clientID}
	success := dbConfig.DeleteRestRoute(a...)
	if success == true {
		fmt.Print("Deleted route: ")
		fmt.Println(routeID)
	} else {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestDbConfig_DeleteClient(t *testing.T) {
	a := []interface{}{clientID}
	success := dbConfig.DeleteClient(a...)
	if success == true {
		fmt.Print("Deleted client ")
		fmt.Println(clientID)
	} else {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestDbConfig_CloseDb(t *testing.T) {
	if connected == true {
		rtn := dbConfig.CloseDb()
		if rtn != true {
			fmt.Println("database close failed")
			t.Fail()
		} else {
			fmt.Println("database close in database package")
		}
	}
}
