## Branch Naming Convention

### Regular Branches

- main: Release branch
- staging: smoke and soak test before merging to main
- development: All feature branch merges in here, then to staging

### Temporary Branches

- Feature Branches: feature-<ISSUE_NUMBER>-node-aggregator-mechanism
- Bug Branches: bug-<ISSUE_NUMBER>-node-networking-failing-on-amd64
- WIP Branches: wip-<ISSUE_NUMBER>-protocol-scaffold-setup
- Hot Fix Branches: hot-fix-<ISSUE_NUMBER>-race-condition-on-node

## Commits Message

- Must start with any of the follow:
  - wip: 
  - feature:
  - bug:
  - hot-fix:
- Changes of more than one file must include a commit description


### Pre Commit Hook

NB: Make sure to run `snyk auth` to authenticate CLI
```shell
#!/bin/sh

yarn --cwd oracle-contract prettier:write
snyk test --severity-threshold=high --all-projects --detection-depth=4
cd data-provider && go fmt .
golangci-lint run . -v -E goimports --fix
cd ../node && go fmt .
golangci-lint run . -v -E goimports --fix
cd ..
$HOME/go/bin/goimports -w data-provider/*
$HOME/go/bin/goimports -w node/*
```