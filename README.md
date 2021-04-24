# 01011010

## Introduction and Public REST Interface

**TLDR; There is only one public endpoint: `GET /shippingports/:shippingPortId`**

ShippingPorts is a really simple REST service from where you can ONLY retrieve information about shipping ports around the
world by passing the unique identifier of the shipping port to the retrieval endpoint (that is the only one for now). 

For example, if running on a local machine, you can retrieve information about the shipping port MUSIK hitting the 
service address: `http://localhost:8080/shippingports/MUSIK` which yelds the following json encoded string on the 
response body and status code `200`.

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
    }
}
```

In case you hit a malformed or no existent shipping port code, the response will be the status code `404` `not found`.
If something awful happens inside our service, it will respond some various `5xx` codes relative to the problem.


## Running the service

**TLDR; pull the docker compose image and run.**

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
+ mod database
