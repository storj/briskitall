#!/usr/bin/env bash

set -e

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
TESTENV=.test.env

BRISKITALL_CHAINID=1337
export BRISKITALL_CHAINID


has_cmd() {
    builtin type -P "$1" &> /dev/null
}

load-env() {
    [ ! -f "${TESTENV}" ] || source "${TESTENV}"
}

write-env() {
    cat << EOF > "${TESTENV}"
CONTAINER="${CONTAINER}"
CONTAINERPORT="${CONTAINERPORT}"
MULTISIG_CONTRACT_ADDRESS="${MULTISIG_CONTRACT_ADDRESS}"
TOKEN_CONTRACT_ADDRESS="${TOKEN_CONTRACT_ADDRESS}"
EOF
}

start-geth() {
    load-env
	CONTAINER=$(docker run -d --rm \
		-p :8545 \
		ethereum/client-go \
        --dev --datadir . --http --http.addr 0.0.0.0 --http.api eth,web3,net,debug)
    CONTAINERPORT=$(docker port "${CONTAINER}" 8545/tcp | head -n1 | cut -d ':' -f2)
    write-env
    while ! docker exec "${CONTAINER}" geth attach --datadir . --exec "accounts" &> /dev/null ; do
        echo waiting for geth...
        sleep 1
    done
}

stop-geth() {
    load-env
    [ -z "${CONTAINER}" ] || docker stop "${CONTAINER}" > /dev/null
    CONTAINER=
    CONTAINERPORT=
    write-env
}

exec-geth() {
    load-env
    docker exec "${CONTAINER}" geth attach --datadir . --exec "$1"
}

fund() {
    if exec-geth "miner.stop; eth.sendTransaction({from: eth.accounts[0], to: \"$1\", value: web3.toWei(10, \"ether\")}); miner.start; admin.sleepBlocks(1);" > /dev/null; then
        echo "funded $1"
    else
        echo "failed to fund $1"
    fi
}

run() {
    load-env
    /usr/bin/env \
        BRISKITALL_NODE_URL="http://localhost:${CONTAINERPORT}" \
        BRISKITALL_MULTISIG_CONTRACT_ADDRESS="${MULTISIG_CONTRACT_ADDRESS}" \
        BRISKITALL_TOKEN_CONTRACT_ADDRESS="${TOKEN_CONTRACT_ADDRESS}" \
        go run . "$@"
}

load-accounts() {
    ADDRESS=( "$(cat test/testdata/account-0.addr)" )
    ADDRESS+=( "$(cat test/testdata/account-1.addr)" )
    ADDRESS+=( "$(cat test/testdata/account-2.addr)" )
    ADDRESS+=( "$(cat test/testdata/account-3.addr)" )

    KEYFILE=( test/testdata/account-0.key )
    KEYFILE+=( test/testdata/account-1.key )
    KEYFILE+=( test/testdata/account-2.key )
    KEYFILE+=( test/testdata/account-3.key )
}

list-accounts() {
    load-accounts
    for addr in "${ADDRESS[@]}"; do
        echo "$addr"
    done
}

deploy-multisig() {
    load-env
    load-accounts
    MULTISIG_CONTRACT_ADDRESS=$(run test deploy multisig \
        "${ADDRESS[0]}" "${ADDRESS[1]}" \
        --sender-key-file "${KEYFILE[0]}" --jsonout | jq .ContractAddress)
    write-env
}

deploy-token() {
    load-env
    load-accounts
	if [ -z "${MULTISIG_CONTRACT_ADDRESS}" ]; then
		echo "Multisig contract is not deployed yet"
		exit 2
	fi
    TOKEN_CONTRACT_ADDRESS=$(run test deploy token \
        "${MULTISIG_CONTRACT_ADDRESS}" \
        --sender-key-file "${KEYFILE[0]}" --jsonout | jq .ContractAddress)
    write-env
}

up() {
    load-accounts
    start-geth
    for addr in "${ADDRESS[@]}"; do
        fund "${addr}"
    done
    deploy-multisig
    deploy-token
}

down() {
    stop-geth
    rm -f ./"${TESTENV}"
}

if ! has_cmd jq; then
    echo "jq is required"
    exit 2
fi

cd "${SCRIPT_DIR}"
"$@"
