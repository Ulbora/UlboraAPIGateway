package handlers

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

//HandleRestRouteSuperPost HandleRestRouteSuperPost
func (h Handler) HandleRestRouteSuperPost(w http.ResponseWriter, r *http.Request) {
	var gatewayDB mng.GatewayDB
	gatewayDB.DbConfig = h.DbConfig
	gatewayDB.GwCacheHost = getCacheHost()

	auth := getAuth(r)
	rrsme := new(uoauth.Claim)
	rrsme.Role = "superAdmin"
	rrsme.Scope = "write"
	w.Header().Set("Content-Type", "application/json")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		switch r.Method {
		case "POST":
			rrsme.URI = "/ulbora/rs/gwRestRouteSuper/add"
			var valid bool
			if testMode {
				valid = true
			} else {
				valid = auth.Authorize(rrsme)
			}
			if !valid {
				w.WriteHeader(http.StatusUnauthorized)
			} else {
				rt := new(mng.RestRoute)
				decoder := json.NewDecoder(r.Body)
				error := decoder.Decode(&rt)
				if error != nil {
					log.Println(error.Error())
					http.Error(w, error.Error(), http.StatusBadRequest)
				} else if rt.ClientID == 0 || rt.Route == "" {
					http.Error(w, "bad request", http.StatusBadRequest)
				} else {
					resOut := gatewayDB.InsertRestRoute(rt)
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
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

//HandleRestRouteSuperPut HandleRestRouteSuperPut
func (h Handler) HandleRestRouteSuperPut(w http.ResponseWriter, r *http.Request) {
	var gatewayDB mng.GatewayDB
	gatewayDB.DbConfig = h.DbConfig
	gatewayDB.GwCacheHost = getCacheHost()

	auth := getAuth(r)
	rspme := new(uoauth.Claim)
	rspme.Role = "superAdmin"
	rspme.Scope = "write"
	w.Header().Set("Content-Type", "application/json")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		switch r.Method {
		case "PUT":
			rspme.URI = "/ulbora/rs/gwRestRouteSuper/update"
			var valid bool
			if testMode {
				valid = true
			} else {
				valid = auth.Authorize(rspme)
			}
			if !valid {
				w.WriteHeader(http.StatusUnauthorized)
			} else {
				rt := new(mng.RestRoute)
				decoder := json.NewDecoder(r.Body)
				error := decoder.Decode(&rt)
				if error != nil {
					log.Println(error.Error())
					http.Error(w, error.Error(), http.StatusBadRequest)
				} else if rt.ClientID == 0 || rt.ID == 0 || rt.Route == "" {
					http.Error(w, "bad request in update", http.StatusBadRequest)
				} else {
					resOut := gatewayDB.UpdateRestRoute(rt)
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
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

//HandleRestRouteSuperGet HandleRestRouteSuperGet
func (h Handler) HandleRestRouteSuperGet(w http.ResponseWriter, r *http.Request) {
	var gatewayDB mng.GatewayDB
	gatewayDB.DbConfig = h.DbConfig
	gatewayDB.GwCacheHost = getCacheHost()

	auth := getAuth(r)
	rsgme := new(uoauth.Claim)
	rsgme.Role = "superAdmin"

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	var id int64
	var sgClientID int64
	var errID error
	var errCID error

	if vars != nil {
		id, errID = strconv.ParseInt(vars["id"], 10, 0)
		sgClientID, errCID = strconv.ParseInt(vars["clientId"], 10, 0)
	} else {
		var idStr = r.URL.Query().Get("id")
		id, errID = strconv.ParseInt(idStr, 10, 0)

		var clientIDStr = r.URL.Query().Get("clientId")
		sgClientID, errCID = strconv.ParseInt(clientIDStr, 10, 0)
	}
	if errID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	if errCID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	//fmt.Print("id is: ")
	//fmt.Println(id)
	switch r.Method {
	case "GET":
		rsgme.URI = "/ulbora/rs/gwRestRouteSuper/get"
		rsgme.Scope = "read"
		var valid bool
		if testMode {
			valid = true
		} else {
			valid = auth.Authorize(rsgme)
		}
		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			rt := new(mng.RestRoute)
			rt.ID = id
			rt.ClientID = sgClientID
			resOut := gatewayDB.GetRestRoute(rt)
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
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

//HandleRestRouteSuperDelete HandleRestRouteSuperDelete
func (h Handler) HandleRestRouteSuperDelete(w http.ResponseWriter, r *http.Request) {
	var gatewayDB mng.GatewayDB
	gatewayDB.DbConfig = h.DbConfig
	gatewayDB.GwCacheHost = getCacheHost()

	auth := getAuth(r)
	rsdme := new(uoauth.Claim)
	rsdme.Role = "superAdmin"

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	var id int64
	var scClientID int64
	var errID error
	var errCID error

	if vars != nil {
		id, errID = strconv.ParseInt(vars["id"], 10, 0)
		scClientID, errCID = strconv.ParseInt(vars["clientId"], 10, 0)
	} else {
		var idStr = r.URL.Query().Get("id")
		id, errID = strconv.ParseInt(idStr, 10, 0)

		var clientIDStr = r.URL.Query().Get("clientId")
		scClientID, errCID = strconv.ParseInt(clientIDStr, 10, 0)
	}
	if errID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	if errCID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	//fmt.Print("id is: ")
	//fmt.Println(id)
	switch r.Method {
	case "DELETE":
		rsdme.URI = "/ulbora/rs/gwRestRouteSuper/delete"
		rsdme.Scope = "write"
		var valid bool
		if testMode {
			valid = true
		} else {
			valid = auth.Authorize(rsdme)
		}
		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			rt := new(mng.RestRoute)
			rt.ID = id
			rt.ClientID = scClientID
			resOut := gatewayDB.DeleteRestRoute(rt)
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
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

//HandleRestRouteSuperList HandleRestRouteSuperList
func (h Handler) HandleRestRouteSuperList(w http.ResponseWriter, r *http.Request) {
	var gatewayDB mng.GatewayDB
	gatewayDB.DbConfig = h.DbConfig
	gatewayDB.GwCacheHost = getCacheHost()

	auth := getAuth(r)
	me := new(uoauth.Claim)
	me.Role = "superAdmin"
	me.Scope = "read"
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	//var id int64
	var clientID int64
	//var errID error
	var rslErrCID error

	if vars != nil {
		//id, errID = strconv.ParseInt(vars["id"], 10, 0)
		clientID, rslErrCID = strconv.ParseInt(vars["clientId"], 10, 0)
	} else {
		//var idStr = r.URL.Query().Get("id")
		//id, errID = strconv.ParseInt(idStr, 10, 0)

		var clientIDStr = r.URL.Query().Get("clientId")
		clientID, rslErrCID = strconv.ParseInt(clientIDStr, 10, 0)
	}
	//if errID != nil {
	//	http.Error(w, "bad request", http.StatusBadRequest)
	//}
	if rslErrCID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}

	// clientID, errCID := strconv.ParseInt(vars["clientId"], 10, 0)
	// if errCID != nil {
	// 	http.Error(w, "bad request", http.StatusBadRequest)
	// }
	switch r.Method {
	case "GET":
		me.URI = "/ulbora/rs/gwRestRouteSuper/list"
		var valid bool
		if testMode {
			valid = true
		} else {
			valid = auth.Authorize(me)
		}
		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			rt := new(mng.RestRoute)
			rt.ClientID = clientID
			resOut := gatewayDB.GetRestRouteList(rt)
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
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
