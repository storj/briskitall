#!/usr/bin/env bash

set -eo pipefail

cd -- "$( dirname -- "${BASH_SOURCE[0]}" )"

PRYSMCTL="gcr.io/prysmaticlabs/prysm/cmd/prysmctl:latest"
GETHALLTOOLS="ethereum/client-go:alltools-latest"
KEYSTORE="./runtime/keystore"

dockerrun() {
    docker run -u "$(id -u):$(id -g)" "$@"
}

prepare_runtime() {
    rm -rf ./runtime
    mkdir ./runtime
    cp -r ./{config,runtime}/consensus
    cp -r ./{config,runtime}/geth
}

generate_env() {
    cat << EOF > .env
CHAINID="$(cat config/chainid)"
ROOT_ACCOUNT="$(cat ../testdata/root.addr)"
DOCKER_USER="$(id -u):$(id -g)"
EOF
}

update_geth_genesis() {
    jq --arg chainid "$(cat ./config/chainid)" '.config.chainId = ($chainid | tonumber)' \
        ./runtime/geth/genesis.json > ./runtime/geth/genesis.json.tmp
    mv ./runtime/geth/genesis.json{.tmp,}

    while read -r addr; do
        echo "Seeding $addr with 1 ETH..."
        jq --arg addr "$addr" '.alloc[$addr] = { "balance": "0xDE0B6B3A7640000" }' \
            ./runtime/geth/genesis.json > ./runtime/geth/genesis.json.tmp
        mv ./runtime/geth/genesis.json{.tmp,}
    done < seed.txt
}

generate_beacon_chain_genesis() {
    dockerrun \
        --rm -v "./runtime:/runtime" \
        "${PRYSMCTL}" testnet generate-genesis \
            --fork=bellatrix \
            --num-validators=64 \
            --output-ssz=/runtime/consensus/genesis.ssz \
            --chain-config-file=/runtime/consensus/config.yaml \
            --geth-genesis-json-in=/runtime/geth/genesis.json \
            --geth-genesis-json-out=/runtime/geth/genesis.json
}

generate_beacon_chain_auth_secret() {
    openssl rand -hex 32 | tr -d "\n" > "./runtime/jwt.hex"
    cp ./runtime/jwt.hex ./runtime/consensus/jwt.hex
    mv ./runtime/jwt.hex ./runtime/geth/jwt.hex
}

init_geth() {
    dockerrun --rm \
        -v "./runtime:/runtime" \
        "${GETHALLTOOLS}" geth \
            --datadir /runtime/geth \
            init /runtime/geth/genesis.json
}

bootstrap() {
    prepare_runtime
    generate_env
    update_geth_genesis
    generate_beacon_chain_genesis
    generate_beacon_chain_auth_secret
    init_geth
}

if [ -z "$1" ]; then
    bootstrap
else
    "$@"
fi
