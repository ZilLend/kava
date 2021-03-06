package cdp

// DO NOT EDIT - generated by aliasgen tool (github.com/rhuairahrighairidh/aliasgen)

import (
	"github.com/kava-labs/kava/x/cdp/keeper"
	"github.com/kava-labs/kava/x/cdp/types"
)

const (
	BaseDigitFactor                         = keeper.BaseDigitFactor
	AttributeKeyCdpID                       = types.AttributeKeyCdpID
	AttributeKeyDeposit                     = types.AttributeKeyDeposit
	AttributeKeyError                       = types.AttributeKeyError
	AttributeValueCategory                  = types.AttributeValueCategory
	DefaultParamspace                       = types.DefaultParamspace
	EventTypeBeginBlockerFatal              = types.EventTypeBeginBlockerFatal
	EventTypeCdpClose                       = types.EventTypeCdpClose
	EventTypeCdpDeposit                     = types.EventTypeCdpDeposit
	EventTypeCdpDraw                        = types.EventTypeCdpDraw
	EventTypeCdpLiquidation                 = types.EventTypeCdpLiquidation
	EventTypeCdpRepay                       = types.EventTypeCdpRepay
	EventTypeCdpWithdrawal                  = types.EventTypeCdpWithdrawal
	EventTypeCreateCdp                      = types.EventTypeCreateCdp
	LiquidatorMacc                          = types.LiquidatorMacc
	ModuleName                              = types.ModuleName
	QuerierRoute                            = types.QuerierRoute
	QueryGetAccounts                        = types.QueryGetAccounts
	QueryGetCdp                             = types.QueryGetCdp
	QueryGetCdpDeposits                     = types.QueryGetCdpDeposits
	QueryGetCdps                            = types.QueryGetCdps
	QueryGetCdpsByCollateralType            = types.QueryGetCdpsByCollateralType
	QueryGetCdpsByCollateralization         = types.QueryGetCdpsByCollateralization
	QueryGetParams                          = types.QueryGetParams
	QueryGetPreviousSavingsDistributionTime = types.QueryGetPreviousSavingsDistributionTime
	QueryGetSavingsRateDistributed          = types.QueryGetSavingsRateDistributed
	RestCollateralType                      = types.RestCollateralType
	RestOwner                               = types.RestOwner
	RestRatio                               = types.RestRatio
	RouterKey                               = types.RouterKey
	SavingsRateMacc                         = types.SavingsRateMacc
	StoreKey                                = types.StoreKey
)

var (
	// function aliases
	FilterCDPs                         = keeper.FilterCDPs
	FindIntersection                   = keeper.FindIntersection
	NewKeeper                          = keeper.NewKeeper
	NewQuerier                         = keeper.NewQuerier
	CdpKey                             = types.CdpKey
	CollateralRatioBytes               = types.CollateralRatioBytes
	CollateralRatioIterKey             = types.CollateralRatioIterKey
	CollateralRatioKey                 = types.CollateralRatioKey
	DefaultGenesisState                = types.DefaultGenesisState
	DefaultParams                      = types.DefaultParams
	DenomIterKey                       = types.DenomIterKey
	DepositIterKey                     = types.DepositIterKey
	DepositKey                         = types.DepositKey
	GetCdpIDBytes                      = types.GetCdpIDBytes
	GetCdpIDFromBytes                  = types.GetCdpIDFromBytes
	NewAugmentedCDP                    = types.NewAugmentedCDP
	NewCDP                             = types.NewCDP
	NewCDPWithFees                     = types.NewCDPWithFees
	NewCollateralParam                 = types.NewCollateralParam
	NewDebtParam                       = types.NewDebtParam
	NewDeposit                         = types.NewDeposit
	NewGenesisState                    = types.NewGenesisState
	NewMsgCreateCDP                    = types.NewMsgCreateCDP
	NewMsgDeposit                      = types.NewMsgDeposit
	NewMsgDrawDebt                     = types.NewMsgDrawDebt
	NewMsgRepayDebt                    = types.NewMsgRepayDebt
	NewMsgWithdraw                     = types.NewMsgWithdraw
	NewParams                          = types.NewParams
	NewQueryCdpDeposits                = types.NewQueryCdpDeposits
	NewQueryCdpParams                  = types.NewQueryCdpParams
	NewQueryCdpsByCollateralTypeParams = types.NewQueryCdpsByCollateralTypeParams
	NewQueryCdpsByRatioParams          = types.NewQueryCdpsByRatioParams
	NewQueryCdpsParams                 = types.NewQueryCdpsParams
	ParamKeyTable                      = types.ParamKeyTable
	ParseDecBytes                      = types.ParseDecBytes
	RegisterCodec                      = types.RegisterCodec
	RelativePow                        = types.RelativePow
	SortableDecBytes                   = types.SortableDecBytes
	SplitCdpKey                        = types.SplitCdpKey
	SplitCollateralRatioIterKey        = types.SplitCollateralRatioIterKey
	SplitCollateralRatioKey            = types.SplitCollateralRatioKey
	SplitDenomIterKey                  = types.SplitDenomIterKey
	SplitDepositIterKey                = types.SplitDepositIterKey
	SplitDepositKey                    = types.SplitDepositKey
	ValidSortableDec                   = types.ValidSortableDec

	// variable aliases
	CdpIDKey                            = types.CdpIDKey
	CdpIDKeyPrefix                      = types.CdpIDKeyPrefix
	CdpKeyPrefix                        = types.CdpKeyPrefix
	CollateralRatioIndexPrefix          = types.CollateralRatioIndexPrefix
	DebtDenomKey                        = types.DebtDenomKey
	DefaultCdpStartingID                = types.DefaultCdpStartingID
	DefaultCircuitBreaker               = types.DefaultCircuitBreaker
	DefaultCollateralParams             = types.DefaultCollateralParams
	DefaultDebtDenom                    = types.DefaultDebtDenom
	DefaultDebtLot                      = types.DefaultDebtLot
	DefaultDebtParam                    = types.DefaultDebtParam
	DefaultDebtThreshold                = types.DefaultDebtThreshold
	DefaultGlobalDebt                   = types.DefaultGlobalDebt
	DefaultGovDenom                     = types.DefaultGovDenom
	DefaultPreviousDistributionTime     = types.DefaultPreviousDistributionTime
	DefaultSavingsDistributionFrequency = types.DefaultSavingsDistributionFrequency
	DefaultSavingsRateDistributed       = types.DefaultSavingsRateDistributed
	DefaultStableDenom                  = types.DefaultStableDenom
	DefaultSurplusLot                   = types.DefaultSurplusLot
	DefaultSurplusThreshold             = types.DefaultSurplusThreshold
	DepositKeyPrefix                    = types.DepositKeyPrefix
	ErrBelowDebtFloor                   = types.ErrBelowDebtFloor
	ErrCdpAlreadyExists                 = types.ErrCdpAlreadyExists
	ErrCdpNotAvailable                  = types.ErrCdpNotAvailable
	ErrCdpNotFound                      = types.ErrCdpNotFound
	ErrCollateralNotSupported           = types.ErrCollateralNotSupported
	ErrDebtNotSupported                 = types.ErrDebtNotSupported
	ErrDenomPrefixNotFound              = types.ErrDenomPrefixNotFound
	ErrDepositNotAvailable              = types.ErrDepositNotAvailable
	ErrDepositNotFound                  = types.ErrDepositNotFound
	ErrExceedsDebtLimit                 = types.ErrExceedsDebtLimit
	ErrInvalidCollateral                = types.ErrInvalidCollateral
	ErrInvalidCollateralLength          = types.ErrInvalidCollateralLength
	ErrInvalidCollateralRatio           = types.ErrInvalidCollateralRatio
	ErrInvalidDebtRequest               = types.ErrInvalidDebtRequest
	ErrInvalidDeposit                   = types.ErrInvalidDeposit
	ErrInvalidPayment                   = types.ErrInvalidPayment
	ErrInvalidWithdrawAmount            = types.ErrInvalidWithdrawAmount
	ErrLoadingAugmentedCDP              = types.ErrLoadingAugmentedCDP
	ErrPricefeedDown                    = types.ErrPricefeedDown
	GovDenomKey                         = types.GovDenomKey
	KeyCircuitBreaker                   = types.KeyCircuitBreaker
	KeyCollateralParams                 = types.KeyCollateralParams
	KeyDebtLot                          = types.KeyDebtLot
	KeyDebtParam                        = types.KeyDebtParam
	KeyDebtThreshold                    = types.KeyDebtThreshold
	KeyDistributionFrequency            = types.KeyDistributionFrequency
	KeyGlobalDebtLimit                  = types.KeyGlobalDebtLimit
	KeySavingsRateDistributed           = types.KeySavingsRateDistributed
	KeySurplusLot                       = types.KeySurplusLot
	KeySurplusThreshold                 = types.KeySurplusThreshold
	MaxSortableDec                      = types.MaxSortableDec
	ModuleCdc                           = types.ModuleCdc
	PreviousDistributionTimeKey         = types.PreviousDistributionTimeKey
	PricefeedStatusKeyPrefix            = types.PricefeedStatusKeyPrefix
	PrincipalKeyPrefix                  = types.PrincipalKeyPrefix
	SavingsRateDistributedKey           = types.SavingsRateDistributedKey
)

type (
	Keeper                          = keeper.Keeper
	AccountKeeper                   = types.AccountKeeper
	AuctionKeeper                   = types.AuctionKeeper
	AugmentedCDP                    = types.AugmentedCDP
	AugmentedCDPs                   = types.AugmentedCDPs
	CDP                             = types.CDP
	CDPs                            = types.CDPs
	CollateralParam                 = types.CollateralParam
	CollateralParams                = types.CollateralParams
	DebtParam                       = types.DebtParam
	DebtParams                      = types.DebtParams
	Deposit                         = types.Deposit
	Deposits                        = types.Deposits
	GenesisState                    = types.GenesisState
	MsgCreateCDP                    = types.MsgCreateCDP
	MsgDeposit                      = types.MsgDeposit
	MsgDrawDebt                     = types.MsgDrawDebt
	MsgRepayDebt                    = types.MsgRepayDebt
	MsgWithdraw                     = types.MsgWithdraw
	Params                          = types.Params
	PricefeedKeeper                 = types.PricefeedKeeper
	QueryCdpDeposits                = types.QueryCdpDeposits
	QueryCdpParams                  = types.QueryCdpParams
	QueryCdpsByCollateralTypeParams = types.QueryCdpsByCollateralTypeParams
	QueryCdpsByRatioParams          = types.QueryCdpsByRatioParams
	QueryCdpsParams                 = types.QueryCdpsParams
	SupplyKeeper                    = types.SupplyKeeper
)
