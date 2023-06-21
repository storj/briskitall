package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/zeebo/clingy"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	env := clingy.Environment{
		Name: "briskitall",
	}
	ok, err := env.Run(ctx, commands)
	switch {
	case err != nil:
		fmt.Fprintln(os.Stderr, "Error:", err.Error())
		os.Exit(2)
	case !ok:
		os.Exit(1)
	}
}

func commands(cmds clingy.Commands) {
	cmds.Group("query", "Query commands", func() {
		cmds.Group("eth", "Query ETH", func() {
			cmds.New("balance", "Retrieve the ETH balance of an address", new(cmdQueryETHBalance))
		})
		cmds.Group("multisig", "Query the multisig contract", func() {
			cmds.New("requirement", "Query the multisig contract confirmation requirement", new(cmdQueryMultiSigRequirement))
			cmds.Group("owner", "Query multisig contract ownership", func() {
				cmds.New("list", "List multisig contract owners", new(cmdQueryMultiSigOwnerList))
			})
			cmds.Group("transaction", "Query multisig transactions", func() {
				cmds.New("list", "List multisig transactions", new(cmdQueryMultiSigTransactionList))
				cmds.New("status", "Retrieve multisig transaction status", new(cmdQueryMultiSigTransactionStatus))
			})
		})
		cmds.Group("token", "Query the token contract", func() {
			cmds.New("allowance", "Retrieve the token allowance of an address", new(cmdQueryTokenAllowance))
			cmds.New("balance", "Retrieve the token balance of an address", new(cmdQueryTokenBalance))
			cmds.New("upgrade-master", "Retrieve the token upgrade master", new(cmdQueryTokenUpgradeMaster))
		})
		cmds.Group("usb-wallet", "Query USB wallets", func() {
			cmds.Group("account", "Query USB wallet accounts", func() {
				cmds.New("list", "List USB wallet accounts", new(cmdQueryUSBWalletAccountList))
			})
		})
	})

	cmds.Break()
	cmds.Group("submit", "Submit transaction to the multisig contract", func() {
		cmds.Group("eth", "Submit ETH transactions", func() {
			cmds.New("transfer", "Submit a transaction to transfer ETH", new(cmdSubmitETHTransfer))
		})
		cmds.Group("multisig", "Submit multisig transactions", func() {
			cmds.Group("requirement", "Submit requirement transactions", func() {
				cmds.New("change", "Submit a transaction to change the confirmation requirement", new(cmdSubmitMultisigRequirementChange))
			})
			cmds.Group("owner", "Submit owner transactions", func() {
				cmds.New("add", "Submit a transaction to add an owner", new(cmdSubmitMultisigOwnerAdd))
				cmds.New("remove", "Submit a transaction to remove an owner", new(cmdSubmitMultisigOwnerRemove))
				cmds.New("replace", "Submit a transaction to replace an owner", new(cmdSubmitMultisigOwnerReplace))
			})
		})
		cmds.Group("token", "Submit token transactions", func() {
			cmds.New("approve", "Submit a transaction to approve a spender for an amount", new(cmdSubmitTokenApprove))
			cmds.New("set-upgrade-master", "Submit a transaction to set the upgrade master", new(cmdSubmitTokenSetUpgradeMaster))
			cmds.New("transfer", "Submit a transaction to transfer tokens", new(cmdSubmitTokenTransfer))
			cmds.New("transfer-from", "Submit a transaction to transfer tokens based on allowance", new(cmdSubmitTokenTransferFrom))
		})
		cmds.New("call", "Submit a transaction to execute a call to an arbitrary contract", new(cmdSubmitContractCall))
	})
	cmds.New("confirm", "Confirm a pending transaction", new(cmdConfirm))
	cmds.New("execute", "Execute a confirmed transaction", new(cmdExecute))
	cmds.New("revoke", "Revoke confirmation on a pending transaction", new(cmdRevoke))

	cmds.Break()
	cmds.Group("test", "Run test commands", func() {
		cmds.Group("deploy", "Submit transaction to the multisig contract", func() {
			cmds.New("multisig", "Deploy the MultiSigWalletWithDailyLimit contract", new(cmdTestDeployMultiSig))
			cmds.New("token", "Deploy the CentrallyIssuedToken contract", new(cmdTestDeployToken))
		})
		cmds.Group("eth", "Convenience ETH commands for testing", func() {
			cmds.New("transfer", "Transfer ETH", new(cmdTestETHTransfer))
		})
	})
}
