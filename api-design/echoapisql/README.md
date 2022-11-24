```
docker compose up -d
go run server.go
go test -v -tags=integration
docker compose down
```
