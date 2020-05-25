# contracts

##Abstract
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


