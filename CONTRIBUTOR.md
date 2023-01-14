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