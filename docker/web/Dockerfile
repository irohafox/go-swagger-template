FROM golang:1.16.4-alpine
RUN apk update && apk add git

RUN GO111MODULE=off
RUN go get -u github.com/cosmtrek/air
RUN go get -u github.com/go-swagger/go-swagger/cmd/swagger
ENV GO111MODULE="on"

RUN mkdir /go/src/golang-rest-api-template
WORKDIR /go/src/golang-rest-api-template
ADD . /go/src/golang-rest-api-template
