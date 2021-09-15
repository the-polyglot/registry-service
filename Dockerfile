# build executable binary
FROM golang:alpine as builder

COPY go.mod go.sum  /go/src/github.com/tkl/registry-service/

WORKDIR /go/src/github.com/tkl/registry-service

# fetch dependencies
RUN go mod download
RUN go mod verify
# RUN go mod tidy

COPY . /go/src/github.com/tkl/registry-service

ENV CGO_ENABLED=0 
RUN GOOS=darwin GOARCH=amd64 go build -ldflags="-g -02" -o /usr/bin/registry-service

FROM scratch

RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

COPY --from=builder /go/src/github.com/tkl/registry-service/build/registry-service /usr/bin/registry-service

EXPOSE 8080

ENTRYPOINT ["/usr/bin/registry-service"]

