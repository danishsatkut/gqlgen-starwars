sep = =============

travis: setup tests;

setup: ; $(info $(M) $(sep) Performing setup $(sep))
	go mod download

gqlgen: ; $(info $(M) $(sep) Generating graphql server code $(sep))
	gqlgen version
	gqlgen generate --config ./gqlgen.yml

tests: gqlgen ; $(info $(M) $(sep) Running tests $(sep))
	go test ./...

.PHONY: gqlgen travis tests setup
