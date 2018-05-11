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

//HandleBreakerPost HandleBreakerPost
func (h Handler) HandleBreakerPost(w http.ResponseWriter, r *http.Request) {
	var cbDB cb.CircuitBreaker
	cbDB.DbConfig = h.DbConfig
	cbDB.CacheHost = getCacheHost()
	auth := getAuth(r)
	bpme := new(uoauth.Claim)
	bpme.Role = "admin"
	bpme.Scope = "write"
	w.Header().Set("Content-Type", "application/json")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		switch r.Method {
		case "POST":
			bpme.URI = "/ulbora/rs/gwBreaker/add"
			var valid bool
			if testMode {
				valid = true
			} else {
				valid = auth.Authorize(bpme)
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
				} else if bk.RestRouteID == 0 || bk.RouteURIID == 0 {
					http.Error(w, "bad request", http.StatusBadRequest)
				} else {
					bk.ClientID = auth.ClientID
					bpsuc, err := cbDB.InsertBreaker(bk)
					//fmt.Print("response: ")
					//fmt.Println(resOut)
					var bpres BreakerResponse
					bpres.Success = bpsuc
					if err != nil {
						bpres.Error = err.Error()
						log.Println(err.Error())
						//http.Error(w, "json output failed", http.StatusInternalServerError)
					}
					resJSON, cerr := json.Marshal(bpres)
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

//HandleBreakerPut HandleBreakerPut
func (h Handler) HandleBreakerPut(w http.ResponseWriter, r *http.Request) {
	var cbDB cb.CircuitBreaker
	cbDB.DbConfig = h.DbConfig
	cbDB.CacheHost = getCacheHost()
	auth := getAuth(r)
	bpume := new(uoauth.Claim)
	bpume.Role = "admin"
	bpume.Scope = "write"
	w.Header().Set("Content-Type", "application/json")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		switch r.Method {
		case "PUT":
			bpume.URI = "/ulbora/rs/gwBreaker/update"
			var valid bool
			if testMode {
				valid = true
			} else {
				valid = auth.Authorize(bpume)
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
				} else if bk.ID == 0 || bk.RestRouteID == 0 || bk.RouteURIID == 0 {
					http.Error(w, "bad request in update", http.StatusBadRequest)
				} else {
					bk.ClientID = auth.ClientID
					suc, err := cbDB.UpdateBreaker(bk)
					//fmt.Print("response: ")
					//fmt.Println(resOut)
					var bures BreakerResponse
					bures.Success = suc
					if err != nil {
						bures.Error = err.Error()
						log.Println(error.Error())
						http.Error(w, "json output failed", http.StatusInternalServerError)
					}
					resJSON, cerr := json.Marshal(bures)
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

//HandleBreakerReset HandleBreakerReset
func (h Handler) HandleBreakerReset(w http.ResponseWriter, r *http.Request) {
	var cbDB cb.CircuitBreaker
	cbDB.DbConfig = h.DbConfig
	cbDB.CacheHost = getCacheHost()
	auth := getAuth(r)
	brme := new(uoauth.Claim)
	brme.Role = "admin"
	brme.Scope = "write"
	w.Header().Set("Content-Type", "application/json")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		switch r.Method {
		case "POST":
			brme.URI = "/ulbora/rs/gwBreaker/reset"
			var valid bool
			if testMode {
				valid = true
			} else {
				valid = auth.Authorize(brme)
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
				} else if bk.RouteURIID == 0 {
					http.Error(w, "bad request", http.StatusBadRequest)
				} else {
					bk.ClientID = auth.ClientID
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

//HandleBreakerGet HandleBreakerGet
func (h Handler) HandleBreakerGet(w http.ResponseWriter, r *http.Request) {
	var cbDB cb.CircuitBreaker
	cbDB.DbConfig = h.DbConfig
	cbDB.CacheHost = getCacheHost()
	auth := getAuth(r)
	bgme := new(uoauth.Claim)
	bgme.Role = "admin"

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	//var clientID int64
	//var errCID error

	var routeID int64
	var errRID error

	var UID int64
	var errUID error

	if vars != nil {
		routeID, errRID = strconv.ParseInt(vars["routeId"], 10, 0)
		UID, errUID = strconv.ParseInt(vars["urlId"], 10, 0)
	} else {

		var routeIDStr = r.URL.Query().Get("routeId")
		routeID, errRID = strconv.ParseInt(routeIDStr, 10, 0)

		var urlIDStr = r.URL.Query().Get("urlId")
		UID, errUID = strconv.ParseInt(urlIDStr, 10, 0)
	}

	if errRID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	if errUID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}

	switch r.Method {
	case "GET":
		bgme.URI = "/ulbora/rs/gwBreaker/get"
		bgme.Scope = "read"
		var valid bool
		if testMode {
			valid = true
		} else {
			valid = auth.Authorize(bgme)
		}
		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			bk := new(cb.Breaker)
			bk.ClientID = auth.ClientID
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

//HandleBreakerDelete HandleBreakerDelete
func (h Handler) HandleBreakerDelete(w http.ResponseWriter, r *http.Request) {
	var cbDB cb.CircuitBreaker
	cbDB.DbConfig = h.DbConfig
	cbDB.CacheHost = getCacheHost()
	auth := getAuth(r)
	bdme := new(uoauth.Claim)
	bdme.Role = "admin"

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	var routeID int64
	var errRID error

	var UID int64
	var errUID error

	if vars != nil {
		routeID, errRID = strconv.ParseInt(vars["routeId"], 10, 0)
		UID, errUID = strconv.ParseInt(vars["urlId"], 10, 0)
	} else {

		var routeIDStr = r.URL.Query().Get("routeId")
		routeID, errRID = strconv.ParseInt(routeIDStr, 10, 0)

		var urlIDStr = r.URL.Query().Get("urlId")
		UID, errUID = strconv.ParseInt(urlIDStr, 10, 0)
	}

	if errRID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	if errUID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}

	//fmt.Print("id is: ")
	//fmt.Println(id)
	switch r.Method {
	case "DELETE":
		bdme.URI = "/ulbora/rs/gwBreaker/delete"
		bdme.Scope = "write"
		var valid bool
		if testMode {
			valid = true
		} else {
			valid = auth.Authorize(bdme)
		}
		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			bk := new(cb.Breaker)
			bk.ClientID = auth.ClientID
			bk.RestRouteID = routeID
			bk.RouteURIID = UID
			suc := cbDB.DeleteBreaker(bk)
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

//HandleBreakerStatus HandleBreakerStatus
func (h Handler) HandleBreakerStatus(w http.ResponseWriter, r *http.Request) {
	var cbDB cb.CircuitBreaker
	cbDB.DbConfig = h.DbConfig
	cbDB.CacheHost = getCacheHost()
	auth := getAuth(r)
	me := new(uoauth.Claim)
	me.Role = "admin"

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	var UID int64
	var errUID error

	if vars != nil {
		UID, errUID = strconv.ParseInt(vars["urlId"], 10, 0)
	} else {

		var urlIDStr = r.URL.Query().Get("urlId")
		UID, errUID = strconv.ParseInt(urlIDStr, 10, 0)
	}

	if errUID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}

	//fmt.Print("id is: ")
	//fmt.Println(id)
	switch r.Method {
	case "GET":
		me.URI = "/ulbora/rs/gwBreaker/status"
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

			resOut := cbDB.GetStatus(auth.ClientID, UID)
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
