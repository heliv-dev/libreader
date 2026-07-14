export GOPRIVATE=dev.azure.com
makeFileDir := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: tools
tools:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

.PHONY: sqlc
sqlc:
	docker run -it --entrypoint /bin/sh --rm -v  .:/src -w /src/sql/queries unibeautify/sqlformat /src/sql/queries/_format.sh
	docker run --rm -v ${makeFileDir}:/src -w /src kjconroy/sqlc generate -f ./sql/sqlcdbx.yaml


.PHONY: format
format:
	go install golang.org/x/tools/cmd/goimports@latest
	goimports -w -l $(find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: migrate
migrate:
	migrate -source file://sql/migrations -database postgres://youruser:yourpassword@127.0.0.1:5432/library?sslmode=disable up
