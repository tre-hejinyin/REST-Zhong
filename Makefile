.PHONY: build run lint up down db-up db-in mig-cr clean postgres createdb dropdb migrate migrateup migratedown \
coverage test install

install:
	go get github.com/cosmtrek/air@v1.15.1
	go get google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go get -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

up: db-up sl migrate
	DOCKER_BUILDKIT=1 docker-compose up --build server

air: export DB_HOST=127.0.0.1
air: export DB_PORT=5432
air: export DB_USER=test
air: export DB_PASSWORD=test
air: export DB_NAME=template
air: export TZ=Asia/Tokyo
air: export REDIS_HOST=127.0.0.1:6379
air: export TOKEN_SECURITY_KEY=3b93f2d6700480b63c71f8d1f802e878
air:
	air

down:
	docker-compose down

sl:
	sleep 3

run: up

build:
	go build

lint:
	golangci-lint run

clean:
	docker system prune --volumes -f

db-up:
	DOCKER_BUILDKIT=1 docker-compose up --build -d db

db-in:
	docker-compose exec db bash

mig-cr:
	migrate create -ext sql -dir db/migrations -seq ${name}

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.1

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrate:
	migrate -path db/migrations -database 'postgres://test:test@127.0.0.1:5432/template?sslmode=disable' -verbose up

migrateup:
	migrate -path db/migrations -database ${POSTGRESQL_URL} -verbose up

migratedown:
	migrate -path db/migrations -database ${POSTGRESQL_URL} -verbose down

migrate-create:
	migrate create -ext sql -dir migrations/example1 -seq create_users_table

coverage:
	go tool cover -html=c.out -o coverage.html

test:
	go test -covermode=count -coverprofile=c.out
	go tool cover -html=c.out -o coverage.html
