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
	db "UlboraApiGateway/database"
	"fmt"
)

//GatewayResponse res
type GatewayResponse struct {
	Success bool  `json:"success"`
	ID      int64 `json:"id"`
}

//Client client
type Client struct {
	ClientID int64  `json:"id"`
	APIKey   string `json:"apiKey"`
	Enabled  bool   `json:"enabled"`
	Level    string `json:"level"`
}

//RestRoute rest route
type RestRoute struct {
	ID       int64  `json:"id"`
	Route    string `json:"route"`
	ClientID int64  `json:"clientId"`
}

//RouteURL url
type RouteURL struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	Active   bool   `json:"active"`
	RouteID  int64  `json:"routeId"`
	ClientID int64  `json:"clientId"`
}

//GatewayRouteURL url
type GatewayRouteURL struct {
	Route  string `json:"route"`
	Name   string `json:"name"`
	URL    string `json:"url"`
	Active bool   `json:"active"`
}

//GatewayDB db config
type GatewayDB struct {
	DbConfig db.DbConfig
}

//GatewayRoutes gateway routes
type GatewayRoutes struct {
	Route       string
	ClientID    int64
	GwDB        GatewayDB
	GwCacheHost string
}

//ConnectDb to database
func (db *GatewayDB) ConnectDb() bool {
	rtn := db.DbConfig.ConnectDb()
	if rtn == true {
		fmt.Println("db connect")
	}
	return rtn
}

//CloseDb connection to database
func (db *GatewayDB) CloseDb() bool {
	rtn := db.DbConfig.CloseDb()
	if rtn == true {
		fmt.Println("db connect closed")
	}
	return rtn
}
