FROM golang:1.12-alpine3.9 AS builder

RUN apk add git

WORKDIR /app
COPY . .

WORKDIR /app/src
RUN go mod download
RUN go build -o web-scientist main.go

FROM alpine:3.9 AS runner

WORKDIR /app

COPY --from=builder /app/src/web-scientist /app

CMD ["./web-scientist"]
