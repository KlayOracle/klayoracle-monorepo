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
- Docker Desktop
- Cockroach cli
- Mac/Linux machine

### Node Runner

- Generate a self-signed certificate for each node. [see guide using open ssl](setup-guide/openssl).
- Use different authority or wildcard matching authority of each node you intend to run.
- In `node/config.yml`, set the `key` and the `pem` you used to generate the certificate.
- Give the certificate pem file to data providers for authenticating with your Node. You can also run your own data providers.
- Update your organization details in `node/config.yml`. The website must match the authority used to sign the certificate; an example will be if the authority is `*.origineum.com`, the website can be `node-1.origineum.com`. Otherwise, the node service will fail to start.
- In `.env`, `HOST_IP` is the DNS for reaching your node. If you're running a docker container, you can override it and other `env` variables by passing from the terminal.
- In `.env` `PRIVATE_KEY` is your Node signer with enough Klay tokens to pay for the request.
- In `.env` `COCKROACH_DNS_URL` is your full connection string to [https://cockroachlabs.cloud/](https://cockroachlabs.cloud/) cluster. A free account will suffice for a considerable period.
- In `.env` `OAUTH_TOKEN` is the Oauth token given to data providers using your Node.
- Run `make gomodtidy`, `make node-tables`, `make node-server-nolog HOST_IP=0.0.0.0:50054` or `make node-server HOST_IP=0.0.0.0:50054` on your local machine to test it locally.

### DP Runner

- Using sample [KLAY_USD](data-provider/feeds/MANTLE_USD.json) and [WEMIX_USD](data-provider/feeds/WEMIX_USD.json), add all the feeds your data provider will serve to the consumer contract.
- Use the setup guide to understand and add a compatible feed. [Guide here](https://klayoracle.gitbook.io/v1.0.0/data-providers/)
- Using the `make` command, generate a unique 32 bytes string identifier for each adapter feed. Example `make adapter-id-gen ADAPTERS="KLAY_USD.json WEMIX_USD.json"` will generate `adapterId` for each feed.
- Optionally if you prefer to prettifier your feed after generating `adapterId`, use https://jsonformatter.curiousconcept.com/.
- Deploy [OracleProvider](https://github.com/KlayOracle/klayoracle-monorepo/blob/development/oracle-contract/contracts/OracleProviderSample.sol) contract for each feed and replace the `oracleAddress` key.
- Test if your Adapter will run successfully when sent to Node by running `make adapter-dry-run ADAPTERS="KLAY_USD.json WEMIX_USD.json"`
- Add the certificate of the Node your data provider will be using to `data-provider/certs/node`. Update the path in `data-provider/config.yml`.
- Update your organization details in `data-provider/config.yml`.
- In `.env`, `HOST_IP` is the dns for reaching your data provider. As a rule of thumb, if you are running multiple data providers, don't run on the same host,
- Add the list of feeds to be loaded to `data-provider/config.yml`. Any `feed` name not listed will not be sent to the node for aggregation.
- Run `make gomodtidy`, `make dp-client-nolog HOST_IP=0.0.0.0:50002` or `make dp-client HOST_IP=0.0.0.0:50002` on your local machine to test it locally.

For detailed instructions, please look at the KlayOracle full [documentation](https://klayoracle.gitbook.io/v1.0.0/).

### Available Price Feed


| 	                           | **Feed**   	                      | **Oracle Address**	 |
|-----------------------------|-----------------------------------|---------------------|
| **Mantle (Testnet)**  	     | 	                                 | 	                   |
| 	                           | MANTLE/USD  	                     | 0x	                 |
| 	                           | DOT/USD  	                        | 	0x                 |
| 	                           | AVAX/USD  	                       | 	0x                 |
| 	                           | AMZN/USD  	                       | 	0x                 |
| 	                           | ALGO/USD  	                       | 	0x                 |
| 	                           | AAPL/USD  	                       | 	0x                 |
| 	                           | BTC/USD  	                        | 0x	                 |
| 	                           | ETH/USD  	                        | 0x	                 |
| 	                           | BNB/USD  	                        | 0x	                 |
| 	                           | SOL/USD  	                        | 0x	                 |
| 	                           | LTC/USD  	                        | 0x	                 |
| 	                           | MATIC/USD  	                      | 	 0x                |
| 	                           | XRP/USD  	                        | 	 0x                |
| 	                           | PEPE/USD  	                       | 	 0x                |
| 	                           | TSLA/USD (Tesla)  	               | 	 0x                |
| 	                           | NVDA/USD (Nivida)  	              | 	 0x                |
| 	                           | MSFT/USD (Microsoft)  	           | 	 0x                |
| 	                           | META/USD (Facebook - META)  	     | 	 0x                |
| 	                           | JPM/USD (JP Morgan)  	            | 	 0x                |
| 	                           | GOOGL/USD (Alphabet Inc. CI A)  	 | 	 0x                |
| 	                           | GOOG/USD (Alphabet Inc. CI C)  	  | 	 0x                |
| **Klaytn Baobab (Testnet)** | 	                                 | 	                   |
| 	                           | MANTLE/USD  	                     | 0x	                 |
| 	                           | DOT/USD  	                        | 	0x                 |
| 	                           | AVAX/USD  	                       | 	0x                 |
| 	                           | AMZN/USD  	                       | 	0x                 |
| 	                           | ALGO/USD  	                       | 	0x                 |
| 	                           | AAPL/USD  	                       | 	0x                 |
| 	                           | BTC/USD  	                        | 0x	                 |
| 	                           | ETH/USD  	                        | 0x	                 |
| 	                           | BNB/USD  	                        | 0x	                 |
| 	                           | SOL/USD  	                        | 0x	                 |
| 	                           | LTC/USD  	                        | 0x	                 |
| 	                           | MATIC/USD  	                      | 	 0x                |
| 	                           | XRP/USD  	                        | 	 0x                |
| 	                           | PEPE/USD  	                       | 	 0x                |
| 	                           | TSLA/USD (Tesla)  	               | 	 0x                |
| 	                           | NVDA/USD (Nivida)  	              | 	 0x                |
| 	                           | MSFT/USD (Microsoft)  	           | 	 0x                |
| 	                           | META/USD (Facebook - META)  	     | 	 0x                |
| 	                           | JPM/USD (JP Morgan)  	            | 	 0x                |
| 	                           | GOOGL/USD (Alphabet Inc. CI A)  	 | 	 0x                |
| 	                           | GOOG/USD (Alphabet Inc. CI C)  	  | 	 0x                |
| **Klaytn Baobab (Mainnet)** | 	                                 | 	                   |
| 	                           | MANTLE/USD  	                     | 0x	                 |
| 	                           | DOT/USD  	                        | 	0x                 |
| 	                           | AVAX/USD  	                       | 	0x                 |
| 	                           | AMZN/USD  	                       | 	0x                 |
| 	                           | ALGO/USD  	                       | 	0x                 |
| 	                           | AAPL/USD  	                       | 	0x                 |
| 	                           | BTC/USD  	                        | 0x	                 |
| 	                           | ETH/USD  	                        | 0x	                 |
| 	                           | BNB/USD  	                        | 0x	                 |
| 	                           | SOL/USD  	                        | 0x	                 |
| 	                           | LTC/USD  	                        | 0x	                 |
| 	                           | MATIC/USD  	                      | 	 0x                |
| 	                           | XRP/USD  	                        | 	 0x                |
| 	                           | PEPE/USD  	                       | 	 0x                |
| 	                           | TSLA/USD (Tesla)  	               | 	 0x                |
| 	                           | NVDA/USD (Nivida)  	              | 	 0x                |
| 	                           | MSFT/USD (Microsoft)  	           | 	 0x                |
| 	                           | META/USD (Facebook - META)  	     | 	 0x                |
| 	                           | JPM/USD (JP Morgan)  	            | 	 0x                |
| 	                           | GOOGL/USD (Alphabet Inc. CI A)  	 | 	 0x                |
| 	                           | GOOG/USD (Alphabet Inc. CI C)  	  | 	 0x                |


>NB: Divide response by 1e9

### Video: Fetch Price Feed

[![Fetch KLAY/USD Price Feed using DigiOracle formerly KlayOracle](https://user-images.githubusercontent.com/7295729/257072009-19734d6b-1f6e-4cc7-8301-e6ffae5a7665.png)](https://www.youtube.com/watch?v=pJJK9vz_Y_Q)
