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

func handleBreakerSuperChange(w http.ResponseWriter, r *http.Request) {
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
			me.URI = "/ulbora/rs/gwBreakerSuper/add"
			valid := auth.Authorize(me)
			if valid != true {
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
		case "PUT":
			me.URI = "/ulbora/rs/gwBreakerSuper/update"
			valid := auth.Authorize(me)
			if valid != true {
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
		}
	}
}

func handleBreakerSuperReset(w http.ResponseWriter, r *http.Request) {
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
			me.URI = "/ulbora/rs/gwBreakerSuper/reset"
			valid := auth.Authorize(me)
			if valid != true {
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
		}
	}
}

func handleBreakerSuper(w http.ResponseWriter, r *http.Request) {
	auth := getAuth(r)
	me := new(uoauth.Claim)
	me.Role = "superAdmin"

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	clientID, errCID := strconv.ParseInt(vars["clientId"], 10, 0)
	if errCID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	routeID, errRID := strconv.ParseInt(vars["routeId"], 10, 0)
	if errRID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	UID, errUID := strconv.ParseInt(vars["urlId"], 10, 0)
	if errUID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	//fmt.Print("id is: ")
	//fmt.Println(id)
	switch r.Method {
	case "GET":
		me.URI = "/ulbora/rs/gwBreakerSuper/get"
		me.Scope = "read"
		valid := auth.Authorize(me)
		if valid != true {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			bk := new(cb.Breaker)
			bk.ClientID = clientID
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

	case "DELETE":
		me.URI = "/ulbora/rs/gwBreakerSuper/delete"
		me.Scope = "write"
		valid := auth.Authorize(me)
		if valid != true {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			bk := new(cb.Breaker)
			bk.ClientID = clientID
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
	}
}

func handleBreakerStatusSuper(w http.ResponseWriter, r *http.Request) {
	auth := getAuth(r)
	me := new(uoauth.Claim)
	me.Role = "superAdmin"

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	clientID, errCID := strconv.ParseInt(vars["clientId"], 10, 0)
	if errCID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	// routeID, errRID := strconv.ParseInt(vars["routeId"], 10, 0)
	// if errRID != nil {
	// 	http.Error(w, "bad request", http.StatusBadRequest)
	// }
	UID, errUID := strconv.ParseInt(vars["urlId"], 10, 0)
	if errUID != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	//fmt.Print("id is: ")
	//fmt.Println(id)
	switch r.Method {
	case "GET":
		me.URI = "/ulbora/rs/gwBreakerSuper/status"
		me.Scope = "read"
		valid := auth.Authorize(me)
		if valid != true {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			// bk := new(cb.Breaker)
			// bk.ClientID = clientID
			// bk.RestRouteID = routeID
			// bk.RouteURIID = UID
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
	}
}

// func handleRestRouteSuperList(w http.ResponseWriter, r *http.Request) {
// 	auth := getAuth(r)
// 	me := new(uoauth.Claim)
// 	me.Role = "superAdmin"
// 	me.Scope = "read"
// 	w.Header().Set("Content-Type", "application/json")
// 	vars := mux.Vars(r)
// 	clientID, errCID := strconv.ParseInt(vars["clientId"], 10, 0)
// 	if errCID != nil {
// 		http.Error(w, "bad request", http.StatusBadRequest)
// 	}
// 	switch r.Method {
// 	case "GET":
// 		me.URI = "/rs/gwRestRouteSuper/list"
// 		valid := auth.Authorize(me)
// 		if valid != true {
// 			w.WriteHeader(http.StatusUnauthorized)
// 		} else {
// 			rt := new(mng.RestRoute)
// 			rt.ClientID = clientID
// 			resOut := gatewayDB.GetRestRouteList(rt)
// 			//fmt.Print("response: ")
// 			//fmt.Println(resOut)
// 			resJSON, err := json.Marshal(resOut)
// 			//fmt.Print("response json: ")
// 			//fmt.Println(string(resJSON))
// 			if err != nil {
// 				log.Println(err.Error())
// 				http.Error(w, "json output failed", http.StatusInternalServerError)
// 			}
// 			w.WriteHeader(http.StatusOK)
// 			if string(resJSON) == "null" {
// 				fmt.Fprint(w, "[]")
// 			} else {
// 				fmt.Fprint(w, string(resJSON))
// 			}
// 		}
// 	}
// }
