FROM golang:1.19 AS builder
WORKDIR /go/src

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o app cmd/main.go

FROM alpine:latest

WORKDIR /go/bin

RUN apk add git

COPY --from=builder  /go/src/app app

CMD [ "/go/bin/app" ]