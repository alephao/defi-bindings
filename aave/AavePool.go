// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aave

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

// AaveMetaData contains all meta data concerning the Aave contract.
var AaveMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_borrowRateMode\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_borrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_originationFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_borrowBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"_referral\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"_referral\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_protocolFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"FlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_collateral\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_purchaseAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_liquidatedCollateralAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_accruedBorrowInterest\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_receiveAToken\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"LiquidationCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_collateral\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_feeLiquidated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_liquidatedCollateralForFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"OriginationFeeLiquidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newStableRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_borrowBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"RebalanceStableBorrowRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"RedeemUnderlying\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_repayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountMinusFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_borrowBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"ReserveUsedAsCollateralDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"ReserveUsedAsCollateralEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newRateMode\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_borrowBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"LENDINGPOOL_REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"UINT_MAX_VALUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addressesProvider\",\"outputs\":[{\"internalType\":\"contractLendingPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interestRateMode\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_referralCode\",\"type\":\"uint16\"}],\"name\":\"borrow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"core\",\"outputs\":[{\"internalType\":\"contractLendingPoolCore\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dataProvider\",\"outputs\":[{\"internalType\":\"contractLendingPoolDataProvider\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_referralCode\",\"type\":\"uint16\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_params\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveConfigurationData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationBonus\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"interestRateStrategyAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"borrowingEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"stableBorrowRateEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowsStable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowsVariable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"averageStableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableBorrowIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserAccountData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalLiquidityETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCollateralETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowsETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFeesETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBorrowsETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentLiquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserReserveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"currentATokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentBorrowBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalBorrowBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRateMode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originationFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableBorrowIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractLendingPoolAddressesProvider\",\"name\":\"_addressesProvider\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_purchaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_receiveAToken\",\"type\":\"bool\"}],\"name\":\"liquidationCall\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"parametersProvider\",\"outputs\":[{\"internalType\":\"contractLendingPoolParametersProvider\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"rebalanceStableBorrowRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_aTokenBalanceAfterRedeem\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_onBehalfOf\",\"type\":\"address\"}],\"name\":\"repay\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_useAsCollateral\",\"type\":\"bool\"}],\"name\":\"setUserUseReserveAsCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"swapBorrowRateMode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AaveABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveMetaData.ABI instead.
var AaveABI = AaveMetaData.ABI

// Aave is an auto generated Go binding around an Ethereum contract.
type Aave struct {
	AaveCaller     // Read-only binding to the contract
	AaveTransactor // Write-only binding to the contract
	AaveFilterer   // Log filterer for contract events
}

// AaveCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveSession struct {
	Contract     *Aave             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveCallerSession struct {
	Contract *AaveCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AaveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveTransactorSession struct {
	Contract     *AaveTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveRaw struct {
	Contract *Aave // Generic contract binding to access the raw methods on
}

// AaveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveCallerRaw struct {
	Contract *AaveCaller // Generic read-only contract binding to access the raw methods on
}

// AaveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveTransactorRaw struct {
	Contract *AaveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAave creates a new instance of Aave, bound to a specific deployed contract.
func NewAave(address common.Address, backend bind.ContractBackend) (*Aave, error) {
	contract, err := bindAave(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aave{AaveCaller: AaveCaller{contract: contract}, AaveTransactor: AaveTransactor{contract: contract}, AaveFilterer: AaveFilterer{contract: contract}}, nil
}

// NewAaveCaller creates a new read-only instance of Aave, bound to a specific deployed contract.
func NewAaveCaller(address common.Address, caller bind.ContractCaller) (*AaveCaller, error) {
	contract, err := bindAave(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveCaller{contract: contract}, nil
}

// NewAaveTransactor creates a new write-only instance of Aave, bound to a specific deployed contract.
func NewAaveTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveTransactor, error) {
	contract, err := bindAave(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveTransactor{contract: contract}, nil
}

// NewAaveFilterer creates a new log filterer instance of Aave, bound to a specific deployed contract.
func NewAaveFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveFilterer, error) {
	contract, err := bindAave(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveFilterer{contract: contract}, nil
}

// bindAave binds a generic wrapper to an already deployed contract.
func bindAave(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aave *AaveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aave.Contract.AaveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aave *AaveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aave.Contract.AaveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aave *AaveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aave.Contract.AaveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aave *AaveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aave.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aave *AaveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aave.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aave *AaveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aave.Contract.contract.Transact(opts, method, params...)
}

// LENDINGPOOLREVISION is a free data retrieval call binding the contract method 0x8afaff02.
//
// Solidity: function LENDINGPOOL_REVISION() view returns(uint256)
func (_Aave *AaveCaller) LENDINGPOOLREVISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "LENDINGPOOL_REVISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LENDINGPOOLREVISION is a free data retrieval call binding the contract method 0x8afaff02.
//
// Solidity: function LENDINGPOOL_REVISION() view returns(uint256)
func (_Aave *AaveSession) LENDINGPOOLREVISION() (*big.Int, error) {
	return _Aave.Contract.LENDINGPOOLREVISION(&_Aave.CallOpts)
}

// LENDINGPOOLREVISION is a free data retrieval call binding the contract method 0x8afaff02.
//
// Solidity: function LENDINGPOOL_REVISION() view returns(uint256)
func (_Aave *AaveCallerSession) LENDINGPOOLREVISION() (*big.Int, error) {
	return _Aave.Contract.LENDINGPOOLREVISION(&_Aave.CallOpts)
}

// UINTMAXVALUE is a free data retrieval call binding the contract method 0xd0fc81d2.
//
// Solidity: function UINT_MAX_VALUE() view returns(uint256)
func (_Aave *AaveCaller) UINTMAXVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "UINT_MAX_VALUE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UINTMAXVALUE is a free data retrieval call binding the contract method 0xd0fc81d2.
//
// Solidity: function UINT_MAX_VALUE() view returns(uint256)
func (_Aave *AaveSession) UINTMAXVALUE() (*big.Int, error) {
	return _Aave.Contract.UINTMAXVALUE(&_Aave.CallOpts)
}

// UINTMAXVALUE is a free data retrieval call binding the contract method 0xd0fc81d2.
//
// Solidity: function UINT_MAX_VALUE() view returns(uint256)
func (_Aave *AaveCallerSession) UINTMAXVALUE() (*big.Int, error) {
	return _Aave.Contract.UINTMAXVALUE(&_Aave.CallOpts)
}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_Aave *AaveCaller) AddressesProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "addressesProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_Aave *AaveSession) AddressesProvider() (common.Address, error) {
	return _Aave.Contract.AddressesProvider(&_Aave.CallOpts)
}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_Aave *AaveCallerSession) AddressesProvider() (common.Address, error) {
	return _Aave.Contract.AddressesProvider(&_Aave.CallOpts)
}

// Core is a free data retrieval call binding the contract method 0xf2f4eb26.
//
// Solidity: function core() view returns(address)
func (_Aave *AaveCaller) Core(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "core")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Core is a free data retrieval call binding the contract method 0xf2f4eb26.
//
// Solidity: function core() view returns(address)
func (_Aave *AaveSession) Core() (common.Address, error) {
	return _Aave.Contract.Core(&_Aave.CallOpts)
}

// Core is a free data retrieval call binding the contract method 0xf2f4eb26.
//
// Solidity: function core() view returns(address)
func (_Aave *AaveCallerSession) Core() (common.Address, error) {
	return _Aave.Contract.Core(&_Aave.CallOpts)
}

// DataProvider is a free data retrieval call binding the contract method 0xb334ed86.
//
// Solidity: function dataProvider() view returns(address)
func (_Aave *AaveCaller) DataProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "dataProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataProvider is a free data retrieval call binding the contract method 0xb334ed86.
//
// Solidity: function dataProvider() view returns(address)
func (_Aave *AaveSession) DataProvider() (common.Address, error) {
	return _Aave.Contract.DataProvider(&_Aave.CallOpts)
}

// DataProvider is a free data retrieval call binding the contract method 0xb334ed86.
//
// Solidity: function dataProvider() view returns(address)
func (_Aave *AaveCallerSession) DataProvider() (common.Address, error) {
	return _Aave.Contract.DataProvider(&_Aave.CallOpts)
}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address _reserve) view returns(uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, address interestRateStrategyAddress, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive)
func (_Aave *AaveCaller) GetReserveConfigurationData(opts *bind.CallOpts, _reserve common.Address) (struct {
	Ltv                         *big.Int
	LiquidationThreshold        *big.Int
	LiquidationBonus            *big.Int
	InterestRateStrategyAddress common.Address
	UsageAsCollateralEnabled    bool
	BorrowingEnabled            bool
	StableBorrowRateEnabled     bool
	IsActive                    bool
}, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserveConfigurationData", _reserve)

	outstruct := new(struct {
		Ltv                         *big.Int
		LiquidationThreshold        *big.Int
		LiquidationBonus            *big.Int
		InterestRateStrategyAddress common.Address
		UsageAsCollateralEnabled    bool
		BorrowingEnabled            bool
		StableBorrowRateEnabled     bool
		IsActive                    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ltv = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LiquidationThreshold = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LiquidationBonus = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.InterestRateStrategyAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.UsageAsCollateralEnabled = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.BorrowingEnabled = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.StableBorrowRateEnabled = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.IsActive = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address _reserve) view returns(uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, address interestRateStrategyAddress, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive)
func (_Aave *AaveSession) GetReserveConfigurationData(_reserve common.Address) (struct {
	Ltv                         *big.Int
	LiquidationThreshold        *big.Int
	LiquidationBonus            *big.Int
	InterestRateStrategyAddress common.Address
	UsageAsCollateralEnabled    bool
	BorrowingEnabled            bool
	StableBorrowRateEnabled     bool
	IsActive                    bool
}, error) {
	return _Aave.Contract.GetReserveConfigurationData(&_Aave.CallOpts, _reserve)
}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address _reserve) view returns(uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, address interestRateStrategyAddress, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive)
func (_Aave *AaveCallerSession) GetReserveConfigurationData(_reserve common.Address) (struct {
	Ltv                         *big.Int
	LiquidationThreshold        *big.Int
	LiquidationBonus            *big.Int
	InterestRateStrategyAddress common.Address
	UsageAsCollateralEnabled    bool
	BorrowingEnabled            bool
	StableBorrowRateEnabled     bool
	IsActive                    bool
}, error) {
	return _Aave.Contract.GetReserveConfigurationData(&_Aave.CallOpts, _reserve)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address _reserve) view returns(uint256 totalLiquidity, uint256 availableLiquidity, uint256 totalBorrowsStable, uint256 totalBorrowsVariable, uint256 liquidityRate, uint256 variableBorrowRate, uint256 stableBorrowRate, uint256 averageStableBorrowRate, uint256 utilizationRate, uint256 liquidityIndex, uint256 variableBorrowIndex, address aTokenAddress, uint40 lastUpdateTimestamp)
func (_Aave *AaveCaller) GetReserveData(opts *bind.CallOpts, _reserve common.Address) (struct {
	TotalLiquidity          *big.Int
	AvailableLiquidity      *big.Int
	TotalBorrowsStable      *big.Int
	TotalBorrowsVariable    *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	UtilizationRate         *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	ATokenAddress           common.Address
	LastUpdateTimestamp     *big.Int
}, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserveData", _reserve)

	outstruct := new(struct {
		TotalLiquidity          *big.Int
		AvailableLiquidity      *big.Int
		TotalBorrowsStable      *big.Int
		TotalBorrowsVariable    *big.Int
		LiquidityRate           *big.Int
		VariableBorrowRate      *big.Int
		StableBorrowRate        *big.Int
		AverageStableBorrowRate *big.Int
		UtilizationRate         *big.Int
		LiquidityIndex          *big.Int
		VariableBorrowIndex     *big.Int
		ATokenAddress           common.Address
		LastUpdateTimestamp     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalLiquidity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AvailableLiquidity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalBorrowsStable = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalBorrowsVariable = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LiquidityRate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.VariableBorrowRate = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.StableBorrowRate = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.AverageStableBorrowRate = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.UtilizationRate = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.LiquidityIndex = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.VariableBorrowIndex = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.ATokenAddress = *abi.ConvertType(out[11], new(common.Address)).(*common.Address)
	outstruct.LastUpdateTimestamp = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address _reserve) view returns(uint256 totalLiquidity, uint256 availableLiquidity, uint256 totalBorrowsStable, uint256 totalBorrowsVariable, uint256 liquidityRate, uint256 variableBorrowRate, uint256 stableBorrowRate, uint256 averageStableBorrowRate, uint256 utilizationRate, uint256 liquidityIndex, uint256 variableBorrowIndex, address aTokenAddress, uint40 lastUpdateTimestamp)
func (_Aave *AaveSession) GetReserveData(_reserve common.Address) (struct {
	TotalLiquidity          *big.Int
	AvailableLiquidity      *big.Int
	TotalBorrowsStable      *big.Int
	TotalBorrowsVariable    *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	UtilizationRate         *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	ATokenAddress           common.Address
	LastUpdateTimestamp     *big.Int
}, error) {
	return _Aave.Contract.GetReserveData(&_Aave.CallOpts, _reserve)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address _reserve) view returns(uint256 totalLiquidity, uint256 availableLiquidity, uint256 totalBorrowsStable, uint256 totalBorrowsVariable, uint256 liquidityRate, uint256 variableBorrowRate, uint256 stableBorrowRate, uint256 averageStableBorrowRate, uint256 utilizationRate, uint256 liquidityIndex, uint256 variableBorrowIndex, address aTokenAddress, uint40 lastUpdateTimestamp)
func (_Aave *AaveCallerSession) GetReserveData(_reserve common.Address) (struct {
	TotalLiquidity          *big.Int
	AvailableLiquidity      *big.Int
	TotalBorrowsStable      *big.Int
	TotalBorrowsVariable    *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	UtilizationRate         *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	ATokenAddress           common.Address
	LastUpdateTimestamp     *big.Int
}, error) {
	return _Aave.Contract.GetReserveData(&_Aave.CallOpts, _reserve)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(address[])
func (_Aave *AaveCaller) GetReserves(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserves")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(address[])
func (_Aave *AaveSession) GetReserves() ([]common.Address, error) {
	return _Aave.Contract.GetReserves(&_Aave.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(address[])
func (_Aave *AaveCallerSession) GetReserves() ([]common.Address, error) {
	return _Aave.Contract.GetReserves(&_Aave.CallOpts)
}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address _user) view returns(uint256 totalLiquidityETH, uint256 totalCollateralETH, uint256 totalBorrowsETH, uint256 totalFeesETH, uint256 availableBorrowsETH, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_Aave *AaveCaller) GetUserAccountData(opts *bind.CallOpts, _user common.Address) (struct {
	TotalLiquidityETH           *big.Int
	TotalCollateralETH          *big.Int
	TotalBorrowsETH             *big.Int
	TotalFeesETH                *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getUserAccountData", _user)

	outstruct := new(struct {
		TotalLiquidityETH           *big.Int
		TotalCollateralETH          *big.Int
		TotalBorrowsETH             *big.Int
		TotalFeesETH                *big.Int
		AvailableBorrowsETH         *big.Int
		CurrentLiquidationThreshold *big.Int
		Ltv                         *big.Int
		HealthFactor                *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalLiquidityETH = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalCollateralETH = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalBorrowsETH = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalFeesETH = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AvailableBorrowsETH = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.CurrentLiquidationThreshold = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Ltv = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.HealthFactor = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address _user) view returns(uint256 totalLiquidityETH, uint256 totalCollateralETH, uint256 totalBorrowsETH, uint256 totalFeesETH, uint256 availableBorrowsETH, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_Aave *AaveSession) GetUserAccountData(_user common.Address) (struct {
	TotalLiquidityETH           *big.Int
	TotalCollateralETH          *big.Int
	TotalBorrowsETH             *big.Int
	TotalFeesETH                *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	return _Aave.Contract.GetUserAccountData(&_Aave.CallOpts, _user)
}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address _user) view returns(uint256 totalLiquidityETH, uint256 totalCollateralETH, uint256 totalBorrowsETH, uint256 totalFeesETH, uint256 availableBorrowsETH, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_Aave *AaveCallerSession) GetUserAccountData(_user common.Address) (struct {
	TotalLiquidityETH           *big.Int
	TotalCollateralETH          *big.Int
	TotalBorrowsETH             *big.Int
	TotalFeesETH                *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	return _Aave.Contract.GetUserAccountData(&_Aave.CallOpts, _user)
}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address _reserve, address _user) view returns(uint256 currentATokenBalance, uint256 currentBorrowBalance, uint256 principalBorrowBalance, uint256 borrowRateMode, uint256 borrowRate, uint256 liquidityRate, uint256 originationFee, uint256 variableBorrowIndex, uint256 lastUpdateTimestamp, bool usageAsCollateralEnabled)
func (_Aave *AaveCaller) GetUserReserveData(opts *bind.CallOpts, _reserve common.Address, _user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentBorrowBalance     *big.Int
	PrincipalBorrowBalance   *big.Int
	BorrowRateMode           *big.Int
	BorrowRate               *big.Int
	LiquidityRate            *big.Int
	OriginationFee           *big.Int
	VariableBorrowIndex      *big.Int
	LastUpdateTimestamp      *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getUserReserveData", _reserve, _user)

	outstruct := new(struct {
		CurrentATokenBalance     *big.Int
		CurrentBorrowBalance     *big.Int
		PrincipalBorrowBalance   *big.Int
		BorrowRateMode           *big.Int
		BorrowRate               *big.Int
		LiquidityRate            *big.Int
		OriginationFee           *big.Int
		VariableBorrowIndex      *big.Int
		LastUpdateTimestamp      *big.Int
		UsageAsCollateralEnabled bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrentATokenBalance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CurrentBorrowBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PrincipalBorrowBalance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BorrowRateMode = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.BorrowRate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.LiquidityRate = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.OriginationFee = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.VariableBorrowIndex = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.LastUpdateTimestamp = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.UsageAsCollateralEnabled = *abi.ConvertType(out[9], new(bool)).(*bool)

	return *outstruct, err

}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address _reserve, address _user) view returns(uint256 currentATokenBalance, uint256 currentBorrowBalance, uint256 principalBorrowBalance, uint256 borrowRateMode, uint256 borrowRate, uint256 liquidityRate, uint256 originationFee, uint256 variableBorrowIndex, uint256 lastUpdateTimestamp, bool usageAsCollateralEnabled)
func (_Aave *AaveSession) GetUserReserveData(_reserve common.Address, _user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentBorrowBalance     *big.Int
	PrincipalBorrowBalance   *big.Int
	BorrowRateMode           *big.Int
	BorrowRate               *big.Int
	LiquidityRate            *big.Int
	OriginationFee           *big.Int
	VariableBorrowIndex      *big.Int
	LastUpdateTimestamp      *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	return _Aave.Contract.GetUserReserveData(&_Aave.CallOpts, _reserve, _user)
}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address _reserve, address _user) view returns(uint256 currentATokenBalance, uint256 currentBorrowBalance, uint256 principalBorrowBalance, uint256 borrowRateMode, uint256 borrowRate, uint256 liquidityRate, uint256 originationFee, uint256 variableBorrowIndex, uint256 lastUpdateTimestamp, bool usageAsCollateralEnabled)
func (_Aave *AaveCallerSession) GetUserReserveData(_reserve common.Address, _user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentBorrowBalance     *big.Int
	PrincipalBorrowBalance   *big.Int
	BorrowRateMode           *big.Int
	BorrowRate               *big.Int
	LiquidityRate            *big.Int
	OriginationFee           *big.Int
	VariableBorrowIndex      *big.Int
	LastUpdateTimestamp      *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	return _Aave.Contract.GetUserReserveData(&_Aave.CallOpts, _reserve, _user)
}

// ParametersProvider is a free data retrieval call binding the contract method 0x58707e06.
//
// Solidity: function parametersProvider() view returns(address)
func (_Aave *AaveCaller) ParametersProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "parametersProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParametersProvider is a free data retrieval call binding the contract method 0x58707e06.
//
// Solidity: function parametersProvider() view returns(address)
func (_Aave *AaveSession) ParametersProvider() (common.Address, error) {
	return _Aave.Contract.ParametersProvider(&_Aave.CallOpts)
}

// ParametersProvider is a free data retrieval call binding the contract method 0x58707e06.
//
// Solidity: function parametersProvider() view returns(address)
func (_Aave *AaveCallerSession) ParametersProvider() (common.Address, error) {
	return _Aave.Contract.ParametersProvider(&_Aave.CallOpts)
}

// Borrow is a paid mutator transaction binding the contract method 0xc858f5f9.
//
// Solidity: function borrow(address _reserve, uint256 _amount, uint256 _interestRateMode, uint16 _referralCode) returns()
func (_Aave *AaveTransactor) Borrow(opts *bind.TransactOpts, _reserve common.Address, _amount *big.Int, _interestRateMode *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "borrow", _reserve, _amount, _interestRateMode, _referralCode)
}

// Borrow is a paid mutator transaction binding the contract method 0xc858f5f9.
//
// Solidity: function borrow(address _reserve, uint256 _amount, uint256 _interestRateMode, uint16 _referralCode) returns()
func (_Aave *AaveSession) Borrow(_reserve common.Address, _amount *big.Int, _interestRateMode *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _Aave.Contract.Borrow(&_Aave.TransactOpts, _reserve, _amount, _interestRateMode, _referralCode)
}

// Borrow is a paid mutator transaction binding the contract method 0xc858f5f9.
//
// Solidity: function borrow(address _reserve, uint256 _amount, uint256 _interestRateMode, uint16 _referralCode) returns()
func (_Aave *AaveTransactorSession) Borrow(_reserve common.Address, _amount *big.Int, _interestRateMode *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _Aave.Contract.Borrow(&_Aave.TransactOpts, _reserve, _amount, _interestRateMode, _referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xd2d0e066.
//
// Solidity: function deposit(address _reserve, uint256 _amount, uint16 _referralCode) payable returns()
func (_Aave *AaveTransactor) Deposit(opts *bind.TransactOpts, _reserve common.Address, _amount *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "deposit", _reserve, _amount, _referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xd2d0e066.
//
// Solidity: function deposit(address _reserve, uint256 _amount, uint16 _referralCode) payable returns()
func (_Aave *AaveSession) Deposit(_reserve common.Address, _amount *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _Aave.Contract.Deposit(&_Aave.TransactOpts, _reserve, _amount, _referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xd2d0e066.
//
// Solidity: function deposit(address _reserve, uint256 _amount, uint16 _referralCode) payable returns()
func (_Aave *AaveTransactorSession) Deposit(_reserve common.Address, _amount *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _Aave.Contract.Deposit(&_Aave.TransactOpts, _reserve, _amount, _referralCode)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address _receiver, address _reserve, uint256 _amount, bytes _params) returns()
func (_Aave *AaveTransactor) FlashLoan(opts *bind.TransactOpts, _receiver common.Address, _reserve common.Address, _amount *big.Int, _params []byte) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "flashLoan", _receiver, _reserve, _amount, _params)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address _receiver, address _reserve, uint256 _amount, bytes _params) returns()
func (_Aave *AaveSession) FlashLoan(_receiver common.Address, _reserve common.Address, _amount *big.Int, _params []byte) (*types.Transaction, error) {
	return _Aave.Contract.FlashLoan(&_Aave.TransactOpts, _receiver, _reserve, _amount, _params)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address _receiver, address _reserve, uint256 _amount, bytes _params) returns()
func (_Aave *AaveTransactorSession) FlashLoan(_receiver common.Address, _reserve common.Address, _amount *big.Int, _params []byte) (*types.Transaction, error) {
	return _Aave.Contract.FlashLoan(&_Aave.TransactOpts, _receiver, _reserve, _amount, _params)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _addressesProvider) returns()
func (_Aave *AaveTransactor) Initialize(opts *bind.TransactOpts, _addressesProvider common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "initialize", _addressesProvider)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _addressesProvider) returns()
func (_Aave *AaveSession) Initialize(_addressesProvider common.Address) (*types.Transaction, error) {
	return _Aave.Contract.Initialize(&_Aave.TransactOpts, _addressesProvider)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _addressesProvider) returns()
func (_Aave *AaveTransactorSession) Initialize(_addressesProvider common.Address) (*types.Transaction, error) {
	return _Aave.Contract.Initialize(&_Aave.TransactOpts, _addressesProvider)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address _collateral, address _reserve, address _user, uint256 _purchaseAmount, bool _receiveAToken) payable returns()
func (_Aave *AaveTransactor) LiquidationCall(opts *bind.TransactOpts, _collateral common.Address, _reserve common.Address, _user common.Address, _purchaseAmount *big.Int, _receiveAToken bool) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "liquidationCall", _collateral, _reserve, _user, _purchaseAmount, _receiveAToken)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address _collateral, address _reserve, address _user, uint256 _purchaseAmount, bool _receiveAToken) payable returns()
func (_Aave *AaveSession) LiquidationCall(_collateral common.Address, _reserve common.Address, _user common.Address, _purchaseAmount *big.Int, _receiveAToken bool) (*types.Transaction, error) {
	return _Aave.Contract.LiquidationCall(&_Aave.TransactOpts, _collateral, _reserve, _user, _purchaseAmount, _receiveAToken)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address _collateral, address _reserve, address _user, uint256 _purchaseAmount, bool _receiveAToken) payable returns()
func (_Aave *AaveTransactorSession) LiquidationCall(_collateral common.Address, _reserve common.Address, _user common.Address, _purchaseAmount *big.Int, _receiveAToken bool) (*types.Transaction, error) {
	return _Aave.Contract.LiquidationCall(&_Aave.TransactOpts, _collateral, _reserve, _user, _purchaseAmount, _receiveAToken)
}

// RebalanceStableBorrowRate is a paid mutator transaction binding the contract method 0xcd112382.
//
// Solidity: function rebalanceStableBorrowRate(address _reserve, address _user) returns()
func (_Aave *AaveTransactor) RebalanceStableBorrowRate(opts *bind.TransactOpts, _reserve common.Address, _user common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "rebalanceStableBorrowRate", _reserve, _user)
}

// RebalanceStableBorrowRate is a paid mutator transaction binding the contract method 0xcd112382.
//
// Solidity: function rebalanceStableBorrowRate(address _reserve, address _user) returns()
func (_Aave *AaveSession) RebalanceStableBorrowRate(_reserve common.Address, _user common.Address) (*types.Transaction, error) {
	return _Aave.Contract.RebalanceStableBorrowRate(&_Aave.TransactOpts, _reserve, _user)
}

// RebalanceStableBorrowRate is a paid mutator transaction binding the contract method 0xcd112382.
//
// Solidity: function rebalanceStableBorrowRate(address _reserve, address _user) returns()
func (_Aave *AaveTransactorSession) RebalanceStableBorrowRate(_reserve common.Address, _user common.Address) (*types.Transaction, error) {
	return _Aave.Contract.RebalanceStableBorrowRate(&_Aave.TransactOpts, _reserve, _user)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x9895e3d8.
//
// Solidity: function redeemUnderlying(address _reserve, address _user, uint256 _amount, uint256 _aTokenBalanceAfterRedeem) returns()
func (_Aave *AaveTransactor) RedeemUnderlying(opts *bind.TransactOpts, _reserve common.Address, _user common.Address, _amount *big.Int, _aTokenBalanceAfterRedeem *big.Int) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "redeemUnderlying", _reserve, _user, _amount, _aTokenBalanceAfterRedeem)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x9895e3d8.
//
// Solidity: function redeemUnderlying(address _reserve, address _user, uint256 _amount, uint256 _aTokenBalanceAfterRedeem) returns()
func (_Aave *AaveSession) RedeemUnderlying(_reserve common.Address, _user common.Address, _amount *big.Int, _aTokenBalanceAfterRedeem *big.Int) (*types.Transaction, error) {
	return _Aave.Contract.RedeemUnderlying(&_Aave.TransactOpts, _reserve, _user, _amount, _aTokenBalanceAfterRedeem)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x9895e3d8.
//
// Solidity: function redeemUnderlying(address _reserve, address _user, uint256 _amount, uint256 _aTokenBalanceAfterRedeem) returns()
func (_Aave *AaveTransactorSession) RedeemUnderlying(_reserve common.Address, _user common.Address, _amount *big.Int, _aTokenBalanceAfterRedeem *big.Int) (*types.Transaction, error) {
	return _Aave.Contract.RedeemUnderlying(&_Aave.TransactOpts, _reserve, _user, _amount, _aTokenBalanceAfterRedeem)
}

// Repay is a paid mutator transaction binding the contract method 0x5ceae9c4.
//
// Solidity: function repay(address _reserve, uint256 _amount, address _onBehalfOf) payable returns()
func (_Aave *AaveTransactor) Repay(opts *bind.TransactOpts, _reserve common.Address, _amount *big.Int, _onBehalfOf common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "repay", _reserve, _amount, _onBehalfOf)
}

// Repay is a paid mutator transaction binding the contract method 0x5ceae9c4.
//
// Solidity: function repay(address _reserve, uint256 _amount, address _onBehalfOf) payable returns()
func (_Aave *AaveSession) Repay(_reserve common.Address, _amount *big.Int, _onBehalfOf common.Address) (*types.Transaction, error) {
	return _Aave.Contract.Repay(&_Aave.TransactOpts, _reserve, _amount, _onBehalfOf)
}

// Repay is a paid mutator transaction binding the contract method 0x5ceae9c4.
//
// Solidity: function repay(address _reserve, uint256 _amount, address _onBehalfOf) payable returns()
func (_Aave *AaveTransactorSession) Repay(_reserve common.Address, _amount *big.Int, _onBehalfOf common.Address) (*types.Transaction, error) {
	return _Aave.Contract.Repay(&_Aave.TransactOpts, _reserve, _amount, _onBehalfOf)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address _reserve, bool _useAsCollateral) returns()
func (_Aave *AaveTransactor) SetUserUseReserveAsCollateral(opts *bind.TransactOpts, _reserve common.Address, _useAsCollateral bool) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "setUserUseReserveAsCollateral", _reserve, _useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address _reserve, bool _useAsCollateral) returns()
func (_Aave *AaveSession) SetUserUseReserveAsCollateral(_reserve common.Address, _useAsCollateral bool) (*types.Transaction, error) {
	return _Aave.Contract.SetUserUseReserveAsCollateral(&_Aave.TransactOpts, _reserve, _useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address _reserve, bool _useAsCollateral) returns()
func (_Aave *AaveTransactorSession) SetUserUseReserveAsCollateral(_reserve common.Address, _useAsCollateral bool) (*types.Transaction, error) {
	return _Aave.Contract.SetUserUseReserveAsCollateral(&_Aave.TransactOpts, _reserve, _useAsCollateral)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x48ca1300.
//
// Solidity: function swapBorrowRateMode(address _reserve) returns()
func (_Aave *AaveTransactor) SwapBorrowRateMode(opts *bind.TransactOpts, _reserve common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "swapBorrowRateMode", _reserve)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x48ca1300.
//
// Solidity: function swapBorrowRateMode(address _reserve) returns()
func (_Aave *AaveSession) SwapBorrowRateMode(_reserve common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SwapBorrowRateMode(&_Aave.TransactOpts, _reserve)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x48ca1300.
//
// Solidity: function swapBorrowRateMode(address _reserve) returns()
func (_Aave *AaveTransactorSession) SwapBorrowRateMode(_reserve common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SwapBorrowRateMode(&_Aave.TransactOpts, _reserve)
}

// AaveBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the Aave contract.
type AaveBorrowIterator struct {
	Event *AaveBorrow // Event containing the contract specifics and raw log

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
func (it *AaveBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveBorrow)
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
		it.Event = new(AaveBorrow)
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
func (it *AaveBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveBorrow represents a Borrow event raised by the Aave contract.
type AaveBorrow struct {
	Reserve               common.Address
	User                  common.Address
	Amount                *big.Int
	BorrowRateMode        *big.Int
	BorrowRate            *big.Int
	OriginationFee        *big.Int
	BorrowBalanceIncrease *big.Int
	Referral              uint16
	Timestamp             *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x1e77446728e5558aa1b7e81e0cdab9cc1b075ba893b740600c76a315c2caa553.
//
// Solidity: event Borrow(address indexed _reserve, address indexed _user, uint256 _amount, uint256 _borrowRateMode, uint256 _borrowRate, uint256 _originationFee, uint256 _borrowBalanceIncrease, uint16 indexed _referral, uint256 _timestamp)
func (_Aave *AaveFilterer) FilterBorrow(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address, _referral []uint16) (*AaveBorrowIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	var _referralRule []interface{}
	for _, _referralItem := range _referral {
		_referralRule = append(_referralRule, _referralItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "Borrow", _reserveRule, _userRule, _referralRule)
	if err != nil {
		return nil, err
	}
	return &AaveBorrowIterator{contract: _Aave.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x1e77446728e5558aa1b7e81e0cdab9cc1b075ba893b740600c76a315c2caa553.
//
// Solidity: event Borrow(address indexed _reserve, address indexed _user, uint256 _amount, uint256 _borrowRateMode, uint256 _borrowRate, uint256 _originationFee, uint256 _borrowBalanceIncrease, uint16 indexed _referral, uint256 _timestamp)
func (_Aave *AaveFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *AaveBorrow, _reserve []common.Address, _user []common.Address, _referral []uint16) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	var _referralRule []interface{}
	for _, _referralItem := range _referral {
		_referralRule = append(_referralRule, _referralItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "Borrow", _reserveRule, _userRule, _referralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveBorrow)
				if err := _Aave.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x1e77446728e5558aa1b7e81e0cdab9cc1b075ba893b740600c76a315c2caa553.
//
// Solidity: event Borrow(address indexed _reserve, address indexed _user, uint256 _amount, uint256 _borrowRateMode, uint256 _borrowRate, uint256 _originationFee, uint256 _borrowBalanceIncrease, uint16 indexed _referral, uint256 _timestamp)
func (_Aave *AaveFilterer) ParseBorrow(log types.Log) (*AaveBorrow, error) {
	event := new(AaveBorrow)
	if err := _Aave.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Aave contract.
type AaveDepositIterator struct {
	Event *AaveDeposit // Event containing the contract specifics and raw log

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
func (it *AaveDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveDeposit)
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
		it.Event = new(AaveDeposit)
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
func (it *AaveDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveDeposit represents a Deposit event raised by the Aave contract.
type AaveDeposit struct {
	Reserve   common.Address
	User      common.Address
	Amount    *big.Int
	Referral  uint16
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xc12c57b1c73a2c3a2ea4613e9476abb3d8d146857aab7329e24243fb59710c82.
//
// Solidity: event Deposit(address indexed _reserve, address indexed _user, uint256 _amount, uint16 indexed _referral, uint256 _timestamp)
func (_Aave *AaveFilterer) FilterDeposit(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address, _referral []uint16) (*AaveDepositIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	var _referralRule []interface{}
	for _, _referralItem := range _referral {
		_referralRule = append(_referralRule, _referralItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "Deposit", _reserveRule, _userRule, _referralRule)
	if err != nil {
		return nil, err
	}
	return &AaveDepositIterator{contract: _Aave.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xc12c57b1c73a2c3a2ea4613e9476abb3d8d146857aab7329e24243fb59710c82.
//
// Solidity: event Deposit(address indexed _reserve, address indexed _user, uint256 _amount, uint16 indexed _referral, uint256 _timestamp)
func (_Aave *AaveFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *AaveDeposit, _reserve []common.Address, _user []common.Address, _referral []uint16) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	var _referralRule []interface{}
	for _, _referralItem := range _referral {
		_referralRule = append(_referralRule, _referralItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "Deposit", _reserveRule, _userRule, _referralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveDeposit)
				if err := _Aave.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xc12c57b1c73a2c3a2ea4613e9476abb3d8d146857aab7329e24243fb59710c82.
//
// Solidity: event Deposit(address indexed _reserve, address indexed _user, uint256 _amount, uint16 indexed _referral, uint256 _timestamp)
func (_Aave *AaveFilterer) ParseDeposit(log types.Log) (*AaveDeposit, error) {
	event := new(AaveDeposit)
	if err := _Aave.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveFlashLoanIterator is returned from FilterFlashLoan and is used to iterate over the raw logs and unpacked data for FlashLoan events raised by the Aave contract.
type AaveFlashLoanIterator struct {
	Event *AaveFlashLoan // Event containing the contract specifics and raw log

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
func (it *AaveFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveFlashLoan)
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
		it.Event = new(AaveFlashLoan)
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
func (it *AaveFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveFlashLoan represents a FlashLoan event raised by the Aave contract.
type AaveFlashLoan struct {
	Target      common.Address
	Reserve     common.Address
	Amount      *big.Int
	TotalFee    *big.Int
	ProtocolFee *big.Int
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFlashLoan is a free log retrieval operation binding the contract event 0x5b8f46461c1dd69fb968f1a003acee221ea3e19540e350233b612ddb43433b55.
//
// Solidity: event FlashLoan(address indexed _target, address indexed _reserve, uint256 _amount, uint256 _totalFee, uint256 _protocolFee, uint256 _timestamp)
func (_Aave *AaveFilterer) FilterFlashLoan(opts *bind.FilterOpts, _target []common.Address, _reserve []common.Address) (*AaveFlashLoanIterator, error) {

	var _targetRule []interface{}
	for _, _targetItem := range _target {
		_targetRule = append(_targetRule, _targetItem)
	}
	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "FlashLoan", _targetRule, _reserveRule)
	if err != nil {
		return nil, err
	}
	return &AaveFlashLoanIterator{contract: _Aave.contract, event: "FlashLoan", logs: logs, sub: sub}, nil
}

// WatchFlashLoan is a free log subscription operation binding the contract event 0x5b8f46461c1dd69fb968f1a003acee221ea3e19540e350233b612ddb43433b55.
//
// Solidity: event FlashLoan(address indexed _target, address indexed _reserve, uint256 _amount, uint256 _totalFee, uint256 _protocolFee, uint256 _timestamp)
func (_Aave *AaveFilterer) WatchFlashLoan(opts *bind.WatchOpts, sink chan<- *AaveFlashLoan, _target []common.Address, _reserve []common.Address) (event.Subscription, error) {

	var _targetRule []interface{}
	for _, _targetItem := range _target {
		_targetRule = append(_targetRule, _targetItem)
	}
	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "FlashLoan", _targetRule, _reserveRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveFlashLoan)
				if err := _Aave.contract.UnpackLog(event, "FlashLoan", log); err != nil {
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

// ParseFlashLoan is a log parse operation binding the contract event 0x5b8f46461c1dd69fb968f1a003acee221ea3e19540e350233b612ddb43433b55.
//
// Solidity: event FlashLoan(address indexed _target, address indexed _reserve, uint256 _amount, uint256 _totalFee, uint256 _protocolFee, uint256 _timestamp)
func (_Aave *AaveFilterer) ParseFlashLoan(log types.Log) (*AaveFlashLoan, error) {
	event := new(AaveFlashLoan)
	if err := _Aave.contract.UnpackLog(event, "FlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveLiquidationCallIterator is returned from FilterLiquidationCall and is used to iterate over the raw logs and unpacked data for LiquidationCall events raised by the Aave contract.
type AaveLiquidationCallIterator struct {
	Event *AaveLiquidationCall // Event containing the contract specifics and raw log

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
func (it *AaveLiquidationCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveLiquidationCall)
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
		it.Event = new(AaveLiquidationCall)
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
func (it *AaveLiquidationCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveLiquidationCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveLiquidationCall represents a LiquidationCall event raised by the Aave contract.
type AaveLiquidationCall struct {
	Collateral                 common.Address
	Reserve                    common.Address
	User                       common.Address
	PurchaseAmount             *big.Int
	LiquidatedCollateralAmount *big.Int
	AccruedBorrowInterest      *big.Int
	Liquidator                 common.Address
	ReceiveAToken              bool
	Timestamp                  *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterLiquidationCall is a free log retrieval operation binding the contract event 0x56864757fd5b1fc9f38f5f3a981cd8ae512ce41b902cf73fc506ee369c6bc237.
//
// Solidity: event LiquidationCall(address indexed _collateral, address indexed _reserve, address indexed _user, uint256 _purchaseAmount, uint256 _liquidatedCollateralAmount, uint256 _accruedBorrowInterest, address _liquidator, bool _receiveAToken, uint256 _timestamp)
func (_Aave *AaveFilterer) FilterLiquidationCall(opts *bind.FilterOpts, _collateral []common.Address, _reserve []common.Address, _user []common.Address) (*AaveLiquidationCallIterator, error) {

	var _collateralRule []interface{}
	for _, _collateralItem := range _collateral {
		_collateralRule = append(_collateralRule, _collateralItem)
	}
	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "LiquidationCall", _collateralRule, _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return &AaveLiquidationCallIterator{contract: _Aave.contract, event: "LiquidationCall", logs: logs, sub: sub}, nil
}

// WatchLiquidationCall is a free log subscription operation binding the contract event 0x56864757fd5b1fc9f38f5f3a981cd8ae512ce41b902cf73fc506ee369c6bc237.
//
// Solidity: event LiquidationCall(address indexed _collateral, address indexed _reserve, address indexed _user, uint256 _purchaseAmount, uint256 _liquidatedCollateralAmount, uint256 _accruedBorrowInterest, address _liquidator, bool _receiveAToken, uint256 _timestamp)
func (_Aave *AaveFilterer) WatchLiquidationCall(opts *bind.WatchOpts, sink chan<- *AaveLiquidationCall, _collateral []common.Address, _reserve []common.Address, _user []common.Address) (event.Subscription, error) {

	var _collateralRule []interface{}
	for _, _collateralItem := range _collateral {
		_collateralRule = append(_collateralRule, _collateralItem)
	}
	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "LiquidationCall", _collateralRule, _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveLiquidationCall)
				if err := _Aave.contract.UnpackLog(event, "LiquidationCall", log); err != nil {
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

// ParseLiquidationCall is a log parse operation binding the contract event 0x56864757fd5b1fc9f38f5f3a981cd8ae512ce41b902cf73fc506ee369c6bc237.
//
// Solidity: event LiquidationCall(address indexed _collateral, address indexed _reserve, address indexed _user, uint256 _purchaseAmount, uint256 _liquidatedCollateralAmount, uint256 _accruedBorrowInterest, address _liquidator, bool _receiveAToken, uint256 _timestamp)
func (_Aave *AaveFilterer) ParseLiquidationCall(log types.Log) (*AaveLiquidationCall, error) {
	event := new(AaveLiquidationCall)
	if err := _Aave.contract.UnpackLog(event, "LiquidationCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveOriginationFeeLiquidatedIterator is returned from FilterOriginationFeeLiquidated and is used to iterate over the raw logs and unpacked data for OriginationFeeLiquidated events raised by the Aave contract.
type AaveOriginationFeeLiquidatedIterator struct {
	Event *AaveOriginationFeeLiquidated // Event containing the contract specifics and raw log

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
func (it *AaveOriginationFeeLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveOriginationFeeLiquidated)
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
		it.Event = new(AaveOriginationFeeLiquidated)
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
func (it *AaveOriginationFeeLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveOriginationFeeLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveOriginationFeeLiquidated represents a OriginationFeeLiquidated event raised by the Aave contract.
type AaveOriginationFeeLiquidated struct {
	Collateral                 common.Address
	Reserve                    common.Address
	User                       common.Address
	FeeLiquidated              *big.Int
	LiquidatedCollateralForFee *big.Int
	Timestamp                  *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterOriginationFeeLiquidated is a free log retrieval operation binding the contract event 0x36ca8b16d61dc13b1062adff83e3778ab92d14f9e35bfe9fd1283e02b13fb0a1.
//
// Solidity: event OriginationFeeLiquidated(address indexed _collateral, address indexed _reserve, address indexed _user, uint256 _feeLiquidated, uint256 _liquidatedCollateralForFee, uint256 _timestamp)
func (_Aave *AaveFilterer) FilterOriginationFeeLiquidated(opts *bind.FilterOpts, _collateral []common.Address, _reserve []common.Address, _user []common.Address) (*AaveOriginationFeeLiquidatedIterator, error) {

	var _collateralRule []interface{}
	for _, _collateralItem := range _collateral {
		_collateralRule = append(_collateralRule, _collateralItem)
	}
	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "OriginationFeeLiquidated", _collateralRule, _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return &AaveOriginationFeeLiquidatedIterator{contract: _Aave.contract, event: "OriginationFeeLiquidated", logs: logs, sub: sub}, nil
}

// WatchOriginationFeeLiquidated is a free log subscription operation binding the contract event 0x36ca8b16d61dc13b1062adff83e3778ab92d14f9e35bfe9fd1283e02b13fb0a1.
//
// Solidity: event OriginationFeeLiquidated(address indexed _collateral, address indexed _reserve, address indexed _user, uint256 _feeLiquidated, uint256 _liquidatedCollateralForFee, uint256 _timestamp)
func (_Aave *AaveFilterer) WatchOriginationFeeLiquidated(opts *bind.WatchOpts, sink chan<- *AaveOriginationFeeLiquidated, _collateral []common.Address, _reserve []common.Address, _user []common.Address) (event.Subscription, error) {

	var _collateralRule []interface{}
	for _, _collateralItem := range _collateral {
		_collateralRule = append(_collateralRule, _collateralItem)
	}
	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "OriginationFeeLiquidated", _collateralRule, _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveOriginationFeeLiquidated)
				if err := _Aave.contract.UnpackLog(event, "OriginationFeeLiquidated", log); err != nil {
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

// ParseOriginationFeeLiquidated is a log parse operation binding the contract event 0x36ca8b16d61dc13b1062adff83e3778ab92d14f9e35bfe9fd1283e02b13fb0a1.
//
// Solidity: event OriginationFeeLiquidated(address indexed _collateral, address indexed _reserve, address indexed _user, uint256 _feeLiquidated, uint256 _liquidatedCollateralForFee, uint256 _timestamp)
func (_Aave *AaveFilterer) ParseOriginationFeeLiquidated(log types.Log) (*AaveOriginationFeeLiquidated, error) {
	event := new(AaveOriginationFeeLiquidated)
	if err := _Aave.contract.UnpackLog(event, "OriginationFeeLiquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveRebalanceStableBorrowRateIterator is returned from FilterRebalanceStableBorrowRate and is used to iterate over the raw logs and unpacked data for RebalanceStableBorrowRate events raised by the Aave contract.
type AaveRebalanceStableBorrowRateIterator struct {
	Event *AaveRebalanceStableBorrowRate // Event containing the contract specifics and raw log

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
func (it *AaveRebalanceStableBorrowRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveRebalanceStableBorrowRate)
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
		it.Event = new(AaveRebalanceStableBorrowRate)
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
func (it *AaveRebalanceStableBorrowRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveRebalanceStableBorrowRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveRebalanceStableBorrowRate represents a RebalanceStableBorrowRate event raised by the Aave contract.
type AaveRebalanceStableBorrowRate struct {
	Reserve               common.Address
	User                  common.Address
	NewStableRate         *big.Int
	BorrowBalanceIncrease *big.Int
	Timestamp             *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRebalanceStableBorrowRate is a free log retrieval operation binding the contract event 0x5050ad184862424ee0852d1838d355ad65bed1e5e6da67ac9a2dac1922677f60.
//
// Solidity: event RebalanceStableBorrowRate(address indexed _reserve, address indexed _user, uint256 _newStableRate, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_Aave *AaveFilterer) FilterRebalanceStableBorrowRate(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address) (*AaveRebalanceStableBorrowRateIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "RebalanceStableBorrowRate", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return &AaveRebalanceStableBorrowRateIterator{contract: _Aave.contract, event: "RebalanceStableBorrowRate", logs: logs, sub: sub}, nil
}

// WatchRebalanceStableBorrowRate is a free log subscription operation binding the contract event 0x5050ad184862424ee0852d1838d355ad65bed1e5e6da67ac9a2dac1922677f60.
//
// Solidity: event RebalanceStableBorrowRate(address indexed _reserve, address indexed _user, uint256 _newStableRate, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_Aave *AaveFilterer) WatchRebalanceStableBorrowRate(opts *bind.WatchOpts, sink chan<- *AaveRebalanceStableBorrowRate, _reserve []common.Address, _user []common.Address) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "RebalanceStableBorrowRate", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveRebalanceStableBorrowRate)
				if err := _Aave.contract.UnpackLog(event, "RebalanceStableBorrowRate", log); err != nil {
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

// ParseRebalanceStableBorrowRate is a log parse operation binding the contract event 0x5050ad184862424ee0852d1838d355ad65bed1e5e6da67ac9a2dac1922677f60.
//
// Solidity: event RebalanceStableBorrowRate(address indexed _reserve, address indexed _user, uint256 _newStableRate, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_Aave *AaveFilterer) ParseRebalanceStableBorrowRate(log types.Log) (*AaveRebalanceStableBorrowRate, error) {
	event := new(AaveRebalanceStableBorrowRate)
	if err := _Aave.contract.UnpackLog(event, "RebalanceStableBorrowRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveRedeemUnderlyingIterator is returned from FilterRedeemUnderlying and is used to iterate over the raw logs and unpacked data for RedeemUnderlying events raised by the Aave contract.
type AaveRedeemUnderlyingIterator struct {
	Event *AaveRedeemUnderlying // Event containing the contract specifics and raw log

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
func (it *AaveRedeemUnderlyingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveRedeemUnderlying)
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
		it.Event = new(AaveRedeemUnderlying)
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
func (it *AaveRedeemUnderlyingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveRedeemUnderlyingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveRedeemUnderlying represents a RedeemUnderlying event raised by the Aave contract.
type AaveRedeemUnderlying struct {
	Reserve   common.Address
	User      common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRedeemUnderlying is a free log retrieval operation binding the contract event 0x9c4ed599cd8555b9c1e8cd7643240d7d71eb76b792948c49fcb4d411f7b6b3c6.
//
// Solidity: event RedeemUnderlying(address indexed _reserve, address indexed _user, uint256 _amount, uint256 _timestamp)
func (_Aave *AaveFilterer) FilterRedeemUnderlying(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address) (*AaveRedeemUnderlyingIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "RedeemUnderlying", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return &AaveRedeemUnderlyingIterator{contract: _Aave.contract, event: "RedeemUnderlying", logs: logs, sub: sub}, nil
}

// WatchRedeemUnderlying is a free log subscription operation binding the contract event 0x9c4ed599cd8555b9c1e8cd7643240d7d71eb76b792948c49fcb4d411f7b6b3c6.
//
// Solidity: event RedeemUnderlying(address indexed _reserve, address indexed _user, uint256 _amount, uint256 _timestamp)
func (_Aave *AaveFilterer) WatchRedeemUnderlying(opts *bind.WatchOpts, sink chan<- *AaveRedeemUnderlying, _reserve []common.Address, _user []common.Address) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "RedeemUnderlying", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveRedeemUnderlying)
				if err := _Aave.contract.UnpackLog(event, "RedeemUnderlying", log); err != nil {
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

// ParseRedeemUnderlying is a log parse operation binding the contract event 0x9c4ed599cd8555b9c1e8cd7643240d7d71eb76b792948c49fcb4d411f7b6b3c6.
//
// Solidity: event RedeemUnderlying(address indexed _reserve, address indexed _user, uint256 _amount, uint256 _timestamp)
func (_Aave *AaveFilterer) ParseRedeemUnderlying(log types.Log) (*AaveRedeemUnderlying, error) {
	event := new(AaveRedeemUnderlying)
	if err := _Aave.contract.UnpackLog(event, "RedeemUnderlying", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the Aave contract.
type AaveRepayIterator struct {
	Event *AaveRepay // Event containing the contract specifics and raw log

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
func (it *AaveRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveRepay)
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
		it.Event = new(AaveRepay)
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
func (it *AaveRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveRepay represents a Repay event raised by the Aave contract.
type AaveRepay struct {
	Reserve               common.Address
	User                  common.Address
	Repayer               common.Address
	AmountMinusFees       *big.Int
	Fees                  *big.Int
	BorrowBalanceIncrease *big.Int
	Timestamp             *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0xb718f0b14f03d8c3adf35b15e3da52421b042ac879e5a689011a8b1e0036773d.
//
// Solidity: event Repay(address indexed _reserve, address indexed _user, address indexed _repayer, uint256 _amountMinusFees, uint256 _fees, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_Aave *AaveFilterer) FilterRepay(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address, _repayer []common.Address) (*AaveRepayIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _repayerRule []interface{}
	for _, _repayerItem := range _repayer {
		_repayerRule = append(_repayerRule, _repayerItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "Repay", _reserveRule, _userRule, _repayerRule)
	if err != nil {
		return nil, err
	}
	return &AaveRepayIterator{contract: _Aave.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0xb718f0b14f03d8c3adf35b15e3da52421b042ac879e5a689011a8b1e0036773d.
//
// Solidity: event Repay(address indexed _reserve, address indexed _user, address indexed _repayer, uint256 _amountMinusFees, uint256 _fees, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_Aave *AaveFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *AaveRepay, _reserve []common.Address, _user []common.Address, _repayer []common.Address) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _repayerRule []interface{}
	for _, _repayerItem := range _repayer {
		_repayerRule = append(_repayerRule, _repayerItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "Repay", _reserveRule, _userRule, _repayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveRepay)
				if err := _Aave.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0xb718f0b14f03d8c3adf35b15e3da52421b042ac879e5a689011a8b1e0036773d.
//
// Solidity: event Repay(address indexed _reserve, address indexed _user, address indexed _repayer, uint256 _amountMinusFees, uint256 _fees, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_Aave *AaveFilterer) ParseRepay(log types.Log) (*AaveRepay, error) {
	event := new(AaveRepay)
	if err := _Aave.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveReserveUsedAsCollateralDisabledIterator is returned from FilterReserveUsedAsCollateralDisabled and is used to iterate over the raw logs and unpacked data for ReserveUsedAsCollateralDisabled events raised by the Aave contract.
type AaveReserveUsedAsCollateralDisabledIterator struct {
	Event *AaveReserveUsedAsCollateralDisabled // Event containing the contract specifics and raw log

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
func (it *AaveReserveUsedAsCollateralDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveReserveUsedAsCollateralDisabled)
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
		it.Event = new(AaveReserveUsedAsCollateralDisabled)
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
func (it *AaveReserveUsedAsCollateralDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveReserveUsedAsCollateralDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveReserveUsedAsCollateralDisabled represents a ReserveUsedAsCollateralDisabled event raised by the Aave contract.
type AaveReserveUsedAsCollateralDisabled struct {
	Reserve common.Address
	User    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReserveUsedAsCollateralDisabled is a free log retrieval operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed _reserve, address indexed _user)
func (_Aave *AaveFilterer) FilterReserveUsedAsCollateralDisabled(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address) (*AaveReserveUsedAsCollateralDisabledIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "ReserveUsedAsCollateralDisabled", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return &AaveReserveUsedAsCollateralDisabledIterator{contract: _Aave.contract, event: "ReserveUsedAsCollateralDisabled", logs: logs, sub: sub}, nil
}

// WatchReserveUsedAsCollateralDisabled is a free log subscription operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed _reserve, address indexed _user)
func (_Aave *AaveFilterer) WatchReserveUsedAsCollateralDisabled(opts *bind.WatchOpts, sink chan<- *AaveReserveUsedAsCollateralDisabled, _reserve []common.Address, _user []common.Address) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "ReserveUsedAsCollateralDisabled", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveReserveUsedAsCollateralDisabled)
				if err := _Aave.contract.UnpackLog(event, "ReserveUsedAsCollateralDisabled", log); err != nil {
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

// ParseReserveUsedAsCollateralDisabled is a log parse operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed _reserve, address indexed _user)
func (_Aave *AaveFilterer) ParseReserveUsedAsCollateralDisabled(log types.Log) (*AaveReserveUsedAsCollateralDisabled, error) {
	event := new(AaveReserveUsedAsCollateralDisabled)
	if err := _Aave.contract.UnpackLog(event, "ReserveUsedAsCollateralDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveReserveUsedAsCollateralEnabledIterator is returned from FilterReserveUsedAsCollateralEnabled and is used to iterate over the raw logs and unpacked data for ReserveUsedAsCollateralEnabled events raised by the Aave contract.
type AaveReserveUsedAsCollateralEnabledIterator struct {
	Event *AaveReserveUsedAsCollateralEnabled // Event containing the contract specifics and raw log

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
func (it *AaveReserveUsedAsCollateralEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveReserveUsedAsCollateralEnabled)
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
		it.Event = new(AaveReserveUsedAsCollateralEnabled)
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
func (it *AaveReserveUsedAsCollateralEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveReserveUsedAsCollateralEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveReserveUsedAsCollateralEnabled represents a ReserveUsedAsCollateralEnabled event raised by the Aave contract.
type AaveReserveUsedAsCollateralEnabled struct {
	Reserve common.Address
	User    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReserveUsedAsCollateralEnabled is a free log retrieval operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed _reserve, address indexed _user)
func (_Aave *AaveFilterer) FilterReserveUsedAsCollateralEnabled(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address) (*AaveReserveUsedAsCollateralEnabledIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "ReserveUsedAsCollateralEnabled", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return &AaveReserveUsedAsCollateralEnabledIterator{contract: _Aave.contract, event: "ReserveUsedAsCollateralEnabled", logs: logs, sub: sub}, nil
}

// WatchReserveUsedAsCollateralEnabled is a free log subscription operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed _reserve, address indexed _user)
func (_Aave *AaveFilterer) WatchReserveUsedAsCollateralEnabled(opts *bind.WatchOpts, sink chan<- *AaveReserveUsedAsCollateralEnabled, _reserve []common.Address, _user []common.Address) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "ReserveUsedAsCollateralEnabled", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveReserveUsedAsCollateralEnabled)
				if err := _Aave.contract.UnpackLog(event, "ReserveUsedAsCollateralEnabled", log); err != nil {
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

// ParseReserveUsedAsCollateralEnabled is a log parse operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed _reserve, address indexed _user)
func (_Aave *AaveFilterer) ParseReserveUsedAsCollateralEnabled(log types.Log) (*AaveReserveUsedAsCollateralEnabled, error) {
	event := new(AaveReserveUsedAsCollateralEnabled)
	if err := _Aave.contract.UnpackLog(event, "ReserveUsedAsCollateralEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Aave contract.
type AaveSwapIterator struct {
	Event *AaveSwap // Event containing the contract specifics and raw log

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
func (it *AaveSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveSwap)
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
		it.Event = new(AaveSwap)
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
func (it *AaveSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveSwap represents a Swap event raised by the Aave contract.
type AaveSwap struct {
	Reserve               common.Address
	User                  common.Address
	NewRateMode           *big.Int
	NewRate               *big.Int
	BorrowBalanceIncrease *big.Int
	Timestamp             *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xb3e2773606abfd36b5bd91394b3a54d1398336c65005baf7bf7a05efeffaf75b.
//
// Solidity: event Swap(address indexed _reserve, address indexed _user, uint256 _newRateMode, uint256 _newRate, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_Aave *AaveFilterer) FilterSwap(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address) (*AaveSwapIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "Swap", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return &AaveSwapIterator{contract: _Aave.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xb3e2773606abfd36b5bd91394b3a54d1398336c65005baf7bf7a05efeffaf75b.
//
// Solidity: event Swap(address indexed _reserve, address indexed _user, uint256 _newRateMode, uint256 _newRate, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_Aave *AaveFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *AaveSwap, _reserve []common.Address, _user []common.Address) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "Swap", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveSwap)
				if err := _Aave.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xb3e2773606abfd36b5bd91394b3a54d1398336c65005baf7bf7a05efeffaf75b.
//
// Solidity: event Swap(address indexed _reserve, address indexed _user, uint256 _newRateMode, uint256 _newRate, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_Aave *AaveFilterer) ParseSwap(log types.Log) (*AaveSwap, error) {
	event := new(AaveSwap)
	if err := _Aave.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
