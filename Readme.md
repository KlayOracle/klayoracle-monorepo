![codeql](https://github.com/KlayOracle/klayoracle-monorepo/actions/workflows/codeql.yml/badge.svg)
![dependency](https://github.com/KlayOracle/klayoracle-monorepo/actions/workflows/dependency-review.yml/badge.svg)
![linter](https://github.com/KlayOracle/klayoracle-monorepo/actions/workflows/golang-ci.yml/badge.svg)
![unit testing](https://github.com/KlayOracle/klayoracle-monorepo/actions/workflows/gounit-test.yml/badge.svg)
![slither analysis](https://github.com/KlayOracle/klayoracle-monorepo/actions/workflows/slither.yml/badge.svg)
![snyk check](https://github.com/KlayOracle/klayoracle-monorepo/actions/workflows/snyk.yml/badge.svg)

## Official KlayOracle monorepo

KlayOracle is a decentralized oracle for the Klaytn blockchain. This monorepo contains all core packages that make up the KlayOracle protocol.

## 📚 Learn More

Docs
- [Developer documentation](https://klayoracle.gitbook.io/v1.0.0/)
- [Architecture Decision Records](/docs/arch)

Website: [klayoracle.com](https://klayoracle.com)

Blog: [klayoracle.medium.com](https://klayoracle.medium.com)

### 📦 Packages

- [data-provider](/data-provider): For Data Providers to serve data feeds to Data Subscribers.
- [node](/node): For Node runners to serve multiple Data Providers requests.
- [oracle-contract](/oracle-contract): Deployed by Data Providers for each data feed. 


### 📬 Contact

Feel free to email engineering@klayoracle.com

## Connect with KlayOracle

[Website](https://klayoracle.com) | [Dorahacks](https://community.dorahacks.io/t/klayoracle-an-open-source-oracle-framework-to-securely-integrate-off-chain-data-with-klaytn-smart-contracts/58) | [Twitter](https://twitter.com/klayoracle) | [Medium](https://klayoracle.medium.com/) | [Telegram](https://t.me/klayoracle) | [YouTube](https://www.youtube.com/@klayoracle)

## Contributors

- [Oluwafemi Alofe](https://www.linkedin.com/in/oluwafemialofe/)
- [Paul Oladimeji](https://www.linkedin.com/in/pauloladimeji/)

## Checklist

### Node Runner

- Generate a selfsigned certificate for each node. [see guide using open ssl](setup-guide/openssl).
- Use different authority or wildcard matching authority of each node you intend to run.
- In `node/config.yml`, set the `key` and the `pem` you used to generate the certificate.
- Give the certificate pem file to data providers for authenticating with your Node. You can also run your own data providers.
- Update your organization details in `node/config.yml`. Website must match authority used to sign certificate, an example will be if authority is `*.origineum.com`, website can be `node-1.origineum.com`. Otherwise, node service will fail to start.
- In `.env`, `HOST_IP` is the dns for reaching your node. If you're running a docker container, you can override it and other `env` variable by passing from terminal.
- In `.env` `PRIVATE_KEY` is your Node signer with enough Klay tokens to pay for request.
- In `.env` `COCKROACH_DNS_URL` is your full connection string to [https://cockroachlabs.cloud/](https://cockroachlabs.cloud/) cluster. A free account will suffice for a considerable period.
- In `.env` `OAUTH_TOKEN` is the Oauth token given to data providers using your Node.
- Run `make gomodtidy`, `make node-tables`, `make node-server-nolog HOST_IP=0.0.0.0:50054` or `make node-server HOST_IP=0.0.0.0:50054` on your local machine to test it locally

Refer to the KlayOracle full [documentation](https://klayoracle.gitbook.io/v1.0.0/) for detailed instruction.

### DP Runner

