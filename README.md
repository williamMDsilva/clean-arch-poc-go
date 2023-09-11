# DATABASE

```bash
migrate create -ext sql -dir database/migrations -seq create_product_table
```

# local dependencies

https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

```bash
brew install golang-migrate
```

# project dependencies

https://github.com/jackc/pgx/wiki/Getting-started-with-pgx
https://github.com/gorilla/mux
https://github.com/booscaaa/go-paginate
https://github.com/spf13/viper
https://github.com/stretchr/testify
https://github.com/pashagolub/pgxmock
https://github.com/golang-migrate/migrate

**_Banco de dados_**: PostgresSQL

**_Libs no go_**

- Pgx
- Mux
- Go-paginate
- Viper
- Testify
- Pgx Mock
- Migrate

# docs

http://localhost:3000/swagger/

```bash
swag init -d adapter/http --parseDependency --parseInternal --parseDepth 2 -o adapter/http/docs
```

# Run

```bash
go run adapter/http/main.go
```

# tests

### Test verbose

```bash
go test ./... -v
```

### Cover profile

```bash
go test -coverprofile cover.out ./...
go tool cover -html=cover.out -o cover.html
```
