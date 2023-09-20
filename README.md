## TopHackerNews

List top stories on HackerNews via it's open [API](https://github.com/HackerNews/API) .

### How to run locally

```
# start api
go mod download
APP_ENV=local go run server.go

# start web
cd web
npm install
npm run dev

# open website
open http://localhost:3000
```

### Tech Stack

- golang
- svelte
- graphql
