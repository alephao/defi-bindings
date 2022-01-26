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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableBorrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableBorrowIndex\",\"type\":\"uint256\"}],\"name\":\"ReserveUpdated\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"CORE_REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"activateReserve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addressesProvider\",\"outputs\":[{\"internalType\":\"contractLendingPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"deactivateReserve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"disableBorrowingOnReserve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"disableReserveAsCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"disableReserveStableBorrowRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_stableBorrowRateEnabled\",\"type\":\"bool\"}],\"name\":\"enableBorrowingOnReserve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_baseLTVasCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_liquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_liquidationBonus\",\"type\":\"uint256\"}],\"name\":\"enableReserveAsCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"enableReserveStableBorrowRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"freezeReserve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveATokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveAvailableLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveConfiguration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveCurrentAverageStableBorrowRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveCurrentLiquidityRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveCurrentStableBorrowRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveCurrentVariableBorrowRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveInterestRateStrategyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveIsActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveIsFreezed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveIsStableBorrowRateEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveLastUpdate\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"timestamp\",\"type\":\"uint40\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveLiquidationBonus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveLiquidationThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveLiquidityCumulativeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveNormalizedIncome\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveTotalBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveTotalBorrowsStable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveTotalBorrowsVariable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveTotalLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveUtilizationRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveVariableBorrowsCumulativeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserBasicReserveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserBorrowBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserCurrentBorrowRateMode\",\"outputs\":[{\"internalType\":\"enumCoreLibrary.InterestRateMode\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserCurrentStableBorrowRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserLastUpdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserOriginationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserUnderlyingAssetBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserVariableBorrowCumulativeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_decimals\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_interestRateStrategyAddress\",\"type\":\"address\"}],\"name\":\"initReserve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractLendingPoolAddressesProvider\",\"name\":\"_addressesProvider\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"isReserveBorrowingEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"isReserveUsageAsCollateralEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"isUserAllowedToBorrowAtStable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"isUserUseReserveAsCollateralEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lendingPoolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"}],\"name\":\"liquidateFee\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"refreshConfiguration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserveToRemove\",\"type\":\"address\"}],\"name\":\"removeLastAddedReserve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reservesList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ltv\",\"type\":\"uint256\"}],\"name\":\"setReserveBaseLTVasCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_decimals\",\"type\":\"uint256\"}],\"name\":\"setReserveDecimals\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rateStrategyAddress\",\"type\":\"address\"}],\"name\":\"setReserveInterestRateStrategyAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_bonus\",\"type\":\"uint256\"}],\"name\":\"setReserveLiquidationBonus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"setReserveLiquidationThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_useAsCollateral\",\"type\":\"bool\"}],\"name\":\"setUserUseReserveAsCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"}],\"name\":\"transferToFeeCollectionAddress\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferToReserve\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferToUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"unfreezeReserve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountBorrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_borrowFee\",\"type\":\"uint256\"},{\"internalType\":\"enumCoreLibrary.InterestRateMode\",\"name\":\"_rateMode\",\"type\":\"uint8\"}],\"name\":\"updateStateOnBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isFirstDeposit\",\"type\":\"bool\"}],\"name\":\"updateStateOnDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_availableLiquidityBefore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_income\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_protocolFee\",\"type\":\"uint256\"}],\"name\":\"updateStateOnFlashLoan\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_principalReserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collateralReserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountToLiquidate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_collateralToLiquidate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeLiquidated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_liquidatedCollateralForFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_balanceIncrease\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_liquidatorReceivesAToken\",\"type\":\"bool\"}],\"name\":\"updateStateOnLiquidation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_balanceIncrease\",\"type\":\"uint256\"}],\"name\":\"updateStateOnRebalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountRedeemed\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_userRedeemedEverything\",\"type\":\"bool\"}],\"name\":\"updateStateOnRedeem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_paybackAmountMinusFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_originationFeeRepaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_balanceIncrease\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_repaidWholeLoan\",\"type\":\"bool\"}],\"name\":\"updateStateOnRepay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_principalBorrowBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_compoundedBorrowBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_balanceIncrease\",\"type\":\"uint256\"},{\"internalType\":\"enumCoreLibrary.InterestRateMode\",\"name\":\"_currentRateMode\",\"type\":\"uint8\"}],\"name\":\"updateStateOnSwapRate\",\"outputs\":[{\"internalType\":\"enumCoreLibrary.InterestRateMode\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// COREREVISION is a free data retrieval call binding the contract method 0x0c7de4e9.
//
// Solidity: function CORE_REVISION() view returns(uint256)
func (_Aave *AaveCaller) COREREVISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "CORE_REVISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COREREVISION is a free data retrieval call binding the contract method 0x0c7de4e9.
//
// Solidity: function CORE_REVISION() view returns(uint256)
func (_Aave *AaveSession) COREREVISION() (*big.Int, error) {
	return _Aave.Contract.COREREVISION(&_Aave.CallOpts)
}

// COREREVISION is a free data retrieval call binding the contract method 0x0c7de4e9.
//
// Solidity: function CORE_REVISION() view returns(uint256)
func (_Aave *AaveCallerSession) COREREVISION() (*big.Int, error) {
	return _Aave.Contract.COREREVISION(&_Aave.CallOpts)
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

// GetReserveATokenAddress is a free data retrieval call binding the contract method 0x34b3beee.
//
// Solidity: function getReserveATokenAddress(address _reserve) view returns(address)
func (_Aave *AaveCaller) GetReserveATokenAddress(opts *bind.CallOpts, _reserve common.Address) (common.Address, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserveATokenAddress", _reserve)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReserveATokenAddress is a free data retrieval call binding the contract method 0x34b3beee.
//
// Solidity: function getReserveATokenAddress(address _reserve) view returns(address)
func (_Aave *AaveSession) GetReserveATokenAddress(_reserve common.Address) (common.Address, error) {
	return _Aave.Contract.GetReserveATokenAddress(&_Aave.CallOpts, _reserve)
}

// GetReserveATokenAddress is a free data retrieval call binding the contract method 0x34b3beee.
//
// Solidity: function getReserveATokenAddress(address _reserve) view returns(address)
func (_Aave *AaveCallerSession) GetReserveATokenAddress(_reserve common.Address) (common.Address, error) {
	return _Aave.Contract.GetReserveATokenAddress(&_Aave.CallOpts, _reserve)
}

// GetReserveAvailableLiquidity is a free data retrieval call binding the contract method 0xe2403019.
//
// Solidity: function getReserveAvailableLiquidity(address _reserve) view returns(uint256)
func (_Aave *AaveCaller) GetReserveAvailableLiquidity(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserveAvailableLiquidity", _reserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveAvailableLiquidity is a free data retrieval call binding the contract method 0xe2403019.
//
// Solidity: function getReserveAvailableLiquidity(address _reserve) view returns(uint256)
func (_Aave *AaveSession) GetReserveAvailableLiquidity(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveAvailableLiquidity(&_Aave.CallOpts, _reserve)
}

// GetReserveAvailableLiquidity is a free data retrieval call binding the contract method 0xe2403019.
//
// Solidity: function getReserveAvailableLiquidity(address _reserve) view returns(uint256)
func (_Aave *AaveCallerSession) GetReserveAvailableLiquidity(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveAvailableLiquidity(&_Aave.CallOpts, _reserve)
}

// GetReserveConfiguration is a free data retrieval call binding the contract method 0x5fc526ff.
//
// Solidity: function getReserveConfiguration(address _reserve) view returns(uint256, uint256, uint256, bool)
func (_Aave *AaveCaller) GetReserveConfiguration(opts *bind.CallOpts, _reserve common.Address) (*big.Int, *big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserveConfiguration", _reserve)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(bool)).(*bool)

	return out0, out1, out2, out3, err

}

// GetReserveConfiguration is a free data retrieval call binding the contract method 0x5fc526ff.
//
// Solidity: function getReserveConfiguration(address _reserve) view returns(uint256, uint256, uint256, bool)
func (_Aave *AaveSession) GetReserveConfiguration(_reserve common.Address) (*big.Int, *big.Int, *big.Int, bool, error) {
	return _Aave.Contract.GetReserveConfiguration(&_Aave.CallOpts, _reserve)
}

// GetReserveConfiguration is a free data retrieval call binding the contract method 0x5fc526ff.
//
// Solidity: function getReserveConfiguration(address _reserve) view returns(uint256, uint256, uint256, bool)
func (_Aave *AaveCallerSession) GetReserveConfiguration(_reserve common.Address) (*big.Int, *big.Int, *big.Int, bool, error) {
	return _Aave.Contract.GetReserveConfiguration(&_Aave.CallOpts, _reserve)
}

// GetReserveCurrentAverageStableBorrowRate is a free data retrieval call binding the contract method 0x46bc0f28.
//
// Solidity: function getReserveCurrentAverageStableBorrowRate(address _reserve) view returns(uint256)
func (_Aave *AaveCaller) GetReserveCurrentAverageStableBorrowRate(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserveCurrentAverageStableBorrowRate", _reserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveCurrentAverageStableBorrowRate is a free data retrieval call binding the contract method 0x46bc0f28.
//
// Solidity: function getReserveCurrentAverageStableBorrowRate(address _reserve) view returns(uint256)
func (_Aave *AaveSession) GetReserveCurrentAverageStableBorrowRate(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveCurrentAverageStableBorrowRate(&_Aave.CallOpts, _reserve)
}

// GetReserveCurrentAverageStableBorrowRate is a free data retrieval call binding the contract method 0x46bc0f28.
//
// Solidity: function getReserveCurrentAverageStableBorrowRate(address _reserve) view returns(uint256)
func (_Aave *AaveCallerSession) GetReserveCurrentAverageStableBorrowRate(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveCurrentAverageStableBorrowRate(&_Aave.CallOpts, _reserve)
}

// GetReserveCurrentLiquidityRate is a free data retrieval call binding the contract method 0xc540148e.
//
// Solidity: function getReserveCurrentLiquidityRate(address _reserve) view returns(uint256)
func (_Aave *AaveCaller) GetReserveCurrentLiquidityRate(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserveCurrentLiquidityRate", _reserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveCurrentLiquidityRate is a free data retrieval call binding the contract method 0xc540148e.
//
// Solidity: function getReserveCurrentLiquidityRate(address _reserve) view returns(uint256)
func (_Aave *AaveSession) GetReserveCurrentLiquidityRate(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveCurrentLiquidityRate(&_Aave.CallOpts, _reserve)
}

// GetReserveCurrentLiquidityRate is a free data retrieval call binding the contract method 0xc540148e.
//
// Solidity: function getReserveCurrentLiquidityRate(address _reserve) view returns(uint256)
func (_Aave *AaveCallerSession) GetReserveCurrentLiquidityRate(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveCurrentLiquidityRate(&_Aave.CallOpts, _reserve)
}

// GetReserveCurrentStableBorrowRate is a free data retrieval call binding the contract method 0x88079d88.
//
// Solidity: function getReserveCurrentStableBorrowRate(address _reserve) view returns(uint256)
func (_Aave *AaveCaller) GetReserveCurrentStableBorrowRate(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserveCurrentStableBorrowRate", _reserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveCurrentStableBorrowRate is a free data retrieval call binding the contract method 0x88079d88.
//
// Solidity: function getReserveCurrentStableBorrowRate(address _reserve) view returns(uint256)
func (_Aave *AaveSession) GetReserveCurrentStableBorrowRate(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveCurrentStableBorrowRate(&_Aave.CallOpts, _reserve)
}

// GetReserveCurrentStableBorrowRate is a free data retrieval call binding the contract method 0x88079d88.
//
// Solidity: function getReserveCurrentStableBorrowRate(address _reserve) view returns(uint256)
func (_Aave *AaveCallerSession) GetReserveCurrentStableBorrowRate(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveCurrentStableBorrowRate(&_Aave.CallOpts, _reserve)
}

// GetReserveCurrentVariableBorrowRate is a free data retrieval call binding the contract method 0x906c0a41.
//
// Solidity: function getReserveCurrentVariableBorrowRate(address _reserve) view returns(uint256)
func (_Aave *AaveCaller) GetReserveCurrentVariableBorrowRate(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserveCurrentVariableBorrowRate", _reserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveCurrentVariableBorrowRate is a free data retrieval call binding the contract method 0x906c0a41.
//
// Solidity: function getReserveCurrentVariableBorrowRate(address _reserve) view returns(uint256)
func (_Aave *AaveSession) GetReserveCurrentVariableBorrowRate(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveCurrentVariableBorrowRate(&_Aave.CallOpts, _reserve)
}

// GetReserveCurrentVariableBorrowRate is a free data retrieval call binding the contract method 0x906c0a41.
//
// Solidity: function getReserveCurrentVariableBorrowRate(address _reserve) view returns(uint256)
func (_Aave *AaveCallerSession) GetReserveCurrentVariableBorrowRate(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveCurrentVariableBorrowRate(&_Aave.CallOpts, _reserve)
}

// GetReserveDecimals is a free data retrieval call binding the contract method 0xa2353fdc.
//
// Solidity: function getReserveDecimals(address _reserve) view returns(uint256)
func (_Aave *AaveCaller) GetReserveDecimals(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserveDecimals", _reserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveDecimals is a free data retrieval call binding the contract method 0xa2353fdc.
//
// Solidity: function getReserveDecimals(address _reserve) view returns(uint256)
func (_Aave *AaveSession) GetReserveDecimals(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveDecimals(&_Aave.CallOpts, _reserve)
}

// GetReserveDecimals is a free data retrieval call binding the contract method 0xa2353fdc.
//
// Solidity: function getReserveDecimals(address _reserve) view returns(uint256)
func (_Aave *AaveCallerSession) GetReserveDecimals(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveDecimals(&_Aave.CallOpts, _reserve)
}

// GetReserveInterestRateStrategyAddress is a free data retrieval call binding the contract method 0x6ae14416.
//
// Solidity: function getReserveInterestRateStrategyAddress(address _reserve) view returns(address)
func (_Aave *AaveCaller) GetReserveInterestRateStrategyAddress(opts *bind.CallOpts, _reserve common.Address) (common.Address, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserveInterestRateStrategyAddress", _reserve)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReserveInterestRateStrategyAddress is a free data retrieval call binding the contract method 0x6ae14416.
//
// Solidity: function getReserveInterestRateStrategyAddress(address _reserve) view returns(address)
func (_Aave *AaveSession) GetReserveInterestRateStrategyAddress(_reserve common.Address) (common.Address, error) {
	return _Aave.Contract.GetReserveInterestRateStrategyAddress(&_Aave.CallOpts, _reserve)
}

// GetReserveInterestRateStrategyAddress is a free data retrieval call binding the contract method 0x6ae14416.
//
// Solidity: function getReserveInterestRateStrategyAddress(address _reserve) view returns(address)
func (_Aave *AaveCallerSession) GetReserveInterestRateStrategyAddress(_reserve common.Address) (common.Address, error) {
	return _Aave.Contract.GetReserveInterestRateStrategyAddress(&_Aave.CallOpts, _reserve)
}

// GetReserveIsActive is a free data retrieval call binding the contract method 0x05075d6e.
//
// Solidity: function getReserveIsActive(address _reserve) view returns(bool)
func (_Aave *AaveCaller) GetReserveIsActive(opts *bind.CallOpts, _reserve common.Address) (bool, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserveIsActive", _reserve)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetReserveIsActive is a free data retrieval call binding the contract method 0x05075d6e.
//
// Solidity: function getReserveIsActive(address _reserve) view returns(bool)
func (_Aave *AaveSession) GetReserveIsActive(_reserve common.Address) (bool, error) {
	return _Aave.Contract.GetReserveIsActive(&_Aave.CallOpts, _reserve)
}

// GetReserveIsActive is a free data retrieval call binding the contract method 0x05075d6e.
//
// Solidity: function getReserveIsActive(address _reserve) view returns(bool)
func (_Aave *AaveCallerSession) GetReserveIsActive(_reserve common.Address) (bool, error) {
	return _Aave.Contract.GetReserveIsActive(&_Aave.CallOpts, _reserve)
}

// GetReserveIsFreezed is a free data retrieval call binding the contract method 0x4a08accb.
//
// Solidity: function getReserveIsFreezed(address _reserve) view returns(bool)
func (_Aave *AaveCaller) GetReserveIsFreezed(opts *bind.CallOpts, _reserve common.Address) (bool, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserveIsFreezed", _reserve)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetReserveIsFreezed is a free data retrieval call binding the contract method 0x4a08accb.
//
// Solidity: function getReserveIsFreezed(address _reserve) view returns(bool)
func (_Aave *AaveSession) GetReserveIsFreezed(_reserve common.Address) (bool, error) {
	return _Aave.Contract.GetReserveIsFreezed(&_Aave.CallOpts, _reserve)
}

// GetReserveIsFreezed is a free data retrieval call binding the contract method 0x4a08accb.
//
// Solidity: function getReserveIsFreezed(address _reserve) view returns(bool)
func (_Aave *AaveCallerSession) GetReserveIsFreezed(_reserve common.Address) (bool, error) {
	return _Aave.Contract.GetReserveIsFreezed(&_Aave.CallOpts, _reserve)
}

// GetReserveIsStableBorrowRateEnabled is a free data retrieval call binding the contract method 0x9e674848.
//
// Solidity: function getReserveIsStableBorrowRateEnabled(address _reserve) view returns(bool)
func (_Aave *AaveCaller) GetReserveIsStableBorrowRateEnabled(opts *bind.CallOpts, _reserve common.Address) (bool, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserveIsStableBorrowRateEnabled", _reserve)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetReserveIsStableBorrowRateEnabled is a free data retrieval call binding the contract method 0x9e674848.
//
// Solidity: function getReserveIsStableBorrowRateEnabled(address _reserve) view returns(bool)
func (_Aave *AaveSession) GetReserveIsStableBorrowRateEnabled(_reserve common.Address) (bool, error) {
	return _Aave.Contract.GetReserveIsStableBorrowRateEnabled(&_Aave.CallOpts, _reserve)
}

// GetReserveIsStableBorrowRateEnabled is a free data retrieval call binding the contract method 0x9e674848.
//
// Solidity: function getReserveIsStableBorrowRateEnabled(address _reserve) view returns(bool)
func (_Aave *AaveCallerSession) GetReserveIsStableBorrowRateEnabled(_reserve common.Address) (bool, error) {
	return _Aave.Contract.GetReserveIsStableBorrowRateEnabled(&_Aave.CallOpts, _reserve)
}

// GetReserveLastUpdate is a free data retrieval call binding the contract method 0x4f144609.
//
// Solidity: function getReserveLastUpdate(address _reserve) view returns(uint40 timestamp)
func (_Aave *AaveCaller) GetReserveLastUpdate(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserveLastUpdate", _reserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveLastUpdate is a free data retrieval call binding the contract method 0x4f144609.
//
// Solidity: function getReserveLastUpdate(address _reserve) view returns(uint40 timestamp)
func (_Aave *AaveSession) GetReserveLastUpdate(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveLastUpdate(&_Aave.CallOpts, _reserve)
}

// GetReserveLastUpdate is a free data retrieval call binding the contract method 0x4f144609.
//
// Solidity: function getReserveLastUpdate(address _reserve) view returns(uint40 timestamp)
func (_Aave *AaveCallerSession) GetReserveLastUpdate(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveLastUpdate(&_Aave.CallOpts, _reserve)
}

// GetReserveLiquidationBonus is a free data retrieval call binding the contract method 0xc76a6c9c.
//
// Solidity: function getReserveLiquidationBonus(address _reserve) view returns(uint256)
func (_Aave *AaveCaller) GetReserveLiquidationBonus(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserveLiquidationBonus", _reserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveLiquidationBonus is a free data retrieval call binding the contract method 0xc76a6c9c.
//
// Solidity: function getReserveLiquidationBonus(address _reserve) view returns(uint256)
func (_Aave *AaveSession) GetReserveLiquidationBonus(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveLiquidationBonus(&_Aave.CallOpts, _reserve)
}

// GetReserveLiquidationBonus is a free data retrieval call binding the contract method 0xc76a6c9c.
//
// Solidity: function getReserveLiquidationBonus(address _reserve) view returns(uint256)
func (_Aave *AaveCallerSession) GetReserveLiquidationBonus(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveLiquidationBonus(&_Aave.CallOpts, _reserve)
}

// GetReserveLiquidationThreshold is a free data retrieval call binding the contract method 0xf054ab46.
//
// Solidity: function getReserveLiquidationThreshold(address _reserve) view returns(uint256)
func (_Aave *AaveCaller) GetReserveLiquidationThreshold(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserveLiquidationThreshold", _reserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveLiquidationThreshold is a free data retrieval call binding the contract method 0xf054ab46.
//
// Solidity: function getReserveLiquidationThreshold(address _reserve) view returns(uint256)
func (_Aave *AaveSession) GetReserveLiquidationThreshold(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveLiquidationThreshold(&_Aave.CallOpts, _reserve)
}

// GetReserveLiquidationThreshold is a free data retrieval call binding the contract method 0xf054ab46.
//
// Solidity: function getReserveLiquidationThreshold(address _reserve) view returns(uint256)
func (_Aave *AaveCallerSession) GetReserveLiquidationThreshold(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveLiquidationThreshold(&_Aave.CallOpts, _reserve)
}

// GetReserveLiquidityCumulativeIndex is a free data retrieval call binding the contract method 0xbd7fd79a.
//
// Solidity: function getReserveLiquidityCumulativeIndex(address _reserve) view returns(uint256)
func (_Aave *AaveCaller) GetReserveLiquidityCumulativeIndex(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserveLiquidityCumulativeIndex", _reserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveLiquidityCumulativeIndex is a free data retrieval call binding the contract method 0xbd7fd79a.
//
// Solidity: function getReserveLiquidityCumulativeIndex(address _reserve) view returns(uint256)
func (_Aave *AaveSession) GetReserveLiquidityCumulativeIndex(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveLiquidityCumulativeIndex(&_Aave.CallOpts, _reserve)
}

// GetReserveLiquidityCumulativeIndex is a free data retrieval call binding the contract method 0xbd7fd79a.
//
// Solidity: function getReserveLiquidityCumulativeIndex(address _reserve) view returns(uint256)
func (_Aave *AaveCallerSession) GetReserveLiquidityCumulativeIndex(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveLiquidityCumulativeIndex(&_Aave.CallOpts, _reserve)
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address _reserve) view returns(uint256)
func (_Aave *AaveCaller) GetReserveNormalizedIncome(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserveNormalizedIncome", _reserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address _reserve) view returns(uint256)
func (_Aave *AaveSession) GetReserveNormalizedIncome(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveNormalizedIncome(&_Aave.CallOpts, _reserve)
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address _reserve) view returns(uint256)
func (_Aave *AaveCallerSession) GetReserveNormalizedIncome(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveNormalizedIncome(&_Aave.CallOpts, _reserve)
}

// GetReserveTotalBorrows is a free data retrieval call binding the contract method 0xe6d18190.
//
// Solidity: function getReserveTotalBorrows(address _reserve) view returns(uint256)
func (_Aave *AaveCaller) GetReserveTotalBorrows(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserveTotalBorrows", _reserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveTotalBorrows is a free data retrieval call binding the contract method 0xe6d18190.
//
// Solidity: function getReserveTotalBorrows(address _reserve) view returns(uint256)
func (_Aave *AaveSession) GetReserveTotalBorrows(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveTotalBorrows(&_Aave.CallOpts, _reserve)
}

// GetReserveTotalBorrows is a free data retrieval call binding the contract method 0xe6d18190.
//
// Solidity: function getReserveTotalBorrows(address _reserve) view returns(uint256)
func (_Aave *AaveCallerSession) GetReserveTotalBorrows(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveTotalBorrows(&_Aave.CallOpts, _reserve)
}

// GetReserveTotalBorrowsStable is a free data retrieval call binding the contract method 0x7f90fec5.
//
// Solidity: function getReserveTotalBorrowsStable(address _reserve) view returns(uint256)
func (_Aave *AaveCaller) GetReserveTotalBorrowsStable(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserveTotalBorrowsStable", _reserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveTotalBorrowsStable is a free data retrieval call binding the contract method 0x7f90fec5.
//
// Solidity: function getReserveTotalBorrowsStable(address _reserve) view returns(uint256)
func (_Aave *AaveSession) GetReserveTotalBorrowsStable(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveTotalBorrowsStable(&_Aave.CallOpts, _reserve)
}

// GetReserveTotalBorrowsStable is a free data retrieval call binding the contract method 0x7f90fec5.
//
// Solidity: function getReserveTotalBorrowsStable(address _reserve) view returns(uint256)
func (_Aave *AaveCallerSession) GetReserveTotalBorrowsStable(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveTotalBorrowsStable(&_Aave.CallOpts, _reserve)
}

// GetReserveTotalBorrowsVariable is a free data retrieval call binding the contract method 0x98bd4737.
//
// Solidity: function getReserveTotalBorrowsVariable(address _reserve) view returns(uint256)
func (_Aave *AaveCaller) GetReserveTotalBorrowsVariable(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserveTotalBorrowsVariable", _reserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveTotalBorrowsVariable is a free data retrieval call binding the contract method 0x98bd4737.
//
// Solidity: function getReserveTotalBorrowsVariable(address _reserve) view returns(uint256)
func (_Aave *AaveSession) GetReserveTotalBorrowsVariable(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveTotalBorrowsVariable(&_Aave.CallOpts, _reserve)
}

// GetReserveTotalBorrowsVariable is a free data retrieval call binding the contract method 0x98bd4737.
//
// Solidity: function getReserveTotalBorrowsVariable(address _reserve) view returns(uint256)
func (_Aave *AaveCallerSession) GetReserveTotalBorrowsVariable(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveTotalBorrowsVariable(&_Aave.CallOpts, _reserve)
}

// GetReserveTotalLiquidity is a free data retrieval call binding the contract method 0xc33cfd90.
//
// Solidity: function getReserveTotalLiquidity(address _reserve) view returns(uint256)
func (_Aave *AaveCaller) GetReserveTotalLiquidity(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserveTotalLiquidity", _reserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveTotalLiquidity is a free data retrieval call binding the contract method 0xc33cfd90.
//
// Solidity: function getReserveTotalLiquidity(address _reserve) view returns(uint256)
func (_Aave *AaveSession) GetReserveTotalLiquidity(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveTotalLiquidity(&_Aave.CallOpts, _reserve)
}

// GetReserveTotalLiquidity is a free data retrieval call binding the contract method 0xc33cfd90.
//
// Solidity: function getReserveTotalLiquidity(address _reserve) view returns(uint256)
func (_Aave *AaveCallerSession) GetReserveTotalLiquidity(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveTotalLiquidity(&_Aave.CallOpts, _reserve)
}

// GetReserveUtilizationRate is a free data retrieval call binding the contract method 0xbfacad84.
//
// Solidity: function getReserveUtilizationRate(address _reserve) view returns(uint256)
func (_Aave *AaveCaller) GetReserveUtilizationRate(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserveUtilizationRate", _reserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveUtilizationRate is a free data retrieval call binding the contract method 0xbfacad84.
//
// Solidity: function getReserveUtilizationRate(address _reserve) view returns(uint256)
func (_Aave *AaveSession) GetReserveUtilizationRate(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveUtilizationRate(&_Aave.CallOpts, _reserve)
}

// GetReserveUtilizationRate is a free data retrieval call binding the contract method 0xbfacad84.
//
// Solidity: function getReserveUtilizationRate(address _reserve) view returns(uint256)
func (_Aave *AaveCallerSession) GetReserveUtilizationRate(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveUtilizationRate(&_Aave.CallOpts, _reserve)
}

// GetReserveVariableBorrowsCumulativeIndex is a free data retrieval call binding the contract method 0xb701d093.
//
// Solidity: function getReserveVariableBorrowsCumulativeIndex(address _reserve) view returns(uint256)
func (_Aave *AaveCaller) GetReserveVariableBorrowsCumulativeIndex(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getReserveVariableBorrowsCumulativeIndex", _reserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveVariableBorrowsCumulativeIndex is a free data retrieval call binding the contract method 0xb701d093.
//
// Solidity: function getReserveVariableBorrowsCumulativeIndex(address _reserve) view returns(uint256)
func (_Aave *AaveSession) GetReserveVariableBorrowsCumulativeIndex(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveVariableBorrowsCumulativeIndex(&_Aave.CallOpts, _reserve)
}

// GetReserveVariableBorrowsCumulativeIndex is a free data retrieval call binding the contract method 0xb701d093.
//
// Solidity: function getReserveVariableBorrowsCumulativeIndex(address _reserve) view returns(uint256)
func (_Aave *AaveCallerSession) GetReserveVariableBorrowsCumulativeIndex(_reserve common.Address) (*big.Int, error) {
	return _Aave.Contract.GetReserveVariableBorrowsCumulativeIndex(&_Aave.CallOpts, _reserve)
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

// GetUserBasicReserveData is a free data retrieval call binding the contract method 0xe10076ad.
//
// Solidity: function getUserBasicReserveData(address _reserve, address _user) view returns(uint256, uint256, uint256, bool)
func (_Aave *AaveCaller) GetUserBasicReserveData(opts *bind.CallOpts, _reserve common.Address, _user common.Address) (*big.Int, *big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getUserBasicReserveData", _reserve, _user)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(bool)).(*bool)

	return out0, out1, out2, out3, err

}

// GetUserBasicReserveData is a free data retrieval call binding the contract method 0xe10076ad.
//
// Solidity: function getUserBasicReserveData(address _reserve, address _user) view returns(uint256, uint256, uint256, bool)
func (_Aave *AaveSession) GetUserBasicReserveData(_reserve common.Address, _user common.Address) (*big.Int, *big.Int, *big.Int, bool, error) {
	return _Aave.Contract.GetUserBasicReserveData(&_Aave.CallOpts, _reserve, _user)
}

// GetUserBasicReserveData is a free data retrieval call binding the contract method 0xe10076ad.
//
// Solidity: function getUserBasicReserveData(address _reserve, address _user) view returns(uint256, uint256, uint256, bool)
func (_Aave *AaveCallerSession) GetUserBasicReserveData(_reserve common.Address, _user common.Address) (*big.Int, *big.Int, *big.Int, bool, error) {
	return _Aave.Contract.GetUserBasicReserveData(&_Aave.CallOpts, _reserve, _user)
}

// GetUserBorrowBalances is a free data retrieval call binding the contract method 0x9fb8afcd.
//
// Solidity: function getUserBorrowBalances(address _reserve, address _user) view returns(uint256, uint256, uint256)
func (_Aave *AaveCaller) GetUserBorrowBalances(opts *bind.CallOpts, _reserve common.Address, _user common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getUserBorrowBalances", _reserve, _user)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetUserBorrowBalances is a free data retrieval call binding the contract method 0x9fb8afcd.
//
// Solidity: function getUserBorrowBalances(address _reserve, address _user) view returns(uint256, uint256, uint256)
func (_Aave *AaveSession) GetUserBorrowBalances(_reserve common.Address, _user common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Aave.Contract.GetUserBorrowBalances(&_Aave.CallOpts, _reserve, _user)
}

// GetUserBorrowBalances is a free data retrieval call binding the contract method 0x9fb8afcd.
//
// Solidity: function getUserBorrowBalances(address _reserve, address _user) view returns(uint256, uint256, uint256)
func (_Aave *AaveCallerSession) GetUserBorrowBalances(_reserve common.Address, _user common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Aave.Contract.GetUserBorrowBalances(&_Aave.CallOpts, _reserve, _user)
}

// GetUserCurrentBorrowRateMode is a free data retrieval call binding the contract method 0x1ca19f19.
//
// Solidity: function getUserCurrentBorrowRateMode(address _reserve, address _user) view returns(uint8)
func (_Aave *AaveCaller) GetUserCurrentBorrowRateMode(opts *bind.CallOpts, _reserve common.Address, _user common.Address) (uint8, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getUserCurrentBorrowRateMode", _reserve, _user)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetUserCurrentBorrowRateMode is a free data retrieval call binding the contract method 0x1ca19f19.
//
// Solidity: function getUserCurrentBorrowRateMode(address _reserve, address _user) view returns(uint8)
func (_Aave *AaveSession) GetUserCurrentBorrowRateMode(_reserve common.Address, _user common.Address) (uint8, error) {
	return _Aave.Contract.GetUserCurrentBorrowRateMode(&_Aave.CallOpts, _reserve, _user)
}

// GetUserCurrentBorrowRateMode is a free data retrieval call binding the contract method 0x1ca19f19.
//
// Solidity: function getUserCurrentBorrowRateMode(address _reserve, address _user) view returns(uint8)
func (_Aave *AaveCallerSession) GetUserCurrentBorrowRateMode(_reserve common.Address, _user common.Address) (uint8, error) {
	return _Aave.Contract.GetUserCurrentBorrowRateMode(&_Aave.CallOpts, _reserve, _user)
}

// GetUserCurrentStableBorrowRate is a free data retrieval call binding the contract method 0x6fffab0c.
//
// Solidity: function getUserCurrentStableBorrowRate(address _reserve, address _user) view returns(uint256)
func (_Aave *AaveCaller) GetUserCurrentStableBorrowRate(opts *bind.CallOpts, _reserve common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getUserCurrentStableBorrowRate", _reserve, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserCurrentStableBorrowRate is a free data retrieval call binding the contract method 0x6fffab0c.
//
// Solidity: function getUserCurrentStableBorrowRate(address _reserve, address _user) view returns(uint256)
func (_Aave *AaveSession) GetUserCurrentStableBorrowRate(_reserve common.Address, _user common.Address) (*big.Int, error) {
	return _Aave.Contract.GetUserCurrentStableBorrowRate(&_Aave.CallOpts, _reserve, _user)
}

// GetUserCurrentStableBorrowRate is a free data retrieval call binding the contract method 0x6fffab0c.
//
// Solidity: function getUserCurrentStableBorrowRate(address _reserve, address _user) view returns(uint256)
func (_Aave *AaveCallerSession) GetUserCurrentStableBorrowRate(_reserve common.Address, _user common.Address) (*big.Int, error) {
	return _Aave.Contract.GetUserCurrentStableBorrowRate(&_Aave.CallOpts, _reserve, _user)
}

// GetUserLastUpdate is a free data retrieval call binding the contract method 0x66d103f3.
//
// Solidity: function getUserLastUpdate(address _reserve, address _user) view returns(uint256 timestamp)
func (_Aave *AaveCaller) GetUserLastUpdate(opts *bind.CallOpts, _reserve common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getUserLastUpdate", _reserve, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserLastUpdate is a free data retrieval call binding the contract method 0x66d103f3.
//
// Solidity: function getUserLastUpdate(address _reserve, address _user) view returns(uint256 timestamp)
func (_Aave *AaveSession) GetUserLastUpdate(_reserve common.Address, _user common.Address) (*big.Int, error) {
	return _Aave.Contract.GetUserLastUpdate(&_Aave.CallOpts, _reserve, _user)
}

// GetUserLastUpdate is a free data retrieval call binding the contract method 0x66d103f3.
//
// Solidity: function getUserLastUpdate(address _reserve, address _user) view returns(uint256 timestamp)
func (_Aave *AaveCallerSession) GetUserLastUpdate(_reserve common.Address, _user common.Address) (*big.Int, error) {
	return _Aave.Contract.GetUserLastUpdate(&_Aave.CallOpts, _reserve, _user)
}

// GetUserOriginationFee is a free data retrieval call binding the contract method 0xfeab31ac.
//
// Solidity: function getUserOriginationFee(address _reserve, address _user) view returns(uint256)
func (_Aave *AaveCaller) GetUserOriginationFee(opts *bind.CallOpts, _reserve common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getUserOriginationFee", _reserve, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserOriginationFee is a free data retrieval call binding the contract method 0xfeab31ac.
//
// Solidity: function getUserOriginationFee(address _reserve, address _user) view returns(uint256)
func (_Aave *AaveSession) GetUserOriginationFee(_reserve common.Address, _user common.Address) (*big.Int, error) {
	return _Aave.Contract.GetUserOriginationFee(&_Aave.CallOpts, _reserve, _user)
}

// GetUserOriginationFee is a free data retrieval call binding the contract method 0xfeab31ac.
//
// Solidity: function getUserOriginationFee(address _reserve, address _user) view returns(uint256)
func (_Aave *AaveCallerSession) GetUserOriginationFee(_reserve common.Address, _user common.Address) (*big.Int, error) {
	return _Aave.Contract.GetUserOriginationFee(&_Aave.CallOpts, _reserve, _user)
}

// GetUserUnderlyingAssetBalance is a free data retrieval call binding the contract method 0x18a4dbca.
//
// Solidity: function getUserUnderlyingAssetBalance(address _reserve, address _user) view returns(uint256)
func (_Aave *AaveCaller) GetUserUnderlyingAssetBalance(opts *bind.CallOpts, _reserve common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getUserUnderlyingAssetBalance", _reserve, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserUnderlyingAssetBalance is a free data retrieval call binding the contract method 0x18a4dbca.
//
// Solidity: function getUserUnderlyingAssetBalance(address _reserve, address _user) view returns(uint256)
func (_Aave *AaveSession) GetUserUnderlyingAssetBalance(_reserve common.Address, _user common.Address) (*big.Int, error) {
	return _Aave.Contract.GetUserUnderlyingAssetBalance(&_Aave.CallOpts, _reserve, _user)
}

// GetUserUnderlyingAssetBalance is a free data retrieval call binding the contract method 0x18a4dbca.
//
// Solidity: function getUserUnderlyingAssetBalance(address _reserve, address _user) view returns(uint256)
func (_Aave *AaveCallerSession) GetUserUnderlyingAssetBalance(_reserve common.Address, _user common.Address) (*big.Int, error) {
	return _Aave.Contract.GetUserUnderlyingAssetBalance(&_Aave.CallOpts, _reserve, _user)
}

// GetUserVariableBorrowCumulativeIndex is a free data retrieval call binding the contract method 0xf6ea8d76.
//
// Solidity: function getUserVariableBorrowCumulativeIndex(address _reserve, address _user) view returns(uint256)
func (_Aave *AaveCaller) GetUserVariableBorrowCumulativeIndex(opts *bind.CallOpts, _reserve common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "getUserVariableBorrowCumulativeIndex", _reserve, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserVariableBorrowCumulativeIndex is a free data retrieval call binding the contract method 0xf6ea8d76.
//
// Solidity: function getUserVariableBorrowCumulativeIndex(address _reserve, address _user) view returns(uint256)
func (_Aave *AaveSession) GetUserVariableBorrowCumulativeIndex(_reserve common.Address, _user common.Address) (*big.Int, error) {
	return _Aave.Contract.GetUserVariableBorrowCumulativeIndex(&_Aave.CallOpts, _reserve, _user)
}

// GetUserVariableBorrowCumulativeIndex is a free data retrieval call binding the contract method 0xf6ea8d76.
//
// Solidity: function getUserVariableBorrowCumulativeIndex(address _reserve, address _user) view returns(uint256)
func (_Aave *AaveCallerSession) GetUserVariableBorrowCumulativeIndex(_reserve common.Address, _user common.Address) (*big.Int, error) {
	return _Aave.Contract.GetUserVariableBorrowCumulativeIndex(&_Aave.CallOpts, _reserve, _user)
}

// IsReserveBorrowingEnabled is a free data retrieval call binding the contract method 0x5cf2e656.
//
// Solidity: function isReserveBorrowingEnabled(address _reserve) view returns(bool)
func (_Aave *AaveCaller) IsReserveBorrowingEnabled(opts *bind.CallOpts, _reserve common.Address) (bool, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "isReserveBorrowingEnabled", _reserve)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReserveBorrowingEnabled is a free data retrieval call binding the contract method 0x5cf2e656.
//
// Solidity: function isReserveBorrowingEnabled(address _reserve) view returns(bool)
func (_Aave *AaveSession) IsReserveBorrowingEnabled(_reserve common.Address) (bool, error) {
	return _Aave.Contract.IsReserveBorrowingEnabled(&_Aave.CallOpts, _reserve)
}

// IsReserveBorrowingEnabled is a free data retrieval call binding the contract method 0x5cf2e656.
//
// Solidity: function isReserveBorrowingEnabled(address _reserve) view returns(bool)
func (_Aave *AaveCallerSession) IsReserveBorrowingEnabled(_reserve common.Address) (bool, error) {
	return _Aave.Contract.IsReserveBorrowingEnabled(&_Aave.CallOpts, _reserve)
}

// IsReserveUsageAsCollateralEnabled is a free data retrieval call binding the contract method 0x18f9bbae.
//
// Solidity: function isReserveUsageAsCollateralEnabled(address _reserve) view returns(bool)
func (_Aave *AaveCaller) IsReserveUsageAsCollateralEnabled(opts *bind.CallOpts, _reserve common.Address) (bool, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "isReserveUsageAsCollateralEnabled", _reserve)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReserveUsageAsCollateralEnabled is a free data retrieval call binding the contract method 0x18f9bbae.
//
// Solidity: function isReserveUsageAsCollateralEnabled(address _reserve) view returns(bool)
func (_Aave *AaveSession) IsReserveUsageAsCollateralEnabled(_reserve common.Address) (bool, error) {
	return _Aave.Contract.IsReserveUsageAsCollateralEnabled(&_Aave.CallOpts, _reserve)
}

// IsReserveUsageAsCollateralEnabled is a free data retrieval call binding the contract method 0x18f9bbae.
//
// Solidity: function isReserveUsageAsCollateralEnabled(address _reserve) view returns(bool)
func (_Aave *AaveCallerSession) IsReserveUsageAsCollateralEnabled(_reserve common.Address) (bool, error) {
	return _Aave.Contract.IsReserveUsageAsCollateralEnabled(&_Aave.CallOpts, _reserve)
}

// IsUserAllowedToBorrowAtStable is a free data retrieval call binding the contract method 0xe2174d86.
//
// Solidity: function isUserAllowedToBorrowAtStable(address _reserve, address _user, uint256 _amount) view returns(bool)
func (_Aave *AaveCaller) IsUserAllowedToBorrowAtStable(opts *bind.CallOpts, _reserve common.Address, _user common.Address, _amount *big.Int) (bool, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "isUserAllowedToBorrowAtStable", _reserve, _user, _amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUserAllowedToBorrowAtStable is a free data retrieval call binding the contract method 0xe2174d86.
//
// Solidity: function isUserAllowedToBorrowAtStable(address _reserve, address _user, uint256 _amount) view returns(bool)
func (_Aave *AaveSession) IsUserAllowedToBorrowAtStable(_reserve common.Address, _user common.Address, _amount *big.Int) (bool, error) {
	return _Aave.Contract.IsUserAllowedToBorrowAtStable(&_Aave.CallOpts, _reserve, _user, _amount)
}

// IsUserAllowedToBorrowAtStable is a free data retrieval call binding the contract method 0xe2174d86.
//
// Solidity: function isUserAllowedToBorrowAtStable(address _reserve, address _user, uint256 _amount) view returns(bool)
func (_Aave *AaveCallerSession) IsUserAllowedToBorrowAtStable(_reserve common.Address, _user common.Address, _amount *big.Int) (bool, error) {
	return _Aave.Contract.IsUserAllowedToBorrowAtStable(&_Aave.CallOpts, _reserve, _user, _amount)
}

// IsUserUseReserveAsCollateralEnabled is a free data retrieval call binding the contract method 0x9e3c4f3b.
//
// Solidity: function isUserUseReserveAsCollateralEnabled(address _reserve, address _user) view returns(bool)
func (_Aave *AaveCaller) IsUserUseReserveAsCollateralEnabled(opts *bind.CallOpts, _reserve common.Address, _user common.Address) (bool, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "isUserUseReserveAsCollateralEnabled", _reserve, _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUserUseReserveAsCollateralEnabled is a free data retrieval call binding the contract method 0x9e3c4f3b.
//
// Solidity: function isUserUseReserveAsCollateralEnabled(address _reserve, address _user) view returns(bool)
func (_Aave *AaveSession) IsUserUseReserveAsCollateralEnabled(_reserve common.Address, _user common.Address) (bool, error) {
	return _Aave.Contract.IsUserUseReserveAsCollateralEnabled(&_Aave.CallOpts, _reserve, _user)
}

// IsUserUseReserveAsCollateralEnabled is a free data retrieval call binding the contract method 0x9e3c4f3b.
//
// Solidity: function isUserUseReserveAsCollateralEnabled(address _reserve, address _user) view returns(bool)
func (_Aave *AaveCallerSession) IsUserUseReserveAsCollateralEnabled(_reserve common.Address, _user common.Address) (bool, error) {
	return _Aave.Contract.IsUserUseReserveAsCollateralEnabled(&_Aave.CallOpts, _reserve, _user)
}

// LendingPoolAddress is a free data retrieval call binding the contract method 0xd3ae26b3.
//
// Solidity: function lendingPoolAddress() view returns(address)
func (_Aave *AaveCaller) LendingPoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "lendingPoolAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LendingPoolAddress is a free data retrieval call binding the contract method 0xd3ae26b3.
//
// Solidity: function lendingPoolAddress() view returns(address)
func (_Aave *AaveSession) LendingPoolAddress() (common.Address, error) {
	return _Aave.Contract.LendingPoolAddress(&_Aave.CallOpts)
}

// LendingPoolAddress is a free data retrieval call binding the contract method 0xd3ae26b3.
//
// Solidity: function lendingPoolAddress() view returns(address)
func (_Aave *AaveCallerSession) LendingPoolAddress() (common.Address, error) {
	return _Aave.Contract.LendingPoolAddress(&_Aave.CallOpts)
}

// ReservesList is a free data retrieval call binding the contract method 0x4fe7a6e5.
//
// Solidity: function reservesList(uint256 ) view returns(address)
func (_Aave *AaveCaller) ReservesList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Aave.contract.Call(opts, &out, "reservesList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReservesList is a free data retrieval call binding the contract method 0x4fe7a6e5.
//
// Solidity: function reservesList(uint256 ) view returns(address)
func (_Aave *AaveSession) ReservesList(arg0 *big.Int) (common.Address, error) {
	return _Aave.Contract.ReservesList(&_Aave.CallOpts, arg0)
}

// ReservesList is a free data retrieval call binding the contract method 0x4fe7a6e5.
//
// Solidity: function reservesList(uint256 ) view returns(address)
func (_Aave *AaveCallerSession) ReservesList(arg0 *big.Int) (common.Address, error) {
	return _Aave.Contract.ReservesList(&_Aave.CallOpts, arg0)
}

// ActivateReserve is a paid mutator transaction binding the contract method 0xb75d6f34.
//
// Solidity: function activateReserve(address _reserve) returns()
func (_Aave *AaveTransactor) ActivateReserve(opts *bind.TransactOpts, _reserve common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "activateReserve", _reserve)
}

// ActivateReserve is a paid mutator transaction binding the contract method 0xb75d6f34.
//
// Solidity: function activateReserve(address _reserve) returns()
func (_Aave *AaveSession) ActivateReserve(_reserve common.Address) (*types.Transaction, error) {
	return _Aave.Contract.ActivateReserve(&_Aave.TransactOpts, _reserve)
}

// ActivateReserve is a paid mutator transaction binding the contract method 0xb75d6f34.
//
// Solidity: function activateReserve(address _reserve) returns()
func (_Aave *AaveTransactorSession) ActivateReserve(_reserve common.Address) (*types.Transaction, error) {
	return _Aave.Contract.ActivateReserve(&_Aave.TransactOpts, _reserve)
}

// DeactivateReserve is a paid mutator transaction binding the contract method 0x3e72a454.
//
// Solidity: function deactivateReserve(address _reserve) returns()
func (_Aave *AaveTransactor) DeactivateReserve(opts *bind.TransactOpts, _reserve common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "deactivateReserve", _reserve)
}

// DeactivateReserve is a paid mutator transaction binding the contract method 0x3e72a454.
//
// Solidity: function deactivateReserve(address _reserve) returns()
func (_Aave *AaveSession) DeactivateReserve(_reserve common.Address) (*types.Transaction, error) {
	return _Aave.Contract.DeactivateReserve(&_Aave.TransactOpts, _reserve)
}

// DeactivateReserve is a paid mutator transaction binding the contract method 0x3e72a454.
//
// Solidity: function deactivateReserve(address _reserve) returns()
func (_Aave *AaveTransactorSession) DeactivateReserve(_reserve common.Address) (*types.Transaction, error) {
	return _Aave.Contract.DeactivateReserve(&_Aave.TransactOpts, _reserve)
}

// DisableBorrowingOnReserve is a paid mutator transaction binding the contract method 0xa8dc0f45.
//
// Solidity: function disableBorrowingOnReserve(address _reserve) returns()
func (_Aave *AaveTransactor) DisableBorrowingOnReserve(opts *bind.TransactOpts, _reserve common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "disableBorrowingOnReserve", _reserve)
}

// DisableBorrowingOnReserve is a paid mutator transaction binding the contract method 0xa8dc0f45.
//
// Solidity: function disableBorrowingOnReserve(address _reserve) returns()
func (_Aave *AaveSession) DisableBorrowingOnReserve(_reserve common.Address) (*types.Transaction, error) {
	return _Aave.Contract.DisableBorrowingOnReserve(&_Aave.TransactOpts, _reserve)
}

// DisableBorrowingOnReserve is a paid mutator transaction binding the contract method 0xa8dc0f45.
//
// Solidity: function disableBorrowingOnReserve(address _reserve) returns()
func (_Aave *AaveTransactorSession) DisableBorrowingOnReserve(_reserve common.Address) (*types.Transaction, error) {
	return _Aave.Contract.DisableBorrowingOnReserve(&_Aave.TransactOpts, _reserve)
}

// DisableReserveAsCollateral is a paid mutator transaction binding the contract method 0xe8ae2f5b.
//
// Solidity: function disableReserveAsCollateral(address _reserve) returns()
func (_Aave *AaveTransactor) DisableReserveAsCollateral(opts *bind.TransactOpts, _reserve common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "disableReserveAsCollateral", _reserve)
}

// DisableReserveAsCollateral is a paid mutator transaction binding the contract method 0xe8ae2f5b.
//
// Solidity: function disableReserveAsCollateral(address _reserve) returns()
func (_Aave *AaveSession) DisableReserveAsCollateral(_reserve common.Address) (*types.Transaction, error) {
	return _Aave.Contract.DisableReserveAsCollateral(&_Aave.TransactOpts, _reserve)
}

// DisableReserveAsCollateral is a paid mutator transaction binding the contract method 0xe8ae2f5b.
//
// Solidity: function disableReserveAsCollateral(address _reserve) returns()
func (_Aave *AaveTransactorSession) DisableReserveAsCollateral(_reserve common.Address) (*types.Transaction, error) {
	return _Aave.Contract.DisableReserveAsCollateral(&_Aave.TransactOpts, _reserve)
}

// DisableReserveStableBorrowRate is a paid mutator transaction binding the contract method 0xb8c0f745.
//
// Solidity: function disableReserveStableBorrowRate(address _reserve) returns()
func (_Aave *AaveTransactor) DisableReserveStableBorrowRate(opts *bind.TransactOpts, _reserve common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "disableReserveStableBorrowRate", _reserve)
}

// DisableReserveStableBorrowRate is a paid mutator transaction binding the contract method 0xb8c0f745.
//
// Solidity: function disableReserveStableBorrowRate(address _reserve) returns()
func (_Aave *AaveSession) DisableReserveStableBorrowRate(_reserve common.Address) (*types.Transaction, error) {
	return _Aave.Contract.DisableReserveStableBorrowRate(&_Aave.TransactOpts, _reserve)
}

// DisableReserveStableBorrowRate is a paid mutator transaction binding the contract method 0xb8c0f745.
//
// Solidity: function disableReserveStableBorrowRate(address _reserve) returns()
func (_Aave *AaveTransactorSession) DisableReserveStableBorrowRate(_reserve common.Address) (*types.Transaction, error) {
	return _Aave.Contract.DisableReserveStableBorrowRate(&_Aave.TransactOpts, _reserve)
}

// EnableBorrowingOnReserve is a paid mutator transaction binding the contract method 0xeede87c1.
//
// Solidity: function enableBorrowingOnReserve(address _reserve, bool _stableBorrowRateEnabled) returns()
func (_Aave *AaveTransactor) EnableBorrowingOnReserve(opts *bind.TransactOpts, _reserve common.Address, _stableBorrowRateEnabled bool) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "enableBorrowingOnReserve", _reserve, _stableBorrowRateEnabled)
}

// EnableBorrowingOnReserve is a paid mutator transaction binding the contract method 0xeede87c1.
//
// Solidity: function enableBorrowingOnReserve(address _reserve, bool _stableBorrowRateEnabled) returns()
func (_Aave *AaveSession) EnableBorrowingOnReserve(_reserve common.Address, _stableBorrowRateEnabled bool) (*types.Transaction, error) {
	return _Aave.Contract.EnableBorrowingOnReserve(&_Aave.TransactOpts, _reserve, _stableBorrowRateEnabled)
}

// EnableBorrowingOnReserve is a paid mutator transaction binding the contract method 0xeede87c1.
//
// Solidity: function enableBorrowingOnReserve(address _reserve, bool _stableBorrowRateEnabled) returns()
func (_Aave *AaveTransactorSession) EnableBorrowingOnReserve(_reserve common.Address, _stableBorrowRateEnabled bool) (*types.Transaction, error) {
	return _Aave.Contract.EnableBorrowingOnReserve(&_Aave.TransactOpts, _reserve, _stableBorrowRateEnabled)
}

// EnableReserveAsCollateral is a paid mutator transaction binding the contract method 0xa5bc826c.
//
// Solidity: function enableReserveAsCollateral(address _reserve, uint256 _baseLTVasCollateral, uint256 _liquidationThreshold, uint256 _liquidationBonus) returns()
func (_Aave *AaveTransactor) EnableReserveAsCollateral(opts *bind.TransactOpts, _reserve common.Address, _baseLTVasCollateral *big.Int, _liquidationThreshold *big.Int, _liquidationBonus *big.Int) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "enableReserveAsCollateral", _reserve, _baseLTVasCollateral, _liquidationThreshold, _liquidationBonus)
}

// EnableReserveAsCollateral is a paid mutator transaction binding the contract method 0xa5bc826c.
//
// Solidity: function enableReserveAsCollateral(address _reserve, uint256 _baseLTVasCollateral, uint256 _liquidationThreshold, uint256 _liquidationBonus) returns()
func (_Aave *AaveSession) EnableReserveAsCollateral(_reserve common.Address, _baseLTVasCollateral *big.Int, _liquidationThreshold *big.Int, _liquidationBonus *big.Int) (*types.Transaction, error) {
	return _Aave.Contract.EnableReserveAsCollateral(&_Aave.TransactOpts, _reserve, _baseLTVasCollateral, _liquidationThreshold, _liquidationBonus)
}

// EnableReserveAsCollateral is a paid mutator transaction binding the contract method 0xa5bc826c.
//
// Solidity: function enableReserveAsCollateral(address _reserve, uint256 _baseLTVasCollateral, uint256 _liquidationThreshold, uint256 _liquidationBonus) returns()
func (_Aave *AaveTransactorSession) EnableReserveAsCollateral(_reserve common.Address, _baseLTVasCollateral *big.Int, _liquidationThreshold *big.Int, _liquidationBonus *big.Int) (*types.Transaction, error) {
	return _Aave.Contract.EnableReserveAsCollateral(&_Aave.TransactOpts, _reserve, _baseLTVasCollateral, _liquidationThreshold, _liquidationBonus)
}

// EnableReserveStableBorrowRate is a paid mutator transaction binding the contract method 0xdae4c4e7.
//
// Solidity: function enableReserveStableBorrowRate(address _reserve) returns()
func (_Aave *AaveTransactor) EnableReserveStableBorrowRate(opts *bind.TransactOpts, _reserve common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "enableReserveStableBorrowRate", _reserve)
}

// EnableReserveStableBorrowRate is a paid mutator transaction binding the contract method 0xdae4c4e7.
//
// Solidity: function enableReserveStableBorrowRate(address _reserve) returns()
func (_Aave *AaveSession) EnableReserveStableBorrowRate(_reserve common.Address) (*types.Transaction, error) {
	return _Aave.Contract.EnableReserveStableBorrowRate(&_Aave.TransactOpts, _reserve)
}

// EnableReserveStableBorrowRate is a paid mutator transaction binding the contract method 0xdae4c4e7.
//
// Solidity: function enableReserveStableBorrowRate(address _reserve) returns()
func (_Aave *AaveTransactorSession) EnableReserveStableBorrowRate(_reserve common.Address) (*types.Transaction, error) {
	return _Aave.Contract.EnableReserveStableBorrowRate(&_Aave.TransactOpts, _reserve)
}

// FreezeReserve is a paid mutator transaction binding the contract method 0x7aca76eb.
//
// Solidity: function freezeReserve(address _reserve) returns()
func (_Aave *AaveTransactor) FreezeReserve(opts *bind.TransactOpts, _reserve common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "freezeReserve", _reserve)
}

// FreezeReserve is a paid mutator transaction binding the contract method 0x7aca76eb.
//
// Solidity: function freezeReserve(address _reserve) returns()
func (_Aave *AaveSession) FreezeReserve(_reserve common.Address) (*types.Transaction, error) {
	return _Aave.Contract.FreezeReserve(&_Aave.TransactOpts, _reserve)
}

// FreezeReserve is a paid mutator transaction binding the contract method 0x7aca76eb.
//
// Solidity: function freezeReserve(address _reserve) returns()
func (_Aave *AaveTransactorSession) FreezeReserve(_reserve common.Address) (*types.Transaction, error) {
	return _Aave.Contract.FreezeReserve(&_Aave.TransactOpts, _reserve)
}

// InitReserve is a paid mutator transaction binding the contract method 0x45330a40.
//
// Solidity: function initReserve(address _reserve, address _aTokenAddress, uint256 _decimals, address _interestRateStrategyAddress) returns()
func (_Aave *AaveTransactor) InitReserve(opts *bind.TransactOpts, _reserve common.Address, _aTokenAddress common.Address, _decimals *big.Int, _interestRateStrategyAddress common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "initReserve", _reserve, _aTokenAddress, _decimals, _interestRateStrategyAddress)
}

// InitReserve is a paid mutator transaction binding the contract method 0x45330a40.
//
// Solidity: function initReserve(address _reserve, address _aTokenAddress, uint256 _decimals, address _interestRateStrategyAddress) returns()
func (_Aave *AaveSession) InitReserve(_reserve common.Address, _aTokenAddress common.Address, _decimals *big.Int, _interestRateStrategyAddress common.Address) (*types.Transaction, error) {
	return _Aave.Contract.InitReserve(&_Aave.TransactOpts, _reserve, _aTokenAddress, _decimals, _interestRateStrategyAddress)
}

// InitReserve is a paid mutator transaction binding the contract method 0x45330a40.
//
// Solidity: function initReserve(address _reserve, address _aTokenAddress, uint256 _decimals, address _interestRateStrategyAddress) returns()
func (_Aave *AaveTransactorSession) InitReserve(_reserve common.Address, _aTokenAddress common.Address, _decimals *big.Int, _interestRateStrategyAddress common.Address) (*types.Transaction, error) {
	return _Aave.Contract.InitReserve(&_Aave.TransactOpts, _reserve, _aTokenAddress, _decimals, _interestRateStrategyAddress)
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

// LiquidateFee is a paid mutator transaction binding the contract method 0x8f385c22.
//
// Solidity: function liquidateFee(address _token, uint256 _amount, address _destination) payable returns()
func (_Aave *AaveTransactor) LiquidateFee(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _destination common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "liquidateFee", _token, _amount, _destination)
}

// LiquidateFee is a paid mutator transaction binding the contract method 0x8f385c22.
//
// Solidity: function liquidateFee(address _token, uint256 _amount, address _destination) payable returns()
func (_Aave *AaveSession) LiquidateFee(_token common.Address, _amount *big.Int, _destination common.Address) (*types.Transaction, error) {
	return _Aave.Contract.LiquidateFee(&_Aave.TransactOpts, _token, _amount, _destination)
}

// LiquidateFee is a paid mutator transaction binding the contract method 0x8f385c22.
//
// Solidity: function liquidateFee(address _token, uint256 _amount, address _destination) payable returns()
func (_Aave *AaveTransactorSession) LiquidateFee(_token common.Address, _amount *big.Int, _destination common.Address) (*types.Transaction, error) {
	return _Aave.Contract.LiquidateFee(&_Aave.TransactOpts, _token, _amount, _destination)
}

// RefreshConfiguration is a paid mutator transaction binding the contract method 0x64681083.
//
// Solidity: function refreshConfiguration() returns()
func (_Aave *AaveTransactor) RefreshConfiguration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "refreshConfiguration")
}

// RefreshConfiguration is a paid mutator transaction binding the contract method 0x64681083.
//
// Solidity: function refreshConfiguration() returns()
func (_Aave *AaveSession) RefreshConfiguration() (*types.Transaction, error) {
	return _Aave.Contract.RefreshConfiguration(&_Aave.TransactOpts)
}

// RefreshConfiguration is a paid mutator transaction binding the contract method 0x64681083.
//
// Solidity: function refreshConfiguration() returns()
func (_Aave *AaveTransactorSession) RefreshConfiguration() (*types.Transaction, error) {
	return _Aave.Contract.RefreshConfiguration(&_Aave.TransactOpts)
}

// RemoveLastAddedReserve is a paid mutator transaction binding the contract method 0xd06e2ec1.
//
// Solidity: function removeLastAddedReserve(address _reserveToRemove) returns()
func (_Aave *AaveTransactor) RemoveLastAddedReserve(opts *bind.TransactOpts, _reserveToRemove common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "removeLastAddedReserve", _reserveToRemove)
}

// RemoveLastAddedReserve is a paid mutator transaction binding the contract method 0xd06e2ec1.
//
// Solidity: function removeLastAddedReserve(address _reserveToRemove) returns()
func (_Aave *AaveSession) RemoveLastAddedReserve(_reserveToRemove common.Address) (*types.Transaction, error) {
	return _Aave.Contract.RemoveLastAddedReserve(&_Aave.TransactOpts, _reserveToRemove)
}

// RemoveLastAddedReserve is a paid mutator transaction binding the contract method 0xd06e2ec1.
//
// Solidity: function removeLastAddedReserve(address _reserveToRemove) returns()
func (_Aave *AaveTransactorSession) RemoveLastAddedReserve(_reserveToRemove common.Address) (*types.Transaction, error) {
	return _Aave.Contract.RemoveLastAddedReserve(&_Aave.TransactOpts, _reserveToRemove)
}

// SetReserveBaseLTVasCollateral is a paid mutator transaction binding the contract method 0xd466016f.
//
// Solidity: function setReserveBaseLTVasCollateral(address _reserve, uint256 _ltv) returns()
func (_Aave *AaveTransactor) SetReserveBaseLTVasCollateral(opts *bind.TransactOpts, _reserve common.Address, _ltv *big.Int) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "setReserveBaseLTVasCollateral", _reserve, _ltv)
}

// SetReserveBaseLTVasCollateral is a paid mutator transaction binding the contract method 0xd466016f.
//
// Solidity: function setReserveBaseLTVasCollateral(address _reserve, uint256 _ltv) returns()
func (_Aave *AaveSession) SetReserveBaseLTVasCollateral(_reserve common.Address, _ltv *big.Int) (*types.Transaction, error) {
	return _Aave.Contract.SetReserveBaseLTVasCollateral(&_Aave.TransactOpts, _reserve, _ltv)
}

// SetReserveBaseLTVasCollateral is a paid mutator transaction binding the contract method 0xd466016f.
//
// Solidity: function setReserveBaseLTVasCollateral(address _reserve, uint256 _ltv) returns()
func (_Aave *AaveTransactorSession) SetReserveBaseLTVasCollateral(_reserve common.Address, _ltv *big.Int) (*types.Transaction, error) {
	return _Aave.Contract.SetReserveBaseLTVasCollateral(&_Aave.TransactOpts, _reserve, _ltv)
}

// SetReserveDecimals is a paid mutator transaction binding the contract method 0x66bbd928.
//
// Solidity: function setReserveDecimals(address _reserve, uint256 _decimals) returns()
func (_Aave *AaveTransactor) SetReserveDecimals(opts *bind.TransactOpts, _reserve common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "setReserveDecimals", _reserve, _decimals)
}

// SetReserveDecimals is a paid mutator transaction binding the contract method 0x66bbd928.
//
// Solidity: function setReserveDecimals(address _reserve, uint256 _decimals) returns()
func (_Aave *AaveSession) SetReserveDecimals(_reserve common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _Aave.Contract.SetReserveDecimals(&_Aave.TransactOpts, _reserve, _decimals)
}

// SetReserveDecimals is a paid mutator transaction binding the contract method 0x66bbd928.
//
// Solidity: function setReserveDecimals(address _reserve, uint256 _decimals) returns()
func (_Aave *AaveTransactorSession) SetReserveDecimals(_reserve common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _Aave.Contract.SetReserveDecimals(&_Aave.TransactOpts, _reserve, _decimals)
}

// SetReserveInterestRateStrategyAddress is a paid mutator transaction binding the contract method 0x1d2118f9.
//
// Solidity: function setReserveInterestRateStrategyAddress(address _reserve, address _rateStrategyAddress) returns()
func (_Aave *AaveTransactor) SetReserveInterestRateStrategyAddress(opts *bind.TransactOpts, _reserve common.Address, _rateStrategyAddress common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "setReserveInterestRateStrategyAddress", _reserve, _rateStrategyAddress)
}

// SetReserveInterestRateStrategyAddress is a paid mutator transaction binding the contract method 0x1d2118f9.
//
// Solidity: function setReserveInterestRateStrategyAddress(address _reserve, address _rateStrategyAddress) returns()
func (_Aave *AaveSession) SetReserveInterestRateStrategyAddress(_reserve common.Address, _rateStrategyAddress common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SetReserveInterestRateStrategyAddress(&_Aave.TransactOpts, _reserve, _rateStrategyAddress)
}

// SetReserveInterestRateStrategyAddress is a paid mutator transaction binding the contract method 0x1d2118f9.
//
// Solidity: function setReserveInterestRateStrategyAddress(address _reserve, address _rateStrategyAddress) returns()
func (_Aave *AaveTransactorSession) SetReserveInterestRateStrategyAddress(_reserve common.Address, _rateStrategyAddress common.Address) (*types.Transaction, error) {
	return _Aave.Contract.SetReserveInterestRateStrategyAddress(&_Aave.TransactOpts, _reserve, _rateStrategyAddress)
}

// SetReserveLiquidationBonus is a paid mutator transaction binding the contract method 0x70fb84f4.
//
// Solidity: function setReserveLiquidationBonus(address _reserve, uint256 _bonus) returns()
func (_Aave *AaveTransactor) SetReserveLiquidationBonus(opts *bind.TransactOpts, _reserve common.Address, _bonus *big.Int) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "setReserveLiquidationBonus", _reserve, _bonus)
}

// SetReserveLiquidationBonus is a paid mutator transaction binding the contract method 0x70fb84f4.
//
// Solidity: function setReserveLiquidationBonus(address _reserve, uint256 _bonus) returns()
func (_Aave *AaveSession) SetReserveLiquidationBonus(_reserve common.Address, _bonus *big.Int) (*types.Transaction, error) {
	return _Aave.Contract.SetReserveLiquidationBonus(&_Aave.TransactOpts, _reserve, _bonus)
}

// SetReserveLiquidationBonus is a paid mutator transaction binding the contract method 0x70fb84f4.
//
// Solidity: function setReserveLiquidationBonus(address _reserve, uint256 _bonus) returns()
func (_Aave *AaveTransactorSession) SetReserveLiquidationBonus(_reserve common.Address, _bonus *big.Int) (*types.Transaction, error) {
	return _Aave.Contract.SetReserveLiquidationBonus(&_Aave.TransactOpts, _reserve, _bonus)
}

// SetReserveLiquidationThreshold is a paid mutator transaction binding the contract method 0x3443a14b.
//
// Solidity: function setReserveLiquidationThreshold(address _reserve, uint256 _threshold) returns()
func (_Aave *AaveTransactor) SetReserveLiquidationThreshold(opts *bind.TransactOpts, _reserve common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "setReserveLiquidationThreshold", _reserve, _threshold)
}

// SetReserveLiquidationThreshold is a paid mutator transaction binding the contract method 0x3443a14b.
//
// Solidity: function setReserveLiquidationThreshold(address _reserve, uint256 _threshold) returns()
func (_Aave *AaveSession) SetReserveLiquidationThreshold(_reserve common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Aave.Contract.SetReserveLiquidationThreshold(&_Aave.TransactOpts, _reserve, _threshold)
}

// SetReserveLiquidationThreshold is a paid mutator transaction binding the contract method 0x3443a14b.
//
// Solidity: function setReserveLiquidationThreshold(address _reserve, uint256 _threshold) returns()
func (_Aave *AaveTransactorSession) SetReserveLiquidationThreshold(_reserve common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Aave.Contract.SetReserveLiquidationThreshold(&_Aave.TransactOpts, _reserve, _threshold)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0xfa51854c.
//
// Solidity: function setUserUseReserveAsCollateral(address _reserve, address _user, bool _useAsCollateral) returns()
func (_Aave *AaveTransactor) SetUserUseReserveAsCollateral(opts *bind.TransactOpts, _reserve common.Address, _user common.Address, _useAsCollateral bool) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "setUserUseReserveAsCollateral", _reserve, _user, _useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0xfa51854c.
//
// Solidity: function setUserUseReserveAsCollateral(address _reserve, address _user, bool _useAsCollateral) returns()
func (_Aave *AaveSession) SetUserUseReserveAsCollateral(_reserve common.Address, _user common.Address, _useAsCollateral bool) (*types.Transaction, error) {
	return _Aave.Contract.SetUserUseReserveAsCollateral(&_Aave.TransactOpts, _reserve, _user, _useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0xfa51854c.
//
// Solidity: function setUserUseReserveAsCollateral(address _reserve, address _user, bool _useAsCollateral) returns()
func (_Aave *AaveTransactorSession) SetUserUseReserveAsCollateral(_reserve common.Address, _user common.Address, _useAsCollateral bool) (*types.Transaction, error) {
	return _Aave.Contract.SetUserUseReserveAsCollateral(&_Aave.TransactOpts, _reserve, _user, _useAsCollateral)
}

// TransferToFeeCollectionAddress is a paid mutator transaction binding the contract method 0xc7d14237.
//
// Solidity: function transferToFeeCollectionAddress(address _token, address _user, uint256 _amount, address _destination) payable returns()
func (_Aave *AaveTransactor) TransferToFeeCollectionAddress(opts *bind.TransactOpts, _token common.Address, _user common.Address, _amount *big.Int, _destination common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "transferToFeeCollectionAddress", _token, _user, _amount, _destination)
}

// TransferToFeeCollectionAddress is a paid mutator transaction binding the contract method 0xc7d14237.
//
// Solidity: function transferToFeeCollectionAddress(address _token, address _user, uint256 _amount, address _destination) payable returns()
func (_Aave *AaveSession) TransferToFeeCollectionAddress(_token common.Address, _user common.Address, _amount *big.Int, _destination common.Address) (*types.Transaction, error) {
	return _Aave.Contract.TransferToFeeCollectionAddress(&_Aave.TransactOpts, _token, _user, _amount, _destination)
}

// TransferToFeeCollectionAddress is a paid mutator transaction binding the contract method 0xc7d14237.
//
// Solidity: function transferToFeeCollectionAddress(address _token, address _user, uint256 _amount, address _destination) payable returns()
func (_Aave *AaveTransactorSession) TransferToFeeCollectionAddress(_token common.Address, _user common.Address, _amount *big.Int, _destination common.Address) (*types.Transaction, error) {
	return _Aave.Contract.TransferToFeeCollectionAddress(&_Aave.TransactOpts, _token, _user, _amount, _destination)
}

// TransferToReserve is a paid mutator transaction binding the contract method 0x28fcf4d3.
//
// Solidity: function transferToReserve(address _reserve, address _user, uint256 _amount) payable returns()
func (_Aave *AaveTransactor) TransferToReserve(opts *bind.TransactOpts, _reserve common.Address, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "transferToReserve", _reserve, _user, _amount)
}

// TransferToReserve is a paid mutator transaction binding the contract method 0x28fcf4d3.
//
// Solidity: function transferToReserve(address _reserve, address _user, uint256 _amount) payable returns()
func (_Aave *AaveSession) TransferToReserve(_reserve common.Address, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Aave.Contract.TransferToReserve(&_Aave.TransactOpts, _reserve, _user, _amount)
}

// TransferToReserve is a paid mutator transaction binding the contract method 0x28fcf4d3.
//
// Solidity: function transferToReserve(address _reserve, address _user, uint256 _amount) payable returns()
func (_Aave *AaveTransactorSession) TransferToReserve(_reserve common.Address, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Aave.Contract.TransferToReserve(&_Aave.TransactOpts, _reserve, _user, _amount)
}

// TransferToUser is a paid mutator transaction binding the contract method 0xfa93b2a5.
//
// Solidity: function transferToUser(address _reserve, address _user, uint256 _amount) returns()
func (_Aave *AaveTransactor) TransferToUser(opts *bind.TransactOpts, _reserve common.Address, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "transferToUser", _reserve, _user, _amount)
}

// TransferToUser is a paid mutator transaction binding the contract method 0xfa93b2a5.
//
// Solidity: function transferToUser(address _reserve, address _user, uint256 _amount) returns()
func (_Aave *AaveSession) TransferToUser(_reserve common.Address, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Aave.Contract.TransferToUser(&_Aave.TransactOpts, _reserve, _user, _amount)
}

// TransferToUser is a paid mutator transaction binding the contract method 0xfa93b2a5.
//
// Solidity: function transferToUser(address _reserve, address _user, uint256 _amount) returns()
func (_Aave *AaveTransactorSession) TransferToUser(_reserve common.Address, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Aave.Contract.TransferToUser(&_Aave.TransactOpts, _reserve, _user, _amount)
}

// UnfreezeReserve is a paid mutator transaction binding the contract method 0xef1f9373.
//
// Solidity: function unfreezeReserve(address _reserve) returns()
func (_Aave *AaveTransactor) UnfreezeReserve(opts *bind.TransactOpts, _reserve common.Address) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "unfreezeReserve", _reserve)
}

// UnfreezeReserve is a paid mutator transaction binding the contract method 0xef1f9373.
//
// Solidity: function unfreezeReserve(address _reserve) returns()
func (_Aave *AaveSession) UnfreezeReserve(_reserve common.Address) (*types.Transaction, error) {
	return _Aave.Contract.UnfreezeReserve(&_Aave.TransactOpts, _reserve)
}

// UnfreezeReserve is a paid mutator transaction binding the contract method 0xef1f9373.
//
// Solidity: function unfreezeReserve(address _reserve) returns()
func (_Aave *AaveTransactorSession) UnfreezeReserve(_reserve common.Address) (*types.Transaction, error) {
	return _Aave.Contract.UnfreezeReserve(&_Aave.TransactOpts, _reserve)
}

// UpdateStateOnBorrow is a paid mutator transaction binding the contract method 0x37ac6fe4.
//
// Solidity: function updateStateOnBorrow(address _reserve, address _user, uint256 _amountBorrowed, uint256 _borrowFee, uint8 _rateMode) returns(uint256, uint256)
func (_Aave *AaveTransactor) UpdateStateOnBorrow(opts *bind.TransactOpts, _reserve common.Address, _user common.Address, _amountBorrowed *big.Int, _borrowFee *big.Int, _rateMode uint8) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "updateStateOnBorrow", _reserve, _user, _amountBorrowed, _borrowFee, _rateMode)
}

// UpdateStateOnBorrow is a paid mutator transaction binding the contract method 0x37ac6fe4.
//
// Solidity: function updateStateOnBorrow(address _reserve, address _user, uint256 _amountBorrowed, uint256 _borrowFee, uint8 _rateMode) returns(uint256, uint256)
func (_Aave *AaveSession) UpdateStateOnBorrow(_reserve common.Address, _user common.Address, _amountBorrowed *big.Int, _borrowFee *big.Int, _rateMode uint8) (*types.Transaction, error) {
	return _Aave.Contract.UpdateStateOnBorrow(&_Aave.TransactOpts, _reserve, _user, _amountBorrowed, _borrowFee, _rateMode)
}

// UpdateStateOnBorrow is a paid mutator transaction binding the contract method 0x37ac6fe4.
//
// Solidity: function updateStateOnBorrow(address _reserve, address _user, uint256 _amountBorrowed, uint256 _borrowFee, uint8 _rateMode) returns(uint256, uint256)
func (_Aave *AaveTransactorSession) UpdateStateOnBorrow(_reserve common.Address, _user common.Address, _amountBorrowed *big.Int, _borrowFee *big.Int, _rateMode uint8) (*types.Transaction, error) {
	return _Aave.Contract.UpdateStateOnBorrow(&_Aave.TransactOpts, _reserve, _user, _amountBorrowed, _borrowFee, _rateMode)
}

// UpdateStateOnDeposit is a paid mutator transaction binding the contract method 0xbcd6ffa4.
//
// Solidity: function updateStateOnDeposit(address _reserve, address _user, uint256 _amount, bool _isFirstDeposit) returns()
func (_Aave *AaveTransactor) UpdateStateOnDeposit(opts *bind.TransactOpts, _reserve common.Address, _user common.Address, _amount *big.Int, _isFirstDeposit bool) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "updateStateOnDeposit", _reserve, _user, _amount, _isFirstDeposit)
}

// UpdateStateOnDeposit is a paid mutator transaction binding the contract method 0xbcd6ffa4.
//
// Solidity: function updateStateOnDeposit(address _reserve, address _user, uint256 _amount, bool _isFirstDeposit) returns()
func (_Aave *AaveSession) UpdateStateOnDeposit(_reserve common.Address, _user common.Address, _amount *big.Int, _isFirstDeposit bool) (*types.Transaction, error) {
	return _Aave.Contract.UpdateStateOnDeposit(&_Aave.TransactOpts, _reserve, _user, _amount, _isFirstDeposit)
}

// UpdateStateOnDeposit is a paid mutator transaction binding the contract method 0xbcd6ffa4.
//
// Solidity: function updateStateOnDeposit(address _reserve, address _user, uint256 _amount, bool _isFirstDeposit) returns()
func (_Aave *AaveTransactorSession) UpdateStateOnDeposit(_reserve common.Address, _user common.Address, _amount *big.Int, _isFirstDeposit bool) (*types.Transaction, error) {
	return _Aave.Contract.UpdateStateOnDeposit(&_Aave.TransactOpts, _reserve, _user, _amount, _isFirstDeposit)
}

// UpdateStateOnFlashLoan is a paid mutator transaction binding the contract method 0x09ac2953.
//
// Solidity: function updateStateOnFlashLoan(address _reserve, uint256 _availableLiquidityBefore, uint256 _income, uint256 _protocolFee) returns()
func (_Aave *AaveTransactor) UpdateStateOnFlashLoan(opts *bind.TransactOpts, _reserve common.Address, _availableLiquidityBefore *big.Int, _income *big.Int, _protocolFee *big.Int) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "updateStateOnFlashLoan", _reserve, _availableLiquidityBefore, _income, _protocolFee)
}

// UpdateStateOnFlashLoan is a paid mutator transaction binding the contract method 0x09ac2953.
//
// Solidity: function updateStateOnFlashLoan(address _reserve, uint256 _availableLiquidityBefore, uint256 _income, uint256 _protocolFee) returns()
func (_Aave *AaveSession) UpdateStateOnFlashLoan(_reserve common.Address, _availableLiquidityBefore *big.Int, _income *big.Int, _protocolFee *big.Int) (*types.Transaction, error) {
	return _Aave.Contract.UpdateStateOnFlashLoan(&_Aave.TransactOpts, _reserve, _availableLiquidityBefore, _income, _protocolFee)
}

// UpdateStateOnFlashLoan is a paid mutator transaction binding the contract method 0x09ac2953.
//
// Solidity: function updateStateOnFlashLoan(address _reserve, uint256 _availableLiquidityBefore, uint256 _income, uint256 _protocolFee) returns()
func (_Aave *AaveTransactorSession) UpdateStateOnFlashLoan(_reserve common.Address, _availableLiquidityBefore *big.Int, _income *big.Int, _protocolFee *big.Int) (*types.Transaction, error) {
	return _Aave.Contract.UpdateStateOnFlashLoan(&_Aave.TransactOpts, _reserve, _availableLiquidityBefore, _income, _protocolFee)
}

// UpdateStateOnLiquidation is a paid mutator transaction binding the contract method 0x68beb4d6.
//
// Solidity: function updateStateOnLiquidation(address _principalReserve, address _collateralReserve, address _user, uint256 _amountToLiquidate, uint256 _collateralToLiquidate, uint256 _feeLiquidated, uint256 _liquidatedCollateralForFee, uint256 _balanceIncrease, bool _liquidatorReceivesAToken) returns()
func (_Aave *AaveTransactor) UpdateStateOnLiquidation(opts *bind.TransactOpts, _principalReserve common.Address, _collateralReserve common.Address, _user common.Address, _amountToLiquidate *big.Int, _collateralToLiquidate *big.Int, _feeLiquidated *big.Int, _liquidatedCollateralForFee *big.Int, _balanceIncrease *big.Int, _liquidatorReceivesAToken bool) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "updateStateOnLiquidation", _principalReserve, _collateralReserve, _user, _amountToLiquidate, _collateralToLiquidate, _feeLiquidated, _liquidatedCollateralForFee, _balanceIncrease, _liquidatorReceivesAToken)
}

// UpdateStateOnLiquidation is a paid mutator transaction binding the contract method 0x68beb4d6.
//
// Solidity: function updateStateOnLiquidation(address _principalReserve, address _collateralReserve, address _user, uint256 _amountToLiquidate, uint256 _collateralToLiquidate, uint256 _feeLiquidated, uint256 _liquidatedCollateralForFee, uint256 _balanceIncrease, bool _liquidatorReceivesAToken) returns()
func (_Aave *AaveSession) UpdateStateOnLiquidation(_principalReserve common.Address, _collateralReserve common.Address, _user common.Address, _amountToLiquidate *big.Int, _collateralToLiquidate *big.Int, _feeLiquidated *big.Int, _liquidatedCollateralForFee *big.Int, _balanceIncrease *big.Int, _liquidatorReceivesAToken bool) (*types.Transaction, error) {
	return _Aave.Contract.UpdateStateOnLiquidation(&_Aave.TransactOpts, _principalReserve, _collateralReserve, _user, _amountToLiquidate, _collateralToLiquidate, _feeLiquidated, _liquidatedCollateralForFee, _balanceIncrease, _liquidatorReceivesAToken)
}

// UpdateStateOnLiquidation is a paid mutator transaction binding the contract method 0x68beb4d6.
//
// Solidity: function updateStateOnLiquidation(address _principalReserve, address _collateralReserve, address _user, uint256 _amountToLiquidate, uint256 _collateralToLiquidate, uint256 _feeLiquidated, uint256 _liquidatedCollateralForFee, uint256 _balanceIncrease, bool _liquidatorReceivesAToken) returns()
func (_Aave *AaveTransactorSession) UpdateStateOnLiquidation(_principalReserve common.Address, _collateralReserve common.Address, _user common.Address, _amountToLiquidate *big.Int, _collateralToLiquidate *big.Int, _feeLiquidated *big.Int, _liquidatedCollateralForFee *big.Int, _balanceIncrease *big.Int, _liquidatorReceivesAToken bool) (*types.Transaction, error) {
	return _Aave.Contract.UpdateStateOnLiquidation(&_Aave.TransactOpts, _principalReserve, _collateralReserve, _user, _amountToLiquidate, _collateralToLiquidate, _feeLiquidated, _liquidatedCollateralForFee, _balanceIncrease, _liquidatorReceivesAToken)
}

// UpdateStateOnRebalance is a paid mutator transaction binding the contract method 0xaf825b07.
//
// Solidity: function updateStateOnRebalance(address _reserve, address _user, uint256 _balanceIncrease) returns(uint256)
func (_Aave *AaveTransactor) UpdateStateOnRebalance(opts *bind.TransactOpts, _reserve common.Address, _user common.Address, _balanceIncrease *big.Int) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "updateStateOnRebalance", _reserve, _user, _balanceIncrease)
}

// UpdateStateOnRebalance is a paid mutator transaction binding the contract method 0xaf825b07.
//
// Solidity: function updateStateOnRebalance(address _reserve, address _user, uint256 _balanceIncrease) returns(uint256)
func (_Aave *AaveSession) UpdateStateOnRebalance(_reserve common.Address, _user common.Address, _balanceIncrease *big.Int) (*types.Transaction, error) {
	return _Aave.Contract.UpdateStateOnRebalance(&_Aave.TransactOpts, _reserve, _user, _balanceIncrease)
}

// UpdateStateOnRebalance is a paid mutator transaction binding the contract method 0xaf825b07.
//
// Solidity: function updateStateOnRebalance(address _reserve, address _user, uint256 _balanceIncrease) returns(uint256)
func (_Aave *AaveTransactorSession) UpdateStateOnRebalance(_reserve common.Address, _user common.Address, _balanceIncrease *big.Int) (*types.Transaction, error) {
	return _Aave.Contract.UpdateStateOnRebalance(&_Aave.TransactOpts, _reserve, _user, _balanceIncrease)
}

// UpdateStateOnRedeem is a paid mutator transaction binding the contract method 0xafcdbea3.
//
// Solidity: function updateStateOnRedeem(address _reserve, address _user, uint256 _amountRedeemed, bool _userRedeemedEverything) returns()
func (_Aave *AaveTransactor) UpdateStateOnRedeem(opts *bind.TransactOpts, _reserve common.Address, _user common.Address, _amountRedeemed *big.Int, _userRedeemedEverything bool) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "updateStateOnRedeem", _reserve, _user, _amountRedeemed, _userRedeemedEverything)
}

// UpdateStateOnRedeem is a paid mutator transaction binding the contract method 0xafcdbea3.
//
// Solidity: function updateStateOnRedeem(address _reserve, address _user, uint256 _amountRedeemed, bool _userRedeemedEverything) returns()
func (_Aave *AaveSession) UpdateStateOnRedeem(_reserve common.Address, _user common.Address, _amountRedeemed *big.Int, _userRedeemedEverything bool) (*types.Transaction, error) {
	return _Aave.Contract.UpdateStateOnRedeem(&_Aave.TransactOpts, _reserve, _user, _amountRedeemed, _userRedeemedEverything)
}

// UpdateStateOnRedeem is a paid mutator transaction binding the contract method 0xafcdbea3.
//
// Solidity: function updateStateOnRedeem(address _reserve, address _user, uint256 _amountRedeemed, bool _userRedeemedEverything) returns()
func (_Aave *AaveTransactorSession) UpdateStateOnRedeem(_reserve common.Address, _user common.Address, _amountRedeemed *big.Int, _userRedeemedEverything bool) (*types.Transaction, error) {
	return _Aave.Contract.UpdateStateOnRedeem(&_Aave.TransactOpts, _reserve, _user, _amountRedeemed, _userRedeemedEverything)
}

// UpdateStateOnRepay is a paid mutator transaction binding the contract method 0xda12d96f.
//
// Solidity: function updateStateOnRepay(address _reserve, address _user, uint256 _paybackAmountMinusFees, uint256 _originationFeeRepaid, uint256 _balanceIncrease, bool _repaidWholeLoan) returns()
func (_Aave *AaveTransactor) UpdateStateOnRepay(opts *bind.TransactOpts, _reserve common.Address, _user common.Address, _paybackAmountMinusFees *big.Int, _originationFeeRepaid *big.Int, _balanceIncrease *big.Int, _repaidWholeLoan bool) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "updateStateOnRepay", _reserve, _user, _paybackAmountMinusFees, _originationFeeRepaid, _balanceIncrease, _repaidWholeLoan)
}

// UpdateStateOnRepay is a paid mutator transaction binding the contract method 0xda12d96f.
//
// Solidity: function updateStateOnRepay(address _reserve, address _user, uint256 _paybackAmountMinusFees, uint256 _originationFeeRepaid, uint256 _balanceIncrease, bool _repaidWholeLoan) returns()
func (_Aave *AaveSession) UpdateStateOnRepay(_reserve common.Address, _user common.Address, _paybackAmountMinusFees *big.Int, _originationFeeRepaid *big.Int, _balanceIncrease *big.Int, _repaidWholeLoan bool) (*types.Transaction, error) {
	return _Aave.Contract.UpdateStateOnRepay(&_Aave.TransactOpts, _reserve, _user, _paybackAmountMinusFees, _originationFeeRepaid, _balanceIncrease, _repaidWholeLoan)
}

// UpdateStateOnRepay is a paid mutator transaction binding the contract method 0xda12d96f.
//
// Solidity: function updateStateOnRepay(address _reserve, address _user, uint256 _paybackAmountMinusFees, uint256 _originationFeeRepaid, uint256 _balanceIncrease, bool _repaidWholeLoan) returns()
func (_Aave *AaveTransactorSession) UpdateStateOnRepay(_reserve common.Address, _user common.Address, _paybackAmountMinusFees *big.Int, _originationFeeRepaid *big.Int, _balanceIncrease *big.Int, _repaidWholeLoan bool) (*types.Transaction, error) {
	return _Aave.Contract.UpdateStateOnRepay(&_Aave.TransactOpts, _reserve, _user, _paybackAmountMinusFees, _originationFeeRepaid, _balanceIncrease, _repaidWholeLoan)
}

// UpdateStateOnSwapRate is a paid mutator transaction binding the contract method 0xf6148311.
//
// Solidity: function updateStateOnSwapRate(address _reserve, address _user, uint256 _principalBorrowBalance, uint256 _compoundedBorrowBalance, uint256 _balanceIncrease, uint8 _currentRateMode) returns(uint8, uint256)
func (_Aave *AaveTransactor) UpdateStateOnSwapRate(opts *bind.TransactOpts, _reserve common.Address, _user common.Address, _principalBorrowBalance *big.Int, _compoundedBorrowBalance *big.Int, _balanceIncrease *big.Int, _currentRateMode uint8) (*types.Transaction, error) {
	return _Aave.contract.Transact(opts, "updateStateOnSwapRate", _reserve, _user, _principalBorrowBalance, _compoundedBorrowBalance, _balanceIncrease, _currentRateMode)
}

// UpdateStateOnSwapRate is a paid mutator transaction binding the contract method 0xf6148311.
//
// Solidity: function updateStateOnSwapRate(address _reserve, address _user, uint256 _principalBorrowBalance, uint256 _compoundedBorrowBalance, uint256 _balanceIncrease, uint8 _currentRateMode) returns(uint8, uint256)
func (_Aave *AaveSession) UpdateStateOnSwapRate(_reserve common.Address, _user common.Address, _principalBorrowBalance *big.Int, _compoundedBorrowBalance *big.Int, _balanceIncrease *big.Int, _currentRateMode uint8) (*types.Transaction, error) {
	return _Aave.Contract.UpdateStateOnSwapRate(&_Aave.TransactOpts, _reserve, _user, _principalBorrowBalance, _compoundedBorrowBalance, _balanceIncrease, _currentRateMode)
}

// UpdateStateOnSwapRate is a paid mutator transaction binding the contract method 0xf6148311.
//
// Solidity: function updateStateOnSwapRate(address _reserve, address _user, uint256 _principalBorrowBalance, uint256 _compoundedBorrowBalance, uint256 _balanceIncrease, uint8 _currentRateMode) returns(uint8, uint256)
func (_Aave *AaveTransactorSession) UpdateStateOnSwapRate(_reserve common.Address, _user common.Address, _principalBorrowBalance *big.Int, _compoundedBorrowBalance *big.Int, _balanceIncrease *big.Int, _currentRateMode uint8) (*types.Transaction, error) {
	return _Aave.Contract.UpdateStateOnSwapRate(&_Aave.TransactOpts, _reserve, _user, _principalBorrowBalance, _compoundedBorrowBalance, _balanceIncrease, _currentRateMode)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Aave *AaveTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Aave.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Aave *AaveSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Aave.Contract.Fallback(&_Aave.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Aave *AaveTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Aave.Contract.Fallback(&_Aave.TransactOpts, calldata)
}

// AaveReserveUpdatedIterator is returned from FilterReserveUpdated and is used to iterate over the raw logs and unpacked data for ReserveUpdated events raised by the Aave contract.
type AaveReserveUpdatedIterator struct {
	Event *AaveReserveUpdated // Event containing the contract specifics and raw log

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
func (it *AaveReserveUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveReserveUpdated)
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
		it.Event = new(AaveReserveUpdated)
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
func (it *AaveReserveUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveReserveUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveReserveUpdated represents a ReserveUpdated event raised by the Aave contract.
type AaveReserveUpdated struct {
	Reserve             common.Address
	LiquidityRate       *big.Int
	StableBorrowRate    *big.Int
	VariableBorrowRate  *big.Int
	LiquidityIndex      *big.Int
	VariableBorrowIndex *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReserveUpdated is a free log retrieval operation binding the contract event 0x04e4f521f16fcfd987978b05504262c2a2db223844977aab000e5accedb2d725.
//
// Solidity: event ReserveUpdated(address indexed reserve, uint256 liquidityRate, uint256 stableBorrowRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_Aave *AaveFilterer) FilterReserveUpdated(opts *bind.FilterOpts, reserve []common.Address) (*AaveReserveUpdatedIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _Aave.contract.FilterLogs(opts, "ReserveUpdated", reserveRule)
	if err != nil {
		return nil, err
	}
	return &AaveReserveUpdatedIterator{contract: _Aave.contract, event: "ReserveUpdated", logs: logs, sub: sub}, nil
}

// WatchReserveUpdated is a free log subscription operation binding the contract event 0x04e4f521f16fcfd987978b05504262c2a2db223844977aab000e5accedb2d725.
//
// Solidity: event ReserveUpdated(address indexed reserve, uint256 liquidityRate, uint256 stableBorrowRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_Aave *AaveFilterer) WatchReserveUpdated(opts *bind.WatchOpts, sink chan<- *AaveReserveUpdated, reserve []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _Aave.contract.WatchLogs(opts, "ReserveUpdated", reserveRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveReserveUpdated)
				if err := _Aave.contract.UnpackLog(event, "ReserveUpdated", log); err != nil {
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

// ParseReserveUpdated is a log parse operation binding the contract event 0x04e4f521f16fcfd987978b05504262c2a2db223844977aab000e5accedb2d725.
//
// Solidity: event ReserveUpdated(address indexed reserve, uint256 liquidityRate, uint256 stableBorrowRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_Aave *AaveFilterer) ParseReserveUpdated(log types.Log) (*AaveReserveUpdated, error) {
	event := new(AaveReserveUpdated)
	if err := _Aave.contract.UnpackLog(event, "ReserveUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
