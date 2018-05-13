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
	cb "UlboraApiGateway/circuitbreaker"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	uoauth "github.com/Ulbora/go-ulbora-oauth2"
	"github.com/gorilla/mux"
)

// BreakerResponse BreakerResponse
type BreakerResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

//HandleBreakerSuperPost HandleBreakerSuperPost
func (h Handler) HandleBreakerSuperPost(w http.ResponseWriter, r *http.Request) {
	var cbDB cb.CircuitBreaker
	cbDB.DbConfig = h.DbConfig
	cbDB.CacheHost = getCacheHost()
	auth := getAuth(r)
	bspme := new(uoauth.Claim)
	bspme.Role = "superAdmin"
	bspme.Scope = "write"
	w.Header().Set("Content-Type", "application/json")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		switch r.Method {
		case "POST":
			bspme.URI = "/ulbora/rs/gwBreakerSuper/add"
			var valid bool
			if testMode {
				valid = true
			} else {
				valid = auth.Authorize(bspme)
			}
			if !valid {
				w.WriteHeader(http.StatusUnauthorized)
			} else {
				bk := new(cb.Breaker)
				decoder := json.NewDecoder(r.Body)
				error := decoder.Decode(&bk)
				if error != nil {
					log.Println(error.Error())
					http.Error(w, error.Error(), http.StatusBadRequest)
				} else if bk.ClientID == 0 || bk.RestRouteID == 0 || bk.RouteURIID == 0 {
					http.Error(w, "bad request", http.StatusBadRequest)
				} else {
					suc, err := cbDB.InsertBreaker(bk)
					//fmt.Print("response: ")
					//fmt.Println(resOut)
					var res BreakerResponse
					res.Success = suc
					if err != nil {
						res.Error = err.Error()
						log.Println(err.Error())
						//http.Error(w, "json output failed", http.StatusInternalServerError)
					}
					resJSON, cerr := json.Marshal(res)
					if cerr != nil {
						log.Println(cerr.Error())
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

//HandleBreakerSuperPut HandleBreakerSuperPut
func (h Handler) HandleBreakerSuperPut(w http.ResponseWriter, r *http.Request) {
	var cbDB cb.CircuitBreaker
	cbDB.DbConfig = h.DbConfig
	cbDB.CacheHost = getCacheHost()
	auth := getAuth(r)
	bsume := new(uoauth.Claim)
	bsume.Role = "superAdmin"
	bsume.Scope = "write"
	w.Header().Set("Content-Type", "application/json")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		switch r.Method {
		case "PUT":
			bsume.URI = "/ulbora/rs/gwBreakerSuper/update"
			var valid bool
			if testMode {
				valid = true
			} else {
				valid = auth.Authorize(bsume)
			}
			if !valid {
				w.WriteHeader(http.StatusUnauthorized)
			} else {
				bk := new(cb.Breaker)
				decoder := json.NewDecoder(r.Body)
				error := decoder.Decode(&bk)
				if error != nil {
					log.Println(error.Error())
					http.Error(w, error.Error(), http.StatusBadRequest)
				} else if bk.ID == 0 || bk.ClientID == 0 || bk.RestRouteID == 0 || bk.RouteURIID == 0 {
					http.Error(w, "bad request in update", http.StatusBadRequest)
				} else {
					suc, err := cbDB.UpdateBreaker(bk)
					//fmt.Print("response: ")
					//fmt.Println(resOut)
					var res BreakerResponse
					res.Success = suc
					if err != nil {
						res.Error = err.Error()
						log.Println(error.Error())
						//http.Error(w, "json output failed", http.StatusInternalServerError)
					}
					resJSON, cerr := json.Marshal(res)
					if cerr != nil {
						log.Println(cerr.Error())
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

//HandleBreakerSuperReset HandleBreakerSuperReset
func (h Handler) HandleBreakerSuperReset(w http.ResponseWriter, r *http.Request) {
	var cbDB cb.CircuitBreaker
	cbDB.DbConfig = h.DbConfig
	cbDB.CacheHost = getCacheHost()
	auth := getAuth(r)
	bsrme := new(uoauth.Claim)
	bsrme.Role = "superAdmin"
	bsrme.Scope = "write"
	w.Header().Set("Content-Type", "application/json")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		switch r.Method {
		case "POST":
			bsrme.URI = "/ulbora/rs/gwBreakerSuper/reset"
			var valid bool
			if testMode {
				valid = true
			} else {
				valid = auth.Authorize(bsrme)
			}
			if !valid {
				w.WriteHeader(http.StatusUnauthorized)
			} else {
				bk := new(cb.Breaker)
				decoder := json.NewDecoder(r.Body)
				error := decoder.Decode(&bk)
				if error != nil {
					log.Println(error.Error())
					http.Error(w, error.Error(), http.StatusBadRequest)
				} else if bk.ClientID == 0 || bk.RouteURIID == 0 {
					http.Error(w, "bad request", http.StatusBadRequest)
				} else {
					cbDB.Reset(bk.ClientID, bk.RouteURIID)
					//fmt.Print("response: ")
					//fmt.Println(resOut)
					var res BreakerResponse
					res.Success = true
					resJSON, cerr := json.Marshal(res)
					if cerr != nil {
						log.Println(cerr.Error())
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

//HandleBreakerSuperGet HandleBreakerSuperGet
func (h Handler) HandleBreakerSuperGet(w http.ResponseWriter, r *http.Request) {
	var cbDB cb.CircuitBreaker
	cbDB.DbConfig = h.DbConfig
	cbDB.CacheHost = getCacheHost()
	auth := getAuth(r)
	bsgme := new(uoauth.Claim)
	bsgme.Role = "superAdmin"

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	var bsgClientID int64
	var errCID error

	var routeID int64
	var errRIDbsg error

	var UID int64
	var errUID error

	if vars != nil {
		bsgClientID, errCID = strconv.ParseInt(vars["clientId"], 10, 0)
		routeID, errRIDbsg = strconv.ParseInt(vars["routeId"], 10, 0)
		UID, errUID = strconv.ParseInt(vars["urlId"], 10, 0)
	} else {
		var clientIDStr = r.URL.Query().Get("clientId")
		bsgClientID, errCID = strconv.ParseInt(clientIDStr, 10, 0)

		var routeIDStr = r.URL.Query().Get("routeId")
		routeID, errRIDbsg = strconv.ParseInt(routeIDStr, 10, 0)

		var urlIDStr = r.URL.Query().Get("urlId")
		UID, errUID = strconv.ParseInt(urlIDStr, 10, 0)
	}
	if errCID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	if errRIDbsg != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	if errUID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}

	//fmt.Print("id is: ")
	//fmt.Println(id)
	switch r.Method {
	case "GET":
		bsgme.URI = "/ulbora/rs/gwBreakerSuper/get"
		bsgme.Scope = "read"
		var valid bool
		if testMode {
			valid = true
		} else {
			valid = auth.Authorize(bsgme)
		}
		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			bk := new(cb.Breaker)
			bk.ClientID = bsgClientID
			bk.RestRouteID = routeID
			bk.RouteURIID = UID
			resOut := cbDB.GetBreaker(bk)
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

//HandleBreakerSuperDelete HandleBreakerSuperDelete
func (h Handler) HandleBreakerSuperDelete(w http.ResponseWriter, r *http.Request) {
	var cbDB cb.CircuitBreaker
	cbDB.DbConfig = h.DbConfig
	cbDB.CacheHost = getCacheHost()
	auth := getAuth(r)
	bsdme := new(uoauth.Claim)
	bsdme.Role = "superAdmin"

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	var bsdClientID int64
	var errCID error

	var routeID int64
	var errRID error

	var UID int64
	var errUIDbsd error

	if vars != nil {
		bsdClientID, errCID = strconv.ParseInt(vars["clientId"], 10, 0)
		routeID, errRID = strconv.ParseInt(vars["routeId"], 10, 0)
		UID, errUIDbsd = strconv.ParseInt(vars["urlId"], 10, 0)
	} else {
		var clientIDStr = r.URL.Query().Get("clientId")
		bsdClientID, errCID = strconv.ParseInt(clientIDStr, 10, 0)

		var routeIDStr = r.URL.Query().Get("routeId")
		routeID, errRID = strconv.ParseInt(routeIDStr, 10, 0)

		var urlIDStr = r.URL.Query().Get("urlId")
		UID, errUIDbsd = strconv.ParseInt(urlIDStr, 10, 0)
	}
	if errCID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	if errRID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	if errUIDbsd != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}

	//fmt.Print("id is: ")
	//fmt.Println(id)
	switch r.Method {
	case "DELETE":
		bsdme.URI = "/ulbora/rs/gwBreakerSuper/delete"
		bsdme.Scope = "write"
		var valid bool
		if testMode {
			valid = true
		} else {
			valid = auth.Authorize(bsdme)
		}
		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			bk := new(cb.Breaker)
			bk.ClientID = bsdClientID
			bk.RestRouteID = routeID
			bk.RouteURIID = UID
			suc := cbDB.DeleteBreaker(bk)
			//fmt.Print("response: ")
			//fmt.Println(resOut)
			var res BreakerResponse
			res.Success = suc
			//fmt.Print("response: ")
			//fmt.Println(resOut)
			resJSON, err := json.Marshal(res)
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

//HandleBreakerStatusSuper HandleBreakerStatusSuper
func (h Handler) HandleBreakerStatusSuper(w http.ResponseWriter, r *http.Request) {
	var cbDB cb.CircuitBreaker
	cbDB.DbConfig = h.DbConfig
	cbDB.CacheHost = getCacheHost()
	auth := getAuth(r)
	me := new(uoauth.Claim)
	me.Role = "superAdmin"

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	var clientID int64
	var errCID error

	var UID int64
	var errUID error

	if vars != nil {
		clientID, errCID = strconv.ParseInt(vars["clientId"], 10, 0)
		UID, errUID = strconv.ParseInt(vars["urlId"], 10, 0)
	} else {
		var clientIDStr = r.URL.Query().Get("clientId")
		clientID, errCID = strconv.ParseInt(clientIDStr, 10, 0)

		var urlIDStr = r.URL.Query().Get("urlId")
		UID, errUID = strconv.ParseInt(urlIDStr, 10, 0)
	}
	if errCID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}

	if errUID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}

	//fmt.Print("clientId is: ")
	//fmt.Println(clientID)
	//fmt.Print("id is: ")
	//fmt.Println(UID)
	switch r.Method {
	case "GET":
		me.URI = "/ulbora/rs/gwBreakerSuper/status"
		me.Scope = "read"
		var valid bool
		if testMode {
			valid = true
		} else {
			valid = auth.Authorize(me)
		}
		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			resOut := cbDB.GetStatus(clientID, UID)
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
