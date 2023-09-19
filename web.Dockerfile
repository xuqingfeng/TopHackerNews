FROM node:18 as builder

WORKDIR /src/web

COPY web/package* ./

RUN npm ci

COPY . /src

RUN npm run build

FROM nginx:alpine

COPY --from=builder /src/web/dist/ /usr/share/nginx/html/
