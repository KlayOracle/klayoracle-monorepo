#!/bin/sh

WD="$(pwd)"
DP="$WD"/data-provider
NODE="$WD"/node
OC="$WD"/oracle-contract

cd "$OC" && yarn prettier:write
cd "$OC" && yarn vun:check

cd "$DP"
go fmt .
golangci-lint run . -v -E goimports --fix --color=always --config="$WD"/.golangci.yml
"$HOME"/go/bin/goimports -w .

cd "$NODE"
go fmt .
golangci-lint run . -v -E goimports --fix --color=always --config="$WD"/.golangci.yml
"$HOME"/go/bin/goimports -w .

cd "$WD"
snyk test --severity-threshold=high --all-projects --detection-depth=4