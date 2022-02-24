BUF_VERSION:=1.0.0-rc9

# Database configuration
export DATABASE_NAME=bookstoresvc_development
export DATABASE_USER=root
export DATABASE_PASSWORD=
export DATABASE_HOST=localhost
export DATABASE_PORT=5432
export DATABASE_SSL_MODE=disable
export DATABASE_URL=postgres://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=${DATABASE_SSL_MODE}

setup:
	init
	install
	db-create
	db-migrate
	db-seed

init:
	go mod vendor
	docker run -v $$(pwd):/src -w /src --rm bufbuild/buf:$(BUF_VERSION) mod update

build:
	mkdir out || echo "out directory already exists"
	go build -o out/bookstoresvc main.go

develop:
	go run main.go

clean:
	rm -rf ./proto/*.go

generate:
	docker run -v ${PWD}:/src -w /src --rm bufbuild/buf:$(BUF_VERSION) generate

lint:
	docker run -v ${PWD}:/src -w /src --rm bufbuild/buf:$(BUF_VERSION) lint
	docker run -v ${PWD}:/src -w /src --rm bufbuild/buf:$(BUF_VERSION) breaking --against 'https://github.com/fikrirnurhidayat/bookstoresvc.git#branch=main'

db-migrate:
	docker run -v ${PWD}/db/migrations:/migrations \
		--rm -it --network host migrate/migrate \
		--path=/migrations/ \
		--database ${DATABASE_URL} up

db-seed:
	psql -a -f ./db/seeds.sql ${DATABASE_NAME}

db-rollback:
	docker run -v ${PWD}/db/migrations:/migrations \
		--rm -it --network host migrate/migrate \
		--path=/migrations/ \
		--database ${DATABASE_URL} down

db-create:
	createdb ${DATABASE_NAME}

db-drop:
	dropdb ${DATABASE_NAME}
