FROM golang:1.21 as builder

WORKDIR /src

COPY . .

RUN go mod download && \
    go build -o thn-api server.go

FROM alpine:3.18

LABEL org.opencontainers.image.source https://github.com/xuqingfeng/TopHackerNews

WORKDIR /app

# https://stackoverflow.com/questions/66963068/docker-alpine-executable-binary-not-found-even-if-in-path
RUN apk update && apk add curl libc6-compat ca-certificates

COPY --from=builder /src/thn-api /app/thn-api

HEALTHCHECK --interval=2m --timeout=5s --start-period=5s \
    CMD curl -f http://127.0.0.1:8080/ || exit 1

EXPOSE 8080

CMD ["./thn-api"]
