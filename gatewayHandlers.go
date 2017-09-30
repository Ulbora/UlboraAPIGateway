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

package main

import (
	mgr "UlboraApiGateway/managers"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func handleActiveRoute(w http.ResponseWriter, r *http.Request) {
	var gwr mgr.GatewayRoutes
	gwr.ClientID = 1
	gwr.GwCacheHost = getCacheHost()
	gwr.GwDB = gatewayDB

	var rtn string
	var rtnCode int
	vars := mux.Vars(r)
	route := vars["route"]
	fpath := vars["fpath"]
	code := r.URL.Query()
	gwr.Route = route
	rts := gwr.GetGatewayRoutes(true, "")

	fmt.Print("route: ")
	fmt.Println(route)
	fmt.Print("fpath: ")
	fmt.Println(fpath)
	fmt.Print("code: ")
	fmt.Println(code)
	fmt.Println("Found url: " + rts.URL)
	if rts.URL == "" {
		fmt.Println("No route found")
		rtnCode = 400
		rtn = "bad route"
	} else {
		switch r.Method {
		case "POST", "PUT", "PATCH":
			fmt.Print("found routes: ")
			fmt.Println(rts)
			var spath = rts.URL + "/" + fpath + parseQueryString(code)
			fmt.Print("spath: ")
			fmt.Println(spath)
			//body := r.Body.Read()
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Print("Body: ")
				fmt.Println(string(body))
			}
			req, rErr := http.NewRequest(r.Method, spath, bytes.NewBuffer(body))
			if rErr != nil {
				fmt.Print("request err: ")
				fmt.Println(rErr)
			} else {
				req.Header.Set("Content-Type", r.Header.Get("Content-Type"))
				client := &http.Client{}
				resp, cErr := client.Do(req)
				if cErr != nil {
					fmt.Print("Request err: ")
					fmt.Println(cErr)
					rtnCode = 400
					rtn = cErr.Error()
				} else {
					defer resp.Body.Close()
					respbody, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						fmt.Println(err)
						rtnCode = 500
						rtn = err.Error()
					} else {
						rtn = string(respbody)
						fmt.Print("Resp Body: ")
						fmt.Println(rtn)
						rtnCode = resp.StatusCode
						w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
					}
				}
			}
		case "GET":
			var spath = rts.URL + "/" + fpath + parseQueryString(code)
			fmt.Print("api path: ")
			fmt.Println(spath)
			resp, err := http.Get(spath)
			fmt.Print("res: ")
			fmt.Println(resp)
			if err != nil {
				fmt.Println(err)
				rtnCode = 400
				rtn = err.Error()
			} else {
				defer resp.Body.Close()
				respbody, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
					rtnCode = 500
					rtn = err.Error()
				} else {
					rtn = string(respbody)
					fmt.Print("Resp Body: ")
					fmt.Println(rtn)
					rtnCode = resp.StatusCode
					w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
				}
			}
		case "DELETE":
			var spath = rts.URL + "/" + fpath + parseQueryString(code)
			fmt.Print("fpath: ")
			fmt.Println(fpath)
			code := r.URL.Query()
			fmt.Println(code)
			req, rErr := http.NewRequest(r.Method, spath, nil)
			if rErr != nil {
				fmt.Print("request err: ")
				fmt.Println(rErr)
			} else {
				//req.Header.Set("Content-Type", r.Header.Get("Content-Type"))
				client := &http.Client{}
				resp, cErr := client.Do(req)
				if cErr != nil {
					fmt.Print("Request err: ")
					fmt.Println(cErr)
					rtnCode = 400
					rtn = cErr.Error()
				} else {
					defer resp.Body.Close()
					respbody, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						fmt.Println(err)
						rtnCode = 500
						rtn = err.Error()
					} else {
						rtn = string(respbody)
						fmt.Print("Resp Body: ")
						fmt.Println(rtn)
						rtnCode = resp.StatusCode
						w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
					}
				}
			}
		case "OPTIONS":
			var spath = rts.URL + "/" + fpath + parseQueryString(code)
			req, rErr := http.NewRequest(r.Method, spath, nil)
			if rErr != nil {
				fmt.Print("request err: ")
				fmt.Println(rErr)
			} else {
				client := &http.Client{}
				resp, cErr := client.Do(req)
				if cErr != nil {
					fmt.Print("Request err: ")
					fmt.Println(cErr)
					rtnCode = 400
					rtn = cErr.Error()
				} else {
					defer resp.Body.Close()
					respbody, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						fmt.Println(err)
						rtnCode = 500
						rtn = err.Error()
					} else {
						rtn = string(respbody)
						fmt.Print("Resp Body: ")
						fmt.Println(rtn)
						rtnCode = resp.StatusCode
						w.Header().Set("access-control-allow-headers", resp.Header.Get("access-control-allow-headers"))
						w.Header().Set("access-control-allow-methods", resp.Header.Get("access-control-allow-methods"))
						w.Header().Set("access-control-allow-origin", resp.Header.Get("access-control-allow-origin"))
						w.Header().Set("connection", resp.Header.Get("connection"))
						w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
					}
				}
			}
		}
	}
	w.WriteHeader(rtnCode)
	fmt.Fprint(w, rtn)
}
