## Get url params: 
```
https://golangcode.com/get-a-url-parameter-from-a-request/
```
## The sample JSON body expected by POST calls:
```
{
	"firstname": "rambo",
	"lastname": "Z",
	"address": {
		"city": "jalgaon",
		"state": "MH",
		"street": "1st",
		"zip": "95000"
	},
	"email": "abcd@circus.com",
	"password": "<pwd>"
}
```
### Run userapi dockers with ELB:
```
docker-compose scale userapi=3 proxy=1
```
