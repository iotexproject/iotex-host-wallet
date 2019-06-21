iotex-host-wallet
=================

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
// 2. config wallet private key to config.yaml
// 3. install mongo and change connection to config.yaml

// login mongo and insert service apikey and public key
db.service.save({name:"vita", api_key:"API-3d1b9d29-3a7e-4163-aeda-6255411c27b9", status:"normal", public_key: "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAh66fsHrXdxd1WhM1jPMwG4LpkVenHrw6amX+n0i7ks264b+qrD8vGOhWE7iaYM7jKURAv5th3EVgSs+njFgz0fFgHe5ntkkyEnzE+Gyd+P0Bne6Ve2T6uHYOFLcPMNFLkzfpX07YR3eUEehtunwjqgdSokKOl0QMfRbTmtSJBeb7DVSjlHWv8VZp+W6Lj/M6dpOU0EDpFmKO6PC112YRvIbiUjT4tNOyugIy3oWAHvmGD2pAw/LlY4N2lP08mntsT3nWIwpuVKWPkXcIttI6+9JdoERs2lyda2rojSJj2V7OStufPeYqjcXVErLuFWG5rHWLRdnm0z3fz2+wxO4fowIDAQAB", created_at:1560396059})

// generate master key with password
./tools -p 123455

// start wallet
echo 123456 | nohup ./wallet &
```
