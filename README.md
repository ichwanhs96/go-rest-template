## Setting up environment
This docker compose file used to run server and postgresql db in our local machine using docker
- run command "docker-compose -f docker-compose.yml up" to run everything
- run command "docker-compose up -d postgres" to run db only

## How to run
Do `go run main.go`, server run and listen at port 8090

## Routes
| Path | Method | Header | Description |
| ---- | ------ | ------ | ----------- |
| /hello | GET  |        | hello world |