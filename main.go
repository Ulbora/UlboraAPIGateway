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
	"fmt"
	"net/http"
	"os"

	mgr "UlboraApiGateway/managers"

	"github.com/gorilla/mux"
)

type authHeader struct {
	token    string
	clientID int64
	userID   string
	hashed   bool
}

var gatewayDB mgr.GatewayDB

//var gwr mgr.GatewayRoutes

func main() {

	if os.Getenv("MYSQL_PORT_3306_TCP_ADDR") != "" {
		gatewayDB.DbConfig.Host = os.Getenv("MYSQL_PORT_3306_TCP_ADDR")
	} else if os.Getenv("DATABASE_HOST") != "" {
		gatewayDB.DbConfig.Host = os.Getenv("DATABASE_HOST")
	} else {
		gatewayDB.DbConfig.Host = "localhost:3306"
	}

	if os.Getenv("DATABASE_USER_NAME") != "" {
		gatewayDB.DbConfig.DbUser = os.Getenv("DATABASE_USER_NAME")
	} else {
		gatewayDB.DbConfig.DbUser = "admin"
	}

	if os.Getenv("DATABASE_USER_PASSWORD") != "" {
		gatewayDB.DbConfig.DbPw = os.Getenv("DATABASE_USER_PASSWORD")
	} else {
		gatewayDB.DbConfig.DbPw = "admin"
	}

	if os.Getenv("DATABASE_NAME") != "" {
		gatewayDB.DbConfig.DatabaseName = os.Getenv("DATABASE_NAME")
	} else {
		gatewayDB.DbConfig.DatabaseName = "ulbora_api_gateway"
	}
	gatewayDB.ConnectDb()
	defer gatewayDB.CloseDb()
	//gwr.GwDB = gatewayDB

	fmt.Println("Api Gateway running!")
	router := mux.NewRouter()
	//super admin client services
	router.HandleFunc("/rs/gwClient/add", handleClientChange)
	router.HandleFunc("/rs/gwClient/update", handleClientChange)
	router.HandleFunc("/rs/gwClient/get/{clientId}", handleClient)
	router.HandleFunc("/rs/gwClient/list", handleClientList)
	router.HandleFunc("/rs/gwClient/delete/{clientId}", handleClient)

	// super admin restRoute services
	router.HandleFunc("/rs/gwRestRouteSuper/add", handleRestRouteSuperChange)
	router.HandleFunc("/rs/gwRestRouteSuper/update", handleRestRouteSuperChange)
	router.HandleFunc("/rs/gwRestRouteSuper/get/{id}/{clientId}", handleRestRouteSuper)
	router.HandleFunc("/rs/gwRestRouteSuper/list/{clientId}", handleRestRouteSuperList)
	router.HandleFunc("/rs/gwRestRouteSuper/delete/{id}/{clientId}", handleRestRouteSuper)

	// super admin routeUrl services
	router.HandleFunc("/rs/gwRouteUrlSuper/add", handleRouteURLSuperChange)
	router.HandleFunc("/rs/gwRouteUrlSuper/update", handleRouteURLSuperChange)
	router.HandleFunc("/rs/gwRouteUrlSuper/get/{id}/{routeId}/{clientId}", handleRouteURLSuper)
	router.HandleFunc("/rs/gwRouteUrlSuper/list/{routeId}/{clientId}", handleRouteURLSuperList)
	router.HandleFunc("/rs/gwRouteUrlSuper/delete/{id}/{routeId}/{clientId}", handleRouteURLSuper)
	router.HandleFunc("/rs/gwRouteUrlSuper/activate", handleRouteURLActivateSuper)

	// admin restRoute services
	router.HandleFunc("/rs/gwRestRoute/add", handleRestRouteChange)
	router.HandleFunc("/rs/gwRestRoute/update", handleRestRouteChange)
	router.HandleFunc("/rs/gwRestRoute/get/{id}", handleRestRoute)
	router.HandleFunc("/rs/gwRestRoute/list", handleRestRouteList)
	router.HandleFunc("/rs/gwRestRoute/delete/{id}", handleRestRoute)

	// admin routeUrl services
	router.HandleFunc("/rs/gwRouteUrl/add", handleRouteURLChange)
	router.HandleFunc("/rs/gwRouteUrl/update", handleRouteURLChange)
	router.HandleFunc("/rs/gwRouteUrl/get/{id}/{routeId}", handleRouteURL)
	router.HandleFunc("/rs/gwRouteUrl/list/{routeId}", handleRouteURLList)
	router.HandleFunc("/rs/gwRouteUrl/delete/{id}/{routeId}", handleRouteURL)
	router.HandleFunc("/rs/gwRouteUrl/activate", handleRouteURLActivate)

	//gateway routes
	router.HandleFunc("/np/{route}/{rname}/{fpath:[^.]+}", handleGwRoute)
	router.HandleFunc("/{route}/{fpath:[^.]+}", handleGwRoute)
	http.ListenAndServe(":3011", router)
}
