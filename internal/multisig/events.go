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

	blockNumber() uint64
}

type confirmationEvent contract.MultiSigWalletConfirmation

func (e *confirmationEvent) String() string {
	return fmt.Sprintf("Confirmation(%s,%s)", e.Sender, e.TransactionId)
}

func (e *confirmationEvent) blockNumber() uint64 { return e.Raw.BlockNumber }

type revocationEvent contract.MultiSigWalletRevocation

func (e *revocationEvent) String() string {
	return fmt.Sprintf("Revocation(%s,%s)", e.Sender, e.TransactionId)
}

func (e *revocationEvent) blockNumber() uint64 { return e.Raw.BlockNumber }

type submissionEvent contract.MultiSigWalletSubmission

func (e *submissionEvent) String() string {
	return fmt.Sprintf("Submission(%s)", e.TransactionId)
}

func (e *submissionEvent) blockNumber() uint64 { return e.Raw.BlockNumber }

type executionEvent contract.MultiSigWalletExecution

func (e *executionEvent) String() string {
	return fmt.Sprintf("Execution(%s)", e.TransactionId)
}

func (e *executionEvent) blockNumber() uint64 { return e.Raw.BlockNumber }

type executionFailureEvent contract.MultiSigWalletExecutionFailure

func (e *executionFailureEvent) String() string {
	return fmt.Sprintf("ExecutionFailure(%s)", e.TransactionId)
}

func (e *executionFailureEvent) blockNumber() uint64 { return e.Raw.BlockNumber }

type depositEvent contract.MultiSigWalletDeposit

func (e *depositEvent) String() string {
	return fmt.Sprintf("Deposit(%s, %s)", e.Sender, e.Value)
}

func (e *depositEvent) blockNumber() uint64 { return e.Raw.BlockNumber }

type ownerAdditionEvent contract.MultiSigWalletOwnerAddition

func (e *ownerAdditionEvent) String() string {
	return fmt.Sprintf("OwnerAddition(%s)", e.Owner)
}

func (e *ownerAdditionEvent) blockNumber() uint64 { return e.Raw.BlockNumber }

type ownerRemovalEvent contract.MultiSigWalletOwnerRemoval

func (e *ownerRemovalEvent) String() string {
	return fmt.Sprintf("OwnerRemoval(%s)", e.Owner)
}

func (e *ownerRemovalEvent) blockNumber() uint64 { return e.Raw.BlockNumber }

type requirementChangeEvent contract.MultiSigWalletRequirementChange

func (e *requirementChangeEvent) String() string {
	return fmt.Sprintf("RequirementChanged(%s)", e.Required)
}

func (e *requirementChangeEvent) blockNumber() uint64 { return e.Raw.BlockNumber }

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
		return events[i].blockNumber() < events[j].blockNumber()
	})
}
