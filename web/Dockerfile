FROM node:18 as builder

WORKDIR /src

COPY package* ./

RUN npm ci

COPY . .

RUN npm run build

FROM nginx:alpine

COPY --from=builder /src/dist/ /usr/share/nginx/html/
