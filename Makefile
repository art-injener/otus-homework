DATABASE ?= mysql
CUR_USER ?= root
PASSWORD ?= mysql
HOST ?= localhost:3307
DB_NAME ?= dbase

.PHONY: migrate_up
migrate_up:
	docker image build -t custom_migrate ./db
	docker run --network host custom_migrate -path=/migrations/ -database "${DATABASE}://${CUR_USER}:${PASSWORD}@tcp(${HOST})/${DB_NAME}" up

.PHONY: build
build:
	go build -o ./bin/server ./cmd/webserver/main.go

.PHONY: run
run:
	go run cmd/webserver/main.go


.PHONY: install-lint
install-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.42.1

.PHONY: lint
lint:
	golangci-lint run

.PHONY: docker-build
docker-build:
	docker build  -t webservice .

.PHONY: docker-run
docker-run:
	docker run  --network=host --restart=always -d webservice