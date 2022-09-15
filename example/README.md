# Load balancer example
## Diagram
```mermaid
graph LR
A[Client] -->|127.0.0.1:3000/hello| B{Load Balancer}    
    B -->|127.0.0.1:5001/hello| C[First backend]
    B -->|127.0.0.1:5002/hello| D[Second backend]
    B -->|127.0.0.1:5003/hello| E[Third backend]
```
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
$ curl -XPOST -d "" 127.0.0.1:3000/hello 
Hello First

# Response of "Second" backend
$ curl -XPOST -d "" 127.0.0.1:3000/hello 
Hello Second
```
