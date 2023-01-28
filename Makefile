
.PHONY: proto-installed
proto-installed:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi

.PHONY:proto-dp
proto-dp: proto-installed
	@protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        data-provider/protoadapter/*.proto

.PHONY: gomodtidy-dp
gomodtidy-dp: ## Run go mod tidy on all modules.
	cd ./data-provider && go mod tidy

.PHONY:test-dp
test-dp:
	@cd ./data-provider && go test -v ./...