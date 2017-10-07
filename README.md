Ulbora API Gateway
==============

Multi-user API gateway with self service portal: http://www.myapigateway.com

Copyright (C) 2016 Ulbora Labs Inc. (www.ulboralabs.com)
All rights reserved.

Copyright (C) 2016 Ken Williamson
All rights reserved.

Certain inventions and disclosures in this file may be claimed within
patents owned or patent applications filed by Ulbora Labs Inc., or third
parties.


User Admin Portal: https://github.com/Ulbora/ApiGatewayUserPortal.git


# Using API Gateway Routes
The Ulbora API Gateway routes REST services calls to endpoint assigned through the user portal.

## Headers For Gateway
- clientId: Your assigned client id
- apiKey: Your assigned API Key
- Any other headers required for your micro services

## Allowed HTTP Methods
- POST
- PUT
- PATCH
- GET
- DELETE
- OPTIONS


## Gateway Routes
### Local Non-Prod
- http://localhost:3011/np/routeID/routeName/yourRoute
- (example): http://localhost:3011/np/challenge/blue/rs/challenge/en_us?g=g&b=b
- Note: 
- routeID is: challenge
- routeName is: blue

### Local Prod

- http://localhost:3011/routeID/yourRoute
- (example): http://localhost:3011/challenge/rs/challenge?name=sam&age=44
- Note: 
- routeID is: challenge

## Add Client


## Headers
- Content-Type: application/json (for POST and PUT)
- Authorization: Bearer aToken (POST, PUT, and DELETE. No token required for get services.)
- clientId: clientId (example 33477)