.PHONY: all
all: build build-frontend fmt vet lint test

ALL_PACKAGES=$(shell go list ./... | grep -v -e "vendor" -e "frontend")
UNIT_TEST_PACKAGES=$(shell  go list ./... | grep -v -e "vendor" -e "frontend")
DB_NAME="shop_dev"
TEST_DB_NAME="shop_test"
TEST_DB_PORT=5432
DB_PORT=5432
APP_EXECUTABLE="out/shop"

setup:
	go get -u golang.org/x/lint/golint
	sudo npm install -g @angular/cli

compile:
	mkdir -p out/
	go build -o $(APP_EXECUTABLE)

build: compile fmt vet lint

build-frontend:
	cd ./frontend;npm install; ng build --base-href /homepage/;

install:
	go install ./...

fmt:
	go fmt $(ALL_PACKAGES)

vet:
	go vet $(ALL_PACKAGES)

lint:
	@for p in $(UNIT_TEST_PACKAGES); do \
		echo "==> Linting $$p"; \
		golint $$p | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; } \
	done


test: testdb.drop testdb.create testdb.migrate
	ENVIRONMENT=test go test $(UNIT_TEST_PACKAGES) -race

db.setup: db.create db.migrate

db.create:
	createdb -p $(DB_PORT) -Opostgres -Eutf8 $(DB_NAME)

db.migrate:
	$(APP_EXECUTABLE) migrate

db.rollback:
	$(APP_EXECUTABLE) rollback

db.drop:
	dropdb -p $(DB_PORT) --if-exists -Upostgres $(DB_NAME)

db.reset: db.drop db.create db.migrate


testdb.create:
	createdb  -p $(TEST_DB_PORT) -Opostgres -Eutf8 $(TEST_DB_NAME)

testdb.migrate:
	ENVIRONMENT=test $(APP_EXECUTABLE) migrate

testdb.drop:
	dropdb -p $(TEST_DB_PORT) --if-exists -Upostgres $(TEST_DB_NAME)

testdb.reset: testdb.drop testdb.create testdb.migrate

test-cover-html:
	@echo "mode: count" > coverage-all.out
	$(foreach pkg, $(ALL_PACKAGES),\
	go test -coverprofile=coverage.out -covermode=count $(pkg);\
	tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html=coverage-all.out -o out/coverage.html

copy-config:
	cp application.yml.sample application.yml

copy-config-ci:
	cp ci.yml.sample application.yml
