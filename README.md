# goDDOS
academic purpose ddos server & client written in go

https://en.wikipedia.org/wiki/Denial-of-service_attack

![Alt text](https://raw.githubusercontent.com/arnaucode/goDDOS/master/concept.png "concept")



### v1
This version works as:
- the client every X time try to connect to the server
- when the server gets a new connection from a client, stores that connection in memory
- when a new order is written in the cmd of the server, sends that order to all the stored clients
- each client receive the order and starts the attack

example command line on the server:
```
ddos http://website.com 10

ddos [website] [num of gets for each client]
```

### v2
This version works as:
- the server have an endpoint where is placed the current order
- each client every X time, performs a GET to the endpoint of the server, to see if there is any new order
- when a new order is written in the cmd of the server, it will return that order in the server endpoint
- then, the clients goes to the endpoint to get the order
- once the client have the order, starts the attack

example command line on the server:
```
http://website.com 10

[website] [num of gets for each client]
```
