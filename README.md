# gRPC_auth
gRPC SSO service

```bash
go mod download 
go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./migrations
go run cmd/sso/main.go --config=./config/local.yaml
```

[Protobuf contract](https://github.com/sotchenkov/protos)

[Reference](https://github.com/GolangLessons/sso)
