#!/usr/bin/env bash

set -eo pipefail

cd -- "$( dirname -- "${BASH_SOURCE[0]}" )"

run() {
    /usr/bin/env \
        BRISKITALL_CHAINID="$(cat ./config/chainid)" \
        BRISKITALL_NODE_URL="http://localhost:8545" \
        BRISKITALL_CLEF_ENDPOINT="./runtime/clef/clef.ipc" \
        go run ../.. "$@"
}

run "$@"
