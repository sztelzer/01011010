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
COPY shippingportsclient/ ./shippingportsclient/
COPY shippingportsprotocol/ ./shippingportsprotocol/
COPY shippingportsserver/ ./shippingportsserver/

# build clientApp
WORKDIR /build/shippingportsclient
RUN go mod download
RUN go build -o /dist/shippingportsclientapp .

# build serverApp
WORKDIR /build/shippingportsserver
RUN go mod download
RUN go build -o /dist/shippingportsserverapp .


FROM scratch AS shippingportsclient

COPY --from=builder /dist/shippingportsclientapp /

EXPOSE 8080

ENTRYPOINT ["./shippingportsclientapp"]


FROM scratch AS shippingportsserver

COPY --from=builder /dist/shippingportsserverapp /

EXPOSE 50051

ENTRYPOINT ["./shippingportsserverapp"]