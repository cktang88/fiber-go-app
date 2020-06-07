# fiber-go-app

- go modules
- gofiber/fiber
- https://github.com/cosmtrek/air for live-reload

Libs

- https://github.com/dominikh/go-tools for more static analysis
- use https://github.com/tarent/loginsrv if it had support for 2FA/TOTP...
  - probably just use https://github.com/firebase/firebase-admin-go

### dev

```bash
# to start hot reloading server
$ chmod +x air
$ ./air
```

### go modules

`go get _` to add new dependency
`go mod tidy` to prune deps in `go.mod`
`go mod vendor` copy installed deps to `/vendor`
