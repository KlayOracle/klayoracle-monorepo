## Data Providers

### Step 1

Following the [docs](https://klayoracle.gitbook.io/v1.0.0/), setup adapters in the feed folder

Example 

```json
{
  "active": true,
  "name": "WEMIX_USD",
  "job_type": "DATA_FEED",
  "adapter_id":"",
  "oracle_address": "0xCC4377b912c4517Fe895817c6a7c6937D92A70B3",
  "category": 2,
  "feeds": [
    {
      "url": "https://min-api.cryptocompare.com/data/pricemultifull?fsyms=WEMIX&tsyms=USD",
      "request_type": 0,
      "headers": [{"field": {"Content-Type" : "application/json"}}],
      "reducers": [{"function": "PARSE","args": ["RAW","WEMIX","USD","PRICE"]},{"function": "MUL","args": ["1000000000"]}]
    },
    {
      "url": "https://rest.coinapi.io/v1/exchangerate/WEMIX/USD",
      "request_type": 1,
      "headers": [{"field": {"X-CoinAPI-Key": "${X_COIN_API_KEY}"}}],
      "reducers": [{"function": "PARSE","args": ["rate"]},{"function": "MUL","args": ["1000000000"]}]
    }
  ]
}
```

Add environment variables in the `.env` file. 
E.g. In the sample above, `${X_COIN_API_KEY}` will be substituted using the value of `X_COIN_API_KEY` in the .env file.


## Step 2

Run:

```shell
ADAPTERS="KLAY_USD.json WEMIX_USD.json" make adapter-id-gen
```

Replace `KLAY_USD.json WEMIX_USD.json` with you list of Adapters to generate a new `adapter_id` for, seperated by single space.

## Step 3 : Start Node

```shell
PORT=50051 make node-image
HOST_PORT=50051 NODE_PORT=50051 make node-container
```

While making image port `50051` will be exposed and while running container `HOST_PORT` on host is bind to `NODE_PORT`.

## Step 4 : Run 4 DP

In Dev "0.0.0.0:50002", "0.0.0.0:50003", "0.0.0.0:50004" are bootstrap data providers.

Run 3 DP containers for bootstraps and several others. The bootstrap data providers will register all joining dp and you should see them in container logs as they join.

```shell
PORT=50052 make dp-image



In reality, all of this will be running on different host, independently. So when the host is down all the nodes or dp won't go offline at once.