NETWORK=$(shell docker network list | grep "klayoracle")
WD=$(shell pwd)
HOST_IP := "0.0.0.0:50051"
NODE_PORT := "50051"
DP_PORT := "50002"
NODE_IMAGE=klayoracle-node:v1.0.2-beta
DP_IMAGE=klayoracle-dp:v1.0.2-beta
SHELL := /bin/bash


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

.PHONY: cockroach-network
cockroach-network:
	@if [ -z "$(NETWORK)" ];	then	\
    	 docker network create klayoracle;		\
    else \
    	echo "$(NETWORK)";	\
    fi

.PHONY: cockroach-1
cockroach-1:
	@docker run -d \
    --name=roach1 \
    --hostname=roach1 \
    --net=klayoracle \
    -p 26257:26257 \
    -p 8081:8080 \
    -v "./setup-guide/volumes/cockroachdb/roach1:/cockroach/cockroach-data" \
    -v "./setup-guide/volumes/cockroachdb/migration:/cockroach/migration" \
    cockroachdb/cockroach:v23.1.8 start \
      --advertise-addr=roach1:26357 \
      --http-addr=roach1:8080 \
      --listen-addr=roach1:26357 \
      --sql-addr=roach1:26257 \
      --insecure \
      --join=roach1:26357,roach2:26357,roach3:26357

.PHONY: cockroach-2
cockroach-2:
	@docker run -d \
       --name=roach2 \
       --hostname=roach2 \
       --net=klayoracle \
       -p 26258:26258 \
       -p 8082:8081 \
       -v "./setup-guide/volumes/cockroachdb/roach2:/cockroach/cockroach-data" \
       cockroachdb/cockroach:v23.1.8 start \
         --advertise-addr=roach2:26357 \
         --http-addr=roach2:8081 \
         --listen-addr=roach2:26357 \
         --sql-addr=roach2:26258 \
         --insecure \
         --join=roach1:26357,roach2:26357,roach3:26357

.PHONY: cockroach-3
cockroach-3:
	@docker run -d \
       --name=roach3 \
       --hostname=roach3 \
       --net=klayoracle \
       -p 26259:26259 \
       -p 8083:8082 \
       -v "./setup-guide/volumes/cockroachdb/roach3:/cockroach/cockroach-data" \
       cockroachdb/cockroach:v23.1.8 start \
         --advertise-addr=roach3:26357 \
         --http-addr=roach3:8082 \
         --listen-addr=roach3:26357 \
         --sql-addr=roach3:26259 \
         --insecure \
         --join=roach1:26357,roach2:26357,roach3:26357

.PHONY: sleep-60seconds
sleep-60seconds:
	@echo "Preparing DB..."; \
    sleep 10; \
    echo "Finalizing DB..."; \

.PHONY: prepare-db-volume
prepare-db-volume:
	@for var in $(shell rm -r ./setup-guide/volumes/cockroachdb/roach{1,2,3}/*); do "roaching..." ; done;

.PHONY: cockroach-init
cockroach-init:
	@docker exec -it roach1 ./cockroach --host=roach1:26357 init --insecure

.PHONY: cockroach-migrate
cockroach-migrate:
	@cockroach sql --file=./setup-guide/volumes/cockroachdb/migration/dbinit.sql --insecure --host=0.0.0.0:26258
#	@docker exec -it roach1 ./cockroach sql --file=./migration/dbinit.sql --host=roach1:26258 --insecure

.PHONY: cockroach-cluster
cockroach-cluster: prepare-db-volume cockroach-network cockroach-1 cockroach-2 cockroach-3 sleep-60seconds cockroach-init cockroach-migrate

.PHONEY: export-var
export-var:
	@for var in $(shell cat ${TARGET}); do export "$${var}"; done;

#Deprecated, use `make cockroach-cluster` instead
.PHONEY: devnet-tables
devnet-tables:
	@for var in $(shell cat setup-guide/volumes/nodes/nd{1,2,3,4,5}/db.var); do export "$${var}" && make node-tables ; done;

#Deprecated
.PHONY: devnet-cluster
devnet-cluster:
	@make docker-network; NODE_IMAGE=${NODE_IMAGE} NODE_PORT=${NODE_PORT} make node-image; DP_PORT=${DP_PORT} DP_IMAGE=${DP_IMAGE}  make dp-image
	@make devnet-tables
	@docker compose up --detach

.PHONY: oracle-cluster
oracle-cluster:
	@make docker-network; NODE_IMAGE=${NODE_IMAGE} NODE_PORT=${NODE_PORT} make node-image; DP_PORT=${DP_PORT} DP_IMAGE=${DP_IMAGE}  make dp-image
	@make cockroach-cluster
	@docker compose --file compose.yml up --detach

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
    	cd node/utils && go run compile_contract/compile_contract.go;	\
    fi