// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package compound

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

// CompoundMetaData contains all meta data concerning the Compound contract.
var CompoundMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pauseState\",\"type\":\"bool\"}],\"name\":\"ActionPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSpeed\",\"type\":\"uint256\"}],\"name\":\"CompSpeedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compBorrowIndex\",\"type\":\"uint256\"}],\"name\":\"DistributedBorrowerComp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"supplier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compSupplyIndex\",\"type\":\"uint256\"}],\"name\":\"DistributedSupplierComp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isComped\",\"type\":\"bool\"}],\"name\":\"MarketComped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MarketEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MarketExited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"MarketListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBorrowCap\",\"type\":\"uint256\"}],\"name\":\"NewBorrowCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldBorrowCapGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newBorrowCapGuardian\",\"type\":\"address\"}],\"name\":\"NewBorrowCapGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCloseFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCloseFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewCloseFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCollateralFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCollateralFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewCollateralFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCompRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCompRate\",\"type\":\"uint256\"}],\"name\":\"NewCompRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLiquidationIncentiveMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLiquidationIncentiveMantissa\",\"type\":\"uint256\"}],\"name\":\"NewLiquidationIncentive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMaxAssets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxAssets\",\"type\":\"uint256\"}],\"name\":\"NewMaxAssets\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPauseGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPauseGuardian\",\"type\":\"address\"}],\"name\":\"NewPauseGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractPriceOracle\",\"name\":\"oldPriceOracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPriceOracle\",\"name\":\"newPriceOracle\",\"type\":\"address\"}],\"name\":\"NewPriceOracle\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cTokens\",\"type\":\"address[]\"}],\"name\":\"_addCompMarkets\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractUnitroller\",\"name\":\"unitroller\",\"type\":\"address\"}],\"name\":\"_become\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_borrowGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"_dropCompMarket\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_mintGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBorrowCapGuardian\",\"type\":\"address\"}],\"name\":\"_setBorrowCapGuardian\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setBorrowPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCloseFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setCloseFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newCollateralFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setCollateralFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"compRate_\",\"type\":\"uint256\"}],\"name\":\"_setCompRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLiquidationIncentiveMantissa\",\"type\":\"uint256\"}],\"name\":\"_setLiquidationIncentive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken[]\",\"name\":\"cTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"newBorrowCaps\",\"type\":\"uint256[]\"}],\"name\":\"_setMarketBorrowCaps\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxAssets\",\"type\":\"uint256\"}],\"name\":\"_setMaxAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setMintPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPauseGuardian\",\"type\":\"address\"}],\"name\":\"_setPauseGuardian\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"newOracle\",\"type\":\"address\"}],\"name\":\"_setPriceOracle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setSeizePaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setTransferPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"_supportMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountAssets\",\"outputs\":[{\"internalType\":\"contractCToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allMarkets\",\"outputs\":[{\"internalType\":\"contractCToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowCapGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowCaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"checkMembership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"claimComp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closeFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compAccrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compBorrowState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compBorrowerIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compClaimThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compInitialIndex\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"\",\"type\":\"uint224\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compSpeeds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compSupplierIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compSupplyState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cTokens\",\"type\":\"address[]\"}],\"name\":\"enterMarkets\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenAddress\",\"type\":\"address\"}],\"name\":\"exitMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllMarkets\",\"outputs\":[{\"internalType\":\"contractCToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAssetsIn\",\"outputs\":[{\"internalType\":\"contractCToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCompAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenModify\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"getHypotheticalAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isComptroller\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateCalculateSeizeTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"liquidationIncentiveMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isListed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"collateralFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isComped\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mintAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualMintAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"mintVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pauseGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingComptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"refreshCompSpeeds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowerIndex\",\"type\":\"uint256\"}],\"name\":\"repayBorrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"seizeGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transferGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CompoundABI is the input ABI used to generate the binding from.
// Deprecated: Use CompoundMetaData.ABI instead.
var CompoundABI = CompoundMetaData.ABI

// Compound is an auto generated Go binding around an Ethereum contract.
type Compound struct {
	CompoundCaller     // Read-only binding to the contract
	CompoundTransactor // Write-only binding to the contract
	CompoundFilterer   // Log filterer for contract events
}

// CompoundCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompoundCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompoundTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompoundFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompoundSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompoundSession struct {
	Contract     *Compound         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CompoundCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompoundCallerSession struct {
	Contract *CompoundCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CompoundTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompoundTransactorSession struct {
	Contract     *CompoundTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CompoundRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompoundRaw struct {
	Contract *Compound // Generic contract binding to access the raw methods on
}

// CompoundCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompoundCallerRaw struct {
	Contract *CompoundCaller // Generic read-only contract binding to access the raw methods on
}

// CompoundTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompoundTransactorRaw struct {
	Contract *CompoundTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompound creates a new instance of Compound, bound to a specific deployed contract.
func NewCompound(address common.Address, backend bind.ContractBackend) (*Compound, error) {
	contract, err := bindCompound(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Compound{CompoundCaller: CompoundCaller{contract: contract}, CompoundTransactor: CompoundTransactor{contract: contract}, CompoundFilterer: CompoundFilterer{contract: contract}}, nil
}

// NewCompoundCaller creates a new read-only instance of Compound, bound to a specific deployed contract.
func NewCompoundCaller(address common.Address, caller bind.ContractCaller) (*CompoundCaller, error) {
	contract, err := bindCompound(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundCaller{contract: contract}, nil
}

// NewCompoundTransactor creates a new write-only instance of Compound, bound to a specific deployed contract.
func NewCompoundTransactor(address common.Address, transactor bind.ContractTransactor) (*CompoundTransactor, error) {
	contract, err := bindCompound(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompoundTransactor{contract: contract}, nil
}

// NewCompoundFilterer creates a new log filterer instance of Compound, bound to a specific deployed contract.
func NewCompoundFilterer(address common.Address, filterer bind.ContractFilterer) (*CompoundFilterer, error) {
	contract, err := bindCompound(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompoundFilterer{contract: contract}, nil
}

// bindCompound binds a generic wrapper to an already deployed contract.
func bindCompound(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CompoundABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compound *CompoundRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Compound.Contract.CompoundCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compound *CompoundRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compound.Contract.CompoundTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compound *CompoundRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compound.Contract.CompoundTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compound *CompoundCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Compound.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compound *CompoundTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compound.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compound *CompoundTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compound.Contract.contract.Transact(opts, method, params...)
}

// PBorrowGuardianPaused is a free data retrieval call binding the contract method 0xe6653f3d.
//
// Solidity: function _borrowGuardianPaused() view returns(bool)
func (_Compound *CompoundCaller) PBorrowGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "_borrowGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PBorrowGuardianPaused is a free data retrieval call binding the contract method 0xe6653f3d.
//
// Solidity: function _borrowGuardianPaused() view returns(bool)
func (_Compound *CompoundSession) PBorrowGuardianPaused() (bool, error) {
	return _Compound.Contract.PBorrowGuardianPaused(&_Compound.CallOpts)
}

// PBorrowGuardianPaused is a free data retrieval call binding the contract method 0xe6653f3d.
//
// Solidity: function _borrowGuardianPaused() view returns(bool)
func (_Compound *CompoundCallerSession) PBorrowGuardianPaused() (bool, error) {
	return _Compound.Contract.PBorrowGuardianPaused(&_Compound.CallOpts)
}

// PMintGuardianPaused is a free data retrieval call binding the contract method 0x3c94786f.
//
// Solidity: function _mintGuardianPaused() view returns(bool)
func (_Compound *CompoundCaller) PMintGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "_mintGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PMintGuardianPaused is a free data retrieval call binding the contract method 0x3c94786f.
//
// Solidity: function _mintGuardianPaused() view returns(bool)
func (_Compound *CompoundSession) PMintGuardianPaused() (bool, error) {
	return _Compound.Contract.PMintGuardianPaused(&_Compound.CallOpts)
}

// PMintGuardianPaused is a free data retrieval call binding the contract method 0x3c94786f.
//
// Solidity: function _mintGuardianPaused() view returns(bool)
func (_Compound *CompoundCallerSession) PMintGuardianPaused() (bool, error) {
	return _Compound.Contract.PMintGuardianPaused(&_Compound.CallOpts)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_Compound *CompoundCaller) AccountAssets(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "accountAssets", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_Compound *CompoundSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Compound.Contract.AccountAssets(&_Compound.CallOpts, arg0, arg1)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_Compound *CompoundCallerSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Compound.Contract.AccountAssets(&_Compound.CallOpts, arg0, arg1)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Compound *CompoundCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Compound *CompoundSession) Admin() (common.Address, error) {
	return _Compound.Contract.Admin(&_Compound.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Compound *CompoundCallerSession) Admin() (common.Address, error) {
	return _Compound.Contract.Admin(&_Compound.CallOpts)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_Compound *CompoundCaller) AllMarkets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "allMarkets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_Compound *CompoundSession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _Compound.Contract.AllMarkets(&_Compound.CallOpts, arg0)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_Compound *CompoundCallerSession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _Compound.Contract.AllMarkets(&_Compound.CallOpts, arg0)
}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_Compound *CompoundCaller) BorrowCapGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "borrowCapGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_Compound *CompoundSession) BorrowCapGuardian() (common.Address, error) {
	return _Compound.Contract.BorrowCapGuardian(&_Compound.CallOpts)
}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_Compound *CompoundCallerSession) BorrowCapGuardian() (common.Address, error) {
	return _Compound.Contract.BorrowCapGuardian(&_Compound.CallOpts)
}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_Compound *CompoundCaller) BorrowCaps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "borrowCaps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_Compound *CompoundSession) BorrowCaps(arg0 common.Address) (*big.Int, error) {
	return _Compound.Contract.BorrowCaps(&_Compound.CallOpts, arg0)
}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_Compound *CompoundCallerSession) BorrowCaps(arg0 common.Address) (*big.Int, error) {
	return _Compound.Contract.BorrowCaps(&_Compound.CallOpts, arg0)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_Compound *CompoundCaller) BorrowGuardianPaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "borrowGuardianPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_Compound *CompoundSession) BorrowGuardianPaused(arg0 common.Address) (bool, error) {
	return _Compound.Contract.BorrowGuardianPaused(&_Compound.CallOpts, arg0)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_Compound *CompoundCallerSession) BorrowGuardianPaused(arg0 common.Address) (bool, error) {
	return _Compound.Contract.BorrowGuardianPaused(&_Compound.CallOpts, arg0)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address cToken) view returns(bool)
func (_Compound *CompoundCaller) CheckMembership(opts *bind.CallOpts, account common.Address, cToken common.Address) (bool, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "checkMembership", account, cToken)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address cToken) view returns(bool)
func (_Compound *CompoundSession) CheckMembership(account common.Address, cToken common.Address) (bool, error) {
	return _Compound.Contract.CheckMembership(&_Compound.CallOpts, account, cToken)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address cToken) view returns(bool)
func (_Compound *CompoundCallerSession) CheckMembership(account common.Address, cToken common.Address) (bool, error) {
	return _Compound.Contract.CheckMembership(&_Compound.CallOpts, account, cToken)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_Compound *CompoundCaller) CloseFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "closeFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_Compound *CompoundSession) CloseFactorMantissa() (*big.Int, error) {
	return _Compound.Contract.CloseFactorMantissa(&_Compound.CallOpts)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_Compound *CompoundCallerSession) CloseFactorMantissa() (*big.Int, error) {
	return _Compound.Contract.CloseFactorMantissa(&_Compound.CallOpts)
}

// CompAccrued is a free data retrieval call binding the contract method 0xcc7ebdc4.
//
// Solidity: function compAccrued(address ) view returns(uint256)
func (_Compound *CompoundCaller) CompAccrued(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "compAccrued", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompAccrued is a free data retrieval call binding the contract method 0xcc7ebdc4.
//
// Solidity: function compAccrued(address ) view returns(uint256)
func (_Compound *CompoundSession) CompAccrued(arg0 common.Address) (*big.Int, error) {
	return _Compound.Contract.CompAccrued(&_Compound.CallOpts, arg0)
}

// CompAccrued is a free data retrieval call binding the contract method 0xcc7ebdc4.
//
// Solidity: function compAccrued(address ) view returns(uint256)
func (_Compound *CompoundCallerSession) CompAccrued(arg0 common.Address) (*big.Int, error) {
	return _Compound.Contract.CompAccrued(&_Compound.CallOpts, arg0)
}

// CompBorrowState is a free data retrieval call binding the contract method 0x8c57804e.
//
// Solidity: function compBorrowState(address ) view returns(uint224 index, uint32 block)
func (_Compound *CompoundCaller) CompBorrowState(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "compBorrowState", arg0)

	outstruct := new(struct {
		Index *big.Int
		Block uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Block = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// CompBorrowState is a free data retrieval call binding the contract method 0x8c57804e.
//
// Solidity: function compBorrowState(address ) view returns(uint224 index, uint32 block)
func (_Compound *CompoundSession) CompBorrowState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Compound.Contract.CompBorrowState(&_Compound.CallOpts, arg0)
}

// CompBorrowState is a free data retrieval call binding the contract method 0x8c57804e.
//
// Solidity: function compBorrowState(address ) view returns(uint224 index, uint32 block)
func (_Compound *CompoundCallerSession) CompBorrowState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Compound.Contract.CompBorrowState(&_Compound.CallOpts, arg0)
}

// CompBorrowerIndex is a free data retrieval call binding the contract method 0xca0af043.
//
// Solidity: function compBorrowerIndex(address , address ) view returns(uint256)
func (_Compound *CompoundCaller) CompBorrowerIndex(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "compBorrowerIndex", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompBorrowerIndex is a free data retrieval call binding the contract method 0xca0af043.
//
// Solidity: function compBorrowerIndex(address , address ) view returns(uint256)
func (_Compound *CompoundSession) CompBorrowerIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Compound.Contract.CompBorrowerIndex(&_Compound.CallOpts, arg0, arg1)
}

// CompBorrowerIndex is a free data retrieval call binding the contract method 0xca0af043.
//
// Solidity: function compBorrowerIndex(address , address ) view returns(uint256)
func (_Compound *CompoundCallerSession) CompBorrowerIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Compound.Contract.CompBorrowerIndex(&_Compound.CallOpts, arg0, arg1)
}

// CompClaimThreshold is a free data retrieval call binding the contract method 0x747026c9.
//
// Solidity: function compClaimThreshold() view returns(uint256)
func (_Compound *CompoundCaller) CompClaimThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "compClaimThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompClaimThreshold is a free data retrieval call binding the contract method 0x747026c9.
//
// Solidity: function compClaimThreshold() view returns(uint256)
func (_Compound *CompoundSession) CompClaimThreshold() (*big.Int, error) {
	return _Compound.Contract.CompClaimThreshold(&_Compound.CallOpts)
}

// CompClaimThreshold is a free data retrieval call binding the contract method 0x747026c9.
//
// Solidity: function compClaimThreshold() view returns(uint256)
func (_Compound *CompoundCallerSession) CompClaimThreshold() (*big.Int, error) {
	return _Compound.Contract.CompClaimThreshold(&_Compound.CallOpts)
}

// CompInitialIndex is a free data retrieval call binding the contract method 0xa7f0e231.
//
// Solidity: function compInitialIndex() view returns(uint224)
func (_Compound *CompoundCaller) CompInitialIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "compInitialIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompInitialIndex is a free data retrieval call binding the contract method 0xa7f0e231.
//
// Solidity: function compInitialIndex() view returns(uint224)
func (_Compound *CompoundSession) CompInitialIndex() (*big.Int, error) {
	return _Compound.Contract.CompInitialIndex(&_Compound.CallOpts)
}

// CompInitialIndex is a free data retrieval call binding the contract method 0xa7f0e231.
//
// Solidity: function compInitialIndex() view returns(uint224)
func (_Compound *CompoundCallerSession) CompInitialIndex() (*big.Int, error) {
	return _Compound.Contract.CompInitialIndex(&_Compound.CallOpts)
}

// CompRate is a free data retrieval call binding the contract method 0xaa900754.
//
// Solidity: function compRate() view returns(uint256)
func (_Compound *CompoundCaller) CompRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "compRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompRate is a free data retrieval call binding the contract method 0xaa900754.
//
// Solidity: function compRate() view returns(uint256)
func (_Compound *CompoundSession) CompRate() (*big.Int, error) {
	return _Compound.Contract.CompRate(&_Compound.CallOpts)
}

// CompRate is a free data retrieval call binding the contract method 0xaa900754.
//
// Solidity: function compRate() view returns(uint256)
func (_Compound *CompoundCallerSession) CompRate() (*big.Int, error) {
	return _Compound.Contract.CompRate(&_Compound.CallOpts)
}

// CompSpeeds is a free data retrieval call binding the contract method 0x1d7b33d7.
//
// Solidity: function compSpeeds(address ) view returns(uint256)
func (_Compound *CompoundCaller) CompSpeeds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "compSpeeds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompSpeeds is a free data retrieval call binding the contract method 0x1d7b33d7.
//
// Solidity: function compSpeeds(address ) view returns(uint256)
func (_Compound *CompoundSession) CompSpeeds(arg0 common.Address) (*big.Int, error) {
	return _Compound.Contract.CompSpeeds(&_Compound.CallOpts, arg0)
}

// CompSpeeds is a free data retrieval call binding the contract method 0x1d7b33d7.
//
// Solidity: function compSpeeds(address ) view returns(uint256)
func (_Compound *CompoundCallerSession) CompSpeeds(arg0 common.Address) (*big.Int, error) {
	return _Compound.Contract.CompSpeeds(&_Compound.CallOpts, arg0)
}

// CompSupplierIndex is a free data retrieval call binding the contract method 0xb21be7fd.
//
// Solidity: function compSupplierIndex(address , address ) view returns(uint256)
func (_Compound *CompoundCaller) CompSupplierIndex(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "compSupplierIndex", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompSupplierIndex is a free data retrieval call binding the contract method 0xb21be7fd.
//
// Solidity: function compSupplierIndex(address , address ) view returns(uint256)
func (_Compound *CompoundSession) CompSupplierIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Compound.Contract.CompSupplierIndex(&_Compound.CallOpts, arg0, arg1)
}

// CompSupplierIndex is a free data retrieval call binding the contract method 0xb21be7fd.
//
// Solidity: function compSupplierIndex(address , address ) view returns(uint256)
func (_Compound *CompoundCallerSession) CompSupplierIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Compound.Contract.CompSupplierIndex(&_Compound.CallOpts, arg0, arg1)
}

// CompSupplyState is a free data retrieval call binding the contract method 0x6b79c38d.
//
// Solidity: function compSupplyState(address ) view returns(uint224 index, uint32 block)
func (_Compound *CompoundCaller) CompSupplyState(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "compSupplyState", arg0)

	outstruct := new(struct {
		Index *big.Int
		Block uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Block = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// CompSupplyState is a free data retrieval call binding the contract method 0x6b79c38d.
//
// Solidity: function compSupplyState(address ) view returns(uint224 index, uint32 block)
func (_Compound *CompoundSession) CompSupplyState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Compound.Contract.CompSupplyState(&_Compound.CallOpts, arg0)
}

// CompSupplyState is a free data retrieval call binding the contract method 0x6b79c38d.
//
// Solidity: function compSupplyState(address ) view returns(uint224 index, uint32 block)
func (_Compound *CompoundCallerSession) CompSupplyState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Compound.Contract.CompSupplyState(&_Compound.CallOpts, arg0)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Compound *CompoundCaller) ComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "comptrollerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Compound *CompoundSession) ComptrollerImplementation() (common.Address, error) {
	return _Compound.Contract.ComptrollerImplementation(&_Compound.CallOpts)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Compound *CompoundCallerSession) ComptrollerImplementation() (common.Address, error) {
	return _Compound.Contract.ComptrollerImplementation(&_Compound.CallOpts)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_Compound *CompoundCaller) GetAccountLiquidity(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "getAccountLiquidity", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_Compound *CompoundSession) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Compound.Contract.GetAccountLiquidity(&_Compound.CallOpts, account)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_Compound *CompoundCallerSession) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Compound.Contract.GetAccountLiquidity(&_Compound.CallOpts, account)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_Compound *CompoundCaller) GetAllMarkets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "getAllMarkets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_Compound *CompoundSession) GetAllMarkets() ([]common.Address, error) {
	return _Compound.Contract.GetAllMarkets(&_Compound.CallOpts)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_Compound *CompoundCallerSession) GetAllMarkets() ([]common.Address, error) {
	return _Compound.Contract.GetAllMarkets(&_Compound.CallOpts)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_Compound *CompoundCaller) GetAssetsIn(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "getAssetsIn", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_Compound *CompoundSession) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _Compound.Contract.GetAssetsIn(&_Compound.CallOpts, account)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_Compound *CompoundCallerSession) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _Compound.Contract.GetAssetsIn(&_Compound.CallOpts, account)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_Compound *CompoundCaller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_Compound *CompoundSession) GetBlockNumber() (*big.Int, error) {
	return _Compound.Contract.GetBlockNumber(&_Compound.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_Compound *CompoundCallerSession) GetBlockNumber() (*big.Int, error) {
	return _Compound.Contract.GetBlockNumber(&_Compound.CallOpts)
}

// GetCompAddress is a free data retrieval call binding the contract method 0x9d1b5a0a.
//
// Solidity: function getCompAddress() view returns(address)
func (_Compound *CompoundCaller) GetCompAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "getCompAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCompAddress is a free data retrieval call binding the contract method 0x9d1b5a0a.
//
// Solidity: function getCompAddress() view returns(address)
func (_Compound *CompoundSession) GetCompAddress() (common.Address, error) {
	return _Compound.Contract.GetCompAddress(&_Compound.CallOpts)
}

// GetCompAddress is a free data retrieval call binding the contract method 0x9d1b5a0a.
//
// Solidity: function getCompAddress() view returns(address)
func (_Compound *CompoundCallerSession) GetCompAddress() (common.Address, error) {
	return _Compound.Contract.GetCompAddress(&_Compound.CallOpts)
}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address cTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_Compound *CompoundCaller) GetHypotheticalAccountLiquidity(opts *bind.CallOpts, account common.Address, cTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "getHypotheticalAccountLiquidity", account, cTokenModify, redeemTokens, borrowAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address cTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_Compound *CompoundSession) GetHypotheticalAccountLiquidity(account common.Address, cTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Compound.Contract.GetHypotheticalAccountLiquidity(&_Compound.CallOpts, account, cTokenModify, redeemTokens, borrowAmount)
}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address cTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_Compound *CompoundCallerSession) GetHypotheticalAccountLiquidity(account common.Address, cTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Compound.Contract.GetHypotheticalAccountLiquidity(&_Compound.CallOpts, account, cTokenModify, redeemTokens, borrowAmount)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_Compound *CompoundCaller) IsComptroller(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "isComptroller")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_Compound *CompoundSession) IsComptroller() (bool, error) {
	return _Compound.Contract.IsComptroller(&_Compound.CallOpts)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_Compound *CompoundCallerSession) IsComptroller() (bool, error) {
	return _Compound.Contract.IsComptroller(&_Compound.CallOpts)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address cTokenBorrowed, address cTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Compound *CompoundCaller) LiquidateCalculateSeizeTokens(opts *bind.CallOpts, cTokenBorrowed common.Address, cTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "liquidateCalculateSeizeTokens", cTokenBorrowed, cTokenCollateral, actualRepayAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address cTokenBorrowed, address cTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Compound *CompoundSession) LiquidateCalculateSeizeTokens(cTokenBorrowed common.Address, cTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Compound.Contract.LiquidateCalculateSeizeTokens(&_Compound.CallOpts, cTokenBorrowed, cTokenCollateral, actualRepayAmount)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address cTokenBorrowed, address cTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Compound *CompoundCallerSession) LiquidateCalculateSeizeTokens(cTokenBorrowed common.Address, cTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Compound.Contract.LiquidateCalculateSeizeTokens(&_Compound.CallOpts, cTokenBorrowed, cTokenCollateral, actualRepayAmount)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_Compound *CompoundCaller) LiquidationIncentiveMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "liquidationIncentiveMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_Compound *CompoundSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _Compound.Contract.LiquidationIncentiveMantissa(&_Compound.CallOpts)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_Compound *CompoundCallerSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _Compound.Contract.LiquidationIncentiveMantissa(&_Compound.CallOpts)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isComped)
func (_Compound *CompoundCaller) Markets(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsComped                 bool
}, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "markets", arg0)

	outstruct := new(struct {
		IsListed                 bool
		CollateralFactorMantissa *big.Int
		IsComped                 bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsListed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.CollateralFactorMantissa = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsComped = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isComped)
func (_Compound *CompoundSession) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsComped                 bool
}, error) {
	return _Compound.Contract.Markets(&_Compound.CallOpts, arg0)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isComped)
func (_Compound *CompoundCallerSession) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsComped                 bool
}, error) {
	return _Compound.Contract.Markets(&_Compound.CallOpts, arg0)
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_Compound *CompoundCaller) MaxAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "maxAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_Compound *CompoundSession) MaxAssets() (*big.Int, error) {
	return _Compound.Contract.MaxAssets(&_Compound.CallOpts)
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_Compound *CompoundCallerSession) MaxAssets() (*big.Int, error) {
	return _Compound.Contract.MaxAssets(&_Compound.CallOpts)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_Compound *CompoundCaller) MintGuardianPaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "mintGuardianPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_Compound *CompoundSession) MintGuardianPaused(arg0 common.Address) (bool, error) {
	return _Compound.Contract.MintGuardianPaused(&_Compound.CallOpts, arg0)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_Compound *CompoundCallerSession) MintGuardianPaused(arg0 common.Address) (bool, error) {
	return _Compound.Contract.MintGuardianPaused(&_Compound.CallOpts, arg0)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Compound *CompoundCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Compound *CompoundSession) Oracle() (common.Address, error) {
	return _Compound.Contract.Oracle(&_Compound.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Compound *CompoundCallerSession) Oracle() (common.Address, error) {
	return _Compound.Contract.Oracle(&_Compound.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Compound *CompoundCaller) PauseGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "pauseGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Compound *CompoundSession) PauseGuardian() (common.Address, error) {
	return _Compound.Contract.PauseGuardian(&_Compound.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Compound *CompoundCallerSession) PauseGuardian() (common.Address, error) {
	return _Compound.Contract.PauseGuardian(&_Compound.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Compound *CompoundCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Compound *CompoundSession) PendingAdmin() (common.Address, error) {
	return _Compound.Contract.PendingAdmin(&_Compound.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Compound *CompoundCallerSession) PendingAdmin() (common.Address, error) {
	return _Compound.Contract.PendingAdmin(&_Compound.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Compound *CompoundCaller) PendingComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "pendingComptrollerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Compound *CompoundSession) PendingComptrollerImplementation() (common.Address, error) {
	return _Compound.Contract.PendingComptrollerImplementation(&_Compound.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Compound *CompoundCallerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _Compound.Contract.PendingComptrollerImplementation(&_Compound.CallOpts)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_Compound *CompoundCaller) SeizeGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "seizeGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_Compound *CompoundSession) SeizeGuardianPaused() (bool, error) {
	return _Compound.Contract.SeizeGuardianPaused(&_Compound.CallOpts)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_Compound *CompoundCallerSession) SeizeGuardianPaused() (bool, error) {
	return _Compound.Contract.SeizeGuardianPaused(&_Compound.CallOpts)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_Compound *CompoundCaller) TransferGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Compound.contract.Call(opts, &out, "transferGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_Compound *CompoundSession) TransferGuardianPaused() (bool, error) {
	return _Compound.Contract.TransferGuardianPaused(&_Compound.CallOpts)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_Compound *CompoundCallerSession) TransferGuardianPaused() (bool, error) {
	return _Compound.Contract.TransferGuardianPaused(&_Compound.CallOpts)
}

// AddCompMarkets is a paid mutator transaction binding the contract method 0xce485c5e.
//
// Solidity: function _addCompMarkets(address[] cTokens) returns()
func (_Compound *CompoundTransactor) AddCompMarkets(opts *bind.TransactOpts, cTokens []common.Address) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "_addCompMarkets", cTokens)
}

// AddCompMarkets is a paid mutator transaction binding the contract method 0xce485c5e.
//
// Solidity: function _addCompMarkets(address[] cTokens) returns()
func (_Compound *CompoundSession) AddCompMarkets(cTokens []common.Address) (*types.Transaction, error) {
	return _Compound.Contract.AddCompMarkets(&_Compound.TransactOpts, cTokens)
}

// AddCompMarkets is a paid mutator transaction binding the contract method 0xce485c5e.
//
// Solidity: function _addCompMarkets(address[] cTokens) returns()
func (_Compound *CompoundTransactorSession) AddCompMarkets(cTokens []common.Address) (*types.Transaction, error) {
	return _Compound.Contract.AddCompMarkets(&_Compound.TransactOpts, cTokens)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_Compound *CompoundTransactor) Become(opts *bind.TransactOpts, unitroller common.Address) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "_become", unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_Compound *CompoundSession) Become(unitroller common.Address) (*types.Transaction, error) {
	return _Compound.Contract.Become(&_Compound.TransactOpts, unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_Compound *CompoundTransactorSession) Become(unitroller common.Address) (*types.Transaction, error) {
	return _Compound.Contract.Become(&_Compound.TransactOpts, unitroller)
}

// DropCompMarket is a paid mutator transaction binding the contract method 0x3aa729b4.
//
// Solidity: function _dropCompMarket(address cToken) returns()
func (_Compound *CompoundTransactor) DropCompMarket(opts *bind.TransactOpts, cToken common.Address) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "_dropCompMarket", cToken)
}

// DropCompMarket is a paid mutator transaction binding the contract method 0x3aa729b4.
//
// Solidity: function _dropCompMarket(address cToken) returns()
func (_Compound *CompoundSession) DropCompMarket(cToken common.Address) (*types.Transaction, error) {
	return _Compound.Contract.DropCompMarket(&_Compound.TransactOpts, cToken)
}

// DropCompMarket is a paid mutator transaction binding the contract method 0x3aa729b4.
//
// Solidity: function _dropCompMarket(address cToken) returns()
func (_Compound *CompoundTransactorSession) DropCompMarket(cToken common.Address) (*types.Transaction, error) {
	return _Compound.Contract.DropCompMarket(&_Compound.TransactOpts, cToken)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_Compound *CompoundTransactor) SetBorrowCapGuardian(opts *bind.TransactOpts, newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "_setBorrowCapGuardian", newBorrowCapGuardian)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_Compound *CompoundSession) SetBorrowCapGuardian(newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _Compound.Contract.SetBorrowCapGuardian(&_Compound.TransactOpts, newBorrowCapGuardian)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_Compound *CompoundTransactorSession) SetBorrowCapGuardian(newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _Compound.Contract.SetBorrowCapGuardian(&_Compound.TransactOpts, newBorrowCapGuardian)
}

// SetBorrowPaused is a paid mutator transaction binding the contract method 0x18c882a5.
//
// Solidity: function _setBorrowPaused(address cToken, bool state) returns(bool)
func (_Compound *CompoundTransactor) SetBorrowPaused(opts *bind.TransactOpts, cToken common.Address, state bool) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "_setBorrowPaused", cToken, state)
}

// SetBorrowPaused is a paid mutator transaction binding the contract method 0x18c882a5.
//
// Solidity: function _setBorrowPaused(address cToken, bool state) returns(bool)
func (_Compound *CompoundSession) SetBorrowPaused(cToken common.Address, state bool) (*types.Transaction, error) {
	return _Compound.Contract.SetBorrowPaused(&_Compound.TransactOpts, cToken, state)
}

// SetBorrowPaused is a paid mutator transaction binding the contract method 0x18c882a5.
//
// Solidity: function _setBorrowPaused(address cToken, bool state) returns(bool)
func (_Compound *CompoundTransactorSession) SetBorrowPaused(cToken common.Address, state bool) (*types.Transaction, error) {
	return _Compound.Contract.SetBorrowPaused(&_Compound.TransactOpts, cToken, state)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_Compound *CompoundTransactor) SetCloseFactor(opts *bind.TransactOpts, newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "_setCloseFactor", newCloseFactorMantissa)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_Compound *CompoundSession) SetCloseFactor(newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.SetCloseFactor(&_Compound.TransactOpts, newCloseFactorMantissa)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_Compound *CompoundTransactorSession) SetCloseFactor(newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.SetCloseFactor(&_Compound.TransactOpts, newCloseFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address cToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_Compound *CompoundTransactor) SetCollateralFactor(opts *bind.TransactOpts, cToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "_setCollateralFactor", cToken, newCollateralFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address cToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_Compound *CompoundSession) SetCollateralFactor(cToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.SetCollateralFactor(&_Compound.TransactOpts, cToken, newCollateralFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address cToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_Compound *CompoundTransactorSession) SetCollateralFactor(cToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.SetCollateralFactor(&_Compound.TransactOpts, cToken, newCollateralFactorMantissa)
}

// SetCompRate is a paid mutator transaction binding the contract method 0x6a491112.
//
// Solidity: function _setCompRate(uint256 compRate_) returns()
func (_Compound *CompoundTransactor) SetCompRate(opts *bind.TransactOpts, compRate_ *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "_setCompRate", compRate_)
}

// SetCompRate is a paid mutator transaction binding the contract method 0x6a491112.
//
// Solidity: function _setCompRate(uint256 compRate_) returns()
func (_Compound *CompoundSession) SetCompRate(compRate_ *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.SetCompRate(&_Compound.TransactOpts, compRate_)
}

// SetCompRate is a paid mutator transaction binding the contract method 0x6a491112.
//
// Solidity: function _setCompRate(uint256 compRate_) returns()
func (_Compound *CompoundTransactorSession) SetCompRate(compRate_ *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.SetCompRate(&_Compound.TransactOpts, compRate_)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_Compound *CompoundTransactor) SetLiquidationIncentive(opts *bind.TransactOpts, newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "_setLiquidationIncentive", newLiquidationIncentiveMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_Compound *CompoundSession) SetLiquidationIncentive(newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.SetLiquidationIncentive(&_Compound.TransactOpts, newLiquidationIncentiveMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_Compound *CompoundTransactorSession) SetLiquidationIncentive(newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.SetLiquidationIncentive(&_Compound.TransactOpts, newLiquidationIncentiveMantissa)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] cTokens, uint256[] newBorrowCaps) returns()
func (_Compound *CompoundTransactor) SetMarketBorrowCaps(opts *bind.TransactOpts, cTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "_setMarketBorrowCaps", cTokens, newBorrowCaps)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] cTokens, uint256[] newBorrowCaps) returns()
func (_Compound *CompoundSession) SetMarketBorrowCaps(cTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _Compound.Contract.SetMarketBorrowCaps(&_Compound.TransactOpts, cTokens, newBorrowCaps)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] cTokens, uint256[] newBorrowCaps) returns()
func (_Compound *CompoundTransactorSession) SetMarketBorrowCaps(cTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _Compound.Contract.SetMarketBorrowCaps(&_Compound.TransactOpts, cTokens, newBorrowCaps)
}

// SetMaxAssets is a paid mutator transaction binding the contract method 0xd9226ced.
//
// Solidity: function _setMaxAssets(uint256 newMaxAssets) returns(uint256)
func (_Compound *CompoundTransactor) SetMaxAssets(opts *bind.TransactOpts, newMaxAssets *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "_setMaxAssets", newMaxAssets)
}

// SetMaxAssets is a paid mutator transaction binding the contract method 0xd9226ced.
//
// Solidity: function _setMaxAssets(uint256 newMaxAssets) returns(uint256)
func (_Compound *CompoundSession) SetMaxAssets(newMaxAssets *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.SetMaxAssets(&_Compound.TransactOpts, newMaxAssets)
}

// SetMaxAssets is a paid mutator transaction binding the contract method 0xd9226ced.
//
// Solidity: function _setMaxAssets(uint256 newMaxAssets) returns(uint256)
func (_Compound *CompoundTransactorSession) SetMaxAssets(newMaxAssets *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.SetMaxAssets(&_Compound.TransactOpts, newMaxAssets)
}

// SetMintPaused is a paid mutator transaction binding the contract method 0x3bcf7ec1.
//
// Solidity: function _setMintPaused(address cToken, bool state) returns(bool)
func (_Compound *CompoundTransactor) SetMintPaused(opts *bind.TransactOpts, cToken common.Address, state bool) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "_setMintPaused", cToken, state)
}

// SetMintPaused is a paid mutator transaction binding the contract method 0x3bcf7ec1.
//
// Solidity: function _setMintPaused(address cToken, bool state) returns(bool)
func (_Compound *CompoundSession) SetMintPaused(cToken common.Address, state bool) (*types.Transaction, error) {
	return _Compound.Contract.SetMintPaused(&_Compound.TransactOpts, cToken, state)
}

// SetMintPaused is a paid mutator transaction binding the contract method 0x3bcf7ec1.
//
// Solidity: function _setMintPaused(address cToken, bool state) returns(bool)
func (_Compound *CompoundTransactorSession) SetMintPaused(cToken common.Address, state bool) (*types.Transaction, error) {
	return _Compound.Contract.SetMintPaused(&_Compound.TransactOpts, cToken, state)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_Compound *CompoundTransactor) SetPauseGuardian(opts *bind.TransactOpts, newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "_setPauseGuardian", newPauseGuardian)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_Compound *CompoundSession) SetPauseGuardian(newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Compound.Contract.SetPauseGuardian(&_Compound.TransactOpts, newPauseGuardian)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_Compound *CompoundTransactorSession) SetPauseGuardian(newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Compound.Contract.SetPauseGuardian(&_Compound.TransactOpts, newPauseGuardian)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_Compound *CompoundTransactor) SetPriceOracle(opts *bind.TransactOpts, newOracle common.Address) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "_setPriceOracle", newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_Compound *CompoundSession) SetPriceOracle(newOracle common.Address) (*types.Transaction, error) {
	return _Compound.Contract.SetPriceOracle(&_Compound.TransactOpts, newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_Compound *CompoundTransactorSession) SetPriceOracle(newOracle common.Address) (*types.Transaction, error) {
	return _Compound.Contract.SetPriceOracle(&_Compound.TransactOpts, newOracle)
}

// SetSeizePaused is a paid mutator transaction binding the contract method 0x2d70db78.
//
// Solidity: function _setSeizePaused(bool state) returns(bool)
func (_Compound *CompoundTransactor) SetSeizePaused(opts *bind.TransactOpts, state bool) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "_setSeizePaused", state)
}

// SetSeizePaused is a paid mutator transaction binding the contract method 0x2d70db78.
//
// Solidity: function _setSeizePaused(bool state) returns(bool)
func (_Compound *CompoundSession) SetSeizePaused(state bool) (*types.Transaction, error) {
	return _Compound.Contract.SetSeizePaused(&_Compound.TransactOpts, state)
}

// SetSeizePaused is a paid mutator transaction binding the contract method 0x2d70db78.
//
// Solidity: function _setSeizePaused(bool state) returns(bool)
func (_Compound *CompoundTransactorSession) SetSeizePaused(state bool) (*types.Transaction, error) {
	return _Compound.Contract.SetSeizePaused(&_Compound.TransactOpts, state)
}

// SetTransferPaused is a paid mutator transaction binding the contract method 0x8ebf6364.
//
// Solidity: function _setTransferPaused(bool state) returns(bool)
func (_Compound *CompoundTransactor) SetTransferPaused(opts *bind.TransactOpts, state bool) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "_setTransferPaused", state)
}

// SetTransferPaused is a paid mutator transaction binding the contract method 0x8ebf6364.
//
// Solidity: function _setTransferPaused(bool state) returns(bool)
func (_Compound *CompoundSession) SetTransferPaused(state bool) (*types.Transaction, error) {
	return _Compound.Contract.SetTransferPaused(&_Compound.TransactOpts, state)
}

// SetTransferPaused is a paid mutator transaction binding the contract method 0x8ebf6364.
//
// Solidity: function _setTransferPaused(bool state) returns(bool)
func (_Compound *CompoundTransactorSession) SetTransferPaused(state bool) (*types.Transaction, error) {
	return _Compound.Contract.SetTransferPaused(&_Compound.TransactOpts, state)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address cToken) returns(uint256)
func (_Compound *CompoundTransactor) SupportMarket(opts *bind.TransactOpts, cToken common.Address) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "_supportMarket", cToken)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address cToken) returns(uint256)
func (_Compound *CompoundSession) SupportMarket(cToken common.Address) (*types.Transaction, error) {
	return _Compound.Contract.SupportMarket(&_Compound.TransactOpts, cToken)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address cToken) returns(uint256)
func (_Compound *CompoundTransactorSession) SupportMarket(cToken common.Address) (*types.Transaction, error) {
	return _Compound.Contract.SupportMarket(&_Compound.TransactOpts, cToken)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address cToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_Compound *CompoundTransactor) BorrowAllowed(opts *bind.TransactOpts, cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "borrowAllowed", cToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address cToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_Compound *CompoundSession) BorrowAllowed(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.BorrowAllowed(&_Compound.TransactOpts, cToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address cToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_Compound *CompoundTransactorSession) BorrowAllowed(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.BorrowAllowed(&_Compound.TransactOpts, cToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address cToken, address borrower, uint256 borrowAmount) returns()
func (_Compound *CompoundTransactor) BorrowVerify(opts *bind.TransactOpts, cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "borrowVerify", cToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address cToken, address borrower, uint256 borrowAmount) returns()
func (_Compound *CompoundSession) BorrowVerify(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.BorrowVerify(&_Compound.TransactOpts, cToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address cToken, address borrower, uint256 borrowAmount) returns()
func (_Compound *CompoundTransactorSession) BorrowVerify(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.BorrowVerify(&_Compound.TransactOpts, cToken, borrower, borrowAmount)
}

// ClaimComp is a paid mutator transaction binding the contract method 0xe9af0292.
//
// Solidity: function claimComp(address holder) returns()
func (_Compound *CompoundTransactor) ClaimComp(opts *bind.TransactOpts, holder common.Address) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "claimComp", holder)
}

// ClaimComp is a paid mutator transaction binding the contract method 0xe9af0292.
//
// Solidity: function claimComp(address holder) returns()
func (_Compound *CompoundSession) ClaimComp(holder common.Address) (*types.Transaction, error) {
	return _Compound.Contract.ClaimComp(&_Compound.TransactOpts, holder)
}

// ClaimComp is a paid mutator transaction binding the contract method 0xe9af0292.
//
// Solidity: function claimComp(address holder) returns()
func (_Compound *CompoundTransactorSession) ClaimComp(holder common.Address) (*types.Transaction, error) {
	return _Compound.Contract.ClaimComp(&_Compound.TransactOpts, holder)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] cTokens) returns(uint256[])
func (_Compound *CompoundTransactor) EnterMarkets(opts *bind.TransactOpts, cTokens []common.Address) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "enterMarkets", cTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] cTokens) returns(uint256[])
func (_Compound *CompoundSession) EnterMarkets(cTokens []common.Address) (*types.Transaction, error) {
	return _Compound.Contract.EnterMarkets(&_Compound.TransactOpts, cTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] cTokens) returns(uint256[])
func (_Compound *CompoundTransactorSession) EnterMarkets(cTokens []common.Address) (*types.Transaction, error) {
	return _Compound.Contract.EnterMarkets(&_Compound.TransactOpts, cTokens)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address cTokenAddress) returns(uint256)
func (_Compound *CompoundTransactor) ExitMarket(opts *bind.TransactOpts, cTokenAddress common.Address) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "exitMarket", cTokenAddress)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address cTokenAddress) returns(uint256)
func (_Compound *CompoundSession) ExitMarket(cTokenAddress common.Address) (*types.Transaction, error) {
	return _Compound.Contract.ExitMarket(&_Compound.TransactOpts, cTokenAddress)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address cTokenAddress) returns(uint256)
func (_Compound *CompoundTransactorSession) ExitMarket(cTokenAddress common.Address) (*types.Transaction, error) {
	return _Compound.Contract.ExitMarket(&_Compound.TransactOpts, cTokenAddress)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_Compound *CompoundTransactor) LiquidateBorrowAllowed(opts *bind.TransactOpts, cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "liquidateBorrowAllowed", cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_Compound *CompoundSession) LiquidateBorrowAllowed(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.LiquidateBorrowAllowed(&_Compound.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_Compound *CompoundTransactorSession) LiquidateBorrowAllowed(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.LiquidateBorrowAllowed(&_Compound.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_Compound *CompoundTransactor) LiquidateBorrowVerify(opts *bind.TransactOpts, cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "liquidateBorrowVerify", cTokenBorrowed, cTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_Compound *CompoundSession) LiquidateBorrowVerify(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.LiquidateBorrowVerify(&_Compound.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_Compound *CompoundTransactorSession) LiquidateBorrowVerify(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.LiquidateBorrowVerify(&_Compound.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address cToken, address minter, uint256 mintAmount) returns(uint256)
func (_Compound *CompoundTransactor) MintAllowed(opts *bind.TransactOpts, cToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "mintAllowed", cToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address cToken, address minter, uint256 mintAmount) returns(uint256)
func (_Compound *CompoundSession) MintAllowed(cToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.MintAllowed(&_Compound.TransactOpts, cToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address cToken, address minter, uint256 mintAmount) returns(uint256)
func (_Compound *CompoundTransactorSession) MintAllowed(cToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.MintAllowed(&_Compound.TransactOpts, cToken, minter, mintAmount)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address cToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_Compound *CompoundTransactor) MintVerify(opts *bind.TransactOpts, cToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "mintVerify", cToken, minter, actualMintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address cToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_Compound *CompoundSession) MintVerify(cToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.MintVerify(&_Compound.TransactOpts, cToken, minter, actualMintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address cToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_Compound *CompoundTransactorSession) MintVerify(cToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.MintVerify(&_Compound.TransactOpts, cToken, minter, actualMintAmount, mintTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address cToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_Compound *CompoundTransactor) RedeemAllowed(opts *bind.TransactOpts, cToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "redeemAllowed", cToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address cToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_Compound *CompoundSession) RedeemAllowed(cToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.RedeemAllowed(&_Compound.TransactOpts, cToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address cToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_Compound *CompoundTransactorSession) RedeemAllowed(cToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.RedeemAllowed(&_Compound.TransactOpts, cToken, redeemer, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address cToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_Compound *CompoundTransactor) RedeemVerify(opts *bind.TransactOpts, cToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "redeemVerify", cToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address cToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_Compound *CompoundSession) RedeemVerify(cToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.RedeemVerify(&_Compound.TransactOpts, cToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address cToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_Compound *CompoundTransactorSession) RedeemVerify(cToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.RedeemVerify(&_Compound.TransactOpts, cToken, redeemer, redeemAmount, redeemTokens)
}

// RefreshCompSpeeds is a paid mutator transaction binding the contract method 0x4d8e5037.
//
// Solidity: function refreshCompSpeeds() returns()
func (_Compound *CompoundTransactor) RefreshCompSpeeds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "refreshCompSpeeds")
}

// RefreshCompSpeeds is a paid mutator transaction binding the contract method 0x4d8e5037.
//
// Solidity: function refreshCompSpeeds() returns()
func (_Compound *CompoundSession) RefreshCompSpeeds() (*types.Transaction, error) {
	return _Compound.Contract.RefreshCompSpeeds(&_Compound.TransactOpts)
}

// RefreshCompSpeeds is a paid mutator transaction binding the contract method 0x4d8e5037.
//
// Solidity: function refreshCompSpeeds() returns()
func (_Compound *CompoundTransactorSession) RefreshCompSpeeds() (*types.Transaction, error) {
	return _Compound.Contract.RefreshCompSpeeds(&_Compound.TransactOpts)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address cToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_Compound *CompoundTransactor) RepayBorrowAllowed(opts *bind.TransactOpts, cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "repayBorrowAllowed", cToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address cToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_Compound *CompoundSession) RepayBorrowAllowed(cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.RepayBorrowAllowed(&_Compound.TransactOpts, cToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address cToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_Compound *CompoundTransactorSession) RepayBorrowAllowed(cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.RepayBorrowAllowed(&_Compound.TransactOpts, cToken, payer, borrower, repayAmount)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address cToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_Compound *CompoundTransactor) RepayBorrowVerify(opts *bind.TransactOpts, cToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "repayBorrowVerify", cToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address cToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_Compound *CompoundSession) RepayBorrowVerify(cToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.RepayBorrowVerify(&_Compound.TransactOpts, cToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address cToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_Compound *CompoundTransactorSession) RepayBorrowVerify(cToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.RepayBorrowVerify(&_Compound.TransactOpts, cToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Compound *CompoundTransactor) SeizeAllowed(opts *bind.TransactOpts, cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "seizeAllowed", cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Compound *CompoundSession) SeizeAllowed(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.SeizeAllowed(&_Compound.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Compound *CompoundTransactorSession) SeizeAllowed(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.SeizeAllowed(&_Compound.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Compound *CompoundTransactor) SeizeVerify(opts *bind.TransactOpts, cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "seizeVerify", cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Compound *CompoundSession) SeizeVerify(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.SeizeVerify(&_Compound.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Compound *CompoundTransactorSession) SeizeVerify(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.SeizeVerify(&_Compound.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address cToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_Compound *CompoundTransactor) TransferAllowed(opts *bind.TransactOpts, cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "transferAllowed", cToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address cToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_Compound *CompoundSession) TransferAllowed(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.TransferAllowed(&_Compound.TransactOpts, cToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address cToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_Compound *CompoundTransactorSession) TransferAllowed(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.TransferAllowed(&_Compound.TransactOpts, cToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address cToken, address src, address dst, uint256 transferTokens) returns()
func (_Compound *CompoundTransactor) TransferVerify(opts *bind.TransactOpts, cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Compound.contract.Transact(opts, "transferVerify", cToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address cToken, address src, address dst, uint256 transferTokens) returns()
func (_Compound *CompoundSession) TransferVerify(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.TransferVerify(&_Compound.TransactOpts, cToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address cToken, address src, address dst, uint256 transferTokens) returns()
func (_Compound *CompoundTransactorSession) TransferVerify(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Compound.Contract.TransferVerify(&_Compound.TransactOpts, cToken, src, dst, transferTokens)
}

// CompoundActionPausedIterator is returned from FilterActionPaused and is used to iterate over the raw logs and unpacked data for ActionPaused events raised by the Compound contract.
type CompoundActionPausedIterator struct {
	Event *CompoundActionPaused // Event containing the contract specifics and raw log

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
func (it *CompoundActionPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompoundActionPaused)
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
		it.Event = new(CompoundActionPaused)
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
func (it *CompoundActionPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompoundActionPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompoundActionPaused represents a ActionPaused event raised by the Compound contract.
type CompoundActionPaused struct {
	Action     string
	PauseState bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterActionPaused is a free log retrieval operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_Compound *CompoundFilterer) FilterActionPaused(opts *bind.FilterOpts) (*CompoundActionPausedIterator, error) {

	logs, sub, err := _Compound.contract.FilterLogs(opts, "ActionPaused")
	if err != nil {
		return nil, err
	}
	return &CompoundActionPausedIterator{contract: _Compound.contract, event: "ActionPaused", logs: logs, sub: sub}, nil
}

// WatchActionPaused is a free log subscription operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_Compound *CompoundFilterer) WatchActionPaused(opts *bind.WatchOpts, sink chan<- *CompoundActionPaused) (event.Subscription, error) {

	logs, sub, err := _Compound.contract.WatchLogs(opts, "ActionPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompoundActionPaused)
				if err := _Compound.contract.UnpackLog(event, "ActionPaused", log); err != nil {
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

// ParseActionPaused is a log parse operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_Compound *CompoundFilterer) ParseActionPaused(log types.Log) (*CompoundActionPaused, error) {
	event := new(CompoundActionPaused)
	if err := _Compound.contract.UnpackLog(event, "ActionPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompoundCompSpeedUpdatedIterator is returned from FilterCompSpeedUpdated and is used to iterate over the raw logs and unpacked data for CompSpeedUpdated events raised by the Compound contract.
type CompoundCompSpeedUpdatedIterator struct {
	Event *CompoundCompSpeedUpdated // Event containing the contract specifics and raw log

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
func (it *CompoundCompSpeedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompoundCompSpeedUpdated)
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
		it.Event = new(CompoundCompSpeedUpdated)
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
func (it *CompoundCompSpeedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompoundCompSpeedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompoundCompSpeedUpdated represents a CompSpeedUpdated event raised by the Compound contract.
type CompoundCompSpeedUpdated struct {
	CToken   common.Address
	NewSpeed *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCompSpeedUpdated is a free log retrieval operation binding the contract event 0x2ab93f65628379309f36cb125e90d7c902454a545c4f8b8cb0794af75c24b807.
//
// Solidity: event CompSpeedUpdated(address indexed cToken, uint256 newSpeed)
func (_Compound *CompoundFilterer) FilterCompSpeedUpdated(opts *bind.FilterOpts, cToken []common.Address) (*CompoundCompSpeedUpdatedIterator, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}

	logs, sub, err := _Compound.contract.FilterLogs(opts, "CompSpeedUpdated", cTokenRule)
	if err != nil {
		return nil, err
	}
	return &CompoundCompSpeedUpdatedIterator{contract: _Compound.contract, event: "CompSpeedUpdated", logs: logs, sub: sub}, nil
}

// WatchCompSpeedUpdated is a free log subscription operation binding the contract event 0x2ab93f65628379309f36cb125e90d7c902454a545c4f8b8cb0794af75c24b807.
//
// Solidity: event CompSpeedUpdated(address indexed cToken, uint256 newSpeed)
func (_Compound *CompoundFilterer) WatchCompSpeedUpdated(opts *bind.WatchOpts, sink chan<- *CompoundCompSpeedUpdated, cToken []common.Address) (event.Subscription, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}

	logs, sub, err := _Compound.contract.WatchLogs(opts, "CompSpeedUpdated", cTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompoundCompSpeedUpdated)
				if err := _Compound.contract.UnpackLog(event, "CompSpeedUpdated", log); err != nil {
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

// ParseCompSpeedUpdated is a log parse operation binding the contract event 0x2ab93f65628379309f36cb125e90d7c902454a545c4f8b8cb0794af75c24b807.
//
// Solidity: event CompSpeedUpdated(address indexed cToken, uint256 newSpeed)
func (_Compound *CompoundFilterer) ParseCompSpeedUpdated(log types.Log) (*CompoundCompSpeedUpdated, error) {
	event := new(CompoundCompSpeedUpdated)
	if err := _Compound.contract.UnpackLog(event, "CompSpeedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompoundDistributedBorrowerCompIterator is returned from FilterDistributedBorrowerComp and is used to iterate over the raw logs and unpacked data for DistributedBorrowerComp events raised by the Compound contract.
type CompoundDistributedBorrowerCompIterator struct {
	Event *CompoundDistributedBorrowerComp // Event containing the contract specifics and raw log

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
func (it *CompoundDistributedBorrowerCompIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompoundDistributedBorrowerComp)
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
		it.Event = new(CompoundDistributedBorrowerComp)
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
func (it *CompoundDistributedBorrowerCompIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompoundDistributedBorrowerCompIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompoundDistributedBorrowerComp represents a DistributedBorrowerComp event raised by the Compound contract.
type CompoundDistributedBorrowerComp struct {
	CToken          common.Address
	Borrower        common.Address
	CompDelta       *big.Int
	CompBorrowIndex *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDistributedBorrowerComp is a free log retrieval operation binding the contract event 0x1fc3ecc087d8d2d15e23d0032af5a47059c3892d003d8e139fdcb6bb327c99a6.
//
// Solidity: event DistributedBorrowerComp(address indexed cToken, address indexed borrower, uint256 compDelta, uint256 compBorrowIndex)
func (_Compound *CompoundFilterer) FilterDistributedBorrowerComp(opts *bind.FilterOpts, cToken []common.Address, borrower []common.Address) (*CompoundDistributedBorrowerCompIterator, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Compound.contract.FilterLogs(opts, "DistributedBorrowerComp", cTokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &CompoundDistributedBorrowerCompIterator{contract: _Compound.contract, event: "DistributedBorrowerComp", logs: logs, sub: sub}, nil
}

// WatchDistributedBorrowerComp is a free log subscription operation binding the contract event 0x1fc3ecc087d8d2d15e23d0032af5a47059c3892d003d8e139fdcb6bb327c99a6.
//
// Solidity: event DistributedBorrowerComp(address indexed cToken, address indexed borrower, uint256 compDelta, uint256 compBorrowIndex)
func (_Compound *CompoundFilterer) WatchDistributedBorrowerComp(opts *bind.WatchOpts, sink chan<- *CompoundDistributedBorrowerComp, cToken []common.Address, borrower []common.Address) (event.Subscription, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Compound.contract.WatchLogs(opts, "DistributedBorrowerComp", cTokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompoundDistributedBorrowerComp)
				if err := _Compound.contract.UnpackLog(event, "DistributedBorrowerComp", log); err != nil {
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

// ParseDistributedBorrowerComp is a log parse operation binding the contract event 0x1fc3ecc087d8d2d15e23d0032af5a47059c3892d003d8e139fdcb6bb327c99a6.
//
// Solidity: event DistributedBorrowerComp(address indexed cToken, address indexed borrower, uint256 compDelta, uint256 compBorrowIndex)
func (_Compound *CompoundFilterer) ParseDistributedBorrowerComp(log types.Log) (*CompoundDistributedBorrowerComp, error) {
	event := new(CompoundDistributedBorrowerComp)
	if err := _Compound.contract.UnpackLog(event, "DistributedBorrowerComp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompoundDistributedSupplierCompIterator is returned from FilterDistributedSupplierComp and is used to iterate over the raw logs and unpacked data for DistributedSupplierComp events raised by the Compound contract.
type CompoundDistributedSupplierCompIterator struct {
	Event *CompoundDistributedSupplierComp // Event containing the contract specifics and raw log

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
func (it *CompoundDistributedSupplierCompIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompoundDistributedSupplierComp)
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
		it.Event = new(CompoundDistributedSupplierComp)
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
func (it *CompoundDistributedSupplierCompIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompoundDistributedSupplierCompIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompoundDistributedSupplierComp represents a DistributedSupplierComp event raised by the Compound contract.
type CompoundDistributedSupplierComp struct {
	CToken          common.Address
	Supplier        common.Address
	CompDelta       *big.Int
	CompSupplyIndex *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDistributedSupplierComp is a free log retrieval operation binding the contract event 0x2caecd17d02f56fa897705dcc740da2d237c373f70686f4e0d9bd3bf0400ea7a.
//
// Solidity: event DistributedSupplierComp(address indexed cToken, address indexed supplier, uint256 compDelta, uint256 compSupplyIndex)
func (_Compound *CompoundFilterer) FilterDistributedSupplierComp(opts *bind.FilterOpts, cToken []common.Address, supplier []common.Address) (*CompoundDistributedSupplierCompIterator, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}
	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _Compound.contract.FilterLogs(opts, "DistributedSupplierComp", cTokenRule, supplierRule)
	if err != nil {
		return nil, err
	}
	return &CompoundDistributedSupplierCompIterator{contract: _Compound.contract, event: "DistributedSupplierComp", logs: logs, sub: sub}, nil
}

// WatchDistributedSupplierComp is a free log subscription operation binding the contract event 0x2caecd17d02f56fa897705dcc740da2d237c373f70686f4e0d9bd3bf0400ea7a.
//
// Solidity: event DistributedSupplierComp(address indexed cToken, address indexed supplier, uint256 compDelta, uint256 compSupplyIndex)
func (_Compound *CompoundFilterer) WatchDistributedSupplierComp(opts *bind.WatchOpts, sink chan<- *CompoundDistributedSupplierComp, cToken []common.Address, supplier []common.Address) (event.Subscription, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}
	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _Compound.contract.WatchLogs(opts, "DistributedSupplierComp", cTokenRule, supplierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompoundDistributedSupplierComp)
				if err := _Compound.contract.UnpackLog(event, "DistributedSupplierComp", log); err != nil {
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

// ParseDistributedSupplierComp is a log parse operation binding the contract event 0x2caecd17d02f56fa897705dcc740da2d237c373f70686f4e0d9bd3bf0400ea7a.
//
// Solidity: event DistributedSupplierComp(address indexed cToken, address indexed supplier, uint256 compDelta, uint256 compSupplyIndex)
func (_Compound *CompoundFilterer) ParseDistributedSupplierComp(log types.Log) (*CompoundDistributedSupplierComp, error) {
	event := new(CompoundDistributedSupplierComp)
	if err := _Compound.contract.UnpackLog(event, "DistributedSupplierComp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompoundFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the Compound contract.
type CompoundFailureIterator struct {
	Event *CompoundFailure // Event containing the contract specifics and raw log

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
func (it *CompoundFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompoundFailure)
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
		it.Event = new(CompoundFailure)
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
func (it *CompoundFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompoundFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompoundFailure represents a Failure event raised by the Compound contract.
type CompoundFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Compound *CompoundFilterer) FilterFailure(opts *bind.FilterOpts) (*CompoundFailureIterator, error) {

	logs, sub, err := _Compound.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &CompoundFailureIterator{contract: _Compound.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Compound *CompoundFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *CompoundFailure) (event.Subscription, error) {

	logs, sub, err := _Compound.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompoundFailure)
				if err := _Compound.contract.UnpackLog(event, "Failure", log); err != nil {
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

// ParseFailure is a log parse operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Compound *CompoundFilterer) ParseFailure(log types.Log) (*CompoundFailure, error) {
	event := new(CompoundFailure)
	if err := _Compound.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompoundMarketCompedIterator is returned from FilterMarketComped and is used to iterate over the raw logs and unpacked data for MarketComped events raised by the Compound contract.
type CompoundMarketCompedIterator struct {
	Event *CompoundMarketComped // Event containing the contract specifics and raw log

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
func (it *CompoundMarketCompedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompoundMarketComped)
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
		it.Event = new(CompoundMarketComped)
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
func (it *CompoundMarketCompedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompoundMarketCompedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompoundMarketComped represents a MarketComped event raised by the Compound contract.
type CompoundMarketComped struct {
	CToken   common.Address
	IsComped bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketComped is a free log retrieval operation binding the contract event 0x93c1f3e36ed71139f466a4ce8c9751790e2e33f5afb2df0dcfb3aeabe55d5aa2.
//
// Solidity: event MarketComped(address cToken, bool isComped)
func (_Compound *CompoundFilterer) FilterMarketComped(opts *bind.FilterOpts) (*CompoundMarketCompedIterator, error) {

	logs, sub, err := _Compound.contract.FilterLogs(opts, "MarketComped")
	if err != nil {
		return nil, err
	}
	return &CompoundMarketCompedIterator{contract: _Compound.contract, event: "MarketComped", logs: logs, sub: sub}, nil
}

// WatchMarketComped is a free log subscription operation binding the contract event 0x93c1f3e36ed71139f466a4ce8c9751790e2e33f5afb2df0dcfb3aeabe55d5aa2.
//
// Solidity: event MarketComped(address cToken, bool isComped)
func (_Compound *CompoundFilterer) WatchMarketComped(opts *bind.WatchOpts, sink chan<- *CompoundMarketComped) (event.Subscription, error) {

	logs, sub, err := _Compound.contract.WatchLogs(opts, "MarketComped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompoundMarketComped)
				if err := _Compound.contract.UnpackLog(event, "MarketComped", log); err != nil {
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

// ParseMarketComped is a log parse operation binding the contract event 0x93c1f3e36ed71139f466a4ce8c9751790e2e33f5afb2df0dcfb3aeabe55d5aa2.
//
// Solidity: event MarketComped(address cToken, bool isComped)
func (_Compound *CompoundFilterer) ParseMarketComped(log types.Log) (*CompoundMarketComped, error) {
	event := new(CompoundMarketComped)
	if err := _Compound.contract.UnpackLog(event, "MarketComped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompoundMarketEnteredIterator is returned from FilterMarketEntered and is used to iterate over the raw logs and unpacked data for MarketEntered events raised by the Compound contract.
type CompoundMarketEnteredIterator struct {
	Event *CompoundMarketEntered // Event containing the contract specifics and raw log

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
func (it *CompoundMarketEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompoundMarketEntered)
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
		it.Event = new(CompoundMarketEntered)
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
func (it *CompoundMarketEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompoundMarketEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompoundMarketEntered represents a MarketEntered event raised by the Compound contract.
type CompoundMarketEntered struct {
	CToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketEntered is a free log retrieval operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address cToken, address account)
func (_Compound *CompoundFilterer) FilterMarketEntered(opts *bind.FilterOpts) (*CompoundMarketEnteredIterator, error) {

	logs, sub, err := _Compound.contract.FilterLogs(opts, "MarketEntered")
	if err != nil {
		return nil, err
	}
	return &CompoundMarketEnteredIterator{contract: _Compound.contract, event: "MarketEntered", logs: logs, sub: sub}, nil
}

// WatchMarketEntered is a free log subscription operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address cToken, address account)
func (_Compound *CompoundFilterer) WatchMarketEntered(opts *bind.WatchOpts, sink chan<- *CompoundMarketEntered) (event.Subscription, error) {

	logs, sub, err := _Compound.contract.WatchLogs(opts, "MarketEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompoundMarketEntered)
				if err := _Compound.contract.UnpackLog(event, "MarketEntered", log); err != nil {
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

// ParseMarketEntered is a log parse operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address cToken, address account)
func (_Compound *CompoundFilterer) ParseMarketEntered(log types.Log) (*CompoundMarketEntered, error) {
	event := new(CompoundMarketEntered)
	if err := _Compound.contract.UnpackLog(event, "MarketEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompoundMarketExitedIterator is returned from FilterMarketExited and is used to iterate over the raw logs and unpacked data for MarketExited events raised by the Compound contract.
type CompoundMarketExitedIterator struct {
	Event *CompoundMarketExited // Event containing the contract specifics and raw log

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
func (it *CompoundMarketExitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompoundMarketExited)
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
		it.Event = new(CompoundMarketExited)
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
func (it *CompoundMarketExitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompoundMarketExitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompoundMarketExited represents a MarketExited event raised by the Compound contract.
type CompoundMarketExited struct {
	CToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketExited is a free log retrieval operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address cToken, address account)
func (_Compound *CompoundFilterer) FilterMarketExited(opts *bind.FilterOpts) (*CompoundMarketExitedIterator, error) {

	logs, sub, err := _Compound.contract.FilterLogs(opts, "MarketExited")
	if err != nil {
		return nil, err
	}
	return &CompoundMarketExitedIterator{contract: _Compound.contract, event: "MarketExited", logs: logs, sub: sub}, nil
}

// WatchMarketExited is a free log subscription operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address cToken, address account)
func (_Compound *CompoundFilterer) WatchMarketExited(opts *bind.WatchOpts, sink chan<- *CompoundMarketExited) (event.Subscription, error) {

	logs, sub, err := _Compound.contract.WatchLogs(opts, "MarketExited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompoundMarketExited)
				if err := _Compound.contract.UnpackLog(event, "MarketExited", log); err != nil {
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

// ParseMarketExited is a log parse operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address cToken, address account)
func (_Compound *CompoundFilterer) ParseMarketExited(log types.Log) (*CompoundMarketExited, error) {
	event := new(CompoundMarketExited)
	if err := _Compound.contract.UnpackLog(event, "MarketExited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompoundMarketListedIterator is returned from FilterMarketListed and is used to iterate over the raw logs and unpacked data for MarketListed events raised by the Compound contract.
type CompoundMarketListedIterator struct {
	Event *CompoundMarketListed // Event containing the contract specifics and raw log

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
func (it *CompoundMarketListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompoundMarketListed)
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
		it.Event = new(CompoundMarketListed)
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
func (it *CompoundMarketListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompoundMarketListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompoundMarketListed represents a MarketListed event raised by the Compound contract.
type CompoundMarketListed struct {
	CToken common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketListed is a free log retrieval operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address cToken)
func (_Compound *CompoundFilterer) FilterMarketListed(opts *bind.FilterOpts) (*CompoundMarketListedIterator, error) {

	logs, sub, err := _Compound.contract.FilterLogs(opts, "MarketListed")
	if err != nil {
		return nil, err
	}
	return &CompoundMarketListedIterator{contract: _Compound.contract, event: "MarketListed", logs: logs, sub: sub}, nil
}

// WatchMarketListed is a free log subscription operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address cToken)
func (_Compound *CompoundFilterer) WatchMarketListed(opts *bind.WatchOpts, sink chan<- *CompoundMarketListed) (event.Subscription, error) {

	logs, sub, err := _Compound.contract.WatchLogs(opts, "MarketListed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompoundMarketListed)
				if err := _Compound.contract.UnpackLog(event, "MarketListed", log); err != nil {
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

// ParseMarketListed is a log parse operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address cToken)
func (_Compound *CompoundFilterer) ParseMarketListed(log types.Log) (*CompoundMarketListed, error) {
	event := new(CompoundMarketListed)
	if err := _Compound.contract.UnpackLog(event, "MarketListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompoundNewBorrowCapIterator is returned from FilterNewBorrowCap and is used to iterate over the raw logs and unpacked data for NewBorrowCap events raised by the Compound contract.
type CompoundNewBorrowCapIterator struct {
	Event *CompoundNewBorrowCap // Event containing the contract specifics and raw log

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
func (it *CompoundNewBorrowCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompoundNewBorrowCap)
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
		it.Event = new(CompoundNewBorrowCap)
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
func (it *CompoundNewBorrowCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompoundNewBorrowCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompoundNewBorrowCap represents a NewBorrowCap event raised by the Compound contract.
type CompoundNewBorrowCap struct {
	CToken       common.Address
	NewBorrowCap *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewBorrowCap is a free log retrieval operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed cToken, uint256 newBorrowCap)
func (_Compound *CompoundFilterer) FilterNewBorrowCap(opts *bind.FilterOpts, cToken []common.Address) (*CompoundNewBorrowCapIterator, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}

	logs, sub, err := _Compound.contract.FilterLogs(opts, "NewBorrowCap", cTokenRule)
	if err != nil {
		return nil, err
	}
	return &CompoundNewBorrowCapIterator{contract: _Compound.contract, event: "NewBorrowCap", logs: logs, sub: sub}, nil
}

// WatchNewBorrowCap is a free log subscription operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed cToken, uint256 newBorrowCap)
func (_Compound *CompoundFilterer) WatchNewBorrowCap(opts *bind.WatchOpts, sink chan<- *CompoundNewBorrowCap, cToken []common.Address) (event.Subscription, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}

	logs, sub, err := _Compound.contract.WatchLogs(opts, "NewBorrowCap", cTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompoundNewBorrowCap)
				if err := _Compound.contract.UnpackLog(event, "NewBorrowCap", log); err != nil {
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

// ParseNewBorrowCap is a log parse operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed cToken, uint256 newBorrowCap)
func (_Compound *CompoundFilterer) ParseNewBorrowCap(log types.Log) (*CompoundNewBorrowCap, error) {
	event := new(CompoundNewBorrowCap)
	if err := _Compound.contract.UnpackLog(event, "NewBorrowCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompoundNewBorrowCapGuardianIterator is returned from FilterNewBorrowCapGuardian and is used to iterate over the raw logs and unpacked data for NewBorrowCapGuardian events raised by the Compound contract.
type CompoundNewBorrowCapGuardianIterator struct {
	Event *CompoundNewBorrowCapGuardian // Event containing the contract specifics and raw log

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
func (it *CompoundNewBorrowCapGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompoundNewBorrowCapGuardian)
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
		it.Event = new(CompoundNewBorrowCapGuardian)
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
func (it *CompoundNewBorrowCapGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompoundNewBorrowCapGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompoundNewBorrowCapGuardian represents a NewBorrowCapGuardian event raised by the Compound contract.
type CompoundNewBorrowCapGuardian struct {
	OldBorrowCapGuardian common.Address
	NewBorrowCapGuardian common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewBorrowCapGuardian is a free log retrieval operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_Compound *CompoundFilterer) FilterNewBorrowCapGuardian(opts *bind.FilterOpts) (*CompoundNewBorrowCapGuardianIterator, error) {

	logs, sub, err := _Compound.contract.FilterLogs(opts, "NewBorrowCapGuardian")
	if err != nil {
		return nil, err
	}
	return &CompoundNewBorrowCapGuardianIterator{contract: _Compound.contract, event: "NewBorrowCapGuardian", logs: logs, sub: sub}, nil
}

// WatchNewBorrowCapGuardian is a free log subscription operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_Compound *CompoundFilterer) WatchNewBorrowCapGuardian(opts *bind.WatchOpts, sink chan<- *CompoundNewBorrowCapGuardian) (event.Subscription, error) {

	logs, sub, err := _Compound.contract.WatchLogs(opts, "NewBorrowCapGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompoundNewBorrowCapGuardian)
				if err := _Compound.contract.UnpackLog(event, "NewBorrowCapGuardian", log); err != nil {
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

// ParseNewBorrowCapGuardian is a log parse operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_Compound *CompoundFilterer) ParseNewBorrowCapGuardian(log types.Log) (*CompoundNewBorrowCapGuardian, error) {
	event := new(CompoundNewBorrowCapGuardian)
	if err := _Compound.contract.UnpackLog(event, "NewBorrowCapGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompoundNewCloseFactorIterator is returned from FilterNewCloseFactor and is used to iterate over the raw logs and unpacked data for NewCloseFactor events raised by the Compound contract.
type CompoundNewCloseFactorIterator struct {
	Event *CompoundNewCloseFactor // Event containing the contract specifics and raw log

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
func (it *CompoundNewCloseFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompoundNewCloseFactor)
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
		it.Event = new(CompoundNewCloseFactor)
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
func (it *CompoundNewCloseFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompoundNewCloseFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompoundNewCloseFactor represents a NewCloseFactor event raised by the Compound contract.
type CompoundNewCloseFactor struct {
	OldCloseFactorMantissa *big.Int
	NewCloseFactorMantissa *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterNewCloseFactor is a free log retrieval operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Compound *CompoundFilterer) FilterNewCloseFactor(opts *bind.FilterOpts) (*CompoundNewCloseFactorIterator, error) {

	logs, sub, err := _Compound.contract.FilterLogs(opts, "NewCloseFactor")
	if err != nil {
		return nil, err
	}
	return &CompoundNewCloseFactorIterator{contract: _Compound.contract, event: "NewCloseFactor", logs: logs, sub: sub}, nil
}

// WatchNewCloseFactor is a free log subscription operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Compound *CompoundFilterer) WatchNewCloseFactor(opts *bind.WatchOpts, sink chan<- *CompoundNewCloseFactor) (event.Subscription, error) {

	logs, sub, err := _Compound.contract.WatchLogs(opts, "NewCloseFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompoundNewCloseFactor)
				if err := _Compound.contract.UnpackLog(event, "NewCloseFactor", log); err != nil {
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

// ParseNewCloseFactor is a log parse operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Compound *CompoundFilterer) ParseNewCloseFactor(log types.Log) (*CompoundNewCloseFactor, error) {
	event := new(CompoundNewCloseFactor)
	if err := _Compound.contract.UnpackLog(event, "NewCloseFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompoundNewCollateralFactorIterator is returned from FilterNewCollateralFactor and is used to iterate over the raw logs and unpacked data for NewCollateralFactor events raised by the Compound contract.
type CompoundNewCollateralFactorIterator struct {
	Event *CompoundNewCollateralFactor // Event containing the contract specifics and raw log

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
func (it *CompoundNewCollateralFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompoundNewCollateralFactor)
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
		it.Event = new(CompoundNewCollateralFactor)
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
func (it *CompoundNewCollateralFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompoundNewCollateralFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompoundNewCollateralFactor represents a NewCollateralFactor event raised by the Compound contract.
type CompoundNewCollateralFactor struct {
	CToken                      common.Address
	OldCollateralFactorMantissa *big.Int
	NewCollateralFactorMantissa *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterNewCollateralFactor is a free log retrieval operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address cToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Compound *CompoundFilterer) FilterNewCollateralFactor(opts *bind.FilterOpts) (*CompoundNewCollateralFactorIterator, error) {

	logs, sub, err := _Compound.contract.FilterLogs(opts, "NewCollateralFactor")
	if err != nil {
		return nil, err
	}
	return &CompoundNewCollateralFactorIterator{contract: _Compound.contract, event: "NewCollateralFactor", logs: logs, sub: sub}, nil
}

// WatchNewCollateralFactor is a free log subscription operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address cToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Compound *CompoundFilterer) WatchNewCollateralFactor(opts *bind.WatchOpts, sink chan<- *CompoundNewCollateralFactor) (event.Subscription, error) {

	logs, sub, err := _Compound.contract.WatchLogs(opts, "NewCollateralFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompoundNewCollateralFactor)
				if err := _Compound.contract.UnpackLog(event, "NewCollateralFactor", log); err != nil {
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

// ParseNewCollateralFactor is a log parse operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address cToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Compound *CompoundFilterer) ParseNewCollateralFactor(log types.Log) (*CompoundNewCollateralFactor, error) {
	event := new(CompoundNewCollateralFactor)
	if err := _Compound.contract.UnpackLog(event, "NewCollateralFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompoundNewCompRateIterator is returned from FilterNewCompRate and is used to iterate over the raw logs and unpacked data for NewCompRate events raised by the Compound contract.
type CompoundNewCompRateIterator struct {
	Event *CompoundNewCompRate // Event containing the contract specifics and raw log

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
func (it *CompoundNewCompRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompoundNewCompRate)
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
		it.Event = new(CompoundNewCompRate)
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
func (it *CompoundNewCompRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompoundNewCompRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompoundNewCompRate represents a NewCompRate event raised by the Compound contract.
type CompoundNewCompRate struct {
	OldCompRate *big.Int
	NewCompRate *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewCompRate is a free log retrieval operation binding the contract event 0xc227c9272633c3a307d9845bf2bc2509cefb20d655b5f3c1002d8e1e3f22c8b0.
//
// Solidity: event NewCompRate(uint256 oldCompRate, uint256 newCompRate)
func (_Compound *CompoundFilterer) FilterNewCompRate(opts *bind.FilterOpts) (*CompoundNewCompRateIterator, error) {

	logs, sub, err := _Compound.contract.FilterLogs(opts, "NewCompRate")
	if err != nil {
		return nil, err
	}
	return &CompoundNewCompRateIterator{contract: _Compound.contract, event: "NewCompRate", logs: logs, sub: sub}, nil
}

// WatchNewCompRate is a free log subscription operation binding the contract event 0xc227c9272633c3a307d9845bf2bc2509cefb20d655b5f3c1002d8e1e3f22c8b0.
//
// Solidity: event NewCompRate(uint256 oldCompRate, uint256 newCompRate)
func (_Compound *CompoundFilterer) WatchNewCompRate(opts *bind.WatchOpts, sink chan<- *CompoundNewCompRate) (event.Subscription, error) {

	logs, sub, err := _Compound.contract.WatchLogs(opts, "NewCompRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompoundNewCompRate)
				if err := _Compound.contract.UnpackLog(event, "NewCompRate", log); err != nil {
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

// ParseNewCompRate is a log parse operation binding the contract event 0xc227c9272633c3a307d9845bf2bc2509cefb20d655b5f3c1002d8e1e3f22c8b0.
//
// Solidity: event NewCompRate(uint256 oldCompRate, uint256 newCompRate)
func (_Compound *CompoundFilterer) ParseNewCompRate(log types.Log) (*CompoundNewCompRate, error) {
	event := new(CompoundNewCompRate)
	if err := _Compound.contract.UnpackLog(event, "NewCompRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompoundNewLiquidationIncentiveIterator is returned from FilterNewLiquidationIncentive and is used to iterate over the raw logs and unpacked data for NewLiquidationIncentive events raised by the Compound contract.
type CompoundNewLiquidationIncentiveIterator struct {
	Event *CompoundNewLiquidationIncentive // Event containing the contract specifics and raw log

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
func (it *CompoundNewLiquidationIncentiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompoundNewLiquidationIncentive)
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
		it.Event = new(CompoundNewLiquidationIncentive)
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
func (it *CompoundNewLiquidationIncentiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompoundNewLiquidationIncentiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompoundNewLiquidationIncentive represents a NewLiquidationIncentive event raised by the Compound contract.
type CompoundNewLiquidationIncentive struct {
	OldLiquidationIncentiveMantissa *big.Int
	NewLiquidationIncentiveMantissa *big.Int
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterNewLiquidationIncentive is a free log retrieval operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Compound *CompoundFilterer) FilterNewLiquidationIncentive(opts *bind.FilterOpts) (*CompoundNewLiquidationIncentiveIterator, error) {

	logs, sub, err := _Compound.contract.FilterLogs(opts, "NewLiquidationIncentive")
	if err != nil {
		return nil, err
	}
	return &CompoundNewLiquidationIncentiveIterator{contract: _Compound.contract, event: "NewLiquidationIncentive", logs: logs, sub: sub}, nil
}

// WatchNewLiquidationIncentive is a free log subscription operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Compound *CompoundFilterer) WatchNewLiquidationIncentive(opts *bind.WatchOpts, sink chan<- *CompoundNewLiquidationIncentive) (event.Subscription, error) {

	logs, sub, err := _Compound.contract.WatchLogs(opts, "NewLiquidationIncentive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompoundNewLiquidationIncentive)
				if err := _Compound.contract.UnpackLog(event, "NewLiquidationIncentive", log); err != nil {
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

// ParseNewLiquidationIncentive is a log parse operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Compound *CompoundFilterer) ParseNewLiquidationIncentive(log types.Log) (*CompoundNewLiquidationIncentive, error) {
	event := new(CompoundNewLiquidationIncentive)
	if err := _Compound.contract.UnpackLog(event, "NewLiquidationIncentive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompoundNewMaxAssetsIterator is returned from FilterNewMaxAssets and is used to iterate over the raw logs and unpacked data for NewMaxAssets events raised by the Compound contract.
type CompoundNewMaxAssetsIterator struct {
	Event *CompoundNewMaxAssets // Event containing the contract specifics and raw log

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
func (it *CompoundNewMaxAssetsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompoundNewMaxAssets)
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
		it.Event = new(CompoundNewMaxAssets)
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
func (it *CompoundNewMaxAssetsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompoundNewMaxAssetsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompoundNewMaxAssets represents a NewMaxAssets event raised by the Compound contract.
type CompoundNewMaxAssets struct {
	OldMaxAssets *big.Int
	NewMaxAssets *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewMaxAssets is a free log retrieval operation binding the contract event 0x7093cf1eb653f749c3ff531d6df7f92764536a7fa0d13530cd26e070780c32ea.
//
// Solidity: event NewMaxAssets(uint256 oldMaxAssets, uint256 newMaxAssets)
func (_Compound *CompoundFilterer) FilterNewMaxAssets(opts *bind.FilterOpts) (*CompoundNewMaxAssetsIterator, error) {

	logs, sub, err := _Compound.contract.FilterLogs(opts, "NewMaxAssets")
	if err != nil {
		return nil, err
	}
	return &CompoundNewMaxAssetsIterator{contract: _Compound.contract, event: "NewMaxAssets", logs: logs, sub: sub}, nil
}

// WatchNewMaxAssets is a free log subscription operation binding the contract event 0x7093cf1eb653f749c3ff531d6df7f92764536a7fa0d13530cd26e070780c32ea.
//
// Solidity: event NewMaxAssets(uint256 oldMaxAssets, uint256 newMaxAssets)
func (_Compound *CompoundFilterer) WatchNewMaxAssets(opts *bind.WatchOpts, sink chan<- *CompoundNewMaxAssets) (event.Subscription, error) {

	logs, sub, err := _Compound.contract.WatchLogs(opts, "NewMaxAssets")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompoundNewMaxAssets)
				if err := _Compound.contract.UnpackLog(event, "NewMaxAssets", log); err != nil {
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

// ParseNewMaxAssets is a log parse operation binding the contract event 0x7093cf1eb653f749c3ff531d6df7f92764536a7fa0d13530cd26e070780c32ea.
//
// Solidity: event NewMaxAssets(uint256 oldMaxAssets, uint256 newMaxAssets)
func (_Compound *CompoundFilterer) ParseNewMaxAssets(log types.Log) (*CompoundNewMaxAssets, error) {
	event := new(CompoundNewMaxAssets)
	if err := _Compound.contract.UnpackLog(event, "NewMaxAssets", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompoundNewPauseGuardianIterator is returned from FilterNewPauseGuardian and is used to iterate over the raw logs and unpacked data for NewPauseGuardian events raised by the Compound contract.
type CompoundNewPauseGuardianIterator struct {
	Event *CompoundNewPauseGuardian // Event containing the contract specifics and raw log

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
func (it *CompoundNewPauseGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompoundNewPauseGuardian)
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
		it.Event = new(CompoundNewPauseGuardian)
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
func (it *CompoundNewPauseGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompoundNewPauseGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompoundNewPauseGuardian represents a NewPauseGuardian event raised by the Compound contract.
type CompoundNewPauseGuardian struct {
	OldPauseGuardian common.Address
	NewPauseGuardian common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewPauseGuardian is a free log retrieval operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Compound *CompoundFilterer) FilterNewPauseGuardian(opts *bind.FilterOpts) (*CompoundNewPauseGuardianIterator, error) {

	logs, sub, err := _Compound.contract.FilterLogs(opts, "NewPauseGuardian")
	if err != nil {
		return nil, err
	}
	return &CompoundNewPauseGuardianIterator{contract: _Compound.contract, event: "NewPauseGuardian", logs: logs, sub: sub}, nil
}

// WatchNewPauseGuardian is a free log subscription operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Compound *CompoundFilterer) WatchNewPauseGuardian(opts *bind.WatchOpts, sink chan<- *CompoundNewPauseGuardian) (event.Subscription, error) {

	logs, sub, err := _Compound.contract.WatchLogs(opts, "NewPauseGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompoundNewPauseGuardian)
				if err := _Compound.contract.UnpackLog(event, "NewPauseGuardian", log); err != nil {
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

// ParseNewPauseGuardian is a log parse operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Compound *CompoundFilterer) ParseNewPauseGuardian(log types.Log) (*CompoundNewPauseGuardian, error) {
	event := new(CompoundNewPauseGuardian)
	if err := _Compound.contract.UnpackLog(event, "NewPauseGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CompoundNewPriceOracleIterator is returned from FilterNewPriceOracle and is used to iterate over the raw logs and unpacked data for NewPriceOracle events raised by the Compound contract.
type CompoundNewPriceOracleIterator struct {
	Event *CompoundNewPriceOracle // Event containing the contract specifics and raw log

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
func (it *CompoundNewPriceOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CompoundNewPriceOracle)
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
		it.Event = new(CompoundNewPriceOracle)
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
func (it *CompoundNewPriceOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CompoundNewPriceOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CompoundNewPriceOracle represents a NewPriceOracle event raised by the Compound contract.
type CompoundNewPriceOracle struct {
	OldPriceOracle common.Address
	NewPriceOracle common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewPriceOracle is a free log retrieval operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Compound *CompoundFilterer) FilterNewPriceOracle(opts *bind.FilterOpts) (*CompoundNewPriceOracleIterator, error) {

	logs, sub, err := _Compound.contract.FilterLogs(opts, "NewPriceOracle")
	if err != nil {
		return nil, err
	}
	return &CompoundNewPriceOracleIterator{contract: _Compound.contract, event: "NewPriceOracle", logs: logs, sub: sub}, nil
}

// WatchNewPriceOracle is a free log subscription operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Compound *CompoundFilterer) WatchNewPriceOracle(opts *bind.WatchOpts, sink chan<- *CompoundNewPriceOracle) (event.Subscription, error) {

	logs, sub, err := _Compound.contract.WatchLogs(opts, "NewPriceOracle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CompoundNewPriceOracle)
				if err := _Compound.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
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

// ParseNewPriceOracle is a log parse operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Compound *CompoundFilterer) ParseNewPriceOracle(log types.Log) (*CompoundNewPriceOracle, error) {
	event := new(CompoundNewPriceOracle)
	if err := _Compound.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
