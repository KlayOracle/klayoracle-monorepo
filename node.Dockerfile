#Build Image
FROM golang:1.19-buster as build
RUN go version
WORKDIR /klayoracle

RUN mkdir node
RUN mkdir data-provider

COPY Makefile ./
COPY node /node/
COPY data-provider /data-provider/

COPY node/go.mod node/go.sum /node/
COPY data-provider/go.mod data-provider/go.sum /data-provider/

WORKDIR /data-provider
RUN go mod tidy

WORKDIR /node
RUN go mod tidy
RUN go build -o kloc-node . && cp -r . /var/klayoracle
# RUN rm /var/klayoracle/.env /var/klayoracle/config.yaml /var/klayoracle/certs/host/x509 -r

##Final Image
FROM ubuntu:20.04

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates

COPY --from=build /var/klayoracle /var/klayoracle

ARG PORT
EXPOSE $PORT
ENV WORK_DIR=/var/klayoracle

CMD ./var/klayoracle/kloc-node