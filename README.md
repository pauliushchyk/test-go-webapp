# test-go-webapp

## start up

1. install [dep](https://golang.github.io/dep/docs/introduction.html)
2. run `dep ensure`
3. restore database if it doesn't exist (use backup file)
4. change configuration `config/development.toml` or `config/production.toml`
5. run `go test -run ''` for checking tests
6. run `go run main.go` for running app
