// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BurnableTokenMetaData contains all meta data concerning the BurnableToken contract.
var BurnableTokenMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"burnAmount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BURN_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"burnedAmount\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]",
	Bin: "0x6060604052341561000f57600080fd5b6106fe8061001e6000396000f3006060604052600436106100985763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461009d57806318160ddd146100e057806323b872dd1461010557806342966c681461013a57806370a0823114610152578063a9059cbb1461017e578063dd62ed3e146101ad578063eefa597b146101df578063fccc2813146101f2575b600080fd5b34156100a857600080fd5b6100cc73ffffffffffffffffffffffffffffffffffffffff6004351660243561022e565b604051901515815260200160405180910390f35b34156100eb57600080fd5b6100f36102f0565b60405190815260200160405180910390f35b341561011057600080fd5b6100cc73ffffffffffffffffffffffffffffffffffffffff600435811690602435166044356102f6565b341561014557600080fd5b61015060043561042b565b005b341561015d57600080fd5b6100f373ffffffffffffffffffffffffffffffffffffffff60043516610534565b341561018957600080fd5b6100cc73ffffffffffffffffffffffffffffffffffffffff6004351660243561055c565b34156101b857600080fd5b6100f373ffffffffffffffffffffffffffffffffffffffff60043581169060243516610649565b34156101ea57600080fd5b6100cc610681565b34156101fd57600080fd5b610205610686565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b60008115801590610270575073ffffffffffffffffffffffffffffffffffffffff33811660009081526002602090815260408083209387168352929052205415155b1561027a57600080fd5b73ffffffffffffffffffffffffffffffffffffffff338116600081815260026020908152604080832094881680845294909152908190208590557f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b73ffffffffffffffffffffffffffffffffffffffff8084166000908152600260209081526040808320338516845282528083205493861683526001909152812054909190610344908461068b565b73ffffffffffffffffffffffffffffffffffffffff808616600090815260016020526040808220939093559087168152205461038090846106af565b73ffffffffffffffffffffffffffffffffffffffff86166000908152600160205260409020556103b081846106af565b73ffffffffffffffffffffffffffffffffffffffff808716600081815260026020908152604080832033861684529091529081902093909355908616917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9086905190815260200160405180910390a3506001949350505050565b3373ffffffffffffffffffffffffffffffffffffffff811660009081526001602052604090205461045c90836106af565b73ffffffffffffffffffffffffffffffffffffffff82166000908152600160205260408120919091555461049090836106af565b6000557f696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7818360405173ffffffffffffffffffffffffffffffffffffffff909216825260208201526040908101905180910390a1600073ffffffffffffffffffffffffffffffffffffffff82167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff1660009081526001602052604090205490565b60006040604436101561056e57600080fd5b73ffffffffffffffffffffffffffffffffffffffff331660009081526001602052604090205461059e90846106af565b73ffffffffffffffffffffffffffffffffffffffff33811660009081526001602052604080822093909355908616815220546105da908461068b565b73ffffffffffffffffffffffffffffffffffffffff80861660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9086905190815260200160405180910390a35060019392505050565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260026020908152604080832093909416825291909152205490565b600181565b600081565b60008282016106a88482108015906106a35750838210155b6106c3565b9392505050565b60006106bd838311156106c3565b50900390565b8015156106cf57600080fd5b505600a165627a7a72305820be7c06ef28e94d5d33915e331d2f030c78dd98684efe160fbdad1595d1ea86a10029",
}

// BurnableTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use BurnableTokenMetaData.ABI instead.
var BurnableTokenABI = BurnableTokenMetaData.ABI

// BurnableTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BurnableTokenMetaData.Bin instead.
var BurnableTokenBin = BurnableTokenMetaData.Bin

// DeployBurnableToken deploys a new Ethereum contract, binding an instance of BurnableToken to it.
func DeployBurnableToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BurnableToken, error) {
	parsed, err := BurnableTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnableTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnableToken{BurnableTokenCaller: BurnableTokenCaller{contract: contract}, BurnableTokenTransactor: BurnableTokenTransactor{contract: contract}, BurnableTokenFilterer: BurnableTokenFilterer{contract: contract}}, nil
}

// BurnableToken is an auto generated Go binding around an Ethereum contract.
type BurnableToken struct {
	BurnableTokenCaller     // Read-only binding to the contract
	BurnableTokenTransactor // Write-only binding to the contract
	BurnableTokenFilterer   // Log filterer for contract events
}

// BurnableTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BurnableTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnableTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BurnableTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnableTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BurnableTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnableTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BurnableTokenSession struct {
	Contract     *BurnableToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BurnableTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BurnableTokenCallerSession struct {
	Contract *BurnableTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BurnableTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BurnableTokenTransactorSession struct {
	Contract     *BurnableTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BurnableTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BurnableTokenRaw struct {
	Contract *BurnableToken // Generic contract binding to access the raw methods on
}

// BurnableTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BurnableTokenCallerRaw struct {
	Contract *BurnableTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BurnableTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BurnableTokenTransactorRaw struct {
	Contract *BurnableTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBurnableToken creates a new instance of BurnableToken, bound to a specific deployed contract.
func NewBurnableToken(address common.Address, backend bind.ContractBackend) (*BurnableToken, error) {
	contract, err := bindBurnableToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnableToken{BurnableTokenCaller: BurnableTokenCaller{contract: contract}, BurnableTokenTransactor: BurnableTokenTransactor{contract: contract}, BurnableTokenFilterer: BurnableTokenFilterer{contract: contract}}, nil
}

// NewBurnableTokenCaller creates a new read-only instance of BurnableToken, bound to a specific deployed contract.
func NewBurnableTokenCaller(address common.Address, caller bind.ContractCaller) (*BurnableTokenCaller, error) {
	contract, err := bindBurnableToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenCaller{contract: contract}, nil
}

// NewBurnableTokenTransactor creates a new write-only instance of BurnableToken, bound to a specific deployed contract.
func NewBurnableTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnableTokenTransactor, error) {
	contract, err := bindBurnableToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenTransactor{contract: contract}, nil
}

// NewBurnableTokenFilterer creates a new log filterer instance of BurnableToken, bound to a specific deployed contract.
func NewBurnableTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnableTokenFilterer, error) {
	contract, err := bindBurnableToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenFilterer{contract: contract}, nil
}

// bindBurnableToken binds a generic wrapper to an already deployed contract.
func bindBurnableToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnableTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnableToken *BurnableTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnableToken.Contract.BurnableTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnableToken *BurnableTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnableToken.Contract.BurnableTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnableToken *BurnableTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnableToken.Contract.BurnableTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnableToken *BurnableTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnableToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnableToken *BurnableTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnableToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnableToken *BurnableTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnableToken.Contract.contract.Transact(opts, method, params...)
}

// BURNADDRESS is a free data retrieval call binding the contract method 0xfccc2813.
//
// Solidity: function BURN_ADDRESS() view returns(address)
func (_BurnableToken *BurnableTokenCaller) BURNADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnableToken.contract.Call(opts, &out, "BURN_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNADDRESS is a free data retrieval call binding the contract method 0xfccc2813.
//
// Solidity: function BURN_ADDRESS() view returns(address)
func (_BurnableToken *BurnableTokenSession) BURNADDRESS() (common.Address, error) {
	return _BurnableToken.Contract.BURNADDRESS(&_BurnableToken.CallOpts)
}

// BURNADDRESS is a free data retrieval call binding the contract method 0xfccc2813.
//
// Solidity: function BURN_ADDRESS() view returns(address)
func (_BurnableToken *BurnableTokenCallerSession) BURNADDRESS() (common.Address, error) {
	return _BurnableToken.Contract.BURNADDRESS(&_BurnableToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_BurnableToken *BurnableTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BurnableToken.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_BurnableToken *BurnableTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _BurnableToken.Contract.Allowance(&_BurnableToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_BurnableToken *BurnableTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _BurnableToken.Contract.Allowance(&_BurnableToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_BurnableToken *BurnableTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BurnableToken.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_BurnableToken *BurnableTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BurnableToken.Contract.BalanceOf(&_BurnableToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_BurnableToken *BurnableTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BurnableToken.Contract.BalanceOf(&_BurnableToken.CallOpts, _owner)
}

// IsToken is a free data retrieval call binding the contract method 0xeefa597b.
//
// Solidity: function isToken() view returns(bool)
func (_BurnableToken *BurnableTokenCaller) IsToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BurnableToken.contract.Call(opts, &out, "isToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsToken is a free data retrieval call binding the contract method 0xeefa597b.
//
// Solidity: function isToken() view returns(bool)
func (_BurnableToken *BurnableTokenSession) IsToken() (bool, error) {
	return _BurnableToken.Contract.IsToken(&_BurnableToken.CallOpts)
}

// IsToken is a free data retrieval call binding the contract method 0xeefa597b.
//
// Solidity: function isToken() view returns(bool)
func (_BurnableToken *BurnableTokenCallerSession) IsToken() (bool, error) {
	return _BurnableToken.Contract.IsToken(&_BurnableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BurnableToken *BurnableTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnableToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BurnableToken *BurnableTokenSession) TotalSupply() (*big.Int, error) {
	return _BurnableToken.Contract.TotalSupply(&_BurnableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BurnableToken *BurnableTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BurnableToken.Contract.TotalSupply(&_BurnableToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_BurnableToken *BurnableTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_BurnableToken *BurnableTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.Approve(&_BurnableToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_BurnableToken *BurnableTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.Approve(&_BurnableToken.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 burnAmount) returns()
func (_BurnableToken *BurnableTokenTransactor) Burn(opts *bind.TransactOpts, burnAmount *big.Int) (*types.Transaction, error) {
	return _BurnableToken.contract.Transact(opts, "burn", burnAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 burnAmount) returns()
func (_BurnableToken *BurnableTokenSession) Burn(burnAmount *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.Burn(&_BurnableToken.TransactOpts, burnAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 burnAmount) returns()
func (_BurnableToken *BurnableTokenTransactorSession) Burn(burnAmount *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.Burn(&_BurnableToken.TransactOpts, burnAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_BurnableToken *BurnableTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_BurnableToken *BurnableTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.Transfer(&_BurnableToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_BurnableToken *BurnableTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.Transfer(&_BurnableToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_BurnableToken *BurnableTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_BurnableToken *BurnableTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.TransferFrom(&_BurnableToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_BurnableToken *BurnableTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.TransferFrom(&_BurnableToken.TransactOpts, _from, _to, _value)
}

// BurnableTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BurnableToken contract.
type BurnableTokenApprovalIterator struct {
	Event *BurnableTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BurnableTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnableTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BurnableTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BurnableTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnableTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnableTokenApproval represents a Approval event raised by the BurnableToken contract.
type BurnableTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BurnableToken *BurnableTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BurnableTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnableToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenApprovalIterator{contract: _BurnableToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BurnableToken *BurnableTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BurnableTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnableToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnableTokenApproval)
				if err := _BurnableToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BurnableToken *BurnableTokenFilterer) ParseApproval(log types.Log) (*BurnableTokenApproval, error) {
	event := new(BurnableTokenApproval)
	if err := _BurnableToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BurnableTokenBurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the BurnableToken contract.
type BurnableTokenBurnedIterator struct {
	Event *BurnableTokenBurned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BurnableTokenBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnableTokenBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BurnableTokenBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BurnableTokenBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnableTokenBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnableTokenBurned represents a Burned event raised by the BurnableToken contract.
type BurnableTokenBurned struct {
	Burner       common.Address
	BurnedAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: event Burned(address burner, uint256 burnedAmount)
func (_BurnableToken *BurnableTokenFilterer) FilterBurned(opts *bind.FilterOpts) (*BurnableTokenBurnedIterator, error) {

	logs, sub, err := _BurnableToken.contract.FilterLogs(opts, "Burned")
	if err != nil {
		return nil, err
	}
	return &BurnableTokenBurnedIterator{contract: _BurnableToken.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: event Burned(address burner, uint256 burnedAmount)
func (_BurnableToken *BurnableTokenFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *BurnableTokenBurned) (event.Subscription, error) {

	logs, sub, err := _BurnableToken.contract.WatchLogs(opts, "Burned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnableTokenBurned)
				if err := _BurnableToken.contract.UnpackLog(event, "Burned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurned is a log parse operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: event Burned(address burner, uint256 burnedAmount)
func (_BurnableToken *BurnableTokenFilterer) ParseBurned(log types.Log) (*BurnableTokenBurned, error) {
	event := new(BurnableTokenBurned)
	if err := _BurnableToken.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BurnableTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BurnableToken contract.
type BurnableTokenTransferIterator struct {
	Event *BurnableTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BurnableTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnableTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BurnableTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BurnableTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnableTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnableTokenTransfer represents a Transfer event raised by the BurnableToken contract.
type BurnableTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BurnableToken *BurnableTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnableTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnableToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenTransferIterator{contract: _BurnableToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BurnableToken *BurnableTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnableTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnableToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnableTokenTransfer)
				if err := _BurnableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BurnableToken *BurnableTokenFilterer) ParseTransfer(log types.Log) (*BurnableTokenTransfer, error) {
	event := new(BurnableTokenTransfer)
	if err := _BurnableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CentrallyIssuedTokenMetaData contains all meta data concerning the CentrallyIssuedToken contract.
var CentrallyIssuedTokenMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"burnAmount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"upgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"upgradeAgent\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"upgradeMaster\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUpgradeState\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canUpgrade\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalUpgraded\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"setUpgradeAgent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BURN_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"master\",\"type\":\"address\"}],\"name\":\"setUpgradeMaster\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"name\":\"_decimals\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Upgrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"UpgradeAgentSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"burnedAmount\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]",
	Bin: "0x606060405234156200001057600080fd5b604051620010bc380380620010bc83398101604052808051919060200180518201919060200180518201919060200180519190602001805160038054600160a060020a031916600160a060020a0389161790559150600690508480516200007c929160200190620000c2565b50600783805162000092929160200190620000c2565b506000828155600891909155600160a060020a039094168452600160205260409093209290925550620001679050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200010557805160ff191683800117855562000135565b8280016001018555821562000135579182015b828111156200013557825182559160200191906001019062000118565b506200014392915062000147565b5090565b6200016491905b808211156200014357600081556001016200014e565b90565b610f4580620001776000396000f3006060604052600436106101115763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde038114610116578063095ea7b3146101a057806318160ddd146101e357806323b872dd14610208578063313ce5671461023d57806342966c681461025057806345977d03146102685780635de4ccb01461027e578063600440cb146102ba57806370a08231146102cd5780638444b391146102f957806395d89b41146103305780639738968c14610343578063a9059cbb14610356578063c752ff6214610385578063d7e7088a14610398578063dd62ed3e146103c4578063eefa597b146103f6578063fccc281314610409578063ffeb7d751461041c575b600080fd5b341561012157600080fd5b610129610448565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561016557808201518382015260200161014d565b50505050905090810190601f1680156101925780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101ab57600080fd5b6101cf73ffffffffffffffffffffffffffffffffffffffff600435166024356104e6565b604051901515815260200160405180910390f35b34156101ee57600080fd5b6101f66105a8565b60405190815260200160405180910390f35b341561021357600080fd5b6101cf73ffffffffffffffffffffffffffffffffffffffff600435811690602435166044356105ae565b341561024857600080fd5b6101f66106e3565b341561025b57600080fd5b6102666004356106e9565b005b341561027357600080fd5b6102666004356107f2565b341561028957600080fd5b6102916109a3565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156102c557600080fd5b6102916109bf565b34156102d857600080fd5b6101f673ffffffffffffffffffffffffffffffffffffffff600435166109db565b341561030457600080fd5b61030c610a03565b6040518082600481111561031c57fe5b60ff16815260200191505060405180910390f35b341561033b57600080fd5b610129610a5a565b341561034e57600080fd5b6101cf610ac5565b341561036157600080fd5b6101cf73ffffffffffffffffffffffffffffffffffffffff60043516602435610aca565b341561039057600080fd5b6101f6610bb7565b34156103a357600080fd5b61026673ffffffffffffffffffffffffffffffffffffffff60043516610bbd565b34156103cf57600080fd5b6101f673ffffffffffffffffffffffffffffffffffffffff60043581169060243516610dff565b341561040157600080fd5b6101cf610e37565b341561041457600080fd5b610291610e3c565b341561042757600080fd5b61026673ffffffffffffffffffffffffffffffffffffffff60043516610e41565b60068054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104de5780601f106104b3576101008083540402835291602001916104de565b820191906000526020600020905b8154815290600101906020018083116104c157829003601f168201915b505050505081565b60008115801590610528575073ffffffffffffffffffffffffffffffffffffffff33811660009081526002602090815260408083209387168352929052205415155b1561053257600080fd5b73ffffffffffffffffffffffffffffffffffffffff338116600081815260026020908152604080832094881680845294909152908190208590557f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b73ffffffffffffffffffffffffffffffffffffffff80841660009081526002602090815260408083203385168452825280832054938616835260019091528120549091906105fc9084610ed2565b73ffffffffffffffffffffffffffffffffffffffff80861660009081526001602052604080822093909355908716815220546106389084610ef6565b73ffffffffffffffffffffffffffffffffffffffff86166000908152600160205260409020556106688184610ef6565b73ffffffffffffffffffffffffffffffffffffffff808716600081815260026020908152604080832033861684529091529081902093909355908616917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9086905190815260200160405180910390a3506001949350505050565b60085481565b3373ffffffffffffffffffffffffffffffffffffffff811660009081526001602052604090205461071a9083610ef6565b73ffffffffffffffffffffffffffffffffffffffff82166000908152600160205260408120919091555461074e9083610ef6565b6000557f696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7818360405173ffffffffffffffffffffffffffffffffffffffff909216825260208201526040908101905180910390a1600073ffffffffffffffffffffffffffffffffffffffff82167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a35050565b60006107fc610a03565b9050600381600481111561080c57fe5b14806108235750600481600481111561082157fe5b145b151561082e57600080fd5b81151561083a57600080fd5b73ffffffffffffffffffffffffffffffffffffffff331660009081526001602052604090205461086a9083610ef6565b73ffffffffffffffffffffffffffffffffffffffff33166000908152600160205260408120919091555461089e9083610ef6565b6000556005546108ae9083610ed2565b60055560045473ffffffffffffffffffffffffffffffffffffffff1663753e88e533846040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815273ffffffffffffffffffffffffffffffffffffffff90921660048301526024820152604401600060405180830381600087803b151561093a57600080fd5b6102c65a03f1151561094b57600080fd5b505060045473ffffffffffffffffffffffffffffffffffffffff908116915033167f7e5c344a8141a805725cb476f76c6953b842222b967edd1f78ddb6e8b3f397ac8460405190815260200160405180910390a35050565b60045473ffffffffffffffffffffffffffffffffffffffff1681565b60035473ffffffffffffffffffffffffffffffffffffffff1681565b73ffffffffffffffffffffffffffffffffffffffff1660009081526001602052604090205490565b6000610a0d610ac5565b1515610a1b57506001610a57565b60045473ffffffffffffffffffffffffffffffffffffffff161515610a4257506002610a57565b6005541515610a5357506003610a57565b5060045b90565b60078054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104de5780601f106104b3576101008083540402835291602001916104de565b600190565b600060406044361015610adc57600080fd5b73ffffffffffffffffffffffffffffffffffffffff3316600090815260016020526040902054610b0c9084610ef6565b73ffffffffffffffffffffffffffffffffffffffff3381166000908152600160205260408082209390935590861681522054610b489084610ed2565b73ffffffffffffffffffffffffffffffffffffffff80861660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9086905190815260200160405180910390a35060019392505050565b60055481565b610bc5610ac5565b1515610bd057600080fd5b73ffffffffffffffffffffffffffffffffffffffff81161515610bf257600080fd5b6003543373ffffffffffffffffffffffffffffffffffffffff908116911614610c1a57600080fd5b6004610c24610a03565b6004811115610c2f57fe5b1415610c3a57600080fd5b600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169190911791829055166361d3d7a66000604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1515610cd657600080fd5b6102c65a03f11515610ce757600080fd5b505050604051805190501515610cfc57600080fd5b60008054600454909173ffffffffffffffffffffffffffffffffffffffff90911690634b2ba0dd90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1515610d7257600080fd5b6102c65a03f11515610d8357600080fd5b50505060405180519050141515610d9957600080fd5b6004547f7845d5aa74cc410e35571258d954f23b82276e160fe8c188fa80566580f279cc9073ffffffffffffffffffffffffffffffffffffffff1660405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a150565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260026020908152604080832093909416825291909152205490565b600181565b600081565b73ffffffffffffffffffffffffffffffffffffffff81161515610e6357600080fd5b6003543373ffffffffffffffffffffffffffffffffffffffff908116911614610e8b57600080fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6000828201610eef848210801590610eea5750838210155b610f0a565b9392505050565b6000610f0483831115610f0a565b50900390565b801515610f1657600080fd5b505600a165627a7a72305820c2cf36a7a6541eda5f1638c532027e097ed6c96c67e09a96cabc6eba7faff3920029",
}

// CentrallyIssuedTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use CentrallyIssuedTokenMetaData.ABI instead.
var CentrallyIssuedTokenABI = CentrallyIssuedTokenMetaData.ABI

// CentrallyIssuedTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CentrallyIssuedTokenMetaData.Bin instead.
var CentrallyIssuedTokenBin = CentrallyIssuedTokenMetaData.Bin

// DeployCentrallyIssuedToken deploys a new Ethereum contract, binding an instance of CentrallyIssuedToken to it.
func DeployCentrallyIssuedToken(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _name string, _symbol string, _totalSupply *big.Int, _decimals *big.Int) (common.Address, *types.Transaction, *CentrallyIssuedToken, error) {
	parsed, err := CentrallyIssuedTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CentrallyIssuedTokenBin), backend, _owner, _name, _symbol, _totalSupply, _decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CentrallyIssuedToken{CentrallyIssuedTokenCaller: CentrallyIssuedTokenCaller{contract: contract}, CentrallyIssuedTokenTransactor: CentrallyIssuedTokenTransactor{contract: contract}, CentrallyIssuedTokenFilterer: CentrallyIssuedTokenFilterer{contract: contract}}, nil
}

// CentrallyIssuedToken is an auto generated Go binding around an Ethereum contract.
type CentrallyIssuedToken struct {
	CentrallyIssuedTokenCaller     // Read-only binding to the contract
	CentrallyIssuedTokenTransactor // Write-only binding to the contract
	CentrallyIssuedTokenFilterer   // Log filterer for contract events
}

// CentrallyIssuedTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type CentrallyIssuedTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CentrallyIssuedTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CentrallyIssuedTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CentrallyIssuedTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CentrallyIssuedTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CentrallyIssuedTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CentrallyIssuedTokenSession struct {
	Contract     *CentrallyIssuedToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CentrallyIssuedTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CentrallyIssuedTokenCallerSession struct {
	Contract *CentrallyIssuedTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// CentrallyIssuedTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CentrallyIssuedTokenTransactorSession struct {
	Contract     *CentrallyIssuedTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// CentrallyIssuedTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type CentrallyIssuedTokenRaw struct {
	Contract *CentrallyIssuedToken // Generic contract binding to access the raw methods on
}

// CentrallyIssuedTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CentrallyIssuedTokenCallerRaw struct {
	Contract *CentrallyIssuedTokenCaller // Generic read-only contract binding to access the raw methods on
}

// CentrallyIssuedTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CentrallyIssuedTokenTransactorRaw struct {
	Contract *CentrallyIssuedTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCentrallyIssuedToken creates a new instance of CentrallyIssuedToken, bound to a specific deployed contract.
func NewCentrallyIssuedToken(address common.Address, backend bind.ContractBackend) (*CentrallyIssuedToken, error) {
	contract, err := bindCentrallyIssuedToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CentrallyIssuedToken{CentrallyIssuedTokenCaller: CentrallyIssuedTokenCaller{contract: contract}, CentrallyIssuedTokenTransactor: CentrallyIssuedTokenTransactor{contract: contract}, CentrallyIssuedTokenFilterer: CentrallyIssuedTokenFilterer{contract: contract}}, nil
}

// NewCentrallyIssuedTokenCaller creates a new read-only instance of CentrallyIssuedToken, bound to a specific deployed contract.
func NewCentrallyIssuedTokenCaller(address common.Address, caller bind.ContractCaller) (*CentrallyIssuedTokenCaller, error) {
	contract, err := bindCentrallyIssuedToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CentrallyIssuedTokenCaller{contract: contract}, nil
}

// NewCentrallyIssuedTokenTransactor creates a new write-only instance of CentrallyIssuedToken, bound to a specific deployed contract.
func NewCentrallyIssuedTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*CentrallyIssuedTokenTransactor, error) {
	contract, err := bindCentrallyIssuedToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CentrallyIssuedTokenTransactor{contract: contract}, nil
}

// NewCentrallyIssuedTokenFilterer creates a new log filterer instance of CentrallyIssuedToken, bound to a specific deployed contract.
func NewCentrallyIssuedTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*CentrallyIssuedTokenFilterer, error) {
	contract, err := bindCentrallyIssuedToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CentrallyIssuedTokenFilterer{contract: contract}, nil
}

// bindCentrallyIssuedToken binds a generic wrapper to an already deployed contract.
func bindCentrallyIssuedToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CentrallyIssuedTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CentrallyIssuedToken *CentrallyIssuedTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CentrallyIssuedToken.Contract.CentrallyIssuedTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CentrallyIssuedToken *CentrallyIssuedTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CentrallyIssuedToken.Contract.CentrallyIssuedTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CentrallyIssuedToken *CentrallyIssuedTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CentrallyIssuedToken.Contract.CentrallyIssuedTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CentrallyIssuedToken *CentrallyIssuedTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CentrallyIssuedToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CentrallyIssuedToken *CentrallyIssuedTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CentrallyIssuedToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CentrallyIssuedToken *CentrallyIssuedTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CentrallyIssuedToken.Contract.contract.Transact(opts, method, params...)
}

// BURNADDRESS is a free data retrieval call binding the contract method 0xfccc2813.
//
// Solidity: function BURN_ADDRESS() view returns(address)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCaller) BURNADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CentrallyIssuedToken.contract.Call(opts, &out, "BURN_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNADDRESS is a free data retrieval call binding the contract method 0xfccc2813.
//
// Solidity: function BURN_ADDRESS() view returns(address)
func (_CentrallyIssuedToken *CentrallyIssuedTokenSession) BURNADDRESS() (common.Address, error) {
	return _CentrallyIssuedToken.Contract.BURNADDRESS(&_CentrallyIssuedToken.CallOpts)
}

// BURNADDRESS is a free data retrieval call binding the contract method 0xfccc2813.
//
// Solidity: function BURN_ADDRESS() view returns(address)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCallerSession) BURNADDRESS() (common.Address, error) {
	return _CentrallyIssuedToken.Contract.BURNADDRESS(&_CentrallyIssuedToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CentrallyIssuedToken.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_CentrallyIssuedToken *CentrallyIssuedTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _CentrallyIssuedToken.Contract.Allowance(&_CentrallyIssuedToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _CentrallyIssuedToken.Contract.Allowance(&_CentrallyIssuedToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CentrallyIssuedToken.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_CentrallyIssuedToken *CentrallyIssuedTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _CentrallyIssuedToken.Contract.BalanceOf(&_CentrallyIssuedToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _CentrallyIssuedToken.Contract.BalanceOf(&_CentrallyIssuedToken.CallOpts, _owner)
}

// CanUpgrade is a free data retrieval call binding the contract method 0x9738968c.
//
// Solidity: function canUpgrade() view returns(bool)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCaller) CanUpgrade(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CentrallyIssuedToken.contract.Call(opts, &out, "canUpgrade")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanUpgrade is a free data retrieval call binding the contract method 0x9738968c.
//
// Solidity: function canUpgrade() view returns(bool)
func (_CentrallyIssuedToken *CentrallyIssuedTokenSession) CanUpgrade() (bool, error) {
	return _CentrallyIssuedToken.Contract.CanUpgrade(&_CentrallyIssuedToken.CallOpts)
}

// CanUpgrade is a free data retrieval call binding the contract method 0x9738968c.
//
// Solidity: function canUpgrade() view returns(bool)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCallerSession) CanUpgrade() (bool, error) {
	return _CentrallyIssuedToken.Contract.CanUpgrade(&_CentrallyIssuedToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CentrallyIssuedToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CentrallyIssuedToken *CentrallyIssuedTokenSession) Decimals() (*big.Int, error) {
	return _CentrallyIssuedToken.Contract.Decimals(&_CentrallyIssuedToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCallerSession) Decimals() (*big.Int, error) {
	return _CentrallyIssuedToken.Contract.Decimals(&_CentrallyIssuedToken.CallOpts)
}

// GetUpgradeState is a free data retrieval call binding the contract method 0x8444b391.
//
// Solidity: function getUpgradeState() view returns(uint8)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCaller) GetUpgradeState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CentrallyIssuedToken.contract.Call(opts, &out, "getUpgradeState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetUpgradeState is a free data retrieval call binding the contract method 0x8444b391.
//
// Solidity: function getUpgradeState() view returns(uint8)
func (_CentrallyIssuedToken *CentrallyIssuedTokenSession) GetUpgradeState() (uint8, error) {
	return _CentrallyIssuedToken.Contract.GetUpgradeState(&_CentrallyIssuedToken.CallOpts)
}

// GetUpgradeState is a free data retrieval call binding the contract method 0x8444b391.
//
// Solidity: function getUpgradeState() view returns(uint8)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCallerSession) GetUpgradeState() (uint8, error) {
	return _CentrallyIssuedToken.Contract.GetUpgradeState(&_CentrallyIssuedToken.CallOpts)
}

// IsToken is a free data retrieval call binding the contract method 0xeefa597b.
//
// Solidity: function isToken() view returns(bool)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCaller) IsToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CentrallyIssuedToken.contract.Call(opts, &out, "isToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsToken is a free data retrieval call binding the contract method 0xeefa597b.
//
// Solidity: function isToken() view returns(bool)
func (_CentrallyIssuedToken *CentrallyIssuedTokenSession) IsToken() (bool, error) {
	return _CentrallyIssuedToken.Contract.IsToken(&_CentrallyIssuedToken.CallOpts)
}

// IsToken is a free data retrieval call binding the contract method 0xeefa597b.
//
// Solidity: function isToken() view returns(bool)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCallerSession) IsToken() (bool, error) {
	return _CentrallyIssuedToken.Contract.IsToken(&_CentrallyIssuedToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CentrallyIssuedToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CentrallyIssuedToken *CentrallyIssuedTokenSession) Name() (string, error) {
	return _CentrallyIssuedToken.Contract.Name(&_CentrallyIssuedToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCallerSession) Name() (string, error) {
	return _CentrallyIssuedToken.Contract.Name(&_CentrallyIssuedToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CentrallyIssuedToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CentrallyIssuedToken *CentrallyIssuedTokenSession) Symbol() (string, error) {
	return _CentrallyIssuedToken.Contract.Symbol(&_CentrallyIssuedToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCallerSession) Symbol() (string, error) {
	return _CentrallyIssuedToken.Contract.Symbol(&_CentrallyIssuedToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CentrallyIssuedToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CentrallyIssuedToken *CentrallyIssuedTokenSession) TotalSupply() (*big.Int, error) {
	return _CentrallyIssuedToken.Contract.TotalSupply(&_CentrallyIssuedToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _CentrallyIssuedToken.Contract.TotalSupply(&_CentrallyIssuedToken.CallOpts)
}

// TotalUpgraded is a free data retrieval call binding the contract method 0xc752ff62.
//
// Solidity: function totalUpgraded() view returns(uint256)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCaller) TotalUpgraded(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CentrallyIssuedToken.contract.Call(opts, &out, "totalUpgraded")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalUpgraded is a free data retrieval call binding the contract method 0xc752ff62.
//
// Solidity: function totalUpgraded() view returns(uint256)
func (_CentrallyIssuedToken *CentrallyIssuedTokenSession) TotalUpgraded() (*big.Int, error) {
	return _CentrallyIssuedToken.Contract.TotalUpgraded(&_CentrallyIssuedToken.CallOpts)
}

// TotalUpgraded is a free data retrieval call binding the contract method 0xc752ff62.
//
// Solidity: function totalUpgraded() view returns(uint256)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCallerSession) TotalUpgraded() (*big.Int, error) {
	return _CentrallyIssuedToken.Contract.TotalUpgraded(&_CentrallyIssuedToken.CallOpts)
}

// UpgradeAgent is a free data retrieval call binding the contract method 0x5de4ccb0.
//
// Solidity: function upgradeAgent() view returns(address)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCaller) UpgradeAgent(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CentrallyIssuedToken.contract.Call(opts, &out, "upgradeAgent")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpgradeAgent is a free data retrieval call binding the contract method 0x5de4ccb0.
//
// Solidity: function upgradeAgent() view returns(address)
func (_CentrallyIssuedToken *CentrallyIssuedTokenSession) UpgradeAgent() (common.Address, error) {
	return _CentrallyIssuedToken.Contract.UpgradeAgent(&_CentrallyIssuedToken.CallOpts)
}

// UpgradeAgent is a free data retrieval call binding the contract method 0x5de4ccb0.
//
// Solidity: function upgradeAgent() view returns(address)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCallerSession) UpgradeAgent() (common.Address, error) {
	return _CentrallyIssuedToken.Contract.UpgradeAgent(&_CentrallyIssuedToken.CallOpts)
}

// UpgradeMaster is a free data retrieval call binding the contract method 0x600440cb.
//
// Solidity: function upgradeMaster() view returns(address)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCaller) UpgradeMaster(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CentrallyIssuedToken.contract.Call(opts, &out, "upgradeMaster")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpgradeMaster is a free data retrieval call binding the contract method 0x600440cb.
//
// Solidity: function upgradeMaster() view returns(address)
func (_CentrallyIssuedToken *CentrallyIssuedTokenSession) UpgradeMaster() (common.Address, error) {
	return _CentrallyIssuedToken.Contract.UpgradeMaster(&_CentrallyIssuedToken.CallOpts)
}

// UpgradeMaster is a free data retrieval call binding the contract method 0x600440cb.
//
// Solidity: function upgradeMaster() view returns(address)
func (_CentrallyIssuedToken *CentrallyIssuedTokenCallerSession) UpgradeMaster() (common.Address, error) {
	return _CentrallyIssuedToken.Contract.UpgradeMaster(&_CentrallyIssuedToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_CentrallyIssuedToken *CentrallyIssuedTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CentrallyIssuedToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_CentrallyIssuedToken *CentrallyIssuedTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CentrallyIssuedToken.Contract.Approve(&_CentrallyIssuedToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_CentrallyIssuedToken *CentrallyIssuedTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CentrallyIssuedToken.Contract.Approve(&_CentrallyIssuedToken.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 burnAmount) returns()
func (_CentrallyIssuedToken *CentrallyIssuedTokenTransactor) Burn(opts *bind.TransactOpts, burnAmount *big.Int) (*types.Transaction, error) {
	return _CentrallyIssuedToken.contract.Transact(opts, "burn", burnAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 burnAmount) returns()
func (_CentrallyIssuedToken *CentrallyIssuedTokenSession) Burn(burnAmount *big.Int) (*types.Transaction, error) {
	return _CentrallyIssuedToken.Contract.Burn(&_CentrallyIssuedToken.TransactOpts, burnAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 burnAmount) returns()
func (_CentrallyIssuedToken *CentrallyIssuedTokenTransactorSession) Burn(burnAmount *big.Int) (*types.Transaction, error) {
	return _CentrallyIssuedToken.Contract.Burn(&_CentrallyIssuedToken.TransactOpts, burnAmount)
}

// SetUpgradeAgent is a paid mutator transaction binding the contract method 0xd7e7088a.
//
// Solidity: function setUpgradeAgent(address agent) returns()
func (_CentrallyIssuedToken *CentrallyIssuedTokenTransactor) SetUpgradeAgent(opts *bind.TransactOpts, agent common.Address) (*types.Transaction, error) {
	return _CentrallyIssuedToken.contract.Transact(opts, "setUpgradeAgent", agent)
}

// SetUpgradeAgent is a paid mutator transaction binding the contract method 0xd7e7088a.
//
// Solidity: function setUpgradeAgent(address agent) returns()
func (_CentrallyIssuedToken *CentrallyIssuedTokenSession) SetUpgradeAgent(agent common.Address) (*types.Transaction, error) {
	return _CentrallyIssuedToken.Contract.SetUpgradeAgent(&_CentrallyIssuedToken.TransactOpts, agent)
}

// SetUpgradeAgent is a paid mutator transaction binding the contract method 0xd7e7088a.
//
// Solidity: function setUpgradeAgent(address agent) returns()
func (_CentrallyIssuedToken *CentrallyIssuedTokenTransactorSession) SetUpgradeAgent(agent common.Address) (*types.Transaction, error) {
	return _CentrallyIssuedToken.Contract.SetUpgradeAgent(&_CentrallyIssuedToken.TransactOpts, agent)
}

// SetUpgradeMaster is a paid mutator transaction binding the contract method 0xffeb7d75.
//
// Solidity: function setUpgradeMaster(address master) returns()
func (_CentrallyIssuedToken *CentrallyIssuedTokenTransactor) SetUpgradeMaster(opts *bind.TransactOpts, master common.Address) (*types.Transaction, error) {
	return _CentrallyIssuedToken.contract.Transact(opts, "setUpgradeMaster", master)
}

// SetUpgradeMaster is a paid mutator transaction binding the contract method 0xffeb7d75.
//
// Solidity: function setUpgradeMaster(address master) returns()
func (_CentrallyIssuedToken *CentrallyIssuedTokenSession) SetUpgradeMaster(master common.Address) (*types.Transaction, error) {
	return _CentrallyIssuedToken.Contract.SetUpgradeMaster(&_CentrallyIssuedToken.TransactOpts, master)
}

// SetUpgradeMaster is a paid mutator transaction binding the contract method 0xffeb7d75.
//
// Solidity: function setUpgradeMaster(address master) returns()
func (_CentrallyIssuedToken *CentrallyIssuedTokenTransactorSession) SetUpgradeMaster(master common.Address) (*types.Transaction, error) {
	return _CentrallyIssuedToken.Contract.SetUpgradeMaster(&_CentrallyIssuedToken.TransactOpts, master)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_CentrallyIssuedToken *CentrallyIssuedTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CentrallyIssuedToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_CentrallyIssuedToken *CentrallyIssuedTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CentrallyIssuedToken.Contract.Transfer(&_CentrallyIssuedToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_CentrallyIssuedToken *CentrallyIssuedTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CentrallyIssuedToken.Contract.Transfer(&_CentrallyIssuedToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_CentrallyIssuedToken *CentrallyIssuedTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CentrallyIssuedToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_CentrallyIssuedToken *CentrallyIssuedTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CentrallyIssuedToken.Contract.TransferFrom(&_CentrallyIssuedToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_CentrallyIssuedToken *CentrallyIssuedTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CentrallyIssuedToken.Contract.TransferFrom(&_CentrallyIssuedToken.TransactOpts, _from, _to, _value)
}

// Upgrade is a paid mutator transaction binding the contract method 0x45977d03.
//
// Solidity: function upgrade(uint256 value) returns()
func (_CentrallyIssuedToken *CentrallyIssuedTokenTransactor) Upgrade(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _CentrallyIssuedToken.contract.Transact(opts, "upgrade", value)
}

// Upgrade is a paid mutator transaction binding the contract method 0x45977d03.
//
// Solidity: function upgrade(uint256 value) returns()
func (_CentrallyIssuedToken *CentrallyIssuedTokenSession) Upgrade(value *big.Int) (*types.Transaction, error) {
	return _CentrallyIssuedToken.Contract.Upgrade(&_CentrallyIssuedToken.TransactOpts, value)
}

// Upgrade is a paid mutator transaction binding the contract method 0x45977d03.
//
// Solidity: function upgrade(uint256 value) returns()
func (_CentrallyIssuedToken *CentrallyIssuedTokenTransactorSession) Upgrade(value *big.Int) (*types.Transaction, error) {
	return _CentrallyIssuedToken.Contract.Upgrade(&_CentrallyIssuedToken.TransactOpts, value)
}

// CentrallyIssuedTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CentrallyIssuedToken contract.
type CentrallyIssuedTokenApprovalIterator struct {
	Event *CentrallyIssuedTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CentrallyIssuedTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CentrallyIssuedTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CentrallyIssuedTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CentrallyIssuedTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CentrallyIssuedTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CentrallyIssuedTokenApproval represents a Approval event raised by the CentrallyIssuedToken contract.
type CentrallyIssuedTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CentrallyIssuedToken *CentrallyIssuedTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CentrallyIssuedTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CentrallyIssuedToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CentrallyIssuedTokenApprovalIterator{contract: _CentrallyIssuedToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CentrallyIssuedToken *CentrallyIssuedTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CentrallyIssuedTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CentrallyIssuedToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CentrallyIssuedTokenApproval)
				if err := _CentrallyIssuedToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CentrallyIssuedToken *CentrallyIssuedTokenFilterer) ParseApproval(log types.Log) (*CentrallyIssuedTokenApproval, error) {
	event := new(CentrallyIssuedTokenApproval)
	if err := _CentrallyIssuedToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CentrallyIssuedTokenBurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the CentrallyIssuedToken contract.
type CentrallyIssuedTokenBurnedIterator struct {
	Event *CentrallyIssuedTokenBurned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CentrallyIssuedTokenBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CentrallyIssuedTokenBurned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CentrallyIssuedTokenBurned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CentrallyIssuedTokenBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CentrallyIssuedTokenBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CentrallyIssuedTokenBurned represents a Burned event raised by the CentrallyIssuedToken contract.
type CentrallyIssuedTokenBurned struct {
	Burner       common.Address
	BurnedAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: event Burned(address burner, uint256 burnedAmount)
func (_CentrallyIssuedToken *CentrallyIssuedTokenFilterer) FilterBurned(opts *bind.FilterOpts) (*CentrallyIssuedTokenBurnedIterator, error) {

	logs, sub, err := _CentrallyIssuedToken.contract.FilterLogs(opts, "Burned")
	if err != nil {
		return nil, err
	}
	return &CentrallyIssuedTokenBurnedIterator{contract: _CentrallyIssuedToken.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: event Burned(address burner, uint256 burnedAmount)
func (_CentrallyIssuedToken *CentrallyIssuedTokenFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *CentrallyIssuedTokenBurned) (event.Subscription, error) {

	logs, sub, err := _CentrallyIssuedToken.contract.WatchLogs(opts, "Burned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CentrallyIssuedTokenBurned)
				if err := _CentrallyIssuedToken.contract.UnpackLog(event, "Burned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurned is a log parse operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: event Burned(address burner, uint256 burnedAmount)
func (_CentrallyIssuedToken *CentrallyIssuedTokenFilterer) ParseBurned(log types.Log) (*CentrallyIssuedTokenBurned, error) {
	event := new(CentrallyIssuedTokenBurned)
	if err := _CentrallyIssuedToken.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CentrallyIssuedTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CentrallyIssuedToken contract.
type CentrallyIssuedTokenTransferIterator struct {
	Event *CentrallyIssuedTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CentrallyIssuedTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CentrallyIssuedTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CentrallyIssuedTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CentrallyIssuedTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CentrallyIssuedTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CentrallyIssuedTokenTransfer represents a Transfer event raised by the CentrallyIssuedToken contract.
type CentrallyIssuedTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CentrallyIssuedToken *CentrallyIssuedTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CentrallyIssuedTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CentrallyIssuedToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CentrallyIssuedTokenTransferIterator{contract: _CentrallyIssuedToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CentrallyIssuedToken *CentrallyIssuedTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CentrallyIssuedTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CentrallyIssuedToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CentrallyIssuedTokenTransfer)
				if err := _CentrallyIssuedToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CentrallyIssuedToken *CentrallyIssuedTokenFilterer) ParseTransfer(log types.Log) (*CentrallyIssuedTokenTransfer, error) {
	event := new(CentrallyIssuedTokenTransfer)
	if err := _CentrallyIssuedToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CentrallyIssuedTokenUpgradeIterator is returned from FilterUpgrade and is used to iterate over the raw logs and unpacked data for Upgrade events raised by the CentrallyIssuedToken contract.
type CentrallyIssuedTokenUpgradeIterator struct {
	Event *CentrallyIssuedTokenUpgrade // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CentrallyIssuedTokenUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CentrallyIssuedTokenUpgrade)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CentrallyIssuedTokenUpgrade)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CentrallyIssuedTokenUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CentrallyIssuedTokenUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CentrallyIssuedTokenUpgrade represents a Upgrade event raised by the CentrallyIssuedToken contract.
type CentrallyIssuedTokenUpgrade struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpgrade is a free log retrieval operation binding the contract event 0x7e5c344a8141a805725cb476f76c6953b842222b967edd1f78ddb6e8b3f397ac.
//
// Solidity: event Upgrade(address indexed _from, address indexed _to, uint256 _value)
func (_CentrallyIssuedToken *CentrallyIssuedTokenFilterer) FilterUpgrade(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*CentrallyIssuedTokenUpgradeIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _CentrallyIssuedToken.contract.FilterLogs(opts, "Upgrade", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &CentrallyIssuedTokenUpgradeIterator{contract: _CentrallyIssuedToken.contract, event: "Upgrade", logs: logs, sub: sub}, nil
}

// WatchUpgrade is a free log subscription operation binding the contract event 0x7e5c344a8141a805725cb476f76c6953b842222b967edd1f78ddb6e8b3f397ac.
//
// Solidity: event Upgrade(address indexed _from, address indexed _to, uint256 _value)
func (_CentrallyIssuedToken *CentrallyIssuedTokenFilterer) WatchUpgrade(opts *bind.WatchOpts, sink chan<- *CentrallyIssuedTokenUpgrade, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _CentrallyIssuedToken.contract.WatchLogs(opts, "Upgrade", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CentrallyIssuedTokenUpgrade)
				if err := _CentrallyIssuedToken.contract.UnpackLog(event, "Upgrade", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgrade is a log parse operation binding the contract event 0x7e5c344a8141a805725cb476f76c6953b842222b967edd1f78ddb6e8b3f397ac.
//
// Solidity: event Upgrade(address indexed _from, address indexed _to, uint256 _value)
func (_CentrallyIssuedToken *CentrallyIssuedTokenFilterer) ParseUpgrade(log types.Log) (*CentrallyIssuedTokenUpgrade, error) {
	event := new(CentrallyIssuedTokenUpgrade)
	if err := _CentrallyIssuedToken.contract.UnpackLog(event, "Upgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CentrallyIssuedTokenUpgradeAgentSetIterator is returned from FilterUpgradeAgentSet and is used to iterate over the raw logs and unpacked data for UpgradeAgentSet events raised by the CentrallyIssuedToken contract.
type CentrallyIssuedTokenUpgradeAgentSetIterator struct {
	Event *CentrallyIssuedTokenUpgradeAgentSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CentrallyIssuedTokenUpgradeAgentSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CentrallyIssuedTokenUpgradeAgentSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CentrallyIssuedTokenUpgradeAgentSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CentrallyIssuedTokenUpgradeAgentSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CentrallyIssuedTokenUpgradeAgentSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CentrallyIssuedTokenUpgradeAgentSet represents a UpgradeAgentSet event raised by the CentrallyIssuedToken contract.
type CentrallyIssuedTokenUpgradeAgentSet struct {
	Agent common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpgradeAgentSet is a free log retrieval operation binding the contract event 0x7845d5aa74cc410e35571258d954f23b82276e160fe8c188fa80566580f279cc.
//
// Solidity: event UpgradeAgentSet(address agent)
func (_CentrallyIssuedToken *CentrallyIssuedTokenFilterer) FilterUpgradeAgentSet(opts *bind.FilterOpts) (*CentrallyIssuedTokenUpgradeAgentSetIterator, error) {

	logs, sub, err := _CentrallyIssuedToken.contract.FilterLogs(opts, "UpgradeAgentSet")
	if err != nil {
		return nil, err
	}
	return &CentrallyIssuedTokenUpgradeAgentSetIterator{contract: _CentrallyIssuedToken.contract, event: "UpgradeAgentSet", logs: logs, sub: sub}, nil
}

// WatchUpgradeAgentSet is a free log subscription operation binding the contract event 0x7845d5aa74cc410e35571258d954f23b82276e160fe8c188fa80566580f279cc.
//
// Solidity: event UpgradeAgentSet(address agent)
func (_CentrallyIssuedToken *CentrallyIssuedTokenFilterer) WatchUpgradeAgentSet(opts *bind.WatchOpts, sink chan<- *CentrallyIssuedTokenUpgradeAgentSet) (event.Subscription, error) {

	logs, sub, err := _CentrallyIssuedToken.contract.WatchLogs(opts, "UpgradeAgentSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CentrallyIssuedTokenUpgradeAgentSet)
				if err := _CentrallyIssuedToken.contract.UnpackLog(event, "UpgradeAgentSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgradeAgentSet is a log parse operation binding the contract event 0x7845d5aa74cc410e35571258d954f23b82276e160fe8c188fa80566580f279cc.
//
// Solidity: event UpgradeAgentSet(address agent)
func (_CentrallyIssuedToken *CentrallyIssuedTokenFilterer) ParseUpgradeAgentSet(log types.Log) (*CentrallyIssuedTokenUpgradeAgentSet, error) {
	event := new(CentrallyIssuedTokenUpgradeAgentSet)
	if err := _CentrallyIssuedToken.contract.UnpackLog(event, "UpgradeAgentSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20MetaData contains all meta data concerning the ERC20 contract.
var ERC20MetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"ok\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"ok\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"ok\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]",
}

// ERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20MetaData.ABI instead.
var ERC20ABI = ERC20MetaData.ABI

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "balanceOf", who)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool ok)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool ok)
func (_ERC20 *ERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool ok)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool ok)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool ok)
func (_ERC20 *ERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool ok)
func (_ERC20 *ERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool ok)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool ok)
func (_ERC20 *ERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool ok)
func (_ERC20 *ERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) {
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FactoryMetaData contains all meta data concerning the Factory contract.
var FactoryMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isInstantiation\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"instantiations\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"getInstantiationCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"instantiation\",\"type\":\"address\"}],\"name\":\"ContractInstantiation\",\"type\":\"event\"}]",
	Bin: "0x6060604052341561000f57600080fd5b6101de8061001e6000396000f3006060604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632f4f3316811461005b57806357183c821461009b5780638f838478146100f3575b600080fd5b341561006657600080fd5b61008773ffffffffffffffffffffffffffffffffffffffff60043516610131565b604051901515815260200160405180910390f35b34156100a657600080fd5b6100ca73ffffffffffffffffffffffffffffffffffffffff60043516602435610146565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156100fe57600080fd5b61011f73ffffffffffffffffffffffffffffffffffffffff6004351661018a565b60405190815260200160405180910390f35b60006020819052908152604090205460ff1681565b60016020528160005260406000208181548110151561016157fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff169150829050565b73ffffffffffffffffffffffffffffffffffffffff16600090815260016020526040902054905600a165627a7a72305820b133705a55d8fa2fef3244a7a922c24e12be88f7c6595ab0a3aede1b7d49d4e10029",
}

// FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use FactoryMetaData.ABI instead.
var FactoryABI = FactoryMetaData.ABI

// FactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FactoryMetaData.Bin instead.
var FactoryBin = FactoryMetaData.Bin

// DeployFactory deploys a new Ethereum contract, binding an instance of Factory to it.
func DeployFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Factory, error) {
	parsed, err := FactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

// Factory is an auto generated Go binding around an Ethereum contract.
type Factory struct {
	FactoryCaller     // Read-only binding to the contract
	FactoryTransactor // Write-only binding to the contract
	FactoryFilterer   // Log filterer for contract events
}

// FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FactorySession struct {
	Contract     *Factory          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FactoryCallerSession struct {
	Contract *FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FactoryTransactorSession struct {
	Contract     *FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FactoryRaw struct {
	Contract *Factory // Generic contract binding to access the raw methods on
}

// FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FactoryCallerRaw struct {
	Contract *FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FactoryTransactorRaw struct {
	Contract *FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFactory creates a new instance of Factory, bound to a specific deployed contract.
func NewFactory(address common.Address, backend bind.ContractBackend) (*Factory, error) {
	contract, err := bindFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

// NewFactoryCaller creates a new read-only instance of Factory, bound to a specific deployed contract.
func NewFactoryCaller(address common.Address, caller bind.ContractCaller) (*FactoryCaller, error) {
	contract, err := bindFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryCaller{contract: contract}, nil
}

// NewFactoryTransactor creates a new write-only instance of Factory, bound to a specific deployed contract.
func NewFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*FactoryTransactor, error) {
	contract, err := bindFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryTransactor{contract: contract}, nil
}

// NewFactoryFilterer creates a new log filterer instance of Factory, bound to a specific deployed contract.
func NewFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*FactoryFilterer, error) {
	contract, err := bindFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FactoryFilterer{contract: contract}, nil
}

// bindFactory binds a generic wrapper to an already deployed contract.
func bindFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transact(opts, method, params...)
}

// GetInstantiationCount is a free data retrieval call binding the contract method 0x8f838478.
//
// Solidity: function getInstantiationCount(address creator) view returns(uint256)
func (_Factory *FactoryCaller) GetInstantiationCount(opts *bind.CallOpts, creator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "getInstantiationCount", creator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInstantiationCount is a free data retrieval call binding the contract method 0x8f838478.
//
// Solidity: function getInstantiationCount(address creator) view returns(uint256)
func (_Factory *FactorySession) GetInstantiationCount(creator common.Address) (*big.Int, error) {
	return _Factory.Contract.GetInstantiationCount(&_Factory.CallOpts, creator)
}

// GetInstantiationCount is a free data retrieval call binding the contract method 0x8f838478.
//
// Solidity: function getInstantiationCount(address creator) view returns(uint256)
func (_Factory *FactoryCallerSession) GetInstantiationCount(creator common.Address) (*big.Int, error) {
	return _Factory.Contract.GetInstantiationCount(&_Factory.CallOpts, creator)
}

// Instantiations is a free data retrieval call binding the contract method 0x57183c82.
//
// Solidity: function instantiations(address , uint256 ) view returns(address)
func (_Factory *FactoryCaller) Instantiations(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "instantiations", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Instantiations is a free data retrieval call binding the contract method 0x57183c82.
//
// Solidity: function instantiations(address , uint256 ) view returns(address)
func (_Factory *FactorySession) Instantiations(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Factory.Contract.Instantiations(&_Factory.CallOpts, arg0, arg1)
}

// Instantiations is a free data retrieval call binding the contract method 0x57183c82.
//
// Solidity: function instantiations(address , uint256 ) view returns(address)
func (_Factory *FactoryCallerSession) Instantiations(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Factory.Contract.Instantiations(&_Factory.CallOpts, arg0, arg1)
}

// IsInstantiation is a free data retrieval call binding the contract method 0x2f4f3316.
//
// Solidity: function isInstantiation(address ) view returns(bool)
func (_Factory *FactoryCaller) IsInstantiation(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Factory.contract.Call(opts, &out, "isInstantiation", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInstantiation is a free data retrieval call binding the contract method 0x2f4f3316.
//
// Solidity: function isInstantiation(address ) view returns(bool)
func (_Factory *FactorySession) IsInstantiation(arg0 common.Address) (bool, error) {
	return _Factory.Contract.IsInstantiation(&_Factory.CallOpts, arg0)
}

// IsInstantiation is a free data retrieval call binding the contract method 0x2f4f3316.
//
// Solidity: function isInstantiation(address ) view returns(bool)
func (_Factory *FactoryCallerSession) IsInstantiation(arg0 common.Address) (bool, error) {
	return _Factory.Contract.IsInstantiation(&_Factory.CallOpts, arg0)
}

// FactoryContractInstantiationIterator is returned from FilterContractInstantiation and is used to iterate over the raw logs and unpacked data for ContractInstantiation events raised by the Factory contract.
type FactoryContractInstantiationIterator struct {
	Event *FactoryContractInstantiation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FactoryContractInstantiationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryContractInstantiation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FactoryContractInstantiation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FactoryContractInstantiationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryContractInstantiationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryContractInstantiation represents a ContractInstantiation event raised by the Factory contract.
type FactoryContractInstantiation struct {
	Sender        common.Address
	Instantiation common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterContractInstantiation is a free log retrieval operation binding the contract event 0x4fb057ad4a26ed17a57957fa69c306f11987596069b89521c511fc9a894e6161.
//
// Solidity: event ContractInstantiation(address sender, address instantiation)
func (_Factory *FactoryFilterer) FilterContractInstantiation(opts *bind.FilterOpts) (*FactoryContractInstantiationIterator, error) {

	logs, sub, err := _Factory.contract.FilterLogs(opts, "ContractInstantiation")
	if err != nil {
		return nil, err
	}
	return &FactoryContractInstantiationIterator{contract: _Factory.contract, event: "ContractInstantiation", logs: logs, sub: sub}, nil
}

// WatchContractInstantiation is a free log subscription operation binding the contract event 0x4fb057ad4a26ed17a57957fa69c306f11987596069b89521c511fc9a894e6161.
//
// Solidity: event ContractInstantiation(address sender, address instantiation)
func (_Factory *FactoryFilterer) WatchContractInstantiation(opts *bind.WatchOpts, sink chan<- *FactoryContractInstantiation) (event.Subscription, error) {

	logs, sub, err := _Factory.contract.WatchLogs(opts, "ContractInstantiation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryContractInstantiation)
				if err := _Factory.contract.UnpackLog(event, "ContractInstantiation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContractInstantiation is a log parse operation binding the contract event 0x4fb057ad4a26ed17a57957fa69c306f11987596069b89521c511fc9a894e6161.
//
// Solidity: event ContractInstantiation(address sender, address instantiation)
func (_Factory *FactoryFilterer) ParseContractInstantiation(log types.Log) (*FactoryContractInstantiation, error) {
	event := new(FactoryContractInstantiation)
	if err := _Factory.contract.UnpackLog(event, "ContractInstantiation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigWalletMetaData contains all meta data concerning the MultiSigWallet contract.
var MultiSigWalletMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"removeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"revokeConfirmation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"confirmations\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"pending\",\"type\":\"bool\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"addOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"isConfirmed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmationCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"from\",\"type\":\"uint256\"},{\"name\":\"to\",\"type\":\"uint256\"},{\"name\":\"pending\",\"type\":\"bool\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionIds\",\"outputs\":[{\"name\":\"_transactionIds\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmations\",\"outputs\":[{\"name\":\"_confirmations\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transactionCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_required\",\"type\":\"uint256\"}],\"name\":\"changeRequirement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"confirmTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"submitTransaction\",\"outputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OWNER_COUNT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"required\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"replaceOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"},{\"name\":\"_required\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Confirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Revocation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Submission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Execution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"ExecutionFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"RequirementChange\",\"type\":\"event\"}]",
	Bin: "0x606060405234156200001057600080fd5b60405162001b3338038062001b3383398101604052808051820191906020018051915060009050825182603282111580156200004c5750818111155b80156200005857508015155b80156200006457508115155b15156200007057600080fd5b600092505b84518310156200014157600260008685815181106200009057fe5b90602001906020020151600160a060020a0316815260208101919091526040016000205460ff16158015620000e35750848381518110620000cd57fe5b90602001906020020151600160a060020a031615155b1515620000ef57600080fd5b6001600260008786815181106200010257fe5b90602001906020020151600160a060020a031681526020810191909152604001600020805460ff19169115159190911790556001929092019162000075565b60038580516200015692916020019062000168565b50505060049190915550620001fe9050565b828054828255906000526020600020908101928215620001c2579160200282015b82811115620001c25782518254600160a060020a031916600160a060020a03919091161782556020929092019160019091019062000189565b50620001d0929150620001d4565b5090565b620001fb91905b80821115620001d0578054600160a060020a0319168155600101620001db565b90565b611925806200020e6000396000f30060606040526004361061011c5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663025e7c278114610172578063173825d9146101b157806320ea8d86146101dd5780632f54bf6e146101f35780633411c81c1461023357806354741525146102625780637065cb4814610291578063784547a7146102bd5780638b51d13f146102d35780639ace38c2146102e9578063a0e67e2b146103d3578063a8abe69a14610439578063b5dc40c31461045c578063b77bf60014610472578063ba51a6df14610485578063c01a8c841461049b578063c6427474146104b1578063d74f8edd14610523578063dc8452cd14610536578063e20056e614610549578063ee22610b1461057b575b6000341115610170573373ffffffffffffffffffffffffffffffffffffffff167fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c3460405190815260200160405180910390a25b005b341561017d57600080fd5b610188600435610591565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156101bc57600080fd5b61017073ffffffffffffffffffffffffffffffffffffffff600435166105c6565b34156101e857600080fd5b610170600435610853565b34156101fe57600080fd5b61021f73ffffffffffffffffffffffffffffffffffffffff60043516610976565b604051901515815260200160405180910390f35b341561023e57600080fd5b61021f60043573ffffffffffffffffffffffffffffffffffffffff6024351661098b565b341561026d57600080fd5b61027f600435151560243515156109ab565b60405190815260200160405180910390f35b341561029c57600080fd5b61017073ffffffffffffffffffffffffffffffffffffffff60043516610a17565b34156102c857600080fd5b61021f600435610bca565b34156102de57600080fd5b61027f600435610c5b565b34156102f457600080fd5b6102ff600435610cd7565b60405173ffffffffffffffffffffffffffffffffffffffff85168152602081018490528115156060820152608060408201818152845460027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61010060018416150201909116049183018290529060a0830190859080156103c15780601f10610396576101008083540402835291602001916103c1565b820191906000526020600020905b8154815290600101906020018083116103a457829003601f168201915b50509550505050505060405180910390f35b34156103de57600080fd5b6103e6610d18565b60405160208082528190810183818151815260200191508051906020019060200280838360005b8381101561042557808201518382015260200161040d565b505050509050019250505060405180910390f35b341561044457600080fd5b6103e660043560243560443515156064351515610d8e565b341561046757600080fd5b6103e6600435610eb6565b341561047d57600080fd5b61027f61104e565b341561049057600080fd5b610170600435611054565b34156104a657600080fd5b610170600435611101565b34156104bc57600080fd5b61027f6004803573ffffffffffffffffffffffffffffffffffffffff169060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061124195505050505050565b341561052e57600080fd5b61027f611260565b341561054157600080fd5b61027f611265565b341561055457600080fd5b61017073ffffffffffffffffffffffffffffffffffffffff6004358116906024351661126b565b341561058657600080fd5b6101706004356114b7565b600380548290811061059f57fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b60003073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561060257600080fd5b73ffffffffffffffffffffffffffffffffffffffff8216600090815260026020526040902054829060ff16151561063857600080fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260026020526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905591505b6003547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018210156107c1578273ffffffffffffffffffffffffffffffffffffffff166003838154811015156106d857fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1614156107b657600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190811061073057fe5b6000918252602090912001546003805473ffffffffffffffffffffffffffffffffffffffff909216918490811061076357fe5b600091825260209091200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff929092169190911790556107c1565b600190910190610686565b600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906107f29082611826565b50600354600454111561080b5760035461080b90611054565b8273ffffffffffffffffffffffffffffffffffffffff167f8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b9060405160405180910390a2505050565b3373ffffffffffffffffffffffffffffffffffffffff811660009081526002602052604090205460ff16151561088857600080fd5b60008281526001602090815260408083203373ffffffffffffffffffffffffffffffffffffffff8116855292529091205483919060ff1615156108ca57600080fd5b600084815260208190526040902060030154849060ff16156108eb57600080fd5b600085815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff331680855292529182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905586917ff6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9905160405180910390a35050505050565b60026020526000908152604090205460ff1681565b600160209081526000928352604080842090915290825290205460ff1681565b6000805b600554811015610a10578380156109d8575060008181526020819052604090206003015460ff16155b806109fc57508280156109fc575060008181526020819052604090206003015460ff165b15610a08576001820191505b6001016109af565b5092915050565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610a5157600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116600090815260026020526040902054819060ff1615610a8657600080fd5b8173ffffffffffffffffffffffffffffffffffffffff81161515610aa957600080fd5b60038054905060010160045460328211158015610ac65750818111155b8015610ad157508015155b8015610adc57508115155b1515610ae757600080fd5b73ffffffffffffffffffffffffffffffffffffffff8516600090815260026020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660019081179091556003805490918101610b4a8382611826565b50600091825260209091200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff87169081179091557ff39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d60405160405180910390a25050505050565b600080805b600354811015610c545760008481526001602052604081206003805491929184908110610bf857fe5b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff16835282019290925260400190205460ff1615610c39576001820191505b600454821415610c4c5760019250610c54565b600101610bcf565b5050919050565b6000805b600354811015610cd15760008381526001602052604081206003805491929184908110610c8857fe5b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff16835282019290925260400190205460ff1615610cc9576001820191505b600101610c5f565b50919050565b600060208190529081526040902080546001820154600383015473ffffffffffffffffffffffffffffffffffffffff90921692909160029091019060ff1684565b610d2061184f565b6003805480602002602001604051908101604052809291908181526020018280548015610d8357602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610d58575b505050505090505b90565b610d9661184f565b610d9e61184f565b600080600554604051805910610db15750595b9080825280602002602001820160405250925060009150600090505b600554811015610e4657858015610df6575060008181526020819052604090206003015460ff16155b80610e1a5750848015610e1a575060008181526020819052604090206003015460ff165b15610e3e5780838381518110610e2c57fe5b60209081029091010152600191909101905b600101610dcd565b878703604051805910610e565750595b908082528060200260200182016040525093508790505b86811015610eab57828181518110610e8157fe5b906020019060200201518489830381518110610e9957fe5b60209081029091010152600101610e6d565b505050949350505050565b610ebe61184f565b610ec661184f565b6003546000908190604051805910610edb5750595b9080825280602002602001820160405250925060009150600090505b600354811015610fca5760008581526001602052604081206003805491929184908110610f2057fe5b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff16835282019290925260400190205460ff1615610fc2576003805482908110610f6857fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16838381518110610f9657fe5b73ffffffffffffffffffffffffffffffffffffffff909216602092830290910190910152600191909101905b600101610ef7565b81604051805910610fd85750595b90808252806020026020018201604052509350600090505b818110156110465782818151811061100457fe5b9060200190602002015184828151811061101a57fe5b73ffffffffffffffffffffffffffffffffffffffff909216602092830290910190910152600101610ff0565b505050919050565b60055481565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561108e57600080fd5b60035481603282118015906110a35750818111155b80156110ae57508015155b80156110b957508115155b15156110c457600080fd5b60048390557fa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a8360405190815260200160405180910390a1505050565b3373ffffffffffffffffffffffffffffffffffffffff811660009081526002602052604090205460ff16151561113657600080fd5b600082815260208190526040902054829073ffffffffffffffffffffffffffffffffffffffff16151561116857600080fd5b60008381526001602090815260408083203373ffffffffffffffffffffffffffffffffffffffff8116855292529091205484919060ff16156111a957600080fd5b600085815260016020818152604080842073ffffffffffffffffffffffffffffffffffffffff331680865292529283902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690921790915586917f4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef905160405180910390a361123a856114b7565b5050505050565b600061124e8484846116d9565b905061125981611101565b9392505050565b603281565b60045481565b60003073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156112a757600080fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260026020526040902054839060ff1615156112dd57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260026020526040902054839060ff161561131257600080fd5b600092505b6003548310156113dd578473ffffffffffffffffffffffffffffffffffffffff1660038481548110151561134757fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1614156113d2578360038481548110151561137f57fe5b600091825260209091200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff929092169190911790556113dd565b600190920191611317565b73ffffffffffffffffffffffffffffffffffffffff80861660008181526002602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff009081169091559388168252908190208054909316600117909255907f8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90905160405180910390a28373ffffffffffffffffffffffffffffffffffffffff167ff39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d60405160405180910390a25050505050565b3373ffffffffffffffffffffffffffffffffffffffff811660009081526002602052604081205490919060ff1615156114ef57600080fd5b60008381526001602090815260408083203373ffffffffffffffffffffffffffffffffffffffff8116855292529091205484919060ff16151561153157600080fd5b600085815260208190526040902060030154859060ff161561155257600080fd5b61155b86610bca565b156116d157600086815260208190526040908190206003810180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915581549082015491975073ffffffffffffffffffffffffffffffffffffffff169160028801905180828054600181600116156101000203166002900480156116285780601f106115fd57610100808354040283529160200191611628565b820191906000526020600020905b81548152906001019060200180831161160b57829003601f168201915b505091505060006040518083038185876187965a03f1925050501561167957857f33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed7560405160405180910390a26116d1565b857f526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b7923660405160405180910390a26003850180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690555b505050505050565b60008373ffffffffffffffffffffffffffffffffffffffff811615156116fe57600080fd5b60055491506080604051908101604090815273ffffffffffffffffffffffffffffffffffffffff8716825260208083018790528183018690526000606084018190528581529081905220815181547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff91909116178155602082015181600101556040820151816002019080516117ae929160200190611861565b50606082015160039190910180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001691151591909117905550600580546001019055817fc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e5160405160405180910390a2509392505050565b81548183558181151161184a5760008381526020902061184a9181019083016118df565b505050565b60206040519081016040526000815290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106118a257805160ff19168380011785556118cf565b828001600101855582156118cf579182015b828111156118cf5782518255916020019190600101906118b4565b506118db9291506118df565b5090565b610d8b91905b808211156118db57600081556001016118e55600a165627a7a723058208a4165c7738d17a26a421505b5f200adc00843577c1253abbca6dcd244b26c090029",
}

// MultiSigWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use MultiSigWalletMetaData.ABI instead.
var MultiSigWalletABI = MultiSigWalletMetaData.ABI

// MultiSigWalletBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MultiSigWalletMetaData.Bin instead.
var MultiSigWalletBin = MultiSigWalletMetaData.Bin

// DeployMultiSigWallet deploys a new Ethereum contract, binding an instance of MultiSigWallet to it.
func DeployMultiSigWallet(auth *bind.TransactOpts, backend bind.ContractBackend, _owners []common.Address, _required *big.Int) (common.Address, *types.Transaction, *MultiSigWallet, error) {
	parsed, err := MultiSigWalletMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MultiSigWalletBin), backend, _owners, _required)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MultiSigWallet{MultiSigWalletCaller: MultiSigWalletCaller{contract: contract}, MultiSigWalletTransactor: MultiSigWalletTransactor{contract: contract}, MultiSigWalletFilterer: MultiSigWalletFilterer{contract: contract}}, nil
}

// MultiSigWallet is an auto generated Go binding around an Ethereum contract.
type MultiSigWallet struct {
	MultiSigWalletCaller     // Read-only binding to the contract
	MultiSigWalletTransactor // Write-only binding to the contract
	MultiSigWalletFilterer   // Log filterer for contract events
}

// MultiSigWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiSigWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiSigWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiSigWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiSigWalletSession struct {
	Contract     *MultiSigWallet   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultiSigWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiSigWalletCallerSession struct {
	Contract *MultiSigWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MultiSigWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiSigWalletTransactorSession struct {
	Contract     *MultiSigWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MultiSigWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiSigWalletRaw struct {
	Contract *MultiSigWallet // Generic contract binding to access the raw methods on
}

// MultiSigWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiSigWalletCallerRaw struct {
	Contract *MultiSigWalletCaller // Generic read-only contract binding to access the raw methods on
}

// MultiSigWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiSigWalletTransactorRaw struct {
	Contract *MultiSigWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiSigWallet creates a new instance of MultiSigWallet, bound to a specific deployed contract.
func NewMultiSigWallet(address common.Address, backend bind.ContractBackend) (*MultiSigWallet, error) {
	contract, err := bindMultiSigWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiSigWallet{MultiSigWalletCaller: MultiSigWalletCaller{contract: contract}, MultiSigWalletTransactor: MultiSigWalletTransactor{contract: contract}, MultiSigWalletFilterer: MultiSigWalletFilterer{contract: contract}}, nil
}

// NewMultiSigWalletCaller creates a new read-only instance of MultiSigWallet, bound to a specific deployed contract.
func NewMultiSigWalletCaller(address common.Address, caller bind.ContractCaller) (*MultiSigWalletCaller, error) {
	contract, err := bindMultiSigWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletCaller{contract: contract}, nil
}

// NewMultiSigWalletTransactor creates a new write-only instance of MultiSigWallet, bound to a specific deployed contract.
func NewMultiSigWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiSigWalletTransactor, error) {
	contract, err := bindMultiSigWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletTransactor{contract: contract}, nil
}

// NewMultiSigWalletFilterer creates a new log filterer instance of MultiSigWallet, bound to a specific deployed contract.
func NewMultiSigWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiSigWalletFilterer, error) {
	contract, err := bindMultiSigWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletFilterer{contract: contract}, nil
}

// bindMultiSigWallet binds a generic wrapper to an already deployed contract.
func bindMultiSigWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MultiSigWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSigWallet *MultiSigWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiSigWallet.Contract.MultiSigWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSigWallet *MultiSigWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.MultiSigWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSigWallet *MultiSigWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.MultiSigWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSigWallet *MultiSigWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiSigWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSigWallet *MultiSigWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSigWallet *MultiSigWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.contract.Transact(opts, method, params...)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() view returns(uint256)
func (_MultiSigWallet *MultiSigWalletCaller) MAXOWNERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiSigWallet.contract.Call(opts, &out, "MAX_OWNER_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() view returns(uint256)
func (_MultiSigWallet *MultiSigWalletSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _MultiSigWallet.Contract.MAXOWNERCOUNT(&_MultiSigWallet.CallOpts)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() view returns(uint256)
func (_MultiSigWallet *MultiSigWalletCallerSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _MultiSigWallet.Contract.MAXOWNERCOUNT(&_MultiSigWallet.CallOpts)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) view returns(bool)
func (_MultiSigWallet *MultiSigWalletCaller) Confirmations(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _MultiSigWallet.contract.Call(opts, &out, "confirmations", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) view returns(bool)
func (_MultiSigWallet *MultiSigWalletSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _MultiSigWallet.Contract.Confirmations(&_MultiSigWallet.CallOpts, arg0, arg1)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) view returns(bool)
func (_MultiSigWallet *MultiSigWalletCallerSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _MultiSigWallet.Contract.Confirmations(&_MultiSigWallet.CallOpts, arg0, arg1)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) view returns(uint256 count)
func (_MultiSigWallet *MultiSigWalletCaller) GetConfirmationCount(opts *bind.CallOpts, transactionId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MultiSigWallet.contract.Call(opts, &out, "getConfirmationCount", transactionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) view returns(uint256 count)
func (_MultiSigWallet *MultiSigWalletSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _MultiSigWallet.Contract.GetConfirmationCount(&_MultiSigWallet.CallOpts, transactionId)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) view returns(uint256 count)
func (_MultiSigWallet *MultiSigWalletCallerSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _MultiSigWallet.Contract.GetConfirmationCount(&_MultiSigWallet.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) view returns(address[] _confirmations)
func (_MultiSigWallet *MultiSigWalletCaller) GetConfirmations(opts *bind.CallOpts, transactionId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _MultiSigWallet.contract.Call(opts, &out, "getConfirmations", transactionId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) view returns(address[] _confirmations)
func (_MultiSigWallet *MultiSigWalletSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _MultiSigWallet.Contract.GetConfirmations(&_MultiSigWallet.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) view returns(address[] _confirmations)
func (_MultiSigWallet *MultiSigWalletCallerSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _MultiSigWallet.Contract.GetConfirmations(&_MultiSigWallet.CallOpts, transactionId)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_MultiSigWallet *MultiSigWalletCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MultiSigWallet.contract.Call(opts, &out, "getOwners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_MultiSigWallet *MultiSigWalletSession) GetOwners() ([]common.Address, error) {
	return _MultiSigWallet.Contract.GetOwners(&_MultiSigWallet.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_MultiSigWallet *MultiSigWalletCallerSession) GetOwners() ([]common.Address, error) {
	return _MultiSigWallet.Contract.GetOwners(&_MultiSigWallet.CallOpts)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) view returns(uint256 count)
func (_MultiSigWallet *MultiSigWalletCaller) GetTransactionCount(opts *bind.CallOpts, pending bool, executed bool) (*big.Int, error) {
	var out []interface{}
	err := _MultiSigWallet.contract.Call(opts, &out, "getTransactionCount", pending, executed)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) view returns(uint256 count)
func (_MultiSigWallet *MultiSigWalletSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _MultiSigWallet.Contract.GetTransactionCount(&_MultiSigWallet.CallOpts, pending, executed)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) view returns(uint256 count)
func (_MultiSigWallet *MultiSigWalletCallerSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _MultiSigWallet.Contract.GetTransactionCount(&_MultiSigWallet.CallOpts, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) view returns(uint256[] _transactionIds)
func (_MultiSigWallet *MultiSigWalletCaller) GetTransactionIds(opts *bind.CallOpts, from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	var out []interface{}
	err := _MultiSigWallet.contract.Call(opts, &out, "getTransactionIds", from, to, pending, executed)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) view returns(uint256[] _transactionIds)
func (_MultiSigWallet *MultiSigWalletSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _MultiSigWallet.Contract.GetTransactionIds(&_MultiSigWallet.CallOpts, from, to, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) view returns(uint256[] _transactionIds)
func (_MultiSigWallet *MultiSigWalletCallerSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _MultiSigWallet.Contract.GetTransactionIds(&_MultiSigWallet.CallOpts, from, to, pending, executed)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) view returns(bool)
func (_MultiSigWallet *MultiSigWalletCaller) IsConfirmed(opts *bind.CallOpts, transactionId *big.Int) (bool, error) {
	var out []interface{}
	err := _MultiSigWallet.contract.Call(opts, &out, "isConfirmed", transactionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) view returns(bool)
func (_MultiSigWallet *MultiSigWalletSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _MultiSigWallet.Contract.IsConfirmed(&_MultiSigWallet.CallOpts, transactionId)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) view returns(bool)
func (_MultiSigWallet *MultiSigWalletCallerSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _MultiSigWallet.Contract.IsConfirmed(&_MultiSigWallet.CallOpts, transactionId)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_MultiSigWallet *MultiSigWalletCaller) IsOwner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MultiSigWallet.contract.Call(opts, &out, "isOwner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_MultiSigWallet *MultiSigWalletSession) IsOwner(arg0 common.Address) (bool, error) {
	return _MultiSigWallet.Contract.IsOwner(&_MultiSigWallet.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_MultiSigWallet *MultiSigWalletCallerSession) IsOwner(arg0 common.Address) (bool, error) {
	return _MultiSigWallet.Contract.IsOwner(&_MultiSigWallet.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_MultiSigWallet *MultiSigWalletCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MultiSigWallet.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_MultiSigWallet *MultiSigWalletSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _MultiSigWallet.Contract.Owners(&_MultiSigWallet.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_MultiSigWallet *MultiSigWalletCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _MultiSigWallet.Contract.Owners(&_MultiSigWallet.CallOpts, arg0)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() view returns(uint256)
func (_MultiSigWallet *MultiSigWalletCaller) Required(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiSigWallet.contract.Call(opts, &out, "required")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() view returns(uint256)
func (_MultiSigWallet *MultiSigWalletSession) Required() (*big.Int, error) {
	return _MultiSigWallet.Contract.Required(&_MultiSigWallet.CallOpts)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() view returns(uint256)
func (_MultiSigWallet *MultiSigWalletCallerSession) Required() (*big.Int, error) {
	return _MultiSigWallet.Contract.Required(&_MultiSigWallet.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_MultiSigWallet *MultiSigWalletCaller) TransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiSigWallet.contract.Call(opts, &out, "transactionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_MultiSigWallet *MultiSigWalletSession) TransactionCount() (*big.Int, error) {
	return _MultiSigWallet.Contract.TransactionCount(&_MultiSigWallet.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_MultiSigWallet *MultiSigWalletCallerSession) TransactionCount() (*big.Int, error) {
	return _MultiSigWallet.Contract.TransactionCount(&_MultiSigWallet.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address destination, uint256 value, bytes data, bool executed)
func (_MultiSigWallet *MultiSigWalletCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	var out []interface{}
	err := _MultiSigWallet.contract.Call(opts, &out, "transactions", arg0)

	outstruct := new(struct {
		Destination common.Address
		Value       *big.Int
		Data        []byte
		Executed    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Destination = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Data = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Executed = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address destination, uint256 value, bytes data, bool executed)
func (_MultiSigWallet *MultiSigWalletSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _MultiSigWallet.Contract.Transactions(&_MultiSigWallet.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address destination, uint256 value, bytes data, bool executed)
func (_MultiSigWallet *MultiSigWalletCallerSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _MultiSigWallet.Contract.Transactions(&_MultiSigWallet.CallOpts, arg0)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_MultiSigWallet *MultiSigWalletTransactor) AddOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _MultiSigWallet.contract.Transact(opts, "addOwner", owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_MultiSigWallet *MultiSigWalletSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.AddOwner(&_MultiSigWallet.TransactOpts, owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_MultiSigWallet *MultiSigWalletTransactorSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.AddOwner(&_MultiSigWallet.TransactOpts, owner)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_MultiSigWallet *MultiSigWalletTransactor) ChangeRequirement(opts *bind.TransactOpts, _required *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.contract.Transact(opts, "changeRequirement", _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_MultiSigWallet *MultiSigWalletSession) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.ChangeRequirement(&_MultiSigWallet.TransactOpts, _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_MultiSigWallet *MultiSigWalletTransactorSession) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.ChangeRequirement(&_MultiSigWallet.TransactOpts, _required)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_MultiSigWallet *MultiSigWalletTransactor) ConfirmTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.contract.Transact(opts, "confirmTransaction", transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_MultiSigWallet *MultiSigWalletSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.ConfirmTransaction(&_MultiSigWallet.TransactOpts, transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_MultiSigWallet *MultiSigWalletTransactorSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.ConfirmTransaction(&_MultiSigWallet.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_MultiSigWallet *MultiSigWalletTransactor) ExecuteTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.contract.Transact(opts, "executeTransaction", transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_MultiSigWallet *MultiSigWalletSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.ExecuteTransaction(&_MultiSigWallet.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_MultiSigWallet *MultiSigWalletTransactorSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.ExecuteTransaction(&_MultiSigWallet.TransactOpts, transactionId)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_MultiSigWallet *MultiSigWalletTransactor) RemoveOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _MultiSigWallet.contract.Transact(opts, "removeOwner", owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_MultiSigWallet *MultiSigWalletSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.RemoveOwner(&_MultiSigWallet.TransactOpts, owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_MultiSigWallet *MultiSigWalletTransactorSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.RemoveOwner(&_MultiSigWallet.TransactOpts, owner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_MultiSigWallet *MultiSigWalletTransactor) ReplaceOwner(opts *bind.TransactOpts, owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MultiSigWallet.contract.Transact(opts, "replaceOwner", owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_MultiSigWallet *MultiSigWalletSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.ReplaceOwner(&_MultiSigWallet.TransactOpts, owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_MultiSigWallet *MultiSigWalletTransactorSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.ReplaceOwner(&_MultiSigWallet.TransactOpts, owner, newOwner)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_MultiSigWallet *MultiSigWalletTransactor) RevokeConfirmation(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.contract.Transact(opts, "revokeConfirmation", transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_MultiSigWallet *MultiSigWalletSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.RevokeConfirmation(&_MultiSigWallet.TransactOpts, transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_MultiSigWallet *MultiSigWalletTransactorSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.RevokeConfirmation(&_MultiSigWallet.TransactOpts, transactionId)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_MultiSigWallet *MultiSigWalletTransactor) SubmitTransaction(opts *bind.TransactOpts, destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultiSigWallet.contract.Transact(opts, "submitTransaction", destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_MultiSigWallet *MultiSigWalletSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.SubmitTransaction(&_MultiSigWallet.TransactOpts, destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_MultiSigWallet *MultiSigWalletTransactorSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.SubmitTransaction(&_MultiSigWallet.TransactOpts, destination, value, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MultiSigWallet *MultiSigWalletTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _MultiSigWallet.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MultiSigWallet *MultiSigWalletSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.Fallback(&_MultiSigWallet.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MultiSigWallet *MultiSigWalletTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MultiSigWallet.Contract.Fallback(&_MultiSigWallet.TransactOpts, calldata)
}

// MultiSigWalletConfirmationIterator is returned from FilterConfirmation and is used to iterate over the raw logs and unpacked data for Confirmation events raised by the MultiSigWallet contract.
type MultiSigWalletConfirmationIterator struct {
	Event *MultiSigWalletConfirmation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletConfirmation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletConfirmation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletConfirmation represents a Confirmation event raised by the MultiSigWallet contract.
type MultiSigWalletConfirmation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterConfirmation is a free log retrieval operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWallet *MultiSigWalletFilterer) FilterConfirmation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*MultiSigWalletConfirmationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.FilterLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletConfirmationIterator{contract: _MultiSigWallet.contract, event: "Confirmation", logs: logs, sub: sub}, nil
}

// WatchConfirmation is a free log subscription operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWallet *MultiSigWalletFilterer) WatchConfirmation(opts *bind.WatchOpts, sink chan<- *MultiSigWalletConfirmation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.WatchLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletConfirmation)
				if err := _MultiSigWallet.contract.UnpackLog(event, "Confirmation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConfirmation is a log parse operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWallet *MultiSigWalletFilterer) ParseConfirmation(log types.Log) (*MultiSigWalletConfirmation, error) {
	event := new(MultiSigWalletConfirmation)
	if err := _MultiSigWallet.contract.UnpackLog(event, "Confirmation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigWalletDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the MultiSigWallet contract.
type MultiSigWalletDepositIterator struct {
	Event *MultiSigWalletDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletDeposit represents a Deposit event raised by the MultiSigWallet contract.
type MultiSigWalletDeposit struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_MultiSigWallet *MultiSigWalletFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*MultiSigWalletDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MultiSigWallet.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletDepositIterator{contract: _MultiSigWallet.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_MultiSigWallet *MultiSigWalletFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MultiSigWalletDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MultiSigWallet.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletDeposit)
				if err := _MultiSigWallet.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_MultiSigWallet *MultiSigWalletFilterer) ParseDeposit(log types.Log) (*MultiSigWalletDeposit, error) {
	event := new(MultiSigWalletDeposit)
	if err := _MultiSigWallet.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigWalletExecutionIterator is returned from FilterExecution and is used to iterate over the raw logs and unpacked data for Execution events raised by the MultiSigWallet contract.
type MultiSigWalletExecutionIterator struct {
	Event *MultiSigWalletExecution // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletExecution)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletExecution)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletExecution represents a Execution event raised by the MultiSigWallet contract.
type MultiSigWalletExecution struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecution is a free log retrieval operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(uint256 indexed transactionId)
func (_MultiSigWallet *MultiSigWalletFilterer) FilterExecution(opts *bind.FilterOpts, transactionId []*big.Int) (*MultiSigWalletExecutionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.FilterLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletExecutionIterator{contract: _MultiSigWallet.contract, event: "Execution", logs: logs, sub: sub}, nil
}

// WatchExecution is a free log subscription operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(uint256 indexed transactionId)
func (_MultiSigWallet *MultiSigWalletFilterer) WatchExecution(opts *bind.WatchOpts, sink chan<- *MultiSigWalletExecution, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.WatchLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletExecution)
				if err := _MultiSigWallet.contract.UnpackLog(event, "Execution", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecution is a log parse operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(uint256 indexed transactionId)
func (_MultiSigWallet *MultiSigWalletFilterer) ParseExecution(log types.Log) (*MultiSigWalletExecution, error) {
	event := new(MultiSigWalletExecution)
	if err := _MultiSigWallet.contract.UnpackLog(event, "Execution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigWalletExecutionFailureIterator is returned from FilterExecutionFailure and is used to iterate over the raw logs and unpacked data for ExecutionFailure events raised by the MultiSigWallet contract.
type MultiSigWalletExecutionFailureIterator struct {
	Event *MultiSigWalletExecutionFailure // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletExecutionFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletExecutionFailure)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletExecutionFailure)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletExecutionFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletExecutionFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletExecutionFailure represents a ExecutionFailure event raised by the MultiSigWallet contract.
type MultiSigWalletExecutionFailure struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionFailure is a free log retrieval operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 indexed transactionId)
func (_MultiSigWallet *MultiSigWalletFilterer) FilterExecutionFailure(opts *bind.FilterOpts, transactionId []*big.Int) (*MultiSigWalletExecutionFailureIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.FilterLogs(opts, "ExecutionFailure", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletExecutionFailureIterator{contract: _MultiSigWallet.contract, event: "ExecutionFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFailure is a free log subscription operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 indexed transactionId)
func (_MultiSigWallet *MultiSigWalletFilterer) WatchExecutionFailure(opts *bind.WatchOpts, sink chan<- *MultiSigWalletExecutionFailure, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.WatchLogs(opts, "ExecutionFailure", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletExecutionFailure)
				if err := _MultiSigWallet.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutionFailure is a log parse operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 indexed transactionId)
func (_MultiSigWallet *MultiSigWalletFilterer) ParseExecutionFailure(log types.Log) (*MultiSigWalletExecutionFailure, error) {
	event := new(MultiSigWalletExecutionFailure)
	if err := _MultiSigWallet.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigWalletOwnerAdditionIterator is returned from FilterOwnerAddition and is used to iterate over the raw logs and unpacked data for OwnerAddition events raised by the MultiSigWallet contract.
type MultiSigWalletOwnerAdditionIterator struct {
	Event *MultiSigWalletOwnerAddition // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletOwnerAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletOwnerAddition)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletOwnerAddition)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletOwnerAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletOwnerAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletOwnerAddition represents a OwnerAddition event raised by the MultiSigWallet contract.
type MultiSigWalletOwnerAddition struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerAddition is a free log retrieval operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_MultiSigWallet *MultiSigWalletFilterer) FilterOwnerAddition(opts *bind.FilterOpts, owner []common.Address) (*MultiSigWalletOwnerAdditionIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSigWallet.contract.FilterLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletOwnerAdditionIterator{contract: _MultiSigWallet.contract, event: "OwnerAddition", logs: logs, sub: sub}, nil
}

// WatchOwnerAddition is a free log subscription operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_MultiSigWallet *MultiSigWalletFilterer) WatchOwnerAddition(opts *bind.WatchOpts, sink chan<- *MultiSigWalletOwnerAddition, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSigWallet.contract.WatchLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletOwnerAddition)
				if err := _MultiSigWallet.contract.UnpackLog(event, "OwnerAddition", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnerAddition is a log parse operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_MultiSigWallet *MultiSigWalletFilterer) ParseOwnerAddition(log types.Log) (*MultiSigWalletOwnerAddition, error) {
	event := new(MultiSigWalletOwnerAddition)
	if err := _MultiSigWallet.contract.UnpackLog(event, "OwnerAddition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigWalletOwnerRemovalIterator is returned from FilterOwnerRemoval and is used to iterate over the raw logs and unpacked data for OwnerRemoval events raised by the MultiSigWallet contract.
type MultiSigWalletOwnerRemovalIterator struct {
	Event *MultiSigWalletOwnerRemoval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletOwnerRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletOwnerRemoval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletOwnerRemoval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletOwnerRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletOwnerRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletOwnerRemoval represents a OwnerRemoval event raised by the MultiSigWallet contract.
type MultiSigWalletOwnerRemoval struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerRemoval is a free log retrieval operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_MultiSigWallet *MultiSigWalletFilterer) FilterOwnerRemoval(opts *bind.FilterOpts, owner []common.Address) (*MultiSigWalletOwnerRemovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSigWallet.contract.FilterLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletOwnerRemovalIterator{contract: _MultiSigWallet.contract, event: "OwnerRemoval", logs: logs, sub: sub}, nil
}

// WatchOwnerRemoval is a free log subscription operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_MultiSigWallet *MultiSigWalletFilterer) WatchOwnerRemoval(opts *bind.WatchOpts, sink chan<- *MultiSigWalletOwnerRemoval, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSigWallet.contract.WatchLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletOwnerRemoval)
				if err := _MultiSigWallet.contract.UnpackLog(event, "OwnerRemoval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnerRemoval is a log parse operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_MultiSigWallet *MultiSigWalletFilterer) ParseOwnerRemoval(log types.Log) (*MultiSigWalletOwnerRemoval, error) {
	event := new(MultiSigWalletOwnerRemoval)
	if err := _MultiSigWallet.contract.UnpackLog(event, "OwnerRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigWalletRequirementChangeIterator is returned from FilterRequirementChange and is used to iterate over the raw logs and unpacked data for RequirementChange events raised by the MultiSigWallet contract.
type MultiSigWalletRequirementChangeIterator struct {
	Event *MultiSigWalletRequirementChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletRequirementChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletRequirementChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletRequirementChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletRequirementChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletRequirementChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletRequirementChange represents a RequirementChange event raised by the MultiSigWallet contract.
type MultiSigWalletRequirementChange struct {
	Required *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequirementChange is a free log retrieval operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_MultiSigWallet *MultiSigWalletFilterer) FilterRequirementChange(opts *bind.FilterOpts) (*MultiSigWalletRequirementChangeIterator, error) {

	logs, sub, err := _MultiSigWallet.contract.FilterLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletRequirementChangeIterator{contract: _MultiSigWallet.contract, event: "RequirementChange", logs: logs, sub: sub}, nil
}

// WatchRequirementChange is a free log subscription operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_MultiSigWallet *MultiSigWalletFilterer) WatchRequirementChange(opts *bind.WatchOpts, sink chan<- *MultiSigWalletRequirementChange) (event.Subscription, error) {

	logs, sub, err := _MultiSigWallet.contract.WatchLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletRequirementChange)
				if err := _MultiSigWallet.contract.UnpackLog(event, "RequirementChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequirementChange is a log parse operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_MultiSigWallet *MultiSigWalletFilterer) ParseRequirementChange(log types.Log) (*MultiSigWalletRequirementChange, error) {
	event := new(MultiSigWalletRequirementChange)
	if err := _MultiSigWallet.contract.UnpackLog(event, "RequirementChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigWalletRevocationIterator is returned from FilterRevocation and is used to iterate over the raw logs and unpacked data for Revocation events raised by the MultiSigWallet contract.
type MultiSigWalletRevocationIterator struct {
	Event *MultiSigWalletRevocation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletRevocationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletRevocation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletRevocation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletRevocationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletRevocationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletRevocation represents a Revocation event raised by the MultiSigWallet contract.
type MultiSigWalletRevocation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRevocation is a free log retrieval operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWallet *MultiSigWalletFilterer) FilterRevocation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*MultiSigWalletRevocationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.FilterLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletRevocationIterator{contract: _MultiSigWallet.contract, event: "Revocation", logs: logs, sub: sub}, nil
}

// WatchRevocation is a free log subscription operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWallet *MultiSigWalletFilterer) WatchRevocation(opts *bind.WatchOpts, sink chan<- *MultiSigWalletRevocation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.WatchLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletRevocation)
				if err := _MultiSigWallet.contract.UnpackLog(event, "Revocation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevocation is a log parse operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWallet *MultiSigWalletFilterer) ParseRevocation(log types.Log) (*MultiSigWalletRevocation, error) {
	event := new(MultiSigWalletRevocation)
	if err := _MultiSigWallet.contract.UnpackLog(event, "Revocation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigWalletSubmissionIterator is returned from FilterSubmission and is used to iterate over the raw logs and unpacked data for Submission events raised by the MultiSigWallet contract.
type MultiSigWalletSubmissionIterator struct {
	Event *MultiSigWalletSubmission // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletSubmission)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletSubmission)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletSubmission represents a Submission event raised by the MultiSigWallet contract.
type MultiSigWalletSubmission struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSubmission is a free log retrieval operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_MultiSigWallet *MultiSigWalletFilterer) FilterSubmission(opts *bind.FilterOpts, transactionId []*big.Int) (*MultiSigWalletSubmissionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.FilterLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletSubmissionIterator{contract: _MultiSigWallet.contract, event: "Submission", logs: logs, sub: sub}, nil
}

// WatchSubmission is a free log subscription operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_MultiSigWallet *MultiSigWalletFilterer) WatchSubmission(opts *bind.WatchOpts, sink chan<- *MultiSigWalletSubmission, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWallet.contract.WatchLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletSubmission)
				if err := _MultiSigWallet.contract.UnpackLog(event, "Submission", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubmission is a log parse operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_MultiSigWallet *MultiSigWalletFilterer) ParseSubmission(log types.Log) (*MultiSigWalletSubmission, error) {
	event := new(MultiSigWalletSubmission)
	if err := _MultiSigWallet.contract.UnpackLog(event, "Submission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigWalletWithDailyLimitMetaData contains all meta data concerning the MultiSigWalletWithDailyLimit contract.
var MultiSigWalletWithDailyLimitMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"removeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"revokeConfirmation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"confirmations\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"calcMaxWithdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"pending\",\"type\":\"bool\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dailyLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastDay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"addOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"isConfirmed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmationCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"from\",\"type\":\"uint256\"},{\"name\":\"to\",\"type\":\"uint256\"},{\"name\":\"pending\",\"type\":\"bool\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionIds\",\"outputs\":[{\"name\":\"_transactionIds\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmations\",\"outputs\":[{\"name\":\"_confirmations\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transactionCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_required\",\"type\":\"uint256\"}],\"name\":\"changeRequirement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"confirmTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"submitTransaction\",\"outputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_dailyLimit\",\"type\":\"uint256\"}],\"name\":\"changeDailyLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OWNER_COUNT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"required\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"replaceOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spentToday\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"},{\"name\":\"_required\",\"type\":\"uint256\"},{\"name\":\"_dailyLimit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"dailyLimit\",\"type\":\"uint256\"}],\"name\":\"DailyLimitChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Confirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Revocation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Submission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Execution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"ExecutionFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"RequirementChange\",\"type\":\"event\"}]",
	Bin: "0x606060405234156200001057600080fd5b60405162001d5f38038062001d5f8339810160405280805182019190602001805191906020018051915083905082600082518260328211158015620000555750818111155b80156200006157508015155b80156200006d57508115155b15156200007957600080fd5b600092505b84518310156200014a57600260008685815181106200009957fe5b90602001906020020151600160a060020a0316815260208101919091526040016000205460ff16158015620000ec5750848381518110620000d657fe5b90602001906020020151600160a060020a031615155b1515620000f857600080fd5b6001600260008786815181106200010b57fe5b90602001906020020151600160a060020a031681526020810191909152604001600020805460ff1916911515919091179055600192909201916200007e565b60038580516200015f92916020019062000176565b5050506004919091555050600655506200020c9050565b828054828255906000526020600020908101928215620001d0579160200282015b82811115620001d05782518254600160a060020a031916600160a060020a03919091161782556020929092019160019091019062000197565b50620001de929150620001e2565b5090565b6200020991905b80821115620001de578054600160a060020a0319168155600101620001e9565b90565b611b43806200021c6000396000f3006060604052600436106101535763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663025e7c2781146101a9578063173825d9146101e857806320ea8d86146102145780632f54bf6e1461022a5780633411c81c1461026a5780634bc9fdc21461029957806354741525146102be57806367eeba0c146102db5780636b0c932d146102ee5780637065cb4814610301578063784547a71461032d5780638b51d13f146103435780639ace38c214610359578063a0e67e2b14610443578063a8abe69a146104a9578063b5dc40c3146104cc578063b77bf600146104e2578063ba51a6df146104f5578063c01a8c841461050b578063c642747414610521578063cea0862114610593578063d74f8edd146105a9578063dc8452cd146105bc578063e20056e6146105cf578063ee22610b14610601578063f059cf2b14610617575b60003411156101a7573373ffffffffffffffffffffffffffffffffffffffff167fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c3460405190815260200160405180910390a25b005b34156101b457600080fd5b6101bf60043561062a565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156101f357600080fd5b6101a773ffffffffffffffffffffffffffffffffffffffff6004351661065f565b341561021f57600080fd5b6101a76004356108ec565b341561023557600080fd5b61025673ffffffffffffffffffffffffffffffffffffffff60043516610a0f565b604051901515815260200160405180910390f35b341561027557600080fd5b61025660043573ffffffffffffffffffffffffffffffffffffffff60243516610a24565b34156102a457600080fd5b6102ac610a44565b60405190815260200160405180910390f35b34156102c957600080fd5b6102ac60043515156024351515610a7e565b34156102e657600080fd5b6102ac610aea565b34156102f957600080fd5b6102ac610af0565b341561030c57600080fd5b6101a773ffffffffffffffffffffffffffffffffffffffff60043516610af6565b341561033857600080fd5b610256600435610ca9565b341561034e57600080fd5b6102ac600435610d3a565b341561036457600080fd5b61036f600435610db6565b60405173ffffffffffffffffffffffffffffffffffffffff85168152602081018490528115156060820152608060408201818152845460027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61010060018416150201909116049183018290529060a0830190859080156104315780601f1061040657610100808354040283529160200191610431565b820191906000526020600020905b81548152906001019060200180831161041457829003601f168201915b50509550505050505060405180910390f35b341561044e57600080fd5b610456610df7565b60405160208082528190810183818151815260200191508051906020019060200280838360005b8381101561049557808201518382015260200161047d565b505050509050019250505060405180910390f35b34156104b457600080fd5b61045660043560243560443515156064351515610e6c565b34156104d757600080fd5b610456600435610f94565b34156104ed57600080fd5b6102ac61112c565b341561050057600080fd5b6101a7600435611132565b341561051657600080fd5b6101a76004356111df565b341561052c57600080fd5b6102ac6004803573ffffffffffffffffffffffffffffffffffffffff169060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061131f95505050505050565b341561059e57600080fd5b6101a760043561133e565b34156105b457600080fd5b6102ac6113b3565b34156105c757600080fd5b6102ac6113b8565b34156105da57600080fd5b6101a773ffffffffffffffffffffffffffffffffffffffff600435811690602435166113be565b341561060c57600080fd5b6101a760043561160a565b341561062257600080fd5b6102ac6118a9565b600380548290811061063857fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b60003073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561069b57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8216600090815260026020526040902054829060ff1615156106d157600080fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260026020526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905591505b6003547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0182101561085a578273ffffffffffffffffffffffffffffffffffffffff1660038381548110151561077157fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16141561084f57600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81019081106107c957fe5b6000918252602090912001546003805473ffffffffffffffffffffffffffffffffffffffff90921691849081106107fc57fe5b600091825260209091200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905561085a565b60019091019061071f565b600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019061088b9082611a44565b5060035460045411156108a4576003546108a490611132565b8273ffffffffffffffffffffffffffffffffffffffff167f8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b9060405160405180910390a2505050565b3373ffffffffffffffffffffffffffffffffffffffff811660009081526002602052604090205460ff16151561092157600080fd5b60008281526001602090815260408083203373ffffffffffffffffffffffffffffffffffffffff8116855292529091205483919060ff16151561096357600080fd5b600084815260208190526040902060030154849060ff161561098457600080fd5b600085815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff331680855292529182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905586917ff6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9905160405180910390a35050505050565b60026020526000908152604090205460ff1681565b600160209081526000928352604080842090915290825290205460ff1681565b60006007546201518001421115610a5e5750600654610a7b565b6008546006541015610a7257506000610a7b565b50600854600654035b90565b6000805b600554811015610ae357838015610aab575060008181526020819052604090206003015460ff16155b80610acf5750828015610acf575060008181526020819052604090206003015460ff165b15610adb576001820191505b600101610a82565b5092915050565b60065481565b60075481565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610b3057600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116600090815260026020526040902054819060ff1615610b6557600080fd5b8173ffffffffffffffffffffffffffffffffffffffff81161515610b8857600080fd5b60038054905060010160045460328211158015610ba55750818111155b8015610bb057508015155b8015610bbb57508115155b1515610bc657600080fd5b73ffffffffffffffffffffffffffffffffffffffff8516600090815260026020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660019081179091556003805490918101610c298382611a44565b50600091825260209091200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff87169081179091557ff39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d60405160405180910390a25050505050565b600080805b600354811015610d335760008481526001602052604081206003805491929184908110610cd757fe5b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff16835282019290925260400190205460ff1615610d18576001820191505b600454821415610d2b5760019250610d33565b600101610cae565b5050919050565b6000805b600354811015610db05760008381526001602052604081206003805491929184908110610d6757fe5b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff16835282019290925260400190205460ff1615610da8576001820191505b600101610d3e565b50919050565b600060208190529081526040902080546001820154600383015473ffffffffffffffffffffffffffffffffffffffff90921692909160029091019060ff1684565b610dff611a6d565b6003805480602002602001604051908101604052809291908181526020018280548015610e6257602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610e37575b5050505050905090565b610e74611a6d565b610e7c611a6d565b600080600554604051805910610e8f5750595b9080825280602002602001820160405250925060009150600090505b600554811015610f2457858015610ed4575060008181526020819052604090206003015460ff16155b80610ef85750848015610ef8575060008181526020819052604090206003015460ff165b15610f1c5780838381518110610f0a57fe5b60209081029091010152600191909101905b600101610eab565b878703604051805910610f345750595b908082528060200260200182016040525093508790505b86811015610f8957828181518110610f5f57fe5b906020019060200201518489830381518110610f7757fe5b60209081029091010152600101610f4b565b505050949350505050565b610f9c611a6d565b610fa4611a6d565b6003546000908190604051805910610fb95750595b9080825280602002602001820160405250925060009150600090505b6003548110156110a85760008581526001602052604081206003805491929184908110610ffe57fe5b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff16835282019290925260400190205460ff16156110a057600380548290811061104657fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1683838151811061107457fe5b73ffffffffffffffffffffffffffffffffffffffff909216602092830290910190910152600191909101905b600101610fd5565b816040518059106110b65750595b90808252806020026020018201604052509350600090505b81811015611124578281815181106110e257fe5b906020019060200201518482815181106110f857fe5b73ffffffffffffffffffffffffffffffffffffffff9092166020928302909101909101526001016110ce565b505050919050565b60055481565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561116c57600080fd5b60035481603282118015906111815750818111155b801561118c57508015155b801561119757508115155b15156111a257600080fd5b60048390557fa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a8360405190815260200160405180910390a1505050565b3373ffffffffffffffffffffffffffffffffffffffff811660009081526002602052604090205460ff16151561121457600080fd5b600082815260208190526040902054829073ffffffffffffffffffffffffffffffffffffffff16151561124657600080fd5b60008381526001602090815260408083203373ffffffffffffffffffffffffffffffffffffffff8116855292529091205484919060ff161561128757600080fd5b600085815260016020818152604080842073ffffffffffffffffffffffffffffffffffffffff331680865292529283902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690921790915586917f4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef905160405180910390a36113188561160a565b5050505050565b600061132c8484846118af565b9050611337816111df565b9392505050565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561137857600080fd5b60068190557fc71bdc6afaf9b1aa90a7078191d4fc1adf3bf680fca3183697df6b0dc226bca28160405190815260200160405180910390a150565b603281565b60045481565b60003073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156113fa57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260026020526040902054839060ff16151561143057600080fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260026020526040902054839060ff161561146557600080fd5b600092505b600354831015611530578473ffffffffffffffffffffffffffffffffffffffff1660038481548110151561149a57fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16141561152557836003848154811015156114d257fe5b600091825260209091200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055611530565b60019092019161146a565b73ffffffffffffffffffffffffffffffffffffffff80861660008181526002602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff009081169091559388168252908190208054909316600117909255907f8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90905160405180910390a28373ffffffffffffffffffffffffffffffffffffffff167ff39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d60405160405180910390a25050505050565b3373ffffffffffffffffffffffffffffffffffffffff81166000908152600260205260408120549091829160ff16151561164357600080fd5b60008481526001602090815260408083203373ffffffffffffffffffffffffffffffffffffffff8116855292529091205485919060ff16151561168557600080fd5b600086815260208190526040902060030154869060ff16156116a657600080fd5b600087815260208190526040902095506116bf87610ca9565b9450848061171057506002808701547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff610100600183161502011604158015611710575061171086600101546119fc565b156118a0576003860180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558415156117585760018601546008805490910190555b8554600187015473ffffffffffffffffffffffffffffffffffffffff909116906002880160405180828054600181600116156101000203166002900480156117e15780601f106117b6576101008083540402835291602001916117e1565b820191906000526020600020905b8154815290600101906020018083116117c457829003601f168201915b505091505060006040518083038185876187965a03f1925050501561183257867f33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed7560405160405180910390a26118a0565b867f526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b7923660405160405180910390a26003860180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558415156118a0576001860154600880549190910390555b50505050505050565b60085481565b60008373ffffffffffffffffffffffffffffffffffffffff811615156118d457600080fd5b60055491506080604051908101604090815273ffffffffffffffffffffffffffffffffffffffff8716825260208083018790528183018690526000606084018190528581529081905220815181547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9190911617815560208201518160010155604082015181600201908051611984929160200190611a7f565b50606082015160039190910180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001691151591909117905550600580546001019055817fc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e5160405160405180910390a2509392505050565b60006007546201518001421115611a17574260075560006008555b60065482600854011180611a2e5750600854828101105b15611a3b57506000611a3f565b5060015b919050565b815481835581811511611a6857600083815260209020611a68918101908301611afd565b505050565b60206040519081016040526000815290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611ac057805160ff1916838001178555611aed565b82800160010185558215611aed579182015b82811115611aed578251825591602001919060010190611ad2565b50611af9929150611afd565b5090565b610a7b91905b80821115611af95760008155600101611b035600a165627a7a723058200539fd531f22087f46f0832a5aa805799b2bed0f9c39d750f323fbe532884ead0029",
}

// MultiSigWalletWithDailyLimitABI is the input ABI used to generate the binding from.
// Deprecated: Use MultiSigWalletWithDailyLimitMetaData.ABI instead.
var MultiSigWalletWithDailyLimitABI = MultiSigWalletWithDailyLimitMetaData.ABI

// MultiSigWalletWithDailyLimitBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MultiSigWalletWithDailyLimitMetaData.Bin instead.
var MultiSigWalletWithDailyLimitBin = MultiSigWalletWithDailyLimitMetaData.Bin

// DeployMultiSigWalletWithDailyLimit deploys a new Ethereum contract, binding an instance of MultiSigWalletWithDailyLimit to it.
func DeployMultiSigWalletWithDailyLimit(auth *bind.TransactOpts, backend bind.ContractBackend, _owners []common.Address, _required *big.Int, _dailyLimit *big.Int) (common.Address, *types.Transaction, *MultiSigWalletWithDailyLimit, error) {
	parsed, err := MultiSigWalletWithDailyLimitMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MultiSigWalletWithDailyLimitBin), backend, _owners, _required, _dailyLimit)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MultiSigWalletWithDailyLimit{MultiSigWalletWithDailyLimitCaller: MultiSigWalletWithDailyLimitCaller{contract: contract}, MultiSigWalletWithDailyLimitTransactor: MultiSigWalletWithDailyLimitTransactor{contract: contract}, MultiSigWalletWithDailyLimitFilterer: MultiSigWalletWithDailyLimitFilterer{contract: contract}}, nil
}

// MultiSigWalletWithDailyLimit is an auto generated Go binding around an Ethereum contract.
type MultiSigWalletWithDailyLimit struct {
	MultiSigWalletWithDailyLimitCaller     // Read-only binding to the contract
	MultiSigWalletWithDailyLimitTransactor // Write-only binding to the contract
	MultiSigWalletWithDailyLimitFilterer   // Log filterer for contract events
}

// MultiSigWalletWithDailyLimitCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiSigWalletWithDailyLimitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigWalletWithDailyLimitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiSigWalletWithDailyLimitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigWalletWithDailyLimitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiSigWalletWithDailyLimitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigWalletWithDailyLimitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiSigWalletWithDailyLimitSession struct {
	Contract     *MultiSigWalletWithDailyLimit // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MultiSigWalletWithDailyLimitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiSigWalletWithDailyLimitCallerSession struct {
	Contract *MultiSigWalletWithDailyLimitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// MultiSigWalletWithDailyLimitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiSigWalletWithDailyLimitTransactorSession struct {
	Contract     *MultiSigWalletWithDailyLimitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// MultiSigWalletWithDailyLimitRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiSigWalletWithDailyLimitRaw struct {
	Contract *MultiSigWalletWithDailyLimit // Generic contract binding to access the raw methods on
}

// MultiSigWalletWithDailyLimitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiSigWalletWithDailyLimitCallerRaw struct {
	Contract *MultiSigWalletWithDailyLimitCaller // Generic read-only contract binding to access the raw methods on
}

// MultiSigWalletWithDailyLimitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiSigWalletWithDailyLimitTransactorRaw struct {
	Contract *MultiSigWalletWithDailyLimitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiSigWalletWithDailyLimit creates a new instance of MultiSigWalletWithDailyLimit, bound to a specific deployed contract.
func NewMultiSigWalletWithDailyLimit(address common.Address, backend bind.ContractBackend) (*MultiSigWalletWithDailyLimit, error) {
	contract, err := bindMultiSigWalletWithDailyLimit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimit{MultiSigWalletWithDailyLimitCaller: MultiSigWalletWithDailyLimitCaller{contract: contract}, MultiSigWalletWithDailyLimitTransactor: MultiSigWalletWithDailyLimitTransactor{contract: contract}, MultiSigWalletWithDailyLimitFilterer: MultiSigWalletWithDailyLimitFilterer{contract: contract}}, nil
}

// NewMultiSigWalletWithDailyLimitCaller creates a new read-only instance of MultiSigWalletWithDailyLimit, bound to a specific deployed contract.
func NewMultiSigWalletWithDailyLimitCaller(address common.Address, caller bind.ContractCaller) (*MultiSigWalletWithDailyLimitCaller, error) {
	contract, err := bindMultiSigWalletWithDailyLimit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitCaller{contract: contract}, nil
}

// NewMultiSigWalletWithDailyLimitTransactor creates a new write-only instance of MultiSigWalletWithDailyLimit, bound to a specific deployed contract.
func NewMultiSigWalletWithDailyLimitTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiSigWalletWithDailyLimitTransactor, error) {
	contract, err := bindMultiSigWalletWithDailyLimit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitTransactor{contract: contract}, nil
}

// NewMultiSigWalletWithDailyLimitFilterer creates a new log filterer instance of MultiSigWalletWithDailyLimit, bound to a specific deployed contract.
func NewMultiSigWalletWithDailyLimitFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiSigWalletWithDailyLimitFilterer, error) {
	contract, err := bindMultiSigWalletWithDailyLimit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitFilterer{contract: contract}, nil
}

// bindMultiSigWalletWithDailyLimit binds a generic wrapper to an already deployed contract.
func bindMultiSigWalletWithDailyLimit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MultiSigWalletWithDailyLimitMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiSigWalletWithDailyLimit.Contract.MultiSigWalletWithDailyLimitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.MultiSigWalletWithDailyLimitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.MultiSigWalletWithDailyLimitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiSigWalletWithDailyLimit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.contract.Transact(opts, method, params...)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) MAXOWNERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, &out, "MAX_OWNER_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.MAXOWNERCOUNT(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.MAXOWNERCOUNT(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// CalcMaxWithdraw is a free data retrieval call binding the contract method 0x4bc9fdc2.
//
// Solidity: function calcMaxWithdraw() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) CalcMaxWithdraw(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, &out, "calcMaxWithdraw")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcMaxWithdraw is a free data retrieval call binding the contract method 0x4bc9fdc2.
//
// Solidity: function calcMaxWithdraw() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) CalcMaxWithdraw() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.CalcMaxWithdraw(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// CalcMaxWithdraw is a free data retrieval call binding the contract method 0x4bc9fdc2.
//
// Solidity: function calcMaxWithdraw() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) CalcMaxWithdraw() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.CalcMaxWithdraw(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) view returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) Confirmations(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, &out, "confirmations", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) view returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Confirmations(&_MultiSigWalletWithDailyLimit.CallOpts, arg0, arg1)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) view returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Confirmations(&_MultiSigWalletWithDailyLimit.CallOpts, arg0, arg1)
}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) DailyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, &out, "dailyLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) DailyLimit() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.DailyLimit(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) DailyLimit() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.DailyLimit(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) view returns(uint256 count)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) GetConfirmationCount(opts *bind.CallOpts, transactionId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, &out, "getConfirmationCount", transactionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) view returns(uint256 count)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetConfirmationCount(&_MultiSigWalletWithDailyLimit.CallOpts, transactionId)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) view returns(uint256 count)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetConfirmationCount(&_MultiSigWalletWithDailyLimit.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) view returns(address[] _confirmations)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) GetConfirmations(opts *bind.CallOpts, transactionId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, &out, "getConfirmations", transactionId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) view returns(address[] _confirmations)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetConfirmations(&_MultiSigWalletWithDailyLimit.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) view returns(address[] _confirmations)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetConfirmations(&_MultiSigWalletWithDailyLimit.CallOpts, transactionId)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, &out, "getOwners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) GetOwners() ([]common.Address, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetOwners(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) GetOwners() ([]common.Address, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetOwners(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) view returns(uint256 count)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) GetTransactionCount(opts *bind.CallOpts, pending bool, executed bool) (*big.Int, error) {
	var out []interface{}
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, &out, "getTransactionCount", pending, executed)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) view returns(uint256 count)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetTransactionCount(&_MultiSigWalletWithDailyLimit.CallOpts, pending, executed)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) view returns(uint256 count)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetTransactionCount(&_MultiSigWalletWithDailyLimit.CallOpts, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) view returns(uint256[] _transactionIds)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) GetTransactionIds(opts *bind.CallOpts, from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	var out []interface{}
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, &out, "getTransactionIds", from, to, pending, executed)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) view returns(uint256[] _transactionIds)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetTransactionIds(&_MultiSigWalletWithDailyLimit.CallOpts, from, to, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) view returns(uint256[] _transactionIds)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.GetTransactionIds(&_MultiSigWalletWithDailyLimit.CallOpts, from, to, pending, executed)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) view returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) IsConfirmed(opts *bind.CallOpts, transactionId *big.Int) (bool, error) {
	var out []interface{}
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, &out, "isConfirmed", transactionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) view returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _MultiSigWalletWithDailyLimit.Contract.IsConfirmed(&_MultiSigWalletWithDailyLimit.CallOpts, transactionId)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) view returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _MultiSigWalletWithDailyLimit.Contract.IsConfirmed(&_MultiSigWalletWithDailyLimit.CallOpts, transactionId)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) IsOwner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, &out, "isOwner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) IsOwner(arg0 common.Address) (bool, error) {
	return _MultiSigWalletWithDailyLimit.Contract.IsOwner(&_MultiSigWalletWithDailyLimit.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) IsOwner(arg0 common.Address) (bool, error) {
	return _MultiSigWalletWithDailyLimit.Contract.IsOwner(&_MultiSigWalletWithDailyLimit.CallOpts, arg0)
}

// LastDay is a free data retrieval call binding the contract method 0x6b0c932d.
//
// Solidity: function lastDay() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) LastDay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, &out, "lastDay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastDay is a free data retrieval call binding the contract method 0x6b0c932d.
//
// Solidity: function lastDay() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) LastDay() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.LastDay(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// LastDay is a free data retrieval call binding the contract method 0x6b0c932d.
//
// Solidity: function lastDay() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) LastDay() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.LastDay(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Owners(&_MultiSigWalletWithDailyLimit.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Owners(&_MultiSigWalletWithDailyLimit.CallOpts, arg0)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) Required(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, &out, "required")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) Required() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Required(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) Required() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Required(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// SpentToday is a free data retrieval call binding the contract method 0xf059cf2b.
//
// Solidity: function spentToday() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) SpentToday(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, &out, "spentToday")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpentToday is a free data retrieval call binding the contract method 0xf059cf2b.
//
// Solidity: function spentToday() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) SpentToday() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.SpentToday(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// SpentToday is a free data retrieval call binding the contract method 0xf059cf2b.
//
// Solidity: function spentToday() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) SpentToday() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.SpentToday(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) TransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, &out, "transactionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) TransactionCount() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.TransactionCount(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) TransactionCount() (*big.Int, error) {
	return _MultiSigWalletWithDailyLimit.Contract.TransactionCount(&_MultiSigWalletWithDailyLimit.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address destination, uint256 value, bytes data, bool executed)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	var out []interface{}
	err := _MultiSigWalletWithDailyLimit.contract.Call(opts, &out, "transactions", arg0)

	outstruct := new(struct {
		Destination common.Address
		Value       *big.Int
		Data        []byte
		Executed    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Destination = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Data = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Executed = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address destination, uint256 value, bytes data, bool executed)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Transactions(&_MultiSigWalletWithDailyLimit.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address destination, uint256 value, bytes data, bool executed)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitCallerSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Transactions(&_MultiSigWalletWithDailyLimit.CallOpts, arg0)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) AddOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "addOwner", owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.AddOwner(&_MultiSigWalletWithDailyLimit.TransactOpts, owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.AddOwner(&_MultiSigWalletWithDailyLimit.TransactOpts, owner)
}

// ChangeDailyLimit is a paid mutator transaction binding the contract method 0xcea08621.
//
// Solidity: function changeDailyLimit(uint256 _dailyLimit) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) ChangeDailyLimit(opts *bind.TransactOpts, _dailyLimit *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "changeDailyLimit", _dailyLimit)
}

// ChangeDailyLimit is a paid mutator transaction binding the contract method 0xcea08621.
//
// Solidity: function changeDailyLimit(uint256 _dailyLimit) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) ChangeDailyLimit(_dailyLimit *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ChangeDailyLimit(&_MultiSigWalletWithDailyLimit.TransactOpts, _dailyLimit)
}

// ChangeDailyLimit is a paid mutator transaction binding the contract method 0xcea08621.
//
// Solidity: function changeDailyLimit(uint256 _dailyLimit) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) ChangeDailyLimit(_dailyLimit *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ChangeDailyLimit(&_MultiSigWalletWithDailyLimit.TransactOpts, _dailyLimit)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) ChangeRequirement(opts *bind.TransactOpts, _required *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "changeRequirement", _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ChangeRequirement(&_MultiSigWalletWithDailyLimit.TransactOpts, _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ChangeRequirement(&_MultiSigWalletWithDailyLimit.TransactOpts, _required)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) ConfirmTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "confirmTransaction", transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ConfirmTransaction(&_MultiSigWalletWithDailyLimit.TransactOpts, transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ConfirmTransaction(&_MultiSigWalletWithDailyLimit.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) ExecuteTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "executeTransaction", transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ExecuteTransaction(&_MultiSigWalletWithDailyLimit.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ExecuteTransaction(&_MultiSigWalletWithDailyLimit.TransactOpts, transactionId)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) RemoveOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "removeOwner", owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.RemoveOwner(&_MultiSigWalletWithDailyLimit.TransactOpts, owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.RemoveOwner(&_MultiSigWalletWithDailyLimit.TransactOpts, owner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) ReplaceOwner(opts *bind.TransactOpts, owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "replaceOwner", owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ReplaceOwner(&_MultiSigWalletWithDailyLimit.TransactOpts, owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.ReplaceOwner(&_MultiSigWalletWithDailyLimit.TransactOpts, owner, newOwner)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) RevokeConfirmation(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "revokeConfirmation", transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.RevokeConfirmation(&_MultiSigWalletWithDailyLimit.TransactOpts, transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.RevokeConfirmation(&_MultiSigWalletWithDailyLimit.TransactOpts, transactionId)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) SubmitTransaction(opts *bind.TransactOpts, destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.Transact(opts, "submitTransaction", destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.SubmitTransaction(&_MultiSigWalletWithDailyLimit.TransactOpts, destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.SubmitTransaction(&_MultiSigWalletWithDailyLimit.TransactOpts, destination, value, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Fallback(&_MultiSigWalletWithDailyLimit.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MultiSigWalletWithDailyLimit.Contract.Fallback(&_MultiSigWalletWithDailyLimit.TransactOpts, calldata)
}

// MultiSigWalletWithDailyLimitConfirmationIterator is returned from FilterConfirmation and is used to iterate over the raw logs and unpacked data for Confirmation events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitConfirmationIterator struct {
	Event *MultiSigWalletWithDailyLimitConfirmation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitConfirmation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitConfirmation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitConfirmation represents a Confirmation event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitConfirmation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterConfirmation is a free log retrieval operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterConfirmation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*MultiSigWalletWithDailyLimitConfirmationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitConfirmationIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "Confirmation", logs: logs, sub: sub}, nil
}

// WatchConfirmation is a free log subscription operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchConfirmation(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitConfirmation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitConfirmation)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "Confirmation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConfirmation is a log parse operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) ParseConfirmation(log types.Log) (*MultiSigWalletWithDailyLimitConfirmation, error) {
	event := new(MultiSigWalletWithDailyLimitConfirmation)
	if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "Confirmation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigWalletWithDailyLimitDailyLimitChangeIterator is returned from FilterDailyLimitChange and is used to iterate over the raw logs and unpacked data for DailyLimitChange events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitDailyLimitChangeIterator struct {
	Event *MultiSigWalletWithDailyLimitDailyLimitChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitDailyLimitChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitDailyLimitChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitDailyLimitChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitDailyLimitChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitDailyLimitChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitDailyLimitChange represents a DailyLimitChange event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitDailyLimitChange struct {
	DailyLimit *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDailyLimitChange is a free log retrieval operation binding the contract event 0xc71bdc6afaf9b1aa90a7078191d4fc1adf3bf680fca3183697df6b0dc226bca2.
//
// Solidity: event DailyLimitChange(uint256 dailyLimit)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterDailyLimitChange(opts *bind.FilterOpts) (*MultiSigWalletWithDailyLimitDailyLimitChangeIterator, error) {

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "DailyLimitChange")
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitDailyLimitChangeIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "DailyLimitChange", logs: logs, sub: sub}, nil
}

// WatchDailyLimitChange is a free log subscription operation binding the contract event 0xc71bdc6afaf9b1aa90a7078191d4fc1adf3bf680fca3183697df6b0dc226bca2.
//
// Solidity: event DailyLimitChange(uint256 dailyLimit)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchDailyLimitChange(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitDailyLimitChange) (event.Subscription, error) {

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "DailyLimitChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitDailyLimitChange)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "DailyLimitChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDailyLimitChange is a log parse operation binding the contract event 0xc71bdc6afaf9b1aa90a7078191d4fc1adf3bf680fca3183697df6b0dc226bca2.
//
// Solidity: event DailyLimitChange(uint256 dailyLimit)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) ParseDailyLimitChange(log types.Log) (*MultiSigWalletWithDailyLimitDailyLimitChange, error) {
	event := new(MultiSigWalletWithDailyLimitDailyLimitChange)
	if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "DailyLimitChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigWalletWithDailyLimitDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitDepositIterator struct {
	Event *MultiSigWalletWithDailyLimitDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitDeposit represents a Deposit event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitDeposit struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*MultiSigWalletWithDailyLimitDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitDepositIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitDeposit)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) ParseDeposit(log types.Log) (*MultiSigWalletWithDailyLimitDeposit, error) {
	event := new(MultiSigWalletWithDailyLimitDeposit)
	if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigWalletWithDailyLimitExecutionIterator is returned from FilterExecution and is used to iterate over the raw logs and unpacked data for Execution events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitExecutionIterator struct {
	Event *MultiSigWalletWithDailyLimitExecution // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitExecution)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitExecution)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitExecution represents a Execution event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitExecution struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecution is a free log retrieval operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterExecution(opts *bind.FilterOpts, transactionId []*big.Int) (*MultiSigWalletWithDailyLimitExecutionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitExecutionIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "Execution", logs: logs, sub: sub}, nil
}

// WatchExecution is a free log subscription operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchExecution(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitExecution, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitExecution)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "Execution", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecution is a log parse operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) ParseExecution(log types.Log) (*MultiSigWalletWithDailyLimitExecution, error) {
	event := new(MultiSigWalletWithDailyLimitExecution)
	if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "Execution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigWalletWithDailyLimitExecutionFailureIterator is returned from FilterExecutionFailure and is used to iterate over the raw logs and unpacked data for ExecutionFailure events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitExecutionFailureIterator struct {
	Event *MultiSigWalletWithDailyLimitExecutionFailure // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitExecutionFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitExecutionFailure)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitExecutionFailure)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitExecutionFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitExecutionFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitExecutionFailure represents a ExecutionFailure event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitExecutionFailure struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionFailure is a free log retrieval operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterExecutionFailure(opts *bind.FilterOpts, transactionId []*big.Int) (*MultiSigWalletWithDailyLimitExecutionFailureIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "ExecutionFailure", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitExecutionFailureIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "ExecutionFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFailure is a free log subscription operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchExecutionFailure(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitExecutionFailure, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "ExecutionFailure", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitExecutionFailure)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutionFailure is a log parse operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) ParseExecutionFailure(log types.Log) (*MultiSigWalletWithDailyLimitExecutionFailure, error) {
	event := new(MultiSigWalletWithDailyLimitExecutionFailure)
	if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigWalletWithDailyLimitOwnerAdditionIterator is returned from FilterOwnerAddition and is used to iterate over the raw logs and unpacked data for OwnerAddition events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitOwnerAdditionIterator struct {
	Event *MultiSigWalletWithDailyLimitOwnerAddition // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitOwnerAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitOwnerAddition)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitOwnerAddition)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitOwnerAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitOwnerAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitOwnerAddition represents a OwnerAddition event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitOwnerAddition struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerAddition is a free log retrieval operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterOwnerAddition(opts *bind.FilterOpts, owner []common.Address) (*MultiSigWalletWithDailyLimitOwnerAdditionIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitOwnerAdditionIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "OwnerAddition", logs: logs, sub: sub}, nil
}

// WatchOwnerAddition is a free log subscription operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchOwnerAddition(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitOwnerAddition, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitOwnerAddition)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "OwnerAddition", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnerAddition is a log parse operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) ParseOwnerAddition(log types.Log) (*MultiSigWalletWithDailyLimitOwnerAddition, error) {
	event := new(MultiSigWalletWithDailyLimitOwnerAddition)
	if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "OwnerAddition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigWalletWithDailyLimitOwnerRemovalIterator is returned from FilterOwnerRemoval and is used to iterate over the raw logs and unpacked data for OwnerRemoval events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitOwnerRemovalIterator struct {
	Event *MultiSigWalletWithDailyLimitOwnerRemoval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitOwnerRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitOwnerRemoval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitOwnerRemoval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitOwnerRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitOwnerRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitOwnerRemoval represents a OwnerRemoval event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitOwnerRemoval struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerRemoval is a free log retrieval operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterOwnerRemoval(opts *bind.FilterOpts, owner []common.Address) (*MultiSigWalletWithDailyLimitOwnerRemovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitOwnerRemovalIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "OwnerRemoval", logs: logs, sub: sub}, nil
}

// WatchOwnerRemoval is a free log subscription operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchOwnerRemoval(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitOwnerRemoval, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitOwnerRemoval)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "OwnerRemoval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnerRemoval is a log parse operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) ParseOwnerRemoval(log types.Log) (*MultiSigWalletWithDailyLimitOwnerRemoval, error) {
	event := new(MultiSigWalletWithDailyLimitOwnerRemoval)
	if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "OwnerRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigWalletWithDailyLimitRequirementChangeIterator is returned from FilterRequirementChange and is used to iterate over the raw logs and unpacked data for RequirementChange events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitRequirementChangeIterator struct {
	Event *MultiSigWalletWithDailyLimitRequirementChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitRequirementChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitRequirementChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitRequirementChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitRequirementChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitRequirementChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitRequirementChange represents a RequirementChange event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitRequirementChange struct {
	Required *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequirementChange is a free log retrieval operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterRequirementChange(opts *bind.FilterOpts) (*MultiSigWalletWithDailyLimitRequirementChangeIterator, error) {

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitRequirementChangeIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "RequirementChange", logs: logs, sub: sub}, nil
}

// WatchRequirementChange is a free log subscription operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchRequirementChange(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitRequirementChange) (event.Subscription, error) {

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitRequirementChange)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "RequirementChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequirementChange is a log parse operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) ParseRequirementChange(log types.Log) (*MultiSigWalletWithDailyLimitRequirementChange, error) {
	event := new(MultiSigWalletWithDailyLimitRequirementChange)
	if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "RequirementChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigWalletWithDailyLimitRevocationIterator is returned from FilterRevocation and is used to iterate over the raw logs and unpacked data for Revocation events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitRevocationIterator struct {
	Event *MultiSigWalletWithDailyLimitRevocation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitRevocationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitRevocation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitRevocation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitRevocationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitRevocationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitRevocation represents a Revocation event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitRevocation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRevocation is a free log retrieval operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterRevocation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*MultiSigWalletWithDailyLimitRevocationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitRevocationIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "Revocation", logs: logs, sub: sub}, nil
}

// WatchRevocation is a free log subscription operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchRevocation(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitRevocation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitRevocation)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "Revocation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevocation is a log parse operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) ParseRevocation(log types.Log) (*MultiSigWalletWithDailyLimitRevocation, error) {
	event := new(MultiSigWalletWithDailyLimitRevocation)
	if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "Revocation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigWalletWithDailyLimitSubmissionIterator is returned from FilterSubmission and is used to iterate over the raw logs and unpacked data for Submission events raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitSubmissionIterator struct {
	Event *MultiSigWalletWithDailyLimitSubmission // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletWithDailyLimitSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletWithDailyLimitSubmission)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletWithDailyLimitSubmission)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletWithDailyLimitSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletWithDailyLimitSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletWithDailyLimitSubmission represents a Submission event raised by the MultiSigWalletWithDailyLimit contract.
type MultiSigWalletWithDailyLimitSubmission struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSubmission is a free log retrieval operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) FilterSubmission(opts *bind.FilterOpts, transactionId []*big.Int) (*MultiSigWalletWithDailyLimitSubmissionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.FilterLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletWithDailyLimitSubmissionIterator{contract: _MultiSigWalletWithDailyLimit.contract, event: "Submission", logs: logs, sub: sub}, nil
}

// WatchSubmission is a free log subscription operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) WatchSubmission(opts *bind.WatchOpts, sink chan<- *MultiSigWalletWithDailyLimitSubmission, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletWithDailyLimit.contract.WatchLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletWithDailyLimitSubmission)
				if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "Submission", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubmission is a log parse operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_MultiSigWalletWithDailyLimit *MultiSigWalletWithDailyLimitFilterer) ParseSubmission(log types.Log) (*MultiSigWalletWithDailyLimitSubmission, error) {
	event := new(MultiSigWalletWithDailyLimitSubmission)
	if err := _MultiSigWalletWithDailyLimit.contract.UnpackLog(event, "Submission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a72305820b241df596497ee656076844b8bd2919978aa006b6a635fab8e54cebe0397b8070029",
}

// SafeMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMathMetaData.ABI instead.
var SafeMathABI = SafeMathMetaData.ABI

// SafeMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMathMetaData.Bin instead.
var SafeMathBin = SafeMathMetaData.Bin

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// StandardTokenMetaData contains all meta data concerning the StandardToken contract.
var StandardTokenMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]",
	Bin: "0x6060604052341561000f57600080fd5b6105868061001e6000396000f3006060604052600436106100825763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461008757806318160ddd146100ca57806323b872dd146100ef57806370a0823114610124578063a9059cbb14610150578063dd62ed3e1461017f578063eefa597b146101b1575b600080fd5b341561009257600080fd5b6100b673ffffffffffffffffffffffffffffffffffffffff600435166024356101c4565b604051901515815260200160405180910390f35b34156100d557600080fd5b6100dd610286565b60405190815260200160405180910390f35b34156100fa57600080fd5b6100b673ffffffffffffffffffffffffffffffffffffffff6004358116906024351660443561028c565b341561012f57600080fd5b6100dd73ffffffffffffffffffffffffffffffffffffffff600435166103c1565b341561015b57600080fd5b6100b673ffffffffffffffffffffffffffffffffffffffff600435166024356103e9565b341561018a57600080fd5b6100dd73ffffffffffffffffffffffffffffffffffffffff600435811690602435166104d6565b34156101bc57600080fd5b6100b661050e565b60008115801590610206575073ffffffffffffffffffffffffffffffffffffffff33811660009081526002602090815260408083209387168352929052205415155b1561021057600080fd5b73ffffffffffffffffffffffffffffffffffffffff338116600081815260026020908152604080832094881680845294909152908190208590557f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b73ffffffffffffffffffffffffffffffffffffffff80841660009081526002602090815260408083203385168452825280832054938616835260019091528120549091906102da9084610513565b73ffffffffffffffffffffffffffffffffffffffff80861660009081526001602052604080822093909355908716815220546103169084610537565b73ffffffffffffffffffffffffffffffffffffffff86166000908152600160205260409020556103468184610537565b73ffffffffffffffffffffffffffffffffffffffff808716600081815260026020908152604080832033861684529091529081902093909355908616917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9086905190815260200160405180910390a3506001949350505050565b73ffffffffffffffffffffffffffffffffffffffff1660009081526001602052604090205490565b6000604060443610156103fb57600080fd5b73ffffffffffffffffffffffffffffffffffffffff331660009081526001602052604090205461042b9084610537565b73ffffffffffffffffffffffffffffffffffffffff33811660009081526001602052604080822093909355908616815220546104679084610513565b73ffffffffffffffffffffffffffffffffffffffff80861660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9086905190815260200160405180910390a35060019392505050565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260026020908152604080832093909416825291909152205490565b600181565b600082820161053084821080159061052b5750838210155b61054b565b9392505050565b60006105458383111561054b565b50900390565b80151561055757600080fd5b505600a165627a7a72305820e39ad092e71e0d8ff62cf0c15a0e791cf98d73fe98ed33e2829aecbdafb865dd0029",
}

// StandardTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use StandardTokenMetaData.ABI instead.
var StandardTokenABI = StandardTokenMetaData.ABI

// StandardTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StandardTokenMetaData.Bin instead.
var StandardTokenBin = StandardTokenMetaData.Bin

// DeployStandardToken deploys a new Ethereum contract, binding an instance of StandardToken to it.
func DeployStandardToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StandardToken, error) {
	parsed, err := StandardTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StandardTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// StandardToken is an auto generated Go binding around an Ethereum contract.
type StandardToken struct {
	StandardTokenCaller     // Read-only binding to the contract
	StandardTokenTransactor // Write-only binding to the contract
	StandardTokenFilterer   // Log filterer for contract events
}

// StandardTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandardTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StandardTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardTokenSession struct {
	Contract     *StandardToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StandardTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardTokenCallerSession struct {
	Contract *StandardTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StandardTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardTokenTransactorSession struct {
	Contract     *StandardTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StandardTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StandardTokenRaw struct {
	Contract *StandardToken // Generic contract binding to access the raw methods on
}

// StandardTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardTokenCallerRaw struct {
	Contract *StandardTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StandardTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardTokenTransactorRaw struct {
	Contract *StandardTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardToken creates a new instance of StandardToken, bound to a specific deployed contract.
func NewStandardToken(address common.Address, backend bind.ContractBackend) (*StandardToken, error) {
	contract, err := bindStandardToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// NewStandardTokenCaller creates a new read-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenCaller(address common.Address, caller bind.ContractCaller) (*StandardTokenCaller, error) {
	contract, err := bindStandardToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenCaller{contract: contract}, nil
}

// NewStandardTokenTransactor creates a new write-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardTokenTransactor, error) {
	contract, err := bindStandardToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransactor{contract: contract}, nil
}

// NewStandardTokenFilterer creates a new log filterer instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*StandardTokenFilterer, error) {
	contract, err := bindStandardToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StandardTokenFilterer{contract: contract}, nil
}

// bindStandardToken binds a generic wrapper to an already deployed contract.
func bindStandardToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StandardTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.StandardTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_StandardToken *StandardTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StandardToken.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_StandardToken *StandardTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_StandardToken *StandardTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_StandardToken *StandardTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StandardToken.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_StandardToken *StandardTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_StandardToken *StandardTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// IsToken is a free data retrieval call binding the contract method 0xeefa597b.
//
// Solidity: function isToken() view returns(bool)
func (_StandardToken *StandardTokenCaller) IsToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StandardToken.contract.Call(opts, &out, "isToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsToken is a free data retrieval call binding the contract method 0xeefa597b.
//
// Solidity: function isToken() view returns(bool)
func (_StandardToken *StandardTokenSession) IsToken() (bool, error) {
	return _StandardToken.Contract.IsToken(&_StandardToken.CallOpts)
}

// IsToken is a free data retrieval call binding the contract method 0xeefa597b.
//
// Solidity: function isToken() view returns(bool)
func (_StandardToken *StandardTokenCallerSession) IsToken() (bool, error) {
	return _StandardToken.Contract.IsToken(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StandardToken *StandardTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StandardToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StandardToken *StandardTokenSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StandardToken *StandardTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_StandardToken *StandardTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_StandardToken *StandardTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_StandardToken *StandardTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_StandardToken *StandardTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_StandardToken *StandardTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_StandardToken *StandardTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_StandardToken *StandardTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_StandardToken *StandardTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_StandardToken *StandardTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// StandardTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StandardToken contract.
type StandardTokenApprovalIterator struct {
	Event *StandardTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StandardTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StandardTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StandardTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenApproval represents a Approval event raised by the StandardToken contract.
type StandardTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StandardToken *StandardTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StandardTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenApprovalIterator{contract: _StandardToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StandardToken *StandardTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StandardTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenApproval)
				if err := _StandardToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StandardToken *StandardTokenFilterer) ParseApproval(log types.Log) (*StandardTokenApproval, error) {
	event := new(StandardTokenApproval)
	if err := _StandardToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StandardToken contract.
type StandardTokenTransferIterator struct {
	Event *StandardTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StandardTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StandardTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StandardTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenTransfer represents a Transfer event raised by the StandardToken contract.
type StandardTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StandardToken *StandardTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StandardTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransferIterator{contract: _StandardToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StandardToken *StandardTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StandardTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenTransfer)
				if err := _StandardToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StandardToken *StandardTokenFilterer) ParseTransfer(log types.Log) (*StandardTokenTransfer, error) {
	event := new(StandardTokenTransfer)
	if err := _StandardToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeAgentMetaData contains all meta data concerning the UpgradeAgent contract.
var UpgradeAgentMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"originalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isUpgradeAgent\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"upgradeFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UpgradeAgentABI is the input ABI used to generate the binding from.
// Deprecated: Use UpgradeAgentMetaData.ABI instead.
var UpgradeAgentABI = UpgradeAgentMetaData.ABI

// UpgradeAgent is an auto generated Go binding around an Ethereum contract.
type UpgradeAgent struct {
	UpgradeAgentCaller     // Read-only binding to the contract
	UpgradeAgentTransactor // Write-only binding to the contract
	UpgradeAgentFilterer   // Log filterer for contract events
}

// UpgradeAgentCaller is an auto generated read-only Go binding around an Ethereum contract.
type UpgradeAgentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeAgentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UpgradeAgentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeAgentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UpgradeAgentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeAgentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UpgradeAgentSession struct {
	Contract     *UpgradeAgent     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UpgradeAgentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UpgradeAgentCallerSession struct {
	Contract *UpgradeAgentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// UpgradeAgentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UpgradeAgentTransactorSession struct {
	Contract     *UpgradeAgentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// UpgradeAgentRaw is an auto generated low-level Go binding around an Ethereum contract.
type UpgradeAgentRaw struct {
	Contract *UpgradeAgent // Generic contract binding to access the raw methods on
}

// UpgradeAgentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UpgradeAgentCallerRaw struct {
	Contract *UpgradeAgentCaller // Generic read-only contract binding to access the raw methods on
}

// UpgradeAgentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UpgradeAgentTransactorRaw struct {
	Contract *UpgradeAgentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpgradeAgent creates a new instance of UpgradeAgent, bound to a specific deployed contract.
func NewUpgradeAgent(address common.Address, backend bind.ContractBackend) (*UpgradeAgent, error) {
	contract, err := bindUpgradeAgent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UpgradeAgent{UpgradeAgentCaller: UpgradeAgentCaller{contract: contract}, UpgradeAgentTransactor: UpgradeAgentTransactor{contract: contract}, UpgradeAgentFilterer: UpgradeAgentFilterer{contract: contract}}, nil
}

// NewUpgradeAgentCaller creates a new read-only instance of UpgradeAgent, bound to a specific deployed contract.
func NewUpgradeAgentCaller(address common.Address, caller bind.ContractCaller) (*UpgradeAgentCaller, error) {
	contract, err := bindUpgradeAgent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeAgentCaller{contract: contract}, nil
}

// NewUpgradeAgentTransactor creates a new write-only instance of UpgradeAgent, bound to a specific deployed contract.
func NewUpgradeAgentTransactor(address common.Address, transactor bind.ContractTransactor) (*UpgradeAgentTransactor, error) {
	contract, err := bindUpgradeAgent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeAgentTransactor{contract: contract}, nil
}

// NewUpgradeAgentFilterer creates a new log filterer instance of UpgradeAgent, bound to a specific deployed contract.
func NewUpgradeAgentFilterer(address common.Address, filterer bind.ContractFilterer) (*UpgradeAgentFilterer, error) {
	contract, err := bindUpgradeAgent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpgradeAgentFilterer{contract: contract}, nil
}

// bindUpgradeAgent binds a generic wrapper to an already deployed contract.
func bindUpgradeAgent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UpgradeAgentMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeAgent *UpgradeAgentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeAgent.Contract.UpgradeAgentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeAgent *UpgradeAgentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeAgent.Contract.UpgradeAgentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeAgent *UpgradeAgentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeAgent.Contract.UpgradeAgentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeAgent *UpgradeAgentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeAgent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeAgent *UpgradeAgentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeAgent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeAgent *UpgradeAgentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeAgent.Contract.contract.Transact(opts, method, params...)
}

// IsUpgradeAgent is a free data retrieval call binding the contract method 0x61d3d7a6.
//
// Solidity: function isUpgradeAgent() view returns(bool)
func (_UpgradeAgent *UpgradeAgentCaller) IsUpgradeAgent(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _UpgradeAgent.contract.Call(opts, &out, "isUpgradeAgent")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUpgradeAgent is a free data retrieval call binding the contract method 0x61d3d7a6.
//
// Solidity: function isUpgradeAgent() view returns(bool)
func (_UpgradeAgent *UpgradeAgentSession) IsUpgradeAgent() (bool, error) {
	return _UpgradeAgent.Contract.IsUpgradeAgent(&_UpgradeAgent.CallOpts)
}

// IsUpgradeAgent is a free data retrieval call binding the contract method 0x61d3d7a6.
//
// Solidity: function isUpgradeAgent() view returns(bool)
func (_UpgradeAgent *UpgradeAgentCallerSession) IsUpgradeAgent() (bool, error) {
	return _UpgradeAgent.Contract.IsUpgradeAgent(&_UpgradeAgent.CallOpts)
}

// OriginalSupply is a free data retrieval call binding the contract method 0x4b2ba0dd.
//
// Solidity: function originalSupply() view returns(uint256)
func (_UpgradeAgent *UpgradeAgentCaller) OriginalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeAgent.contract.Call(opts, &out, "originalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OriginalSupply is a free data retrieval call binding the contract method 0x4b2ba0dd.
//
// Solidity: function originalSupply() view returns(uint256)
func (_UpgradeAgent *UpgradeAgentSession) OriginalSupply() (*big.Int, error) {
	return _UpgradeAgent.Contract.OriginalSupply(&_UpgradeAgent.CallOpts)
}

// OriginalSupply is a free data retrieval call binding the contract method 0x4b2ba0dd.
//
// Solidity: function originalSupply() view returns(uint256)
func (_UpgradeAgent *UpgradeAgentCallerSession) OriginalSupply() (*big.Int, error) {
	return _UpgradeAgent.Contract.OriginalSupply(&_UpgradeAgent.CallOpts)
}

// UpgradeFrom is a paid mutator transaction binding the contract method 0x753e88e5.
//
// Solidity: function upgradeFrom(address _from, uint256 _value) returns()
func (_UpgradeAgent *UpgradeAgentTransactor) UpgradeFrom(opts *bind.TransactOpts, _from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _UpgradeAgent.contract.Transact(opts, "upgradeFrom", _from, _value)
}

// UpgradeFrom is a paid mutator transaction binding the contract method 0x753e88e5.
//
// Solidity: function upgradeFrom(address _from, uint256 _value) returns()
func (_UpgradeAgent *UpgradeAgentSession) UpgradeFrom(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _UpgradeAgent.Contract.UpgradeFrom(&_UpgradeAgent.TransactOpts, _from, _value)
}

// UpgradeFrom is a paid mutator transaction binding the contract method 0x753e88e5.
//
// Solidity: function upgradeFrom(address _from, uint256 _value) returns()
func (_UpgradeAgent *UpgradeAgentTransactorSession) UpgradeFrom(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _UpgradeAgent.Contract.UpgradeFrom(&_UpgradeAgent.TransactOpts, _from, _value)
}

// UpgradeableTokenMetaData contains all meta data concerning the UpgradeableToken contract.
var UpgradeableTokenMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"upgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"upgradeAgent\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"upgradeMaster\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUpgradeState\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canUpgrade\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalUpgraded\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"setUpgradeAgent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"master\",\"type\":\"address\"}],\"name\":\"setUpgradeMaster\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_upgradeMaster\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Upgrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"UpgradeAgentSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]",
	Bin: "0x6060604052341561000f57600080fd5b604051602080610c6c8339810160405280805160038054600160a060020a03909216600160a060020a03199092169190911790555050610c18806100546000396000f3006060604052600436106100da5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b381146100df57806318160ddd1461012257806323b872dd1461014757806345977d031461017c5780635de4ccb014610194578063600440cb146101d057806370a08231146101e35780638444b3911461020f5780639738968c14610246578063a9059cbb14610259578063c752ff6214610288578063d7e7088a1461029b578063dd62ed3e146102c7578063eefa597b146102f9578063ffeb7d751461030c575b600080fd5b34156100ea57600080fd5b61010e73ffffffffffffffffffffffffffffffffffffffff60043516602435610338565b604051901515815260200160405180910390f35b341561012d57600080fd5b6101356103fa565b60405190815260200160405180910390f35b341561015257600080fd5b61010e73ffffffffffffffffffffffffffffffffffffffff60043581169060243516604435610400565b341561018757600080fd5b610192600435610535565b005b341561019f57600080fd5b6101a76106e6565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34156101db57600080fd5b6101a7610702565b34156101ee57600080fd5b61013573ffffffffffffffffffffffffffffffffffffffff6004351661071e565b341561021a57600080fd5b610222610746565b6040518082600481111561023257fe5b60ff16815260200191505060405180910390f35b341561025157600080fd5b61010e61079d565b341561026457600080fd5b61010e73ffffffffffffffffffffffffffffffffffffffff600435166024356107a2565b341561029357600080fd5b61013561088f565b34156102a657600080fd5b61019273ffffffffffffffffffffffffffffffffffffffff60043516610895565b34156102d257600080fd5b61013573ffffffffffffffffffffffffffffffffffffffff60043581169060243516610ad7565b341561030457600080fd5b61010e610b0f565b341561031757600080fd5b61019273ffffffffffffffffffffffffffffffffffffffff60043516610b14565b6000811580159061037a575073ffffffffffffffffffffffffffffffffffffffff33811660009081526002602090815260408083209387168352929052205415155b1561038457600080fd5b73ffffffffffffffffffffffffffffffffffffffff338116600081815260026020908152604080832094881680845294909152908190208590557f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a350600192915050565b60005481565b73ffffffffffffffffffffffffffffffffffffffff808416600090815260026020908152604080832033851684528252808320549386168352600190915281205490919061044e9084610ba5565b73ffffffffffffffffffffffffffffffffffffffff808616600090815260016020526040808220939093559087168152205461048a9084610bc9565b73ffffffffffffffffffffffffffffffffffffffff86166000908152600160205260409020556104ba8184610bc9565b73ffffffffffffffffffffffffffffffffffffffff808716600081815260026020908152604080832033861684529091529081902093909355908616917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9086905190815260200160405180910390a3506001949350505050565b600061053f610746565b9050600381600481111561054f57fe5b14806105665750600481600481111561056457fe5b145b151561057157600080fd5b81151561057d57600080fd5b73ffffffffffffffffffffffffffffffffffffffff33166000908152600160205260409020546105ad9083610bc9565b73ffffffffffffffffffffffffffffffffffffffff3316600090815260016020526040812091909155546105e19083610bc9565b6000556005546105f19083610ba5565b60055560045473ffffffffffffffffffffffffffffffffffffffff1663753e88e533846040517c010000000000000000000000000000000000000000000000000000000063ffffffff851602815273ffffffffffffffffffffffffffffffffffffffff90921660048301526024820152604401600060405180830381600087803b151561067d57600080fd5b6102c65a03f1151561068e57600080fd5b505060045473ffffffffffffffffffffffffffffffffffffffff908116915033167f7e5c344a8141a805725cb476f76c6953b842222b967edd1f78ddb6e8b3f397ac8460405190815260200160405180910390a35050565b60045473ffffffffffffffffffffffffffffffffffffffff1681565b60035473ffffffffffffffffffffffffffffffffffffffff1681565b73ffffffffffffffffffffffffffffffffffffffff1660009081526001602052604090205490565b600061075061079d565b151561075e5750600161079a565b60045473ffffffffffffffffffffffffffffffffffffffff1615156107855750600261079a565b60055415156107965750600361079a565b5060045b90565b600190565b6000604060443610156107b457600080fd5b73ffffffffffffffffffffffffffffffffffffffff33166000908152600160205260409020546107e49084610bc9565b73ffffffffffffffffffffffffffffffffffffffff33811660009081526001602052604080822093909355908616815220546108209084610ba5565b73ffffffffffffffffffffffffffffffffffffffff80861660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9086905190815260200160405180910390a35060019392505050565b60055481565b61089d61079d565b15156108a857600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615156108ca57600080fd5b6003543373ffffffffffffffffffffffffffffffffffffffff9081169116146108f257600080fd5b60046108fc610746565b600481111561090757fe5b141561091257600080fd5b600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169190911791829055166361d3d7a66000604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15156109ae57600080fd5b6102c65a03f115156109bf57600080fd5b5050506040518051905015156109d457600080fd5b60008054600454909173ffffffffffffffffffffffffffffffffffffffff90911690634b2ba0dd90604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1515610a4a57600080fd5b6102c65a03f11515610a5b57600080fd5b50505060405180519050141515610a7157600080fd5b6004547f7845d5aa74cc410e35571258d954f23b82276e160fe8c188fa80566580f279cc9073ffffffffffffffffffffffffffffffffffffffff1660405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a150565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260026020908152604080832093909416825291909152205490565b600181565b73ffffffffffffffffffffffffffffffffffffffff81161515610b3657600080fd5b6003543373ffffffffffffffffffffffffffffffffffffffff908116911614610b5e57600080fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6000828201610bc2848210801590610bbd5750838210155b610bdd565b9392505050565b6000610bd783831115610bdd565b50900390565b801515610be957600080fd5b505600a165627a7a723058206b155a02e3458e6fa0189cd002d11ca22eabee5afce7f0e07680c7413808cab60029",
}

// UpgradeableTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use UpgradeableTokenMetaData.ABI instead.
var UpgradeableTokenABI = UpgradeableTokenMetaData.ABI

// UpgradeableTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UpgradeableTokenMetaData.Bin instead.
var UpgradeableTokenBin = UpgradeableTokenMetaData.Bin

// DeployUpgradeableToken deploys a new Ethereum contract, binding an instance of UpgradeableToken to it.
func DeployUpgradeableToken(auth *bind.TransactOpts, backend bind.ContractBackend, _upgradeMaster common.Address) (common.Address, *types.Transaction, *UpgradeableToken, error) {
	parsed, err := UpgradeableTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UpgradeableTokenBin), backend, _upgradeMaster)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UpgradeableToken{UpgradeableTokenCaller: UpgradeableTokenCaller{contract: contract}, UpgradeableTokenTransactor: UpgradeableTokenTransactor{contract: contract}, UpgradeableTokenFilterer: UpgradeableTokenFilterer{contract: contract}}, nil
}

// UpgradeableToken is an auto generated Go binding around an Ethereum contract.
type UpgradeableToken struct {
	UpgradeableTokenCaller     // Read-only binding to the contract
	UpgradeableTokenTransactor // Write-only binding to the contract
	UpgradeableTokenFilterer   // Log filterer for contract events
}

// UpgradeableTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type UpgradeableTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UpgradeableTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UpgradeableTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeableTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UpgradeableTokenSession struct {
	Contract     *UpgradeableToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UpgradeableTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UpgradeableTokenCallerSession struct {
	Contract *UpgradeableTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// UpgradeableTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UpgradeableTokenTransactorSession struct {
	Contract     *UpgradeableTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// UpgradeableTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type UpgradeableTokenRaw struct {
	Contract *UpgradeableToken // Generic contract binding to access the raw methods on
}

// UpgradeableTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UpgradeableTokenCallerRaw struct {
	Contract *UpgradeableTokenCaller // Generic read-only contract binding to access the raw methods on
}

// UpgradeableTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UpgradeableTokenTransactorRaw struct {
	Contract *UpgradeableTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpgradeableToken creates a new instance of UpgradeableToken, bound to a specific deployed contract.
func NewUpgradeableToken(address common.Address, backend bind.ContractBackend) (*UpgradeableToken, error) {
	contract, err := bindUpgradeableToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UpgradeableToken{UpgradeableTokenCaller: UpgradeableTokenCaller{contract: contract}, UpgradeableTokenTransactor: UpgradeableTokenTransactor{contract: contract}, UpgradeableTokenFilterer: UpgradeableTokenFilterer{contract: contract}}, nil
}

// NewUpgradeableTokenCaller creates a new read-only instance of UpgradeableToken, bound to a specific deployed contract.
func NewUpgradeableTokenCaller(address common.Address, caller bind.ContractCaller) (*UpgradeableTokenCaller, error) {
	contract, err := bindUpgradeableToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeableTokenCaller{contract: contract}, nil
}

// NewUpgradeableTokenTransactor creates a new write-only instance of UpgradeableToken, bound to a specific deployed contract.
func NewUpgradeableTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*UpgradeableTokenTransactor, error) {
	contract, err := bindUpgradeableToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeableTokenTransactor{contract: contract}, nil
}

// NewUpgradeableTokenFilterer creates a new log filterer instance of UpgradeableToken, bound to a specific deployed contract.
func NewUpgradeableTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*UpgradeableTokenFilterer, error) {
	contract, err := bindUpgradeableToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpgradeableTokenFilterer{contract: contract}, nil
}

// bindUpgradeableToken binds a generic wrapper to an already deployed contract.
func bindUpgradeableToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UpgradeableTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeableToken *UpgradeableTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeableToken.Contract.UpgradeableTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeableToken *UpgradeableTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeableToken.Contract.UpgradeableTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeableToken *UpgradeableTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeableToken.Contract.UpgradeableTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeableToken *UpgradeableTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeableToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeableToken *UpgradeableTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeableToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeableToken *UpgradeableTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeableToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_UpgradeableToken *UpgradeableTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeableToken.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_UpgradeableToken *UpgradeableTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _UpgradeableToken.Contract.Allowance(&_UpgradeableToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_UpgradeableToken *UpgradeableTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _UpgradeableToken.Contract.Allowance(&_UpgradeableToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_UpgradeableToken *UpgradeableTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeableToken.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_UpgradeableToken *UpgradeableTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _UpgradeableToken.Contract.BalanceOf(&_UpgradeableToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_UpgradeableToken *UpgradeableTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _UpgradeableToken.Contract.BalanceOf(&_UpgradeableToken.CallOpts, _owner)
}

// CanUpgrade is a free data retrieval call binding the contract method 0x9738968c.
//
// Solidity: function canUpgrade() view returns(bool)
func (_UpgradeableToken *UpgradeableTokenCaller) CanUpgrade(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _UpgradeableToken.contract.Call(opts, &out, "canUpgrade")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanUpgrade is a free data retrieval call binding the contract method 0x9738968c.
//
// Solidity: function canUpgrade() view returns(bool)
func (_UpgradeableToken *UpgradeableTokenSession) CanUpgrade() (bool, error) {
	return _UpgradeableToken.Contract.CanUpgrade(&_UpgradeableToken.CallOpts)
}

// CanUpgrade is a free data retrieval call binding the contract method 0x9738968c.
//
// Solidity: function canUpgrade() view returns(bool)
func (_UpgradeableToken *UpgradeableTokenCallerSession) CanUpgrade() (bool, error) {
	return _UpgradeableToken.Contract.CanUpgrade(&_UpgradeableToken.CallOpts)
}

// GetUpgradeState is a free data retrieval call binding the contract method 0x8444b391.
//
// Solidity: function getUpgradeState() view returns(uint8)
func (_UpgradeableToken *UpgradeableTokenCaller) GetUpgradeState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _UpgradeableToken.contract.Call(opts, &out, "getUpgradeState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetUpgradeState is a free data retrieval call binding the contract method 0x8444b391.
//
// Solidity: function getUpgradeState() view returns(uint8)
func (_UpgradeableToken *UpgradeableTokenSession) GetUpgradeState() (uint8, error) {
	return _UpgradeableToken.Contract.GetUpgradeState(&_UpgradeableToken.CallOpts)
}

// GetUpgradeState is a free data retrieval call binding the contract method 0x8444b391.
//
// Solidity: function getUpgradeState() view returns(uint8)
func (_UpgradeableToken *UpgradeableTokenCallerSession) GetUpgradeState() (uint8, error) {
	return _UpgradeableToken.Contract.GetUpgradeState(&_UpgradeableToken.CallOpts)
}

// IsToken is a free data retrieval call binding the contract method 0xeefa597b.
//
// Solidity: function isToken() view returns(bool)
func (_UpgradeableToken *UpgradeableTokenCaller) IsToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _UpgradeableToken.contract.Call(opts, &out, "isToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsToken is a free data retrieval call binding the contract method 0xeefa597b.
//
// Solidity: function isToken() view returns(bool)
func (_UpgradeableToken *UpgradeableTokenSession) IsToken() (bool, error) {
	return _UpgradeableToken.Contract.IsToken(&_UpgradeableToken.CallOpts)
}

// IsToken is a free data retrieval call binding the contract method 0xeefa597b.
//
// Solidity: function isToken() view returns(bool)
func (_UpgradeableToken *UpgradeableTokenCallerSession) IsToken() (bool, error) {
	return _UpgradeableToken.Contract.IsToken(&_UpgradeableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UpgradeableToken *UpgradeableTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeableToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UpgradeableToken *UpgradeableTokenSession) TotalSupply() (*big.Int, error) {
	return _UpgradeableToken.Contract.TotalSupply(&_UpgradeableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UpgradeableToken *UpgradeableTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _UpgradeableToken.Contract.TotalSupply(&_UpgradeableToken.CallOpts)
}

// TotalUpgraded is a free data retrieval call binding the contract method 0xc752ff62.
//
// Solidity: function totalUpgraded() view returns(uint256)
func (_UpgradeableToken *UpgradeableTokenCaller) TotalUpgraded(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeableToken.contract.Call(opts, &out, "totalUpgraded")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalUpgraded is a free data retrieval call binding the contract method 0xc752ff62.
//
// Solidity: function totalUpgraded() view returns(uint256)
func (_UpgradeableToken *UpgradeableTokenSession) TotalUpgraded() (*big.Int, error) {
	return _UpgradeableToken.Contract.TotalUpgraded(&_UpgradeableToken.CallOpts)
}

// TotalUpgraded is a free data retrieval call binding the contract method 0xc752ff62.
//
// Solidity: function totalUpgraded() view returns(uint256)
func (_UpgradeableToken *UpgradeableTokenCallerSession) TotalUpgraded() (*big.Int, error) {
	return _UpgradeableToken.Contract.TotalUpgraded(&_UpgradeableToken.CallOpts)
}

// UpgradeAgent is a free data retrieval call binding the contract method 0x5de4ccb0.
//
// Solidity: function upgradeAgent() view returns(address)
func (_UpgradeableToken *UpgradeableTokenCaller) UpgradeAgent(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UpgradeableToken.contract.Call(opts, &out, "upgradeAgent")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpgradeAgent is a free data retrieval call binding the contract method 0x5de4ccb0.
//
// Solidity: function upgradeAgent() view returns(address)
func (_UpgradeableToken *UpgradeableTokenSession) UpgradeAgent() (common.Address, error) {
	return _UpgradeableToken.Contract.UpgradeAgent(&_UpgradeableToken.CallOpts)
}

// UpgradeAgent is a free data retrieval call binding the contract method 0x5de4ccb0.
//
// Solidity: function upgradeAgent() view returns(address)
func (_UpgradeableToken *UpgradeableTokenCallerSession) UpgradeAgent() (common.Address, error) {
	return _UpgradeableToken.Contract.UpgradeAgent(&_UpgradeableToken.CallOpts)
}

// UpgradeMaster is a free data retrieval call binding the contract method 0x600440cb.
//
// Solidity: function upgradeMaster() view returns(address)
func (_UpgradeableToken *UpgradeableTokenCaller) UpgradeMaster(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UpgradeableToken.contract.Call(opts, &out, "upgradeMaster")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpgradeMaster is a free data retrieval call binding the contract method 0x600440cb.
//
// Solidity: function upgradeMaster() view returns(address)
func (_UpgradeableToken *UpgradeableTokenSession) UpgradeMaster() (common.Address, error) {
	return _UpgradeableToken.Contract.UpgradeMaster(&_UpgradeableToken.CallOpts)
}

// UpgradeMaster is a free data retrieval call binding the contract method 0x600440cb.
//
// Solidity: function upgradeMaster() view returns(address)
func (_UpgradeableToken *UpgradeableTokenCallerSession) UpgradeMaster() (common.Address, error) {
	return _UpgradeableToken.Contract.UpgradeMaster(&_UpgradeableToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_UpgradeableToken *UpgradeableTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _UpgradeableToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_UpgradeableToken *UpgradeableTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _UpgradeableToken.Contract.Approve(&_UpgradeableToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_UpgradeableToken *UpgradeableTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _UpgradeableToken.Contract.Approve(&_UpgradeableToken.TransactOpts, _spender, _value)
}

// SetUpgradeAgent is a paid mutator transaction binding the contract method 0xd7e7088a.
//
// Solidity: function setUpgradeAgent(address agent) returns()
func (_UpgradeableToken *UpgradeableTokenTransactor) SetUpgradeAgent(opts *bind.TransactOpts, agent common.Address) (*types.Transaction, error) {
	return _UpgradeableToken.contract.Transact(opts, "setUpgradeAgent", agent)
}

// SetUpgradeAgent is a paid mutator transaction binding the contract method 0xd7e7088a.
//
// Solidity: function setUpgradeAgent(address agent) returns()
func (_UpgradeableToken *UpgradeableTokenSession) SetUpgradeAgent(agent common.Address) (*types.Transaction, error) {
	return _UpgradeableToken.Contract.SetUpgradeAgent(&_UpgradeableToken.TransactOpts, agent)
}

// SetUpgradeAgent is a paid mutator transaction binding the contract method 0xd7e7088a.
//
// Solidity: function setUpgradeAgent(address agent) returns()
func (_UpgradeableToken *UpgradeableTokenTransactorSession) SetUpgradeAgent(agent common.Address) (*types.Transaction, error) {
	return _UpgradeableToken.Contract.SetUpgradeAgent(&_UpgradeableToken.TransactOpts, agent)
}

// SetUpgradeMaster is a paid mutator transaction binding the contract method 0xffeb7d75.
//
// Solidity: function setUpgradeMaster(address master) returns()
func (_UpgradeableToken *UpgradeableTokenTransactor) SetUpgradeMaster(opts *bind.TransactOpts, master common.Address) (*types.Transaction, error) {
	return _UpgradeableToken.contract.Transact(opts, "setUpgradeMaster", master)
}

// SetUpgradeMaster is a paid mutator transaction binding the contract method 0xffeb7d75.
//
// Solidity: function setUpgradeMaster(address master) returns()
func (_UpgradeableToken *UpgradeableTokenSession) SetUpgradeMaster(master common.Address) (*types.Transaction, error) {
	return _UpgradeableToken.Contract.SetUpgradeMaster(&_UpgradeableToken.TransactOpts, master)
}

// SetUpgradeMaster is a paid mutator transaction binding the contract method 0xffeb7d75.
//
// Solidity: function setUpgradeMaster(address master) returns()
func (_UpgradeableToken *UpgradeableTokenTransactorSession) SetUpgradeMaster(master common.Address) (*types.Transaction, error) {
	return _UpgradeableToken.Contract.SetUpgradeMaster(&_UpgradeableToken.TransactOpts, master)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_UpgradeableToken *UpgradeableTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _UpgradeableToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_UpgradeableToken *UpgradeableTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _UpgradeableToken.Contract.Transfer(&_UpgradeableToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_UpgradeableToken *UpgradeableTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _UpgradeableToken.Contract.Transfer(&_UpgradeableToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_UpgradeableToken *UpgradeableTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _UpgradeableToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_UpgradeableToken *UpgradeableTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _UpgradeableToken.Contract.TransferFrom(&_UpgradeableToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_UpgradeableToken *UpgradeableTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _UpgradeableToken.Contract.TransferFrom(&_UpgradeableToken.TransactOpts, _from, _to, _value)
}

// Upgrade is a paid mutator transaction binding the contract method 0x45977d03.
//
// Solidity: function upgrade(uint256 value) returns()
func (_UpgradeableToken *UpgradeableTokenTransactor) Upgrade(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _UpgradeableToken.contract.Transact(opts, "upgrade", value)
}

// Upgrade is a paid mutator transaction binding the contract method 0x45977d03.
//
// Solidity: function upgrade(uint256 value) returns()
func (_UpgradeableToken *UpgradeableTokenSession) Upgrade(value *big.Int) (*types.Transaction, error) {
	return _UpgradeableToken.Contract.Upgrade(&_UpgradeableToken.TransactOpts, value)
}

// Upgrade is a paid mutator transaction binding the contract method 0x45977d03.
//
// Solidity: function upgrade(uint256 value) returns()
func (_UpgradeableToken *UpgradeableTokenTransactorSession) Upgrade(value *big.Int) (*types.Transaction, error) {
	return _UpgradeableToken.Contract.Upgrade(&_UpgradeableToken.TransactOpts, value)
}

// UpgradeableTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the UpgradeableToken contract.
type UpgradeableTokenApprovalIterator struct {
	Event *UpgradeableTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UpgradeableTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UpgradeableTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UpgradeableTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableTokenApproval represents a Approval event raised by the UpgradeableToken contract.
type UpgradeableTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UpgradeableToken *UpgradeableTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*UpgradeableTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _UpgradeableToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeableTokenApprovalIterator{contract: _UpgradeableToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UpgradeableToken *UpgradeableTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *UpgradeableTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _UpgradeableToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableTokenApproval)
				if err := _UpgradeableToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UpgradeableToken *UpgradeableTokenFilterer) ParseApproval(log types.Log) (*UpgradeableTokenApproval, error) {
	event := new(UpgradeableTokenApproval)
	if err := _UpgradeableToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeableTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the UpgradeableToken contract.
type UpgradeableTokenTransferIterator struct {
	Event *UpgradeableTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UpgradeableTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UpgradeableTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UpgradeableTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableTokenTransfer represents a Transfer event raised by the UpgradeableToken contract.
type UpgradeableTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UpgradeableToken *UpgradeableTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*UpgradeableTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UpgradeableToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeableTokenTransferIterator{contract: _UpgradeableToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UpgradeableToken *UpgradeableTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *UpgradeableTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UpgradeableToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableTokenTransfer)
				if err := _UpgradeableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UpgradeableToken *UpgradeableTokenFilterer) ParseTransfer(log types.Log) (*UpgradeableTokenTransfer, error) {
	event := new(UpgradeableTokenTransfer)
	if err := _UpgradeableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeableTokenUpgradeIterator is returned from FilterUpgrade and is used to iterate over the raw logs and unpacked data for Upgrade events raised by the UpgradeableToken contract.
type UpgradeableTokenUpgradeIterator struct {
	Event *UpgradeableTokenUpgrade // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UpgradeableTokenUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableTokenUpgrade)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UpgradeableTokenUpgrade)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UpgradeableTokenUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableTokenUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableTokenUpgrade represents a Upgrade event raised by the UpgradeableToken contract.
type UpgradeableTokenUpgrade struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpgrade is a free log retrieval operation binding the contract event 0x7e5c344a8141a805725cb476f76c6953b842222b967edd1f78ddb6e8b3f397ac.
//
// Solidity: event Upgrade(address indexed _from, address indexed _to, uint256 _value)
func (_UpgradeableToken *UpgradeableTokenFilterer) FilterUpgrade(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*UpgradeableTokenUpgradeIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _UpgradeableToken.contract.FilterLogs(opts, "Upgrade", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeableTokenUpgradeIterator{contract: _UpgradeableToken.contract, event: "Upgrade", logs: logs, sub: sub}, nil
}

// WatchUpgrade is a free log subscription operation binding the contract event 0x7e5c344a8141a805725cb476f76c6953b842222b967edd1f78ddb6e8b3f397ac.
//
// Solidity: event Upgrade(address indexed _from, address indexed _to, uint256 _value)
func (_UpgradeableToken *UpgradeableTokenFilterer) WatchUpgrade(opts *bind.WatchOpts, sink chan<- *UpgradeableTokenUpgrade, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _UpgradeableToken.contract.WatchLogs(opts, "Upgrade", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableTokenUpgrade)
				if err := _UpgradeableToken.contract.UnpackLog(event, "Upgrade", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgrade is a log parse operation binding the contract event 0x7e5c344a8141a805725cb476f76c6953b842222b967edd1f78ddb6e8b3f397ac.
//
// Solidity: event Upgrade(address indexed _from, address indexed _to, uint256 _value)
func (_UpgradeableToken *UpgradeableTokenFilterer) ParseUpgrade(log types.Log) (*UpgradeableTokenUpgrade, error) {
	event := new(UpgradeableTokenUpgrade)
	if err := _UpgradeableToken.contract.UnpackLog(event, "Upgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeableTokenUpgradeAgentSetIterator is returned from FilterUpgradeAgentSet and is used to iterate over the raw logs and unpacked data for UpgradeAgentSet events raised by the UpgradeableToken contract.
type UpgradeableTokenUpgradeAgentSetIterator struct {
	Event *UpgradeableTokenUpgradeAgentSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UpgradeableTokenUpgradeAgentSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeableTokenUpgradeAgentSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UpgradeableTokenUpgradeAgentSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UpgradeableTokenUpgradeAgentSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeableTokenUpgradeAgentSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeableTokenUpgradeAgentSet represents a UpgradeAgentSet event raised by the UpgradeableToken contract.
type UpgradeableTokenUpgradeAgentSet struct {
	Agent common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpgradeAgentSet is a free log retrieval operation binding the contract event 0x7845d5aa74cc410e35571258d954f23b82276e160fe8c188fa80566580f279cc.
//
// Solidity: event UpgradeAgentSet(address agent)
func (_UpgradeableToken *UpgradeableTokenFilterer) FilterUpgradeAgentSet(opts *bind.FilterOpts) (*UpgradeableTokenUpgradeAgentSetIterator, error) {

	logs, sub, err := _UpgradeableToken.contract.FilterLogs(opts, "UpgradeAgentSet")
	if err != nil {
		return nil, err
	}
	return &UpgradeableTokenUpgradeAgentSetIterator{contract: _UpgradeableToken.contract, event: "UpgradeAgentSet", logs: logs, sub: sub}, nil
}

// WatchUpgradeAgentSet is a free log subscription operation binding the contract event 0x7845d5aa74cc410e35571258d954f23b82276e160fe8c188fa80566580f279cc.
//
// Solidity: event UpgradeAgentSet(address agent)
func (_UpgradeableToken *UpgradeableTokenFilterer) WatchUpgradeAgentSet(opts *bind.WatchOpts, sink chan<- *UpgradeableTokenUpgradeAgentSet) (event.Subscription, error) {

	logs, sub, err := _UpgradeableToken.contract.WatchLogs(opts, "UpgradeAgentSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeableTokenUpgradeAgentSet)
				if err := _UpgradeableToken.contract.UnpackLog(event, "UpgradeAgentSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgradeAgentSet is a log parse operation binding the contract event 0x7845d5aa74cc410e35571258d954f23b82276e160fe8c188fa80566580f279cc.
//
// Solidity: event UpgradeAgentSet(address agent)
func (_UpgradeableToken *UpgradeableTokenFilterer) ParseUpgradeAgentSet(log types.Log) (*UpgradeableTokenUpgradeAgentSet, error) {
	event := new(UpgradeableTokenUpgradeAgentSet)
	if err := _UpgradeableToken.contract.UnpackLog(event, "UpgradeAgentSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

