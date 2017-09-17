package mysqldb

// ContentQuery is a content select query
const (
	ConnectionTestQuery = "SELECT count(*) from client"
	//client
	InsertClientQuery  = "INSERT INTO client (client_id, api_key, enabled, level) VALUES (?, ?, ?, ?) "
	UpdateClientQuery  = "UPDATE client set api_key = ?, enabled = ?, level = ? WHERE client_id = ? "
	ClientGetQuery     = "select * from client WHERE client_id = ? "
	ClientGetListQuery = "select * from client order by client_id "
	ClientDeleteQuery  = "delete from client WHERE client_id = ? "

	// route
	InsertRestRouteQuery  = "INSERT INTO rest_route (route, client_id) VALUES (?, ?) "
	UpdateRestRouteQuery  = "UPDATE rest_route set route = ? WHERE id = ? and client_id = ? "
	RestRouteGetQuery     = "select * from rest_route WHERE id = ? and client_id = ?"
	RestRouteGetListQuery = "select * from rest_route WHERE client_id = ?"
	RestRouteDeleteQuery  = "delete from rest_route WHERE id = ? and client_id = ?"

	// rest route
	InsertRouteURLQuery           = "INSERT INTO route_url (name, url, active, rest_route_id, rest_route_client_id) VALUES (?, ?, ?, ?, ?) "
	UpdateRouteURLQuery           = "UPDATE route_url set name = ?, url = ?, active = ? WHERE id = ? and rest_route_id = ? and rest_route_client_id = ? "
	ActivateRouteURLQuery         = "UPDATE route_url set active = 1 WHERE id = ? and rest_route_id = ? and rest_route_client_id = ? "
	DeactivateOtherRouteURLsQuery = "UPDATE route_url set active = 0 WHERE id != ? and rest_route_id = ? and rest_route_client_id = ? "
	RouteURLGetQuery              = "select * from route_url WHERE id = ? and rest_route_id = ? and rest_route_client_id = ? "
	RouteURLGetListQuery          = "select * from route_url WHERE rest_route_id = ? and rest_route_client_id = ? "
	RouteURLDeleteQuery           = "delete from route_url WHERE id = ? and rest_route_id = ? and rest_route_client_id = ? "

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
		" where c.client_id = ? and rr.route = ? and rl.active = ? "
)
