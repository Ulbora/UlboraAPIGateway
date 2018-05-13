package monitor

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
	db "UlboraApiGateway/database"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"
)

//GatewayPerformanceMonitor error monitor
type GatewayPerformanceMonitor struct {
	DbConfig      db.DbConfig
	CacheHost     string
	CallBatchSize int64
}

//GwPerformance GwPerformance
type GwPerformance struct {
	ID             int64     `json:"id"`
	Calls          int64     `json:"calls"`
	LatencyMsTotal int64     `json:"latencyMsTotal"`
	Entered        time.Time `json:"entered"`
	RouteURIID     int64     `json:"routeUriId"`
	RestRouteID    int64     `json:"routeId"`
	ClientID       int64     `json:"clientId"`
}

type routePerformance struct {
	Calls          int64 `json:"calls"`
	LatencyMsTotal int64 `json:"LatencyMsTotal"`
}

var perCache = make(map[string]routePerformance)
var mu sync.Mutex

//ConnectDb to database
func (g *GatewayPerformanceMonitor) ConnectDb() bool {
	rtn := g.DbConfig.ConnectDb()
	if rtn {
		fmt.Println("db connect")
	}
	return rtn
}

//InsertRoutePerformance insert
func (g *GatewayPerformanceMonitor) InsertRoutePerformance(e *GwPerformance) (bool, error) {
	var success bool
	var err error
	//fmt.Println("testing db")
	//fmt.Println(g.DbConfig)
	dbConnected := g.DbConfig.ConnectionTest()
	//fmt.Print("db connected: ")
	//fmt.Println(dbConnected)
	if !dbConnected {
		fmt.Println("reconnection to closed database")
		g.DbConfig.ConnectDb()
	}
	//var a []interface{}
	a := []interface{}{e.Calls, e.LatencyMsTotal, e.Entered, e.RouteURIID, e.RestRouteID, e.ClientID}
	suc, insID := g.DbConfig.InsertRoutePerformance(a...)
	if suc && insID != -1 {
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

//SaveRoutePerformance updates route performance in cache
func (g *GatewayPerformanceMonitor) SaveRoutePerformance(clientID int64, routeID int64, urlID int64, latency int64) bool {
	g.DeleteRoutePerformance()
	mu.Lock()
	defer mu.Unlock()
	//fmt.Print("per insert client1 :")
	//fmt.Println(clientID)
	var rtn bool
	//var s Status
	key := strconv.FormatInt(clientID, 10) + "perf:" + strconv.FormatInt(urlID, 10)
	//fmt.Print("key: ")
	//fmt.Println(key)

	var found bool
	var useExCache bool
	var cp ch.CProxy
	var rp routePerformance
	if g.CacheHost != "" {
		useExCache = true
		cp.Host = g.CacheHost
		res := cp.Get(key)
		if res.Success {
			rJSON, err := b64.StdEncoding.DecodeString(res.Value)
			//fmt.Print("json from cache: ")
			//fmt.Println(rJSON)
			if err != nil {
				fmt.Println(err)
			} else {
				err := json.Unmarshal([]byte(rJSON), &rp)
				if err != nil {
					fmt.Println(err)
				} else {
					//fmt.Print("cache from server: ")
					//fmt.Println(rp)
					found = res.Success
				}
			}
		} else if res.ServiceFailed {
			useExCache = false
			rp, found = perCache[key]
			//fmt.Print("cache from local after service failed: ")
			//fmt.Println(rp)
		}
	} else {
		rp, found = perCache[key]
		//fmt.Print("cache from local: ")
		//fmt.Println(rp)
	}
	//fmt.Print("per insert client2 :")
	//fmt.Println(clientID)
	//fmt.Print("found :")
	//fmt.Println(found)
	if found {
		rp.Calls = rp.Calls + 1
		rp.LatencyMsTotal = rp.LatencyMsTotal + latency
		if rp.Calls >= g.CallBatchSize {
			//sendToDbAndClear(rp, cp, useExCache)
			//fmt.Print("per insert client3 :")
			//fmt.Println(clientID)
			//fmt.Println("saving to database and clearing cache")
			var p GwPerformance
			p.ClientID = clientID
			p.Calls = rp.Calls
			p.Entered = time.Now()
			p.LatencyMsTotal = rp.LatencyMsTotal
			p.RestRouteID = routeID
			p.RouteURIID = urlID
			//fmt.Print("per insert client :")
			//fmt.Println(clientID)
			suc, err := g.InsertRoutePerformance(&p)
			if err != nil {
				fmt.Println(err)
			}
			//fmt.Print("Insert Res: ")
			//fmt.Println(suc)
			if suc {
				if useExCache {
					//var cp ch.CProxy
					//cp.Host = c.CacheHost
					res := cp.Delete(key)
					if !res.Success {
						delete(perCache, key)
						rtn = true
					} else {
						rtn = true
					}
				} else {
					delete(perCache, key)
					rtn = true
				}
			} else {
				if useExCache {
					suc := saveToCasheServer(key, rp, cp)
					if !suc {
						perCache[key] = rp
						rtn = true
					} else {
						rtn = true
					}
				} else {
					perCache[key] = rp
					rtn = true
				}
			}
		} else {
			if useExCache {
				suc := saveToCasheServer(key, rp, cp)
				if !suc {
					perCache[key] = rp
					rtn = true
				} else {
					rtn = true
				}
			} else {
				perCache[key] = rp
				rtn = true
			}
			//rtn = true
		}
	} else {
		//var rp routePerformance
		rp.Calls = 1
		rp.LatencyMsTotal = latency
		if useExCache {
			suc := saveToCasheServer(key, rp, cp)
			if !suc {
				perCache[key] = rp
			}
		} else {
			perCache[key] = rp
		}
		rtn = true
	}
	return rtn
}

//DeleteRoutePerformance from database
func (g *GatewayPerformanceMonitor) DeleteRoutePerformance() bool {
	a := []interface{}{} //{e.RouteURIID, e.RestRouteID, e.ClientID}
	var success bool
	suc := g.DbConfig.DeleteRoutePerformance(a...)
	if suc {
		success = suc
	} // else {
	//fmt.Println("Failed to delete performance Record")
	//}
	return success
}

//CloseDb connection to database
func (g *GatewayPerformanceMonitor) CloseDb() bool {
	rtn := g.DbConfig.CloseDb()
	if rtn {
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

// func sendToDbAndClear(rp routePerformance, cp ch.CProxy, externalCache bool) {
// 	var p GwPerformance
// 	p.ClientID = clientID
// 	p.Calls = 500
// 	p.Entered = time.Now().Add(time.Hour * -2400)
// 	p.LatencyMsTotal = 10000
// 	p.RestRouteID = routeID
// 	p.RouteURIID = routeURLID
// 	suc, err := gatewayDB.InsertRoutePerformance(&p)
// }

func saveToCasheServer(key string, rp routePerformance, cp ch.CProxy) bool {
	var rtn bool
	//fmt.Print("breakerState being saved in Trip: ")
	//fmt.Println(bs)
	aJSON, err := json.Marshal(rp)
	//fmt.Print("json being saved to server in Trip: ")
	//fmt.Println(aJSON)
	if err != nil {
		fmt.Println(err)
	} else {
		cval := b64.StdEncoding.EncodeToString([]byte(aJSON))
		var i ch.Item
		i.Key = key
		i.Value = cval
		res := cp.Set(&i)
		if !res.Success {
			fmt.Println("Cache server save failed for " + key + ".")
		} else {
			rtn = res.Success
			//fmt.Println("Cache server save success for " + key + ".")
		}
	}
	return rtn
}
