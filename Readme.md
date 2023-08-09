![codeql](https://github.com/KlayOracle/klayoracle-monorepo/actions/workflows/codeql.yml/badge.svg)
![dependency](https://github.com/KlayOracle/klayoracle-monorepo/actions/workflows/dependency-review.yml/badge.svg)
![linter](https://github.com/KlayOracle/klayoracle-monorepo/actions/workflows/golang-ci.yml/badge.svg)
![unit testing](https://github.com/KlayOracle/klayoracle-monorepo/actions/workflows/gounit-test.yml/badge.svg)
![slither analysis](https://github.com/KlayOracle/klayoracle-monorepo/actions/workflows/slither.yml/badge.svg)
![snyk check](https://github.com/KlayOracle/klayoracle-monorepo/actions/workflows/snyk.yml/badge.svg)

## Official DigiOracle monorepo

DigiOracle (formerly KlayOracle) is a blockchain oracle that helps smart contract developers connect Web3 applications to real-world events, data, and assets.

This monorepo contains all core packages that make up the DigiOracle protocol.

## 📚 Learn More

Docs
- [Developer documentation](https://klayoracle.gitbook.io/v1.0.0/)
- [Architecture Decision Records](/docs/arch)

Website: [digioracle.link](https://digioracle.link)

Blog: [digioracle-network.medium.com](https://digioracle-network.medium.com)

### 📦 Packages

- [data-provider](/data-provider): For Data Providers to serve data feeds to Data Subscribers.
- [node](/node): For Node runners to serve multiple Data Providers requests.
- [oracle-contract](/oracle-contract): Deployed by Data Providers for each data feed. 


### 📬 Contact

Feel free to email engineering@klayoracle.com

## Connect with KlayOracle

[Website](https://digioracle.link) | [Dorahacks](https://community.dorahacks.io/t/klayoracle-an-open-source-oracle-framework-to-securely-integrate-off-chain-data-with-klaytn-smart-contracts/58) | [Twitter](https://twitter.com/klayoracle) | [Medium](https://digioracle-network.medium.com/) | [Telegram](https://t.me/klayoracle) | [YouTube](https://www.youtube.com/@klayoracle)

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

- Using sample [KLAY_USD](data-provider/feeds/KLAY_USD.json) and [WEMIX_USD](data-provider/feeds/WEMIX_USD.json), add all the feeds your data provider will serve to the consumer contract.
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

| 	                           | **Feed**   	                      | **Oracle Address**	 | **Decimal** |
|-----------------------------|-----------------------------------|---------------------|-----------|
| **Klaytn Cypress (Mainnet)** | 	                                 | 	                   |  |
| 	                           | KLAY/USD  	                       | 0x62cB9Ccd9F08e27A89A9C8C5CCeEAEf7ddC4043B	                 | 9 |
| 	                           | DOT/USD  	                        | 	0x49573a03fDfBf54766c0BB3C20Db5aAfea749211                 | 9 |
| 	                           | AVAX/USD  	                       | 	0x8D58564D0e7394902a10105BbE183c9A44f02d1d                 | 9 |
| 	                           | AMZN/USD  	                       | 	0xf57761CC1FcF506Da4396E62C5920C1632c30620                 | 9 |
| 	                           | ALGO/USD  	                       | 	0xcbE8E1F7aF9e21192E4260152aDC91c8805Be222                 | 9 |
| 	                           | AAPL/USD  	                       | 	0x722222452eA9bE8dDF62118835394e8C079d2971                 | 9 |
| 	                           | BTC/USD  	                        | 0xc8D7Df132015f44FbB7D535Ca48677cC129566c8	                 | 9 |
| 	                           | ETH/USD  	                        | 0xce4bC34b8718C61901AB76b81e4e6BA8d096eAce	                 | 9 |
| 	                           | BNB/USD  	                        | 0x8D58564D0e7394902a10105BbE183c9A44f02d1d	                 | 9 |
| 	                           | SOL/USD  	                        | 0x70f1bB87C395a377476903AAE8b1A37b76347c8D	                 | 9 |
| 	                           | LTC/USD  	                        | 0x14cc5115AC7c9dE06151C91b6c9465ad381b0688	                 | 9 |
| 	                           | MATIC/USD  	                      | 	 0x1FAD86EfBa4FADCF49E79aBFAFEB0F54Ef7342a0                | 9 |
| 	                           | XRP/USD  	                        | 	 0x52fe38E36A3ec9Dc3dE1819a88411668E456bDE5                | 9 |
| 	                           | PEPE/USD  	                       | 	 0x672589c68b8C92bd0d644d67Cc6a148b18407C80                | 9 |
| 	                           | TSLA/USD (Tesla)  	               | 	 0xc9654D64c0f783d87271c5edA514951436a9a8ea                | 9 |
| 	                           | NVDA/USD (Nivida)  	              | 	 0x81430EDf87Fe20A95d2766a8EA5f31d4D4509a09                | 9 |
| 	                           | MSFT/USD (Microsoft)  	           | 	 0xa618871566f9446AefC671792874fA6628A9BA49                | 9 |
| 	                           | META/USD (Facebook - META)  	     | 	 0x8DFAC6650073d5818Eab9F4a150eDE1B8471050D                | 9 |
| 	                           | JPM/USD (JP Morgan)  	            | 	 0xe2B5222eDA290b79Ee8cb2a6ba67Bfc1C377f972                | 9 |
| 	                           | GOOGL/USD (Alphabet Inc. CI A)  	 | 	 0x65a14cf12730eBF7874eE459F3C9F5d33e909c15                | 9 |
| 	                           | GOOG/USD (Alphabet Inc. CI C)  	  | 	 0x0080f3FF6ca9123Cf069D2C5A35d6463E6Faa0BD                | 9 |
| **Klaytn Baobab (Testnet)** | 	                                 | 	                   |  |
| 	                           | KLAY/USD  	                       | 0xbc884088e406422a3ef39aedd1c546de7ac4be7c	                 | 9 |
| 	                           | WEMIX/USD  	                      | 	0x251b2c534fa7d696b356a96b17bf87a2ad38f394                 | 9 |
| 	                           | DOT/USD  	                        | 	0x                 | 9 |
| 	                           | AVAX/USD  	                       | 	0x                 | 9 |
| 	                           | AMZN/USD  	                       | 	0x                 | 9 |
| 	                           | ALGO/USD  	                       | 	0x                 | 9 |
| 	                           | AAPL/USD  	                       | 	0x                 | 9 |
| 	                           | BTC/USD  	                        | 0xf57761CC1FcF506Da4396E62C5920C1632c30620	                 | 9 |
| 	                           | ETH/USD  	                        | 0x8D58564D0e7394902a10105BbE183c9A44f02d1d	                 | 9 |
| 	                           | BNB/USD  	                        | 0x1FdebbE196D2FeE7Ca98126a8b1f0f42a6E2833f	                 | 9 |
| 	                           | SOL/USD  	                        | 0xc8D7Df132015f44FbB7D535Ca48677cC129566c8	                 | 9 |
| 	                           | LTC/USD  	                        | 0xce4bC34b8718C61901AB76b81e4e6BA8d096eAce	                 | 9 |
| 	                           | MATIC/USD  	                      | 	 0x0080f3FF6ca9123Cf069D2C5A35d6463E6Faa0BD                | 9 |
| 	                           | XRP/USD  	                        | 	 0x                | 9 |
| 	                           | PEPE/USD  	                       | 	 0x                | 9 |
| 	                           | TSLA/USD (Tesla)  	               | 	 0x                | 9 |
| 	                           | NVDA/USD (Nivida)  	              | 	 0x                | 9 |
| 	                           | MSFT/USD (Microsoft)  	           | 	 0x                | 9 |
| 	                           | META/USD (Facebook - META)  	     | 	 0x                | 9 |
| 	                           | JPM/USD (JP Morgan)  	            | 	 0x                | 9 |
| 	                           | GOOGL/USD (Alphabet Inc. CI A)  	 | 	 0x                | 9 |
| 	                           | GOOG/USD (Alphabet Inc. CI C)  	  | 	 0x                | 9 |
| **Mantle (Testnet)**  	     | 	                                 | 	                   |  |
| 	                           | MANTLE/USD  	                     | 0x94bfd13Dc52951a81c77c24350d1911e1C17cd65	                 | 9 |
| 	                           | DOT/USD  	                        | 	0xA0698f9889a1B7906B2C236FAb7a9c5d128fa32a                 | 9 |
| 	                           | AVAX/USD  	                       | 	0x6f785c43FBB5569b7f1A57fe62fBCCE96175612a                 | 9 |
| 	                           | AMZN/USD  	                       | 	0x0fCB49fb0c573d4273126f19EaF6A696B5881DEB                 | 9 |
| 	                           | ALGO/USD  	                       | 	0x98E46422C5f0ce5630782604434B5A89CA29A051                 | 9 |
| 	                           | AAPL/USD  	                       | 	0x3101B216F301E746C3E3b96D6A8AF7983717e506                 | 9 |
| 	                           | BTC/USD  	                        | 0x85190f5303D85c2668275cE7A07758ba22cE0C35	                 | 9 |
| 	                           | ETH/USD  	                        | 0x7274F4c13438Bbf955eA87A883679511F8921f36	                 | 9 |
| 	                           | BNB/USD  	                        | 0x336238318d3a8c421207193173f80101D2d78c7e	                 | 9 |
| 	                           | SOL/USD  	                        | 0xC906384361FD1D58DE1C6f43Fe4Ad9239FFF701b	                 | 9 |
| 	                           | LTC/USD  	                        | 0x7c023184b0949794F8df014915a8deBdE73A432a	                 | 9 |
| 	                           | MATIC/USD  	                      | 	 0x1FAD86EfBa4FADCF49E79aBFAFEB0F54Ef7342a0                | 9 |
| 	                           | XRP/USD  	                        | 	 0x0b84b6580b85e82C068324C379af136a820a28A0                | 9 |
| 	                           | PEPE/USD  	                       | 	 0x3bbE624FCED00B4557dCd1d157d241a9dA0D1654                | 9 |
| 	                           | TSLA/USD (Tesla)  	               | 	 0x5E8466CD85Af8e3DF3FdC2d800c5361d6a6E9F4e                | 9 |
| 	                           | NVDA/USD (Nivida)  	              | 	 0x2070c4d9a6Bf7E291edED112584162A5469F53b4                | 9 |
| 	                           | MSFT/USD (Microsoft)  	           | 	 0xE4725EdF77460Bf1053a8c575c255E0818C8c04A                | 9 |
| 	                           | META/USD (Facebook - META)  	     | 	 0x463B123af02E3d75403Ed79d447aA0E8d7501106                | 9 |
| 	                           | JPM/USD (JP Morgan)  	            | 	 0x81B238602226Ea019876676E9617737e4BDC5447                | 9 |
| 	                           | GOOGL/USD (Alphabet Inc. CI A)  	 | 	 0x87bD12e9FB32d390b776A59aA63b102a7664A5F6                | 9 |
| 	                           | GOOG/USD (Alphabet Inc. CI C)  	  | 	 0x9b18663631A53d891f0C7199bb656A4863a94D5a                | 9 |


### Video: How to Fetch Price Feed from Oracle

[![Fetch KLAY/USD Price Feed using DigiOracle formerly KlayOracle](https://user-images.githubusercontent.com/7295729/257072009-19734d6b-1f6e-4cc7-8301-e6ffae5a7665.png)](https://www.youtube.com/watch?v=pJJK9vz_Y_Q)
