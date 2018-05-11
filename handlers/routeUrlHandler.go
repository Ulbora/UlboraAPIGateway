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

//HandleRouteURLPost HandleRouteURLPost
func (h Handler) HandleRouteURLPost(w http.ResponseWriter, r *http.Request) {
	var gatewayDB mng.GatewayDB
	gatewayDB.DbConfig = h.DbConfig
	gatewayDB.GwCacheHost = getCacheHost()

	auth := getAuth(r)
	rupme := new(uoauth.Claim)
	rupme.Role = "admin"
	rupme.Scope = "write"
	w.Header().Set("Content-Type", "application/json")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		switch r.Method {
		case "POST":
			rupme.URI = "/ulbora/rs/gwRouteUrl/add"
			var valid bool
			if testMode {
				valid = true
			} else {
				valid = auth.Authorize(rupme)
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
				} else if rt.RouteID == 0 || rt.Name == "" || rt.URL == "" {
					http.Error(w, "bad request", http.StatusBadRequest)
				} else {
					rt.ClientID = auth.ClientID
					rt.Active = false
					resOut := gatewayDB.InsertRouteURL(rt)
					//fmt.Print("response: ")
					//fmt.Println(resOut)
					resJSON, err := json.Marshal(resOut)
					if err != nil {
						log.Println(error.Error())
						//.Error(w, "json output failed", http.StatusInternalServerError)
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

//HandleRouteURLPut HandleRouteURLPut
func (h Handler) HandleRouteURLPut(w http.ResponseWriter, r *http.Request) {
	var gatewayDB mng.GatewayDB
	gatewayDB.DbConfig = h.DbConfig
	gatewayDB.GwCacheHost = getCacheHost()

	auth := getAuth(r)
	ruume := new(uoauth.Claim)
	ruume.Role = "admin"
	ruume.Scope = "write"
	w.Header().Set("Content-Type", "application/json")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		switch r.Method {
		case "PUT":
			ruume.URI = "/ulbora/rs/gwRouteUrl/update"
			var valid bool
			if testMode {
				valid = true
			} else {
				valid = auth.Authorize(ruume)
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
				} else if rt.ID == 0 || rt.RouteID == 0 || rt.Name == "" || rt.URL == "" {
					http.Error(w, "bad request in update", http.StatusBadRequest)
				} else {
					rt.ClientID = auth.ClientID
					resOut := gatewayDB.UpdateRouteURL(rt)
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

//HandleRouteURLActivate HandleRouteURLActivate
func (h Handler) HandleRouteURLActivate(w http.ResponseWriter, r *http.Request) {
	var gatewayDB mng.GatewayDB
	gatewayDB.DbConfig = h.DbConfig
	gatewayDB.GwCacheHost = getCacheHost()

	auth := getAuth(r)
	ruame := new(uoauth.Claim)
	ruame.Role = "admin"
	ruame.Scope = "write"
	w.Header().Set("Content-Type", "application/json")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		switch r.Method {
		case "PUT":
			ruame.URI = "/ulbora/rs/gwRouteUrl/activate"
			var valid bool
			if testMode {
				valid = true
			} else {
				valid = auth.Authorize(ruame)
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
				} else if rt.ID == 0 || rt.RouteID == 0 {
					http.Error(w, "bad request in update", http.StatusBadRequest)
				} else {
					rt.ClientID = auth.ClientID
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

//HandleRouteURLGet HandleRouteURLGet
func (h Handler) HandleRouteURLGet(w http.ResponseWriter, r *http.Request) {
	var gatewayDB mng.GatewayDB
	gatewayDB.DbConfig = h.DbConfig
	gatewayDB.GwCacheHost = getCacheHost()

	auth := getAuth(r)
	rugme := new(uoauth.Claim)
	rugme.Role = "admin"

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	var id int64
	//var clientID int64
	var routeID int64
	var errID error
	//var errCID error
	var errRID error

	if vars != nil {
		id, errID = strconv.ParseInt(vars["id"], 10, 0)
		routeID, errRID = strconv.ParseInt(vars["routeId"], 10, 0)
		//clientID, errCID = strconv.ParseInt(vars["clientId"], 10, 0)
	} else {
		var idStr = r.URL.Query().Get("id")
		id, errID = strconv.ParseInt(idStr, 10, 0)

		var routeIDStr = r.URL.Query().Get("routeId")
		routeID, errRID = strconv.ParseInt(routeIDStr, 10, 0)

		//var clientIDStr = r.URL.Query().Get("clientId")
		//clientID, errCID = strconv.ParseInt(clientIDStr, 10, 0)
	}
	if errID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	if errRID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	// if errCID != nil {
	// 	http.Error(w, "bad request", http.StatusBadRequest)
	// }

	// id, errID := strconv.ParseInt(vars["id"], 10, 0)
	// if errID != nil {
	// 	http.Error(w, "bad request", http.StatusBadRequest)
	// }
	// routeID, errRID := strconv.ParseInt(vars["routeId"], 10, 0)
	// if errRID != nil {
	// 	http.Error(w, "bad request", http.StatusBadRequest)
	// }

	//fmt.Print("id is: ")
	//fmt.Println(id)
	switch r.Method {
	case "GET":
		rugme.URI = "/ulbora/rs/gwRouteUrl/get"
		rugme.Scope = "read"
		var valid bool
		if testMode {
			valid = true
		} else {
			valid = auth.Authorize(rugme)
		}
		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			rt := new(mng.RouteURL)
			rt.ID = id
			rt.RouteID = routeID
			rt.ClientID = auth.ClientID
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

//HandleRouteURLDelete HandleRouteURLDelete
func (h Handler) HandleRouteURLDelete(w http.ResponseWriter, r *http.Request) {
	var gatewayDB mng.GatewayDB
	gatewayDB.DbConfig = h.DbConfig
	gatewayDB.GwCacheHost = getCacheHost()

	auth := getAuth(r)
	rudme := new(uoauth.Claim)
	rudme.Role = "admin"

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	var id int64
	//var clientID int64
	var routeID int64
	var errID error
	//var errCID error
	var errRID error

	if vars != nil {
		id, errID = strconv.ParseInt(vars["id"], 10, 0)
		routeID, errRID = strconv.ParseInt(vars["routeId"], 10, 0)
		//clientID, errCID = strconv.ParseInt(vars["clientId"], 10, 0)
	} else {
		var idStr = r.URL.Query().Get("id")
		id, errID = strconv.ParseInt(idStr, 10, 0)

		var routeIDStr = r.URL.Query().Get("routeId")
		routeID, errRID = strconv.ParseInt(routeIDStr, 10, 0)

		//var clientIDStr = r.URL.Query().Get("clientId")
		//clientID, errCID = strconv.ParseInt(clientIDStr, 10, 0)
	}
	if errID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	if errRID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}

	// id, errID := strconv.ParseInt(vars["id"], 10, 0)
	// if errID != nil {
	// 	http.Error(w, "bad request", http.StatusBadRequest)
	// }
	// routeID, errRID := strconv.ParseInt(vars["routeId"], 10, 0)
	// if errRID != nil {
	// 	http.Error(w, "bad request", http.StatusBadRequest)
	// }

	//fmt.Print("id is: ")
	//fmt.Println(id)
	switch r.Method {
	case "DELETE":
		rudme.URI = "/ulbora/rs/gwRouteUrl/delete"
		rudme.Scope = "write"
		var valid bool
		if testMode {
			valid = true
		} else {
			valid = auth.Authorize(rudme)
		}
		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			rt := new(mng.RouteURL)
			rt.ID = id
			rt.RouteID = routeID
			rt.ClientID = auth.ClientID
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

//HandleRouteURLList HandleRouteURLList
func (h Handler) HandleRouteURLList(w http.ResponseWriter, r *http.Request) {
	var gatewayDB mng.GatewayDB
	gatewayDB.DbConfig = h.DbConfig
	gatewayDB.GwCacheHost = getCacheHost()

	auth := getAuth(r)
	rulme := new(uoauth.Claim)
	rulme.Role = "admin"
	rulme.Scope = "read"
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	//var id int64
	//var clientID int64
	var routeID int64
	//var errID error
	//var errCID error
	var errRID error

	if vars != nil {
		//id, errID = strconv.ParseInt(vars["id"], 10, 0)
		routeID, errRID = strconv.ParseInt(vars["routeId"], 10, 0)
		//clientID, errCID = strconv.ParseInt(vars["clientId"], 10, 0)
	} else {
		//	var idStr = r.URL.Query().Get("id")
		//id, errID = strconv.ParseInt(idStr, 10, 0)

		var routeIDStr = r.URL.Query().Get("routeId")
		routeID, errRID = strconv.ParseInt(routeIDStr, 10, 0)

		//var clientIDStr = r.URL.Query().Get("clientId")
		//clientID, errCID = strconv.ParseInt(clientIDStr, 10, 0)
	}
	// if errID != nil {
	// 	http.Error(w, "bad request", http.StatusBadRequest)
	// }
	if errRID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}

	// routeID, errRID := strconv.ParseInt(vars["routeId"], 10, 0)
	// if errRID != nil {
	// 	http.Error(w, "bad request", http.StatusBadRequest)
	// }
	switch r.Method {
	case "GET":
		rulme.URI = "/ulbora/rs/gwRouteUrl/list"
		var valid bool
		if testMode {
			valid = true
		} else {
			valid = auth.Authorize(rulme)
		}
		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			rt := new(mng.RouteURL)
			rt.ClientID = auth.ClientID
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
