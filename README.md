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