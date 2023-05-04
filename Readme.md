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

- Install go version >= 18
- Docker desktop
- Mac/Linux machine

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

### DP Runner

- Using sample [KLAY_USD](data-provider/feeds/KLAY_USD.json) and [WEMIX_USD](data-provider/feeds/WEMIX_USD.json), add all the feeds your data provider will serving to consumer contract.
- Use the setup guide to understand and add a compatible feed. [Guide here](https://klayoracle.gitbook.io/v1.0.0/data-providers/)
- Using the `make` command, generate unique 32 bytes string identifier for each adapter feed. example ` ADAPTERS="KLAY_USD.json WEMIX_USD.json make adapter-id-gen"` will generate `adapterId` for each feed.
- Optionally if you prefer to prettifier your feed after generating `adapterId`, use https://jsonformatter.curiousconcept.com/.
- Deploy [OracleProvider](https://github.com/KlayOracle/klayoracle-monorepo/blob/development/oracle-contract/contracts/OracleProviderSample.sol) contract for each feed and replace the `oracleAddress` key.
- Add the certificate of the Node your data provider will be using to `data-provider/certs/node`. Update the path in `data-provider/config.yml`.
- Update your organization details in `data-provider/config.yml`.
- In `.env`, `HOST_IP` is the dns for reaching your data provider. As a rule of thumb if you are running multiple data provider don't run on same host,
- Add the list of feeds to be loaded to `data-provider/config.yml`. Any `feed` name not listed will not be sent to the node for aggregation.

Refer to the KlayOracle full [documentation](https://klayoracle.gitbook.io/v1.0.0/) for detailed instruction.

### Using Docker



