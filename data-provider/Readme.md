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
      "headers": [{"field": {"X-CoinAPI-Key": "process.env.X_COIN_API_KEY"}}],
      "reducers": [{"function": "PARSE","args": ["rate"]},{"function": "MUL","args": ["1000000000"]}]
    }
  ]
}
```

## Step 2

Run:

```shell
ADAPTERS="KLAY_USD.json WEMIX_USD.json" make adapter-id-gen
```

Replace `KLAY_USD.json WEMIX_USD.json` with you list of Adapters to generate a new `adapter_id` for, seperated by single space.