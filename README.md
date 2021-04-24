# 01011010

Prerequisites on your machine: Docker

Fast start: clone this repository to your local machine.
Run make with the path and filename of your Shipping Ports json file.

``` zsh
git clone https://github.com/sztelzer/01011010
cd 01011010
make BIND_PATH=~/shipping JSON_FILENAME=ports.json
```
If you set a correct path for your json file (both vars) youl see this result after processing it:
```
Attaching to shippingportsclient_1, shippingportsserver_1
shippingportsserver_1  | 2021/04/24 03:24:52 shippingPortsProtocolServer listening on port: :50051
shippingportsclient_1  | 2021/04/24 03:24:53 successfully loaded 1632 shippingPorts from file to server in 363.06375ms
```



This is a set of two services running in docker compose. It's easy to run, but please attention to one detail:
**You must specify a local directory to bind and filename from where the client service will read data to fill the database.**
The specifics of this json file will be discussed further below.

This make command will be your best friend. It tests, builds and runs everything that is necessary. You just need to 
have docker.

## Public REST Interface

**TLDR; There is only one public endpoint, by default in port 8080.**
Example to get Shipping Port with id `MUSIK` (this id is ficticious).
```
curl -X GET localhost:8080/shippingports/MUSIK
```
``` json
{
    "id": "MUSIK",
    "name": "Musik",
    "city": "Musik",
    "country": "Musicalis",
    "alias": [],
    "regions": [],
    "coordinates": [
        -1.00928,
        -11.88878
    ],
    "province": "",
    "timezone": "SharpOnTime",
    "unlocs": [
        "MUSIK"
    ]
}
```
In case you hit a malformed or non existent shipping port code, the response will be the status code `404 not found`.
If you try other http method, will get `invalid request method`
If something awful happens inside our service, it will respond some various `5xx` codes relative to the problem.


## Running the service

**TLDR; having docker and make on your system run make at the root directory of this repository.**
```
01011010 % make BIND_PATH=~/shipping JSON_FILENAME=ports.json
```
The commands in make are cascading dependent, each running all others before:
- `make proto` Update the protocol buffers packages
- `make tidy` go mod tidy all modules
- `make test` go test all tests
- `make build` build go apps
- `make docker` build docker images
- `make run` run docker compose up

`% make == % make all == % make run`

So, just make.

We use cross compilation of go apps to `linux/amd64` by default. If you would like to build it for 
`linux/arm64`, than run `% make ARCH=arm64`.





This project is a collection of two services meant to be used together, organized around common protocol buffers and
interconnected through gRPC. It is built spanning many Go Modules, but don't frail. It's quite simple.

The services are: 
 - ShippingPortsClient - the REST interface, the gateway, conqueror of open worlds.
 - ShippingPortsServer - the database gateway server, connect with gRPC. It makes the interface incredibly easy to
   use thanks to the protocol buffers.
 - Database - the database is... it is anything we may need. For now it's in memory map with a some methods to simplify.

```
Network -> REST -> ShippingPortsClient -> gRPC -> ShippingPortsServer -> Database (Any)
```

Why so many? What are the benefits of this design? First, ShippingPortsClient have a simple interface for public use.
Second, the services are governed by the protocol buffers, so the changes in business logic must happen
before in them and as soon as we change them, we will know early on tests failing.

Decoupling the database from other applications through the gRPC allow us to connect as many different services
we think about, in almost any different languages, extremely fast, securing the access to data from other applications. 

Also, if we are a hit and need to spawn more servers, so be it.

Finally, the database is really away from consumers, and we can better handle their schema relative to the protocol buffers.
 
### One last fact about ShippingPorts service:

Actually, ShippingPorts service has one more super power, but it's an internal secret: it can read a JSON of Shipping Ports from 
a file in some attached directory. Cool beans. It's like Clark Kent on the streets and Superman for the family.

The JSON file format must be as follows, with any size:
``` json
{
    "MUSIK": {
        "name": "Musik",
        "city": "Musik",
        "country": "Musicalis",
        "alias": [],
        "regions": [],
        "coordinates": [
            -1.00928,
            -11.88878
        ],
        "province": "",
        "timezone": "SharpOnTime",
        "unlocs": [
            "MUSIK"
        ]
    },
    "ANOTHER": {},
    ...
}
```

*Maybe, just maybe, it could be another service loading the server. We must talk to the man before changing this.*
*Maybe, just maybe, it should be an array of objects. Just thinking.*

# Deploy Structure
 - Two go modules - services - with Dockerfiles for building docker images
 - One docker-compose.yaml file to arrange them all in unison
 - One make file to bind them all and do everything in one dev run

**So, to rebuild after changes, just run `% make` (where the makefile is, the root).**

**No changes? Just pull the docker compose image and run.**


## Annotations

Most important:
+ mod database
+ review readmes

Bonus points:
+ better closure detection
+ multiple readers
+ add version to shippingports object
+ get many paged

Done
+ env process count
+ graceful shutdown
+ test rest handler 
+ test file block reader  
