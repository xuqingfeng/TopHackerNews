FROM golang:1.21 as builder

WORKDIR /src

COPY . .

RUN go mod download && \
    go build -o thn-api server.go

FROM alpine

WORKDIR /app

RUN apk update && apk add ca-certificates

COPY --from=builder /src/thn-api /app/thn-api

EXPOSE 8080

CMD ["./thn-api"]
