FROM node:18 as builder

WORKDIR /src/web

COPY web/package* ./

RUN npm ci

COPY . /src/

RUN npm run build

FROM nginx:alpine

LABEL org.opencontainers.image.source https://github.com/xuqingfeng/TopHackerNews

HEALTHCHECK --interval=2m --timeout=5s --start-period=5s \
    CMD curl -f http://127.0.0.1/ || exit 1

COPY --from=builder /src/web/dist/ /usr/share/nginx/html/
