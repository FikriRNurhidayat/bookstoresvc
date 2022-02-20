BUF_VERSION:=1.0.0-rc9

update:
	docker run -v $$(pwd):/src -w /src --rm bufbuild/buf:$(BUF_VERSION) mod update

generate:
	docker run -v $$(pwd):/src -w /src --rm bufbuild/buf:$(BUF_VERSION) generate

lint:
	docker run -v $$(pwd):/src -w /src --rm bufbuild/buf:$(BUF_VERSION) lint
