# 01011010

Prerequisites on your machine: Docker (running)

Fast start: clone this repository to your local machine. Run make with the path and filename of your Shipping Ports json
file. Full example with your source json file being: `~/shipping/ports.json`

``` zsh
% git clone https://github.com/sztelzer/01011010
% cd 01011010
% make BIND_PATH=~/shipping JSON_FILENAME=ports.json
```

After building and finally running both services, you'll see something like this:

```
Attaching to shippingportsclient_1, shippingportsserver_1
shippingportsserver_1  | 2021/04/24 04:54:51 shippingPortsProtocolServer listening on port: :50051
shippingportsclient_1  | 2021/04/24 04:54:52 successfully loaded 1632 shippingPorts from file to server in 308.430959ms
```

We don't provide any json or ports data. The specifics of this json file will be discussed further below.

This make command will be your best friend. It also can run tests, builds and runs everything that is necessary during
development. You just need to have docker.

## Public REST Interface

**TLDR; There is only one public endpoint, by default in port 8080.**
Example to get Shipping Port with id `MUSIK` (this id is fictitious).

``` zsh
% curl -X GET localhost:8080/shippingports/MUSIK
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

In case you hit a malformed or non existent shipping port code, the response will be the status code `404 not found`. If
you try other http method, will get `invalid request method`
If something awful happens inside our service, it will respond some various `5xx` codes relative to the problem.

# About the services

This project is a collection of two services meant to be used together, but not limited to, organized around a common
protocol buffer and gRPC. It is built with some Go Modules responsible for each domain problem.

The modules are:

- ShippingPortsClient - the REST client interface. Also responsible for loading the first load of data from a file.
- ShippingPortsProtocol - the definition of data types and available methods between client and server.
- ShippingPortsServer - the database gateway server, connected with gRPC.
- MemDatabase - a simple memory database like a memstore. It can be changed quite easily, but serves the present
  purpose.

```
Network -> REST -> ShippingPortsClient -> gRPC -> ShippingPortsServer -> MemDatabase
```

Why so many? What are the benefits of this design? First, ShippingPortsClient have a simple interface for public use.
Second, the services are governed by the protocol buffers, so the changes in business logic must happen before, and as
soon as we change them, we will know what to change in the services.

Decoupling the database from other applications through the gRPC allow us to connect different clients no covered now,
in different programming languages, extremely fast, securing the access to data structures from other applications.

Finally, the database is really away from consumers, and we can better handle their schema relative to the protocol
buffers. The clients mostly likely don't need to know much about it, just about the protocol.

### About the file loading

The JSON file format must be as follows, but with any size:

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
    "ANOTHER": {...},
    ...
}
```

Check the protocol file definition `shippingPortsProtocol/shippingPorts.proto` to understand the protocol and
structures.

## Some controls

You can control the quantity of LOAD_SHIPPING_PORTS_PUTTERS running. They read from a channel that is fed with objects
read from the file, and send concurrent requests to the server. Too little (1) will probably result in poor performance.
Too much we risk overflowing the system with unneeded switches. It depends on many factors of the system running.
If the default value is not suitable, you can try to augment it.

You find these options in the `docker-compose.yml`

There you can also change the default ports of services.