## User API:
Command: <br/>
```
curl -i -X POST --url http://localhost:8001/apis/ -d 'name=userapi' -d 'upstream_url=http://40.113.244.137:5000' -d 'request_path=/userapi/' -d 'strip_request_path=true'
```
Output: <br/>
```
HTTP/1.1 201 Created
Date: Thu, 02 May 2019 02:51:50 GMT
Content-Type: application/json; charset=utf-8
Transfer-Encoding: chunked
Connection: keep-alive
Access-Control-Allow-Origin: *
Server: kong/0.9.9

{"upstream_url":"http:\/\/40.113.244.137:5000","strip_request_path":true,"request_path":"\/userapi","id":"f7987c8b-cde9-49ba-aff3-e2cadced8fb5","created_at":1556765510000,"preserve_host":false,"name":"userapi"}
```

## Events API: 


## Booking API:


## Dashboard API:
