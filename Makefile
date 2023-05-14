NETWORK=$(shell docker network list | grep "klayoracle")
WD=$(shell pwd)
HOST_IP := "0.0.0.0:50051"
NODE_PORT := "50051"
DP_PORT := "50002"

.PHONY: proto-installed
proto-installed:
	@if ! which protoc > /dev/null; then \
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


.PHONY: gomodtidy
gomodtidy: ## Run go mod tidy on all modules.
	@cd data-provider && go mod tidy && cd - && cd node && go mod tidy

.PHONY: test-dp
test-dp:
	@cd ./data-provider && go test -v ./...

.PHONY: test-node
test-node:
	@cd ./node && go test -v ./...

.PHONY: pre-commit
pre-commit:
	@./pre-commit.sh

.PHONY: build-adapter-id-gen
build-adapter-id-gen:
	@cd ./data-provider/utils/generateadptid && go build -o ../bin/generate_adapter_id generate_adapter_id.go

.PHONY: adapter-id-gen
adapter-id-gen:
	cd data-provider/utils &&  go run generateadptid/generate_adapter_id.go ${ADAPTERS}

.PHONY: adapter-dry-run
adapter-dry-run:
	@cd data-provider/utils &&  go run dryrun/dryrun.go ${ADAPTERS}

.PHONY: node-server
node-server:
	@cd ./node && GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info HOST_IP=$(HOST_IP) environment=local go run main.go

.PHONY: dp-client
dp-client:
	@cd ./data-provider && GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info HOST_IP=$(HOST_IP) go run main.go

.PHONY: node-server-nolog
node-server-nolog:
	@cd ./node && HOST_IP=$(HOST_IP) environment=local go run main.go

.PHONY: dp-client-nolog
dp-client-nolog:
	@cd ./data-provider && HOST_IP=$(HOST_IP) go run main.go

.PHONY: build-node
build-node:
	@cd ./node && go build -o bin/node

.PHONY: node-image
node-image:
	@docker build -t ${NODE_IMAGE} -f node.Dockerfile . --build-arg PORT=${NODE_PORT}

.PHONY: dp-image
dp-image:
	@docker build -t ${DP_IMAGE} -f dp.Dockerfile . --build-arg PORT=${DP_PORT}

.PHONY: docker-network
docker-network:
	@if [ -z "$(NETWORK)" ];	then	\
    	 docker network create klayoracle;		\
    else \
    	echo "$(NETWORK)";	\
    fi

.PHONEY: export-var
export-var:
	@for var in $(shell cat ${TARGET}); do export "$${var}"; done;

.PHONEY: devnet-tables
devnet-tables:
	@for var in $(shell cat setup-guide/volumes/nodes/nd{1,2,3,4}/db.var); do export "$${var}" && make node-tables ; done;

.PHONY: devnet-cluster
devnet-cluster:
	@make docker-network; TARGET=images.var make export-var; make node-image; make dp-image
	@make devnet-tables
	@docker compose up --detach

.PHONY: node-tables
node-tables:
	@cockroach sql --url $(COCKROACH_DNS_URL) --file ${WD}/node/dbinit.sql

.PHONY: solidity-bindings
solidity-bindings:
	@if ! which solc > /dev/null; then \
    		echo "error: solc not installed" >&2; \
    		exit 1; \
     else \
       solc --abi oracle-contract/contracts/*.sol -o node/contracts/build --bin --overwrite @openzeppelin=oracle-contract/node_modules/@openzeppelin ; \
       cd node/utils && go run compile_contract/compile_contract.go ; \
     fi