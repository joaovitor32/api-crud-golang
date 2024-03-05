# Api Crud Golang

### Run

```bash
   go run cmd/server/main.go
```

### Run with Docker

```bash
   docker network create api-crud-golang
   docker volume create --name=crud-mysql
   docker-compose --project-directory ./ -f ./build/docker-compose.yml up 
```

### Rollback migration

```bash
   docker run --network host migrate -path /database -database "mysql://root:secret@tcp(db:3306)/crud" -verbose down -all
```
