package handlers

/*
 Copyright (C) 2017 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2017 Ken Williamson
 All rights reserved.

 Certain inventions and disclosures in this file may be claimed within
 patents owned or patent applications filed by Ulbora Labs LLC., or third
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
	//env "UlboraApiGateway/environment"
	cb "UlboraApiGateway/circuitbreaker"
	mgr "UlboraApiGateway/managers"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

//HandleGetClusterGwRoutes HandleGetClusterGwRoutes
func (h Handler) HandleGetClusterGwRoutes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var gwrg mgr.GatewayRoutes
		gwrg.GwDB.DbConfig = h.DbConfig
		gwrg.GwCacheHost = getCacheHost()
		cid := r.Header.Get("u-client-id")
		gwrg.ClientID, _ = strconv.ParseInt((cid), 10, 0)
		gwrg.APIKey = r.Header.Get("u-api-key")

		//gwr.GwCacheHost = env.GetCacheHost()
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		var route string
		if vars != nil {
			route = vars["route"]
		} else {
			route = r.URL.Query().Get("route")
		}
		gwrg.Route = route
		res := gwrg.GetClusterGwRoutes()
		resJSON, err := json.Marshal(res)
		//fmt.Print("json out: ")
		//fmt.Println(res)
		if err != nil {
			log.Println(err.Error())
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(resJSON))
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

//HandleClearClusterGwRoutes HandleClearClusterGwRoutes
func (h Handler) HandleClearClusterGwRoutes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "DELETE":
		var gwrc mgr.GatewayRoutes
		gwrc.GwDB.DbConfig = h.DbConfig
		gwrc.GwCacheHost = getCacheHost()
		cid := r.Header.Get("u-client-id")
		gwrc.ClientID, _ = strconv.ParseInt((cid), 10, 0)
		gwrc.APIKey = r.Header.Get("u-api-key")

		//gwr.GwCacheHost = env.GetCacheHost()
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		var route string
		if vars != nil {
			route = vars["route"]
		} else {
			route = r.URL.Query().Get("route")
		}
		gwrc.Route = route
		suc := gwrc.ClearClusterGwRoutes()
		var res mgr.ClusterResponse
		res.Success = suc
		resJSON, err := json.Marshal(res)
		//fmt.Print("json out: ")
		//fmt.Println(res)
		if err != nil {
			log.Println(err.Error())
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(resJSON))
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

//HandleTripClusterBreaker HandleTripClusterBreaker
func (h Handler) HandleTripClusterBreaker(w http.ResponseWriter, r *http.Request) {
	var gwrt mgr.GatewayRoutes
	gwrt.GwDB.DbConfig = h.DbConfig
	gwrt.GwCacheHost = getCacheHost()
	cid := r.Header.Get("u-client-id")
	gwrt.ClientID, _ = strconv.ParseInt((cid), 10, 0)
	gwrt.APIKey = r.Header.Get("u-api-key")

	w.Header().Set("Content-Type", "application/json")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		switch r.Method {
		case "POST":
			var b ClusterBreaker
			decoder := json.NewDecoder(r.Body)
			error := decoder.Decode(&b)
			b.ClientID = gwrt.ClientID
			if error != nil {
				log.Println(error.Error())
				http.Error(w, error.Error(), http.StatusBadRequest)
			} else if b.ClientID == 0 || b.RestRouteID == 0 || b.RouteURIID == 0 || b.OpenFailCode == 0 {
				http.Error(w, "bad request", http.StatusBadRequest)
			} else {
				var bk cb.Breaker
				bk.RouteURIID = b.RouteURIID
				bk.FailoverRouteName = b.FailoverRouteName
				bk.FailureCount = b.FailureCount
				bk.FailureThreshold = b.FailureThreshold
				bk.HealthCheckTimeSeconds = b.HealthCheckTimeSeconds
				bk.OpenFailCode = b.OpenFailCode
				bk.RestRouteID = b.RestRouteID
				bk.RouteURIID = b.RouteURIID
				bk.ClientID = gwrt.ClientID
				gwrt.Route = b.Route
				resOut := gwrt.TripClusterBreaker(&bk)
				gwrt.ClearClusterGwRoutes()
				resJSON, err := json.Marshal(resOut)
				if err != nil {
					log.Println(error.Error())
				}
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, string(resJSON))
			}
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

//HandleResetClusterBreaker HandleResetClusterBreaker
func (h Handler) HandleResetClusterBreaker(w http.ResponseWriter, r *http.Request) {
	var gwrr mgr.GatewayRoutes
	gwrr.GwDB.DbConfig = h.DbConfig
	gwrr.GwCacheHost = getCacheHost()
	cid := r.Header.Get("u-client-id")
	gwrr.ClientID, _ = strconv.ParseInt((cid), 10, 0)
	gwrr.APIKey = r.Header.Get("u-api-key")

	w.Header().Set("Content-Type", "application/json")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		switch r.Method {
		case "POST":
			var b ClusterBreaker
			decoder := json.NewDecoder(r.Body)
			error := decoder.Decode(&b)
			b.ClientID = gwrr.ClientID
			gwrr.Route = b.Route
			if error != nil {
				log.Println(error.Error())
				http.Error(w, error.Error(), http.StatusBadRequest)
			} else if b.ClientID == 0 || b.RouteURIID == 0 {
				http.Error(w, "bad request", http.StatusBadRequest)
			} else {
				suc := gwrr.ResetClusterBreaker(b.RouteURIID)
				var res mgr.ClusterResponse
				res.Success = suc
				gwrr.ClearClusterGwRoutes()
				resJSON, err := json.Marshal(res)
				if err != nil {
					log.Println(error.Error())
				}
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, string(resJSON))
			}
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

//HandleClusterSaveRouteError HandleClusterSaveRouteError
func (h Handler) HandleClusterSaveRouteError(w http.ResponseWriter, r *http.Request) {
	cid := r.Header.Get("u-client-id")
	clientID, _ := strconv.ParseInt((cid), 10, 0)

	w.Header().Set("Content-Type", "application/json")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		switch r.Method {
		case "POST":
			var el ErrorLog
			decoder := json.NewDecoder(r.Body)
			error := decoder.Decode(&el)
			el.ClientID = clientID
			if error != nil {
				log.Println(error.Error())
				http.Error(w, error.Error(), http.StatusBadRequest)
			} else if el.ClientID == 0 || el.RouteID == 0 || el.RouteURIID == 0 || el.ErrCode == 0 {
				http.Error(w, "bad request", http.StatusBadRequest)
			} else {
				suc, err := h.ErrDB.SaveRouteError(el.ClientID, el.ErrCode, el.Message, el.RouteID, el.RouteURIID)
				if err != nil {
					log.Println(error.Error())
				}
				var resOut mgr.ClusterResponse
				resOut.Success = suc
				resJSON, err := json.Marshal(resOut)
				if err != nil {
					log.Println(error.Error())
				}
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, string(resJSON))
			}
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

//HandleClusterSaveRoutePerformance HandleClusterSaveRoutePerformance
func (h Handler) HandleClusterSaveRoutePerformance(w http.ResponseWriter, r *http.Request) {
	cid := r.Header.Get("u-client-id")
	clientID, _ := strconv.ParseInt((cid), 10, 0)

	w.Header().Set("Content-Type", "application/json")
	cType := r.Header.Get("Content-Type")
	if cType != "application/json" {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		switch r.Method {
		case "POST":
			var p PerformanceLog
			decoder := json.NewDecoder(r.Body)
			error := decoder.Decode(&p)
			p.ClientID = clientID
			if error != nil {
				log.Println(error.Error())
				http.Error(w, error.Error(), http.StatusBadRequest)
			} else if p.ClientID == 0 || p.RouteID == 0 || p.RouteURIID == 0 {
				http.Error(w, "bad request", http.StatusBadRequest)
			} else {
				suc := h.MonDB.SaveRoutePerformance(p.ClientID, p.RouteID, p.RouteURIID, p.Latency)
				var resOut mgr.ClusterResponse
				resOut.Success = suc
				resJSON, err := json.Marshal(resOut)
				if err != nil {
					log.Println(error.Error())
				}
				w.WriteHeader(http.StatusOK)
				fmt.Fprint(w, string(resJSON))
			}
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
