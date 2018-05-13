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

//HandleRouteURLSuperPost HandleRouteURLSuperPost
func (h Handler) HandleRouteURLSuperPost(w http.ResponseWriter, r *http.Request) {
	var gatewayDB mng.GatewayDB
	gatewayDB.DbConfig = h.DbConfig
	gatewayDB.GwCacheHost = getCacheHost()

	auth := getAuth(r)
	uspme := new(uoauth.Claim)
	uspme.Role = "superAdmin"
	uspme.Scope = "write"
	w.Header().Set("Content-Type", "application/json")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		switch r.Method {
		case "POST":
			uspme.URI = "/ulbora/rs/gwRouteUrlSuper/add"
			var valid bool
			if testMode {
				valid = true
			} else {
				valid = auth.Authorize(uspme)
			}
			if !valid {
				w.WriteHeader(http.StatusUnauthorized)
			} else {
				rt := new(mng.RouteURL)
				decoder := json.NewDecoder(r.Body)
				error := decoder.Decode(&rt)
				if error != nil {
					log.Println(error.Error())
					http.Error(w, error.Error(), http.StatusBadRequest)
				} else if rt.ClientID == 0 || rt.RouteID == 0 || rt.Name == "" || rt.URL == "" {
					http.Error(w, "bad request", http.StatusBadRequest)
				} else {
					rt.Active = false
					resOut := gatewayDB.InsertRouteURL(rt)
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

//HandleRouteURLSuperPut HandleRouteURLSuperPut
func (h Handler) HandleRouteURLSuperPut(w http.ResponseWriter, r *http.Request) {
	var gatewayDB mng.GatewayDB
	gatewayDB.DbConfig = h.DbConfig
	gatewayDB.GwCacheHost = getCacheHost()

	auth := getAuth(r)
	usume := new(uoauth.Claim)
	usume.Role = "superAdmin"
	usume.Scope = "write"
	w.Header().Set("Content-Type", "application/json")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		switch r.Method {
		case "PUT":
			usume.URI = "/ulbora/rs/gwRouteUrlSuper/update"
			var valid bool
			if testMode {
				valid = true
			} else {
				valid = auth.Authorize(usume)
			}
			if !valid {
				w.WriteHeader(http.StatusUnauthorized)
			} else {
				rt := new(mng.RouteURL)
				decoder := json.NewDecoder(r.Body)
				error := decoder.Decode(&rt)
				if error != nil {
					log.Println(error.Error())
					http.Error(w, error.Error(), http.StatusBadRequest)
				} else if rt.ID == 0 || rt.ClientID == 0 || rt.RouteID == 0 || rt.Name == "" || rt.URL == "" {
					http.Error(w, "bad request in update", http.StatusBadRequest)
				} else {
					resOut := gatewayDB.UpdateRouteURL(rt)
					//fmt.Print("response: ")
					//fmt.Println(resOut)
					gatewayDB.Cb.Reset(rt.ClientID, rt.ID)
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

//HandleRouteURLActivateSuper HandleRouteURLActivateSuper
func (h Handler) HandleRouteURLActivateSuper(w http.ResponseWriter, r *http.Request) {
	var gatewayDB mng.GatewayDB
	gatewayDB.DbConfig = h.DbConfig
	gatewayDB.GwCacheHost = getCacheHost()

	auth := getAuth(r)
	usame := new(uoauth.Claim)
	usame.Role = "superAdmin"
	usame.Scope = "write"
	w.Header().Set("Content-Type", "application/json")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		switch r.Method {
		case "PUT":
			usame.URI = "/ulbora/rs/gwRouteUrlSuper/activate"
			var valid bool
			if testMode {
				valid = true
			} else {
				valid = auth.Authorize(usame)
			}
			if !valid {
				w.WriteHeader(http.StatusUnauthorized)
			} else {
				rt := new(mng.RouteURL)
				decoder := json.NewDecoder(r.Body)
				error := decoder.Decode(&rt)
				if error != nil {
					log.Println(error.Error())
					http.Error(w, error.Error(), http.StatusBadRequest)
				} else if rt.ID == 0 || rt.ClientID == 0 || rt.RouteID == 0 {
					http.Error(w, "bad request in update", http.StatusBadRequest)
				} else {
					resOut := gatewayDB.ActivateRouteURL(rt)
					gatewayDB.Cb.Reset(rt.ClientID, rt.ID)
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

//HandleRouteURLSuperGet HandleRouteURLSuperGet
func (h Handler) HandleRouteURLSuperGet(w http.ResponseWriter, r *http.Request) {
	var gatewayDB mng.GatewayDB
	gatewayDB.DbConfig = h.DbConfig
	gatewayDB.GwCacheHost = getCacheHost()

	auth := getAuth(r)
	usgme := new(uoauth.Claim)
	usgme.Role = "superAdmin"

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	var id int64
	var clientID int64
	var routeIDrusg int64
	var errID error
	var errCID error
	var errRID error

	if vars != nil {
		id, errID = strconv.ParseInt(vars["id"], 10, 0)
		routeIDrusg, errRID = strconv.ParseInt(vars["routeId"], 10, 0)
		clientID, errCID = strconv.ParseInt(vars["clientId"], 10, 0)
	} else {
		var idStr = r.URL.Query().Get("id")
		id, errID = strconv.ParseInt(idStr, 10, 0)

		var routeIDStr = r.URL.Query().Get("routeId")
		routeIDrusg, errRID = strconv.ParseInt(routeIDStr, 10, 0)

		var clientIDStr = r.URL.Query().Get("clientId")
		clientID, errCID = strconv.ParseInt(clientIDStr, 10, 0)
	}
	if errID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	if errRID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	if errCID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}

	//fmt.Print("id is: ")
	//fmt.Println(id)
	switch r.Method {
	case "GET":
		usgme.URI = "/ulbora/rs/gwRouteUrlSuper/get"
		usgme.Scope = "read"
		var valid bool
		if testMode {
			valid = true
		} else {
			valid = auth.Authorize(usgme)
		}
		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			rt := new(mng.RouteURL)
			rt.ID = id
			rt.RouteID = routeIDrusg
			rt.ClientID = clientID
			resOut := gatewayDB.GetRouteURL(rt)
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

//HandleRouteURLSuperDelete HandleRouteURLSuperDelete
func (h Handler) HandleRouteURLSuperDelete(w http.ResponseWriter, r *http.Request) {
	var gatewayDB mng.GatewayDB
	gatewayDB.DbConfig = h.DbConfig
	gatewayDB.GwCacheHost = getCacheHost()

	auth := getAuth(r)
	usdme := new(uoauth.Claim)
	usdme.Role = "superAdmin"

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	var id int64
	var clientID int64
	var routeIDrusd int64
	var errID error
	var errCID error
	var errRID error

	if vars != nil {
		id, errID = strconv.ParseInt(vars["id"], 10, 0)
		routeIDrusd, errRID = strconv.ParseInt(vars["routeId"], 10, 0)
		clientID, errCID = strconv.ParseInt(vars["clientId"], 10, 0)
	} else {
		var idStr = r.URL.Query().Get("id")
		id, errID = strconv.ParseInt(idStr, 10, 0)

		var routeIDStr = r.URL.Query().Get("routeId")
		routeIDrusd, errRID = strconv.ParseInt(routeIDStr, 10, 0)

		var clientIDStr = r.URL.Query().Get("clientId")
		clientID, errCID = strconv.ParseInt(clientIDStr, 10, 0)
	}
	if errID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	if errRID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	if errCID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}

	//fmt.Print("id is: ")
	//fmt.Println(id)
	switch r.Method {
	case "DELETE":
		usdme.URI = "/ulbora/rs/gwRouteUrlSuper/delete"
		usdme.Scope = "write"
		var valid bool
		if testMode {
			valid = true
		} else {
			valid = auth.Authorize(usdme)
		}
		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			rt := new(mng.RouteURL)
			rt.ID = id
			rt.RouteID = routeIDrusd
			rt.ClientID = clientID
			resOut := gatewayDB.DeleteRouteURL(rt)
			gatewayDB.Cb.Reset(rt.ClientID, rt.ID)
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

//HandleRouteURLSuperList HandleRouteURLSuperList
func (h Handler) HandleRouteURLSuperList(w http.ResponseWriter, r *http.Request) {
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
	var routeID int64
	//var errID error
	var errCID error
	var errRID error

	if vars != nil {
		//id, errID = strconv.ParseInt(vars["id"], 10, 0)
		routeID, errRID = strconv.ParseInt(vars["routeId"], 10, 0)
		clientID, errCID = strconv.ParseInt(vars["clientId"], 10, 0)
	} else {
		//var idStr = r.URL.Query().Get("id")
		//id, errID = strconv.ParseInt(idStr, 10, 0)

		var routeIDStr = r.URL.Query().Get("routeId")
		routeID, errRID = strconv.ParseInt(routeIDStr, 10, 0)

		var clientIDStr = r.URL.Query().Get("clientId")
		clientID, errCID = strconv.ParseInt(clientIDStr, 10, 0)
	}
	// if errID != nil {
	// 	http.Error(w, "bad request", http.StatusBadRequest)
	// }
	if errRID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	if errCID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}

	switch r.Method {
	case "GET":
		me.URI = "/ulbora/rs/gwRouteUrlSuper/list"
		var valid bool
		if testMode {
			valid = true
		} else {
			valid = auth.Authorize(me)
		}
		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			rt := new(mng.RouteURL)
			rt.ClientID = clientID
			rt.RouteID = routeID
			resOut := gatewayDB.GetRouteURLList(rt)
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
