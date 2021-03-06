package managers

/*
 Copyright (C) 2017 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2017 Ken Williamson
 All rights reserved.

 Certain inventions and disclosures in this file may be claimed within
 patents owned or patent applications filed by Ulbora Labs LLC., or third
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
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	//"strconv"
	cb "UlboraApiGateway/circuitbreaker"
)

//ClusterResponse ClusterResponse
type ClusterResponse struct {
	Success bool `json:"success"`
}

//GetClusterGwRoutes GetClusterGwRoutes
func (gw *GatewayRoutes) GetClusterGwRoutes() *[]GatewayClusterRouteURL {
	//var rtnVal GatewayRouteURL
	var rtn = make([]GatewayClusterRouteURL, 0)
	var cbDB cb.CircuitBreaker
	cbDB.CacheHost = gw.GwCacheHost
	var cp ch.CProxy
	cp.Host = gw.GwCacheHost
	var cid = strconv.FormatInt(gw.ClientID, 10)
	var key = cid + ":cluster:" + gw.Route
	res := cp.Get(key)
	if res.Success {
		rJSON, err := b64.StdEncoding.DecodeString(res.Value)
		if err != nil {
			fmt.Println(err)
		} else {
			err := json.Unmarshal([]byte(rJSON), &rtn)
			if err != nil {
				fmt.Println(err)
			} else {
				//cbs := cbDB.GetStatus(gw.ClientID, )
				for i := range rtn {
					cbs := cbDB.GetStatus(gw.ClientID, rtn[i].URLID)
					rtn[i].CircuitOpen = cbs.Open
					rtn[i].PartialOpen = cbs.PartialOpen
				}

			}
			//fmt.Print("rtn in cluster routes")
			//fmt.Println(rtn)
		}
		//fmt.Println("Found Gateway route in cache for key: " + key)
	} else {
		fmt.Println("Routes not found in cache for key " + key + ", reading db.")

		dbConnected := gw.GwDB.DbConfig.ConnectionTest()
		if !dbConnected {
			fmt.Println("reconnection to closed database")
			gw.GwDB.DbConfig.ConnectDb()
		}
		var a []interface{}
		a = append(a, gw.Route, gw.ClientID, gw.APIKey)
		rowsPtr := gw.GwDB.DbConfig.GetRouteNameURLList(a...)
		//fmt.Print("rows")
		//fmt.Println(rowsPtr)
		if rowsPtr != nil {
			foundRows := rowsPtr.Rows
			for r := range foundRows {
				foundRow := foundRows[r]
				rowContent := parseClusterGatewayRoutesRow(&foundRow, &cbDB, gw.ClientID)
				rtn = append(rtn, *rowContent)
			}
			aJSON, err := json.Marshal(rtn)
			//fmt.Print("rtn")
			//fmt.Println(rtn)
			if err != nil {
				fmt.Println("JSON parser err: ")
				fmt.Println(err)
			} else if len(rtn) == 0 {
				fmt.Println("No records found in database, not saving to cache.")
			} else {
				cval := b64.StdEncoding.EncodeToString([]byte(aJSON))
				var i ch.Item
				i.Key = key
				i.Value = cval
				//fmt.Print("item: ")
				//fmt.Println(i)
				res := cp.Set(&i)
				if !res.Success {
					fmt.Println("Routes not cached from db for key " + key + ".")
				}
			}
		}
	}
	return &rtn
}

//ClearClusterGwRoutes ClearClusterGwRoutes
func (gw *GatewayRoutes) ClearClusterGwRoutes() bool {
	//fmt.Print("gw ")
	//fmt.Println(gw)
	var cp ch.CProxy
	cp.Host = gw.GwCacheHost
	var cid = strconv.FormatInt(gw.ClientID, 10)
	var key = cid + ":cluster:" + gw.Route
	//fmt.Print("key in clear ")
	//fmt.Println(key)
	rtn := cp.Delete(key)
	return rtn.Success
}

//TripClusterBreaker TripClusterBreaker
func (gw *GatewayRoutes) TripClusterBreaker(b *cb.Breaker) ClusterResponse {
	var rtn ClusterResponse
	var cbDB cb.CircuitBreaker
	cbDB.CacheHost = gw.GwCacheHost
	//fmt.Print("host: ")
	//fmt.Println(cbDB.CacheHost)
	b.ClientID = gw.ClientID
	cbDB.Trip(b)
	gw.ClearClusterGwRoutes()
	rtn.Success = true
	return rtn
}

//ResetClusterBreaker ResetClusterBreaker
func (gw *GatewayRoutes) ResetClusterBreaker(urlID int64) bool {
	var cbDB cb.CircuitBreaker
	cbDB.CacheHost = gw.GwCacheHost
	cbDB.Reset(gw.ClientID, urlID)
	gw.ClearClusterGwRoutes()
	return true
}

func parseClusterGatewayRoutesRow(foundRow *[]string, cbDB *cb.CircuitBreaker, cid int64) *GatewayClusterRouteURL {
	var rtn GatewayClusterRouteURL
	if len(*foundRow) > 0 {
		rtn.RouteID, _ = strconv.ParseInt((*foundRow)[0], 10, 0)
		rtn.Route = (*foundRow)[1]
		rtn.URLID, _ = strconv.ParseInt((*foundRow)[2], 10, 0)
		rtn.Name = (*foundRow)[3]
		rtn.URL = (*foundRow)[4]
		active, err := strconv.ParseBool((*foundRow)[5])
		if err != nil {
			fmt.Print(err)
			rtn.Active = false
		} else {
			rtn.Active = active
		}
		cbs := cbDB.GetStatus(cid, rtn.URLID)
		//fmt.Print("breaker: ")
		//fmt.Print(cbs)
		rtn.OpenFailCode = cbs.OpenFailCode
		rtn.PartialOpen = cbs.PartialOpen
		//rtn.FailoverRouteName = cbs.FailoverRouteName
		rtn.CircuitOpen = cbs.Open
		var b cb.Breaker
		b.ClientID = cid
		b.RestRouteID = rtn.RouteID
		b.RouteURIID = rtn.URLID
		cb := cbDB.GetBreaker(&b)
		rtn.FailureThreshold = cb.FailureThreshold
		rtn.HealthCheckTimeSeconds = cb.HealthCheckTimeSeconds
		rtn.OpenFailCode = cb.OpenFailCode
		rtn.FailoverRouteName = cb.FailoverRouteName
	}
	return &rtn
}
