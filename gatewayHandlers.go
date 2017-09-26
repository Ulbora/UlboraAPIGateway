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
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func handleActiveRoute(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	//cType := r.Header.Get("Content-Type")
	//h := r.
	//h.WriteSubset
	//http.Redirect(w, r, path, http.StatusFound)
	switch r.Method {
	case "POST", "PUT", "PATCH":
		vars := mux.Vars(r)
		route := vars["route"]
		fpath := vars["fpath"]
		var spath = "http://localhost:3003" + "/" + fpath
		fmt.Print("route: ")
		fmt.Println(route)
		fmt.Print("fpath: ")
		fmt.Println(fpath)
		code := r.URL.Query()
		fmt.Println(code)
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
			var rtn string
			var rtnCode int
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
				}
				//decoder := json.NewDecoder(resp.Body)
				//error := decoder.Decode(&rtn)
				//if error != nil {
				//log.Println(error.Error())
				//}

			}
			w.WriteHeader(rtnCode)
			fmt.Fprint(w, rtn)
		}

	//case "PUT":

	//case "PATCH":

	case "GET", "DELETE":
		vars := mux.Vars(r)
		route := vars["route"]
		fpath := vars["fpath"]
		fmt.Print("route: ")
		fmt.Println(route)
		fmt.Print("fpath: ")
		fmt.Println(fpath)
		code := r.URL.Query()
		fmt.Println(code)
	//case "DELETE":

	case "OPTIONS":

	}

	w.WriteHeader(http.StatusOK)
}
