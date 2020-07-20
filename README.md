# fiber-go-app

- [go modules](https://blog.golang.org/migrating-to-go-modules) (stable since v1.14)
- [gofiber/fiber](https://github.com/gofiber/fiber)
- https://github.com/cosmtrek/air for live-reload

### Libs

- https://github.com/dominikh/go-tools for more static analysis
- use https://github.com/tarent/loginsrv if it had support for 2FA/TOTP...
  - probably just use https://github.com/firebase/firebase-admin-go

### Dev

```bash
# to start hot reloading server
$ chmod +x air
$ ./air
```

### go modules

`go get _` to add or upgrade dependency

`go mod tidy` to prune deps in `go.mod`

`go mod vendor` copy installed deps to `/vendor`
