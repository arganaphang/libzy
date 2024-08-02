set dotenv-load
format:
	gofumpt -l -w .
	goimports-reviser -rm-unused -set-alias ./...
	golines -w -m 120 *.go

# build -> build application
build:
	go build -o main ./cmd

# run -> application
run:
	./main

# dev -> run build then run it
dev: 
	watchexec -r -c -e go -- just build run

# health -> Hit Health Check Endpoint
health:
	curl -s http://localhost:8000/healthz | jq

# migrate-create -> create migration
migrate-create NAME:
	migrate create -ext sql -dir ./database/migrations -seq {{NAME}}

# migrate-up -> up migration
migrate-up:
	migrate -path ./database/migrations -database "postgres://$DATABASE_USER:$DATABASE_PASS@$DATABASE_HOST:5432/$DATABASE_NAME?sslmode=disable" up

# sqlc-gen -> Generate Model from sql migrations
sqlc-gen:
	sqlc generate
