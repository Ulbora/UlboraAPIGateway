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
package monitor

import (
	db "UlboraApiGateway/database"
	"fmt"
	"strconv"
	"time"
)

//GatewayPerformanceMonitor error monitor
type GatewayPerformanceMonitor struct {
	DbConfig db.DbConfig
}

//GwPerformance GwPerformance
type GwPerformance struct {
	ID             int64
	Calls          int64
	LatencyMsTotal int64
	Entered        time.Time
	RouteURIID     int64
	RestRouteID    int64
	ClientID       int64
}

//ConnectDb to database
func (g *GatewayPerformanceMonitor) ConnectDb() bool {
	rtn := g.DbConfig.ConnectDb()
	if rtn == true {
		fmt.Println("db connect")
	}
	return rtn
}

//InsertRoutePerformance insert
func (g *GatewayPerformanceMonitor) InsertRoutePerformance(e *GwPerformance) (bool, error) {
	var success bool
	var err error
	dbConnected := g.DbConfig.ConnectionTest()
	if !dbConnected {
		fmt.Println("reconnection to closed database")
		g.DbConfig.ConnectDb()
	}
	//var a []interface{}
	a := []interface{}{e.Calls, e.LatencyMsTotal, e.Entered, e.RouteURIID, e.RestRouteID, e.ClientID}
	suc, insID := g.DbConfig.InsertRoutePerformance(a...)
	if suc == true && insID != -1 {
		success = suc
		//fmt.Print("new Id route error id: ")
		//fmt.Println(insID)
	} else {
		err = fmt.Errorf("Failed to insert route performance Record")
	}
	return success, err
}

//GetRoutePerformance from database
func (g *GatewayPerformanceMonitor) GetRoutePerformance(e *GwPerformance) *[]GwPerformance {
	a := []interface{}{e.RouteURIID, e.RestRouteID, e.ClientID}
	var rtn []GwPerformance
	rowsPtr := g.DbConfig.GetRoutePerformance(a...)
	if rowsPtr != nil {
		//print("content row: ")
		//println(rowPtr.Row)
		foundRows := rowsPtr.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := parseRoutePerformanceRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

//DeleteRoutePerformance from database
func (g *GatewayPerformanceMonitor) DeleteRoutePerformance() bool {
	a := []interface{}{} //{e.RouteURIID, e.RestRouteID, e.ClientID}
	var success bool
	suc := g.DbConfig.DeleteRoutePerformance(a...)
	if suc == true {
		success = suc
	} else {
		fmt.Println("Failed to delete performance Record")
	}
	return success
}

//CloseDb connection to database
func (g *GatewayPerformanceMonitor) CloseDb() bool {
	rtn := g.DbConfig.CloseDb()
	if rtn == true {
		fmt.Println("db connect closed")
	}
	return rtn
}

func parseRoutePerformanceRow(foundRow *[]string) *GwPerformance {
	var rtn GwPerformance
	if len(*foundRow) > 0 {
		rtn.ID, _ = strconv.ParseInt((*foundRow)[0], 10, 0)
		rtn.Calls, _ = strconv.ParseInt((*foundRow)[1], 10, 0)
		rtn.LatencyMsTotal, _ = strconv.ParseInt((*foundRow)[2], 10, 0)
		rtn.Entered, _ = time.Parse("2006-01-02 15:04:05", (*foundRow)[3])
		rtn.RouteURIID, _ = strconv.ParseInt((*foundRow)[4], 10, 0)
		rtn.RestRouteID, _ = strconv.ParseInt((*foundRow)[5], 10, 0)
		rtn.ClientID, _ = strconv.ParseInt((*foundRow)[6], 10, 0)
	}
	return &rtn
}
