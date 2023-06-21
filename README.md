# Brisk-It-All

## Name?

We needed a name for a multisig app, and so naturally we opened the Denny's Menu. Brisk-It-All seemed right:

> Slow-cooked brisket burnt ends, crispy diced bacon, two eggs*, aged white cheddar, Diner Q sauce and pickles on grilled artisan bread. Served with wavy-cut fries.

## Summary

```
$ briskitall --summary
Available commands:
    briskitall query eth balance                     Retrieve the ETH balance of an address
    briskitall query multisig requirement            Query the multisig contract confirmation requirement
    briskitall query multisig owner list             List multisig contract owners
    briskitall query multisig transaction list       List multisig transactions
    briskitall query multisig transaction status     Retrieve multisig transaction status
    briskitall query token allowance                 Retrieve the token allowance of an address
    briskitall query token balance                   Retrieve the token balance of an address
    briskitall query token upgrade-master            Retrieve the token upgrade master
    briskitall query usb-wallet account list         List USB wallet accounts
    briskitall submit eth transfer                   Submit a transaction to transfer ETH
    briskitall submit multisig requirement change    Submit a transaction to change the confirmation requirement
    briskitall submit multisig owner add             Submit a transaction to add an owner
    briskitall submit multisig owner remove          Submit a transaction to remove an owner
    briskitall submit multisig owner replace         Submit a transaction to replace an owner
    briskitall submit token approve                  Submit a transaction to approve a spender for an amount
    briskitall submit token set-upgrade-master       Submit a transaction to set the upgrade master
    briskitall submit token transfer                 Submit a transaction to transfer tokens
    briskitall submit token transfer-from            Submit a transaction to transfer tokens based on allowance
    briskitall submit call                           Submit a transaction to execute a call to an arbitrary contract
    briskitall confirm                               Confirm a pending transaction
    briskitall execute                               Execute a confirmed transaction
    briskitall revoke                                Revoke confirmation on a pending transaction
    briskitall test deploy multisig                  Deploy the MultiSigWalletWithDailyLimit contract
    briskitall test deploy token                     Deploy the CentrallyIssuedToken contract
    briskitall test eth transfer                     Transfer ETH
```

## Private Keys

Brisk-It-All can use private keys from the following:

1. Private key files on disk (must be mode 0600):
    ```
    briskitall some command --sender-key-file <path-to-key> ...
    or
    BRISTKITALL_SENDER_KEY_FILE=<path-to-key> briskitall some command ...
    ```

2. Via USB hardware wallets:
    ```
    briskitall some command --usb-wallet-account=<address>...
    or
    BRISKITALL_USB_WALLET_ACCOUNT=<address> briskitall some command ...
    ```

## End-to-End Testing

Instructions are available [here](./test/e2e).

## Risks

> EGGS SERVED OVER-EASY, POACHED, SUNNY-SIDE-UP OR SOFT-BOILED AND STEAKS THAT ARE SERVED RARE OR MEDIUM-RARE MAY BE UNDERCOOKED AND WILL ONLY BE SERVED UPON THE CONSUMERS’ REQUEST. NOTICE: CONSUMING RAW OR UNDERCOOKED MEATS, POULTRY, SEAFOOD, SHELLFISH OR EGGS MAY INCREASE YOUR RISK OF FOODBORNE ILLNESS, ESPECIALLY IF YOU HAVE CERTAIN MEDICAL CONDITIONS.
