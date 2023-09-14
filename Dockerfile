FROM golang:1.21 as builder

WORKDIR /src

COPY . .

RUN go mod download && \
    go build -o thn-api server.go

FROM busybox:stable

WORKDIR /app

COPY --from=builder /src/thn-api /app/thn-api

EXPOSE 8080

CMD ["./thn-api"]
