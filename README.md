### [BACKLOG](BACKLOG.md)


### How to run development

1. Add development.env
2. Run compose
```
docker compose up -d
```
3. Attach to container
```
docker exec -i -t go-echo-dev bash
```
4. Inside container run
```
air
```
4.1. Build manually
```
go build  -o ./tmp/main -buildvcs=false cmd/go-echo/main.go
```
4.2. Run manually
```
go run cmd/go-echo/main.go
```