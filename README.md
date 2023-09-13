# Go Calendar Tutorial

_Backend App for a full stack project written entirely with Go._

## Used tools 🛠️
- For the backend, we use a **PostgreSQL** database, and create all the methods to access and modify securely instances with **Go**.

- The frontend is also written with **Go**, using the package `html/template` of **Go Templates** engine, which is a common choice for server-side rendering in web applications to generate dynamic HTML content on the server and send it to the client's browser.

## Set it up ⚙️
- Postgres development database:
```bash
docker-compose up dev_postgres
```


## Tests ⌨️
- Unit tests:
```
go test ./internal/dao/ -tags unit_test
```
```
go test ./internal/dao/generator/ -tags unit_test -v
```
```
go test ./cmd/mocking/config/ -tags unit_test,mocking_app -v
```

- Integration tests:
```
go test ./internal/dao/integrationtest/ -tags integration_test -v
```
