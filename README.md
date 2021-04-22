# 01011010

## Introduction and Public REST Interface

**TLDR; There is only one public endpoint: `GET /shippingports/:portId`**

You get it. Now, for the details.

ShippingPorts is a simple REST service from where you can ONLY retrieve information about shipping ports around the
world by passing the unique identifier of the shipping port to the retrieval endpoint (that is the only one for now). 
For example, if running on a local machine, you can retrieve information about the shipping port ZWUTA hitting the 
service address: `http://localhost:8080/shippingports/ZWUTA` which yelds the following json encoded string on the 
response body and status code `200`.

``` json
{
    "id": "ZWUTA",
    "name": "Mutare",
    "city": "Mutare",
    "country": "Zimbabwe",
    "alias": [],
    "regions": [],
    "coordinates": [
        32.650351,
        -18.9757714
    ],
    "province": "Manicaland",
    "timezone": "Africa/Harare",
    "unlocs": [
        "ZWUTA"
    ]
}
```

In case you hit a malformed or no existent shipping port code, the response will be the status code `404` `not found`.
If something awful happens inside our service, it will respond some various `5xx` codes relative to the problem.


## Running the service

**TLDR; pull the docker compose image and run.**

This project is a collection of two services meant to be used together, organized around common RPC buffers, and
interconnected through gRPC. It is built spanning two Go Modules, but don't frail. It's quite simple.

The services are: 
 - ShippingPorts - the 'client service', REST interface, the gateway, conqueror of worlds.
 - ShippingPortsServer - the database gateway service, connect through gRPC. It makes the interface incredibly easy to
   use thanks to the protocol buffers.
 - Database - the database is, it is anything we may need.

So: Public Internet/Localhost Network -> REST -> ShippingPorts -> gRPC -> ShippingPortsServer -> Database

Why so many? What are the benefits of this design? First, ShippingPorts have a simple interface for public use.
Second, the services are governed by the protocol buffers, so the changes in business logic must happen
before in them. In Go, as soon as we change them, if changes need to happen on the services, we will know early.

Decoupling the database from other applications through the gRPC allow us to connect as many different services
we think about, in almost any different languages, extremely fast, securing the access to data from other applications. 

Also, if we are a hit and need to spawn more servers, so be it.

Finally, the database is really away from consumers, and we can better handle their schema relative to the RPC protocol.
 
### One last fact about ShippingPorts service:

Actually, ShippingPorts service has one more super power, but it's a secret: it can read a json of Shipping Ports from 
a file in some attached directory. Cool beans. It's like Clark Kent on the streets and Superman for the family.

Maybe, just maybe, it could be another service. We must talk to the man before changing this.

# The bones
 - Two go modules - with Dockerfile for building their images
 - One docker-compose.yaml file to arrange them all in unison
 - One make file to bind them all and do everything in one dev run

**So, to rebuild after changes, just run `% make` (in the root dir of this repo).**

No changes? Just pull the docker compose image and run it.

# Useful Commands

## Generate gRPC package from proto definition
This will generate/rewrite the go package that implements gRPC used in the applications. 
It's the core where everything begins.
```
protoc --go_out=. --go_opt=paths=source_relative \                                      
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    gRPCShippingPorts/shippingPorts.proto
```
