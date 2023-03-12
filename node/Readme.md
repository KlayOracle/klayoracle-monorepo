## Data Providers

From root directory `./klayoracle-monorepo` build image

```shell
docker build -t klayoracle-node:latest -f node.Dockerfile . --build-arg PORT=50051
```

```shell
docker run -d -p 50051:50051 --env-file node/.env klayoracle-node:latest