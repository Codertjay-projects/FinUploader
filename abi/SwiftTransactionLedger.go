// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package transactionLedger

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

// TransactionLedgerMetaData contains all meta data concerning the TransactionLedger contract.
var TransactionLedgerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"NOT_AUTHORIZED\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_transactionReference\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_dateCurrencyAmount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_orderingCustomer\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_beneficiary\",\"type\":\"string\"}],\"name\":\"addTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finFile\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFinFile\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTransaction\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_cid\",\"type\":\"string\"}],\"name\":\"recordFinFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"transactionReference\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dateCurrencyAmount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"orderingCustomer\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"beneficiary\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TransactionLedgerABI is the input ABI used to generate the binding from.
// Deprecated: Use TransactionLedgerMetaData.ABI instead.
var TransactionLedgerABI = TransactionLedgerMetaData.ABI

// TransactionLedger is an auto generated Go binding around an Ethereum contract.
type TransactionLedger struct {
	TransactionLedgerCaller     // Read-only binding to the contract
	TransactionLedgerTransactor // Write-only binding to the contract
	TransactionLedgerFilterer   // Log filterer for contract events
}

// TransactionLedgerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransactionLedgerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionLedgerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransactionLedgerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionLedgerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransactionLedgerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransactionLedgerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransactionLedgerSession struct {
	Contract     *TransactionLedger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TransactionLedgerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransactionLedgerCallerSession struct {
	Contract *TransactionLedgerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// TransactionLedgerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransactionLedgerTransactorSession struct {
	Contract     *TransactionLedgerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TransactionLedgerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransactionLedgerRaw struct {
	Contract *TransactionLedger // Generic contract binding to access the raw methods on
}

// TransactionLedgerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransactionLedgerCallerRaw struct {
	Contract *TransactionLedgerCaller // Generic read-only contract binding to access the raw methods on
}

// TransactionLedgerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransactionLedgerTransactorRaw struct {
	Contract *TransactionLedgerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransactionLedger creates a new instance of TransactionLedger, bound to a specific deployed contract.
func NewTransactionLedger(address common.Address, backend bind.ContractBackend) (*TransactionLedger, error) {
	contract, err := bindTransactionLedger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransactionLedger{TransactionLedgerCaller: TransactionLedgerCaller{contract: contract}, TransactionLedgerTransactor: TransactionLedgerTransactor{contract: contract}, TransactionLedgerFilterer: TransactionLedgerFilterer{contract: contract}}, nil
}

// NewTransactionLedgerCaller creates a new read-only instance of TransactionLedger, bound to a specific deployed contract.
func NewTransactionLedgerCaller(address common.Address, caller bind.ContractCaller) (*TransactionLedgerCaller, error) {
	contract, err := bindTransactionLedger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransactionLedgerCaller{contract: contract}, nil
}

// NewTransactionLedgerTransactor creates a new write-only instance of TransactionLedger, bound to a specific deployed contract.
func NewTransactionLedgerTransactor(address common.Address, transactor bind.ContractTransactor) (*TransactionLedgerTransactor, error) {
	contract, err := bindTransactionLedger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransactionLedgerTransactor{contract: contract}, nil
}

// NewTransactionLedgerFilterer creates a new log filterer instance of TransactionLedger, bound to a specific deployed contract.
func NewTransactionLedgerFilterer(address common.Address, filterer bind.ContractFilterer) (*TransactionLedgerFilterer, error) {
	contract, err := bindTransactionLedger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransactionLedgerFilterer{contract: contract}, nil
}

// bindTransactionLedger binds a generic wrapper to an already deployed contract.
func bindTransactionLedger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TransactionLedgerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransactionLedger *TransactionLedgerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransactionLedger.Contract.TransactionLedgerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransactionLedger *TransactionLedgerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransactionLedger.Contract.TransactionLedgerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransactionLedger *TransactionLedgerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransactionLedger.Contract.TransactionLedgerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransactionLedger *TransactionLedgerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransactionLedger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransactionLedger *TransactionLedgerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransactionLedger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransactionLedger *TransactionLedgerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransactionLedger.Contract.contract.Transact(opts, method, params...)
}

// FinFile is a free data retrieval call binding the contract method 0x07404b4a.
//
// Solidity: function finFile() view returns(string name, string cid)
func (_TransactionLedger *TransactionLedgerCaller) FinFile(opts *bind.CallOpts) (struct {
	Name string
	Cid  string
}, error) {
	var out []interface{}
	err := _TransactionLedger.contract.Call(opts, &out, "finFile")

	outstruct := new(struct {
		Name string
		Cid  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Cid = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// FinFile is a free data retrieval call binding the contract method 0x07404b4a.
//
// Solidity: function finFile() view returns(string name, string cid)
func (_TransactionLedger *TransactionLedgerSession) FinFile() (struct {
	Name string
	Cid  string
}, error) {
	return _TransactionLedger.Contract.FinFile(&_TransactionLedger.CallOpts)
}

// FinFile is a free data retrieval call binding the contract method 0x07404b4a.
//
// Solidity: function finFile() view returns(string name, string cid)
func (_TransactionLedger *TransactionLedgerCallerSession) FinFile() (struct {
	Name string
	Cid  string
}, error) {
	return _TransactionLedger.Contract.FinFile(&_TransactionLedger.CallOpts)
}

// GetFinFile is a free data retrieval call binding the contract method 0x378a8c9b.
//
// Solidity: function getFinFile() view returns(string)
func (_TransactionLedger *TransactionLedgerCaller) GetFinFile(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TransactionLedger.contract.Call(opts, &out, "getFinFile")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetFinFile is a free data retrieval call binding the contract method 0x378a8c9b.
//
// Solidity: function getFinFile() view returns(string)
func (_TransactionLedger *TransactionLedgerSession) GetFinFile() (string, error) {
	return _TransactionLedger.Contract.GetFinFile(&_TransactionLedger.CallOpts)
}

// GetFinFile is a free data retrieval call binding the contract method 0x378a8c9b.
//
// Solidity: function getFinFile() view returns(string)
func (_TransactionLedger *TransactionLedgerCallerSession) GetFinFile() (string, error) {
	return _TransactionLedger.Contract.GetFinFile(&_TransactionLedger.CallOpts)
}

// GetTransaction is a free data retrieval call binding the contract method 0x33ea3dc8.
//
// Solidity: function getTransaction(uint256 _index) view returns(string, string, string, string)
func (_TransactionLedger *TransactionLedgerCaller) GetTransaction(opts *bind.CallOpts, _index *big.Int) (string, string, string, string, error) {
	var out []interface{}
	err := _TransactionLedger.contract.Call(opts, &out, "getTransaction", _index)

	if err != nil {
		return *new(string), *new(string), *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)

	return out0, out1, out2, out3, err

}

// GetTransaction is a free data retrieval call binding the contract method 0x33ea3dc8.
//
// Solidity: function getTransaction(uint256 _index) view returns(string, string, string, string)
func (_TransactionLedger *TransactionLedgerSession) GetTransaction(_index *big.Int) (string, string, string, string, error) {
	return _TransactionLedger.Contract.GetTransaction(&_TransactionLedger.CallOpts, _index)
}

// GetTransaction is a free data retrieval call binding the contract method 0x33ea3dc8.
//
// Solidity: function getTransaction(uint256 _index) view returns(string, string, string, string)
func (_TransactionLedger *TransactionLedgerCallerSession) GetTransaction(_index *big.Int) (string, string, string, string, error) {
	return _TransactionLedger.Contract.GetTransaction(&_TransactionLedger.CallOpts, _index)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(string transactionReference, string dateCurrencyAmount, string orderingCustomer, string beneficiary)
func (_TransactionLedger *TransactionLedgerCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TransactionReference string
	DateCurrencyAmount   string
	OrderingCustomer     string
	Beneficiary          string
}, error) {
	var out []interface{}
	err := _TransactionLedger.contract.Call(opts, &out, "transactions", arg0)

	outstruct := new(struct {
		TransactionReference string
		DateCurrencyAmount   string
		OrderingCustomer     string
		Beneficiary          string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TransactionReference = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.DateCurrencyAmount = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.OrderingCustomer = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Beneficiary = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(string transactionReference, string dateCurrencyAmount, string orderingCustomer, string beneficiary)
func (_TransactionLedger *TransactionLedgerSession) Transactions(arg0 *big.Int) (struct {
	TransactionReference string
	DateCurrencyAmount   string
	OrderingCustomer     string
	Beneficiary          string
}, error) {
	return _TransactionLedger.Contract.Transactions(&_TransactionLedger.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(string transactionReference, string dateCurrencyAmount, string orderingCustomer, string beneficiary)
func (_TransactionLedger *TransactionLedgerCallerSession) Transactions(arg0 *big.Int) (struct {
	TransactionReference string
	DateCurrencyAmount   string
	OrderingCustomer     string
	Beneficiary          string
}, error) {
	return _TransactionLedger.Contract.Transactions(&_TransactionLedger.CallOpts, arg0)
}

// AddTransaction is a paid mutator transaction binding the contract method 0x283c9e41.
//
// Solidity: function addTransaction(string _transactionReference, string _dateCurrencyAmount, string _orderingCustomer, string _beneficiary) returns()
func (_TransactionLedger *TransactionLedgerTransactor) AddTransaction(opts *bind.TransactOpts, _transactionReference string, _dateCurrencyAmount string, _orderingCustomer string, _beneficiary string) (*types.Transaction, error) {
	return _TransactionLedger.contract.Transact(opts, "addTransaction", _transactionReference, _dateCurrencyAmount, _orderingCustomer, _beneficiary)
}

// AddTransaction is a paid mutator transaction binding the contract method 0x283c9e41.
//
// Solidity: function addTransaction(string _transactionReference, string _dateCurrencyAmount, string _orderingCustomer, string _beneficiary) returns()
func (_TransactionLedger *TransactionLedgerSession) AddTransaction(_transactionReference string, _dateCurrencyAmount string, _orderingCustomer string, _beneficiary string) (*types.Transaction, error) {
	return _TransactionLedger.Contract.AddTransaction(&_TransactionLedger.TransactOpts, _transactionReference, _dateCurrencyAmount, _orderingCustomer, _beneficiary)
}

// AddTransaction is a paid mutator transaction binding the contract method 0x283c9e41.
//
// Solidity: function addTransaction(string _transactionReference, string _dateCurrencyAmount, string _orderingCustomer, string _beneficiary) returns()
func (_TransactionLedger *TransactionLedgerTransactorSession) AddTransaction(_transactionReference string, _dateCurrencyAmount string, _orderingCustomer string, _beneficiary string) (*types.Transaction, error) {
	return _TransactionLedger.Contract.AddTransaction(&_TransactionLedger.TransactOpts, _transactionReference, _dateCurrencyAmount, _orderingCustomer, _beneficiary)
}

// RecordFinFile is a paid mutator transaction binding the contract method 0x9280a0f4.
//
// Solidity: function recordFinFile(string _name, string _cid) returns()
func (_TransactionLedger *TransactionLedgerTransactor) RecordFinFile(opts *bind.TransactOpts, _name string, _cid string) (*types.Transaction, error) {
	return _TransactionLedger.contract.Transact(opts, "recordFinFile", _name, _cid)
}

// RecordFinFile is a paid mutator transaction binding the contract method 0x9280a0f4.
//
// Solidity: function recordFinFile(string _name, string _cid) returns()
func (_TransactionLedger *TransactionLedgerSession) RecordFinFile(_name string, _cid string) (*types.Transaction, error) {
	return _TransactionLedger.Contract.RecordFinFile(&_TransactionLedger.TransactOpts, _name, _cid)
}

// RecordFinFile is a paid mutator transaction binding the contract method 0x9280a0f4.
//
// Solidity: function recordFinFile(string _name, string _cid) returns()
func (_TransactionLedger *TransactionLedgerTransactorSession) RecordFinFile(_name string, _cid string) (*types.Transaction, error) {
	return _TransactionLedger.Contract.RecordFinFile(&_TransactionLedger.TransactOpts, _name, _cid)
}
