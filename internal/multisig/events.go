package multisig

import (
	"fmt"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/zeebo/errs"

	"storj.io/briskitall/internal/contract"
)

type Event interface {
	String() string

	BlockNumber() uint64
}

type confirmationEvent contract.MultiSigWalletConfirmation

func (e *confirmationEvent) String() string {
	return fmt.Sprintf("ETH[%s]: Confirmation(%s)", e.Raw.TxHash, e.Sender)
}

func (e *confirmationEvent) BlockNumber() uint64 { return e.Raw.BlockNumber }

type revocationEvent contract.MultiSigWalletRevocation

func (e *revocationEvent) String() string {
	return fmt.Sprintf("ETH[%s]: Revocation(%s)", e.Raw.TxHash, e.Sender)
}

func (e *revocationEvent) BlockNumber() uint64 { return e.Raw.BlockNumber }

type submissionEvent contract.MultiSigWalletSubmission

func (e *submissionEvent) String() string {
	return fmt.Sprintf("ETH[%s]: Submission()", e.Raw.TxHash)
}

func (e *submissionEvent) BlockNumber() uint64 { return e.Raw.BlockNumber }

type executionEvent contract.MultiSigWalletExecution

func (e *executionEvent) String() string {
	return fmt.Sprintf("ETH[%s]: Execution()", e.Raw.TxHash)
}

func (e *executionEvent) BlockNumber() uint64 { return e.Raw.BlockNumber }

type executionFailureEvent contract.MultiSigWalletExecutionFailure

func (e *executionFailureEvent) String() string {
	return fmt.Sprintf("ETH[%s]: ExecutionFailure()", e.Raw.TxHash)
}

func (e *executionFailureEvent) BlockNumber() uint64 { return e.Raw.BlockNumber }

type depositEvent contract.MultiSigWalletDeposit

func (e *depositEvent) String() string {
	return fmt.Sprintf("ETH[%s]: Deposit(%s, %s)", e.Raw.TxHash, e.Sender, e.Value)
}

func (e *depositEvent) BlockNumber() uint64 { return e.Raw.BlockNumber }

type ownerAdditionEvent contract.MultiSigWalletOwnerAddition

func (e *ownerAdditionEvent) String() string {
	return fmt.Sprintf("ETH[%s]: OwnerAddition(%s)", e.Raw.TxHash, e.Owner)
}

func (e *ownerAdditionEvent) BlockNumber() uint64 { return e.Raw.BlockNumber }

type ownerRemovalEvent contract.MultiSigWalletOwnerRemoval

func (e *ownerRemovalEvent) String() string {
	return fmt.Sprintf("ETH[%s]: OwnerRemoval(%s)", e.Raw.TxHash, e.Owner)
}

func (e *ownerRemovalEvent) BlockNumber() uint64 { return e.Raw.BlockNumber }

type requirementChangeEvent contract.MultiSigWalletRequirementChange

func (e *requirementChangeEvent) String() string {
	return fmt.Sprintf("ETH[%s]: RequirementChanged(%s)", e.Raw.TxHash, e.Required)
}

func (e *requirementChangeEvent) BlockNumber() uint64 { return e.Raw.BlockNumber }

func getAllEvents(opts *bind.FilterOpts, filterer *contract.MultiSigWalletWithDailyLimitFilterer) ([]Event, error) {
	events, err := getTransactionEvents(opts, filterer, nil)
	if err != nil {
		return nil, err
	}

	{
		it, err := filterer.FilterDeposit(opts, nil)
		if err != nil {
			return nil, errs.New("failed to filter deposit events: %v", err)
		}
		for it.Next() {
			events = append(events, (*depositEvent)(it.Event))
		}
		if err != nil {
			return nil, errs.New("failed to iterate deposit events: %v", err)
		}
	}

	{
		it, err := filterer.FilterOwnerAddition(opts, nil)
		if err != nil {
			return nil, errs.New("failed to filter owner addition events: %v", err)
		}
		for it.Next() {
			events = append(events, (*ownerAdditionEvent)(it.Event))
		}
		if err != nil {
			return nil, errs.New("failed to iterate owner addition events: %v", err)
		}
	}

	{
		it, err := filterer.FilterOwnerRemoval(opts, nil)
		if err != nil {
			return nil, errs.New("failed to filter owner removal events: %v", err)
		}
		for it.Next() {
			events = append(events, (*ownerRemovalEvent)(it.Event))
		}
		if err != nil {
			return nil, errs.New("failed to iterate owner removal events: %v", err)
		}
	}

	{
		it, err := filterer.FilterRequirementChange(opts)
		if err != nil {
			return nil, errs.New("failed to filter requirement change events: %v", err)
		}
		for it.Next() {
			events = append(events, (*requirementChangeEvent)(it.Event))
		}
		if err != nil {
			return nil, errs.New("failed to iterate requirement change events: %v", err)
		}
	}

	return events, nil
}

func getTransactionEvents(opts *bind.FilterOpts, filterer *contract.MultiSigWalletWithDailyLimitFilterer, transactionIDs []*big.Int) ([]Event, error) {
	var events []Event
	{
		it, err := filterer.FilterConfirmation(opts, nil, transactionIDs)
		if err != nil {
			return nil, errs.New("failed to filter confirmation events: %v", err)
		}
		for it.Next() {
			events = append(events, (*confirmationEvent)(it.Event))
		}
		if err != nil {
			return nil, errs.New("failed to iterate confirmation events: %v", err)
		}
	}

	{
		it, err := filterer.FilterRevocation(opts, nil, transactionIDs)
		if err != nil {
			return nil, errs.New("failed to filter revocation events: %v", err)
		}
		for it.Next() {
			events = append(events, (*revocationEvent)(it.Event))
		}
		if err != nil {
			return nil, errs.New("failed to iterate revocation events: %v", err)
		}
	}

	{
		it, err := filterer.FilterSubmission(opts, transactionIDs)
		if err != nil {
			return nil, errs.New("failed to filter submission events: %v", err)
		}
		for it.Next() {
			events = append(events, (*submissionEvent)(it.Event))
		}
		if err != nil {
			return nil, errs.New("failed to iterate submission events: %v", err)
		}
	}

	{
		it, err := filterer.FilterExecution(opts, transactionIDs)
		if err != nil {
			return nil, errs.New("failed to filter execution events: %v", err)
		}
		for it.Next() {
			events = append(events, (*executionEvent)(it.Event))
		}
		if err != nil {
			return nil, errs.New("failed to iterate execution events: %v", err)
		}
	}

	{
		it, err := filterer.FilterExecutionFailure(opts, transactionIDs)
		if err != nil {
			return nil, errs.New("failed to filter execution failure events: %v", err)
		}
		for it.Next() {
			events = append(events, (*executionFailureEvent)(it.Event))
		}
		if err != nil {
			return nil, errs.New("failed to iterate execution failure events: %v", err)
		}
	}

	return events, nil
}

func sortEvents(events []Event) {
	sort.Slice(events, func(i, j int) bool {
		return events[i].BlockNumber() < events[j].BlockNumber()
	})
}
