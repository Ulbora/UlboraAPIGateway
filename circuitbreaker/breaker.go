package circuitbreaker

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

//CircuitBreaker CircuitBreaker
type CircuitBreaker struct {
	DbConfig  db.DbConfig
	CacheHost string
}

//Breaker Breaker
type Breaker struct {
	ID                     int64     `json:"id"`
	FailureThreshold       int       `json:"failureThreshold"`
	FailureCount           int       `json:"failureCount"`
	LastFailureTime        time.Time `json:"lastFailureTime"`
	HealthCheckTimeSeconds int       `json:"healthCheckTimeSeconds"`
	FailoverRouteName      string    `json:"failoverRouteName"`
	OpenFailCode           int       `json:"openFailCode"`
	RouteURIID             int64     `json:"routeUriId"`
	RestRouteID            int64     `json:"routeId"`
	ClientID               int64     `json:"clientId"`
}

//Status of the circuit breaker
type Status struct {
	Warning           bool   `json:"warning"`
	Open              bool   `json:"open"`
	PartialOpen       bool   `json:"partialOpen"`
	FailoverRouteName string `json:"failoverRouteName"`
	OpenFailCode      int    `json:"openFailCode"`
}

type breakerState struct {
	Threshold              int       `json:"threshold"`
	FailCount              int       `json:"failCount"`
	LastFailureTime        time.Time `json:"lastFailureTime"`
	HealthCheckTimeSeconds int       `json:"healthCheckTimeSeconds"`
	FailoverRouteName      string    `json:"failoverRouteName"`
	OpenFailCode           int       `json:"openFailCode"`
}

var cbCache = make(map[string]breakerState)
var mu sync.Mutex

//ConnectDb to database
func (c *CircuitBreaker) ConnectDb() bool {
	rtn := c.DbConfig.ConnectDb()
	if rtn == true {
		fmt.Println("db connect")
	}
	return rtn
}

//InsertBreaker insert
func (c *CircuitBreaker) InsertBreaker(b *Breaker) (bool, error) {
	var success bool
	var err error
	dbConnected := c.DbConfig.ConnectionTest()
	if !dbConnected {
		fmt.Println("reconnection to closed database")
		c.DbConfig.ConnectDb()
	}
	//var a []interface{}
	a := []interface{}{b.FailureThreshold, b.HealthCheckTimeSeconds, b.FailoverRouteName, b.OpenFailCode,
		b.RouteURIID, b.RestRouteID, b.ClientID}
	suc, insID := c.DbConfig.InsertRouteBreaker(a...)
	if suc == true && insID != -1 {
		success = suc
		//fmt.Print("new Id route error id: ")
		//fmt.Println(insID)
	} else {
		err = fmt.Errorf("Failed to insert circuit breaker Record")
	}
	return success, err
}

//UpdateBreaker in database
func (c *CircuitBreaker) UpdateBreaker(b *Breaker) (bool, error) {
	var success bool
	var err error
	dbConnected := c.DbConfig.ConnectionTest()
	if !dbConnected {
		fmt.Println("reconnection to closed database")
		c.DbConfig.ConnectDb()
	}
	a := []interface{}{b.FailureThreshold, b.HealthCheckTimeSeconds, b.FailoverRouteName, b.OpenFailCode,
		b.ID, b.RouteURIID, b.RestRouteID, b.ClientID}
	suc := c.DbConfig.UpdateRouteBreakerConfig(a...)
	if suc == true {
		success = suc
		c.Reset(b.ClientID, b.RouteURIID)
	} else {
		err = fmt.Errorf("Failed to update circuit breaker config Record")
	}
	return success, err
}

//GetStatus of the circuit breaker
func (c *CircuitBreaker) GetStatus(clientID int64, urlID int64) *Status {
	mu.Lock()
	defer mu.Unlock()
	var s Status
	key := strconv.FormatInt(clientID, 10) + "breaker:" + strconv.FormatInt(urlID, 10)
	//fmt.Print("cache get key: ")
	//fmt.Println(key)
	var cs breakerState
	var found bool
	if c.CacheHost != "" {
		//fmt.Print("cache host: ")
		//fmt.Println(c.CacheHost)
		var cp ch.CProxy
		cp.Host = c.CacheHost
		res := cp.Get(key)
		//fmt.Print("cache read in from server in status: ")
		//fmt.Println(res)
		if res.Success == true {
			rJSON, err := b64.StdEncoding.DecodeString(res.Value)
			//fmt.Print("json from cache: ")
			//fmt.Println(rJSON)
			if err != nil {
				fmt.Println(err)
			} else {
				err := json.Unmarshal([]byte(rJSON), &cs)
				if err != nil {
					fmt.Println(err)
				} else {
					//fmt.Print("cache from server: ")
					//fmt.Println(cs)
					found = res.Success
				}
			}
		} else if res.ServiceFailed == true {
			cs, found = cbCache[key]
			//fmt.Print("cache from local after service failed: ")
			//fmt.Println(cs)
		}
	} else {
		cs, found = cbCache[key]
		//fmt.Print("cache from local: ")
		//fmt.Println(cs)
	}
	if found == true {
		var timeExpired bool
		if cs.HealthCheckTimeSeconds != 0 {
			var expireTime = cs.LastFailureTime.Add(time.Second * time.Duration(cs.HealthCheckTimeSeconds))
			if expireTime.Before(time.Now()) {
				timeExpired = true
			}
		}
		if cs.FailCount >= cs.Threshold && timeExpired != true {
			//fmt.Println("setting open")
			s.Warning = true
			s.Open = true
			s.FailoverRouteName = cs.FailoverRouteName
			s.OpenFailCode = cs.OpenFailCode

		} else if cs.FailCount > 0 {
			//fmt.Println("setting partial")
			s.Warning = true
			s.PartialOpen = true
		}
	}
	return &s
}

//Trip the circuit breaker
func (c *CircuitBreaker) Trip(b *Breaker) {
	mu.Lock()
	defer mu.Unlock()
	//var s Status
	key := strconv.FormatInt(b.ClientID, 10) + "breaker:" + strconv.FormatInt(b.RouteURIID, 10)
	//fmt.Print("key in trip in breaker: ")
	//fmt.Println(key)
	var cp ch.CProxy
	var cs breakerState
	var found bool
	var useExCache bool
	//cs, found = cbCache[key]
	//fmt.Print("CacheHost: ")
	//fmt.Println(c.CacheHost)
	if c.CacheHost != "" {
		useExCache = true
		//var cp ch.CProxy
		cp.Host = c.CacheHost
		res := cp.Get(key)
		if res.Success == true {
			rJSON, err := b64.StdEncoding.DecodeString(res.Value)
			//fmt.Print("json from cache server in Trip: ")
			//fmt.Println(rJSON)
			if err != nil {
				fmt.Println(err)
			} else {
				err := json.Unmarshal([]byte(rJSON), &cs)
				if err != nil {
					fmt.Println(err)
				} else {
					//fmt.Print("cache from server in Trip: ")
					//fmt.Println(cs)
					found = res.Success
				}
			}
		} else if res.ServiceFailed == true {
			useExCache = false
			cs, found = cbCache[key]
			//fmt.Print("cache from local in Trip after service failed: ")
			//fmt.Println(cs)
		}
	} else {
		cs, found = cbCache[key]
		//fmt.Print("cache from local in Trip: ")
		//fmt.Println(cs)
	}
	if found == true {
		//fmt.Print("cache found in Trip: ")
		//fmt.Println(found)
		cs.FailCount = cs.FailCount + 1
		cs.LastFailureTime = time.Now()
		if useExCache == true {
			suc := saveToCasheServer(key, cs, cp)
			if suc != true {
				cbCache[key] = cs
			}
		} else {
			cbCache[key] = cs
		}
	} else if b.ClientID != 0 && b.RouteURIID != 0 {
		//fmt.Print("cache found in Trip: ")
		//fmt.Println(found)
		var bs breakerState
		bs.HealthCheckTimeSeconds = b.HealthCheckTimeSeconds
		bs.LastFailureTime = time.Now()
		bs.Threshold = b.FailureThreshold
		bs.FailCount = 1
		bs.FailoverRouteName = b.FailoverRouteName
		bs.OpenFailCode = b.OpenFailCode
		if useExCache == true {
			suc := saveToCasheServer(key, bs, cp)
			if suc != true {
				cbCache[key] = bs
			}
		} else {
			cbCache[key] = bs
		}
		//cbCache[key] = bs
	}
}

//Reset the circuit breaker
func (c *CircuitBreaker) Reset(clientID int64, urlID int64) {
	mu.Lock()
	defer mu.Unlock()
	key := strconv.FormatInt(clientID, 10) + "breaker:" + strconv.FormatInt(urlID, 10)
	if c.CacheHost != "" {
		var cp ch.CProxy
		cp.Host = c.CacheHost
		res := cp.Delete(key)
		if res.Success != true {
			delete(cbCache, key)
		}
	} else {
		delete(cbCache, key)
	}
}

//GetBreaker from database
func (c *CircuitBreaker) GetBreaker(b *Breaker) *Breaker {
	//fmt.Print("b in cb: ")
	//fmt.Println(b)
	a := []interface{}{b.RouteURIID, b.RestRouteID, b.ClientID}
	var rtn *Breaker
	rowPtr := c.DbConfig.GetBreaker(a...)
	if rowPtr != nil {
		foundRow := rowPtr.Row
		rtn = parseCircuitBreakerRow(&foundRow)
	}
	return rtn
}

//DeleteBreaker from database
func (c *CircuitBreaker) DeleteBreaker(b *Breaker) bool {
	a := []interface{}{b.RouteURIID, b.RestRouteID, b.ClientID}
	var success bool
	suc := c.DbConfig.DeleteBreaker(a...)
	if suc == true {
		success = suc
		c.Reset(b.ClientID, b.RouteURIID)
	} else {
		fmt.Println("Failed to delete breaker Record")
	}
	return success
}

//CloseDb connection to database
func (c *CircuitBreaker) CloseDb() bool {
	rtn := c.DbConfig.CloseDb()
	if rtn == true {
		fmt.Println("db connect closed")
	}
	return rtn
}

func parseCircuitBreakerRow(foundRow *[]string) *Breaker {
	var rtn Breaker
	if len(*foundRow) > 0 {
		rtn.ID, _ = strconv.ParseInt((*foundRow)[0], 10, 0)
		rtn.FailureThreshold, _ = strconv.Atoi((*foundRow)[1])
		rtn.HealthCheckTimeSeconds, _ = strconv.Atoi((*foundRow)[2])
		rtn.FailoverRouteName = (*foundRow)[3]
		rtn.OpenFailCode, _ = strconv.Atoi((*foundRow)[4])
		rtn.FailureCount, _ = strconv.Atoi((*foundRow)[5])
		rtn.LastFailureTime, _ = time.Parse("2006-01-02 15:04:05", (*foundRow)[6])
		rtn.RouteURIID, _ = strconv.ParseInt((*foundRow)[7], 10, 0)
		rtn.RestRouteID, _ = strconv.ParseInt((*foundRow)[8], 10, 0)
		rtn.ClientID, _ = strconv.ParseInt((*foundRow)[9], 10, 0)
	}
	return &rtn
}

func saveToCasheServer(key string, bs breakerState, cp ch.CProxy) bool {
	var rtn bool
	//fmt.Print("breakerState being saved in Trip: ")
	//fmt.Println(bs)
	aJSON, err := json.Marshal(bs)
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
		if res.Success != true {
			fmt.Println("Cache server save failed for " + key + ".")
		} else {
			rtn = res.Success
			//fmt.Println("Cache server save success for " + key + ".")
		}
	}
	return rtn
}
