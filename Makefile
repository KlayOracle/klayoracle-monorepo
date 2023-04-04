NETWORK=$(shell docker network list | grep "klayoracle")

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

.PHONY: proto-node
proto-node: proto-installed
	@protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        node/protonode/*.proto


.PHONY: gomodtidy-dp
gomodtidy-dp: ## Run go mod tidy on all modules.
	cd ./data-provider && go mod tidy

.PHONY: test-dp
test-dp:
	@cd ./data-provider && go test -v ./...

.PHONY: test-node
test-node:
	@cd ./node && go test -v ./...


.PHONY: build-adapter-id-gen
build-adapter-id-gen:
	@cd ./data-provider/utils/generateadptid && go build -o ../bin/generate_adapter_id generate_adapter_id.go

.PHONY: adapter-id-gen
adapter-id-gen: build-adapter-id-gen
	cd data-provider/utils &&  go run generateadptid/generate_adapter_id.go ${ADAPTERS}

#	@if [ -z "$(ADAPTERS)" ]; then \
#		echo "Specify valid adapter files. ADAPTERS="KLAY_USD.json WEMIX_USD.json""; \
#	else \
#		cd data-provider && ./bin/generate_adapter_id $(ADAPTERS); \
#	fi

.PHONY: node-server
node-server:
	@cd ./node && GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info go run main.go

.PHONY: dp-client
dp-client:
	@cd ./data-provider && GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info go run main.go

.PHONY: node-server-nolog
node-server-nolog:
	@cd ./node && go run main.go

.PHONY: dp-client-nolog
dp-client-nolog:
	@cd ./data-provider && go run main.go

.PHONY: build-node
build-node:
	@cd ./node && go build -o bin/node


.PHONY: node-image
node-image:
	@docker build -t klayoracle-node:dev -f node.Dockerfile . --build-arg PORT=${PORT}

.PHONY: dp-image
dp-image:
	@docker build -t klayoracle-dp:dev -f dp.Dockerfile . --build-arg PORT=${PORT}

.PHONY: docker-network
docker-network:
	@if [ -z "$(NETWORK)" ];	then	\
    	 docker network create klayoracle;		\
    else \
    	echo "$(NETWORK)";	\
    fi

.PHONY: net-cluster
net-cluster:
	@make docker-network
	docker run -d --net klayoracle --name bootstrap_node1 --env-file node/.env klayoracle-node:dev
	docker run -d --net klayoracle --name bootstrap_dp1 --env-file data-provider/.env --env HOST_IP=bootstrap_dp1:50002 klayoracle-dp:dev
	docker run -d --net klayoracle --name bootstrap_dp2 --env-file data-provider/.env --env HOST_IP=bootstrap_dp2:50002 klayoracle-dp:dev
	docker run -d --net klayoracle --name bootstrap_dp3 --env-file data-provider/.env --env HOST_IP=bootstrap_dp3:50002 klayoracle-dp:dev
	docker run -d --net klayoracle --name common_dp --env-file data-provider/.env --env HOST_IP=common_dp:50002 klayoracle-dp:dev