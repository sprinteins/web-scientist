FROM golang:1.12-alpine3.9


RUN apk update
RUN apk add git

RUN go get github.com/oxequa/realize

