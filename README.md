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

#### /sleep/{sleep}/

- AllowedMethods: all
- will sleep for `{sleep}` ms
- return 200 with ok

#### /status/{status}/

- AllowedMethods: all
- return `{status}`

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
