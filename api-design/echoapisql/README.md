

```
docker compose up -d
go run server.go
go test -v
docker compose down
```

```
go test -v -tags=integration
```