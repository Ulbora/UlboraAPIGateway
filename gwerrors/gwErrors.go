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
package gwerrors

import (
	db "UlboraApiGateway/database"
	"fmt"
	"time"
)

//GatewayErrorMonitor error monitor
type GatewayErrorMonitor struct {
	DbConfig db.DbConfig
}

//GwError GwError
type GwError struct {
	ID          int64
	Code        int
	Message     string
	Entered     time.Time
	RouteURIID  int64
	RestRouteID int64
	ClientID    int64
}

//ConnectDb to database
func (g *GatewayErrorMonitor) ConnectDb() bool {
	rtn := g.DbConfig.ConnectDb()
	if rtn == true {
		fmt.Println("db connect")
	}
	return rtn
}

//InsertRouteError insert
func (g *GatewayErrorMonitor) InsertRouteError(e *GwError) (bool, error) {
	var success bool
	var err error
	dbConnected := g.DbConfig.ConnectionTest()
	if !dbConnected {
		fmt.Println("reconnection to closed database")
		g.DbConfig.ConnectDb()
	}
	//var a []interface{}
	a := []interface{}{e.Code, e.Message, time.Now().Add(time.Hour * -2400), e.RouteURIID, e.RestRouteID, e.ClientID}
	suc, insID := g.DbConfig.InsertRouteError(a...)
	if suc == true && insID != -1 {
		success = suc
		//fmt.Print("new Id route error id: ")
		//fmt.Println(insID)
	} else {
		err = fmt.Errorf("Failed to insert Error Record")
	}
	return success, err
}

//CloseDb connection to database
func (g *GatewayErrorMonitor) CloseDb() bool {
	rtn := g.DbConfig.CloseDb()
	if rtn == true {
		fmt.Println("db connect closed")
	}
	return rtn
}
