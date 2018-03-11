package handlers

import (
	//"bytes"
	"fmt"
	//"io/ioutil"
	"net/http"
	"time"
)

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

func doDelete(p *passParams) *returnVals {
	//fmt.Print("found routes: ")
	//fmt.Println(rts)
	var rtnVals returnVals
	var rtn string
	var rtnCode int

	//var sTime1 = time.Now()
	var sTime2 time.Time
	var eTime1 time.Time
	if paramsOK(p) {
		var spath = p.rts.URL + "/" + p.fpath + parseQueryString(*p.code)
		//fmt.Print("fpath: ")
		//fmt.Println(fpath)
		//code := r.URL.Query()
		//fmt.Println(code)
		req, rErr := http.NewRequest(p.r.Method, spath, nil)
		if rErr != nil {
			fmt.Print("request err: ")
			fmt.Println(rErr)
			rtnCode = 400
			rtn = rErr.Error()
		} else {
			buildHeaders(p.r, req)
			client := &http.Client{}
			eTime1 = time.Now()
			resp, cErr := client.Do(req)
			sTime2 = time.Now()
			if cErr != nil {
				fmt.Print("Gateway err: ")
				fmt.Println(cErr)
				rtnCode = 400
				rtn = cErr.Error()
				cbk := p.h.CbDB.GetBreaker(p.b)
				p.h.CbDB.Trip(cbk)
				go p.h.ErrDB.SaveRouteError(p.gwr.ClientID, 400, cErr.Error(), p.rts.RouteID, p.rts.URLID)
			} else {
				defer resp.Body.Close()
				respbody, err := processResponse(resp) //:= ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Print("Resp Body err: ")
					fmt.Println(err)
					rtnCode = 500
					rtn = err.Error()
					cbk := p.h.CbDB.GetBreaker(p.b)
					p.h.CbDB.Trip(cbk)
					go p.h.ErrDB.SaveRouteError(p.gwr.ClientID, 500, err.Error(), p.rts.RouteID, p.rts.URLID)
				} else {
					rtn = string(respbody)
					//fmt.Print("Resp Body: ")
					//fmt.Println(rtn)
					rtnCode = resp.StatusCode
					if rtnCode != http.StatusOK {
						go p.h.ErrDB.SaveRouteError(p.gwr.ClientID, rtnCode, resp.Status, p.rts.RouteID, p.rts.URLID)
					} else {
						go p.h.CbDB.Reset(p.gwr.ClientID, p.rts.URLID)
					}
					buildRespHeaders(resp, p.w)
					//w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
				}
			}
		}
	}
	rtnVals.rtnCode = rtnCode
	rtnVals.rtn = rtn
	rtnVals.eTime1 = eTime1
	rtnVals.sTime2 = sTime2
	return &rtnVals
}
