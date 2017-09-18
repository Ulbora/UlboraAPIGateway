/*
 Copyright (C) 2017 Ulbora Labs Inc. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2017 Ken Williamson
 All rights reserved.

 Certain inventions and disclosures in this file may be claimed within
 patents owned or patent applications filed by Ulbora Labs Inc., or third
 parties.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Affero General Public License as published
 by the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Affero General Public License for more details.

 You should have received a copy of the GNU Affero General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package managers

import (
	"fmt"
	"strconv"
)

//InsertClient in database
func (db *GatewayDB) InsertClient(client *Client) *GatewayResponse {
	var rtn GatewayResponse
	dbConnected := db.DbConfig.ConnectionTest()
	if !dbConnected {
		fmt.Println("reconnection to closed database")
		db.DbConfig.ConnectDb()
	}
	var a []interface{}
	a = append(a, client.ClientID, client.APIKey, client.Enabled, client.Level)
	success, insID := db.DbConfig.InsertClient(a...)
	if success == true {
		fmt.Println("inserted record")
	}
	rtn.ID = insID
	rtn.Success = success
	return &rtn
}

//UpdateClient in database
func (db *GatewayDB) UpdateClient(client *Client) *GatewayResponse {
	var rtn GatewayResponse
	dbConnected := db.DbConfig.ConnectionTest()
	if !dbConnected {
		fmt.Println("reconnection to closed database")
		db.DbConfig.ConnectDb()
	}
	var a []interface{}
	a = append(a, client.APIKey, client.Enabled, client.Level, client.ClientID)
	success := db.DbConfig.UpdateClient(a...)
	if success == true {
		fmt.Println("update record")
	}
	rtn.ID = client.ClientID
	rtn.Success = success
	return &rtn
}

//GetClient client from database
func (db *GatewayDB) GetClient(client *Client) *Client {
	var a []interface{}
	a = append(a, client.ClientID)
	var rtn *Client
	rowPtr := db.DbConfig.GetClient(a...)
	if rowPtr != nil {
		//print("content row: ")
		//println(rowPtr.Row)
		foundRow := rowPtr.Row
		rtn = parseClientRow(&foundRow)
	}
	return rtn
}

//GetClientList client
func (db *GatewayDB) GetClientList(client *Client) *[]Client {
	var rtn []Client
	var a []interface{}
	//a = append(a, content.ClientID)
	rowsPtr := db.DbConfig.GetClientList(a...)
	if rowsPtr != nil {
		foundRows := rowsPtr.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := parseClientRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteClient in database
func (db *GatewayDB) DeleteClient(client *Client) *GatewayResponse {
	var rtn GatewayResponse
	dbConnected := db.DbConfig.ConnectionTest()
	if !dbConnected {
		fmt.Println("reconnection to closed database")
		db.DbConfig.ConnectDb()
	}
	var a []interface{}
	a = append(a, client.ClientID)
	success := db.DbConfig.DeleteClient(a...)
	if success == true {
		fmt.Println("deleted record")
	}
	rtn.ID = client.ClientID
	rtn.Success = success
	return &rtn
}

func parseClientRow(foundRow *[]string) *Client {
	var rtn Client
	if len(*foundRow) > 0 {
		ID, errID := strconv.ParseInt((*foundRow)[0], 10, 0)
		if errID != nil {
			fmt.Print(errID)
		} else {
			rtn.ClientID = ID
		}
		rtn.APIKey = (*foundRow)[1]
		enabled, err3 := strconv.ParseBool((*foundRow)[2])
		if err3 != nil {
			fmt.Print(err3)
			rtn.Enabled = false
		} else {
			rtn.Enabled = enabled
		}
		rtn.Level = (*foundRow)[3]
	}
	return &rtn
}
