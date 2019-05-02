## User API
```
curl -i -X POST --url http://localhost:8001/apis/ -d 'name=userapi' -d 'upstream_url=http://168.61.189.227:5000' -d 'request_path=/userapi/' -d 'strip_request_path=true'

HTTP/1.1 201 Created
Date: Thu, 02 May 2019 04:36:25 GMT
Content-Type: application/json; charset=utf-8
Transfer-Encoding: chunked
Connection: keep-alive
Access-Control-Allow-Origin: *
Server: kong/0.9.9

{"upstream_url":"http:\/\/168.61.189.227:5000","strip_request_path":true,"request_path":"\/userapi","id":"b8ad0699-f82e-4362-96b6-cf6eeaf46e4d","created_at":1556771785000,"preserve_host":false,"name":"userapi"}
```

## Events API
```
curl -i -X POST --url http://localhost:8001/apis/ -d 'name=eventapi' -d 'upstream_url=http://52.165.18.94:3000' -d 'request_path=/eventapi/' -d 'strip_request_path=true'

HTTP/1.1 201 Created
Date: Thu, 02 May 2019 04:37:23 GMT
Content-Type: application/json; charset=utf-8
Transfer-Encoding: chunked
Connection: keep-alive
Access-Control-Allow-Origin: *
Server: kong/0.9.9

{"upstream_url":"http:\/\/52.165.18.94:3000","strip_request_path":true,"request_path":"\/eventapi","id":"d96c0175-db24-4442-abd2-66158a98ee74","created_at":1556771843000,"preserve_host":false,"name":"eventapi"}
```

## Booking API
```
curl -i -X POST --url http://localhost:8001/apis/ -d 'name=bookapi' -d 'upstream_url=http://168.61.153.1:4000' -d 'request_path=/bookapi/' -d 'strip_request_path=true'

HTTP/1.1 201 Created
Date: Thu, 02 May 2019 04:38:18 GMT
Content-Type: application/json; charset=utf-8
Transfer-Encoding: chunked
Connection: keep-alive
Access-Control-Allow-Origin: *
Server: kong/0.9.9

{"upstream_url":"http:\/\/168.61.153.1:4000","strip_request_path":true,"request_path":"\/bookapi","id":"bda29a4a-a827-40c2-ba4f-2d688af3ea32","created_at":1556771898000,"preserve_host":false,"name":"bookapi"}
```
