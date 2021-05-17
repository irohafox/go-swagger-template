FROM golang:1.16.4-alpine
RUN apk update && apk add git

RUN mkdir /go/src/golang-rest-api-template
WORKDIR /go/src/golang-rest-api-template
ADD . /go/src/golang-rest-api-template
