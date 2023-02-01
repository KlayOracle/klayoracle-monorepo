
.PHONY: proto-installed
proto-installed:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi

.PHONY: proto-dp
proto-dp: proto-installed
	@protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        data-provider/protoadapter/*.proto

.PHONY: gomodtidy-dp
gomodtidy-dp: ## Run go mod tidy on all modules.
	cd ./data-provider && go mod tidy

.PHONY: test-dp
test-dp:
	@cd ./data-provider && go test -v ./...

.PHONY: build-adapter-id-gen
build-adapter-id-gen:
	@cd ./data-provider/utils/generateadptid && go build -o ../bin/generate_adapter_id generate_adapter_id.go

.PHONY: adapter-id-gen
adapter-id-gen: build-adapter-id-gen
	@if [ -z "$(ADAPTERS)" ]; then \
		echo "Specify valid adapter files. ADAPTERS="KLAY_USD.json WEMIX_USD.json""; \
	else \
		cd data-provider && ./bin/generate_adapter_id $(ADAPTERS); \
	fi
