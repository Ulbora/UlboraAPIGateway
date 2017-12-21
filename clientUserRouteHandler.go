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

	mng "UlboraApiGateway/managers"

	uoauth "github.com/Ulbora/go-ulbora-oauth2"
)

func handleUserClient(w http.ResponseWriter, r *http.Request) {
	auth := getAuth(r)
	me := new(uoauth.Claim)
	me.Role = "admin"

	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	//clientID, errID := strconv.ParseInt(vars["clientId"], 10, 0)
	//if errID != nil {
	//	http.Error(w, "bad request", http.StatusBadRequest)
	//	}
	//fmt.Print("id is: ")
	//fmt.Println(id)
	switch r.Method {
	case "GET":
		me.URI = "/ulbora/rs/gwClientUser/get"
		me.Scope = "read"
		valid := auth.Authorize(me)
		if valid != true {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			client := new(mng.Client)
			client.ClientID = auth.ClientID
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
	}
}
