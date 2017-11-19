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

package circuitbreaker

import (
	db "UlboraApiGateway/database"
	"fmt"
	"strconv"
	"sync"
	"time"
)

//CircuitBreaker CircuitBreaker
type CircuitBreaker struct {
	DbConfig db.DbConfig
}

//Breaker Breaker
type Breaker struct {
	ID                     int64
	FailureThreshold       int
	FailureCount           int
	LastFailureTime        time.Time
	HealthCheckTimeSeconds int
	FailoverRouteName      string
	OpenFailCode           int
	RouteURIID             int64
	RestRouteID            int64
	ClientID               int64
}

//Status of the circuit breaker
type Status struct {
	Warning     bool
	Open        bool
	PartialOpen bool
}

type breakerState struct {
	threshold              int
	failCount              int
	lastFailureTime        time.Time
	healthCheckTimeSeconds int
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
	key := strconv.FormatInt(clientID, 10) + ":" + strconv.FormatInt(urlID, 10)
	cs, found := cbCache[key]
	fmt.Print("cache: ")
	fmt.Println(cs)
	if found == true {
		var timeExpired bool
		if cs.healthCheckTimeSeconds != 0 {
			var expireTime = cs.lastFailureTime.Add(time.Second * time.Duration(cs.healthCheckTimeSeconds))
			if expireTime.Before(time.Now()) {
				timeExpired = true
			}
		}
		if cs.failCount >= cs.threshold && timeExpired != true {
			fmt.Print("setting open")
			s.Warning = true
			s.Open = true
		} else if cs.failCount > 0 {
			fmt.Print("setting partial")
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
	key := strconv.FormatInt(b.ClientID, 10) + ":" + strconv.FormatInt(b.RouteURIID, 10)
	cs, found := cbCache[key]
	if found == true {
		cs.failCount = cs.failCount + 1
		cs.lastFailureTime = time.Now()
		cbCache[key] = cs
	} else {
		var bs breakerState
		bs.healthCheckTimeSeconds = b.HealthCheckTimeSeconds
		bs.lastFailureTime = time.Now()
		bs.threshold = b.FailureThreshold
		bs.failCount = 1
		cbCache[key] = bs
	}
}

//Reset the circuit breaker
func (c *CircuitBreaker) Reset(clientID int64, urlID int64) {
	mu.Lock()
	defer mu.Unlock()
	key := strconv.FormatInt(clientID, 10) + ":" + strconv.FormatInt(urlID, 10)
	delete(cbCache, key)
}

//GetBreaker from database
func (c *CircuitBreaker) GetBreaker(b *Breaker) *Breaker {
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
