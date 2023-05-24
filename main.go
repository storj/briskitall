package main

import (
	"context"
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
	env.Run(ctx, commands)
}

func commands(cmds clingy.Commands) {
	cmds.Group("query", "Query commands", func() {
		cmds.Group("multisig", "Query the multisig contract", func() {
			cmds.New("requirement", "Query the multisig contract confirmation requirement", new(cmdQueryMultiSigRequirement))
			cmds.Group("owner", "Query multisig contract ownership", func() {
				cmds.New("list", "List owners", new(cmdQueryMultiSigOwnerList))
			})
			cmds.Group("transaction", "Query multisig transactions", func() {
				cmds.New("list", "Lists the transactions", new(cmdQueryMultiSigTransactionList))
				cmds.New("status", "Retrieves transaction status", new(cmdQueryMultiSigTransactionStatus))
			})
		})
		cmds.Group("token", "Query the token contract", func() {
			cmds.New("allowance", "Retrieves the token allowance of an address", new(cmdQueryTokenAllowance))
			cmds.New("balance", "Retrieves the token balance of an address", new(cmdQueryTokenBalance))
			cmds.New("upgrade-master", "Retrieves the token upgrade master", new(cmdQueryTokenUpgradeMaster))
		})
	})

	cmds.Break()
	cmds.Group("submit", "Submit transaction to the multisig contract", func() {
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
		cmds.New("call", "Submits a transaction to execute a call to an arbitrary contract", new(cmdSubmitContractCall))
	})
	cmds.New("confirm", "Confirm a pending transaction", new(cmdConfirm))
	cmds.New("execute", "Execute a confirmed transaction", new(cmdExecute))
	cmds.New("revoke", "Revoke confirmation on a pending transaction", new(cmdRevoke))

	cmds.Break()
	cmds.Group("debug", "Submit transaction to the multisig contract", func() {
		cmds.Group("deploy", "Submit transaction to the multisig contract", func() {
			cmds.New("multisig", "Deploys the MultiSigWalletWithDailyLimit contract", new(cmdDebugDeployMultiSig))
			cmds.New("token", "Deploys the CentrallyIssuedToken contract", new(cmdDebugDeployToken))
		})
	})
}
