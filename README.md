# fiber-go-app

- [go modules](https://blog.golang.org/migrating-to-go-modules) (stable since v1.14)
- [gofiber/fiber](https://github.com/gofiber/fiber)
- https://github.com/cosmtrek/air for live-reload

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

### Libs being considered

- https://github.com/dominikh/go-tools for more static analysis
- use https://github.com/tarent/loginsrv if it had support for 2FA/TOTP...
  - probably just use https://github.com/firebase/firebase-admin-go

### db libs

there are no great ORM libs in go:

> ORM requires a very framework oriented development style, which Go is not particularly geared toward lacking both generics and dynamic class loader. It doesn't fit particularly well with Go's focus on simplicity and performance, either. - https://www.reddit.com/r/golang/comments/b9h8xo/the_state_of_orms_in_2019/

- ~~[jmoiron/sqlx](https://github.com/jmoiron/sqlx) - raw sql strings~~
- [dbr](https://github.com/gocraft/dbr) - inspired by `sqlx` and `Squirrel`

* gorm - heavy, raw sql
* xorm - similar to gorm
* https://github.com/go-reform/reform - unclear docs on how to do joins/order/etc.
* upper/db - ?

* on the other hand could use a code generator like https://github.com/volatiletech/sqlboiler or Prisma (go client)

SQL migrations - https://github.com/rubenv/sql-migrate
