FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy all modules to /build (we need all as local dependencies)
COPY memdatabase/ ./memdatabase/
COPY shippingPortsClient/ ./shippingPortsClient/
COPY shippingPortsProtocol/ ./shippingPortsProtocol/
COPY shippingPortsServer/ ./shippingPortsServer/

# build clientApp
WORKDIR /build/shippingPortsClient
RUN go mod download
RUN go build -o /dist/shippingPortsClientApp .

# build serverApp
WORKDIR /build/shippingPortsServer
RUN go mod download
RUN go build -o /dist/shippingPortsServerApp .


FROM scratch AS shippingportsclient

COPY --from=builder /dist/shippingPortsClientApp /

EXPOSE 8080

ENTRYPOINT ["./shippingPortsClientApp"]


FROM scratch AS shippingportsserver

COPY --from=builder /dist/shippingPortsServerApp /

EXPOSE 50051

ENTRYPOINT ["./shippingPortsServerApp"]