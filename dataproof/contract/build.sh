#!/usr/bin/env bash

solcjs --bin --abi DataProof.sol ConvertLib.sol Migrations.sol
abigen --abi DataProof_sol_MetaCoin.abi --bin DataProof_sol_MetaCoin.bin --pkg proof  --out ../api/dataproof.go

rm ./*.bin
rm ./*.abi