# ETHFSx Contracts ![](https://img.shields.io/badge/status-wip-orange.svg?style=flat-square) 

## Abstract
The core function provided by the contract is a data holding contract, in order to decentralize storage system services; it mainly includes: 
- token pledge;
- storage challenge; 
- challenge response;
- challenge verification;
- token revocation;
- token destruction;

## Usage
In order to make it work, we need to install some compile tools and compile solidity file firstly:
```go
~ npm install -g truffle && npm install -g ethereumjs-testrpc  --unsafe-perm
~ cd contracts/dataproof/contract
~ ./build.sh
```
By then, we will find a go file in api dir named dataproof.go

## Deploy
This contract has been deployed in ethereum test net named Rinkeby, you can explorer it by the url as follow:
```text
https://rinkeby.etherscan.io/tx/0x4befef1f86b12fa7d62da486acdddb90656801ead596771bbd0137877dbcd1a4
```


