Ulbora API Gateway
==============

Multi-user lightweight API Gateway with a self service portal: http://www.myapigateway.com

Copyright (C) 2016 Ulbora Labs Inc. (www.ulboralabs.com)
All rights reserved.

Copyright (C) 2016 Ken Williamson
All rights reserved.

Certain inventions and disclosures in this file may be claimed within
patents owned or patent applications filed by Ulbora Labs Inc., or third
parties.


User Admin Portal: https://github.com/Ulbora/ApiGatewayUserPortal.git


# Using API Gateway Routes
- The Ulbora API Gateway routes REST service calls to endpoints assigned through the user portal.
- There can be multiple API endpoints mapped to any route; blue, green, or any name you choose.
- CORS is passed through the gateway, so if CORS is enabled in your REST services, it works automatically in the gateway.
- Authentication is also passed through the gateway. If you pass any type of token in a header, it will be passed through automatically.
- Any other headers are also passed through the gateway automatically.
- The gateway works with both JSON and XML bodies.


## Headers For Gateway Route Calls
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
- yourRoute is: rs/challenge/en_us?g=g&b=b which can be mappend in the user portal to something like https://www.youapi/rs/challenge/en_us?g=g&b=b

### Local Prod

- http://localhost:3011/routeID/yourRoute
- (example): http://localhost:3011/challenge/rs/challenge?name=sam&age=44
- Note: 
- routeID is: challenge
- yourRoute is: /rs/challenge?name=sam&age=44 which can be mappend in the user portal to something like https://www.youapi/rs/challenge?name=sam&age=44

### Active Production Route
The User Admin Portal allows you to make any route URL the active production route with the click of a switch.
Using Non-Prod routes allows you to test services before placing them in production.

## Add Client


## Headers
- Content-Type: application/json (for POST and PUT)
- Authorization: Bearer aToken (POST, PUT, and DELETE. No token required for get services.)
- clientId: clientId (example 33477)