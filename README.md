# echo

Echo server will echo the requests detail into response.

For api testing or benchmarking for your proxies like APIGateway/ReverseProxy/nginx/traefik.

Build on [chi](https://github.com/go-chi/chi) with simple logical, so the performance is good enough.


# run

```
make build

./echo 9001
```

# apis

### /ping/

- AllowedMethods: all
- return 200 with pong

```shell
$ curl -XPOST http://127.0.0.1:9001/ping/

> POST /ping/ HTTP/1.1
> Host: 127.0.0.1:9001
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Sun, 02 Jun 2019 00:19:19 GMT
< Content-Length: 4
< Content-Type: text/plain; charset=utf-8

pong
```

### /sleep/{sleep}/

- AllowedMethods: all
- will sleep for `{sleep}` ms
- return 200 with ok

```shell
$ curl -XPOST http://127.0.0.1:9001/sleep/10/

> POST /sleep/10/ HTTP/1.1
> Host: 127.0.0.1:9001
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Sun, 02 Jun 2019 00:20:16 GMT
< Content-Length: 2
< Content-Type: text/plain; charset=utf-8

ok
```

and do the benchmark

```shell
$ wrk -t8 -c1000 -d10s http://127.0.0.1:9000/sleep/10/
Running 10s test @ http://127.0.0.1:9000/sleep/10/
  8 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    11.84ms   18.96ms 266.92ms   91.97%
    Req/Sec     8.79k     2.04k   14.77k    64.38%
  699458 requests in 10.04s, 117.40MB read
  Socket errors: connect 0, read 916, write 0, timeout 0
  Non-2xx or 3xx responses: 699458
Requests/sec:  69640.30
Transfer/sec:     11.69MB
```

### /status/{status}/

- AllowedMethods: all
- return `{status}`

```
$ curl -XPOST http://127.0.0.1:9001/status/500/

> POST /status/500/ HTTP/1.1
> Host: 127.0.0.1:9001
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 500 Internal Server Error
< Content-Type: text/plain; charset=utf-8
< X-Content-Type-Options: nosniff
< Date: Sun, 02 Jun 2019 00:21:01 GMT
< Content-Length: 22

Internal Server Error
```

### /echo/

- AllowedMethods: all
- return 200, with the request contents in json

```shell
$ curl -XPOST http://127.0.0.1:9001/echo/\?a\=1\&b\=2

> POST /echo/?a=1&b=2 HTTP/1.1
> Host: 127.0.0.1:9001
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=UTF-8
< Date: Sun, 02 Jun 2019 00:17:57 GMT
< Content-Length: 140

{
    "body": "",
    "form":
    {},
    "headers":
    {
        "Accept": ["*/*"],
        "User-Agent": ["curl/7.54.0"]
    },
    "method": "POST",
    "query":
    {
        "a": ["1"],
        "b": ["2"]
    },
    "url": "/echo/"
}
```

### websocket

- /ws/index/ the websocket test page
- /ws/       the websocket api, will broadcast the message to all client

### file upload/download

- /file/download/{size}/   the size unit is KB

```shell
$ curl -XGET http://127.0.0.1:9001/file/download/1024/ -vv > a.txt

> GET /file/download/1024/ HTTP/1.1
> Host: 127.0.0.1:3000
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Content-Control: private, no-transform, no-store, must-revalidate
< Content-Disposition: attachment; filename=data.txt
< Content-Length: 1048576
< Content-Transfer-Encoding: binary
< Content-Type: text/plain; charset=utf-8
< Expires: 0
< Last-Modified: Tue, 02 Jul 2019 08:23:20 GMT
< Date: Tue, 02 Jul 2019 08:23:20 GMT

$ du -h a.txt
1M
```


- /file/upload/{filename}/

```
$ curl -XPOST \
  'http://127.0.0.1:9001/file/upload/data.txt/' \
  -H 'Content-Type: multipart/form-data' \
  -F data.txt=@/Downloads/data.txt -v

> POST /file/upload/data.txt/ HTTP/1.1
> Host: 127.0.0.1:3000
> User-Agent: curl/7.54.0
> Accept: */*
> Content-Length: 10430
> Expect: 100-continue
> Content-Type: multipart/form-data; boundary=------------------------1fb40c77823315c2
>
< HTTP/1.1 100 Continue
< HTTP/1.1 204 No Content
< Date: Tue, 02 Jul 2019 08:25:48 GMT
```

# TODO

- [ ] httpbin apis

