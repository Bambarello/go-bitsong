package track

import (
	"github.com/bitsongofficial/go-bitsong/x/track/keeper"
	"github.com/bitsongofficial/go-bitsong/x/track/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey

	DefaultParamspace = types.DefaultParamspace
)

var (
	// Keeper methods
	NewKeeper  = keeper.NewKeeper
	NewHandler = keeper.NewHandler
	NewQuerier = keeper.NewQuerier

	// Codec
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec

	// Types
	NewTrack = types.NewTrack

	// Msgs
	NewMsgCreateTrack = types.NewMsgCreateTrack
)

type (
	// Keeper
	Keeper = keeper.Keeper

	// Types
	TrackStatus = types.TrackStatus
	Track       = types.Track
	Tracks      = types.Tracks

	Deposits      = types.Deposits
	DepositParams = types.DepositParams

	// Msgs
	MsgCreateTrack = types.MsgCreateTrack
)
