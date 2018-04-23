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
	ch "UlboraApiGateway/cache"
	cb "UlboraApiGateway/circuitbreaker"
	db "UlboraApiGateway/database"
	"fmt"
	"strconv"
)

//GatewayResponse res
type GatewayResponse struct {
	Success bool  `json:"success"`
	ID      int64 `json:"id"`
}

//Client client
type Client struct {
	ClientID int64  `json:"clientId"`
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
	RouteID      int64  `json:"routeId"`
	Route        string `json:"route"`
	URLID        int64  `json:"urlId"`
	Name         string `json:"name"`
	URL          string `json:"url"`
	Active       bool   `json:"active"`
	CircuitOpen  bool   `json:"circuitOpen"`
	OpenFailCode int    `json:"openFailCode"`
}

//GatewayClusterRouteURL url
type GatewayClusterRouteURL struct {
	RouteID                int64  `json:"routeId"`
	Route                  string `json:"route"`
	URLID                  int64  `json:"urlId"`
	Name                   string `json:"name"`
	URL                    string `json:"url"`
	Active                 bool   `json:"active"`
	CircuitOpen            bool   `json:"circuitOpen"`
	OpenFailCode           int    `json:"openFailCode"`
	FailoverRouteName      string `json:"failoverRouteName"`
	FailureThreshold       int    `json:"failureThreshold"`
	HealthCheckTimeSeconds int    `json:"healthCheckTimeSeconds"`
}

//GatewayDB db config
type GatewayDB struct {
	DbConfig    db.DbConfig
	GwCacheHost string
	Cb          cb.CircuitBreaker
}

//GatewayRoutes gateway routes
type GatewayRoutes struct {
	Route       string
	APIKey      string
	ClientID    int64
	GwDB        GatewayDB
	GwCacheHost string
}

//GateStatusResponse GateStatusResponse
type GateStatusResponse struct {
	Success       bool `json:"success"`
	RouteModified bool `json:"routeModified"`
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

func (db *GatewayDB) clearCache(clientID int64, route string) {
	var cp ch.CProxy
	cp.Host = db.GwCacheHost
	var cid = strconv.FormatInt(clientID, 10)
	var key = cid + ":" + route
	cp.Delete(key)
}
