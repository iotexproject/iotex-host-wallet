iotex-host-wallet
=================

This repo shows how to implement hosted wallet on IoTeX network using golang. 

## wallet

Host wallet sources.

## Tools

Master PrivateKey Generator

## Install

### Build

1. Build tools

```bash
cd tools
go mod download
go build .
```

1. Build wallet

```bash
cd wallet
go mod download
go build .
```

### Deploy

```bash
mkdir service
cp tools/tools .
cp wallet/wallet .
cp wallet/config.yaml .

// 1. generate two RSA keypairs, one for service, one for wallet
// 2. config key to config.yaml
// 3. install mongo and change config.yaml

// generate master key with password
./tools -p 123455

// start wallet
echo 123456 | nohup ./wallet &
```
