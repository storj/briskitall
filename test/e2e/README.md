# End-to-End Testing

This directory contains helper scripts used to test brisk-it-all against a
"real" test network (i.e. not developer-mode geth).

## Prerequisites

- Go
- docker
- jq

## Bootstrapping

To bootstrap the test, run:

    $ ./bootstrap.sh

This script prepares a `./runtime` directory with all of the necessary
configuration, including:

1. Geth data directory initialized with:
    - Proof of stake consensus (with signer accounts seeded with 10000 ETH each).
    - Accounts in [seed.txt](./seed.txt) seeded with 1 ETH each.
1. Beacon chain data directory.

# Starting

1. Bring up geth and the prysm beacon chain and validator.

    $ docker-compose up -d

# Playing Around

The `./briskitall.sh` script is a small helper script that wraps execution of `briskitall` with a few helpful environment variables (e.g. the chain ID, the geth node URL).

## Deploy the MultiSig contract

First things first, you probably want to deploy the MultiSig contract. The following command will deploy the contract using the root account with account-0 and account-1 as the initial set of owners:

    $ ./briskitall.sh test deploy multisig $(cat ../testdata/account-0.addr) $(cat ../testdata/account-1.addr) --sender-key-file ../testdata/root.key
    ETH[0x4354d30f214b930e153ddaa6bac05ed210b8fd19e7c5e0abc351be559780bb10]: Sent 📄 
    ETH[0x4354d30f214b930e153ddaa6bac05ed210b8fd19e7c5e0abc351be559780bb10]: Confirmed ✅ 
    MultiSig contract address: 0x7A35a1584FDD8c88B0Fe60f21199CF6eEeCAA0fe

For convenience, you can export the `BRISKATALL_MULTISIG_CONTRACT_ADDRESS` variable with the contract address output by the deploy command. Alternatively you can pass it via the `--multisig-contract-address` flag to all commands that need it:

    $ export BRISKITALL_MULTISIG_CONTRACT_ADDRESS=0x7A35a1584FDD8c88B0Fe60f21199CF6eEeCAA0fe

Then to check that everything is working, you can list the owners:

    $ ./briskitall.sh query multisig owner list
    0x46f40B6B0dFDa35A8b6247489669a83c69804F3a
    0xBA4e70c2dc335aa86c6BF55F80d631Cf846435F0

## Deploy the Token contract

Now the token contract can be deployed with the MultiSig contract as the owner:

    $ ./briskitall.sh test deploy token "${BRISKITALL_MULTISIG_CONTRACT_ADDRESS}" --sender $(cat ../testdata/root.addr)
    ETH[0x2fec25f4d138267e00b1b95fe9bd6c2691ae4d6ce334dbc2f1be9106cba3bcdf]: Sent 📄 
    ETH[0x2fec25f4d138267e00b1b95fe9bd6c2691ae4d6ce334dbc2f1be9106cba3bcdf]: Confirmed ✅ 
    Token contract address: 0x9AB60992B6257c65A89D141f24d4ec7BE5241366

Once again, for convenience, you can export the `BRISKATALL_TOKEN_CONTRACT_ADDRESS` variable with the contract address output by the deploy command. Alternatively you can pass it via the `--token-contract-address` flag to all commands that need it:

    $ export BRISKITALL_TOKEN_CONTRACT_ADDRESS=0x9AB60992B6257c65A89D141f24d4ec7BE5241366

Then to check that everything is working, you can list the token balance of the MultiSig contract:

    $ ./briskitall.sh query token balance "${BRISKITALL_MULTISIG_CONTRACT_ADDRESS}"
    10000000000

# Tear Down

- Stop the docker containers (i.e. `docker compose down`)
- Delete the runtime directory (`rm -rf ./runtime`)
