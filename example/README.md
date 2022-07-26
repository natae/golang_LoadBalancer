# Load balancer example
## Run with docker (Server)
```
# Change directory for copy dependencies.
cd ..

docker build -t test-lb -f example/Dockerfile .

docker run -p 3000:3000 test-lb
```

## Test with curl (Client)
```
# Response of "First" backend
$ curl -XPOST -i -d "" 127.0.0.1:3000/hello 
HTTP/1.1 200 OK
Date: Tue, 26 Jul 2022 15:16:13 GMT
Content-Length: 11
Content-Type: text/plain; charset=utf-8

Hello First

# Response of "Second" backend
$curl -XPOST -i -d "" 127.0.0.1:3000/hello 
HTTP/1.1 200 OK
Date: Tue, 26 Jul 2022 15:16:13 GMT
Content-Length: 11
Content-Type: text/plain; charset=utf-8

Hello Second

```
