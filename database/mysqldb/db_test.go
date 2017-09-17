package mysqldb

import (
	"fmt"
	"strconv"
	"testing"
)

var connected bool
var clientID int64
var insertID int64
var routeID int64
var routeURLID int64
var routeURLID2 int64

func TestConnectDb(t *testing.T) {
	connected = ConnectDb("localhost:3306", "admin", "admin", "ulbora_api_gateway")
	if connected != true {
		fmt.Println("database init failed")
		t.Fail()
	}
}

func TestConnectionTest(t *testing.T) {
	clientID = 2
	var a []interface{}
	rowPtr := ConnectionTest(a...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		//fmt.Print("Records found during test ")
		//fmt.Println(foundRow)
		//fmt.Println("Get results: --------------------------")
		int64Val, err2 := strconv.ParseInt(foundRow[0], 10, 0)
		fmt.Print("Records found during test ")
		fmt.Println(int64Val)
		if err2 != nil {
			fmt.Print(err2)
		}
		if int64Val >= 0 {
			fmt.Println("database connection ok")
		} else {
			fmt.Println("database connection failed")
			t.Fail()
		}
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
}
func TestInsertClient(t *testing.T) {
	var a []interface{}
	a = append(a, clientID, "23ddeee", true, "small")
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success, insID := InsertClient(a...)
	if success == true && insID != -1 {
		insertID = clientID
		fmt.Print("new Id: ")
		fmt.Println(insID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestUpdateClient(t *testing.T) {
	var a []interface{}
	a = append(a, "23ddeee", true, "small", clientID)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success := UpdateClient(a...)
	if success != true {
		fmt.Println("database update failed")
		t.Fail()
	}
}

func TestGetClient(t *testing.T) {
	a := []interface{}{clientID}
	rowPtr := GetClient(a...)
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

func TestGetClientList(t *testing.T) {
	a := []interface{}{}
	rowsPtr := GetClientList(a...)
	if rowsPtr != nil {
		foundRows := rowsPtr.Rows
		fmt.Print("Get clients ")
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

func TestInsertRestRoute(t *testing.T) {
	var a []interface{}
	a = append(a, "mail", clientID)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success, insID := InsertRestRoute(a...)
	if success == true && insID != -1 {
		routeID = insID
		fmt.Print("new Id: ")
		fmt.Println(insID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestUpdateRestRoute(t *testing.T) {
	var a []interface{}
	a = append(a, "mail2", routeID, clientID)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success := UpdateRestRoute(a...)
	if success != true {
		fmt.Println("database update failed")
		t.Fail()
	}
}

func TestGetRestRoute(t *testing.T) {
	a := []interface{}{routeID, clientID}
	rowPtr := GetRestRoute(a...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		fmt.Print("Get ")
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
	} else {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestGetRestRouteList(t *testing.T) {
	a := []interface{}{clientID}
	rowsPtr := GetRestRouteList(a...)
	if rowsPtr != nil {
		foundRows := rowsPtr.Rows
		fmt.Print("Get clients ")
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

func TestInsertRouteURL(t *testing.T) {
	var a []interface{}
	a = append(a, "blue", "http://www.apigateway.com/blue/", true, routeID, clientID)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success, insID := InsertRouteURL(a...)
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
	success2, insID2 := InsertRouteURL(a2...)
	if success2 == true && insID2 != -1 {
		routeURLID2 = insID2
		fmt.Print("new URL Id: ")
		fmt.Println(insID2)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestUpdateRouteURL(t *testing.T) {
	var a []interface{}
	a = append(a, "green", "http://www.apigateway.com/green/", false, routeURLID2, routeID, clientID)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success := UpdateRouteURL(a...)
	if success != true {
		fmt.Println("database update failed")
		t.Fail()
	}
}

func TestActivateRouteURL(t *testing.T) {
	var a []interface{}
	a = append(a, routeURLID2, routeID, clientID)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success := ActivateRouteURL(a...)
	if success != true {
		fmt.Println("database update failed")
		t.Fail()
	}

	var a2 []interface{}
	a2 = append(a2, routeURLID2, routeID, clientID)
	//can also be: a := []interface{}{"test insert", time.Now(), "some content text", 125}
	success2 := DeactivateOtherRouteURLs(a2...)
	if success2 != true {
		fmt.Println("database update failed")
		t.Fail()
	}
}

func TestGetRouteURL(t *testing.T) {
	a := []interface{}{routeURLID2, routeID, clientID}
	rowPtr := GetRouteURL(a...)
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

func TestGetRouteURL2(t *testing.T) {
	a := []interface{}{routeURLID, routeID, clientID}
	rowPtr := GetRouteURL(a...)
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

func TestGetRouteURLList(t *testing.T) {
	a := []interface{}{routeID, clientID}
	rowsPtr := GetRouteURLList(a...)
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

func TestGetRouteURLs(t *testing.T) {
	a := []interface{}{clientID, "mail2"}
	rowsPtr := GetRouteURLs(a...)
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

func TestDeleteClient(t *testing.T) {
	a := []interface{}{clientID}
	success := DeleteClient(a...)
	if success == true {
		fmt.Print("Deleted client ")
		fmt.Println(clientID)
	} else {
		fmt.Println("database delete failed")
		t.Fail()
	}

	// a2 := []interface{}{insertID2, 125}
	// success2 := DeleteContent(a2...)
	// if success2 == true {
	// 	fmt.Print("Deleted ")
	// 	fmt.Println(insertID2)
	// } else {
	// 	fmt.Println("database delete failed")
	// 	t.Fail()
	// }
}

// func TestDeleteRestRoute(t *testing.T) {
// 	a := []interface{}{routeID, clientID}
// 	success := DeleteRestRoute(a...)
// 	if success == true {
// 		fmt.Print("Deleted ")
// 		fmt.Println(routeID)
// 	} else {
// 		fmt.Println("database delete failed")
// 		t.Fail()
// 	}

// 	// a2 := []interface{}{insertID2, 125}
// 	// success2 := DeleteContent(a2...)
// 	// if success2 == true {
// 	// 	fmt.Print("Deleted ")
// 	// 	fmt.Println(insertID2)
// 	// } else {
// 	// 	fmt.Println("database delete failed")
// 	// 	t.Fail()
// 	// }
// }

// func TestDeleteRouteURL(t *testing.T) {
// 	a := []interface{}{routeURLID, routeID, clientID}
// 	success := DeleteRouteURL(a...)
// 	if success == true {
// 		fmt.Print("Deleted route url: ")
// 		fmt.Println(routeURLID)
// 	} else {
// 		fmt.Println("database delete failed")
// 		t.Fail()
// 	}
// }

func TestCloseDb(t *testing.T) {
	if connected == true {
		rtn := CloseDb()
		if rtn != true {
			fmt.Println("database close failed")
			t.Fail()
		}
	}
}
