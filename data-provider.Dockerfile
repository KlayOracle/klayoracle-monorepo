#Build Image
FROM golang:1.19-buster as gobuild
RUN go version
WORKDIR /klayoracle

COPY Makefile ./
COPY go.mod go.sum ./
RUN go mod download

COPY data-provider data-provider

RUN make node-build-provider

#Final Image
FROM ubuntu:20.04

#Install Dependecies needed
#Pass ENV and ARG
#Install kloc command
#Check health of node

CMD ["kloc", "run"]