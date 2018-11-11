// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package exchange

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// ExchangeABI is the input ABI used to generate the binding from.
const ExchangeABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orderValues\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orders\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cancelledOrders\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LogDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LogBuy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LogWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"placeOrder\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Exchange is an auto generated Go binding around an Ethereum contract.
type Exchange struct {
	ExchangeCaller     // Read-only binding to the contract
	ExchangeTransactor // Write-only binding to the contract
	ExchangeFilterer   // Log filterer for contract events
}

// ExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeSession struct {
	Contract     *Exchange         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeCallerSession struct {
	Contract *ExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeTransactorSession struct {
	Contract     *ExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeRaw struct {
	Contract *Exchange // Generic contract binding to access the raw methods on
}

// ExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeCallerRaw struct {
	Contract *ExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeTransactorRaw struct {
	Contract *ExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchange creates a new instance of Exchange, bound to a specific deployed contract.
func NewExchange(address common.Address, backend bind.ContractBackend) (*Exchange, error) {
	contract, err := bindExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Exchange{ExchangeCaller: ExchangeCaller{contract: contract}, ExchangeTransactor: ExchangeTransactor{contract: contract}, ExchangeFilterer: ExchangeFilterer{contract: contract}}, nil
}

// NewExchangeCaller creates a new read-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeCaller(address common.Address, caller bind.ContractCaller) (*ExchangeCaller, error) {
	contract, err := bindExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeCaller{contract: contract}, nil
}

// NewExchangeTransactor creates a new write-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeTransactor, error) {
	contract, err := bindExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeTransactor{contract: contract}, nil
}

// NewExchangeFilterer creates a new log filterer instance of Exchange, bound to a specific deployed contract.
func NewExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeFilterer, error) {
	contract, err := bindExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeFilterer{contract: contract}, nil
}

// bindExchange binds a generic wrapper to an already deployed contract.
func bindExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.ExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transact(opts, method, params...)
}

// CancelledOrders is a free data retrieval call binding the contract method 0xbc15e601.
//
// Solidity: function cancelledOrders( uint256) constant returns(bool)
func (_Exchange *ExchangeCaller) CancelledOrders(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "cancelledOrders", arg0)
	return *ret0, err
}

// CancelledOrders is a free data retrieval call binding the contract method 0xbc15e601.
//
// Solidity: function cancelledOrders( uint256) constant returns(bool)
func (_Exchange *ExchangeSession) CancelledOrders(arg0 *big.Int) (bool, error) {
	return _Exchange.Contract.CancelledOrders(&_Exchange.CallOpts, arg0)
}

// CancelledOrders is a free data retrieval call binding the contract method 0xbc15e601.
//
// Solidity: function cancelledOrders( uint256) constant returns(bool)
func (_Exchange *ExchangeCallerSession) CancelledOrders(arg0 *big.Int) (bool, error) {
	return _Exchange.Contract.CancelledOrders(&_Exchange.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits( address) constant returns(uint256)
func (_Exchange *ExchangeCaller) Deposits(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "deposits", arg0)
	return *ret0, err
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits( address) constant returns(uint256)
func (_Exchange *ExchangeSession) Deposits(arg0 common.Address) (*big.Int, error) {
	return _Exchange.Contract.Deposits(&_Exchange.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits( address) constant returns(uint256)
func (_Exchange *ExchangeCallerSession) Deposits(arg0 common.Address) (*big.Int, error) {
	return _Exchange.Contract.Deposits(&_Exchange.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Exchange *ExchangeCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Exchange *ExchangeSession) IsOwner() (bool, error) {
	return _Exchange.Contract.IsOwner(&_Exchange.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Exchange *ExchangeCallerSession) IsOwner() (bool, error) {
	return _Exchange.Contract.IsOwner(&_Exchange.CallOpts)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() constant returns(uint256)
func (_Exchange *ExchangeCaller) NextId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "nextId")
	return *ret0, err
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() constant returns(uint256)
func (_Exchange *ExchangeSession) NextId() (*big.Int, error) {
	return _Exchange.Contract.NextId(&_Exchange.CallOpts)
}

// NextId is a free data retrieval call binding the contract method 0x61b8ce8c.
//
// Solidity: function nextId() constant returns(uint256)
func (_Exchange *ExchangeCallerSession) NextId() (*big.Int, error) {
	return _Exchange.Contract.NextId(&_Exchange.CallOpts)
}

// OrderValues is a free data retrieval call binding the contract method 0x33ed0e5f.
//
// Solidity: function orderValues( uint256) constant returns(uint256)
func (_Exchange *ExchangeCaller) OrderValues(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "orderValues", arg0)
	return *ret0, err
}

// OrderValues is a free data retrieval call binding the contract method 0x33ed0e5f.
//
// Solidity: function orderValues( uint256) constant returns(uint256)
func (_Exchange *ExchangeSession) OrderValues(arg0 *big.Int) (*big.Int, error) {
	return _Exchange.Contract.OrderValues(&_Exchange.CallOpts, arg0)
}

// OrderValues is a free data retrieval call binding the contract method 0x33ed0e5f.
//
// Solidity: function orderValues( uint256) constant returns(uint256)
func (_Exchange *ExchangeCallerSession) OrderValues(arg0 *big.Int) (*big.Int, error) {
	return _Exchange.Contract.OrderValues(&_Exchange.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders( uint256) constant returns(bytes32)
func (_Exchange *ExchangeCaller) Orders(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "orders", arg0)
	return *ret0, err
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders( uint256) constant returns(bytes32)
func (_Exchange *ExchangeSession) Orders(arg0 *big.Int) ([32]byte, error) {
	return _Exchange.Contract.Orders(&_Exchange.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders( uint256) constant returns(bytes32)
func (_Exchange *ExchangeCallerSession) Orders(arg0 *big.Int) ([32]byte, error) {
	return _Exchange.Contract.Orders(&_Exchange.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Exchange *ExchangeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Exchange *ExchangeSession) Owner() (common.Address, error) {
	return _Exchange.Contract.Owner(&_Exchange.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Exchange *ExchangeCallerSession) Owner() (common.Address, error) {
	return _Exchange.Contract.Owner(&_Exchange.CallOpts)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(orderId uint256) returns()
func (_Exchange *ExchangeTransactor) CancelOrder(opts *bind.TransactOpts, orderId *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "cancelOrder", orderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(orderId uint256) returns()
func (_Exchange *ExchangeSession) CancelOrder(orderId *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.CancelOrder(&_Exchange.TransactOpts, orderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(orderId uint256) returns()
func (_Exchange *ExchangeTransactorSession) CancelOrder(orderId *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.CancelOrder(&_Exchange.TransactOpts, orderId)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Exchange *ExchangeTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Exchange *ExchangeSession) Deposit() (*types.Transaction, error) {
	return _Exchange.Contract.Deposit(&_Exchange.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Exchange *ExchangeTransactorSession) Deposit() (*types.Transaction, error) {
	return _Exchange.Contract.Deposit(&_Exchange.TransactOpts)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x843f61e2.
//
// Solidity: function placeOrder(amount uint256, price uint256) returns()
func (_Exchange *ExchangeTransactor) PlaceOrder(opts *bind.TransactOpts, amount *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "placeOrder", amount, price)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x843f61e2.
//
// Solidity: function placeOrder(amount uint256, price uint256) returns()
func (_Exchange *ExchangeSession) PlaceOrder(amount *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.PlaceOrder(&_Exchange.TransactOpts, amount, price)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x843f61e2.
//
// Solidity: function placeOrder(amount uint256, price uint256) returns()
func (_Exchange *ExchangeTransactorSession) PlaceOrder(amount *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.PlaceOrder(&_Exchange.TransactOpts, amount, price)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Exchange.Contract.RenounceOwnership(&_Exchange.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Exchange.Contract.RenounceOwnership(&_Exchange.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Exchange *ExchangeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Exchange *ExchangeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TransferOwnership(&_Exchange.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Exchange *ExchangeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TransferOwnership(&_Exchange.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(receiver address, value uint256) returns()
func (_Exchange *ExchangeTransactor) Withdraw(opts *bind.TransactOpts, receiver common.Address, value *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "withdraw", receiver, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(receiver address, value uint256) returns()
func (_Exchange *ExchangeSession) Withdraw(receiver common.Address, value *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Withdraw(&_Exchange.TransactOpts, receiver, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(receiver address, value uint256) returns()
func (_Exchange *ExchangeTransactorSession) Withdraw(receiver common.Address, value *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Withdraw(&_Exchange.TransactOpts, receiver, value)
}

// ExchangeLogBuyIterator is returned from FilterLogBuy and is used to iterate over the raw logs and unpacked data for LogBuy events raised by the Exchange contract.
type ExchangeLogBuyIterator struct {
	Event *ExchangeLogBuy // Event containing the contract specifics and raw log

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
func (it *ExchangeLogBuyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogBuy)
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
		it.Event = new(ExchangeLogBuy)
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
func (it *ExchangeLogBuyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogBuyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogBuy represents a LogBuy event raised by the Exchange contract.
type ExchangeLogBuy struct {
	Sender common.Address
	Amount *big.Int
	Price  *big.Int
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogBuy is a free log retrieval operation binding the contract event 0xc5620ded95cbb91682a998bc6df1a310612e51388b47c88b6dfb3f00d8248ddb.
//
// Solidity: e LogBuy(sender address, amount uint256, price uint256, value uint256)
func (_Exchange *ExchangeFilterer) FilterLogBuy(opts *bind.FilterOpts) (*ExchangeLogBuyIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogBuy")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogBuyIterator{contract: _Exchange.contract, event: "LogBuy", logs: logs, sub: sub}, nil
}

// WatchLogBuy is a free log subscription operation binding the contract event 0xc5620ded95cbb91682a998bc6df1a310612e51388b47c88b6dfb3f00d8248ddb.
//
// Solidity: e LogBuy(sender address, amount uint256, price uint256, value uint256)
func (_Exchange *ExchangeFilterer) WatchLogBuy(opts *bind.WatchOpts, sink chan<- *ExchangeLogBuy) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogBuy")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogBuy)
				if err := _Exchange.contract.UnpackLog(event, "LogBuy", log); err != nil {
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

// ExchangeLogDepositIterator is returned from FilterLogDeposit and is used to iterate over the raw logs and unpacked data for LogDeposit events raised by the Exchange contract.
type ExchangeLogDepositIterator struct {
	Event *ExchangeLogDeposit // Event containing the contract specifics and raw log

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
func (it *ExchangeLogDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogDeposit)
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
		it.Event = new(ExchangeLogDeposit)
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
func (it *ExchangeLogDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogDeposit represents a LogDeposit event raised by the Exchange contract.
type ExchangeLogDeposit struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogDeposit is a free log retrieval operation binding the contract event 0x1b851e1031ef35a238e6c67d0c7991162390df915f70eaf9098dbf0b175a6198.
//
// Solidity: e LogDeposit(sender address, value uint256)
func (_Exchange *ExchangeFilterer) FilterLogDeposit(opts *bind.FilterOpts) (*ExchangeLogDepositIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogDeposit")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogDepositIterator{contract: _Exchange.contract, event: "LogDeposit", logs: logs, sub: sub}, nil
}

// WatchLogDeposit is a free log subscription operation binding the contract event 0x1b851e1031ef35a238e6c67d0c7991162390df915f70eaf9098dbf0b175a6198.
//
// Solidity: e LogDeposit(sender address, value uint256)
func (_Exchange *ExchangeFilterer) WatchLogDeposit(opts *bind.WatchOpts, sink chan<- *ExchangeLogDeposit) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogDeposit)
				if err := _Exchange.contract.UnpackLog(event, "LogDeposit", log); err != nil {
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

// ExchangeLogWithdrawalIterator is returned from FilterLogWithdrawal and is used to iterate over the raw logs and unpacked data for LogWithdrawal events raised by the Exchange contract.
type ExchangeLogWithdrawalIterator struct {
	Event *ExchangeLogWithdrawal // Event containing the contract specifics and raw log

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
func (it *ExchangeLogWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogWithdrawal)
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
		it.Event = new(ExchangeLogWithdrawal)
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
func (it *ExchangeLogWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogWithdrawal represents a LogWithdrawal event raised by the Exchange contract.
type ExchangeLogWithdrawal struct {
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawal is a free log retrieval operation binding the contract event 0xb4214c8c54fc7442f36d3682f59aebaf09358a4431835b30efb29d52cf9e1e91.
//
// Solidity: e LogWithdrawal(receiver address, value uint256)
func (_Exchange *ExchangeFilterer) FilterLogWithdrawal(opts *bind.FilterOpts) (*ExchangeLogWithdrawalIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogWithdrawal")
	if err != nil {
		return nil, err
	}
	return &ExchangeLogWithdrawalIterator{contract: _Exchange.contract, event: "LogWithdrawal", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawal is a free log subscription operation binding the contract event 0xb4214c8c54fc7442f36d3682f59aebaf09358a4431835b30efb29d52cf9e1e91.
//
// Solidity: e LogWithdrawal(receiver address, value uint256)
func (_Exchange *ExchangeFilterer) WatchLogWithdrawal(opts *bind.WatchOpts, sink chan<- *ExchangeLogWithdrawal) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogWithdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogWithdrawal)
				if err := _Exchange.contract.UnpackLog(event, "LogWithdrawal", log); err != nil {
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

// ExchangeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Exchange contract.
type ExchangeOwnershipTransferredIterator struct {
	Event *ExchangeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ExchangeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOwnershipTransferred)
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
		it.Event = new(ExchangeOwnershipTransferred)
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
func (it *ExchangeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOwnershipTransferred represents a OwnershipTransferred event raised by the Exchange contract.
type ExchangeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Exchange *ExchangeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ExchangeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOwnershipTransferredIterator{contract: _Exchange.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Exchange *ExchangeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExchangeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOwnershipTransferred)
				if err := _Exchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
