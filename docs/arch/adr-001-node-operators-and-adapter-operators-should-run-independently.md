
# Architecture Design

**Title:** Node Operators & Adapter Should Run Independently (Architecture Design)

**Date:** 2023-01-14

**Status:** **Draft**

# **Context**

## Original KlayOracle architecture & its problems

In KlayOracle’s original architecture, data feeds, oracle contracts, and nodes were tightly coupled. Node operators were required to set up data feeds and define adapters (a list of data feeds) directly within their nodes.

The challenge this brings up is that: (i) there is an additional cost & time responsibility on node operators to source and pay for data sources in order to run data feeds. (ii) in cases where they are unable to acquire enough data sources, CPU resources will be wasted due to a low number of data feeds, hence making the idea of running nodes unattractive.

It would be more difficult to achieve true scale and decentralization if one party (the node runner) is responsible for providing computational resources as well as paying for different APIs to use as off-chain data sources.

![Deprecated Design of KlayOracle v0.0.1](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/8ef1939b-e235-4fcb-8a93-56c50f60c4eb/Untitled.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Credential=AKIAT73L2G45EIPT3X45%2F20230114%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20230114T102406Z&X-Amz-Expires=86400&X-Amz-Signature=b1ecdb888bb5e7a9c1702a6a33f6bff901e7154641d2b89f455a6d1daf598063&X-Amz-SignedHeaders=host&response-content-disposition=filename%3D%22Untitled.png%22&x-id=GetObject)

Deprecated Design of KlayOracle v0.0.1

# Decision

## Summary of Proposed Solution

Our proposed architecture update focuses on one key theme — the separation of concerns in order to maximize the resources of each participant in the network. In the new KlayOracle, adapters have been renamed as Data Providers and will run independently from nodes.

What if Data Providers focus on sourcing data feeds and adding more over time, while node operators run as many data providers as are willing to use them, so far as their node(s) have the capacity to fulfill their requests?

<aside>
💡 In KlayOracle’s new architecture, data providers can now configure their respective data feeds and oracle contracts and subscribe to one or several **verified nodes** to fetch & compute these data based on the provider’s configuration.

They subscribe to these node runners by paying them a monthly fee to run the computations necessary to fetch data feeds and maintain their truthiness on-chain.

Data providers and nodes will communicate and exchange information via gRPC and Protocol Buffer.

</aside>

Data subscribers (smart contract, API services, Wallets) in turn subscribe to Data Providers in order to access on-chain data, by being whitelisted on the Oracle Contract deployed by the Data Provider.

## Components of the new Architecture

The following section highlights how the main components of KlayOracle’s new architecture interact with each other.

![Architechture KlayOracle V1.png](https://s3.us-west-2.amazonaws.com/secure.notion-static.com/32a5a486-caed-47f0-8469-d5202c690cb4/Architechture_KlayOracle_V1.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Content-Sha256=UNSIGNED-PAYLOAD&X-Amz-Credential=AKIAT73L2G45EIPT3X45%2F20230114%2Fus-west-2%2Fs3%2Faws4_request&X-Amz-Date=20230114T102441Z&X-Amz-Expires=86400&X-Amz-Signature=eb399370ee5c2319d391029325f97b08e3cb71c71264fb0296e50728f3623973&X-Amz-SignedHeaders=host&response-content-disposition=filename%3D%22Architechture%2520KlayOracle%2520V1.png%22&x-id=GetObject)

### Nodes

- Node runners manually register with KlayOracle to become verified on the Network. While anyone can still run nodes on the network without being verified, we intend to make it unattractive to do so, except in a private use case. We’ll also signal and promote only verified nodes for public use cases.
- Before a node is verified, it must stake a certain amount of KlayOracle tokens.
- These tokens may be slashed if a dispute is raised by a KlayOracle token holder within a specific threshold defined in the DisputeContract smart contract.
    - In this case, if the node is found guilty (through a manual voting process by the community), the slashed amount will be distributed to the voters securing the network, and if not found guilty, the amount is retained within the node’s stake.
    - KlayOracle’s dispute resolution process is currently a work-in-progress [1].
- KlayOracle will host 3 — 5 **bootstrapped nodes**, which will be used to coordinate nodes when joining the network.
- Nodes can find and connect to peer nodes (other nodes within the network) by discovery through the bootstrapped nodes [2] using [libp2p](https://libp2p.io/). Peer nodes need to be connected within the network in order to ensure present and future decentralization of the network.
- Components of a **node**:
    - Configuration:
        - Website
        - Contact
        - Whitelist of Data Providers
            - Public key address
    - Environmental Variables (.env)
        - bootstrap: `default false` — checks if the node is a bootstrap node
        - JSON Wallet secret key
    - Core:
        - The node connects to any of the bootstrap nodes to find peers.
        - The node waits for whitelisted data Providers to ping-check for its liveliness.
        - The node registers all whitelisted data providers and stores their details in memory.
        - The node then uses the configuration of each data provider to aggregate data:
            - Fetch data for the data provider when heartbeat time passes
            - Fetch data for the data provider when the deviation threshold is met
        - Before updating on-chain data, ping the Data Provider to check for liveliness.
            - Open Dispute on DisputeContract if liveliness check fails
        - After aggregating data & verifying liveliness, update **Oracle Contract** with a frequency based on the aggregation options set by the Data Provider.
            - Aggregation options — Random ****(Deprecated), Median, Volume Weighted Average.
        - gRPC methods (Unary)
            - Request to List Data Providers
            - Request List of Nodes
            - Request List of Data Fulfilled (time range from Postgres DB)

### Data Providers

- Data Providers will run separately from nodes, as container services hosted on-premise or on the cloud.
- Data providers independently define their configuration options — including their preferred data feeds, data fetching, and aggregation options.
- The data provider can choose to use one or several **verified nodes** to compute data, based on the provider’s configuration options.
- Data Providers choose the node(s) to use by adding their address(es) in their config options. The node discovers the data provider once the provider is online and pings the node for liveliness every 300 seconds.
    
    The node registers and connects to the data provider only if the data provider’s address is whitelisted by the network.
    
- KlayOracle will host 3 — 5 **bootstrapped data providers**, which will be used to coordinate data providers when joining the network.
- Data providers can find and connect to peer data providers by discovery through the bootstrapped data providers [2] using [libp2p](https://libp2p.io/). Peer data providers need to be connected within the network in order to ensure present and future decentralization of the network.
- Before a Data Provider is verified, it stakes KlayOracle tokens in a Contract.
- These tokens may be slashed by the protocol in cases of outage or data mistrust.
    - Cases of outage are automatically resolved by the protocol — an outage is published [*DevOps] by the node whenever it pings the data provider and the provider is offline. In turn, the publish event triggers a service that opens an outage dispute on the DisputeContract smart contract.
    - Cases of data mistrust may be reported by the community. In this case, if the data provider is found guilty, the slashed amount will be distributed to the voters securing the network, and if not found guilty, the amount is retained within the data provider’s stake.
    - KlayOracle’s dispute resolution process is currently a work-in-progress [1].
- **Data Provider** Components:
    - **Data Feeds**: Folder containing a list of data feeds served by a provider. Each list is defined in a JSON file, and the name of each file is the title of the Data feed, e.g `ETH_KLAY.json`
        - **Name**: of Data Feed, e.g ETH_KLAY. Must match the file name in the format {Name}.json
        - **Deviation threshold**: the percentage of deviation from the last returned results should trigger the Node to update the on-chain feed.
        - **Heartbeat**: time that passes before the Node re-fetches data from the data feed, regardless of deviation.
        - **Adapter ID**: a unique string
        - **Type**: Currently defaults to `DATA_FEED`.
        - **Oracle**: Oracle contract that holds on-chain data for this particular data feed.
        - **Feeds**:
            - See [Ref 3](https://www.notion.so/Architecture-Design-c2dca4b749eb4609a7ad9b113adfd529) for a further breakdown of feeds
    - **Config**:
        - Website
        - Contact
        - Fee (in KlayOracle token) — which the subscriber would pay to access the data provider.
        - Node Address (serving this provider)
    - **.ENV**
        - JSON Wallet secret key
        - SSL Certificate
    - **Core:**
        - Generate Wallet
        - Ensure Data Provider’s data feeds (JSON files) are all valid
        - Find and connect to peer data provider
        - Find and connect to node successfully
        - gRPC methods (Unary)
            - Request Health check
            - Provide a particular Data Feed
            - Provide all Data Feeds
        - Pings the Node to register liveliness every 300 secs - Request must be signed using Provider private key wallet for Node to confirm Whitelist
    - Logging
        - [*DevOps]
    - Aggregation Options:
        - Random ****(Deprecated)
        - Median
        - Volume Weighted Average

### Oracle Contract

- Oracle contracts when deployed are tied to specific data Feeds (e.g ETH_KLAY) throughout their lifetime on-chain. Data Providers are responsible for deploying them for each data feed they are serving to data subscribers.
- KlayOracle provides an Abstract contract Oracle contract, which Data Providers extend when writing and deploying their own Oracle implementation contract. The only difference will be overriding the `isWhitelisted` and `beforeFufill` functions, which specify how data providers want to authenticate the data subscribers calling them for the latest data feeds.
    
    ```solidity
    //Example: Whitelist
    function _isWhitelisted() returns(bool) {
    	return whitelist[msg.sender] == true
    }
    
    //Example fees
    function _isWhitelisted() returns(bool) {
    	return IERC20(klayOracleTokenAddress).allowance(msg.sender, address(this)) > oracleFee
    }
    
    function _beforeFufill() returns(bool) {
    	require(IERC20(klayOracleTokenAddress).transferFrom(msg.sender, address(this), oracleFee),"Do not honor!")
    }
    ```
    
- **Oracle Contract** Components:
    - **Core:**
        - Verify that bytes sent from the node are signed by Node (EIP 712) [4, 5]
            - not concluded, as vanilla signing might just be enough
        - Verifies `msg.sender` is expected Node
        - Sends Data Feed to consumer contract with timestamp
        - Keeps track of all requests using a mapping of the hash of Struct holding the Request details to the Request Details.
        - emits Events for a Request, Data updates for Indexing purposes

### Data Subscriber

- []

---

## KlayOracle Explorer

KlayOracle will build and maintain a public dashboard known as ***KlayOracle Explorer*** — which will display all the public nodes & data providers running on the KlayOracle network. Verified nodes will be highlighted on this dashboard.

**Each node on the KlayOracle Node Explorer will have the following details:**

- Data Providers connected to the node
- Data Feed(s) defined by all data providers connected to the node
- Verified status
- Uptime *%*
- Total Uptime
- Key Statistics of I/O requests made by the node.

**Each Data Provider listed on the Explorer will have the following details:**

- Data Feeds
- Last found data
- Previous data history
- Last Update (*time ago*)
- Oracle Address
- Deviation
- Retry
- Requests Details Graph.

**The main dashboard will also highlight key m**etrics of Nodes & Data Providers

# Consequences

- More complex architecture — this will incur significant time and resources.
- Address security concerns between node and data provider's communication.
- Additional responsibility to source for Data Providers separately instead of just Node runners.
- Necessary to deploy KlayOracle’s token in order to use economic incentives to ensure the security of the platform.

# References

- [1] — [Inspiration for the dispute resolution process](https://github.com/redstone-finance/redstone-oracles-monorepo/blob/main/packages/oracle-node/docs/DISPUTE_RESOLUTION.md)
- [2] — [How nodes find each other](https://ethereum-classic-guide.readthedocs.io/en/latest/docs/appendices/how_nodes_find_each_other.html)
- [3] — [Reference to feeds](https://github.com/alofeoluwafemi/klay-oracle/blob/development/adapter/jobs/klay_usd.adpater.json)
- [4] — [EIP-712](https://eips.ethereum.org/EIPS/eip-712)
- [5] — [Sign an EIP-712 message](https://docs.venly.io/widget/widget/sign-a-eip712-message)

# Required Documentation

- Developer Doc
- Architecture Diagram
- Requirements for Data Provider & Node Runners
    - Gas
    - Storage
    - Network
