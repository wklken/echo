# echo
An server will echo the requests detail into response. For api testing / benchmark for your proxies like APIGateway/ReverseProxy.


# run

```
make build

./echo 9001
```

# apis

#### /ping/

- AllowedMethods: all
- return 200 with pong

```
curl -XPOST http://127.0.0.1:9001/ping/

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

#### /sleep/{sleep}/

- AllowedMethods: all
- will sleep for `{sleep}` ms
- return 200 with ok

```
curl -XPOST http://127.0.0.1:9001/sleep/10/

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

```
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

#### /status/{status}/

- AllowedMethods: all
- return `{status}`

```
curl -XPOST http://127.0.0.1:9001/status/500/

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

#### /echo/

- AllowedMethods: all
- return 200, with the request contents in json

```
curl -XPOST http://127.0.0.1:9001/echo/\?a\=1\&b\=2

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
