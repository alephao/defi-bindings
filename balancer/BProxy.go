// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balancer

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
)

// ExchangeProxySwap is an auto generated low-level Go binding around an user-defined struct.
type ExchangeProxySwap struct {
	Pool              common.Address
	TokenIn           common.Address
	TokenOut          common.Address
	SwapAmount        *big.Int
	LimitReturnAmount *big.Int
	MaxPrice          *big.Int
}

// BalancerMetaData contains all meta data concerning the Balancer contract.
var BalancerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeProxy.Swap[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"contractTokenInterface\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractTokenInterface\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTotalAmountOut\",\"type\":\"uint256\"}],\"name\":\"batchSwapExactIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmountOut\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeProxy.Swap[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"contractTokenInterface\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractTokenInterface\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalAmountIn\",\"type\":\"uint256\"}],\"name\":\"batchSwapExactOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmountIn\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeProxy.Swap[][]\",\"name\":\"swapSequences\",\"type\":\"tuple[][]\"},{\"internalType\":\"contractTokenInterface\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractTokenInterface\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTotalAmountOut\",\"type\":\"uint256\"}],\"name\":\"multihopBatchSwapExactIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmountOut\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeProxy.Swap[][]\",\"name\":\"swapSequences\",\"type\":\"tuple[][]\"},{\"internalType\":\"contractTokenInterface\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractTokenInterface\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalAmountIn\",\"type\":\"uint256\"}],\"name\":\"multihopBatchSwapExactOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmountIn\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractTokenInterface\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractTokenInterface\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTotalAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nPools\",\"type\":\"uint256\"}],\"name\":\"smartSwapExactIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmountOut\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractTokenInterface\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractTokenInterface\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nPools\",\"type\":\"uint256\"}],\"name\":\"smartSwapExactOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmountIn\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nPools\",\"type\":\"uint256\"}],\"name\":\"viewSplitExactIn\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeProxy.Swap[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"totalOutput\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nPools\",\"type\":\"uint256\"}],\"name\":\"viewSplitExactOut\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeProxy.Swap[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"totalOutput\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BalancerABI is the input ABI used to generate the binding from.
// Deprecated: Use BalancerMetaData.ABI instead.
var BalancerABI = BalancerMetaData.ABI

// Balancer is an auto generated Go binding around an Ethereum contract.
type Balancer struct {
	BalancerCaller     // Read-only binding to the contract
	BalancerTransactor // Write-only binding to the contract
	BalancerFilterer   // Log filterer for contract events
}

// BalancerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalancerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalancerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalancerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalancerSession struct {
	Contract     *Balancer         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalancerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalancerCallerSession struct {
	Contract *BalancerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BalancerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalancerTransactorSession struct {
	Contract     *BalancerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BalancerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalancerRaw struct {
	Contract *Balancer // Generic contract binding to access the raw methods on
}

// BalancerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalancerCallerRaw struct {
	Contract *BalancerCaller // Generic read-only contract binding to access the raw methods on
}

// BalancerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalancerTransactorRaw struct {
	Contract *BalancerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancer creates a new instance of Balancer, bound to a specific deployed contract.
func NewBalancer(address common.Address, backend bind.ContractBackend) (*Balancer, error) {
	contract, err := bindBalancer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Balancer{BalancerCaller: BalancerCaller{contract: contract}, BalancerTransactor: BalancerTransactor{contract: contract}, BalancerFilterer: BalancerFilterer{contract: contract}}, nil
}

// NewBalancerCaller creates a new read-only instance of Balancer, bound to a specific deployed contract.
func NewBalancerCaller(address common.Address, caller bind.ContractCaller) (*BalancerCaller, error) {
	contract, err := bindBalancer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerCaller{contract: contract}, nil
}

// NewBalancerTransactor creates a new write-only instance of Balancer, bound to a specific deployed contract.
func NewBalancerTransactor(address common.Address, transactor bind.ContractTransactor) (*BalancerTransactor, error) {
	contract, err := bindBalancer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerTransactor{contract: contract}, nil
}

// NewBalancerFilterer creates a new log filterer instance of Balancer, bound to a specific deployed contract.
func NewBalancerFilterer(address common.Address, filterer bind.ContractFilterer) (*BalancerFilterer, error) {
	contract, err := bindBalancer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalancerFilterer{contract: contract}, nil
}

// bindBalancer binds a generic wrapper to an already deployed contract.
func bindBalancer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalancerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balancer *BalancerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balancer.Contract.BalancerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balancer *BalancerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancer.Contract.BalancerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balancer *BalancerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balancer.Contract.BalancerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balancer *BalancerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balancer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balancer *BalancerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balancer *BalancerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balancer.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Balancer *BalancerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Balancer.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Balancer *BalancerSession) IsOwner() (bool, error) {
	return _Balancer.Contract.IsOwner(&_Balancer.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Balancer *BalancerCallerSession) IsOwner() (bool, error) {
	return _Balancer.Contract.IsOwner(&_Balancer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Balancer *BalancerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Balancer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Balancer *BalancerSession) Owner() (common.Address, error) {
	return _Balancer.Contract.Owner(&_Balancer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Balancer *BalancerCallerSession) Owner() (common.Address, error) {
	return _Balancer.Contract.Owner(&_Balancer.CallOpts)
}

// ViewSplitExactIn is a free data retrieval call binding the contract method 0x4b0f93fb.
//
// Solidity: function viewSplitExactIn(address tokenIn, address tokenOut, uint256 swapAmount, uint256 nPools) view returns((address,address,address,uint256,uint256,uint256)[] swaps, uint256 totalOutput)
func (_Balancer *BalancerCaller) ViewSplitExactIn(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address, swapAmount *big.Int, nPools *big.Int) (struct {
	Swaps       []ExchangeProxySwap
	TotalOutput *big.Int
}, error) {
	var out []interface{}
	err := _Balancer.contract.Call(opts, &out, "viewSplitExactIn", tokenIn, tokenOut, swapAmount, nPools)

	outstruct := new(struct {
		Swaps       []ExchangeProxySwap
		TotalOutput *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Swaps = *abi.ConvertType(out[0], new([]ExchangeProxySwap)).(*[]ExchangeProxySwap)
	outstruct.TotalOutput = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ViewSplitExactIn is a free data retrieval call binding the contract method 0x4b0f93fb.
//
// Solidity: function viewSplitExactIn(address tokenIn, address tokenOut, uint256 swapAmount, uint256 nPools) view returns((address,address,address,uint256,uint256,uint256)[] swaps, uint256 totalOutput)
func (_Balancer *BalancerSession) ViewSplitExactIn(tokenIn common.Address, tokenOut common.Address, swapAmount *big.Int, nPools *big.Int) (struct {
	Swaps       []ExchangeProxySwap
	TotalOutput *big.Int
}, error) {
	return _Balancer.Contract.ViewSplitExactIn(&_Balancer.CallOpts, tokenIn, tokenOut, swapAmount, nPools)
}

// ViewSplitExactIn is a free data retrieval call binding the contract method 0x4b0f93fb.
//
// Solidity: function viewSplitExactIn(address tokenIn, address tokenOut, uint256 swapAmount, uint256 nPools) view returns((address,address,address,uint256,uint256,uint256)[] swaps, uint256 totalOutput)
func (_Balancer *BalancerCallerSession) ViewSplitExactIn(tokenIn common.Address, tokenOut common.Address, swapAmount *big.Int, nPools *big.Int) (struct {
	Swaps       []ExchangeProxySwap
	TotalOutput *big.Int
}, error) {
	return _Balancer.Contract.ViewSplitExactIn(&_Balancer.CallOpts, tokenIn, tokenOut, swapAmount, nPools)
}

// ViewSplitExactOut is a free data retrieval call binding the contract method 0x368bb1fc.
//
// Solidity: function viewSplitExactOut(address tokenIn, address tokenOut, uint256 swapAmount, uint256 nPools) view returns((address,address,address,uint256,uint256,uint256)[] swaps, uint256 totalOutput)
func (_Balancer *BalancerCaller) ViewSplitExactOut(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address, swapAmount *big.Int, nPools *big.Int) (struct {
	Swaps       []ExchangeProxySwap
	TotalOutput *big.Int
}, error) {
	var out []interface{}
	err := _Balancer.contract.Call(opts, &out, "viewSplitExactOut", tokenIn, tokenOut, swapAmount, nPools)

	outstruct := new(struct {
		Swaps       []ExchangeProxySwap
		TotalOutput *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Swaps = *abi.ConvertType(out[0], new([]ExchangeProxySwap)).(*[]ExchangeProxySwap)
	outstruct.TotalOutput = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ViewSplitExactOut is a free data retrieval call binding the contract method 0x368bb1fc.
//
// Solidity: function viewSplitExactOut(address tokenIn, address tokenOut, uint256 swapAmount, uint256 nPools) view returns((address,address,address,uint256,uint256,uint256)[] swaps, uint256 totalOutput)
func (_Balancer *BalancerSession) ViewSplitExactOut(tokenIn common.Address, tokenOut common.Address, swapAmount *big.Int, nPools *big.Int) (struct {
	Swaps       []ExchangeProxySwap
	TotalOutput *big.Int
}, error) {
	return _Balancer.Contract.ViewSplitExactOut(&_Balancer.CallOpts, tokenIn, tokenOut, swapAmount, nPools)
}

// ViewSplitExactOut is a free data retrieval call binding the contract method 0x368bb1fc.
//
// Solidity: function viewSplitExactOut(address tokenIn, address tokenOut, uint256 swapAmount, uint256 nPools) view returns((address,address,address,uint256,uint256,uint256)[] swaps, uint256 totalOutput)
func (_Balancer *BalancerCallerSession) ViewSplitExactOut(tokenIn common.Address, tokenOut common.Address, swapAmount *big.Int, nPools *big.Int) (struct {
	Swaps       []ExchangeProxySwap
	TotalOutput *big.Int
}, error) {
	return _Balancer.Contract.ViewSplitExactOut(&_Balancer.CallOpts, tokenIn, tokenOut, swapAmount, nPools)
}

// BatchSwapExactIn is a paid mutator transaction binding the contract method 0x8743ad58.
//
// Solidity: function batchSwapExactIn((address,address,address,uint256,uint256,uint256)[] swaps, address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut) payable returns(uint256 totalAmountOut)
func (_Balancer *BalancerTransactor) BatchSwapExactIn(opts *bind.TransactOpts, swaps []ExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int) (*types.Transaction, error) {
	return _Balancer.contract.Transact(opts, "batchSwapExactIn", swaps, tokenIn, tokenOut, totalAmountIn, minTotalAmountOut)
}

// BatchSwapExactIn is a paid mutator transaction binding the contract method 0x8743ad58.
//
// Solidity: function batchSwapExactIn((address,address,address,uint256,uint256,uint256)[] swaps, address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut) payable returns(uint256 totalAmountOut)
func (_Balancer *BalancerSession) BatchSwapExactIn(swaps []ExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int) (*types.Transaction, error) {
	return _Balancer.Contract.BatchSwapExactIn(&_Balancer.TransactOpts, swaps, tokenIn, tokenOut, totalAmountIn, minTotalAmountOut)
}

// BatchSwapExactIn is a paid mutator transaction binding the contract method 0x8743ad58.
//
// Solidity: function batchSwapExactIn((address,address,address,uint256,uint256,uint256)[] swaps, address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut) payable returns(uint256 totalAmountOut)
func (_Balancer *BalancerTransactorSession) BatchSwapExactIn(swaps []ExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int) (*types.Transaction, error) {
	return _Balancer.Contract.BatchSwapExactIn(&_Balancer.TransactOpts, swaps, tokenIn, tokenOut, totalAmountIn, minTotalAmountOut)
}

// BatchSwapExactOut is a paid mutator transaction binding the contract method 0x2db58134.
//
// Solidity: function batchSwapExactOut((address,address,address,uint256,uint256,uint256)[] swaps, address tokenIn, address tokenOut, uint256 maxTotalAmountIn) payable returns(uint256 totalAmountIn)
func (_Balancer *BalancerTransactor) BatchSwapExactOut(opts *bind.TransactOpts, swaps []ExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, maxTotalAmountIn *big.Int) (*types.Transaction, error) {
	return _Balancer.contract.Transact(opts, "batchSwapExactOut", swaps, tokenIn, tokenOut, maxTotalAmountIn)
}

// BatchSwapExactOut is a paid mutator transaction binding the contract method 0x2db58134.
//
// Solidity: function batchSwapExactOut((address,address,address,uint256,uint256,uint256)[] swaps, address tokenIn, address tokenOut, uint256 maxTotalAmountIn) payable returns(uint256 totalAmountIn)
func (_Balancer *BalancerSession) BatchSwapExactOut(swaps []ExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, maxTotalAmountIn *big.Int) (*types.Transaction, error) {
	return _Balancer.Contract.BatchSwapExactOut(&_Balancer.TransactOpts, swaps, tokenIn, tokenOut, maxTotalAmountIn)
}

// BatchSwapExactOut is a paid mutator transaction binding the contract method 0x2db58134.
//
// Solidity: function batchSwapExactOut((address,address,address,uint256,uint256,uint256)[] swaps, address tokenIn, address tokenOut, uint256 maxTotalAmountIn) payable returns(uint256 totalAmountIn)
func (_Balancer *BalancerTransactorSession) BatchSwapExactOut(swaps []ExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, maxTotalAmountIn *big.Int) (*types.Transaction, error) {
	return _Balancer.Contract.BatchSwapExactOut(&_Balancer.TransactOpts, swaps, tokenIn, tokenOut, maxTotalAmountIn)
}

// MultihopBatchSwapExactIn is a paid mutator transaction binding the contract method 0xe2b39746.
//
// Solidity: function multihopBatchSwapExactIn((address,address,address,uint256,uint256,uint256)[][] swapSequences, address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut) payable returns(uint256 totalAmountOut)
func (_Balancer *BalancerTransactor) MultihopBatchSwapExactIn(opts *bind.TransactOpts, swapSequences [][]ExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int) (*types.Transaction, error) {
	return _Balancer.contract.Transact(opts, "multihopBatchSwapExactIn", swapSequences, tokenIn, tokenOut, totalAmountIn, minTotalAmountOut)
}

// MultihopBatchSwapExactIn is a paid mutator transaction binding the contract method 0xe2b39746.
//
// Solidity: function multihopBatchSwapExactIn((address,address,address,uint256,uint256,uint256)[][] swapSequences, address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut) payable returns(uint256 totalAmountOut)
func (_Balancer *BalancerSession) MultihopBatchSwapExactIn(swapSequences [][]ExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int) (*types.Transaction, error) {
	return _Balancer.Contract.MultihopBatchSwapExactIn(&_Balancer.TransactOpts, swapSequences, tokenIn, tokenOut, totalAmountIn, minTotalAmountOut)
}

// MultihopBatchSwapExactIn is a paid mutator transaction binding the contract method 0xe2b39746.
//
// Solidity: function multihopBatchSwapExactIn((address,address,address,uint256,uint256,uint256)[][] swapSequences, address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut) payable returns(uint256 totalAmountOut)
func (_Balancer *BalancerTransactorSession) MultihopBatchSwapExactIn(swapSequences [][]ExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int) (*types.Transaction, error) {
	return _Balancer.Contract.MultihopBatchSwapExactIn(&_Balancer.TransactOpts, swapSequences, tokenIn, tokenOut, totalAmountIn, minTotalAmountOut)
}

// MultihopBatchSwapExactOut is a paid mutator transaction binding the contract method 0x86b2ecc4.
//
// Solidity: function multihopBatchSwapExactOut((address,address,address,uint256,uint256,uint256)[][] swapSequences, address tokenIn, address tokenOut, uint256 maxTotalAmountIn) payable returns(uint256 totalAmountIn)
func (_Balancer *BalancerTransactor) MultihopBatchSwapExactOut(opts *bind.TransactOpts, swapSequences [][]ExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, maxTotalAmountIn *big.Int) (*types.Transaction, error) {
	return _Balancer.contract.Transact(opts, "multihopBatchSwapExactOut", swapSequences, tokenIn, tokenOut, maxTotalAmountIn)
}

// MultihopBatchSwapExactOut is a paid mutator transaction binding the contract method 0x86b2ecc4.
//
// Solidity: function multihopBatchSwapExactOut((address,address,address,uint256,uint256,uint256)[][] swapSequences, address tokenIn, address tokenOut, uint256 maxTotalAmountIn) payable returns(uint256 totalAmountIn)
func (_Balancer *BalancerSession) MultihopBatchSwapExactOut(swapSequences [][]ExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, maxTotalAmountIn *big.Int) (*types.Transaction, error) {
	return _Balancer.Contract.MultihopBatchSwapExactOut(&_Balancer.TransactOpts, swapSequences, tokenIn, tokenOut, maxTotalAmountIn)
}

// MultihopBatchSwapExactOut is a paid mutator transaction binding the contract method 0x86b2ecc4.
//
// Solidity: function multihopBatchSwapExactOut((address,address,address,uint256,uint256,uint256)[][] swapSequences, address tokenIn, address tokenOut, uint256 maxTotalAmountIn) payable returns(uint256 totalAmountIn)
func (_Balancer *BalancerTransactorSession) MultihopBatchSwapExactOut(swapSequences [][]ExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, maxTotalAmountIn *big.Int) (*types.Transaction, error) {
	return _Balancer.Contract.MultihopBatchSwapExactOut(&_Balancer.TransactOpts, swapSequences, tokenIn, tokenOut, maxTotalAmountIn)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Balancer *BalancerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancer.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Balancer *BalancerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Balancer.Contract.RenounceOwnership(&_Balancer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Balancer *BalancerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Balancer.Contract.RenounceOwnership(&_Balancer.TransactOpts)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_Balancer *BalancerTransactor) SetRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _Balancer.contract.Transact(opts, "setRegistry", _registry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_Balancer *BalancerSession) SetRegistry(_registry common.Address) (*types.Transaction, error) {
	return _Balancer.Contract.SetRegistry(&_Balancer.TransactOpts, _registry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_Balancer *BalancerTransactorSession) SetRegistry(_registry common.Address) (*types.Transaction, error) {
	return _Balancer.Contract.SetRegistry(&_Balancer.TransactOpts, _registry)
}

// SmartSwapExactIn is a paid mutator transaction binding the contract method 0x21b0eb85.
//
// Solidity: function smartSwapExactIn(address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut, uint256 nPools) payable returns(uint256 totalAmountOut)
func (_Balancer *BalancerTransactor) SmartSwapExactIn(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int, nPools *big.Int) (*types.Transaction, error) {
	return _Balancer.contract.Transact(opts, "smartSwapExactIn", tokenIn, tokenOut, totalAmountIn, minTotalAmountOut, nPools)
}

// SmartSwapExactIn is a paid mutator transaction binding the contract method 0x21b0eb85.
//
// Solidity: function smartSwapExactIn(address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut, uint256 nPools) payable returns(uint256 totalAmountOut)
func (_Balancer *BalancerSession) SmartSwapExactIn(tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int, nPools *big.Int) (*types.Transaction, error) {
	return _Balancer.Contract.SmartSwapExactIn(&_Balancer.TransactOpts, tokenIn, tokenOut, totalAmountIn, minTotalAmountOut, nPools)
}

// SmartSwapExactIn is a paid mutator transaction binding the contract method 0x21b0eb85.
//
// Solidity: function smartSwapExactIn(address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut, uint256 nPools) payable returns(uint256 totalAmountOut)
func (_Balancer *BalancerTransactorSession) SmartSwapExactIn(tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int, nPools *big.Int) (*types.Transaction, error) {
	return _Balancer.Contract.SmartSwapExactIn(&_Balancer.TransactOpts, tokenIn, tokenOut, totalAmountIn, minTotalAmountOut, nPools)
}

// SmartSwapExactOut is a paid mutator transaction binding the contract method 0xb40f39ee.
//
// Solidity: function smartSwapExactOut(address tokenIn, address tokenOut, uint256 totalAmountOut, uint256 maxTotalAmountIn, uint256 nPools) payable returns(uint256 totalAmountIn)
func (_Balancer *BalancerTransactor) SmartSwapExactOut(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, totalAmountOut *big.Int, maxTotalAmountIn *big.Int, nPools *big.Int) (*types.Transaction, error) {
	return _Balancer.contract.Transact(opts, "smartSwapExactOut", tokenIn, tokenOut, totalAmountOut, maxTotalAmountIn, nPools)
}

// SmartSwapExactOut is a paid mutator transaction binding the contract method 0xb40f39ee.
//
// Solidity: function smartSwapExactOut(address tokenIn, address tokenOut, uint256 totalAmountOut, uint256 maxTotalAmountIn, uint256 nPools) payable returns(uint256 totalAmountIn)
func (_Balancer *BalancerSession) SmartSwapExactOut(tokenIn common.Address, tokenOut common.Address, totalAmountOut *big.Int, maxTotalAmountIn *big.Int, nPools *big.Int) (*types.Transaction, error) {
	return _Balancer.Contract.SmartSwapExactOut(&_Balancer.TransactOpts, tokenIn, tokenOut, totalAmountOut, maxTotalAmountIn, nPools)
}

// SmartSwapExactOut is a paid mutator transaction binding the contract method 0xb40f39ee.
//
// Solidity: function smartSwapExactOut(address tokenIn, address tokenOut, uint256 totalAmountOut, uint256 maxTotalAmountIn, uint256 nPools) payable returns(uint256 totalAmountIn)
func (_Balancer *BalancerTransactorSession) SmartSwapExactOut(tokenIn common.Address, tokenOut common.Address, totalAmountOut *big.Int, maxTotalAmountIn *big.Int, nPools *big.Int) (*types.Transaction, error) {
	return _Balancer.Contract.SmartSwapExactOut(&_Balancer.TransactOpts, tokenIn, tokenOut, totalAmountOut, maxTotalAmountIn, nPools)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Balancer *BalancerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Balancer.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Balancer *BalancerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Balancer.Contract.TransferOwnership(&_Balancer.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Balancer *BalancerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Balancer.Contract.TransferOwnership(&_Balancer.TransactOpts, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Balancer *BalancerTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Balancer.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Balancer *BalancerSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Balancer.Contract.Fallback(&_Balancer.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Balancer *BalancerTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Balancer.Contract.Fallback(&_Balancer.TransactOpts, calldata)
}

// BalancerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Balancer contract.
type BalancerOwnershipTransferredIterator struct {
	Event *BalancerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BalancerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerOwnershipTransferred)
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
		it.Event = new(BalancerOwnershipTransferred)
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
func (it *BalancerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerOwnershipTransferred represents a OwnershipTransferred event raised by the Balancer contract.
type BalancerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Balancer *BalancerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BalancerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Balancer.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BalancerOwnershipTransferredIterator{contract: _Balancer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Balancer *BalancerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BalancerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Balancer.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerOwnershipTransferred)
				if err := _Balancer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Balancer *BalancerFilterer) ParseOwnershipTransferred(log types.Log) (*BalancerOwnershipTransferred, error) {
	event := new(BalancerOwnershipTransferred)
	if err := _Balancer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
