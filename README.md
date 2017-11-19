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

# Features
- Circuit Breaker
- Health Check
- Self Healing when breaker is open
- Gateway Analytics
- Blue/Green/Active Routes
- Gateway Error Loggin
- Admin Portal (written in Golang)


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


# Super Admin REST Services

## Headers
- Content-Type: application/json (for POST and PUT)
- Authorization: Bearer aToken (POST, PUT, GET and DELETE.)
- clientId: clientId (example 33477)


## Add Client

```
POST:
URL: http://localhost:3011/rs/gwClient/add

Example Request
{  
   "clientId":4,
   "apiKey":"112233",
   "enabled":true,
   "level":"small"
}
  
```

```
Example Response   

{
    "success": true,
    "id": 19
}

```


## Update client

```
PUT:
URL: http://localhost:3011/rs/gwClient/update

Example Request
{  
   "clientId":4,
   "apiKey":"55555",
   "enabled":true,
   "level":"small"
}
  
```

```
Example Response   

{
    "success": true,
    "id": 11
}

```


## Get Client

```
GET:
URL: http://localhost:3011/rs/gwClient/get/4
  
```

```
Example Response   

{  
   "clientId":4,
   "apiKey":"55555",
   "enabled":true,
   "level":"small"
}

```


## Get Client List

```
GET:
URL: http://localhost:3011/rs/gwClient/list
  
```

```
Example Response   

[
    {  
        "clientId":4,
        "apiKey":"555525",
        "enabled":true,
        "level":"small"
    },
    {
        "clientId":5,
        "apiKey":"5553355",
        "enabled":true,
        "level":"small"
    },
    {
        "clientId":6,
        "apiKey":"5555445",
        "enabled":true,
        "level":"small"
    }
]

```


## Delete Client

```
DELETE:
URL: http://localhost:3011/rs/gwClient/delete/1
  
```

```
Example Response   

{
    "success": true,
    "id": 1
}

```



## Add Route

```
POST:
URL: http://localhost:3011/rs/gwRestRouteSuper/add

Example Request
{  
   "clientId":4,
   "route":"mail"
}
  
```

```
Example Response   

{
    "success": true,
    "id": 19
}

```


## Update Route

```
PUT:
URL: http://localhost:3011/rs/gwRestRouteSuper/update

Example Request
{  
   "id": 84,
   "clientId":4,
   "route":"mailChimp"
}
  
```

```
Example Response   

{
    "success": true,
    "id": 84
}

```


## Get Route

```
GET:
URL: http://localhost:3011/rs/gwRestRouteSuper/get/84/4
  
```

```
Example Response   

{  
   "id": 84,
   "clientId":4,
   "route":"mail"
}

```


## Get Route List

```
GET:
URL: http://localhost:3011/rs/gwRestRouteSuper/list/4
  
```

```
Example Response   

[
    {  
        "id": 84,
        "clientId":4,
        "route":"mail"
    },
    {
        "id": 85,
        "clientId":4,
        "route":"content"
    },
    {
        "id": 86,
        "clientId":4,
        "route":"products"
    }
]

```


## Delete Route

```
DELETE:
URL: http://localhost:3011/rs/gwRestRouteSuper/delete/85/4
  
```

```
Example Response   

{
    "success": true,
    "id": 85
}

```




## Add Route URL

```
POST:
URL: http://localhost:3011/rs/gwRouteUrlSuper/add

Example Request
{  
   "routeId":87,
   "clientId": 403,
   "name": "testMail",
   "url": "google.com/test1"
}
  
```

```
Example Response   

{
    "success": true,
    "id": 19
}

```


## Update Route URL

```
PUT:
URL: http://localhost:3011/rs/gwRouteUrlSuper/update

Example Request
{  
   "id": 188,
   "routeId":87,
   "clientId": 403,
   "name": "testMailGreen",
   "url": "google.com/test1"
}
  
```

```
Example Response   

{
    "success": true,
    "id": 188
}

```


## Activate Route URL

```
PUT:
URL: http://localhost:3011/rs/gwRouteUrlSuper/activate

Example Request
{  
   "id": 188,
   "routeId":87,
   "clientId": 403
}
  
```

```
Example Response   

{
    "success": true,
    "id": 188
}

```


## Get Route URL

```
GET:
URL: http://localhost:3011/rs/gwRouteUrlSuper/get/188/87/403
  
```

```
Example Response   

{  
    "id": 188,
    "routeId":87,
    "clientId": 403,
    "name": "testMailGreen",
    "url": "google.com/test1"
}

```


## Get Route URL List

```
GET:
URL: http://localhost:3011/rs/gwRouteUrlSuper/list/87/403
  
```

```
Example Response   

[
    {
        "id": 188,
        "name": "testMailGreen",
        "url": "google.com/test1",
        "active": true,
        "routeId": 87,
        "clientId": 403
    },
    {
        "id": 189,
        "name": "testMail",
        "url": "google.com/test1",
        "active": false,
        "routeId": 87,
        "clientId": 403
    }
]

```


## Delete Route URL

```
DELETE:
URL: http://localhost:3011/rs/gwRouteUrlSuper/delete/190/87/403
  
```

```
Example Response   

{
    "success": true,
    "id": 190
}

```

# Admin REST Services (User Portal)

## Headers
- Content-Type: application/json (for POST and PUT)
- Authorization: Bearer aToken (POST, PUT, GET and DELETE.)
- clientId: clientId (example 33477)




## Add Route

```
POST:
URL: http://localhost:3011/rs/gwRestRoute/add

Example Request
{     
   "route":"mail"
}
  
```

```
Example Response   

{
    "success": true,
    "id": 19
}

```


## Update Route

```
PUT:
URL: http://localhost:3011/rs/gwRestRoute/update

Example Request
{  
   "id": 84,   
   "route":"mailChimp"
}
  
```

```
Example Response   

{
    "success": true,
    "id": 84
}

```


## Get Route

```
GET:
URL: http://localhost:3011/rs/gwRestRoute/get/84
  
```

```
Example Response   

{  
   "id": 84,  
   "route":"mail"
}

```


## Get Route List

```
GET:
URL: http://localhost:3011/rs/gwRestRoute/list
  
```

```
Example Response   

[
    {  
        "id": 84,
        "clientId":4,
        "route":"mail"
    },
    {
        "id": 85,
        "clientId":4,
        "route":"content"
    },
    {
        "id": 86,
        "clientId":4,
        "route":"products"
    }
]

```


## Delete Route

```
DELETE:
URL: http://localhost:3011/rs/gwRestRoute/delete/85
  
```

```
Example Response   

{
    "success": true,
    "id": 85
}

```




## Add Route URL

```
POST:
URL: http://localhost:3011/rs/gwRouteUrl/add

Example Request
{  
   "routeId":87,   
   "name": "testMail",
   "url": "google.com/test1"
}
  
```

```
Example Response   

{
    "success": true,
    "id": 19
}

```


## Update Route URL

```
PUT:
URL: http://localhost:3011/rs/gwRouteUrl/update

Example Request
{  
   "id": 188,
   "routeId":87,   
   "name": "testMailGreen",
   "url": "google.com/test1"
}
  
```

```
Example Response   

{
    "success": true,
    "id": 188
}

```


## Activate Route URL

```
PUT:
URL: http://localhost:3011/rs/gwRouteUrl/activate

Example Request
{  
   "id": 188,
   "routeId":87   
}
  
```

```
Example Response   

{
    "success": true,
    "id": 188
}

```


## Get Route URL

```
GET:
URL: http://localhost:3011/rs/gwRouteUrl/get/188/87
  
```

```
Example Response   

{  
    "id": 188,
    "routeId":87,
    "clientId": 403,
    "name": "testMailGreen",
    "url": "google.com/test1"
}

```


## Get Route URL List

```
GET:
URL: http://localhost:3011/rs/gwRouteUrl/list/87
  
```

```
Example Response   

[
    {
        "id": 188,
        "name": "testMailGreen",
        "url": "google.com/test1",
        "active": true,
        "routeId": 87,
        "clientId": 403
    },
    {
        "id": 189,
        "name": "testMail",
        "url": "google.com/test1",
        "active": false,
        "routeId": 87,
        "clientId": 403
    }
]

```


## Delete Route URL

```
DELETE:
URL: http://localhost:3011/rs/gwRouteUrl/delete/190/87
  
```

```
Example Response   

{
    "success": true,
    "id": 190
}

```