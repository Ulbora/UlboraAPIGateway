package managers

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

import (
	"fmt"
	"strconv"
)

//InsertRouteURL in database
func (db *GatewayDB) InsertRouteURL(ru *RouteURL) *GatewayResponse {
	var rtn GatewayResponse
	dbConnected := db.DbConfig.ConnectionTest()
	if !dbConnected {
		fmt.Println("reconnection to closed database")
		db.DbConfig.ConnectDb()
	}
	var a []interface{}
	a = append(a, ru.Name, ru.URL, ru.Active, ru.RouteID, ru.ClientID)
	success, insID := db.DbConfig.InsertRouteURL(a...)
	if success == true {
		fmt.Println("inserted record")
	}
	rtn.ID = insID
	rtn.Success = success
	return &rtn
}

//UpdateRouteURL in database
func (db *GatewayDB) UpdateRouteURL(ru *RouteURL) *GatewayResponse {
	var rtn GatewayResponse
	dbConnected := db.DbConfig.ConnectionTest()
	if !dbConnected {
		fmt.Println("reconnection to closed database")
		db.DbConfig.ConnectDb()
	}
	var a []interface{}
	a = append(a, ru.Name, ru.URL, ru.ID, ru.RouteID, ru.ClientID)
	success := db.DbConfig.UpdateRouteURL(a...)
	if success == true {
		fmt.Println("update record")
		var rr RestRoute
		rr.ID = ru.RouteID
		rr.ClientID = ru.ClientID
		db.Cb.Reset(ru.ClientID, ru.ID)
		route := db.GetRestRoute(&rr)
		if route != nil {
			db.clearCache(ru.ClientID, route.Route)
		}
	}
	rtn.ID = ru.ID
	rtn.Success = success
	return &rtn
}

//ActivateRouteURL in database
func (db *GatewayDB) ActivateRouteURL(ru *RouteURL) *GatewayResponse {
	var rtn GatewayResponse
	dbConnected := db.DbConfig.ConnectionTest()
	if !dbConnected {
		fmt.Println("reconnection to closed database")
		db.DbConfig.ConnectDb()
	}
	var a []interface{}
	a = append(a, ru.ID, ru.RouteID, ru.ClientID)
	success := db.DbConfig.ActivateRouteURL(a...)
	if success == true {
		fmt.Println("activated urls")
		successd := db.DbConfig.DeactivateOtherRouteURLs(a...)
		if successd == true {
			fmt.Println("deactivated other urls")
			var rr RestRoute
			rr.ID = ru.RouteID
			rr.ClientID = ru.ClientID
			db.Cb.Reset(ru.ClientID, ru.ID)
			route := db.GetRestRoute(&rr)
			if route != nil {
				db.clearCache(ru.ClientID, route.Route)
			}
		}
	}
	rtn.ID = ru.ID
	rtn.Success = success
	return &rtn
}

//GetRouteURL route from database
func (db *GatewayDB) GetRouteURL(ru *RouteURL) *RouteURL {
	var a []interface{}
	a = append(a, ru.ID, ru.RouteID, ru.ClientID)
	var rtn *RouteURL
	rowPtr := db.DbConfig.GetRouteURL(a...)
	if rowPtr != nil {
		//print("content row: ")
		//println(rowPtr.Row)
		foundRow := rowPtr.Row
		rtn = parseRouteURLRow(&foundRow)
	}
	return rtn
}

//GetRouteURLList route
func (db *GatewayDB) GetRouteURLList(ru *RouteURL) *[]RouteURL {
	var rtn []RouteURL
	var a []interface{}
	a = append(a, ru.RouteID, ru.ClientID)
	rowsPtr := db.DbConfig.GetRouteURLList(a...)
	if rowsPtr != nil {
		foundRows := rowsPtr.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := parseRouteURLRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteRouteURL in database
func (db *GatewayDB) DeleteRouteURL(ru *RouteURL) *GatewayResponse {
	var rtn GatewayResponse
	dbConnected := db.DbConfig.ConnectionTest()
	if !dbConnected {
		fmt.Println("reconnection to closed database")
		db.DbConfig.ConnectDb()
	}
	var a []interface{}
	a = append(a, ru.ID, ru.RouteID, ru.ClientID)
	success := db.DbConfig.DeleteRouteURL(a...)
	if success == true {
		fmt.Println("deleted record")
		var rr RestRoute
		rr.ID = ru.RouteID
		rr.ClientID = ru.ClientID
		db.Cb.Reset(ru.ClientID, ru.ID)
		route := db.GetRestRoute(&rr)
		if route != nil {
			db.clearCache(ru.ClientID, route.Route)
		}
	}
	rtn.ID = ru.ID
	rtn.Success = success
	return &rtn
}

func parseRouteURLRow(foundRow *[]string) *RouteURL {
	var rtn RouteURL
	if len(*foundRow) > 0 {
		ID, errID := strconv.ParseInt((*foundRow)[0], 10, 0)
		if errID != nil {
			fmt.Print(errID)
		} else {
			rtn.ID = ID
		}
		rtn.Name = (*foundRow)[1]
		rtn.URL = (*foundRow)[2]
		active, err := strconv.ParseBool((*foundRow)[3])
		if err != nil {
			fmt.Print(err)
			rtn.Active = false
		} else {
			rtn.Active = active
		}
		RID, errRID := strconv.ParseInt((*foundRow)[4], 10, 0)
		if errRID != nil {
			fmt.Print(errRID)
		} else {
			rtn.RouteID = RID
		}
		CID, errID2 := strconv.ParseInt((*foundRow)[5], 10, 0)
		if errID2 != nil {
			fmt.Print(errID2)
		} else {
			rtn.ClientID = CID
		}
	}
	return &rtn
}
