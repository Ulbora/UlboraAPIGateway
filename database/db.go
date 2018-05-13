package database

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
	routeDb "UlboraApiGateway/database/mysqldb"
	"fmt"
	"strconv"
)

//DbConfig db config
type DbConfig struct {
	Host         string
	DbUser       string
	DbPw         string
	DatabaseName string
}

//Row database row
type Row struct {
	Columns []string
	Row     []string
}

//Rows array of database rows
type Rows struct {
	Columns []string
	Rows    [][]string
}

//ConnectDb to database
func (db *DbConfig) ConnectDb() bool {
	rtn := routeDb.ConnectDb(db.Host, db.DbUser, db.DbPw, db.DatabaseName)
	if rtn {
		fmt.Println("db connect in db")
	}
	return rtn
}

//ConnectionTest of database
func (db *DbConfig) ConnectionTest() bool {
	var rtn = false
	//fmt.Print("db in db: ")
	rowPtr := routeDb.ConnectionTest()
	//fmt.Println(rowPtr)
	if rowPtr != nil && len(rowPtr.Row) > 0 {
		foundRow := rowPtr.Row
		int64Val, err2 := strconv.ParseInt(foundRow[0], 10, 0)
		//fmt.Print("Records found during test ")
		//fmt.Println(int64Val)
		if err2 != nil {
			fmt.Print(err2)
		}
		if int64Val >= 0 {
			rtn = true
		}
	}
	return rtn
}

//InsertClient in database
func (db *DbConfig) InsertClient(args ...interface{}) (bool, int64) {
	success, insID := routeDb.InsertClient(args...)
	if success {
		fmt.Println("inserted record")
	}
	return success, insID
}

//UpdateClient in database
func (db *DbConfig) UpdateClient(args ...interface{}) bool {
	success := routeDb.UpdateClient(args...)
	if success {
		fmt.Println("updated record")
	}
	return success
}

//GetClient get a row. Passing in tx allows for transactions
func (db *DbConfig) GetClient(args ...interface{}) *Row {
	var clientRow Row
	rowPtr := routeDb.GetClient(args...)
	if rowPtr != nil {
		clientRow.Columns = rowPtr.Columns
		clientRow.Row = rowPtr.Row
	}
	return &clientRow
}

//GetClientList get a row. Passing in tx allows for transactions
func (db *DbConfig) GetClientList(args ...interface{}) *Rows {
	var clientRows Rows
	rowsPtr := routeDb.GetClientList(args...)
	if rowsPtr != nil {
		clientRows.Columns = rowsPtr.Columns
		clientRows.Rows = rowsPtr.Rows
	}
	return &clientRows
}

//DeleteClient delete
func (db *DbConfig) DeleteClient(args ...interface{}) bool {
	success := routeDb.DeleteClient(args...)
	return success
}

//InsertRoutePerformance in database
func (db *DbConfig) InsertRoutePerformance(args ...interface{}) (bool, int64) {
	success, insID := routeDb.InsertRoutePerformance(args...)
	if success {
		fmt.Println("inserted record")
	}
	return success, insID
}

//GetRoutePerformance get a row. Passing in tx allows for transactions
func (db *DbConfig) GetRoutePerformance(args ...interface{}) *Rows {
	var clientRows Rows
	rowsPtr := routeDb.GetRoutePerformance(args...)
	if rowsPtr != nil {
		clientRows.Columns = rowsPtr.Columns
		clientRows.Rows = rowsPtr.Rows
	}
	return &clientRows
}

//DeleteRoutePerformance delete
func (db *DbConfig) DeleteRoutePerformance(args ...interface{}) bool {
	success := routeDb.DeleteRoutePerformance(args...)
	return success
}

//InsertRouteError in database
func (db *DbConfig) InsertRouteError(args ...interface{}) (bool, int64) {
	success, insID := routeDb.InsertRouteError(args...)
	if success {
		fmt.Println("inserted record")
	}
	return success, insID
}

//GetRouteError get a row. Passing in tx allows for transactions
func (db *DbConfig) GetRouteError(args ...interface{}) *Rows {
	var clientRows Rows
	rowsPtr := routeDb.GetRouteError(args...)
	if rowsPtr != nil {
		clientRows.Columns = rowsPtr.Columns
		clientRows.Rows = rowsPtr.Rows
	}
	return &clientRows
}

//DeleteRouteError delete
func (db *DbConfig) DeleteRouteError(args ...interface{}) bool {
	success := routeDb.DeleteRouteError(args...)
	return success
}

//InsertRouteBreaker in database
func (db *DbConfig) InsertRouteBreaker(args ...interface{}) (bool, int64) {
	success, insID := routeDb.InsertRouteBreaker(args...)
	if success {
		fmt.Println("inserted record")
	}
	return success, insID
}

//UpdateRouteBreakerConfig in database
func (db *DbConfig) UpdateRouteBreakerConfig(args ...interface{}) bool {
	success := routeDb.UpdateRouteBreakerConfig(args...)
	if success {
		fmt.Println("updated record")
	}
	return success
}

//UpdateRouteBreakerFail in database
func (db *DbConfig) UpdateRouteBreakerFail(args ...interface{}) bool {
	success := routeDb.UpdateRouteBreakerFail(args...)
	if success {
		fmt.Println("updated record")
	}
	return success
}

//GetBreaker get a row. Passing in tx allows for transactions
func (db *DbConfig) GetBreaker(args ...interface{}) *Row {
	var clientRow Row
	rowPtr := routeDb.GetBreaker(args...)
	if rowPtr != nil {
		clientRow.Columns = rowPtr.Columns
		clientRow.Row = rowPtr.Row
	}
	return &clientRow
}

//DeleteBreaker delete
func (db *DbConfig) DeleteBreaker(args ...interface{}) bool {
	success := routeDb.DeleteBreaker(args...)
	return success
}

//InsertRestRoute in database
func (db *DbConfig) InsertRestRoute(args ...interface{}) (bool, int64) {
	success, insID := routeDb.InsertRestRoute(args...)
	if success {
		fmt.Println("inserted record")
	}
	return success, insID
}

//UpdateRestRoute in database
func (db *DbConfig) UpdateRestRoute(args ...interface{}) bool {
	success := routeDb.UpdateRestRoute(args...)
	if success {
		fmt.Println("updated record")
	}
	return success
}

//GetRestRoute get a row. Passing in tx allows for transactions
func (db *DbConfig) GetRestRoute(args ...interface{}) *Row {
	var routeRow Row
	rowPtr := routeDb.GetRestRoute(args...)
	if rowPtr != nil {
		routeRow.Columns = rowPtr.Columns
		routeRow.Row = rowPtr.Row
	}
	return &routeRow
}

//GetRestRouteList get a row. Passing in tx allows for transactions
func (db *DbConfig) GetRestRouteList(args ...interface{}) *Rows {
	var routeRows Rows
	rowsPtr := routeDb.GetRestRouteList(args...)
	if rowsPtr != nil {
		routeRows.Columns = rowsPtr.Columns
		routeRows.Rows = rowsPtr.Rows
	}
	return &routeRows
}

//DeleteRestRoute delete
func (db *DbConfig) DeleteRestRoute(args ...interface{}) bool {
	success := routeDb.DeleteRestRoute(args...)
	return success
}

//InsertRouteURL in database
func (db *DbConfig) InsertRouteURL(args ...interface{}) (bool, int64) {
	success, insID := routeDb.InsertRouteURL(args...)
	if success {
		fmt.Println("inserted record")
	}
	return success, insID
}

//UpdateRouteURL in database
func (db *DbConfig) UpdateRouteURL(args ...interface{}) bool {
	success := routeDb.UpdateRouteURL(args...)
	if success {
		fmt.Println("updated record")
	}
	return success
}

//ActivateRouteURL in database
func (db *DbConfig) ActivateRouteURL(args ...interface{}) bool {
	success := routeDb.ActivateRouteURL(args...)
	if success {
		fmt.Println("updated record")
	}
	return success
}

//DeactivateOtherRouteURLs in database
func (db *DbConfig) DeactivateOtherRouteURLs(args ...interface{}) bool {
	success := routeDb.DeactivateOtherRouteURLs(args...)
	if success {
		fmt.Println("updated record")
	}
	return success
}

//GetRouteURL get a row. Passing in tx allows for transactions
func (db *DbConfig) GetRouteURL(args ...interface{}) *Row {
	var routeRow Row
	rowPtr := routeDb.GetRouteURL(args...)
	if rowPtr != nil {
		routeRow.Columns = rowPtr.Columns
		routeRow.Row = rowPtr.Row
	}
	return &routeRow
}

//GetRouteURLList get a row. Passing in tx allows for transactions
func (db *DbConfig) GetRouteURLList(args ...interface{}) *Rows {
	var routeRows Rows
	rowsPtr := routeDb.GetRouteURLList(args...)
	if rowsPtr != nil {
		routeRows.Columns = rowsPtr.Columns
		routeRows.Rows = rowsPtr.Rows
	}
	return &routeRows
}

//GetRouteNameURLList get a row. Passing in tx allows for transactions
func (db *DbConfig) GetRouteNameURLList(args ...interface{}) *Rows {
	var routeRows Rows
	rowsPtr := routeDb.GetRouteNameURLList(args...)
	//fmt.Print("database row: ")
	//fmt.Println(rowsPtr)
	if rowsPtr != nil {
		routeRows.Columns = rowsPtr.Columns
		routeRows.Rows = rowsPtr.Rows
	}
	return &routeRows
}

//DeleteRouteURL delete
func (db *DbConfig) DeleteRouteURL(args ...interface{}) bool {
	success := routeDb.DeleteRouteURL(args...)
	return success
}

//GetRouteURLs get a row. Passing in tx allows for transactions
func (db *DbConfig) GetRouteURLs(args ...interface{}) *Rows {
	var routeRows Rows
	rowsPtr := routeDb.GetRouteURLs(args...)
	if rowsPtr != nil {
		routeRows.Columns = rowsPtr.Columns
		routeRows.Rows = rowsPtr.Rows
	}
	return &routeRows
}

//CloseDb database connection
func (db *DbConfig) CloseDb() bool {
	rtn := routeDb.CloseDb()
	if rtn {
		fmt.Println("db connection closed in db")
	}
	return rtn
}
