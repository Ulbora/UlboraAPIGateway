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

// //GetActiveRouteURL route from database
// func (db *GatewayDB) GetActiveRouteURL(ru *RouteURL) *RouteURL {
// 	var a []interface{}
// 	a = append(a, ru.ID, ru.RouteID, ru.ClientID)
// 	var rtn *RouteURL
// 	rowPtr := db.DbConfig.GetRouteURL(a...)
// 	if rowPtr != nil {
// 		//print("content row: ")
// 		//println(rowPtr.Row)
// 		foundRow := rowPtr.Row
// 		rtn = parseRouteURLRow(&foundRow)
// 	}
// 	return rtn
// }

//GetGatewayRoutes route
func (gw *GatewayRoutes) GetGatewayRoutes() *[]GatewayRouteURL {
	var rtn []GatewayRouteURL

	// check cache for saved value---------

	// //--------------

	// var a []interface{}
	// a = append(a, ru.RouteID, ru.ClientID)
	// rowsPtr := gw.GwDB.DbConfig.GetRouteNameURLList(a...)
	// if rowsPtr != nil {
	// 	foundRows := rowsPtr.Rows
	// 	for r := range foundRows {
	// 		foundRow := foundRows[r]
	// 		rowContent := parseRouteURLsRow(&foundRow)
	// 		rtn = append(rtn, *rowContent)
	// 	}
	// }

	return &rtn
}

func parseGatewayRoutesRow(foundRow *[]string) *RouteURL {
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
