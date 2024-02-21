# gRPC_auth
gRPC SSO service

## Running

```bash
go mod download 
go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./migrations
go run cmd/sso/main.go --config=./config/local.yaml
```

## Testing

```bash
go run ./cmd/migrator/main.go --storage-path=./storage/sso.db --migrations-path=./tests/migrations --migrations-table=migrations_test
go run cmd/sso/main.go --config=./config/local.yaml
go test -v ./tests/auth_register_login_test.go
```

[Protobuf contract](https://github.com/sotchenkov/protos)

[Reference](https://github.com/GolangLessons/sso)
