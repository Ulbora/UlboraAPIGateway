package main

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
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	mng "UlboraApiGateway/managers"

	uoauth "github.com/Ulbora/go-ulbora-oauth2"
	"github.com/gorilla/mux"
)

func handleClientChange(w http.ResponseWriter, r *http.Request) {
	auth := getAuth(r)
	me := new(uoauth.Claim)
	me.Role = "superAdmin"
	me.Scope = "write"
	w.Header().Set("Content-Type", "application/json")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		switch r.Method {
		case "POST":
			me.URI = "/ulbora/rs/gwClient/add"
			valid := auth.Authorize(me)
			if valid != true {
				w.WriteHeader(http.StatusUnauthorized)
			} else {
				client := new(mng.Client)
				decoder := json.NewDecoder(r.Body)
				error := decoder.Decode(&client)
				if error != nil {
					log.Println(error.Error())
					http.Error(w, error.Error(), http.StatusBadRequest)
				} else if client.ClientID == 0 || client.APIKey == "" || client.Level == "" {
					http.Error(w, "bad request", http.StatusBadRequest)
				} else {
					resOut := gatewayDB.InsertClient(client)
					//fmt.Print("response: ")
					//fmt.Println(resOut)
					resJSON, err := json.Marshal(resOut)
					if err != nil {
						log.Println(error.Error())
						//http.Error(w, "json output failed", http.StatusInternalServerError)
					}
					w.WriteHeader(http.StatusOK)
					fmt.Fprint(w, string(resJSON))
				}
			}
		case "PUT":
			me.URI = "/ulbora/rs/gwClient/update"
			valid := auth.Authorize(me)
			if valid != true {
				w.WriteHeader(http.StatusUnauthorized)
			} else {
				client := new(mng.Client)
				decoder := json.NewDecoder(r.Body)
				error := decoder.Decode(&client)
				if error != nil {
					log.Println(error.Error())
					http.Error(w, error.Error(), http.StatusBadRequest)
				} else if client.APIKey == "" || client.Level == "" || client.ClientID == 0 {
					http.Error(w, "bad request in update", http.StatusBadRequest)
				} else {
					resOut := gatewayDB.UpdateClient(client)
					//fmt.Print("response: ")
					//fmt.Println(resOut)
					resJSON, err := json.Marshal(resOut)
					if err != nil {
						log.Println(error.Error())
						//http.Error(w, "json output failed", http.StatusInternalServerError)
					}
					w.WriteHeader(http.StatusOK)
					fmt.Fprint(w, string(resJSON))
				}
			}
		}
	}
}

func handleClient(w http.ResponseWriter, r *http.Request) {
	auth := getAuth(r)
	me := new(uoauth.Claim)
	me.Role = "superAdmin"

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	clientID, errID := strconv.ParseInt(vars["clientId"], 10, 0)
	if errID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	//fmt.Print("id is: ")
	//fmt.Println(id)
	switch r.Method {
	case "GET":
		me.URI = "/ulbora/rs/gwClient/get"
		me.Scope = "read"
		valid := auth.Authorize(me)
		if valid != true {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			client := new(mng.Client)
			client.ClientID = clientID
			resOut := gatewayDB.GetClient(client)
			//fmt.Print("response: ")
			//fmt.Println(resOut)
			resJSON, err := json.Marshal(resOut)
			if err != nil {
				log.Println(err.Error())
				//http.Error(w, "json output failed", http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(resJSON))
		}

	case "DELETE":
		me.URI = "/ulbora/rs/gwClient/delete"
		me.Scope = "write"
		valid := auth.Authorize(me)
		if valid != true {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			client := new(mng.Client)
			client.ClientID = clientID
			resOut := gatewayDB.DeleteClient(client)
			//fmt.Print("response: ")
			//fmt.Println(resOut)
			resJSON, err := json.Marshal(resOut)
			if err != nil {
				log.Println(err.Error())
				//http.Error(w, "json output failed", http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, string(resJSON))
		}
	}
}

func handleClientList(w http.ResponseWriter, r *http.Request) {
	auth := getAuth(r)
	me := new(uoauth.Claim)
	me.Role = "superAdmin"
	me.Scope = "write"
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		me.URI = "/ulbora/rs/gwClient/list"
		valid := auth.Authorize(me)
		if valid != true {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			client := new(mng.Client)
			resOut := gatewayDB.GetClientList(client)
			//fmt.Print("response: ")
			//fmt.Println(resOut)
			resJSON, err := json.Marshal(resOut)
			//fmt.Print("response json: ")
			//fmt.Println(string(resJSON))
			if err != nil {
				log.Println(err.Error())
				//http.Error(w, "json output failed", http.StatusInternalServerError)
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
