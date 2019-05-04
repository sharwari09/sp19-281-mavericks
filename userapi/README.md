## API endpoints:
### Ping
- <b>Method:</b> GET
- <b>URI:</b> 
```
    {{apihost}}:{{apiport}}/ping
```
### User signup
- <b>Method:</b> POST
- <b>URI: </b>
```
    {{apihost}}:{{apiport}}/users/signup
```
- <b>Headers</b>
```
    Content-Type: application/json
```
- <b>Request Attributes</b><br/>
![](https://github.com/nguyensjsu/sp19-281-mavericks/blob/master/images/user-signup-api.PNG)
- <b>Body</b>
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
- <b>Method:</b> POST
- <b>URI:</b> 
```
    {{apihost}}:{{apiport}}/users/signin
```
- <b>Headers</b>
```
    Content-Type: application/json
```
- <b>Request Attributes</b><br/>
![](https://github.com/nguyensjsu/sp19-281-mavericks/blob/master/images/user-login-api.PNG)
- <b>Body</b>
```
    {
    	"email": "devv@gmail.com",
    	"password": "welcome123"
    }
```
### Get user by ID
- <b>Method:</b> GET
- <b>URI:</b>
```
    {{apihost}}:{{apiport}}/users/{user-id}
```
### Get all users
- <b>Method:</b> GET
- <b>URI:</b>
```
    {{apihost}}:{{apiport}}/users
```
### Get user by email
- <b>Method:</b> GET
- <b>URI:</b>
```
    {{apihost}}:{{apiport}}/users?email={{email-id}}
```
### Delete by user ID
- <b>Method:</b> DELETE
- <b>URI:</b>
```
    {{apihost}}:{{apiport}}/users/{{user-id}}
```
### Delete by email ID
- <b>Method:</b> DELETE
- <b>URI:</b> 
```
    {{apihost}}:{{apiport}}/users?email={{email-id}}
```
