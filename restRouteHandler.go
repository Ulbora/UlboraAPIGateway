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
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	mng "UlboraApiGateway/managers"

	uoauth "github.com/Ulbora/go-ulbora-oauth2"
	"github.com/gorilla/mux"
)

func handleRestRouteChange(w http.ResponseWriter, r *http.Request) {
	auth := getAuth(r)
	me := new(uoauth.Claim)
	me.Role = "admin"
	me.Scope = "write"
	w.Header().Set("Content-Type", "application/json")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		switch r.Method {
		case "POST":
			me.URI = "/rs/gwRestRoute/add"
			valid := auth.Authorize(me)
			if valid != true {
				w.WriteHeader(http.StatusUnauthorized)
			} else {
				rt := new(mng.RestRoute)
				decoder := json.NewDecoder(r.Body)
				error := decoder.Decode(&rt)
				if error != nil {
					log.Println(error.Error())
					http.Error(w, error.Error(), http.StatusBadRequest)
				} else if rt.Route == "" {
					http.Error(w, "bad request", http.StatusBadRequest)
				} else {
					rt.ClientID = auth.ClientID
					resOut := gatewayDB.InsertRestRoute(rt)
					//fmt.Print("response: ")
					//fmt.Println(resOut)
					resJSON, err := json.Marshal(resOut)
					if err != nil {
						log.Println(error.Error())
						http.Error(w, "json output failed", http.StatusInternalServerError)
					}
					w.WriteHeader(http.StatusOK)
					fmt.Fprint(w, string(resJSON))
				}
			}
		case "PUT":
			me.URI = "/rs/gwRestRoute/update"
			valid := auth.Authorize(me)
			if valid != true {
				w.WriteHeader(http.StatusUnauthorized)
			} else {
				rt := new(mng.RestRoute)
				decoder := json.NewDecoder(r.Body)
				error := decoder.Decode(&rt)
				if error != nil {
					log.Println(error.Error())
					http.Error(w, error.Error(), http.StatusBadRequest)
				} else if rt.ID == 0 || rt.Route == "" {
					http.Error(w, "bad request in update", http.StatusBadRequest)
				} else {
					rt.ClientID = auth.ClientID
					resOut := gatewayDB.UpdateRestRoute(rt)
					//fmt.Print("response: ")
					//fmt.Println(resOut)
					resJSON, err := json.Marshal(resOut)
					if err != nil {
						log.Println(error.Error())
						http.Error(w, "json output failed", http.StatusInternalServerError)
					}
					w.WriteHeader(http.StatusOK)
					fmt.Fprint(w, string(resJSON))
				}
			}
		}
	}
}

func handleRestRoute(w http.ResponseWriter, r *http.Request) {
	auth := getAuth(r)
	me := new(uoauth.Claim)
	me.Role = "admin"

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, errID := strconv.ParseInt(vars["id"], 10, 0)
	if errID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	//fmt.Print("id is: ")
	//fmt.Println(id)
	switch r.Method {
	case "GET":
		me.URI = "/rs/gwRestRoute/get"
		me.Scope = "read"
		valid := auth.Authorize(me)
		if valid != true {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			rt := new(mng.RestRoute)
			rt.ID = id
			rt.ClientID = auth.ClientID
			resOut := gatewayDB.GetRestRoute(rt)
			//fmt.Print("response: ")
			//fmt.Println(resOut)
			resJSON, err := json.Marshal(resOut)
			if err != nil {
				log.Println(err.Error())
				http.Error(w, "json output failed", http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(resJSON))
		}

	case "DELETE":
		me.URI = "/rs/gwRestRoute/delete"
		me.Scope = "write"
		valid := auth.Authorize(me)
		if valid != true {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			rt := new(mng.RestRoute)
			rt.ID = id
			rt.ClientID = auth.ClientID
			resOut := gatewayDB.DeleteRestRoute(rt)
			//fmt.Print("response: ")
			//fmt.Println(resOut)
			resJSON, err := json.Marshal(resOut)
			if err != nil {
				log.Println(err.Error())
				http.Error(w, "json output failed", http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(resJSON))
		}
	}
}

func handleRestRouteList(w http.ResponseWriter, r *http.Request) {
	auth := getAuth(r)
	me := new(uoauth.Claim)
	me.Role = "admin"
	me.Scope = "read"
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		me.URI = "/rs/gwRestRoute/list"
		valid := auth.Authorize(me)
		if valid != true {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			rt := new(mng.RestRoute)
			rt.ClientID = auth.ClientID
			resOut := gatewayDB.GetRestRouteList(rt)
			//fmt.Print("response: ")
			//fmt.Println(resOut)
			resJSON, err := json.Marshal(resOut)
			//fmt.Print("response json: ")
			//fmt.Println(string(resJSON))
			if err != nil {
				log.Println(err.Error())
				http.Error(w, "json output failed", http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusOK)
			if string(resJSON) == "null" {
				fmt.Fprint(w, "[]")
			} else {
				fmt.Fprint(w, string(resJSON))
			}
		}
	}
}
