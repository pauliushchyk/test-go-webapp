# test-go-webapp

## set up

### installation

1. install [dep](https://golang.github.io/dep/docs/introduction.html)
2. run `dep ensure`

### running app

1. create `Customers` database
2. change configuration `config/development.toml` or `config/production.toml`
3. run `go test -run ''` for checking tests
4. run `go run main.go` for running app

### restoring data

1. restore data by executing INSERT query from `backup` file

Also you can create database before running app and restore structure with data by running queries in `backup` file
