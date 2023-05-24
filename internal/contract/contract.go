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
	Bin: "0x608060405234801561001057600080fd5b506106d7806100206000396000f3006080604052600436106100985763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461009d57806318160ddd146100e257806323b872dd1461010957806342966c681461014057806370a082311461015a578063a9059cbb14610188578063dd62ed3e146101b9578063eefa597b146101ed578063fccc281314610202575b600080fd5b3480156100a957600080fd5b506100ce73ffffffffffffffffffffffffffffffffffffffff60043516602435610240565b604080519115158252519081900360200190f35b3480156100ee57600080fd5b506100f76102fe565b60408051918252519081900360200190f35b34801561011557600080fd5b506100ce73ffffffffffffffffffffffffffffffffffffffff60043581169060243516604435610304565b34801561014c57600080fd5b50610158600435610435565b005b34801561016657600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff60043516610529565b34801561019457600080fd5b506100ce73ffffffffffffffffffffffffffffffffffffffff60043516602435610551565b3480156101c557600080fd5b506100f773ffffffffffffffffffffffffffffffffffffffff60043581169060243516610622565b3480156101f957600080fd5b506100ce61065a565b34801561020e57600080fd5b5061021761065f565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60008115801590610280575033600090815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915290205415155b1561028a57600080fd5b33600081815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a350600192915050565b60005481565b73ffffffffffffffffffffffffffffffffffffffff8084166000908152600260209081526040808320338452825280832054938616835260019091528120549091906103509084610664565b73ffffffffffffffffffffffffffffffffffffffff808616600090815260016020526040808220939093559087168152205461038c9084610688565b73ffffffffffffffffffffffffffffffffffffffff86166000908152600160205260409020556103bc8184610688565b73ffffffffffffffffffffffffffffffffffffffff808716600081815260026020908152604080832033845282529182902094909455805187815290519288169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a3506001949350505050565b3360008181526001602052604090205461044f9083610688565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260016020526040812091909155546104839083610688565b6000556040805173ffffffffffffffffffffffffffffffffffffffff831681526020810184905281517f696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7929181900390910190a160408051838152905160009173ffffffffffffffffffffffffffffffffffffffff8416917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b73ffffffffffffffffffffffffffffffffffffffff1660009081526001602052604090205490565b60006040604436101561056357600080fd5b3360009081526001602052604090205461057d9084610688565b336000908152600160205260408082209290925573ffffffffffffffffffffffffffffffffffffffff8616815220546105b69084610664565b73ffffffffffffffffffffffffffffffffffffffff85166000818152600160209081526040918290209390935580518681529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35060019392505050565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260026020908152604080832093909416825291909152205490565b600181565b600081565b600082820161068184821080159061067c5750838210155b61069c565b9392505050565b60006106968383111561069c565b50900390565b8015156106a857600080fd5b505600a165627a7a72305820d7ded30455eb39c078bd8e72ef48946ba7cc2a999892c03eaa3118bfbfeb81ce0029",
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
	Bin: "0x608060405234801561001057600080fd5b506040516200108c3803806200108c8339810160409081528151602080840151928401516060850151608086015160038054600160a060020a031916600160a060020a0387161790559486018051949690959201939092610076916006918701906100b9565b50825161008a9060079060208601906100b9565b506000828155600891909155600160a060020a0390941684526001602052604090932092909255506101549050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100fa57805160ff1916838001178555610127565b82800160010185558215610127579182015b8281111561012757825182559160200191906001019061010c565b50610133929150610137565b5090565b61015191905b80821115610133576000815560010161013d565b90565b610f2880620001646000396000f3006080604052600436106101115763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde038114610116578063095ea7b3146101a057806318160ddd146101e557806323b872dd1461020c578063313ce5671461024357806342966c681461025857806345977d03146102725780635de4ccb01461028a578063600440cb146102c857806370a08231146102dd5780638444b3911461030b57806395d89b41146103445780639738968c14610359578063a9059cbb1461036e578063c752ff621461039f578063d7e7088a146103b4578063dd62ed3e146103e2578063eefa597b14610416578063fccc28131461042b578063ffeb7d7514610440575b600080fd5b34801561012257600080fd5b5061012b61046e565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561016557818101518382015260200161014d565b50505050905090810190601f1680156101925780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101ac57600080fd5b506101d173ffffffffffffffffffffffffffffffffffffffff6004351660243561051a565b604080519115158252519081900360200190f35b3480156101f157600080fd5b506101fa6105d8565b60408051918252519081900360200190f35b34801561021857600080fd5b506101d173ffffffffffffffffffffffffffffffffffffffff600435811690602435166044356105de565b34801561024f57600080fd5b506101fa61070f565b34801561026457600080fd5b50610270600435610715565b005b34801561027e57600080fd5b50610270600435610809565b34801561029657600080fd5b5061029f610984565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156102d457600080fd5b5061029f6109a0565b3480156102e957600080fd5b506101fa73ffffffffffffffffffffffffffffffffffffffff600435166109bc565b34801561031757600080fd5b506103206109e4565b6040518082600481111561033057fe5b60ff16815260200191505060405180910390f35b34801561035057600080fd5b5061012b610a3b565b34801561036557600080fd5b506101d1610ab4565b34801561037a57600080fd5b506101d173ffffffffffffffffffffffffffffffffffffffff60043516602435610ab9565b3480156103ab57600080fd5b506101fa610b8a565b3480156103c057600080fd5b5061027073ffffffffffffffffffffffffffffffffffffffff60043516610b90565b3480156103ee57600080fd5b506101fa73ffffffffffffffffffffffffffffffffffffffff60043581169060243516610de6565b34801561042257600080fd5b506101d1610e1e565b34801561043757600080fd5b5061029f610e23565b34801561044c57600080fd5b5061027073ffffffffffffffffffffffffffffffffffffffff60043516610e28565b6006805460408051602060026001851615610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190941693909304601f810184900484028201840190925281815292918301828280156105125780601f106104e757610100808354040283529160200191610512565b820191906000526020600020905b8154815290600101906020018083116104f557829003601f168201915b505050505081565b6000811580159061055a575033600090815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915290205415155b1561056457600080fd5b33600081815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a350600192915050565b60005481565b73ffffffffffffffffffffffffffffffffffffffff80841660009081526002602090815260408083203384528252808320549386168352600190915281205490919061062a9084610eb5565b73ffffffffffffffffffffffffffffffffffffffff80861660009081526001602052604080822093909355908716815220546106669084610ed9565b73ffffffffffffffffffffffffffffffffffffffff86166000908152600160205260409020556106968184610ed9565b73ffffffffffffffffffffffffffffffffffffffff808716600081815260026020908152604080832033845282529182902094909455805187815290519288169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a3506001949350505050565b60085481565b3360008181526001602052604090205461072f9083610ed9565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260016020526040812091909155546107639083610ed9565b6000556040805173ffffffffffffffffffffffffffffffffffffffff831681526020810184905281517f696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7929181900390910190a160408051838152905160009173ffffffffffffffffffffffffffffffffffffffff8416917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b60006108136109e4565b9050600381600481111561082357fe5b148061083a5750600481600481111561083857fe5b145b151561084557600080fd5b81151561085157600080fd5b3360009081526001602052604090205461086b9083610ed9565b33600090815260016020526040812091909155546108899083610ed9565b6000556005546108999083610eb5565b60055560048054604080517f753e88e50000000000000000000000000000000000000000000000000000000081523393810193909352602483018590525173ffffffffffffffffffffffffffffffffffffffff9091169163753e88e591604480830192600092919082900301818387803b15801561091657600080fd5b505af115801561092a573d6000803e3d6000fd5b505060045460408051868152905173ffffffffffffffffffffffffffffffffffffffff90921693503392507f7e5c344a8141a805725cb476f76c6953b842222b967edd1f78ddb6e8b3f397ac919081900360200190a35050565b60045473ffffffffffffffffffffffffffffffffffffffff1681565b60035473ffffffffffffffffffffffffffffffffffffffff1681565b73ffffffffffffffffffffffffffffffffffffffff1660009081526001602052604090205490565b60006109ee610ab4565b15156109fc57506001610a38565b60045473ffffffffffffffffffffffffffffffffffffffff161515610a2357506002610a38565b6005541515610a3457506003610a38565b5060045b90565b6007805460408051602060026001851615610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190941693909304601f810184900484028201840190925281815292918301828280156105125780601f106104e757610100808354040283529160200191610512565b600190565b600060406044361015610acb57600080fd5b33600090815260016020526040902054610ae59084610ed9565b336000908152600160205260408082209290925573ffffffffffffffffffffffffffffffffffffffff861681522054610b1e9084610eb5565b73ffffffffffffffffffffffffffffffffffffffff85166000818152600160209081526040918290209390935580518681529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35060019392505050565b60055481565b610b98610ab4565b1515610ba357600080fd5b73ffffffffffffffffffffffffffffffffffffffff81161515610bc557600080fd5b60035473ffffffffffffffffffffffffffffffffffffffff163314610be957600080fd5b6004610bf36109e4565b6004811115610bfe57fe5b1415610c0957600080fd5b600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691909117808355604080517f61d3d7a6000000000000000000000000000000000000000000000000000000008152905191909216926361d3d7a6928082019260209290918290030181600087803b158015610ca157600080fd5b505af1158015610cb5573d6000803e3d6000fd5b505050506040513d6020811015610ccb57600080fd5b50511515610cd857600080fd5b600054600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634b2ba0dd6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015610d6157600080fd5b505af1158015610d75573d6000803e3d6000fd5b505050506040513d6020811015610d8b57600080fd5b505114610d9757600080fd5b6004546040805173ffffffffffffffffffffffffffffffffffffffff9092168252517f7845d5aa74cc410e35571258d954f23b82276e160fe8c188fa80566580f279cc9181900360200190a150565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260026020908152604080832093909416825291909152205490565b600181565b600081565b73ffffffffffffffffffffffffffffffffffffffff81161515610e4a57600080fd5b60035473ffffffffffffffffffffffffffffffffffffffff163314610e6e57600080fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6000828201610ed2848210801590610ecd5750838210155b610eed565b9392505050565b6000610ee783831115610eed565b50900390565b801515610ef957600080fd5b505600a165627a7a72305820f7e5e0e27573efce314578199c439666273bfcc81de4226ac7488f97ad4eb0e40029",
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
	Bin: "0x608060405234801561001057600080fd5b506101e4806100206000396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632f4f3316811461005b57806357183c821461009d5780638f838478146100f7575b600080fd5b34801561006757600080fd5b5061008973ffffffffffffffffffffffffffffffffffffffff60043516610137565b604080519115158252519081900360200190f35b3480156100a957600080fd5b506100ce73ffffffffffffffffffffffffffffffffffffffff6004351660243561014c565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561010357600080fd5b5061012573ffffffffffffffffffffffffffffffffffffffff60043516610190565b60408051918252519081900360200190f35b60006020819052908152604090205460ff1681565b60016020528160005260406000208181548110151561016757fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff169150829050565b73ffffffffffffffffffffffffffffffffffffffff16600090815260016020526040902054905600a165627a7a7230582033c102c21197bb607dc30da38e8d3ab7d0917072f3ca7cff9cbd34f6f92963d50029",
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
	Bin: "0x60806040523480156200001157600080fd5b5060405162001a5338038062001a5383398101604052805160208201519101805190919060009082603282118015906200004b5750818111155b80156200005757508015155b80156200006357508115155b15156200006f57600080fd5b600092505b845183101562000147576002600086858151811015156200009157fe5b6020908102909101810151600160a060020a031682528101919091526040016000205460ff16158015620000e757508483815181101515620000cf57fe5b90602001906020020151600160a060020a0316600014155b1515620000f357600080fd5b60016002600087868151811015156200010857fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff19169115159190911790556001929092019162000074565b84516200015c9060039060208801906200016e565b50505060049190915550620002029050565b828054828255906000526020600020908101928215620001c6579160200282015b82811115620001c65782518254600160a060020a031916600160a060020a039091161782556020909201916001909101906200018f565b50620001d4929150620001d8565b5090565b620001ff91905b80821115620001d4578054600160a060020a0319168155600101620001df565b90565b61184180620002126000396000f30060806040526004361061011c5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663025e7c27811461015e578063173825d91461019f57806320ea8d86146101cd5780632f54bf6e146101e55780633411c81c1461022757806354741525146102585780637065cb4814610289578063784547a7146102b75780638b51d13f146102cf5780639ace38c2146102e7578063a0e67e2b146103bc578063a8abe69a14610421578063b5dc40c314610446578063b77bf6001461045e578063ba51a6df14610473578063c01a8c841461048b578063c6427474146104a3578063d74f8edd14610519578063dc8452cd1461052e578063e20056e614610543578063ee22610b14610577575b600034111561015c5760408051348152905133917fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c919081900360200190a25b005b34801561016a57600080fd5b5061017660043561058f565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156101ab57600080fd5b5061015c73ffffffffffffffffffffffffffffffffffffffff600435166105c4565b3480156101d957600080fd5b5061015c60043561081b565b3480156101f157600080fd5b5061021373ffffffffffffffffffffffffffffffffffffffff600435166108f3565b604080519115158252519081900360200190f35b34801561023357600080fd5b5061021360043573ffffffffffffffffffffffffffffffffffffffff60243516610908565b34801561026457600080fd5b5061027760043515156024351515610928565b60408051918252519081900360200190f35b34801561029557600080fd5b5061015c73ffffffffffffffffffffffffffffffffffffffff60043516610994565b3480156102c357600080fd5b50610213600435610b09565b3480156102db57600080fd5b50610277600435610b9a565b3480156102f357600080fd5b506102ff600435610c16565b604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018060200183151515158152602001828103825284818151815260200191508051906020019080838360005b8381101561037e578181015183820152602001610366565b50505050905090810190601f1680156103ab5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b3480156103c857600080fd5b506103d1610cff565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561040d5781810151838201526020016103f5565b505050509050019250505060405180910390f35b34801561042d57600080fd5b506103d160043560243560443515156064351515610d6f565b34801561045257600080fd5b506103d1600435610ea8565b34801561046a57600080fd5b50610277611055565b34801561047f57600080fd5b5061015c60043561105b565b34801561049757600080fd5b5061015c6004356110da565b3480156104af57600080fd5b50604080516020600460443581810135601f810184900484028501840190955284845261027794823573ffffffffffffffffffffffffffffffffffffffff169460248035953695946064949201919081908401838280828437509497506111d09650505050505050565b34801561052557600080fd5b506102776111ef565b34801561053a57600080fd5b506102776111f4565b34801561054f57600080fd5b5061015c73ffffffffffffffffffffffffffffffffffffffff600435811690602435166111fa565b34801561058357600080fd5b5061015c60043561140a565b600380548290811061059d57fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b60003330146105d257600080fd5b73ffffffffffffffffffffffffffffffffffffffff8216600090815260026020526040902054829060ff16151561060857600080fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260026020526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905591505b6003547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0182101561078b578273ffffffffffffffffffffffffffffffffffffffff166003838154811015156106a857fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16141561078057600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810190811061070057fe5b6000918252602090912001546003805473ffffffffffffffffffffffffffffffffffffffff909216918490811061073357fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061078b565b600190910190610656565b600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01906107bc9082611754565b5060035460045411156107d5576003546107d59061105b565b60405173ffffffffffffffffffffffffffffffffffffffff8416907f8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b9090600090a2505050565b3360008181526002602052604090205460ff16151561083957600080fd5b60008281526001602090815260408083203380855292529091205483919060ff16151561086557600080fd5b600084815260208190526040902060030154849060ff161561088657600080fd5b600085815260016020908152604080832033808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690555187927ff6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e991a35050505050565b60026020526000908152604090205460ff1681565b600160209081526000928352604080842090915290825290205460ff1681565b6000805b60055481101561098d57838015610955575060008181526020819052604090206003015460ff16155b806109795750828015610979575060008181526020819052604090206003015460ff165b15610985576001820191505b60010161092c565b5092915050565b3330146109a057600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116600090815260026020526040902054819060ff16156109d557600080fd5b8173ffffffffffffffffffffffffffffffffffffffff811615156109f857600080fd5b60038054905060010160045460328211158015610a155750818111155b8015610a2057508015155b8015610a2b57508115155b1515610a3657600080fd5b73ffffffffffffffffffffffffffffffffffffffff851660008181526002602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915560038054918201815583527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055517ff39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d9190a25050505050565b600080805b600354811015610b935760008481526001602052604081206003805491929184908110610b3757fe5b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff16835282019290925260400190205460ff1615610b78576001820191505b600454821415610b8b5760019250610b93565b600101610b0e565b5050919050565b6000805b600354811015610c105760008381526001602052604081206003805491929184908110610bc757fe5b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff16835282019290925260400190205460ff1615610c08576001820191505b600101610b9e565b50919050565b60006020818152918152604090819020805460018083015460028085018054875161010095821615959095027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff011691909104601f810188900488028401880190965285835273ffffffffffffffffffffffffffffffffffffffff90931695909491929190830182828015610cec5780601f10610cc157610100808354040283529160200191610cec565b820191906000526020600020905b815481529060010190602001808311610ccf57829003601f168201915b5050506003909301549192505060ff1684565b60606003805480602002602001604051908101604052809291908181526020018280548015610d6457602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610d39575b505050505090505b90565b606080600080600554604051908082528060200260200182016040528015610da1578160200160208202803883390190505b50925060009150600090505b600554811015610e2857858015610dd6575060008181526020819052604090206003015460ff16155b80610dfa5750848015610dfa575060008181526020819052604090206003015460ff165b15610e2057808383815181101515610e0e57fe5b60209081029091010152600191909101905b600101610dad565b878703604051908082528060200260200182016040528015610e54578160200160208202803883390190505b5093508790505b86811015610e9d578281815181101515610e7157fe5b9060200190602002015184898303815181101515610e8b57fe5b60209081029091010152600101610e5b565b505050949350505050565b606080600080600380549050604051908082528060200260200182016040528015610edd578160200160208202803883390190505b50925060009150600090505b600354811015610fc15760008581526001602052604081206003805491929184908110610f1257fe5b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff16835282019290925260400190205460ff1615610fb9576003805482908110610f5a57fe5b600091825260209091200154835173ffffffffffffffffffffffffffffffffffffffff90911690849084908110610f8d57fe5b73ffffffffffffffffffffffffffffffffffffffff909216602092830290910190910152600191909101905b600101610ee9565b81604051908082528060200260200182016040528015610feb578160200160208202803883390190505b509350600090505b8181101561104d57828181518110151561100957fe5b90602001906020020151848281518110151561102157fe5b73ffffffffffffffffffffffffffffffffffffffff909216602092830290910190910152600101610ff3565b505050919050565b60055481565b33301461106757600080fd5b600354816032821180159061107c5750818111155b801561108757508015155b801561109257508115155b151561109d57600080fd5b60048390556040805184815290517fa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a9181900360200190a1505050565b3360008181526002602052604090205460ff1615156110f857600080fd5b600082815260208190526040902054829073ffffffffffffffffffffffffffffffffffffffff16151561112a57600080fd5b60008381526001602090815260408083203380855292529091205484919060ff161561115557600080fd5b600085815260016020818152604080842033808652925280842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016909317909255905187927f4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef91a36111c98561140a565b5050505050565b60006111dd848484611621565b90506111e8816110da565b9392505050565b603281565b60045481565b600033301461120857600080fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260026020526040902054839060ff16151561123e57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260026020526040902054839060ff161561127357600080fd5b600092505b600354831015611338578473ffffffffffffffffffffffffffffffffffffffff166003848154811015156112a857fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16141561132d57836003848154811015156112e057fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611338565b600190920191611278565b73ffffffffffffffffffffffffffffffffffffffff80861660008181526002602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0090811690915593881682528082208054909416600117909355915190917f8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b9091a260405173ffffffffffffffffffffffffffffffffffffffff8516907ff39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d90600090a25050505050565b3360008181526002602052604081205490919060ff16151561142b57600080fd5b60008381526001602090815260408083203380855292529091205484919060ff16151561145757600080fd5b600085815260208190526040902060030154859060ff161561147857600080fd5b61148186610b09565b1561161957600086815260208190526040908190206003810180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001908117909155815481830154935160028085018054959b5073ffffffffffffffffffffffffffffffffffffffff909316959492939192839285927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9183161561010002919091019091160480156115785780601f1061154d57610100808354040283529160200191611578565b820191906000526020600020905b81548152906001019060200180831161155b57829003601f168201915b505091505060006040518083038185875af192505050156115c35760405186907f33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed7590600090a2611619565b60405186907f526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b7923690600090a26003850180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690555b505050505050565b60008373ffffffffffffffffffffffffffffffffffffffff8116151561164657600080fd5b6005546040805160808101825273ffffffffffffffffffffffffffffffffffffffff8881168252602080830189815283850189815260006060860181905287815280845295909520845181547fffffffffffffffffffffffff000000000000000000000000000000000000000016941693909317835551600183015592518051949650919390926116de92600285019291019061177d565b5060609190910151600390910180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001691151591909117905560058054600101905560405182907fc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e5190600090a2509392505050565b815481835581811115611778576000838152602090206117789181019083016117fb565b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106117be57805160ff19168380011785556117eb565b828001600101855582156117eb579182015b828111156117eb5782518255916020019190600101906117d0565b506117f79291506117fb565b5090565b610d6c91905b808211156117f757600081556001016118015600a165627a7a723058207f548769687dc741a4f9cff2873fbbcbb25d6832afb68c844df0d87f4c2b0d920029",
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
	Bin: "0x60806040523480156200001157600080fd5b5060405162001c5238038062001c52833981016040908152815160208301519183015192018051909290839083906000908260328211801590620000555750818111155b80156200006157508015155b80156200006d57508115155b15156200007957600080fd5b600092505b845183101562000151576002600086858151811015156200009b57fe5b6020908102909101810151600160a060020a031682528101919091526040016000205460ff16158015620000f157508483815181101515620000d957fe5b90602001906020020151600160a060020a0316600014155b1515620000fd57600080fd5b60016002600087868151811015156200011257fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff1916911515919091179055600192909201916200007e565b8451620001669060039060208801906200017d565b505050600491909155505060065550620002119050565b828054828255906000526020600020908101928215620001d5579160200282015b82811115620001d55782518254600160a060020a031916600160a060020a039091161782556020909201916001909101906200019e565b50620001e3929150620001e7565b5090565b6200020e91905b80821115620001e3578054600160a060020a0319168155600101620001ee565b90565b611a3180620002216000396000f3006080604052600436106101535763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663025e7c278114610195578063173825d9146101d657806320ea8d86146102045780632f54bf6e1461021c5780633411c81c1461025e5780634bc9fdc21461028f57806354741525146102b657806367eeba0c146102d55780636b0c932d146102ea5780637065cb48146102ff578063784547a71461032d5780638b51d13f146103455780639ace38c21461035d578063a0e67e2b14610432578063a8abe69a14610497578063b5dc40c3146104bc578063b77bf600146104d4578063ba51a6df146104e9578063c01a8c8414610501578063c642747414610519578063cea086211461058f578063d74f8edd146105a7578063dc8452cd146105bc578063e20056e6146105d1578063ee22610b14610605578063f059cf2b1461061d575b60003411156101935760408051348152905133917fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c919081900360200190a25b005b3480156101a157600080fd5b506101ad600435610632565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156101e257600080fd5b5061019373ffffffffffffffffffffffffffffffffffffffff60043516610667565b34801561021057600080fd5b506101936004356108be565b34801561022857600080fd5b5061024a73ffffffffffffffffffffffffffffffffffffffff60043516610996565b604080519115158252519081900360200190f35b34801561026a57600080fd5b5061024a60043573ffffffffffffffffffffffffffffffffffffffff602435166109ab565b34801561029b57600080fd5b506102a46109cb565b60408051918252519081900360200190f35b3480156102c257600080fd5b506102a460043515156024351515610a05565b3480156102e157600080fd5b506102a4610a71565b3480156102f657600080fd5b506102a4610a77565b34801561030b57600080fd5b5061019373ffffffffffffffffffffffffffffffffffffffff60043516610a7d565b34801561033957600080fd5b5061024a600435610bf2565b34801561035157600080fd5b506102a4600435610c83565b34801561036957600080fd5b50610375600435610cff565b604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018060200183151515158152602001828103825284818151815260200191508051906020019080838360005b838110156103f45781810151838201526020016103dc565b50505050905090810190601f1680156104215780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b34801561043e57600080fd5b50610447610de8565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561048357818101518382015260200161046b565b505050509050019250505060405180910390f35b3480156104a357600080fd5b5061044760043560243560443515156064351515610e57565b3480156104c857600080fd5b50610447600435610f90565b3480156104e057600080fd5b506102a461113d565b3480156104f557600080fd5b50610193600435611143565b34801561050d57600080fd5b506101936004356111c2565b34801561052557600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526102a494823573ffffffffffffffffffffffffffffffffffffffff169460248035953695946064949201919081908401838280828437509497506112b89650505050505050565b34801561059b57600080fd5b506101936004356112d7565b3480156105b357600080fd5b506102a461131e565b3480156105c857600080fd5b506102a4611323565b3480156105dd57600080fd5b5061019373ffffffffffffffffffffffffffffffffffffffff60043581169060243516611329565b34801561061157600080fd5b50610193600435611539565b34801561062957600080fd5b506102a46117c3565b600380548290811061064057fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b600033301461067557600080fd5b73ffffffffffffffffffffffffffffffffffffffff8216600090815260026020526040902054829060ff1615156106ab57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260026020526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905591505b6003547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0182101561082e578273ffffffffffffffffffffffffffffffffffffffff1660038381548110151561074b57fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16141561082357600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81019081106107a357fe5b6000918252602090912001546003805473ffffffffffffffffffffffffffffffffffffffff90921691849081106107d657fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061082e565b6001909101906106f9565b600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019061085f9082611944565b5060035460045411156108785760035461087890611143565b60405173ffffffffffffffffffffffffffffffffffffffff8416907f8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b9090600090a2505050565b3360008181526002602052604090205460ff1615156108dc57600080fd5b60008281526001602090815260408083203380855292529091205483919060ff16151561090857600080fd5b600084815260208190526040902060030154849060ff161561092957600080fd5b600085815260016020908152604080832033808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690555187927ff6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e991a35050505050565b60026020526000908152604090205460ff1681565b600160209081526000928352604080842090915290825290205460ff1681565b600060075462015180014211156109e55750600654610a02565b60085460065410156109f957506000610a02565b50600854600654035b90565b6000805b600554811015610a6a57838015610a32575060008181526020819052604090206003015460ff16155b80610a565750828015610a56575060008181526020819052604090206003015460ff165b15610a62576001820191505b600101610a09565b5092915050565b60065481565b60075481565b333014610a8957600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116600090815260026020526040902054819060ff1615610abe57600080fd5b8173ffffffffffffffffffffffffffffffffffffffff81161515610ae157600080fd5b60038054905060010160045460328211158015610afe5750818111155b8015610b0957508015155b8015610b1457508115155b1515610b1f57600080fd5b73ffffffffffffffffffffffffffffffffffffffff851660008181526002602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915560038054918201815583527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180547fffffffffffffffffffffffff00000000000000000000000000000000000000001684179055517ff39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d9190a25050505050565b600080805b600354811015610c7c5760008481526001602052604081206003805491929184908110610c2057fe5b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff16835282019290925260400190205460ff1615610c61576001820191505b600454821415610c745760019250610c7c565b600101610bf7565b5050919050565b6000805b600354811015610cf95760008381526001602052604081206003805491929184908110610cb057fe5b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff16835282019290925260400190205460ff1615610cf1576001820191505b600101610c87565b50919050565b60006020818152918152604090819020805460018083015460028085018054875161010095821615959095027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff011691909104601f810188900488028401880190965285835273ffffffffffffffffffffffffffffffffffffffff90931695909491929190830182828015610dd55780601f10610daa57610100808354040283529160200191610dd5565b820191906000526020600020905b815481529060010190602001808311610db857829003601f168201915b5050506003909301549192505060ff1684565b60606003805480602002602001604051908101604052809291908181526020018280548015610e4d57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610e22575b5050505050905090565b606080600080600554604051908082528060200260200182016040528015610e89578160200160208202803883390190505b50925060009150600090505b600554811015610f1057858015610ebe575060008181526020819052604090206003015460ff16155b80610ee25750848015610ee2575060008181526020819052604090206003015460ff165b15610f0857808383815181101515610ef657fe5b60209081029091010152600191909101905b600101610e95565b878703604051908082528060200260200182016040528015610f3c578160200160208202803883390190505b5093508790505b86811015610f85578281815181101515610f5957fe5b9060200190602002015184898303815181101515610f7357fe5b60209081029091010152600101610f43565b505050949350505050565b606080600080600380549050604051908082528060200260200182016040528015610fc5578160200160208202803883390190505b50925060009150600090505b6003548110156110a95760008581526001602052604081206003805491929184908110610ffa57fe5b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff16835282019290925260400190205460ff16156110a157600380548290811061104257fe5b600091825260209091200154835173ffffffffffffffffffffffffffffffffffffffff9091169084908490811061107557fe5b73ffffffffffffffffffffffffffffffffffffffff909216602092830290910190910152600191909101905b600101610fd1565b816040519080825280602002602001820160405280156110d3578160200160208202803883390190505b509350600090505b818110156111355782818151811015156110f157fe5b90602001906020020151848281518110151561110957fe5b73ffffffffffffffffffffffffffffffffffffffff9092166020928302909101909101526001016110db565b505050919050565b60055481565b33301461114f57600080fd5b60035481603282118015906111645750818111155b801561116f57508015155b801561117a57508115155b151561118557600080fd5b60048390556040805184815290517fa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a9181900360200190a1505050565b3360008181526002602052604090205460ff1615156111e057600080fd5b600082815260208190526040902054829073ffffffffffffffffffffffffffffffffffffffff16151561121257600080fd5b60008381526001602090815260408083203380855292529091205484919060ff161561123d57600080fd5b600085815260016020818152604080842033808652925280842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016909317909255905187927f4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef91a36112b185611539565b5050505050565b60006112c58484846117c9565b90506112d0816111c2565b9392505050565b3330146112e357600080fd5b60068190556040805182815290517fc71bdc6afaf9b1aa90a7078191d4fc1adf3bf680fca3183697df6b0dc226bca29181900360200190a150565b603281565b60045481565b600033301461133757600080fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260026020526040902054839060ff16151561136d57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260026020526040902054839060ff16156113a257600080fd5b600092505b600354831015611467578473ffffffffffffffffffffffffffffffffffffffff166003848154811015156113d757fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16141561145c578360038481548110151561140f57fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611467565b6001909201916113a7565b73ffffffffffffffffffffffffffffffffffffffff80861660008181526002602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0090811690915593881682528082208054909416600117909355915190917f8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b9091a260405173ffffffffffffffffffffffffffffffffffffffff8516907ff39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d90600090a25050505050565b336000818152600260205260408120549091829160ff16151561155b57600080fd5b60008481526001602090815260408083203380855292529091205485919060ff16151561158757600080fd5b600086815260208190526040902060030154869060ff16156115a857600080fd5b600087815260208190526040902095506115c187610bf2565b9450848061161257506002808701547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff610100600183161502011604158015611612575061161286600101546118fc565b156117ba576003860180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905584151561165a5760018601546008805490910190555b8560000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1686600101548760020160405180828054600181600116156101000203166002900480156117035780601f106116d857610100808354040283529160200191611703565b820191906000526020600020905b8154815290600101906020018083116116e657829003601f168201915b505091505060006040518083038185875af1925050501561174e5760405187907f33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed7590600090a26117ba565b60405187907f526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b7923690600090a26003860180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690558415156117ba576001860154600880549190910390555b50505050505050565b60085481565b60008373ffffffffffffffffffffffffffffffffffffffff811615156117ee57600080fd5b6005546040805160808101825273ffffffffffffffffffffffffffffffffffffffff8881168252602080830189815283850189815260006060860181905287815280845295909520845181547fffffffffffffffffffffffff0000000000000000000000000000000000000000169416939093178355516001830155925180519496509193909261188692600285019291019061196d565b5060609190910151600390910180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001691151591909117905560058054600101905560405182907fc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e5190600090a2509392505050565b60006007546201518001421115611917574260075560006008555b6006548260085401118061192e5750600854828101105b1561193b5750600061193f565b5060015b919050565b815481835581811115611968576000838152602090206119689181019083016119eb565b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106119ae57805160ff19168380011785556119db565b828001600101855582156119db579182015b828111156119db5782518255916020019190600101906119c0565b506119e79291506119eb565b5090565b610a0291905b808211156119e757600081556001016119f15600a165627a7a72305820f1f74cb4e860417568381b2430672b76ffb99a658ee6644d318105eeebbe0a720029",
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
	Bin: "0x6080604052348015600f57600080fd5b50603580601d6000396000f3006080604052600080fd00a165627a7a72305820585cdde9f5ee3fa0421f12c587d90e1072d72539e5905729b9320ac0a2a159df0029",
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
	Bin: "0x608060405234801561001057600080fd5b50610570806100206000396000f3006080604052600436106100825763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461008757806318160ddd146100cc57806323b872dd146100f357806370a082311461012a578063a9059cbb14610158578063dd62ed3e14610189578063eefa597b146101bd575b600080fd5b34801561009357600080fd5b506100b873ffffffffffffffffffffffffffffffffffffffff600435166024356101d2565b604080519115158252519081900360200190f35b3480156100d857600080fd5b506100e1610290565b60408051918252519081900360200190f35b3480156100ff57600080fd5b506100b873ffffffffffffffffffffffffffffffffffffffff60043581169060243516604435610296565b34801561013657600080fd5b506100e173ffffffffffffffffffffffffffffffffffffffff600435166103c7565b34801561016457600080fd5b506100b873ffffffffffffffffffffffffffffffffffffffff600435166024356103ef565b34801561019557600080fd5b506100e173ffffffffffffffffffffffffffffffffffffffff600435811690602435166104c0565b3480156101c957600080fd5b506100b86104f8565b60008115801590610212575033600090815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915290205415155b1561021c57600080fd5b33600081815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a350600192915050565b60005481565b73ffffffffffffffffffffffffffffffffffffffff8084166000908152600260209081526040808320338452825280832054938616835260019091528120549091906102e290846104fd565b73ffffffffffffffffffffffffffffffffffffffff808616600090815260016020526040808220939093559087168152205461031e9084610521565b73ffffffffffffffffffffffffffffffffffffffff861660009081526001602052604090205561034e8184610521565b73ffffffffffffffffffffffffffffffffffffffff808716600081815260026020908152604080832033845282529182902094909455805187815290519288169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a3506001949350505050565b73ffffffffffffffffffffffffffffffffffffffff1660009081526001602052604090205490565b60006040604436101561040157600080fd5b3360009081526001602052604090205461041b9084610521565b336000908152600160205260408082209290925573ffffffffffffffffffffffffffffffffffffffff86168152205461045490846104fd565b73ffffffffffffffffffffffffffffffffffffffff85166000818152600160209081526040918290209390935580518681529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35060019392505050565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260026020908152604080832093909416825291909152205490565b600181565b600082820161051a8482108015906105155750838210155b610535565b9392505050565b600061052f83831115610535565b50900390565b80151561054157600080fd5b505600a165627a7a723058206f9c688fab4f07ea29b67c6f89882674efa4034f4e2eca38a476e2bc04f419150029",
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
	Bin: "0x608060405234801561001057600080fd5b50604051602080610c3c833981016040525160038054600160a060020a031916600160a060020a03909216919091179055610bec806100506000396000f3006080604052600436106100da5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b381146100df57806318160ddd1461012457806323b872dd1461014b57806345977d03146101825780635de4ccb01461019c578063600440cb146101da57806370a08231146101ef5780638444b3911461021d5780639738968c14610256578063a9059cbb1461026b578063c752ff621461029c578063d7e7088a146102b1578063dd62ed3e146102df578063eefa597b14610313578063ffeb7d7514610328575b600080fd5b3480156100eb57600080fd5b5061011073ffffffffffffffffffffffffffffffffffffffff60043516602435610356565b604080519115158252519081900360200190f35b34801561013057600080fd5b50610139610414565b60408051918252519081900360200190f35b34801561015757600080fd5b5061011073ffffffffffffffffffffffffffffffffffffffff6004358116906024351660443561041a565b34801561018e57600080fd5b5061019a60043561054b565b005b3480156101a857600080fd5b506101b16106c6565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3480156101e657600080fd5b506101b16106e2565b3480156101fb57600080fd5b5061013973ffffffffffffffffffffffffffffffffffffffff600435166106fe565b34801561022957600080fd5b50610232610726565b6040518082600481111561024257fe5b60ff16815260200191505060405180910390f35b34801561026257600080fd5b5061011061077d565b34801561027757600080fd5b5061011073ffffffffffffffffffffffffffffffffffffffff60043516602435610782565b3480156102a857600080fd5b50610139610853565b3480156102bd57600080fd5b5061019a73ffffffffffffffffffffffffffffffffffffffff60043516610859565b3480156102eb57600080fd5b5061013973ffffffffffffffffffffffffffffffffffffffff60043581169060243516610aaf565b34801561031f57600080fd5b50610110610ae7565b34801561033457600080fd5b5061019a73ffffffffffffffffffffffffffffffffffffffff60043516610aec565b60008115801590610396575033600090815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915290205415155b156103a057600080fd5b33600081815260026020908152604080832073ffffffffffffffffffffffffffffffffffffffff881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a350600192915050565b60005481565b73ffffffffffffffffffffffffffffffffffffffff8084166000908152600260209081526040808320338452825280832054938616835260019091528120549091906104669084610b79565b73ffffffffffffffffffffffffffffffffffffffff80861660009081526001602052604080822093909355908716815220546104a29084610b9d565b73ffffffffffffffffffffffffffffffffffffffff86166000908152600160205260409020556104d28184610b9d565b73ffffffffffffffffffffffffffffffffffffffff808716600081815260026020908152604080832033845282529182902094909455805187815290519288169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a3506001949350505050565b6000610555610726565b9050600381600481111561056557fe5b148061057c5750600481600481111561057a57fe5b145b151561058757600080fd5b81151561059357600080fd5b336000908152600160205260409020546105ad9083610b9d565b33600090815260016020526040812091909155546105cb9083610b9d565b6000556005546105db9083610b79565b60055560048054604080517f753e88e50000000000000000000000000000000000000000000000000000000081523393810193909352602483018590525173ffffffffffffffffffffffffffffffffffffffff9091169163753e88e591604480830192600092919082900301818387803b15801561065857600080fd5b505af115801561066c573d6000803e3d6000fd5b505060045460408051868152905173ffffffffffffffffffffffffffffffffffffffff90921693503392507f7e5c344a8141a805725cb476f76c6953b842222b967edd1f78ddb6e8b3f397ac919081900360200190a35050565b60045473ffffffffffffffffffffffffffffffffffffffff1681565b60035473ffffffffffffffffffffffffffffffffffffffff1681565b73ffffffffffffffffffffffffffffffffffffffff1660009081526001602052604090205490565b600061073061077d565b151561073e5750600161077a565b60045473ffffffffffffffffffffffffffffffffffffffff1615156107655750600261077a565b60055415156107765750600361077a565b5060045b90565b600190565b60006040604436101561079457600080fd5b336000908152600160205260409020546107ae9084610b9d565b336000908152600160205260408082209290925573ffffffffffffffffffffffffffffffffffffffff8616815220546107e79084610b79565b73ffffffffffffffffffffffffffffffffffffffff85166000818152600160209081526040918290209390935580518681529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35060019392505050565b60055481565b61086161077d565b151561086c57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561088e57600080fd5b60035473ffffffffffffffffffffffffffffffffffffffff1633146108b257600080fd5b60046108bc610726565b60048111156108c757fe5b14156108d257600080fd5b600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691909117808355604080517f61d3d7a6000000000000000000000000000000000000000000000000000000008152905191909216926361d3d7a6928082019260209290918290030181600087803b15801561096a57600080fd5b505af115801561097e573d6000803e3d6000fd5b505050506040513d602081101561099457600080fd5b505115156109a157600080fd5b600054600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634b2ba0dd6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015610a2a57600080fd5b505af1158015610a3e573d6000803e3d6000fd5b505050506040513d6020811015610a5457600080fd5b505114610a6057600080fd5b6004546040805173ffffffffffffffffffffffffffffffffffffffff9092168252517f7845d5aa74cc410e35571258d954f23b82276e160fe8c188fa80566580f279cc9181900360200190a150565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260026020908152604080832093909416825291909152205490565b600181565b73ffffffffffffffffffffffffffffffffffffffff81161515610b0e57600080fd5b60035473ffffffffffffffffffffffffffffffffffffffff163314610b3257600080fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6000828201610b96848210801590610b915750838210155b610bb1565b9392505050565b6000610bab83831115610bb1565b50900390565b801515610bbd57600080fd5b505600a165627a7a723058200a0c79ba01cab7fa84c7a3f021f7399c8c21320594155927196d547ff186ec540029",
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

