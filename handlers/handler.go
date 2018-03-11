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
	db "UlboraApiGateway/database"
	gwerr "UlboraApiGateway/gwerrors"
	//gwerr "UlboraApiGateway/gwerrors"
	cb "UlboraApiGateway/circuitbreaker"
	gwmon "UlboraApiGateway/monitor"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	uoauth "github.com/Ulbora/go-ulbora-oauth2"
)

//var errDB gwerr.GatewayErrorMonitor

//Handler Handler
type Handler struct {
	DbConfig db.DbConfig
	ErrDB    gwerr.GatewayErrorMonitor
	MonDB    gwmon.GatewayPerformanceMonitor
	CbDB     cb.CircuitBreaker
}

// //SetManager set manager
// func SetManager(db gwerr.GatewayErrorMonitor) {
// 	errDB = db
// }

type authHeader struct {
	token    string
	clientID int64
	userID   string
	hashed   bool
}

func parseQueryString(vals url.Values) string {
	var rtn = ""
	var first = true
	for key, value := range vals {
		if first == true {
			first = false
			rtn += "?" + key + "=" + value[0]
		} else {
			rtn += "&" + key + "=" + value[0]
		}
	}
	return rtn
}

func getCacheHost() string {
	var rtn = ""
	if os.Getenv("CACHE_HOST") != "" {
		rtn = os.Getenv("CACHE_HOST")
	} else {
		rtn = "http://localhost:3010"
	}
	return rtn
}

func buildHeaders(pr *http.Request, sr *http.Request) {
	h := pr.Header
	for hn, v := range h {
		//fmt.Print("header: ")
		//fmt.Print(hn)
		//fmt.Print(" value: ")
		//fmt.Println(v[0])
		sr.Header.Set(hn, v[0])
	}
}

func buildRespHeaders(pw *http.Response, sw http.ResponseWriter) {
	h := pw.Header
	//var cnt = 0
	for hn, v := range h {
		// cnt++
		// fmt.Print("header: ")
		// fmt.Print(hn)
		// fmt.Print(" value: ")
		// fmt.Println(v[0])
		// if cnt > 5 {
		// 	break
		// }
		sw.Header().Set(hn, v[0])
	}
}

func getAuth(req *http.Request) *uoauth.Oauth {
	changeHeader := getHeaders(req)
	auth := new(uoauth.Oauth)
	auth.Token = changeHeader.token
	auth.ClientID = changeHeader.clientID
	auth.UserID = changeHeader.userID
	auth.Hashed = changeHeader.hashed
	if os.Getenv("OAUTH2_VALIDATION_URI") != "" {
		auth.ValidationURL = os.Getenv("OAUTH2_VALIDATION_URI")
	} else {
		auth.ValidationURL = "http://localhost:3000/rs/token/validate"
	}
	return auth
}

func getHeaders(req *http.Request) *authHeader {
	var rtn = new(authHeader)
	authHeader := req.Header.Get("Authorization")
	tokenArray := strings.Split(authHeader, " ")
	if len(tokenArray) == 2 {
		rtn.token = tokenArray[1]
		//fmt.Println(rtn.token)
	}
	userIDHeader := req.Header.Get("userId")
	rtn.userID = userIDHeader

	clientIDHeader := req.Header.Get("clientId")
	clientID, err := strconv.ParseInt(clientIDHeader, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	rtn.clientID = clientID
	if req.Header.Get("hashed") == "true" {
		rtn.hashed = true
	} else {
		rtn.hashed = false
	}
	//fmt.Println(clientIDHeader)
	//fmt.Println(userIDHeader)
	return rtn
}

func paramsOK(p *passParams) bool {
	var rtn = true
	if p.b == nil || p.code == nil || p.gwr == nil || p.h == nil || p.r == nil || p.rts == nil || p.w == nil {
		rtn = false
	}
	return rtn
}
