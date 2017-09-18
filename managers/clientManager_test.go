package managers

import (
	"fmt"
	"testing"
)

var gatewayDB GatewayDB
var connected bool
var clientID int64
var insertID int64

func TestClient_ConnectDb(t *testing.T) {
	clientID = 4
	gatewayDB.DbConfig.Host = "localhost:3306"
	gatewayDB.DbConfig.DbUser = "admin"
	gatewayDB.DbConfig.DbPw = "admin"
	gatewayDB.DbConfig.DatabaseName = "ulbora_api_gateway"
	connected = gatewayDB.ConnectDb()
	if connected != true {
		t.Fail()
	}
}

func TestClient_InsertClient(t *testing.T) {
	var c Client
	c.APIKey = "12233hgdd"
	c.ClientID = clientID
	c.Enabled = true
	c.Level = "small"

	res := gatewayDB.InsertClient(&c)
	if res.Success == true && res.ID != -1 {
		insertID = clientID
		fmt.Print("new Id: ")
		fmt.Println(insertID)
	} else {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestClient_UpdateClient(t *testing.T) {
	var c Client
	c.APIKey = "12555233hgdd"
	c.ClientID = clientID
	c.Enabled = true
	c.Level = "small"

	res := gatewayDB.UpdateClient(&c)
	if res.Success != true {
		fmt.Println("database update failed")
		t.Fail()
	}
}

func TestClient_GetClient(t *testing.T) {
	var c Client
	c.ClientID = clientID
	res := gatewayDB.GetClient(&c)
	fmt.Println("")
	fmt.Print("found client: ")
	fmt.Println(res)
	if res.ClientID != clientID {
		fmt.Println("database insert failed")
		t.Fail()
	}
}

func TestClient_GetClientList(t *testing.T) {
	var c Client
	res := gatewayDB.GetClientList(&c)
	fmt.Println("")
	fmt.Print("found client list: ")
	fmt.Println(res)
	if len(*res) == 0 {
		fmt.Println("database read failed")
		t.Fail()
	}
}

func TestClient_DeleteClient(t *testing.T) {
	var c Client
	c.ClientID = clientID
	res := gatewayDB.DeleteClient(&c)
	if res.Success != true {
		fmt.Println("database delete failed")
		t.Fail()
	}
}

func TestClient_TestCloseDb(t *testing.T) {
	success := gatewayDB.CloseDb()
	if success != true {
		t.Fail()
	}
}
