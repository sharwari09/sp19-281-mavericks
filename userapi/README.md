## API endpoints:
### Ping
- Method: GET
- URI: 
```
    {{apihost}}:{{apiport}}/ping
```
### User signup
- Method: POST
- URI: 
```
    http://{{apihost}}:{{apiport}}/users/signup
```
- HEADERS
```
    Content-Type: application/json
```
- BODY
```
    {
	    "firstname": "Tony",
	    "lastname": "Stark",
	    "password": "welcome",
	    "address": {
	    	"city": "New York",
	    	"state": "NY",
	    	"street": "11th",
	    	"zip": "55210"
	    },
	    "email": "tony@gmail.com"
    }
```
### User login
- Method: POST
- URI: 
```
    {{apihost}}:{{apiport}}/users/signin
```
- HEADERS
```
    Content-Type: application/json
```
- BODY
```
    {
    	"email": "devv@gmail.com",
    	"password": "welcome123"
    }
```

