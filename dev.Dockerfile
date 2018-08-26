FROM golang:1.10-alpine3.8

RUN apk update
RUN apk add git

RUN go get github.com/oxequa/realize

