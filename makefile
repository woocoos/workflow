version := $(shell /bin/date "+%Y-%m-%d %H:%M")
BUILD_NAME=workflow

build:
	go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o ./cmd/$(BUILD_NAME) ./cmd/main.go
mac:
	GOOS=darwin go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o ./cmd/$(BUILD_NAME)-darwin ./cmd/main.go
	$(if $(shell command -v upx), upx $(BUILD_NAME)-darwin)
win:
	GOOS=windows go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o ./cmd/$(BUILD_NAME).exe ./cmd/main.go
	$(if $(shell command -v upx), upx $(BUILD_NAME).exe)
linux:
	GOOS=linux go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o ./cmd/$(BUILD_NAME)-linux ./cmd/main.go
	$(if $(shell command -v upx), upx $(BUILD_NAME)-linux)
ent-new:
	GOWORK=off go run -mod=mod entgo.io/ent/cmd/ent --target graph/entgen/schema new $(NAME)
migration-init:
	GOWORK=off go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/versioned-migration ./ent/schema
migration:
	GOWORK=off go run -mod=mod ent/migrate/main.go -dsn="$(DSN)" -name=$(NAME)
migration-lint:
	atlas migrate lint --dev-url="$(DSN)" --dir="file://ent/migrate/migrations" --latest=$(LATEST)
migration-apply:
	atlas migrate apply --dev-url="$(DSN)" --dir="file://ent/migrate/migrations" --latest=$(LATEST)
test-db:
	GOWORK=off go run -mod=mod test/initdb.go

.PHONY: gen genent gengql
gen: genent gengql
genent:
	go run graph/entgen/entc.go
gengql:
	go run graph/gqlgen/gqlgen.go