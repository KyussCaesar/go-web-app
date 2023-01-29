.PHONY: setup clean client

PROJECT_GO_VERISON ?= 1.19.5

setup: # ensure your local dev environment is set up
	@chmod +x setup.sh
	./setup.sh

app:
	cd src/app && go generate && cd .. && go build -o ../

client: openapi.yml # build the local client
	npx @openapitools/openapi-generator-cli generate --generator-key go

clean:
