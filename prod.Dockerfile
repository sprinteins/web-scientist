FROM golang:1.12-alpine3.9 AS builder

RUN apk add git

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o web-scientist src/main.go

FROM alpine:3.8 AS runner

WORKDIR /app

COPY --from=builder /app/web-scientist /app

EXPOSE 7070

CMD ["go" , "run", "./web-scientist"]