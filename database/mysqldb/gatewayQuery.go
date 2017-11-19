package mysqldb

/*
 Copyright (C) 2016 Ulbora Labs Inc. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2016 Ken Williamson
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

// ContentQuery is a content select query
const (
	ConnectionTestQuery = "SELECT count(*) from client"
	//client
	InsertClientQuery  = "INSERT INTO client (client_id, api_key, enabled, level) VALUES (?, ?, ?, ?) "
	UpdateClientQuery  = "UPDATE client set api_key = ?, enabled = ?, level = ? WHERE client_id = ? "
	ClientGetQuery     = "select * from client WHERE client_id = ? "
	ClientGetListQuery = "select * from client order by client_id "
	ClientDeleteQuery  = "delete from client WHERE client_id = ? "

	//route performance
	InsertRoutePerformanceQuery = "INSERT INTO route_performance (calls, latency_ms_total, entered, route_url_id, " +
		"route_url_rest_route_id, route_url_rest_route_client_id) " +
		"VALUES(?, ?, ?, ?, ?, ?) "
	RoutePerformanceGetQuery = "SELECT * FROM route_performance WHERE route_url_id = ? and route_url_rest_route_id = ? " +
		"and route_url_rest_route_client_id = ? "
	RoutePerformanceRemoveOldQuery = "DELETE FROM route_performance " +
		"WHERE entered < (NOW()- INTERVAL 90 DAY)"

	//route error
	InsertRouteErrorQuery = "INSERT INTO route_error (code, message, entered, route_url_id, " +
		"route_url_rest_route_id, route_url_rest_route_client_id) " +
		"VALUES(?, ?, ?, ?, ?, ?) "
	RouteErrorGetQuery = "SELECT * FROM route_error WHERE route_url_id = ? and route_url_rest_route_id = ? " +
		"and route_url_rest_route_client_id = ? "
	RouteErrorRemoveOldQuery = "DELETE FROM route_error " +
		"WHERE entered < (NOW()- INTERVAL 90 DAY)"

	//circuit breaker
	InsertRouteBreakerQuery = "INSERT INTO breaker (failure_threshold, health_check_time_seconds, failover_route_name, " +
		"open_fail_code, route_url_id, " +
		"route_url_rest_route_id, route_url_rest_route_client_id) " +
		"VALUES(?, ?, ?, ?, ?, ?, ?) "

	UpdateRouteBreakerConfigQuery = "UPDATE breaker set failure_threshold = ?, health_check_time_seconds = ?, " +
		"failover_route_name = ?, open_fail_code = ? " +
		"WHERE id = ? and route_url_id = ? and route_url_rest_route_id = ? and  route_url_rest_route_client_id = ? "

	UpdateRouteBreakerFailQuery = "UPDATE breaker set failure_count = ?, last_failure_time = ? " +
		"WHERE id = ? and route_url_id = ? and route_url_rest_route_id = ? and  route_url_rest_route_client_id = ? "

	BreakerGetQuery = "select id, failure_threshold, health_check_time_seconds, failover_route_name, open_fail_code, " +
		"failure_count, last_failure_time, route_url_id, route_url_rest_route_id, route_url_rest_route_client_id " +
		"from breaker WHERE route_url_id = ? and route_url_rest_route_id = ? and  " +
		"route_url_rest_route_client_id = ? "

	BreakerDeleteQuery = "delete from breaker WHERE route_url_id = ? and route_url_rest_route_id = ? and " +
		"route_url_rest_route_client_id = ? "

	// route
	InsertRestRouteQuery  = "INSERT INTO rest_route (route, client_id) VALUES (?, ?) "
	UpdateRestRouteQuery  = "UPDATE rest_route set route = ? WHERE id = ? and client_id = ? "
	RestRouteGetQuery     = "select * from rest_route WHERE id = ? and client_id = ?"
	RestRouteGetListQuery = "select * from rest_route WHERE client_id = ?"
	RestRouteDeleteQuery  = "delete from rest_route WHERE id = ? and client_id = ?"

	// rest route
	InsertRouteURLQuery           = "INSERT INTO route_url (name, url, active, rest_route_id, rest_route_client_id) VALUES (?, ?, ?, ?, ?) "
	UpdateRouteURLQuery           = "UPDATE route_url set name = ?, url = ? WHERE id = ? and rest_route_id = ? and rest_route_client_id = ? "
	ActivateRouteURLQuery         = "UPDATE route_url set active = 1 WHERE id = ? and rest_route_id = ? and rest_route_client_id = ? "
	DeactivateOtherRouteURLsQuery = "UPDATE route_url set active = 0 WHERE id != ? and rest_route_id = ? and rest_route_client_id = ? "
	RouteURLGetQuery              = "select * from route_url WHERE id = ? and rest_route_id = ? and rest_route_client_id = ? "
	RouteURLGetListQuery          = "select * from route_url WHERE rest_route_id = ? and rest_route_client_id = ? "
	RouteNameURLGetListQuery      = "select rr.route, ru.name, ru.url, ru.active " +
		"from route_url ru " +
		"inner join rest_route rr " +
		"on ru.rest_route_id = rr.id " +
		"and ru.rest_route_client_id = rr.client_id " +
		"inner join client c " +
		"on c.client_id = rr.client_id " +
		"WHERE rr.route = ? and ru.rest_route_client_id = ? and c.api_key = ? "
	RouteURLDeleteQuery = "delete from route_url WHERE id = ? and rest_route_id = ? and rest_route_client_id = ? and active = 0 "

	GetRouteURLsQuery = "select rr.route, rl.name, rl.url, rl.active " +
		" from rest_route rr inner join client c " +
		" on c.client_id = rr.client_id " +
		" INNER join route_url rl " +
		" on rl.rest_route_id = rr.id " +
		" and rl.rest_route_client_id = rr.client_id " +
		" where c.client_id = ? and rr.route = ? "

	GetActiveRouteURLQuery = "select rr.route, rl.name, rl.url, rl.active " +
		" from rest_route rr inner join client c " +
		" on c.client_id = rr.client_id " +
		" INNER join route_url rl " +
		" on rl.rest_route_id = rr.id " +
		" and rl.rest_route_client_id = rr.client_id " +
		" where c.client_id = ? and rr.route = ? and rl.active = 1 "
)
