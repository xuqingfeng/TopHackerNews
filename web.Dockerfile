FROM node:24 as builder

WORKDIR /src/web

COPY web/package* ./

RUN npm ci

COPY . /src/

RUN npm run build

FROM nginx:alpine

LABEL org.opencontainers.image.source https://github.com/xuqingfeng/TopHackerNews

RUN apk add --no-cache curl

COPY --from=builder /src/web/dist/ /usr/share/nginx/html/

HEALTHCHECK --interval=1m --timeout=5s --start-period=30s --start-interval=5s --retries=3 \
    CMD curl -fsS http://127.0.0.1/ || exit 1
