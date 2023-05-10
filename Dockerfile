# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.20-alpine3.16 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build -o simpleMongoAPI .\cmd\app\main.go

##
## Deploy
##

FROM alpine:3.14.0
WORKDIR /
COPY --from=build /simpleMongoAPI /simpleMongoAPI
ENTRYPOINT ["/simpleMongoAPI"]