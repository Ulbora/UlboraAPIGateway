package handlers

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
	cb "UlboraApiGateway/circuitbreaker"
	mgr "UlboraApiGateway/managers"
	//"bytes"
	"fmt"
	//"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type passParams struct {
	h     *Handler
	rts   *mgr.GatewayRouteURL
	fpath string
	code  *url.Values
	gwr   *mgr.GatewayRoutes
	b     *cb.Breaker
	w     http.ResponseWriter
	r     *http.Request
}

type returnVals struct {
	rtnCode int
	rtn     string
	eTime1  time.Time
	sTime2  time.Time
}

//HandleGwRoute HandleGwRoute
func (h Handler) HandleGwRoute(w http.ResponseWriter, r *http.Request) {
	var gatewayDB mgr.GatewayDB
	gatewayDB.DbConfig = h.DbConfig
	gatewayDB.GwCacheHost = getCacheHost()

	var sTime1 = time.Now()
	var sTime2 time.Time
	var eTime1 time.Time
	var eTime2 time.Time

	var gwr mgr.GatewayRoutes

	cid := r.Header.Get("u-client-id")
	gwr.ClientID, _ = strconv.ParseInt((cid), 10, 0)
	gwr.APIKey = r.Header.Get("u-api-key")
	//fmt.Print("apiKey: ")
	//fmt.Println(gwr.APIKey)
	gwr.GwCacheHost = getCacheHost()
	gwr.GwDB = gatewayDB

	var rtn string
	var rtnCode int

	var route string
	var rName string
	var fpath string
	//var code string

	vars := mux.Vars(r)
	if vars != nil {
		route = vars["route"]
		rName = vars["rname"]
		fpath = vars["fpath"]
	} else {
		route = r.URL.Query().Get("route")
		rName = r.URL.Query().Get("rname")
		fpath = r.URL.Query().Get("fpath")
	}

	code := r.URL.Query()
	gwr.Route = route
	var activeRoute = true
	if rName != "" {
		activeRoute = false
	}
	//fmt.Println("getting route active: " + rName)
	//fmt.Print("active: ")
	//fmt.Println(activeRoute)
	rts := gwr.GetGatewayRoutes(activeRoute, rName)
	//fmt.Print("routes: ")
	//fmt.Println(rts)
	var b cb.Breaker
	b.ClientID = gwr.ClientID
	b.RestRouteID = rts.RouteID
	b.RouteURIID = rts.URLID
	var p passParams
	p.b = &b
	p.code = &code
	p.fpath = fpath
	p.gwr = &gwr
	p.h = &h
	p.r = r
	p.rts = rts
	p.w = w
	// fmt.Print("route: ")
	// fmt.Println(route)
	// fmt.Print("fpath: ")
	// fmt.Println(fpath)
	// fmt.Print("rName: ")
	// fmt.Println(rName)
	// fmt.Print("code: ")
	// fmt.Println(code)
	// fmt.Println("Found url: " + rts.URL)
	if rts.URL == "" {
		fmt.Println("No route found in gateway")
		rtnCode = rts.OpenFailCode
		rtn = "bad route"
		//fmt.Print("found routes: ")
		//fmt.Println(rts)
	} else if rts.CircuitOpen {
		fmt.Println("Circuit breaker is open for this route")
		rtnCode = rts.OpenFailCode
		rtn = "Circuit open"
		//fmt.Print("found route: ")
		//fmt.Println(rts)
	} else {
		switch r.Method {
		case "POST", "PUT", "PATCH":
			pppRtn := doPostPutPatch(&p)
			rtn = pppRtn.rtn
			rtnCode = pppRtn.rtnCode
			eTime1 = pppRtn.eTime1
			sTime2 = pppRtn.sTime2
		case "GET":
			pppRtn := doGet(&p)
			rtn = pppRtn.rtn
			rtnCode = pppRtn.rtnCode
			eTime1 = pppRtn.eTime1
			sTime2 = pppRtn.sTime2
		case "DELETE":
			pppRtn := doDelete(&p)
			rtn = pppRtn.rtn
			rtnCode = pppRtn.rtnCode
			eTime1 = pppRtn.eTime1
			sTime2 = pppRtn.sTime2
		case "OPTIONS":
			pppRtn := doOptions(&p)
			rtn = pppRtn.rtn
			rtnCode = pppRtn.rtnCode
			eTime1 = pppRtn.eTime1
			sTime2 = pppRtn.sTime2
		}
	}
	eTime2 = time.Now()
	dif1 := eTime1.Sub(sTime1)
	dif2 := eTime2.Sub(sTime2)
	tots := dif1.Seconds() + dif2.Seconds()
	//sec := dif.Seconds()
	//fmt.Print("latency sec: ")
	//fmt.Println(tots)
	pms := (tots * 1000000)
	//fmt.Print("latency micros: ")
	//fmt.Println(pms)
	rms := int64(pms + .5)
	//fmt.Print("rounded latency micros: ")
	//fmt.Println(rms)
	go h.MonDB.SaveRoutePerformance(gwr.ClientID, rts.RouteID, rts.URLID, rms)
	w.WriteHeader(rtnCode)
	fmt.Fprint(w, rtn)
}
