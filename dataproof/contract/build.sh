#!/usr/bin/env bash

solcjs --bin --abi DataProof.sol ConvertLib.sol Migrations.sol
abigen --abi DataProof_sol_ProofDataPossession.abi --bin DataProof_sol_ProofDataPossession.bin --pkg proof  --out ../api/dataproof.go

rm ./*.bin
rm ./*.abi