package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	cdptypes "github.com/kava-labs/kava/x/cdp/types"
)

// SupplyKeeper defines the expected supply keeper for module accounts
type SupplyKeeper interface {
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}

// CdpKeeper defines the expected cdp keeper for interacting with cdps
type CdpKeeper interface {
	IterateCdpsByDenom(ctx sdk.Context, denom string, cb func(cdp cdptypes.CDP) (stop bool))
	GetTotalPrincipal(ctx sdk.Context, collateralDenom string, principalDenom string) (total sdk.Int)
}

// AccountKeeper defines the expected keeper interface for interacting with account
type AccountKeeper interface {
	GetModuleAddress(name string) sdk.AccAddress
	GetModuleAccount(ctx sdk.Context, moduleName string) authtypes.ModuleAccountI

	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	SetAccount(ctx sdk.Context, acc authtypes.AccountI)
}
