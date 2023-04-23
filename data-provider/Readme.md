## Data Providers

### Step 1

Following the [docs](https://klayoracle.gitbook.io/v1.0.0/), setup adapters in the feed folder

Example 

```json
{
  "active":true,
  "name":"KLAY_USD",
  "adapterId":"8b7460cccfa0aca303ee85c3fb81c344faad2fbab415adc32b2984008b7efd76",
  "oracleAddress":"0xCC4377b912c4517Fe895817c6a7c6937D92A70B3",
  "category": 2,
  "frequency": 30000000000,
  "feeds":[
    {
      "url":"https://min-api.cryptocompare.com/data/pricemultifull?fsyms=KLAY&tsyms=USD",
      "request_type":0,
      "headers":[
        {
          "field":{
            "Content-Type":"application/json"
          }
        }
      ],
      "reducers":[
        {
          "function":"PARSE",
          "args":[
            ["$.RAW.KLAY.USD.PRICE"]
          ]
        },
        {
          "function":"FLOAT64_MUL_UINT64",
          "args":[
            "1000000000"
          ]
        }
      ],
      "payload": ""
    },
    {
      "url":"https://rest.coinapi.io/v1/exchangerate/KLAY/USD",
      "request_type": 0,
      "headers":[
        {
          "field":{
            "X-CoinAPI-Key":"${X_COIN_API_KEY}"
          }
        }
      ],
      "reducers":[
        {
          "function":"PARSE",
          "args":[
            ["$.rate"]
          ]
        },
        {
          "function":"MUL",
          "args":[
            "1000000000"
          ]
        }
      ]
    }
  ]
}
```

### Understanding JSON path usage for PARSE Reducer

- https://pkg.go.dev/github.com/PaesslerAG/jsonpath#example-Get
- [node](../node/reducer/reducer_test.go)
- https://github.com/PaesslerAG/jsonpath/blob/master/testdata/regression_suite.yaml

Add environment variables in the `.env` file. 
E.g. In the sample above, `${X_COIN_API_KEY}` will be substituted using the value of `X_COIN_API_KEY` in the .env file.


## Step 2

Run:

```shell
ADAPTERS="KLAY_USD.json WEMIX_USD.json" make adapter-id-gen
```

Replace `KLAY_USD.json WEMIX_USD.json` with you list of Adapters to generate a new `adapter_id` for, seperated by single space.

Run:

```shell
ADAPTERS=KLAY_USD.json WEMIX_USD.json make adapter-dry-run
```

This will dry run data aggregation for one round to see if everythign will run fine. Saves you time from sending adapter configuration that won't run on the Node.

> Use https://jsonformatter.org/ to beautify the adapter feeds content, if you want.

## Step 3 : Start Node

```shell
PORT=50051 make node-image
HOST_PORT=50051 NODE_PORT=50051 make node-container
```

While making image port `50051` will be exposed and while running container `HOST_PORT` on host is bind to `NODE_PORT`.

## Step 4 : Run 4 DP

In Dev "0.0.0.0:50002", "0.0.0.0:50003", "0.0.0.0:50004" are bootstrap data providers.

Run 3 DP containers for bootstraps and several others. The bootstrap data providers will register all joining dp and you should see them in container logs as they join.

