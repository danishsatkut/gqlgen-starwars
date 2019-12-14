sep = =============

travis: setup check_gqlgen tests;

build: clean setup check_gqlgen tests ; $(info $(M) $(sep) Building project $(sep))
	go version
	go env
	go build -o ./bin/graphql main.go

clean: ; $(info $(M) $(sep) Cleaning project $(sep))
	$(RM) ./bin/graphql

setup: ; $(info $(M) $(sep) Performing setup $(sep))
	go mod download
	go get github.com/99designs/gqlgen@v0.9.0

server: gqlgen ; $(info $(M) $(sep) Starting dev server $(sep))
	go run ./main.go

gqlgen: ; $(info $(M) $(sep) Generating graphql server code $(sep))
	gqlgen version
	gqlgen generate --config ./gqlgen.yml

check_gqlgen: ; $(info $(M) $(sep) Verify generated code $(sep))
	git diff --exit-code -- ./generated/exec.go
	git diff --exit-code -- ./generated/model.go

tests: gqlgen ; $(info $(M) $(sep) Running tests $(sep))
	go test ./...

.PHONY: build check_gqlgen clean gqlgen server setup tests travis
