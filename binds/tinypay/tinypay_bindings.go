// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tinypay

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

// TinypayMetaData contains all meta data concerning the Tinypay contract.
var TinypayMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"AccountInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"CoinSupported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"tail\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"DepositMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"FundsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newTail\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"PaymentCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldLimit\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newLimit\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"PaymentLimitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"merchant\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"commitHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expiryTime\",\"type\":\"uint64\"}],\"name\":\"PreCommitMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"oldTail\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newTail\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"tailUpdateCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"TailRefreshed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldLimit\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newLimit\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"TailUpdatesLimitSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NATIVE_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addCoinSupport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"opt\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitHash\",\"type\":\"bytes32\"}],\"name\":\"completePayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"tail\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRate\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getSystemStats\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalDeposits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWithdrawals\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"currentFeeRate\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserLimits\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"paymentLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"tailUpdateCount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTailUpdates\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserTail\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"paymaster_\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"feeRate_\",\"type\":\"uint64\"}],\"name\":\"initSystem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isAccountInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isCoinSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"opt\",\"type\":\"bytes\"}],\"name\":\"merchantPrecommit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paymaster\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"newTail\",\"type\":\"bytes\"}],\"name\":\"refreshTail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPaymaster\",\"type\":\"address\"}],\"name\":\"setPaymaster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"}],\"name\":\"setPaymentLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"}],\"name\":\"setTailUpdatesLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newFeeRate\",\"type\":\"uint64\"}],\"name\":\"updateFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// TinypayABI is the input ABI used to generate the binding from.
// Deprecated: Use TinypayMetaData.ABI instead.
var TinypayABI = TinypayMetaData.ABI

// Tinypay is an auto generated Go binding around an Ethereum contract.
type Tinypay struct {
	TinypayCaller     // Read-only binding to the contract
	TinypayTransactor // Write-only binding to the contract
	TinypayFilterer   // Log filterer for contract events
}

// TinypayCaller is an auto generated read-only Go binding around an Ethereum contract.
type TinypayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TinypayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TinypayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TinypayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TinypayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TinypaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TinypaySession struct {
	Contract     *Tinypay          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TinypayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TinypayCallerSession struct {
	Contract *TinypayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TinypayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TinypayTransactorSession struct {
	Contract     *TinypayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TinypayRaw is an auto generated low-level Go binding around an Ethereum contract.
type TinypayRaw struct {
	Contract *Tinypay // Generic contract binding to access the raw methods on
}

// TinypayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TinypayCallerRaw struct {
	Contract *TinypayCaller // Generic read-only contract binding to access the raw methods on
}

// TinypayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TinypayTransactorRaw struct {
	Contract *TinypayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTinypay creates a new instance of Tinypay, bound to a specific deployed contract.
func NewTinypay(address common.Address, backend bind.ContractBackend) (*Tinypay, error) {
	contract, err := bindTinypay(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tinypay{TinypayCaller: TinypayCaller{contract: contract}, TinypayTransactor: TinypayTransactor{contract: contract}, TinypayFilterer: TinypayFilterer{contract: contract}}, nil
}

// NewTinypayCaller creates a new read-only instance of Tinypay, bound to a specific deployed contract.
func NewTinypayCaller(address common.Address, caller bind.ContractCaller) (*TinypayCaller, error) {
	contract, err := bindTinypay(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TinypayCaller{contract: contract}, nil
}

// NewTinypayTransactor creates a new write-only instance of Tinypay, bound to a specific deployed contract.
func NewTinypayTransactor(address common.Address, transactor bind.ContractTransactor) (*TinypayTransactor, error) {
	contract, err := bindTinypay(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TinypayTransactor{contract: contract}, nil
}

// NewTinypayFilterer creates a new log filterer instance of Tinypay, bound to a specific deployed contract.
func NewTinypayFilterer(address common.Address, filterer bind.ContractFilterer) (*TinypayFilterer, error) {
	contract, err := bindTinypay(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TinypayFilterer{contract: contract}, nil
}

// bindTinypay binds a generic wrapper to an already deployed contract.
func bindTinypay(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TinypayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tinypay *TinypayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tinypay.Contract.TinypayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tinypay *TinypayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tinypay.Contract.TinypayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tinypay *TinypayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tinypay.Contract.TinypayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tinypay *TinypayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tinypay.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tinypay *TinypayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tinypay.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tinypay *TinypayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tinypay.Contract.contract.Transact(opts, method, params...)
}

// NATIVETOKEN is a free data retrieval call binding the contract method 0x31f7d964.
//
// Solidity: function NATIVE_TOKEN() view returns(address)
func (_Tinypay *TinypayCaller) NATIVETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tinypay.contract.Call(opts, &out, "NATIVE_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVETOKEN is a free data retrieval call binding the contract method 0x31f7d964.
//
// Solidity: function NATIVE_TOKEN() view returns(address)
func (_Tinypay *TinypaySession) NATIVETOKEN() (common.Address, error) {
	return _Tinypay.Contract.NATIVETOKEN(&_Tinypay.CallOpts)
}

// NATIVETOKEN is a free data retrieval call binding the contract method 0x31f7d964.
//
// Solidity: function NATIVE_TOKEN() view returns(address)
func (_Tinypay *TinypayCallerSession) NATIVETOKEN() (common.Address, error) {
	return _Tinypay.Contract.NATIVETOKEN(&_Tinypay.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Tinypay *TinypayCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tinypay.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Tinypay *TinypaySession) Admin() (common.Address, error) {
	return _Tinypay.Contract.Admin(&_Tinypay.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Tinypay *TinypayCallerSession) Admin() (common.Address, error) {
	return _Tinypay.Contract.Admin(&_Tinypay.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint64)
func (_Tinypay *TinypayCaller) FeeRate(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Tinypay.contract.Call(opts, &out, "feeRate")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint64)
func (_Tinypay *TinypaySession) FeeRate() (uint64, error) {
	return _Tinypay.Contract.FeeRate(&_Tinypay.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint64)
func (_Tinypay *TinypayCallerSession) FeeRate() (uint64, error) {
	return _Tinypay.Contract.FeeRate(&_Tinypay.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xd4fac45d.
//
// Solidity: function getBalance(address user, address token) view returns(uint256)
func (_Tinypay *TinypayCaller) GetBalance(opts *bind.CallOpts, user common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tinypay.contract.Call(opts, &out, "getBalance", user, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xd4fac45d.
//
// Solidity: function getBalance(address user, address token) view returns(uint256)
func (_Tinypay *TinypaySession) GetBalance(user common.Address, token common.Address) (*big.Int, error) {
	return _Tinypay.Contract.GetBalance(&_Tinypay.CallOpts, user, token)
}

// GetBalance is a free data retrieval call binding the contract method 0xd4fac45d.
//
// Solidity: function getBalance(address user, address token) view returns(uint256)
func (_Tinypay *TinypayCallerSession) GetBalance(user common.Address, token common.Address) (*big.Int, error) {
	return _Tinypay.Contract.GetBalance(&_Tinypay.CallOpts, user, token)
}

// GetSystemStats is a free data retrieval call binding the contract method 0x431cefb8.
//
// Solidity: function getSystemStats(address token) view returns(uint256 totalDeposits, uint256 totalWithdrawals, uint64 currentFeeRate)
func (_Tinypay *TinypayCaller) GetSystemStats(opts *bind.CallOpts, token common.Address) (struct {
	TotalDeposits    *big.Int
	TotalWithdrawals *big.Int
	CurrentFeeRate   uint64
}, error) {
	var out []interface{}
	err := _Tinypay.contract.Call(opts, &out, "getSystemStats", token)

	outstruct := new(struct {
		TotalDeposits    *big.Int
		TotalWithdrawals *big.Int
		CurrentFeeRate   uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalDeposits = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalWithdrawals = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CurrentFeeRate = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetSystemStats is a free data retrieval call binding the contract method 0x431cefb8.
//
// Solidity: function getSystemStats(address token) view returns(uint256 totalDeposits, uint256 totalWithdrawals, uint64 currentFeeRate)
func (_Tinypay *TinypaySession) GetSystemStats(token common.Address) (struct {
	TotalDeposits    *big.Int
	TotalWithdrawals *big.Int
	CurrentFeeRate   uint64
}, error) {
	return _Tinypay.Contract.GetSystemStats(&_Tinypay.CallOpts, token)
}

// GetSystemStats is a free data retrieval call binding the contract method 0x431cefb8.
//
// Solidity: function getSystemStats(address token) view returns(uint256 totalDeposits, uint256 totalWithdrawals, uint64 currentFeeRate)
func (_Tinypay *TinypayCallerSession) GetSystemStats(token common.Address) (struct {
	TotalDeposits    *big.Int
	TotalWithdrawals *big.Int
	CurrentFeeRate   uint64
}, error) {
	return _Tinypay.Contract.GetSystemStats(&_Tinypay.CallOpts, token)
}

// GetUserLimits is a free data retrieval call binding the contract method 0xf562e122.
//
// Solidity: function getUserLimits(address user) view returns(uint64 paymentLimit, uint64 tailUpdateCount, uint64 maxTailUpdates)
func (_Tinypay *TinypayCaller) GetUserLimits(opts *bind.CallOpts, user common.Address) (struct {
	PaymentLimit    uint64
	TailUpdateCount uint64
	MaxTailUpdates  uint64
}, error) {
	var out []interface{}
	err := _Tinypay.contract.Call(opts, &out, "getUserLimits", user)

	outstruct := new(struct {
		PaymentLimit    uint64
		TailUpdateCount uint64
		MaxTailUpdates  uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PaymentLimit = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.TailUpdateCount = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.MaxTailUpdates = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetUserLimits is a free data retrieval call binding the contract method 0xf562e122.
//
// Solidity: function getUserLimits(address user) view returns(uint64 paymentLimit, uint64 tailUpdateCount, uint64 maxTailUpdates)
func (_Tinypay *TinypaySession) GetUserLimits(user common.Address) (struct {
	PaymentLimit    uint64
	TailUpdateCount uint64
	MaxTailUpdates  uint64
}, error) {
	return _Tinypay.Contract.GetUserLimits(&_Tinypay.CallOpts, user)
}

// GetUserLimits is a free data retrieval call binding the contract method 0xf562e122.
//
// Solidity: function getUserLimits(address user) view returns(uint64 paymentLimit, uint64 tailUpdateCount, uint64 maxTailUpdates)
func (_Tinypay *TinypayCallerSession) GetUserLimits(user common.Address) (struct {
	PaymentLimit    uint64
	TailUpdateCount uint64
	MaxTailUpdates  uint64
}, error) {
	return _Tinypay.Contract.GetUserLimits(&_Tinypay.CallOpts, user)
}

// GetUserTail is a free data retrieval call binding the contract method 0xdae457eb.
//
// Solidity: function getUserTail(address user) view returns(bytes)
func (_Tinypay *TinypayCaller) GetUserTail(opts *bind.CallOpts, user common.Address) ([]byte, error) {
	var out []interface{}
	err := _Tinypay.contract.Call(opts, &out, "getUserTail", user)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetUserTail is a free data retrieval call binding the contract method 0xdae457eb.
//
// Solidity: function getUserTail(address user) view returns(bytes)
func (_Tinypay *TinypaySession) GetUserTail(user common.Address) ([]byte, error) {
	return _Tinypay.Contract.GetUserTail(&_Tinypay.CallOpts, user)
}

// GetUserTail is a free data retrieval call binding the contract method 0xdae457eb.
//
// Solidity: function getUserTail(address user) view returns(bytes)
func (_Tinypay *TinypayCallerSession) GetUserTail(user common.Address) ([]byte, error) {
	return _Tinypay.Contract.GetUserTail(&_Tinypay.CallOpts, user)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Tinypay *TinypayCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Tinypay.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Tinypay *TinypaySession) Initialized() (bool, error) {
	return _Tinypay.Contract.Initialized(&_Tinypay.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Tinypay *TinypayCallerSession) Initialized() (bool, error) {
	return _Tinypay.Contract.Initialized(&_Tinypay.CallOpts)
}

// IsAccountInitialized is a free data retrieval call binding the contract method 0x8a413710.
//
// Solidity: function isAccountInitialized(address user) view returns(bool)
func (_Tinypay *TinypayCaller) IsAccountInitialized(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _Tinypay.contract.Call(opts, &out, "isAccountInitialized", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAccountInitialized is a free data retrieval call binding the contract method 0x8a413710.
//
// Solidity: function isAccountInitialized(address user) view returns(bool)
func (_Tinypay *TinypaySession) IsAccountInitialized(user common.Address) (bool, error) {
	return _Tinypay.Contract.IsAccountInitialized(&_Tinypay.CallOpts, user)
}

// IsAccountInitialized is a free data retrieval call binding the contract method 0x8a413710.
//
// Solidity: function isAccountInitialized(address user) view returns(bool)
func (_Tinypay *TinypayCallerSession) IsAccountInitialized(user common.Address) (bool, error) {
	return _Tinypay.Contract.IsAccountInitialized(&_Tinypay.CallOpts, user)
}

// IsCoinSupported is a free data retrieval call binding the contract method 0x81236d62.
//
// Solidity: function isCoinSupported(address token) view returns(bool)
func (_Tinypay *TinypayCaller) IsCoinSupported(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _Tinypay.contract.Call(opts, &out, "isCoinSupported", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCoinSupported is a free data retrieval call binding the contract method 0x81236d62.
//
// Solidity: function isCoinSupported(address token) view returns(bool)
func (_Tinypay *TinypaySession) IsCoinSupported(token common.Address) (bool, error) {
	return _Tinypay.Contract.IsCoinSupported(&_Tinypay.CallOpts, token)
}

// IsCoinSupported is a free data retrieval call binding the contract method 0x81236d62.
//
// Solidity: function isCoinSupported(address token) view returns(bool)
func (_Tinypay *TinypayCallerSession) IsCoinSupported(token common.Address) (bool, error) {
	return _Tinypay.Contract.IsCoinSupported(&_Tinypay.CallOpts, token)
}

// Paymaster is a free data retrieval call binding the contract method 0x16e4cbf9.
//
// Solidity: function paymaster() view returns(address)
func (_Tinypay *TinypayCaller) Paymaster(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tinypay.contract.Call(opts, &out, "paymaster")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Paymaster is a free data retrieval call binding the contract method 0x16e4cbf9.
//
// Solidity: function paymaster() view returns(address)
func (_Tinypay *TinypaySession) Paymaster() (common.Address, error) {
	return _Tinypay.Contract.Paymaster(&_Tinypay.CallOpts)
}

// Paymaster is a free data retrieval call binding the contract method 0x16e4cbf9.
//
// Solidity: function paymaster() view returns(address)
func (_Tinypay *TinypayCallerSession) Paymaster() (common.Address, error) {
	return _Tinypay.Contract.Paymaster(&_Tinypay.CallOpts)
}

// AddCoinSupport is a paid mutator transaction binding the contract method 0x39222dd0.
//
// Solidity: function addCoinSupport(address token) returns()
func (_Tinypay *TinypayTransactor) AddCoinSupport(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Tinypay.contract.Transact(opts, "addCoinSupport", token)
}

// AddCoinSupport is a paid mutator transaction binding the contract method 0x39222dd0.
//
// Solidity: function addCoinSupport(address token) returns()
func (_Tinypay *TinypaySession) AddCoinSupport(token common.Address) (*types.Transaction, error) {
	return _Tinypay.Contract.AddCoinSupport(&_Tinypay.TransactOpts, token)
}

// AddCoinSupport is a paid mutator transaction binding the contract method 0x39222dd0.
//
// Solidity: function addCoinSupport(address token) returns()
func (_Tinypay *TinypayTransactorSession) AddCoinSupport(token common.Address) (*types.Transaction, error) {
	return _Tinypay.Contract.AddCoinSupport(&_Tinypay.TransactOpts, token)
}

// CompletePayment is a paid mutator transaction binding the contract method 0x38dbb8e3.
//
// Solidity: function completePayment(address token, bytes opt, address payer, address recipient, uint256 amount, bytes32 commitHash) returns()
func (_Tinypay *TinypayTransactor) CompletePayment(opts *bind.TransactOpts, token common.Address, opt []byte, payer common.Address, recipient common.Address, amount *big.Int, commitHash [32]byte) (*types.Transaction, error) {
	return _Tinypay.contract.Transact(opts, "completePayment", token, opt, payer, recipient, amount, commitHash)
}

// CompletePayment is a paid mutator transaction binding the contract method 0x38dbb8e3.
//
// Solidity: function completePayment(address token, bytes opt, address payer, address recipient, uint256 amount, bytes32 commitHash) returns()
func (_Tinypay *TinypaySession) CompletePayment(token common.Address, opt []byte, payer common.Address, recipient common.Address, amount *big.Int, commitHash [32]byte) (*types.Transaction, error) {
	return _Tinypay.Contract.CompletePayment(&_Tinypay.TransactOpts, token, opt, payer, recipient, amount, commitHash)
}

// CompletePayment is a paid mutator transaction binding the contract method 0x38dbb8e3.
//
// Solidity: function completePayment(address token, bytes opt, address payer, address recipient, uint256 amount, bytes32 commitHash) returns()
func (_Tinypay *TinypayTransactorSession) CompletePayment(token common.Address, opt []byte, payer common.Address, recipient common.Address, amount *big.Int, commitHash [32]byte) (*types.Transaction, error) {
	return _Tinypay.Contract.CompletePayment(&_Tinypay.TransactOpts, token, opt, payer, recipient, amount, commitHash)
}

// Deposit is a paid mutator transaction binding the contract method 0x49bdc2b8.
//
// Solidity: function deposit(address token, uint256 amount, bytes tail) payable returns()
func (_Tinypay *TinypayTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int, tail []byte) (*types.Transaction, error) {
	return _Tinypay.contract.Transact(opts, "deposit", token, amount, tail)
}

// Deposit is a paid mutator transaction binding the contract method 0x49bdc2b8.
//
// Solidity: function deposit(address token, uint256 amount, bytes tail) payable returns()
func (_Tinypay *TinypaySession) Deposit(token common.Address, amount *big.Int, tail []byte) (*types.Transaction, error) {
	return _Tinypay.Contract.Deposit(&_Tinypay.TransactOpts, token, amount, tail)
}

// Deposit is a paid mutator transaction binding the contract method 0x49bdc2b8.
//
// Solidity: function deposit(address token, uint256 amount, bytes tail) payable returns()
func (_Tinypay *TinypayTransactorSession) Deposit(token common.Address, amount *big.Int, tail []byte) (*types.Transaction, error) {
	return _Tinypay.Contract.Deposit(&_Tinypay.TransactOpts, token, amount, tail)
}

// InitSystem is a paid mutator transaction binding the contract method 0xc2b08b2c.
//
// Solidity: function initSystem(address paymaster_, uint64 feeRate_) returns()
func (_Tinypay *TinypayTransactor) InitSystem(opts *bind.TransactOpts, paymaster_ common.Address, feeRate_ uint64) (*types.Transaction, error) {
	return _Tinypay.contract.Transact(opts, "initSystem", paymaster_, feeRate_)
}

// InitSystem is a paid mutator transaction binding the contract method 0xc2b08b2c.
//
// Solidity: function initSystem(address paymaster_, uint64 feeRate_) returns()
func (_Tinypay *TinypaySession) InitSystem(paymaster_ common.Address, feeRate_ uint64) (*types.Transaction, error) {
	return _Tinypay.Contract.InitSystem(&_Tinypay.TransactOpts, paymaster_, feeRate_)
}

// InitSystem is a paid mutator transaction binding the contract method 0xc2b08b2c.
//
// Solidity: function initSystem(address paymaster_, uint64 feeRate_) returns()
func (_Tinypay *TinypayTransactorSession) InitSystem(paymaster_ common.Address, feeRate_ uint64) (*types.Transaction, error) {
	return _Tinypay.Contract.InitSystem(&_Tinypay.TransactOpts, paymaster_, feeRate_)
}

// MerchantPrecommit is a paid mutator transaction binding the contract method 0x432aec1d.
//
// Solidity: function merchantPrecommit(address token, address payer, address recipient, uint256 amount, bytes opt) returns()
func (_Tinypay *TinypayTransactor) MerchantPrecommit(opts *bind.TransactOpts, token common.Address, payer common.Address, recipient common.Address, amount *big.Int, opt []byte) (*types.Transaction, error) {
	return _Tinypay.contract.Transact(opts, "merchantPrecommit", token, payer, recipient, amount, opt)
}

// MerchantPrecommit is a paid mutator transaction binding the contract method 0x432aec1d.
//
// Solidity: function merchantPrecommit(address token, address payer, address recipient, uint256 amount, bytes opt) returns()
func (_Tinypay *TinypaySession) MerchantPrecommit(token common.Address, payer common.Address, recipient common.Address, amount *big.Int, opt []byte) (*types.Transaction, error) {
	return _Tinypay.Contract.MerchantPrecommit(&_Tinypay.TransactOpts, token, payer, recipient, amount, opt)
}

// MerchantPrecommit is a paid mutator transaction binding the contract method 0x432aec1d.
//
// Solidity: function merchantPrecommit(address token, address payer, address recipient, uint256 amount, bytes opt) returns()
func (_Tinypay *TinypayTransactorSession) MerchantPrecommit(token common.Address, payer common.Address, recipient common.Address, amount *big.Int, opt []byte) (*types.Transaction, error) {
	return _Tinypay.Contract.MerchantPrecommit(&_Tinypay.TransactOpts, token, payer, recipient, amount, opt)
}

// RefreshTail is a paid mutator transaction binding the contract method 0x035c5313.
//
// Solidity: function refreshTail(bytes newTail) returns()
func (_Tinypay *TinypayTransactor) RefreshTail(opts *bind.TransactOpts, newTail []byte) (*types.Transaction, error) {
	return _Tinypay.contract.Transact(opts, "refreshTail", newTail)
}

// RefreshTail is a paid mutator transaction binding the contract method 0x035c5313.
//
// Solidity: function refreshTail(bytes newTail) returns()
func (_Tinypay *TinypaySession) RefreshTail(newTail []byte) (*types.Transaction, error) {
	return _Tinypay.Contract.RefreshTail(&_Tinypay.TransactOpts, newTail)
}

// RefreshTail is a paid mutator transaction binding the contract method 0x035c5313.
//
// Solidity: function refreshTail(bytes newTail) returns()
func (_Tinypay *TinypayTransactorSession) RefreshTail(newTail []byte) (*types.Transaction, error) {
	return _Tinypay.Contract.RefreshTail(&_Tinypay.TransactOpts, newTail)
}

// SetPaymaster is a paid mutator transaction binding the contract method 0x2a97fa77.
//
// Solidity: function setPaymaster(address newPaymaster) returns()
func (_Tinypay *TinypayTransactor) SetPaymaster(opts *bind.TransactOpts, newPaymaster common.Address) (*types.Transaction, error) {
	return _Tinypay.contract.Transact(opts, "setPaymaster", newPaymaster)
}

// SetPaymaster is a paid mutator transaction binding the contract method 0x2a97fa77.
//
// Solidity: function setPaymaster(address newPaymaster) returns()
func (_Tinypay *TinypaySession) SetPaymaster(newPaymaster common.Address) (*types.Transaction, error) {
	return _Tinypay.Contract.SetPaymaster(&_Tinypay.TransactOpts, newPaymaster)
}

// SetPaymaster is a paid mutator transaction binding the contract method 0x2a97fa77.
//
// Solidity: function setPaymaster(address newPaymaster) returns()
func (_Tinypay *TinypayTransactorSession) SetPaymaster(newPaymaster common.Address) (*types.Transaction, error) {
	return _Tinypay.Contract.SetPaymaster(&_Tinypay.TransactOpts, newPaymaster)
}

// SetPaymentLimit is a paid mutator transaction binding the contract method 0x9c2a281f.
//
// Solidity: function setPaymentLimit(uint64 limit) returns()
func (_Tinypay *TinypayTransactor) SetPaymentLimit(opts *bind.TransactOpts, limit uint64) (*types.Transaction, error) {
	return _Tinypay.contract.Transact(opts, "setPaymentLimit", limit)
}

// SetPaymentLimit is a paid mutator transaction binding the contract method 0x9c2a281f.
//
// Solidity: function setPaymentLimit(uint64 limit) returns()
func (_Tinypay *TinypaySession) SetPaymentLimit(limit uint64) (*types.Transaction, error) {
	return _Tinypay.Contract.SetPaymentLimit(&_Tinypay.TransactOpts, limit)
}

// SetPaymentLimit is a paid mutator transaction binding the contract method 0x9c2a281f.
//
// Solidity: function setPaymentLimit(uint64 limit) returns()
func (_Tinypay *TinypayTransactorSession) SetPaymentLimit(limit uint64) (*types.Transaction, error) {
	return _Tinypay.Contract.SetPaymentLimit(&_Tinypay.TransactOpts, limit)
}

// SetTailUpdatesLimit is a paid mutator transaction binding the contract method 0x5f542478.
//
// Solidity: function setTailUpdatesLimit(uint64 limit) returns()
func (_Tinypay *TinypayTransactor) SetTailUpdatesLimit(opts *bind.TransactOpts, limit uint64) (*types.Transaction, error) {
	return _Tinypay.contract.Transact(opts, "setTailUpdatesLimit", limit)
}

// SetTailUpdatesLimit is a paid mutator transaction binding the contract method 0x5f542478.
//
// Solidity: function setTailUpdatesLimit(uint64 limit) returns()
func (_Tinypay *TinypaySession) SetTailUpdatesLimit(limit uint64) (*types.Transaction, error) {
	return _Tinypay.Contract.SetTailUpdatesLimit(&_Tinypay.TransactOpts, limit)
}

// SetTailUpdatesLimit is a paid mutator transaction binding the contract method 0x5f542478.
//
// Solidity: function setTailUpdatesLimit(uint64 limit) returns()
func (_Tinypay *TinypayTransactorSession) SetTailUpdatesLimit(limit uint64) (*types.Transaction, error) {
	return _Tinypay.Contract.SetTailUpdatesLimit(&_Tinypay.TransactOpts, limit)
}

// UpdateFeeRate is a paid mutator transaction binding the contract method 0xbb9b767a.
//
// Solidity: function updateFeeRate(uint64 newFeeRate) returns()
func (_Tinypay *TinypayTransactor) UpdateFeeRate(opts *bind.TransactOpts, newFeeRate uint64) (*types.Transaction, error) {
	return _Tinypay.contract.Transact(opts, "updateFeeRate", newFeeRate)
}

// UpdateFeeRate is a paid mutator transaction binding the contract method 0xbb9b767a.
//
// Solidity: function updateFeeRate(uint64 newFeeRate) returns()
func (_Tinypay *TinypaySession) UpdateFeeRate(newFeeRate uint64) (*types.Transaction, error) {
	return _Tinypay.Contract.UpdateFeeRate(&_Tinypay.TransactOpts, newFeeRate)
}

// UpdateFeeRate is a paid mutator transaction binding the contract method 0xbb9b767a.
//
// Solidity: function updateFeeRate(uint64 newFeeRate) returns()
func (_Tinypay *TinypayTransactorSession) UpdateFeeRate(newFeeRate uint64) (*types.Transaction, error) {
	return _Tinypay.Contract.UpdateFeeRate(&_Tinypay.TransactOpts, newFeeRate)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x1095b6d7.
//
// Solidity: function withdrawFee(address token, address to, uint256 amount) returns()
func (_Tinypay *TinypayTransactor) WithdrawFee(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tinypay.contract.Transact(opts, "withdrawFee", token, to, amount)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x1095b6d7.
//
// Solidity: function withdrawFee(address token, address to, uint256 amount) returns()
func (_Tinypay *TinypaySession) WithdrawFee(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tinypay.Contract.WithdrawFee(&_Tinypay.TransactOpts, token, to, amount)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x1095b6d7.
//
// Solidity: function withdrawFee(address token, address to, uint256 amount) returns()
func (_Tinypay *TinypayTransactorSession) WithdrawFee(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tinypay.Contract.WithdrawFee(&_Tinypay.TransactOpts, token, to, amount)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address token, uint256 amount) returns()
func (_Tinypay *TinypayTransactor) WithdrawFunds(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tinypay.contract.Transact(opts, "withdrawFunds", token, amount)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address token, uint256 amount) returns()
func (_Tinypay *TinypaySession) WithdrawFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tinypay.Contract.WithdrawFunds(&_Tinypay.TransactOpts, token, amount)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address token, uint256 amount) returns()
func (_Tinypay *TinypayTransactorSession) WithdrawFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tinypay.Contract.WithdrawFunds(&_Tinypay.TransactOpts, token, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Tinypay *TinypayTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tinypay.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Tinypay *TinypaySession) Receive() (*types.Transaction, error) {
	return _Tinypay.Contract.Receive(&_Tinypay.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Tinypay *TinypayTransactorSession) Receive() (*types.Transaction, error) {
	return _Tinypay.Contract.Receive(&_Tinypay.TransactOpts)
}

// TinypayAccountInitializedIterator is returned from FilterAccountInitialized and is used to iterate over the raw logs and unpacked data for AccountInitialized events raised by the Tinypay contract.
type TinypayAccountInitializedIterator struct {
	Event *TinypayAccountInitialized // Event containing the contract specifics and raw log

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
func (it *TinypayAccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TinypayAccountInitialized)
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
		it.Event = new(TinypayAccountInitialized)
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
func (it *TinypayAccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TinypayAccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TinypayAccountInitialized represents a AccountInitialized event raised by the Tinypay contract.
type TinypayAccountInitialized struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAccountInitialized is a free log retrieval operation binding the contract event 0x9b44a2c7f9f0b5aa4e7da60d8a2325796c57f93a45559d66515743ffbd8a6103.
//
// Solidity: event AccountInitialized(address indexed user)
func (_Tinypay *TinypayFilterer) FilterAccountInitialized(opts *bind.FilterOpts, user []common.Address) (*TinypayAccountInitializedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Tinypay.contract.FilterLogs(opts, "AccountInitialized", userRule)
	if err != nil {
		return nil, err
	}
	return &TinypayAccountInitializedIterator{contract: _Tinypay.contract, event: "AccountInitialized", logs: logs, sub: sub}, nil
}

// WatchAccountInitialized is a free log subscription operation binding the contract event 0x9b44a2c7f9f0b5aa4e7da60d8a2325796c57f93a45559d66515743ffbd8a6103.
//
// Solidity: event AccountInitialized(address indexed user)
func (_Tinypay *TinypayFilterer) WatchAccountInitialized(opts *bind.WatchOpts, sink chan<- *TinypayAccountInitialized, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Tinypay.contract.WatchLogs(opts, "AccountInitialized", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TinypayAccountInitialized)
				if err := _Tinypay.contract.UnpackLog(event, "AccountInitialized", log); err != nil {
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

// ParseAccountInitialized is a log parse operation binding the contract event 0x9b44a2c7f9f0b5aa4e7da60d8a2325796c57f93a45559d66515743ffbd8a6103.
//
// Solidity: event AccountInitialized(address indexed user)
func (_Tinypay *TinypayFilterer) ParseAccountInitialized(log types.Log) (*TinypayAccountInitialized, error) {
	event := new(TinypayAccountInitialized)
	if err := _Tinypay.contract.UnpackLog(event, "AccountInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TinypayCoinSupportedIterator is returned from FilterCoinSupported and is used to iterate over the raw logs and unpacked data for CoinSupported events raised by the Tinypay contract.
type TinypayCoinSupportedIterator struct {
	Event *TinypayCoinSupported // Event containing the contract specifics and raw log

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
func (it *TinypayCoinSupportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TinypayCoinSupported)
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
		it.Event = new(TinypayCoinSupported)
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
func (it *TinypayCoinSupportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TinypayCoinSupportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TinypayCoinSupported represents a CoinSupported event raised by the Tinypay contract.
type TinypayCoinSupported struct {
	Token     common.Address
	Timestamp uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCoinSupported is a free log retrieval operation binding the contract event 0x6ddab6aaca30e17f45dfeeb37269228fafd598b908afedb75005bcde8c715b83.
//
// Solidity: event CoinSupported(address indexed token, uint64 timestamp)
func (_Tinypay *TinypayFilterer) FilterCoinSupported(opts *bind.FilterOpts, token []common.Address) (*TinypayCoinSupportedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Tinypay.contract.FilterLogs(opts, "CoinSupported", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TinypayCoinSupportedIterator{contract: _Tinypay.contract, event: "CoinSupported", logs: logs, sub: sub}, nil
}

// WatchCoinSupported is a free log subscription operation binding the contract event 0x6ddab6aaca30e17f45dfeeb37269228fafd598b908afedb75005bcde8c715b83.
//
// Solidity: event CoinSupported(address indexed token, uint64 timestamp)
func (_Tinypay *TinypayFilterer) WatchCoinSupported(opts *bind.WatchOpts, sink chan<- *TinypayCoinSupported, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Tinypay.contract.WatchLogs(opts, "CoinSupported", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TinypayCoinSupported)
				if err := _Tinypay.contract.UnpackLog(event, "CoinSupported", log); err != nil {
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

// ParseCoinSupported is a log parse operation binding the contract event 0x6ddab6aaca30e17f45dfeeb37269228fafd598b908afedb75005bcde8c715b83.
//
// Solidity: event CoinSupported(address indexed token, uint64 timestamp)
func (_Tinypay *TinypayFilterer) ParseCoinSupported(log types.Log) (*TinypayCoinSupported, error) {
	event := new(TinypayCoinSupported)
	if err := _Tinypay.contract.UnpackLog(event, "CoinSupported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TinypayDepositMadeIterator is returned from FilterDepositMade and is used to iterate over the raw logs and unpacked data for DepositMade events raised by the Tinypay contract.
type TinypayDepositMadeIterator struct {
	Event *TinypayDepositMade // Event containing the contract specifics and raw log

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
func (it *TinypayDepositMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TinypayDepositMade)
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
		it.Event = new(TinypayDepositMade)
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
func (it *TinypayDepositMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TinypayDepositMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TinypayDepositMade represents a DepositMade event raised by the Tinypay contract.
type TinypayDepositMade struct {
	User       common.Address
	Token      common.Address
	Amount     *big.Int
	Tail       []byte
	NewBalance *big.Int
	Timestamp  uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDepositMade is a free log retrieval operation binding the contract event 0x6bda3803844cd2181842ee537512a127c06e832ddd3ed5f1ecd511049194e360.
//
// Solidity: event DepositMade(address indexed user, address indexed token, uint256 amount, bytes tail, uint256 newBalance, uint64 timestamp)
func (_Tinypay *TinypayFilterer) FilterDepositMade(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*TinypayDepositMadeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Tinypay.contract.FilterLogs(opts, "DepositMade", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &TinypayDepositMadeIterator{contract: _Tinypay.contract, event: "DepositMade", logs: logs, sub: sub}, nil
}

// WatchDepositMade is a free log subscription operation binding the contract event 0x6bda3803844cd2181842ee537512a127c06e832ddd3ed5f1ecd511049194e360.
//
// Solidity: event DepositMade(address indexed user, address indexed token, uint256 amount, bytes tail, uint256 newBalance, uint64 timestamp)
func (_Tinypay *TinypayFilterer) WatchDepositMade(opts *bind.WatchOpts, sink chan<- *TinypayDepositMade, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Tinypay.contract.WatchLogs(opts, "DepositMade", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TinypayDepositMade)
				if err := _Tinypay.contract.UnpackLog(event, "DepositMade", log); err != nil {
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

// ParseDepositMade is a log parse operation binding the contract event 0x6bda3803844cd2181842ee537512a127c06e832ddd3ed5f1ecd511049194e360.
//
// Solidity: event DepositMade(address indexed user, address indexed token, uint256 amount, bytes tail, uint256 newBalance, uint64 timestamp)
func (_Tinypay *TinypayFilterer) ParseDepositMade(log types.Log) (*TinypayDepositMade, error) {
	event := new(TinypayDepositMade)
	if err := _Tinypay.contract.UnpackLog(event, "DepositMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TinypayFundsWithdrawnIterator is returned from FilterFundsWithdrawn and is used to iterate over the raw logs and unpacked data for FundsWithdrawn events raised by the Tinypay contract.
type TinypayFundsWithdrawnIterator struct {
	Event *TinypayFundsWithdrawn // Event containing the contract specifics and raw log

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
func (it *TinypayFundsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TinypayFundsWithdrawn)
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
		it.Event = new(TinypayFundsWithdrawn)
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
func (it *TinypayFundsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TinypayFundsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TinypayFundsWithdrawn represents a FundsWithdrawn event raised by the Tinypay contract.
type TinypayFundsWithdrawn struct {
	User       common.Address
	Token      common.Address
	Amount     *big.Int
	NewBalance *big.Int
	Timestamp  uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFundsWithdrawn is a free log retrieval operation binding the contract event 0x1dcaccdd300ec19c43aa96dec60320f5631e8f76a0d6dd35a19e6ef96d1a187d.
//
// Solidity: event FundsWithdrawn(address indexed user, address indexed token, uint256 amount, uint256 newBalance, uint64 timestamp)
func (_Tinypay *TinypayFilterer) FilterFundsWithdrawn(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*TinypayFundsWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Tinypay.contract.FilterLogs(opts, "FundsWithdrawn", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &TinypayFundsWithdrawnIterator{contract: _Tinypay.contract, event: "FundsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFundsWithdrawn is a free log subscription operation binding the contract event 0x1dcaccdd300ec19c43aa96dec60320f5631e8f76a0d6dd35a19e6ef96d1a187d.
//
// Solidity: event FundsWithdrawn(address indexed user, address indexed token, uint256 amount, uint256 newBalance, uint64 timestamp)
func (_Tinypay *TinypayFilterer) WatchFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *TinypayFundsWithdrawn, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Tinypay.contract.WatchLogs(opts, "FundsWithdrawn", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TinypayFundsWithdrawn)
				if err := _Tinypay.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
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

// ParseFundsWithdrawn is a log parse operation binding the contract event 0x1dcaccdd300ec19c43aa96dec60320f5631e8f76a0d6dd35a19e6ef96d1a187d.
//
// Solidity: event FundsWithdrawn(address indexed user, address indexed token, uint256 amount, uint256 newBalance, uint64 timestamp)
func (_Tinypay *TinypayFilterer) ParseFundsWithdrawn(log types.Log) (*TinypayFundsWithdrawn, error) {
	event := new(TinypayFundsWithdrawn)
	if err := _Tinypay.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TinypayPaymentCompletedIterator is returned from FilterPaymentCompleted and is used to iterate over the raw logs and unpacked data for PaymentCompleted events raised by the Tinypay contract.
type TinypayPaymentCompletedIterator struct {
	Event *TinypayPaymentCompleted // Event containing the contract specifics and raw log

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
func (it *TinypayPaymentCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TinypayPaymentCompleted)
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
		it.Event = new(TinypayPaymentCompleted)
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
func (it *TinypayPaymentCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TinypayPaymentCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TinypayPaymentCompleted represents a PaymentCompleted event raised by the Tinypay contract.
type TinypayPaymentCompleted struct {
	Payer     common.Address
	Recipient common.Address
	Token     common.Address
	Amount    *big.Int
	Fee       *big.Int
	NewTail   []byte
	Timestamp uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPaymentCompleted is a free log retrieval operation binding the contract event 0xddb469e464a641bf383b430ca46c464bdd44236156bfb30408d12e10cd1afdbb.
//
// Solidity: event PaymentCompleted(address indexed payer, address indexed recipient, address indexed token, uint256 amount, uint256 fee, bytes newTail, uint64 timestamp)
func (_Tinypay *TinypayFilterer) FilterPaymentCompleted(opts *bind.FilterOpts, payer []common.Address, recipient []common.Address, token []common.Address) (*TinypayPaymentCompletedIterator, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Tinypay.contract.FilterLogs(opts, "PaymentCompleted", payerRule, recipientRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &TinypayPaymentCompletedIterator{contract: _Tinypay.contract, event: "PaymentCompleted", logs: logs, sub: sub}, nil
}

// WatchPaymentCompleted is a free log subscription operation binding the contract event 0xddb469e464a641bf383b430ca46c464bdd44236156bfb30408d12e10cd1afdbb.
//
// Solidity: event PaymentCompleted(address indexed payer, address indexed recipient, address indexed token, uint256 amount, uint256 fee, bytes newTail, uint64 timestamp)
func (_Tinypay *TinypayFilterer) WatchPaymentCompleted(opts *bind.WatchOpts, sink chan<- *TinypayPaymentCompleted, payer []common.Address, recipient []common.Address, token []common.Address) (event.Subscription, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Tinypay.contract.WatchLogs(opts, "PaymentCompleted", payerRule, recipientRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TinypayPaymentCompleted)
				if err := _Tinypay.contract.UnpackLog(event, "PaymentCompleted", log); err != nil {
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

// ParsePaymentCompleted is a log parse operation binding the contract event 0xddb469e464a641bf383b430ca46c464bdd44236156bfb30408d12e10cd1afdbb.
//
// Solidity: event PaymentCompleted(address indexed payer, address indexed recipient, address indexed token, uint256 amount, uint256 fee, bytes newTail, uint64 timestamp)
func (_Tinypay *TinypayFilterer) ParsePaymentCompleted(log types.Log) (*TinypayPaymentCompleted, error) {
	event := new(TinypayPaymentCompleted)
	if err := _Tinypay.contract.UnpackLog(event, "PaymentCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TinypayPaymentLimitUpdatedIterator is returned from FilterPaymentLimitUpdated and is used to iterate over the raw logs and unpacked data for PaymentLimitUpdated events raised by the Tinypay contract.
type TinypayPaymentLimitUpdatedIterator struct {
	Event *TinypayPaymentLimitUpdated // Event containing the contract specifics and raw log

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
func (it *TinypayPaymentLimitUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TinypayPaymentLimitUpdated)
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
		it.Event = new(TinypayPaymentLimitUpdated)
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
func (it *TinypayPaymentLimitUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TinypayPaymentLimitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TinypayPaymentLimitUpdated represents a PaymentLimitUpdated event raised by the Tinypay contract.
type TinypayPaymentLimitUpdated struct {
	User      common.Address
	OldLimit  uint64
	NewLimit  uint64
	Timestamp uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPaymentLimitUpdated is a free log retrieval operation binding the contract event 0x49ff5953104bea65fcf72b51d50d4ecd22f5d51a564a2ce9a584917f089c2bb2.
//
// Solidity: event PaymentLimitUpdated(address indexed user, uint64 oldLimit, uint64 newLimit, uint64 timestamp)
func (_Tinypay *TinypayFilterer) FilterPaymentLimitUpdated(opts *bind.FilterOpts, user []common.Address) (*TinypayPaymentLimitUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Tinypay.contract.FilterLogs(opts, "PaymentLimitUpdated", userRule)
	if err != nil {
		return nil, err
	}
	return &TinypayPaymentLimitUpdatedIterator{contract: _Tinypay.contract, event: "PaymentLimitUpdated", logs: logs, sub: sub}, nil
}

// WatchPaymentLimitUpdated is a free log subscription operation binding the contract event 0x49ff5953104bea65fcf72b51d50d4ecd22f5d51a564a2ce9a584917f089c2bb2.
//
// Solidity: event PaymentLimitUpdated(address indexed user, uint64 oldLimit, uint64 newLimit, uint64 timestamp)
func (_Tinypay *TinypayFilterer) WatchPaymentLimitUpdated(opts *bind.WatchOpts, sink chan<- *TinypayPaymentLimitUpdated, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Tinypay.contract.WatchLogs(opts, "PaymentLimitUpdated", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TinypayPaymentLimitUpdated)
				if err := _Tinypay.contract.UnpackLog(event, "PaymentLimitUpdated", log); err != nil {
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

// ParsePaymentLimitUpdated is a log parse operation binding the contract event 0x49ff5953104bea65fcf72b51d50d4ecd22f5d51a564a2ce9a584917f089c2bb2.
//
// Solidity: event PaymentLimitUpdated(address indexed user, uint64 oldLimit, uint64 newLimit, uint64 timestamp)
func (_Tinypay *TinypayFilterer) ParsePaymentLimitUpdated(log types.Log) (*TinypayPaymentLimitUpdated, error) {
	event := new(TinypayPaymentLimitUpdated)
	if err := _Tinypay.contract.UnpackLog(event, "PaymentLimitUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TinypayPreCommitMadeIterator is returned from FilterPreCommitMade and is used to iterate over the raw logs and unpacked data for PreCommitMade events raised by the Tinypay contract.
type TinypayPreCommitMadeIterator struct {
	Event *TinypayPreCommitMade // Event containing the contract specifics and raw log

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
func (it *TinypayPreCommitMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TinypayPreCommitMade)
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
		it.Event = new(TinypayPreCommitMade)
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
func (it *TinypayPreCommitMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TinypayPreCommitMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TinypayPreCommitMade represents a PreCommitMade event raised by the Tinypay contract.
type TinypayPreCommitMade struct {
	Merchant   common.Address
	Token      common.Address
	CommitHash [32]byte
	ExpiryTime uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPreCommitMade is a free log retrieval operation binding the contract event 0xd752391535a9c0b5f5c3463816f3513b020b4addf2838c77ad6942874530c0f2.
//
// Solidity: event PreCommitMade(address indexed merchant, address indexed token, bytes32 commitHash, uint64 expiryTime)
func (_Tinypay *TinypayFilterer) FilterPreCommitMade(opts *bind.FilterOpts, merchant []common.Address, token []common.Address) (*TinypayPreCommitMadeIterator, error) {

	var merchantRule []interface{}
	for _, merchantItem := range merchant {
		merchantRule = append(merchantRule, merchantItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Tinypay.contract.FilterLogs(opts, "PreCommitMade", merchantRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &TinypayPreCommitMadeIterator{contract: _Tinypay.contract, event: "PreCommitMade", logs: logs, sub: sub}, nil
}

// WatchPreCommitMade is a free log subscription operation binding the contract event 0xd752391535a9c0b5f5c3463816f3513b020b4addf2838c77ad6942874530c0f2.
//
// Solidity: event PreCommitMade(address indexed merchant, address indexed token, bytes32 commitHash, uint64 expiryTime)
func (_Tinypay *TinypayFilterer) WatchPreCommitMade(opts *bind.WatchOpts, sink chan<- *TinypayPreCommitMade, merchant []common.Address, token []common.Address) (event.Subscription, error) {

	var merchantRule []interface{}
	for _, merchantItem := range merchant {
		merchantRule = append(merchantRule, merchantItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Tinypay.contract.WatchLogs(opts, "PreCommitMade", merchantRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TinypayPreCommitMade)
				if err := _Tinypay.contract.UnpackLog(event, "PreCommitMade", log); err != nil {
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

// ParsePreCommitMade is a log parse operation binding the contract event 0xd752391535a9c0b5f5c3463816f3513b020b4addf2838c77ad6942874530c0f2.
//
// Solidity: event PreCommitMade(address indexed merchant, address indexed token, bytes32 commitHash, uint64 expiryTime)
func (_Tinypay *TinypayFilterer) ParsePreCommitMade(log types.Log) (*TinypayPreCommitMade, error) {
	event := new(TinypayPreCommitMade)
	if err := _Tinypay.contract.UnpackLog(event, "PreCommitMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TinypayTailRefreshedIterator is returned from FilterTailRefreshed and is used to iterate over the raw logs and unpacked data for TailRefreshed events raised by the Tinypay contract.
type TinypayTailRefreshedIterator struct {
	Event *TinypayTailRefreshed // Event containing the contract specifics and raw log

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
func (it *TinypayTailRefreshedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TinypayTailRefreshed)
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
		it.Event = new(TinypayTailRefreshed)
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
func (it *TinypayTailRefreshedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TinypayTailRefreshedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TinypayTailRefreshed represents a TailRefreshed event raised by the Tinypay contract.
type TinypayTailRefreshed struct {
	User            common.Address
	OldTail         []byte
	NewTail         []byte
	TailUpdateCount uint64
	Timestamp       uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTailRefreshed is a free log retrieval operation binding the contract event 0x09800539674fe0c3f48f376cf8d8b02e5620874ae0f4238ead3426c0007b53cf.
//
// Solidity: event TailRefreshed(address indexed user, bytes oldTail, bytes newTail, uint64 tailUpdateCount, uint64 timestamp)
func (_Tinypay *TinypayFilterer) FilterTailRefreshed(opts *bind.FilterOpts, user []common.Address) (*TinypayTailRefreshedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Tinypay.contract.FilterLogs(opts, "TailRefreshed", userRule)
	if err != nil {
		return nil, err
	}
	return &TinypayTailRefreshedIterator{contract: _Tinypay.contract, event: "TailRefreshed", logs: logs, sub: sub}, nil
}

// WatchTailRefreshed is a free log subscription operation binding the contract event 0x09800539674fe0c3f48f376cf8d8b02e5620874ae0f4238ead3426c0007b53cf.
//
// Solidity: event TailRefreshed(address indexed user, bytes oldTail, bytes newTail, uint64 tailUpdateCount, uint64 timestamp)
func (_Tinypay *TinypayFilterer) WatchTailRefreshed(opts *bind.WatchOpts, sink chan<- *TinypayTailRefreshed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Tinypay.contract.WatchLogs(opts, "TailRefreshed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TinypayTailRefreshed)
				if err := _Tinypay.contract.UnpackLog(event, "TailRefreshed", log); err != nil {
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

// ParseTailRefreshed is a log parse operation binding the contract event 0x09800539674fe0c3f48f376cf8d8b02e5620874ae0f4238ead3426c0007b53cf.
//
// Solidity: event TailRefreshed(address indexed user, bytes oldTail, bytes newTail, uint64 tailUpdateCount, uint64 timestamp)
func (_Tinypay *TinypayFilterer) ParseTailRefreshed(log types.Log) (*TinypayTailRefreshed, error) {
	event := new(TinypayTailRefreshed)
	if err := _Tinypay.contract.UnpackLog(event, "TailRefreshed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TinypayTailUpdatesLimitSetIterator is returned from FilterTailUpdatesLimitSet and is used to iterate over the raw logs and unpacked data for TailUpdatesLimitSet events raised by the Tinypay contract.
type TinypayTailUpdatesLimitSetIterator struct {
	Event *TinypayTailUpdatesLimitSet // Event containing the contract specifics and raw log

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
func (it *TinypayTailUpdatesLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TinypayTailUpdatesLimitSet)
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
		it.Event = new(TinypayTailUpdatesLimitSet)
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
func (it *TinypayTailUpdatesLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TinypayTailUpdatesLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TinypayTailUpdatesLimitSet represents a TailUpdatesLimitSet event raised by the Tinypay contract.
type TinypayTailUpdatesLimitSet struct {
	User      common.Address
	OldLimit  uint64
	NewLimit  uint64
	Timestamp uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTailUpdatesLimitSet is a free log retrieval operation binding the contract event 0xf38aa731ae46529cfde89aeaea4b5f81871ec7f78675167804f0780be6791a12.
//
// Solidity: event TailUpdatesLimitSet(address indexed user, uint64 oldLimit, uint64 newLimit, uint64 timestamp)
func (_Tinypay *TinypayFilterer) FilterTailUpdatesLimitSet(opts *bind.FilterOpts, user []common.Address) (*TinypayTailUpdatesLimitSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Tinypay.contract.FilterLogs(opts, "TailUpdatesLimitSet", userRule)
	if err != nil {
		return nil, err
	}
	return &TinypayTailUpdatesLimitSetIterator{contract: _Tinypay.contract, event: "TailUpdatesLimitSet", logs: logs, sub: sub}, nil
}

// WatchTailUpdatesLimitSet is a free log subscription operation binding the contract event 0xf38aa731ae46529cfde89aeaea4b5f81871ec7f78675167804f0780be6791a12.
//
// Solidity: event TailUpdatesLimitSet(address indexed user, uint64 oldLimit, uint64 newLimit, uint64 timestamp)
func (_Tinypay *TinypayFilterer) WatchTailUpdatesLimitSet(opts *bind.WatchOpts, sink chan<- *TinypayTailUpdatesLimitSet, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Tinypay.contract.WatchLogs(opts, "TailUpdatesLimitSet", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TinypayTailUpdatesLimitSet)
				if err := _Tinypay.contract.UnpackLog(event, "TailUpdatesLimitSet", log); err != nil {
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

// ParseTailUpdatesLimitSet is a log parse operation binding the contract event 0xf38aa731ae46529cfde89aeaea4b5f81871ec7f78675167804f0780be6791a12.
//
// Solidity: event TailUpdatesLimitSet(address indexed user, uint64 oldLimit, uint64 newLimit, uint64 timestamp)
func (_Tinypay *TinypayFilterer) ParseTailUpdatesLimitSet(log types.Log) (*TinypayTailUpdatesLimitSet, error) {
	event := new(TinypayTailUpdatesLimitSet)
	if err := _Tinypay.contract.UnpackLog(event, "TailUpdatesLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
