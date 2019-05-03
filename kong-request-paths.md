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
curl -i -X POST --url http://localhost:8001/apis/ -d 'name=eventapi' -d 'upstream_url=http://13.67.182.205:3000' -d 'request_path=/eventapi/' -d 'strip_request_path=true'

HTTP/1.1 201 Created
Date: Fri, 03 May 2019 00:28:15 GMT
Content-Type: application/json; charset=utf-8
Transfer-Encoding: chunked
Connection: keep-alive
Access-Control-Allow-Origin: *
Server: kong/0.9.9

{"upstream_url":"http:\/\/13.67.182.205:3000","strip_request_path":true,"request_path":"\/eventapi","id":"67981bba-9640-4189-948f-9f1a6fe42df4","created_at":1556843295000,"preserve_host":false,"name":"eventapi"}
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

## API-Key
```
curl -i -X POST --url http://localhost:8001/consumers/ --data "username=mavericks"

HTTP/1.1 201 Created
Date: Thu, 02 May 2019 04:44:29 GMT
Content-Type: application/json; charset=utf-8
Transfer-Encoding: chunked
Connection: keep-alive
Access-Control-Allow-Origin: *
Server: kong/0.9.9

{"username":"mavericks","created_at":1556772269000,"id":"46b395a5-42dc-4cad-8d6c-066ba8c3e8ba"}

======

curl -i -X POST --url http://localhost:8001/consumers/mavericks/key-auth --data ''

HTTP/1.1 201 Created
Date: Thu, 02 May 2019 04:44:55 GMT
Content-Type: application/json; charset=utf-8
Transfer-Encoding: chunked
Connection: keep-alive
Access-Control-Allow-Origin: *
Server: kong/0.9.9

{"key":"685cd6e92bf045b0b0b217fec6e219c9","consumer_id":"46b395a5-42dc-4cad-8d6c-066ba8c3e8ba","created_at":1556772295000,"id":"6b7e113c-2d67-4ee9-aee1-36f458f5f0ba"}

======

curl -i -X POST http://localhost:8001/apis/eventapi/plugins --data "name=key-auth"

HTTP/1.1 201 Created
Date: Fri, 03 May 2019 00:28:36 GMT
Content-Type: application/json; charset=utf-8
Transfer-Encoding: chunked
Connection: keep-alive
Access-Control-Allow-Origin: *
Server: kong/0.9.9

{"api_id":"67981bba-9640-4189-948f-9f1a6fe42df4","id":"e0c8b911-e26e-4401-aeed-563515faf6e7","created_at":1556843316000,"enabled":true,"name":"key-auth","config":{"key_names":["apikey"],"hide_credentials":false}}


=======

curl -i -X POST http://localhost:8001/apis/userapi/plugins --data "name=key-auth"

HTTP/1.1 201 Created
Date: Thu, 02 May 2019 05:00:45 GMT
Content-Type: application/json; charset=utf-8
Transfer-Encoding: chunked
Connection: keep-alive
Access-Control-Allow-Origin: *
Server: kong/0.9.9

{"api_id":"b8ad0699-f82e-4362-96b6-cf6eeaf46e4d","id":"164f4c8f-d83b-493f-a491-e5a312c02dab","created_at":1556773245000,"enabled":true,"name":"key-auth","config":{"key_names":["apikey"],"hide_credentials":false}}

=======

curl -i -X POST http://localhost:8001/apis/bookapi/plugins --data "name=key-auth"

HTTP/1.1 201 Created
Date: Thu, 02 May 2019 05:01:42 GMT
Content-Type: application/json; charset=utf-8
Transfer-Encoding: chunked
Connection: keep-alive
Access-Control-Allow-Origin: *
Server: kong/0.9.9

{"api_id":"bda29a4a-a827-40c2-ba4f-2d688af3ea32","id":"4d04c002-56e3-478c-8d0d-151d46404d2a","created_at":1556773302000,"enabled":true,"name":"key-auth","config":{"key_names":["apikey"],"hide_credentials":false}}

```
