// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aave

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AaveLPAPABI is the input ABI used to generate the binding from.
const AaveLPAPABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"EthereumAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"FeeProviderUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolConfiguratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolCoreUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolDataProviderUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolLiquidationManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolParametersProviderUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingRateOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PriceOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"ProxyCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"TokenDistributorUpdated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLendingPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"setLendingPoolImpl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLendingPoolCore\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lendingPoolCore\",\"type\":\"address\"}],\"name\":\"setLendingPoolCoreImpl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLendingPoolConfigurator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_configurator\",\"type\":\"address\"}],\"name\":\"setLendingPoolConfiguratorImpl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLendingPoolDataProvider\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"}],\"name\":\"setLendingPoolDataProviderImpl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLendingPoolParametersProvider\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_parametersProvider\",\"type\":\"address\"}],\"name\":\"setLendingPoolParametersProviderImpl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFeeProvider\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeProvider\",\"type\":\"address\"}],\"name\":\"setFeeProviderImpl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLendingPoolLiquidationManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"setLendingPoolLiquidationManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLendingPoolManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lendingPoolManager\",\"type\":\"address\"}],\"name\":\"setLendingPoolManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPriceOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceOracle\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLendingRateOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lendingRateOracle\",\"type\":\"address\"}],\"name\":\"setLendingRateOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTokenDistributor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenDistributor\",\"type\":\"address\"}],\"name\":\"setTokenDistributor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AaveLPAP is an auto generated Go binding around an Ethereum contract.
type AaveLPAP struct {
	AaveLPAPCaller     // Read-only binding to the contract
	AaveLPAPTransactor // Write-only binding to the contract
	AaveLPAPFilterer   // Log filterer for contract events
}

// AaveLPAPCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveLPAPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveLPAPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveLPAPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveLPAPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveLPAPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveLPAPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveLPAPSession struct {
	Contract     *AaveLPAP         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveLPAPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveLPAPCallerSession struct {
	Contract *AaveLPAPCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AaveLPAPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveLPAPTransactorSession struct {
	Contract     *AaveLPAPTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AaveLPAPRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveLPAPRaw struct {
	Contract *AaveLPAP // Generic contract binding to access the raw methods on
}

// AaveLPAPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveLPAPCallerRaw struct {
	Contract *AaveLPAPCaller // Generic read-only contract binding to access the raw methods on
}

// AaveLPAPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveLPAPTransactorRaw struct {
	Contract *AaveLPAPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveLPAP creates a new instance of AaveLPAP, bound to a specific deployed contract.
func NewAaveLPAP(address common.Address, backend bind.ContractBackend) (*AaveLPAP, error) {
	contract, err := bindAaveLPAP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveLPAP{AaveLPAPCaller: AaveLPAPCaller{contract: contract}, AaveLPAPTransactor: AaveLPAPTransactor{contract: contract}, AaveLPAPFilterer: AaveLPAPFilterer{contract: contract}}, nil
}

// NewAaveLPAPCaller creates a new read-only instance of AaveLPAP, bound to a specific deployed contract.
func NewAaveLPAPCaller(address common.Address, caller bind.ContractCaller) (*AaveLPAPCaller, error) {
	contract, err := bindAaveLPAP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveLPAPCaller{contract: contract}, nil
}

// NewAaveLPAPTransactor creates a new write-only instance of AaveLPAP, bound to a specific deployed contract.
func NewAaveLPAPTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveLPAPTransactor, error) {
	contract, err := bindAaveLPAP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveLPAPTransactor{contract: contract}, nil
}

// NewAaveLPAPFilterer creates a new log filterer instance of AaveLPAP, bound to a specific deployed contract.
func NewAaveLPAPFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveLPAPFilterer, error) {
	contract, err := bindAaveLPAP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveLPAPFilterer{contract: contract}, nil
}

// bindAaveLPAP binds a generic wrapper to an already deployed contract.
func bindAaveLPAP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveLPAPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveLPAP *AaveLPAPRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveLPAP.Contract.AaveLPAPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveLPAP *AaveLPAPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveLPAP.Contract.AaveLPAPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveLPAP *AaveLPAPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveLPAP.Contract.AaveLPAPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveLPAP *AaveLPAPCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveLPAP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveLPAP *AaveLPAPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveLPAP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveLPAP *AaveLPAPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveLPAP.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _key) view returns(address)
func (_AaveLPAP *AaveLPAPCaller) GetAddress(opts *bind.CallOpts, _key [32]byte) (common.Address, error) {
	var out []interface{}
	err := _AaveLPAP.contract.Call(opts, &out, "getAddress", _key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _key) view returns(address)
func (_AaveLPAP *AaveLPAPSession) GetAddress(_key [32]byte) (common.Address, error) {
	return _AaveLPAP.Contract.GetAddress(&_AaveLPAP.CallOpts, _key)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _key) view returns(address)
func (_AaveLPAP *AaveLPAPCallerSession) GetAddress(_key [32]byte) (common.Address, error) {
	return _AaveLPAP.Contract.GetAddress(&_AaveLPAP.CallOpts, _key)
}

// GetFeeProvider is a free data retrieval call binding the contract method 0xfbeefc3c.
//
// Solidity: function getFeeProvider() view returns(address)
func (_AaveLPAP *AaveLPAPCaller) GetFeeProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveLPAP.contract.Call(opts, &out, "getFeeProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeProvider is a free data retrieval call binding the contract method 0xfbeefc3c.
//
// Solidity: function getFeeProvider() view returns(address)
func (_AaveLPAP *AaveLPAPSession) GetFeeProvider() (common.Address, error) {
	return _AaveLPAP.Contract.GetFeeProvider(&_AaveLPAP.CallOpts)
}

// GetFeeProvider is a free data retrieval call binding the contract method 0xfbeefc3c.
//
// Solidity: function getFeeProvider() view returns(address)
func (_AaveLPAP *AaveLPAPCallerSession) GetFeeProvider() (common.Address, error) {
	return _AaveLPAP.Contract.GetFeeProvider(&_AaveLPAP.CallOpts)
}

// GetLendingPool is a free data retrieval call binding the contract method 0x0261bf8b.
//
// Solidity: function getLendingPool() view returns(address)
func (_AaveLPAP *AaveLPAPCaller) GetLendingPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveLPAP.contract.Call(opts, &out, "getLendingPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPool is a free data retrieval call binding the contract method 0x0261bf8b.
//
// Solidity: function getLendingPool() view returns(address)
func (_AaveLPAP *AaveLPAPSession) GetLendingPool() (common.Address, error) {
	return _AaveLPAP.Contract.GetLendingPool(&_AaveLPAP.CallOpts)
}

// GetLendingPool is a free data retrieval call binding the contract method 0x0261bf8b.
//
// Solidity: function getLendingPool() view returns(address)
func (_AaveLPAP *AaveLPAPCallerSession) GetLendingPool() (common.Address, error) {
	return _AaveLPAP.Contract.GetLendingPool(&_AaveLPAP.CallOpts)
}

// GetLendingPoolConfigurator is a free data retrieval call binding the contract method 0x85c858b1.
//
// Solidity: function getLendingPoolConfigurator() view returns(address)
func (_AaveLPAP *AaveLPAPCaller) GetLendingPoolConfigurator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveLPAP.contract.Call(opts, &out, "getLendingPoolConfigurator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPoolConfigurator is a free data retrieval call binding the contract method 0x85c858b1.
//
// Solidity: function getLendingPoolConfigurator() view returns(address)
func (_AaveLPAP *AaveLPAPSession) GetLendingPoolConfigurator() (common.Address, error) {
	return _AaveLPAP.Contract.GetLendingPoolConfigurator(&_AaveLPAP.CallOpts)
}

// GetLendingPoolConfigurator is a free data retrieval call binding the contract method 0x85c858b1.
//
// Solidity: function getLendingPoolConfigurator() view returns(address)
func (_AaveLPAP *AaveLPAPCallerSession) GetLendingPoolConfigurator() (common.Address, error) {
	return _AaveLPAP.Contract.GetLendingPoolConfigurator(&_AaveLPAP.CallOpts)
}

// GetLendingPoolCore is a free data retrieval call binding the contract method 0xed6ff760.
//
// Solidity: function getLendingPoolCore() view returns(address)
func (_AaveLPAP *AaveLPAPCaller) GetLendingPoolCore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveLPAP.contract.Call(opts, &out, "getLendingPoolCore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPoolCore is a free data retrieval call binding the contract method 0xed6ff760.
//
// Solidity: function getLendingPoolCore() view returns(address)
func (_AaveLPAP *AaveLPAPSession) GetLendingPoolCore() (common.Address, error) {
	return _AaveLPAP.Contract.GetLendingPoolCore(&_AaveLPAP.CallOpts)
}

// GetLendingPoolCore is a free data retrieval call binding the contract method 0xed6ff760.
//
// Solidity: function getLendingPoolCore() view returns(address)
func (_AaveLPAP *AaveLPAPCallerSession) GetLendingPoolCore() (common.Address, error) {
	return _AaveLPAP.Contract.GetLendingPoolCore(&_AaveLPAP.CallOpts)
}

// GetLendingPoolDataProvider is a free data retrieval call binding the contract method 0x2f58b80d.
//
// Solidity: function getLendingPoolDataProvider() view returns(address)
func (_AaveLPAP *AaveLPAPCaller) GetLendingPoolDataProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveLPAP.contract.Call(opts, &out, "getLendingPoolDataProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPoolDataProvider is a free data retrieval call binding the contract method 0x2f58b80d.
//
// Solidity: function getLendingPoolDataProvider() view returns(address)
func (_AaveLPAP *AaveLPAPSession) GetLendingPoolDataProvider() (common.Address, error) {
	return _AaveLPAP.Contract.GetLendingPoolDataProvider(&_AaveLPAP.CallOpts)
}

// GetLendingPoolDataProvider is a free data retrieval call binding the contract method 0x2f58b80d.
//
// Solidity: function getLendingPoolDataProvider() view returns(address)
func (_AaveLPAP *AaveLPAPCallerSession) GetLendingPoolDataProvider() (common.Address, error) {
	return _AaveLPAP.Contract.GetLendingPoolDataProvider(&_AaveLPAP.CallOpts)
}

// GetLendingPoolLiquidationManager is a free data retrieval call binding the contract method 0x5834eb9a.
//
// Solidity: function getLendingPoolLiquidationManager() view returns(address)
func (_AaveLPAP *AaveLPAPCaller) GetLendingPoolLiquidationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveLPAP.contract.Call(opts, &out, "getLendingPoolLiquidationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPoolLiquidationManager is a free data retrieval call binding the contract method 0x5834eb9a.
//
// Solidity: function getLendingPoolLiquidationManager() view returns(address)
func (_AaveLPAP *AaveLPAPSession) GetLendingPoolLiquidationManager() (common.Address, error) {
	return _AaveLPAP.Contract.GetLendingPoolLiquidationManager(&_AaveLPAP.CallOpts)
}

// GetLendingPoolLiquidationManager is a free data retrieval call binding the contract method 0x5834eb9a.
//
// Solidity: function getLendingPoolLiquidationManager() view returns(address)
func (_AaveLPAP *AaveLPAPCallerSession) GetLendingPoolLiquidationManager() (common.Address, error) {
	return _AaveLPAP.Contract.GetLendingPoolLiquidationManager(&_AaveLPAP.CallOpts)
}

// GetLendingPoolManager is a free data retrieval call binding the contract method 0x33128d59.
//
// Solidity: function getLendingPoolManager() view returns(address)
func (_AaveLPAP *AaveLPAPCaller) GetLendingPoolManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveLPAP.contract.Call(opts, &out, "getLendingPoolManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPoolManager is a free data retrieval call binding the contract method 0x33128d59.
//
// Solidity: function getLendingPoolManager() view returns(address)
func (_AaveLPAP *AaveLPAPSession) GetLendingPoolManager() (common.Address, error) {
	return _AaveLPAP.Contract.GetLendingPoolManager(&_AaveLPAP.CallOpts)
}

// GetLendingPoolManager is a free data retrieval call binding the contract method 0x33128d59.
//
// Solidity: function getLendingPoolManager() view returns(address)
func (_AaveLPAP *AaveLPAPCallerSession) GetLendingPoolManager() (common.Address, error) {
	return _AaveLPAP.Contract.GetLendingPoolManager(&_AaveLPAP.CallOpts)
}

// GetLendingPoolParametersProvider is a free data retrieval call binding the contract method 0x04061d8e.
//
// Solidity: function getLendingPoolParametersProvider() view returns(address)
func (_AaveLPAP *AaveLPAPCaller) GetLendingPoolParametersProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveLPAP.contract.Call(opts, &out, "getLendingPoolParametersProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPoolParametersProvider is a free data retrieval call binding the contract method 0x04061d8e.
//
// Solidity: function getLendingPoolParametersProvider() view returns(address)
func (_AaveLPAP *AaveLPAPSession) GetLendingPoolParametersProvider() (common.Address, error) {
	return _AaveLPAP.Contract.GetLendingPoolParametersProvider(&_AaveLPAP.CallOpts)
}

// GetLendingPoolParametersProvider is a free data retrieval call binding the contract method 0x04061d8e.
//
// Solidity: function getLendingPoolParametersProvider() view returns(address)
func (_AaveLPAP *AaveLPAPCallerSession) GetLendingPoolParametersProvider() (common.Address, error) {
	return _AaveLPAP.Contract.GetLendingPoolParametersProvider(&_AaveLPAP.CallOpts)
}

// GetLendingRateOracle is a free data retrieval call binding the contract method 0x3618abba.
//
// Solidity: function getLendingRateOracle() view returns(address)
func (_AaveLPAP *AaveLPAPCaller) GetLendingRateOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveLPAP.contract.Call(opts, &out, "getLendingRateOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingRateOracle is a free data retrieval call binding the contract method 0x3618abba.
//
// Solidity: function getLendingRateOracle() view returns(address)
func (_AaveLPAP *AaveLPAPSession) GetLendingRateOracle() (common.Address, error) {
	return _AaveLPAP.Contract.GetLendingRateOracle(&_AaveLPAP.CallOpts)
}

// GetLendingRateOracle is a free data retrieval call binding the contract method 0x3618abba.
//
// Solidity: function getLendingRateOracle() view returns(address)
func (_AaveLPAP *AaveLPAPCallerSession) GetLendingRateOracle() (common.Address, error) {
	return _AaveLPAP.Contract.GetLendingRateOracle(&_AaveLPAP.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_AaveLPAP *AaveLPAPCaller) GetPriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveLPAP.contract.Call(opts, &out, "getPriceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_AaveLPAP *AaveLPAPSession) GetPriceOracle() (common.Address, error) {
	return _AaveLPAP.Contract.GetPriceOracle(&_AaveLPAP.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_AaveLPAP *AaveLPAPCallerSession) GetPriceOracle() (common.Address, error) {
	return _AaveLPAP.Contract.GetPriceOracle(&_AaveLPAP.CallOpts)
}

// GetTokenDistributor is a free data retrieval call binding the contract method 0xee891296.
//
// Solidity: function getTokenDistributor() view returns(address)
func (_AaveLPAP *AaveLPAPCaller) GetTokenDistributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveLPAP.contract.Call(opts, &out, "getTokenDistributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenDistributor is a free data retrieval call binding the contract method 0xee891296.
//
// Solidity: function getTokenDistributor() view returns(address)
func (_AaveLPAP *AaveLPAPSession) GetTokenDistributor() (common.Address, error) {
	return _AaveLPAP.Contract.GetTokenDistributor(&_AaveLPAP.CallOpts)
}

// GetTokenDistributor is a free data retrieval call binding the contract method 0xee891296.
//
// Solidity: function getTokenDistributor() view returns(address)
func (_AaveLPAP *AaveLPAPCallerSession) GetTokenDistributor() (common.Address, error) {
	return _AaveLPAP.Contract.GetTokenDistributor(&_AaveLPAP.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_AaveLPAP *AaveLPAPCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AaveLPAP.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_AaveLPAP *AaveLPAPSession) IsOwner() (bool, error) {
	return _AaveLPAP.Contract.IsOwner(&_AaveLPAP.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_AaveLPAP *AaveLPAPCallerSession) IsOwner() (bool, error) {
	return _AaveLPAP.Contract.IsOwner(&_AaveLPAP.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AaveLPAP *AaveLPAPCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveLPAP.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AaveLPAP *AaveLPAPSession) Owner() (common.Address, error) {
	return _AaveLPAP.Contract.Owner(&_AaveLPAP.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AaveLPAP *AaveLPAPCallerSession) Owner() (common.Address, error) {
	return _AaveLPAP.Contract.Owner(&_AaveLPAP.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AaveLPAP *AaveLPAPTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveLPAP.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AaveLPAP *AaveLPAPSession) RenounceOwnership() (*types.Transaction, error) {
	return _AaveLPAP.Contract.RenounceOwnership(&_AaveLPAP.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AaveLPAP *AaveLPAPTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AaveLPAP.Contract.RenounceOwnership(&_AaveLPAP.TransactOpts)
}

// SetFeeProviderImpl is a paid mutator transaction binding the contract method 0x2a62c636.
//
// Solidity: function setFeeProviderImpl(address _feeProvider) returns()
func (_AaveLPAP *AaveLPAPTransactor) SetFeeProviderImpl(opts *bind.TransactOpts, _feeProvider common.Address) (*types.Transaction, error) {
	return _AaveLPAP.contract.Transact(opts, "setFeeProviderImpl", _feeProvider)
}

// SetFeeProviderImpl is a paid mutator transaction binding the contract method 0x2a62c636.
//
// Solidity: function setFeeProviderImpl(address _feeProvider) returns()
func (_AaveLPAP *AaveLPAPSession) SetFeeProviderImpl(_feeProvider common.Address) (*types.Transaction, error) {
	return _AaveLPAP.Contract.SetFeeProviderImpl(&_AaveLPAP.TransactOpts, _feeProvider)
}

// SetFeeProviderImpl is a paid mutator transaction binding the contract method 0x2a62c636.
//
// Solidity: function setFeeProviderImpl(address _feeProvider) returns()
func (_AaveLPAP *AaveLPAPTransactorSession) SetFeeProviderImpl(_feeProvider common.Address) (*types.Transaction, error) {
	return _AaveLPAP.Contract.SetFeeProviderImpl(&_AaveLPAP.TransactOpts, _feeProvider)
}

// SetLendingPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xc12542df.
//
// Solidity: function setLendingPoolConfiguratorImpl(address _configurator) returns()
func (_AaveLPAP *AaveLPAPTransactor) SetLendingPoolConfiguratorImpl(opts *bind.TransactOpts, _configurator common.Address) (*types.Transaction, error) {
	return _AaveLPAP.contract.Transact(opts, "setLendingPoolConfiguratorImpl", _configurator)
}

// SetLendingPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xc12542df.
//
// Solidity: function setLendingPoolConfiguratorImpl(address _configurator) returns()
func (_AaveLPAP *AaveLPAPSession) SetLendingPoolConfiguratorImpl(_configurator common.Address) (*types.Transaction, error) {
	return _AaveLPAP.Contract.SetLendingPoolConfiguratorImpl(&_AaveLPAP.TransactOpts, _configurator)
}

// SetLendingPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xc12542df.
//
// Solidity: function setLendingPoolConfiguratorImpl(address _configurator) returns()
func (_AaveLPAP *AaveLPAPTransactorSession) SetLendingPoolConfiguratorImpl(_configurator common.Address) (*types.Transaction, error) {
	return _AaveLPAP.Contract.SetLendingPoolConfiguratorImpl(&_AaveLPAP.TransactOpts, _configurator)
}

// SetLendingPoolCoreImpl is a paid mutator transaction binding the contract method 0x1c827204.
//
// Solidity: function setLendingPoolCoreImpl(address _lendingPoolCore) returns()
func (_AaveLPAP *AaveLPAPTransactor) SetLendingPoolCoreImpl(opts *bind.TransactOpts, _lendingPoolCore common.Address) (*types.Transaction, error) {
	return _AaveLPAP.contract.Transact(opts, "setLendingPoolCoreImpl", _lendingPoolCore)
}

// SetLendingPoolCoreImpl is a paid mutator transaction binding the contract method 0x1c827204.
//
// Solidity: function setLendingPoolCoreImpl(address _lendingPoolCore) returns()
func (_AaveLPAP *AaveLPAPSession) SetLendingPoolCoreImpl(_lendingPoolCore common.Address) (*types.Transaction, error) {
	return _AaveLPAP.Contract.SetLendingPoolCoreImpl(&_AaveLPAP.TransactOpts, _lendingPoolCore)
}

// SetLendingPoolCoreImpl is a paid mutator transaction binding the contract method 0x1c827204.
//
// Solidity: function setLendingPoolCoreImpl(address _lendingPoolCore) returns()
func (_AaveLPAP *AaveLPAPTransactorSession) SetLendingPoolCoreImpl(_lendingPoolCore common.Address) (*types.Transaction, error) {
	return _AaveLPAP.Contract.SetLendingPoolCoreImpl(&_AaveLPAP.TransactOpts, _lendingPoolCore)
}

// SetLendingPoolDataProviderImpl is a paid mutator transaction binding the contract method 0xbfedc103.
//
// Solidity: function setLendingPoolDataProviderImpl(address _provider) returns()
func (_AaveLPAP *AaveLPAPTransactor) SetLendingPoolDataProviderImpl(opts *bind.TransactOpts, _provider common.Address) (*types.Transaction, error) {
	return _AaveLPAP.contract.Transact(opts, "setLendingPoolDataProviderImpl", _provider)
}

// SetLendingPoolDataProviderImpl is a paid mutator transaction binding the contract method 0xbfedc103.
//
// Solidity: function setLendingPoolDataProviderImpl(address _provider) returns()
func (_AaveLPAP *AaveLPAPSession) SetLendingPoolDataProviderImpl(_provider common.Address) (*types.Transaction, error) {
	return _AaveLPAP.Contract.SetLendingPoolDataProviderImpl(&_AaveLPAP.TransactOpts, _provider)
}

// SetLendingPoolDataProviderImpl is a paid mutator transaction binding the contract method 0xbfedc103.
//
// Solidity: function setLendingPoolDataProviderImpl(address _provider) returns()
func (_AaveLPAP *AaveLPAPTransactorSession) SetLendingPoolDataProviderImpl(_provider common.Address) (*types.Transaction, error) {
	return _AaveLPAP.Contract.SetLendingPoolDataProviderImpl(&_AaveLPAP.TransactOpts, _provider)
}

// SetLendingPoolImpl is a paid mutator transaction binding the contract method 0x5aef021f.
//
// Solidity: function setLendingPoolImpl(address _pool) returns()
func (_AaveLPAP *AaveLPAPTransactor) SetLendingPoolImpl(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _AaveLPAP.contract.Transact(opts, "setLendingPoolImpl", _pool)
}

// SetLendingPoolImpl is a paid mutator transaction binding the contract method 0x5aef021f.
//
// Solidity: function setLendingPoolImpl(address _pool) returns()
func (_AaveLPAP *AaveLPAPSession) SetLendingPoolImpl(_pool common.Address) (*types.Transaction, error) {
	return _AaveLPAP.Contract.SetLendingPoolImpl(&_AaveLPAP.TransactOpts, _pool)
}

// SetLendingPoolImpl is a paid mutator transaction binding the contract method 0x5aef021f.
//
// Solidity: function setLendingPoolImpl(address _pool) returns()
func (_AaveLPAP *AaveLPAPTransactorSession) SetLendingPoolImpl(_pool common.Address) (*types.Transaction, error) {
	return _AaveLPAP.Contract.SetLendingPoolImpl(&_AaveLPAP.TransactOpts, _pool)
}

// SetLendingPoolLiquidationManager is a paid mutator transaction binding the contract method 0x44ce375b.
//
// Solidity: function setLendingPoolLiquidationManager(address _manager) returns()
func (_AaveLPAP *AaveLPAPTransactor) SetLendingPoolLiquidationManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _AaveLPAP.contract.Transact(opts, "setLendingPoolLiquidationManager", _manager)
}

// SetLendingPoolLiquidationManager is a paid mutator transaction binding the contract method 0x44ce375b.
//
// Solidity: function setLendingPoolLiquidationManager(address _manager) returns()
func (_AaveLPAP *AaveLPAPSession) SetLendingPoolLiquidationManager(_manager common.Address) (*types.Transaction, error) {
	return _AaveLPAP.Contract.SetLendingPoolLiquidationManager(&_AaveLPAP.TransactOpts, _manager)
}

// SetLendingPoolLiquidationManager is a paid mutator transaction binding the contract method 0x44ce375b.
//
// Solidity: function setLendingPoolLiquidationManager(address _manager) returns()
func (_AaveLPAP *AaveLPAPTransactorSession) SetLendingPoolLiquidationManager(_manager common.Address) (*types.Transaction, error) {
	return _AaveLPAP.Contract.SetLendingPoolLiquidationManager(&_AaveLPAP.TransactOpts, _manager)
}

// SetLendingPoolManager is a paid mutator transaction binding the contract method 0x40fdcadc.
//
// Solidity: function setLendingPoolManager(address _lendingPoolManager) returns()
func (_AaveLPAP *AaveLPAPTransactor) SetLendingPoolManager(opts *bind.TransactOpts, _lendingPoolManager common.Address) (*types.Transaction, error) {
	return _AaveLPAP.contract.Transact(opts, "setLendingPoolManager", _lendingPoolManager)
}

// SetLendingPoolManager is a paid mutator transaction binding the contract method 0x40fdcadc.
//
// Solidity: function setLendingPoolManager(address _lendingPoolManager) returns()
func (_AaveLPAP *AaveLPAPSession) SetLendingPoolManager(_lendingPoolManager common.Address) (*types.Transaction, error) {
	return _AaveLPAP.Contract.SetLendingPoolManager(&_AaveLPAP.TransactOpts, _lendingPoolManager)
}

// SetLendingPoolManager is a paid mutator transaction binding the contract method 0x40fdcadc.
//
// Solidity: function setLendingPoolManager(address _lendingPoolManager) returns()
func (_AaveLPAP *AaveLPAPTransactorSession) SetLendingPoolManager(_lendingPoolManager common.Address) (*types.Transaction, error) {
	return _AaveLPAP.Contract.SetLendingPoolManager(&_AaveLPAP.TransactOpts, _lendingPoolManager)
}

// SetLendingPoolParametersProviderImpl is a paid mutator transaction binding the contract method 0xa5eface2.
//
// Solidity: function setLendingPoolParametersProviderImpl(address _parametersProvider) returns()
func (_AaveLPAP *AaveLPAPTransactor) SetLendingPoolParametersProviderImpl(opts *bind.TransactOpts, _parametersProvider common.Address) (*types.Transaction, error) {
	return _AaveLPAP.contract.Transact(opts, "setLendingPoolParametersProviderImpl", _parametersProvider)
}

// SetLendingPoolParametersProviderImpl is a paid mutator transaction binding the contract method 0xa5eface2.
//
// Solidity: function setLendingPoolParametersProviderImpl(address _parametersProvider) returns()
func (_AaveLPAP *AaveLPAPSession) SetLendingPoolParametersProviderImpl(_parametersProvider common.Address) (*types.Transaction, error) {
	return _AaveLPAP.Contract.SetLendingPoolParametersProviderImpl(&_AaveLPAP.TransactOpts, _parametersProvider)
}

// SetLendingPoolParametersProviderImpl is a paid mutator transaction binding the contract method 0xa5eface2.
//
// Solidity: function setLendingPoolParametersProviderImpl(address _parametersProvider) returns()
func (_AaveLPAP *AaveLPAPTransactorSession) SetLendingPoolParametersProviderImpl(_parametersProvider common.Address) (*types.Transaction, error) {
	return _AaveLPAP.Contract.SetLendingPoolParametersProviderImpl(&_AaveLPAP.TransactOpts, _parametersProvider)
}

// SetLendingRateOracle is a paid mutator transaction binding the contract method 0x820d1274.
//
// Solidity: function setLendingRateOracle(address _lendingRateOracle) returns()
func (_AaveLPAP *AaveLPAPTransactor) SetLendingRateOracle(opts *bind.TransactOpts, _lendingRateOracle common.Address) (*types.Transaction, error) {
	return _AaveLPAP.contract.Transact(opts, "setLendingRateOracle", _lendingRateOracle)
}

// SetLendingRateOracle is a paid mutator transaction binding the contract method 0x820d1274.
//
// Solidity: function setLendingRateOracle(address _lendingRateOracle) returns()
func (_AaveLPAP *AaveLPAPSession) SetLendingRateOracle(_lendingRateOracle common.Address) (*types.Transaction, error) {
	return _AaveLPAP.Contract.SetLendingRateOracle(&_AaveLPAP.TransactOpts, _lendingRateOracle)
}

// SetLendingRateOracle is a paid mutator transaction binding the contract method 0x820d1274.
//
// Solidity: function setLendingRateOracle(address _lendingRateOracle) returns()
func (_AaveLPAP *AaveLPAPTransactorSession) SetLendingRateOracle(_lendingRateOracle common.Address) (*types.Transaction, error) {
	return _AaveLPAP.Contract.SetLendingRateOracle(&_AaveLPAP.TransactOpts, _lendingRateOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _priceOracle) returns()
func (_AaveLPAP *AaveLPAPTransactor) SetPriceOracle(opts *bind.TransactOpts, _priceOracle common.Address) (*types.Transaction, error) {
	return _AaveLPAP.contract.Transact(opts, "setPriceOracle", _priceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _priceOracle) returns()
func (_AaveLPAP *AaveLPAPSession) SetPriceOracle(_priceOracle common.Address) (*types.Transaction, error) {
	return _AaveLPAP.Contract.SetPriceOracle(&_AaveLPAP.TransactOpts, _priceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _priceOracle) returns()
func (_AaveLPAP *AaveLPAPTransactorSession) SetPriceOracle(_priceOracle common.Address) (*types.Transaction, error) {
	return _AaveLPAP.Contract.SetPriceOracle(&_AaveLPAP.TransactOpts, _priceOracle)
}

// SetTokenDistributor is a paid mutator transaction binding the contract method 0x38280e6b.
//
// Solidity: function setTokenDistributor(address _tokenDistributor) returns()
func (_AaveLPAP *AaveLPAPTransactor) SetTokenDistributor(opts *bind.TransactOpts, _tokenDistributor common.Address) (*types.Transaction, error) {
	return _AaveLPAP.contract.Transact(opts, "setTokenDistributor", _tokenDistributor)
}

// SetTokenDistributor is a paid mutator transaction binding the contract method 0x38280e6b.
//
// Solidity: function setTokenDistributor(address _tokenDistributor) returns()
func (_AaveLPAP *AaveLPAPSession) SetTokenDistributor(_tokenDistributor common.Address) (*types.Transaction, error) {
	return _AaveLPAP.Contract.SetTokenDistributor(&_AaveLPAP.TransactOpts, _tokenDistributor)
}

// SetTokenDistributor is a paid mutator transaction binding the contract method 0x38280e6b.
//
// Solidity: function setTokenDistributor(address _tokenDistributor) returns()
func (_AaveLPAP *AaveLPAPTransactorSession) SetTokenDistributor(_tokenDistributor common.Address) (*types.Transaction, error) {
	return _AaveLPAP.Contract.SetTokenDistributor(&_AaveLPAP.TransactOpts, _tokenDistributor)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AaveLPAP *AaveLPAPTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AaveLPAP.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AaveLPAP *AaveLPAPSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AaveLPAP.Contract.TransferOwnership(&_AaveLPAP.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AaveLPAP *AaveLPAPTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AaveLPAP.Contract.TransferOwnership(&_AaveLPAP.TransactOpts, newOwner)
}

// AaveLPAPEthereumAddressUpdatedIterator is returned from FilterEthereumAddressUpdated and is used to iterate over the raw logs and unpacked data for EthereumAddressUpdated events raised by the AaveLPAP contract.
type AaveLPAPEthereumAddressUpdatedIterator struct {
	Event *AaveLPAPEthereumAddressUpdated // Event containing the contract specifics and raw log

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
func (it *AaveLPAPEthereumAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveLPAPEthereumAddressUpdated)
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
		it.Event = new(AaveLPAPEthereumAddressUpdated)
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
func (it *AaveLPAPEthereumAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveLPAPEthereumAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveLPAPEthereumAddressUpdated represents a EthereumAddressUpdated event raised by the AaveLPAP contract.
type AaveLPAPEthereumAddressUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEthereumAddressUpdated is a free log retrieval operation binding the contract event 0x1941c4aa9ef07ed299ed253b4224020832574525e8db1f3f955c57f395aef829.
//
// Solidity: event EthereumAddressUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) FilterEthereumAddressUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AaveLPAPEthereumAddressUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.FilterLogs(opts, "EthereumAddressUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveLPAPEthereumAddressUpdatedIterator{contract: _AaveLPAP.contract, event: "EthereumAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchEthereumAddressUpdated is a free log subscription operation binding the contract event 0x1941c4aa9ef07ed299ed253b4224020832574525e8db1f3f955c57f395aef829.
//
// Solidity: event EthereumAddressUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) WatchEthereumAddressUpdated(opts *bind.WatchOpts, sink chan<- *AaveLPAPEthereumAddressUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.WatchLogs(opts, "EthereumAddressUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveLPAPEthereumAddressUpdated)
				if err := _AaveLPAP.contract.UnpackLog(event, "EthereumAddressUpdated", log); err != nil {
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

// ParseEthereumAddressUpdated is a log parse operation binding the contract event 0x1941c4aa9ef07ed299ed253b4224020832574525e8db1f3f955c57f395aef829.
//
// Solidity: event EthereumAddressUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) ParseEthereumAddressUpdated(log types.Log) (*AaveLPAPEthereumAddressUpdated, error) {
	event := new(AaveLPAPEthereumAddressUpdated)
	if err := _AaveLPAP.contract.UnpackLog(event, "EthereumAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveLPAPFeeProviderUpdatedIterator is returned from FilterFeeProviderUpdated and is used to iterate over the raw logs and unpacked data for FeeProviderUpdated events raised by the AaveLPAP contract.
type AaveLPAPFeeProviderUpdatedIterator struct {
	Event *AaveLPAPFeeProviderUpdated // Event containing the contract specifics and raw log

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
func (it *AaveLPAPFeeProviderUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveLPAPFeeProviderUpdated)
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
		it.Event = new(AaveLPAPFeeProviderUpdated)
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
func (it *AaveLPAPFeeProviderUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveLPAPFeeProviderUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveLPAPFeeProviderUpdated represents a FeeProviderUpdated event raised by the AaveLPAP contract.
type AaveLPAPFeeProviderUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeProviderUpdated is a free log retrieval operation binding the contract event 0x18e1a55b8eff9c93921eecfa1462d6a8cbb80b3988db3eb14c039e43fdb22661.
//
// Solidity: event FeeProviderUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) FilterFeeProviderUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AaveLPAPFeeProviderUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.FilterLogs(opts, "FeeProviderUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveLPAPFeeProviderUpdatedIterator{contract: _AaveLPAP.contract, event: "FeeProviderUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeProviderUpdated is a free log subscription operation binding the contract event 0x18e1a55b8eff9c93921eecfa1462d6a8cbb80b3988db3eb14c039e43fdb22661.
//
// Solidity: event FeeProviderUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) WatchFeeProviderUpdated(opts *bind.WatchOpts, sink chan<- *AaveLPAPFeeProviderUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.WatchLogs(opts, "FeeProviderUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveLPAPFeeProviderUpdated)
				if err := _AaveLPAP.contract.UnpackLog(event, "FeeProviderUpdated", log); err != nil {
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

// ParseFeeProviderUpdated is a log parse operation binding the contract event 0x18e1a55b8eff9c93921eecfa1462d6a8cbb80b3988db3eb14c039e43fdb22661.
//
// Solidity: event FeeProviderUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) ParseFeeProviderUpdated(log types.Log) (*AaveLPAPFeeProviderUpdated, error) {
	event := new(AaveLPAPFeeProviderUpdated)
	if err := _AaveLPAP.contract.UnpackLog(event, "FeeProviderUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveLPAPLendingPoolConfiguratorUpdatedIterator is returned from FilterLendingPoolConfiguratorUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolConfiguratorUpdated events raised by the AaveLPAP contract.
type AaveLPAPLendingPoolConfiguratorUpdatedIterator struct {
	Event *AaveLPAPLendingPoolConfiguratorUpdated // Event containing the contract specifics and raw log

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
func (it *AaveLPAPLendingPoolConfiguratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveLPAPLendingPoolConfiguratorUpdated)
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
		it.Event = new(AaveLPAPLendingPoolConfiguratorUpdated)
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
func (it *AaveLPAPLendingPoolConfiguratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveLPAPLendingPoolConfiguratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveLPAPLendingPoolConfiguratorUpdated represents a LendingPoolConfiguratorUpdated event raised by the AaveLPAP contract.
type AaveLPAPLendingPoolConfiguratorUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolConfiguratorUpdated is a free log retrieval operation binding the contract event 0xdfabe479bad36782fb1e77fbfddd4e382671713527e4786cfc93a022ae763729.
//
// Solidity: event LendingPoolConfiguratorUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) FilterLendingPoolConfiguratorUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AaveLPAPLendingPoolConfiguratorUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.FilterLogs(opts, "LendingPoolConfiguratorUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveLPAPLendingPoolConfiguratorUpdatedIterator{contract: _AaveLPAP.contract, event: "LendingPoolConfiguratorUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolConfiguratorUpdated is a free log subscription operation binding the contract event 0xdfabe479bad36782fb1e77fbfddd4e382671713527e4786cfc93a022ae763729.
//
// Solidity: event LendingPoolConfiguratorUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) WatchLendingPoolConfiguratorUpdated(opts *bind.WatchOpts, sink chan<- *AaveLPAPLendingPoolConfiguratorUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.WatchLogs(opts, "LendingPoolConfiguratorUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveLPAPLendingPoolConfiguratorUpdated)
				if err := _AaveLPAP.contract.UnpackLog(event, "LendingPoolConfiguratorUpdated", log); err != nil {
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

// ParseLendingPoolConfiguratorUpdated is a log parse operation binding the contract event 0xdfabe479bad36782fb1e77fbfddd4e382671713527e4786cfc93a022ae763729.
//
// Solidity: event LendingPoolConfiguratorUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) ParseLendingPoolConfiguratorUpdated(log types.Log) (*AaveLPAPLendingPoolConfiguratorUpdated, error) {
	event := new(AaveLPAPLendingPoolConfiguratorUpdated)
	if err := _AaveLPAP.contract.UnpackLog(event, "LendingPoolConfiguratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveLPAPLendingPoolCoreUpdatedIterator is returned from FilterLendingPoolCoreUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolCoreUpdated events raised by the AaveLPAP contract.
type AaveLPAPLendingPoolCoreUpdatedIterator struct {
	Event *AaveLPAPLendingPoolCoreUpdated // Event containing the contract specifics and raw log

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
func (it *AaveLPAPLendingPoolCoreUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveLPAPLendingPoolCoreUpdated)
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
		it.Event = new(AaveLPAPLendingPoolCoreUpdated)
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
func (it *AaveLPAPLendingPoolCoreUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveLPAPLendingPoolCoreUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveLPAPLendingPoolCoreUpdated represents a LendingPoolCoreUpdated event raised by the AaveLPAP contract.
type AaveLPAPLendingPoolCoreUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolCoreUpdated is a free log retrieval operation binding the contract event 0x71c226bb2879778ca1298196bf7cc1256baee9d05b496c31ee627d35a471b5b7.
//
// Solidity: event LendingPoolCoreUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) FilterLendingPoolCoreUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AaveLPAPLendingPoolCoreUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.FilterLogs(opts, "LendingPoolCoreUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveLPAPLendingPoolCoreUpdatedIterator{contract: _AaveLPAP.contract, event: "LendingPoolCoreUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolCoreUpdated is a free log subscription operation binding the contract event 0x71c226bb2879778ca1298196bf7cc1256baee9d05b496c31ee627d35a471b5b7.
//
// Solidity: event LendingPoolCoreUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) WatchLendingPoolCoreUpdated(opts *bind.WatchOpts, sink chan<- *AaveLPAPLendingPoolCoreUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.WatchLogs(opts, "LendingPoolCoreUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveLPAPLendingPoolCoreUpdated)
				if err := _AaveLPAP.contract.UnpackLog(event, "LendingPoolCoreUpdated", log); err != nil {
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

// ParseLendingPoolCoreUpdated is a log parse operation binding the contract event 0x71c226bb2879778ca1298196bf7cc1256baee9d05b496c31ee627d35a471b5b7.
//
// Solidity: event LendingPoolCoreUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) ParseLendingPoolCoreUpdated(log types.Log) (*AaveLPAPLendingPoolCoreUpdated, error) {
	event := new(AaveLPAPLendingPoolCoreUpdated)
	if err := _AaveLPAP.contract.UnpackLog(event, "LendingPoolCoreUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveLPAPLendingPoolDataProviderUpdatedIterator is returned from FilterLendingPoolDataProviderUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolDataProviderUpdated events raised by the AaveLPAP contract.
type AaveLPAPLendingPoolDataProviderUpdatedIterator struct {
	Event *AaveLPAPLendingPoolDataProviderUpdated // Event containing the contract specifics and raw log

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
func (it *AaveLPAPLendingPoolDataProviderUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveLPAPLendingPoolDataProviderUpdated)
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
		it.Event = new(AaveLPAPLendingPoolDataProviderUpdated)
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
func (it *AaveLPAPLendingPoolDataProviderUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveLPAPLendingPoolDataProviderUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveLPAPLendingPoolDataProviderUpdated represents a LendingPoolDataProviderUpdated event raised by the AaveLPAP contract.
type AaveLPAPLendingPoolDataProviderUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolDataProviderUpdated is a free log retrieval operation binding the contract event 0x07890d7d10db37434d76ee4f2928fb2dc66227c3b3391206aced4f7bcb59cdb0.
//
// Solidity: event LendingPoolDataProviderUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) FilterLendingPoolDataProviderUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AaveLPAPLendingPoolDataProviderUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.FilterLogs(opts, "LendingPoolDataProviderUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveLPAPLendingPoolDataProviderUpdatedIterator{contract: _AaveLPAP.contract, event: "LendingPoolDataProviderUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolDataProviderUpdated is a free log subscription operation binding the contract event 0x07890d7d10db37434d76ee4f2928fb2dc66227c3b3391206aced4f7bcb59cdb0.
//
// Solidity: event LendingPoolDataProviderUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) WatchLendingPoolDataProviderUpdated(opts *bind.WatchOpts, sink chan<- *AaveLPAPLendingPoolDataProviderUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.WatchLogs(opts, "LendingPoolDataProviderUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveLPAPLendingPoolDataProviderUpdated)
				if err := _AaveLPAP.contract.UnpackLog(event, "LendingPoolDataProviderUpdated", log); err != nil {
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

// ParseLendingPoolDataProviderUpdated is a log parse operation binding the contract event 0x07890d7d10db37434d76ee4f2928fb2dc66227c3b3391206aced4f7bcb59cdb0.
//
// Solidity: event LendingPoolDataProviderUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) ParseLendingPoolDataProviderUpdated(log types.Log) (*AaveLPAPLendingPoolDataProviderUpdated, error) {
	event := new(AaveLPAPLendingPoolDataProviderUpdated)
	if err := _AaveLPAP.contract.UnpackLog(event, "LendingPoolDataProviderUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveLPAPLendingPoolLiquidationManagerUpdatedIterator is returned from FilterLendingPoolLiquidationManagerUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolLiquidationManagerUpdated events raised by the AaveLPAP contract.
type AaveLPAPLendingPoolLiquidationManagerUpdatedIterator struct {
	Event *AaveLPAPLendingPoolLiquidationManagerUpdated // Event containing the contract specifics and raw log

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
func (it *AaveLPAPLendingPoolLiquidationManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveLPAPLendingPoolLiquidationManagerUpdated)
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
		it.Event = new(AaveLPAPLendingPoolLiquidationManagerUpdated)
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
func (it *AaveLPAPLendingPoolLiquidationManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveLPAPLendingPoolLiquidationManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveLPAPLendingPoolLiquidationManagerUpdated represents a LendingPoolLiquidationManagerUpdated event raised by the AaveLPAP contract.
type AaveLPAPLendingPoolLiquidationManagerUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolLiquidationManagerUpdated is a free log retrieval operation binding the contract event 0x1a76cb769b814bc038223da86932b099b20aae03473683e6d98f5c3554e26064.
//
// Solidity: event LendingPoolLiquidationManagerUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) FilterLendingPoolLiquidationManagerUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AaveLPAPLendingPoolLiquidationManagerUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.FilterLogs(opts, "LendingPoolLiquidationManagerUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveLPAPLendingPoolLiquidationManagerUpdatedIterator{contract: _AaveLPAP.contract, event: "LendingPoolLiquidationManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolLiquidationManagerUpdated is a free log subscription operation binding the contract event 0x1a76cb769b814bc038223da86932b099b20aae03473683e6d98f5c3554e26064.
//
// Solidity: event LendingPoolLiquidationManagerUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) WatchLendingPoolLiquidationManagerUpdated(opts *bind.WatchOpts, sink chan<- *AaveLPAPLendingPoolLiquidationManagerUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.WatchLogs(opts, "LendingPoolLiquidationManagerUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveLPAPLendingPoolLiquidationManagerUpdated)
				if err := _AaveLPAP.contract.UnpackLog(event, "LendingPoolLiquidationManagerUpdated", log); err != nil {
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

// ParseLendingPoolLiquidationManagerUpdated is a log parse operation binding the contract event 0x1a76cb769b814bc038223da86932b099b20aae03473683e6d98f5c3554e26064.
//
// Solidity: event LendingPoolLiquidationManagerUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) ParseLendingPoolLiquidationManagerUpdated(log types.Log) (*AaveLPAPLendingPoolLiquidationManagerUpdated, error) {
	event := new(AaveLPAPLendingPoolLiquidationManagerUpdated)
	if err := _AaveLPAP.contract.UnpackLog(event, "LendingPoolLiquidationManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveLPAPLendingPoolManagerUpdatedIterator is returned from FilterLendingPoolManagerUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolManagerUpdated events raised by the AaveLPAP contract.
type AaveLPAPLendingPoolManagerUpdatedIterator struct {
	Event *AaveLPAPLendingPoolManagerUpdated // Event containing the contract specifics and raw log

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
func (it *AaveLPAPLendingPoolManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveLPAPLendingPoolManagerUpdated)
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
		it.Event = new(AaveLPAPLendingPoolManagerUpdated)
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
func (it *AaveLPAPLendingPoolManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveLPAPLendingPoolManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveLPAPLendingPoolManagerUpdated represents a LendingPoolManagerUpdated event raised by the AaveLPAP contract.
type AaveLPAPLendingPoolManagerUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolManagerUpdated is a free log retrieval operation binding the contract event 0xd5280c51185a38d36f7a0f5e56cac6248312bb70d0974782fa5a595127e0e08e.
//
// Solidity: event LendingPoolManagerUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) FilterLendingPoolManagerUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AaveLPAPLendingPoolManagerUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.FilterLogs(opts, "LendingPoolManagerUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveLPAPLendingPoolManagerUpdatedIterator{contract: _AaveLPAP.contract, event: "LendingPoolManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolManagerUpdated is a free log subscription operation binding the contract event 0xd5280c51185a38d36f7a0f5e56cac6248312bb70d0974782fa5a595127e0e08e.
//
// Solidity: event LendingPoolManagerUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) WatchLendingPoolManagerUpdated(opts *bind.WatchOpts, sink chan<- *AaveLPAPLendingPoolManagerUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.WatchLogs(opts, "LendingPoolManagerUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveLPAPLendingPoolManagerUpdated)
				if err := _AaveLPAP.contract.UnpackLog(event, "LendingPoolManagerUpdated", log); err != nil {
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

// ParseLendingPoolManagerUpdated is a log parse operation binding the contract event 0xd5280c51185a38d36f7a0f5e56cac6248312bb70d0974782fa5a595127e0e08e.
//
// Solidity: event LendingPoolManagerUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) ParseLendingPoolManagerUpdated(log types.Log) (*AaveLPAPLendingPoolManagerUpdated, error) {
	event := new(AaveLPAPLendingPoolManagerUpdated)
	if err := _AaveLPAP.contract.UnpackLog(event, "LendingPoolManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveLPAPLendingPoolParametersProviderUpdatedIterator is returned from FilterLendingPoolParametersProviderUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolParametersProviderUpdated events raised by the AaveLPAP contract.
type AaveLPAPLendingPoolParametersProviderUpdatedIterator struct {
	Event *AaveLPAPLendingPoolParametersProviderUpdated // Event containing the contract specifics and raw log

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
func (it *AaveLPAPLendingPoolParametersProviderUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveLPAPLendingPoolParametersProviderUpdated)
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
		it.Event = new(AaveLPAPLendingPoolParametersProviderUpdated)
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
func (it *AaveLPAPLendingPoolParametersProviderUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveLPAPLendingPoolParametersProviderUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveLPAPLendingPoolParametersProviderUpdated represents a LendingPoolParametersProviderUpdated event raised by the AaveLPAP contract.
type AaveLPAPLendingPoolParametersProviderUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolParametersProviderUpdated is a free log retrieval operation binding the contract event 0xce16ea9b2fd7cadddd0f7359971028f50b5ddba33dfae1a9bdea956fabb1e280.
//
// Solidity: event LendingPoolParametersProviderUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) FilterLendingPoolParametersProviderUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AaveLPAPLendingPoolParametersProviderUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.FilterLogs(opts, "LendingPoolParametersProviderUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveLPAPLendingPoolParametersProviderUpdatedIterator{contract: _AaveLPAP.contract, event: "LendingPoolParametersProviderUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolParametersProviderUpdated is a free log subscription operation binding the contract event 0xce16ea9b2fd7cadddd0f7359971028f50b5ddba33dfae1a9bdea956fabb1e280.
//
// Solidity: event LendingPoolParametersProviderUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) WatchLendingPoolParametersProviderUpdated(opts *bind.WatchOpts, sink chan<- *AaveLPAPLendingPoolParametersProviderUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.WatchLogs(opts, "LendingPoolParametersProviderUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveLPAPLendingPoolParametersProviderUpdated)
				if err := _AaveLPAP.contract.UnpackLog(event, "LendingPoolParametersProviderUpdated", log); err != nil {
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

// ParseLendingPoolParametersProviderUpdated is a log parse operation binding the contract event 0xce16ea9b2fd7cadddd0f7359971028f50b5ddba33dfae1a9bdea956fabb1e280.
//
// Solidity: event LendingPoolParametersProviderUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) ParseLendingPoolParametersProviderUpdated(log types.Log) (*AaveLPAPLendingPoolParametersProviderUpdated, error) {
	event := new(AaveLPAPLendingPoolParametersProviderUpdated)
	if err := _AaveLPAP.contract.UnpackLog(event, "LendingPoolParametersProviderUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveLPAPLendingPoolUpdatedIterator is returned from FilterLendingPoolUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolUpdated events raised by the AaveLPAP contract.
type AaveLPAPLendingPoolUpdatedIterator struct {
	Event *AaveLPAPLendingPoolUpdated // Event containing the contract specifics and raw log

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
func (it *AaveLPAPLendingPoolUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveLPAPLendingPoolUpdated)
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
		it.Event = new(AaveLPAPLendingPoolUpdated)
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
func (it *AaveLPAPLendingPoolUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveLPAPLendingPoolUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveLPAPLendingPoolUpdated represents a LendingPoolUpdated event raised by the AaveLPAP contract.
type AaveLPAPLendingPoolUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolUpdated is a free log retrieval operation binding the contract event 0xc4e6c6cdf28d0edbd8bcf071d724d33cc2e7a30be7d06443925656e9cb492aa4.
//
// Solidity: event LendingPoolUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) FilterLendingPoolUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AaveLPAPLendingPoolUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.FilterLogs(opts, "LendingPoolUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveLPAPLendingPoolUpdatedIterator{contract: _AaveLPAP.contract, event: "LendingPoolUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolUpdated is a free log subscription operation binding the contract event 0xc4e6c6cdf28d0edbd8bcf071d724d33cc2e7a30be7d06443925656e9cb492aa4.
//
// Solidity: event LendingPoolUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) WatchLendingPoolUpdated(opts *bind.WatchOpts, sink chan<- *AaveLPAPLendingPoolUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.WatchLogs(opts, "LendingPoolUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveLPAPLendingPoolUpdated)
				if err := _AaveLPAP.contract.UnpackLog(event, "LendingPoolUpdated", log); err != nil {
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

// ParseLendingPoolUpdated is a log parse operation binding the contract event 0xc4e6c6cdf28d0edbd8bcf071d724d33cc2e7a30be7d06443925656e9cb492aa4.
//
// Solidity: event LendingPoolUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) ParseLendingPoolUpdated(log types.Log) (*AaveLPAPLendingPoolUpdated, error) {
	event := new(AaveLPAPLendingPoolUpdated)
	if err := _AaveLPAP.contract.UnpackLog(event, "LendingPoolUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveLPAPLendingRateOracleUpdatedIterator is returned from FilterLendingRateOracleUpdated and is used to iterate over the raw logs and unpacked data for LendingRateOracleUpdated events raised by the AaveLPAP contract.
type AaveLPAPLendingRateOracleUpdatedIterator struct {
	Event *AaveLPAPLendingRateOracleUpdated // Event containing the contract specifics and raw log

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
func (it *AaveLPAPLendingRateOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveLPAPLendingRateOracleUpdated)
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
		it.Event = new(AaveLPAPLendingRateOracleUpdated)
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
func (it *AaveLPAPLendingRateOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveLPAPLendingRateOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveLPAPLendingRateOracleUpdated represents a LendingRateOracleUpdated event raised by the AaveLPAP contract.
type AaveLPAPLendingRateOracleUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingRateOracleUpdated is a free log retrieval operation binding the contract event 0x5c29179aba6942020a8a2d38f65de02fb6b7f784e7f049ed3a3cab97621859b5.
//
// Solidity: event LendingRateOracleUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) FilterLendingRateOracleUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AaveLPAPLendingRateOracleUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.FilterLogs(opts, "LendingRateOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveLPAPLendingRateOracleUpdatedIterator{contract: _AaveLPAP.contract, event: "LendingRateOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingRateOracleUpdated is a free log subscription operation binding the contract event 0x5c29179aba6942020a8a2d38f65de02fb6b7f784e7f049ed3a3cab97621859b5.
//
// Solidity: event LendingRateOracleUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) WatchLendingRateOracleUpdated(opts *bind.WatchOpts, sink chan<- *AaveLPAPLendingRateOracleUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.WatchLogs(opts, "LendingRateOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveLPAPLendingRateOracleUpdated)
				if err := _AaveLPAP.contract.UnpackLog(event, "LendingRateOracleUpdated", log); err != nil {
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

// ParseLendingRateOracleUpdated is a log parse operation binding the contract event 0x5c29179aba6942020a8a2d38f65de02fb6b7f784e7f049ed3a3cab97621859b5.
//
// Solidity: event LendingRateOracleUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) ParseLendingRateOracleUpdated(log types.Log) (*AaveLPAPLendingRateOracleUpdated, error) {
	event := new(AaveLPAPLendingRateOracleUpdated)
	if err := _AaveLPAP.contract.UnpackLog(event, "LendingRateOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveLPAPOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AaveLPAP contract.
type AaveLPAPOwnershipTransferredIterator struct {
	Event *AaveLPAPOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AaveLPAPOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveLPAPOwnershipTransferred)
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
		it.Event = new(AaveLPAPOwnershipTransferred)
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
func (it *AaveLPAPOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveLPAPOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveLPAPOwnershipTransferred represents a OwnershipTransferred event raised by the AaveLPAP contract.
type AaveLPAPOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AaveLPAP *AaveLPAPFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AaveLPAPOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AaveLPAP.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AaveLPAPOwnershipTransferredIterator{contract: _AaveLPAP.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AaveLPAP *AaveLPAPFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AaveLPAPOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AaveLPAP.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveLPAPOwnershipTransferred)
				if err := _AaveLPAP.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AaveLPAP *AaveLPAPFilterer) ParseOwnershipTransferred(log types.Log) (*AaveLPAPOwnershipTransferred, error) {
	event := new(AaveLPAPOwnershipTransferred)
	if err := _AaveLPAP.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveLPAPPriceOracleUpdatedIterator is returned from FilterPriceOracleUpdated and is used to iterate over the raw logs and unpacked data for PriceOracleUpdated events raised by the AaveLPAP contract.
type AaveLPAPPriceOracleUpdatedIterator struct {
	Event *AaveLPAPPriceOracleUpdated // Event containing the contract specifics and raw log

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
func (it *AaveLPAPPriceOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveLPAPPriceOracleUpdated)
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
		it.Event = new(AaveLPAPPriceOracleUpdated)
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
func (it *AaveLPAPPriceOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveLPAPPriceOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveLPAPPriceOracleUpdated represents a PriceOracleUpdated event raised by the AaveLPAP contract.
type AaveLPAPPriceOracleUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPriceOracleUpdated is a free log retrieval operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) FilterPriceOracleUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AaveLPAPPriceOracleUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.FilterLogs(opts, "PriceOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveLPAPPriceOracleUpdatedIterator{contract: _AaveLPAP.contract, event: "PriceOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceOracleUpdated is a free log subscription operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) WatchPriceOracleUpdated(opts *bind.WatchOpts, sink chan<- *AaveLPAPPriceOracleUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.WatchLogs(opts, "PriceOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveLPAPPriceOracleUpdated)
				if err := _AaveLPAP.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
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

// ParsePriceOracleUpdated is a log parse operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) ParsePriceOracleUpdated(log types.Log) (*AaveLPAPPriceOracleUpdated, error) {
	event := new(AaveLPAPPriceOracleUpdated)
	if err := _AaveLPAP.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveLPAPProxyCreatedIterator is returned from FilterProxyCreated and is used to iterate over the raw logs and unpacked data for ProxyCreated events raised by the AaveLPAP contract.
type AaveLPAPProxyCreatedIterator struct {
	Event *AaveLPAPProxyCreated // Event containing the contract specifics and raw log

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
func (it *AaveLPAPProxyCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveLPAPProxyCreated)
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
		it.Event = new(AaveLPAPProxyCreated)
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
func (it *AaveLPAPProxyCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveLPAPProxyCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveLPAPProxyCreated represents a ProxyCreated event raised by the AaveLPAP contract.
type AaveLPAPProxyCreated struct {
	Id         [32]byte
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProxyCreated is a free log retrieval operation binding the contract event 0x1eb35cb4b5bbb23d152f3b4016a5a46c37a07ae930ed0956aba951e231142438.
//
// Solidity: event ProxyCreated(bytes32 id, address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) FilterProxyCreated(opts *bind.FilterOpts, newAddress []common.Address) (*AaveLPAPProxyCreatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.FilterLogs(opts, "ProxyCreated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveLPAPProxyCreatedIterator{contract: _AaveLPAP.contract, event: "ProxyCreated", logs: logs, sub: sub}, nil
}

// WatchProxyCreated is a free log subscription operation binding the contract event 0x1eb35cb4b5bbb23d152f3b4016a5a46c37a07ae930ed0956aba951e231142438.
//
// Solidity: event ProxyCreated(bytes32 id, address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) WatchProxyCreated(opts *bind.WatchOpts, sink chan<- *AaveLPAPProxyCreated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.WatchLogs(opts, "ProxyCreated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveLPAPProxyCreated)
				if err := _AaveLPAP.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
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

// ParseProxyCreated is a log parse operation binding the contract event 0x1eb35cb4b5bbb23d152f3b4016a5a46c37a07ae930ed0956aba951e231142438.
//
// Solidity: event ProxyCreated(bytes32 id, address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) ParseProxyCreated(log types.Log) (*AaveLPAPProxyCreated, error) {
	event := new(AaveLPAPProxyCreated)
	if err := _AaveLPAP.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveLPAPTokenDistributorUpdatedIterator is returned from FilterTokenDistributorUpdated and is used to iterate over the raw logs and unpacked data for TokenDistributorUpdated events raised by the AaveLPAP contract.
type AaveLPAPTokenDistributorUpdatedIterator struct {
	Event *AaveLPAPTokenDistributorUpdated // Event containing the contract specifics and raw log

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
func (it *AaveLPAPTokenDistributorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveLPAPTokenDistributorUpdated)
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
		it.Event = new(AaveLPAPTokenDistributorUpdated)
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
func (it *AaveLPAPTokenDistributorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveLPAPTokenDistributorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveLPAPTokenDistributorUpdated represents a TokenDistributorUpdated event raised by the AaveLPAP contract.
type AaveLPAPTokenDistributorUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokenDistributorUpdated is a free log retrieval operation binding the contract event 0xa8b48a56ec01f81c3615a21ec43e16b925c26293e0801cf6330427f2a687f053.
//
// Solidity: event TokenDistributorUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) FilterTokenDistributorUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*AaveLPAPTokenDistributorUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.FilterLogs(opts, "TokenDistributorUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveLPAPTokenDistributorUpdatedIterator{contract: _AaveLPAP.contract, event: "TokenDistributorUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenDistributorUpdated is a free log subscription operation binding the contract event 0xa8b48a56ec01f81c3615a21ec43e16b925c26293e0801cf6330427f2a687f053.
//
// Solidity: event TokenDistributorUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) WatchTokenDistributorUpdated(opts *bind.WatchOpts, sink chan<- *AaveLPAPTokenDistributorUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveLPAP.contract.WatchLogs(opts, "TokenDistributorUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveLPAPTokenDistributorUpdated)
				if err := _AaveLPAP.contract.UnpackLog(event, "TokenDistributorUpdated", log); err != nil {
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

// ParseTokenDistributorUpdated is a log parse operation binding the contract event 0xa8b48a56ec01f81c3615a21ec43e16b925c26293e0801cf6330427f2a687f053.
//
// Solidity: event TokenDistributorUpdated(address indexed newAddress)
func (_AaveLPAP *AaveLPAPFilterer) ParseTokenDistributorUpdated(log types.Log) (*AaveLPAPTokenDistributorUpdated, error) {
	event := new(AaveLPAPTokenDistributorUpdated)
	if err := _AaveLPAP.contract.UnpackLog(event, "TokenDistributorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
