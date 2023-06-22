#!/usr/bin/env bash

set -e -o pipefail

SCRIPTDIR="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

cd "${SCRIPTDIR}"

SOLC_VERSION=0.4.19
ABIGEN_VERSION=$(go list -m all | grep github.com/ethereum/go-ethereum | awk '{print $2}')

SOLS=( MultiSigWalletWithDailyLimit.sol CentrallyIssuedToken.sol )

cleanup() {
	rm -rf combined.json
    rm -f contract.go.tmp
}
trap cleanup EXIT

# Generate the combined.json file for all contracts
docker run \
    -v "$PWD:/wd" \
    -w /wd \
    --user $(id -u):$(id -g) \
    --rm \
    ethereum/solc:"${SOLC_VERSION}" \
    -o . \
    --combined-json "asm,bin,abi,userdoc,devdoc,metadata" \
    --overwrite \
    --optimize-runs 200 \
    -- \
    "${SOLS[@]}"

# Generate the ABI using the combined.json file
docker run \
    -i --rm \
    --user $(id -u):$(id -g) \
    ethereum/client-go:alltools-v1.11.6 abigen\
    --combined-json - \
    --pkg contract \
    < combined.json > contract.go.tmp

mv contract.go.tmp contract.go
