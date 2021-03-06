package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

/****
 * Track Msg
 ***/

// Track messages types and routes
const (
	TypeMsgCreateTrack = "create_track"
	TypeMsgPlayTrack   = "play_track"
	TypeMsgDeposit     = "deposit"
)

/****************************************
 * MsgCreateTrack
 ****************************************/

var _ sdk.Msg = MsgCreateTrack{}

// MsgCreateTrack defines CreateTrack message
type MsgCreateTrack struct {
	Title       string         `json:"title" yaml:"title"` // Track title
	Audio       string         `json:"audio" yaml:"audio"`
	Image       string         `json:"image" yaml:"image"`
	Duration    string         `json:"duration" yaml:"duration"`
	Hidden      bool           `json:"hidden" yaml:"hidden"`
	Explicit    bool           `json:"explicit" yaml:"explicit"`
	Genre       string         `json:"genre" yaml:"genre"`
	Mood        string         `json:"mood" yaml:"mood"`
	Artists     string         `json:"artists" yaml:"artists"`
	Featuring   string         `json:"featuring" yaml:"featuring"`
	Producers   string         `json:"producers" yaml:"producers"`
	Description string         `json:"description" yaml:"description"`
	Copyright   string         `json:"copyright" yaml:"copyright"`
	Owner       sdk.AccAddress `json:"owner" yaml:"owner"` // Track owner
}

func NewMsgCreateTrack(title, audio, image, duration string, hidden, explicit bool, genre, mood, artists, featuring, producers, description, copyright string, owner sdk.AccAddress) MsgCreateTrack {
	return MsgCreateTrack{
		Title:       title,
		Audio:       audio,
		Image:       image,
		Duration:    duration,
		Hidden:      hidden,
		Explicit:    explicit,
		Genre:       genre,
		Mood:        mood,
		Artists:     artists,
		Featuring:   featuring,
		Producers:   producers,
		Description: description,
		Copyright:   copyright,
		Owner:       owner,
	}
}

//nolint
func (msg MsgCreateTrack) Route() string { return RouterKey }
func (msg MsgCreateTrack) Type() string  { return TypeMsgCreateTrack }

// ValidateBasic
func (msg MsgCreateTrack) ValidateBasic() sdk.Error {
	// TODO:
	// - Add more check

	if len(strings.TrimSpace(msg.Title)) == 0 {
		return ErrInvalidTrackTitle(DefaultCodespace, "track title cannot be blank")
	}

	if len(msg.Title) > MaxTitleLength {
		return ErrInvalidTrackTitle(DefaultCodespace, fmt.Sprintf("track title is longer than max length of %d", MaxTitleLength))
	}

	if len(strings.TrimSpace(msg.Audio)) == 0 {
		return ErrInvalidTrackTitle(DefaultCodespace, "track audio cannot be blank")
	}

	if len(strings.TrimSpace(msg.Image)) == 0 {
		return ErrInvalidTrackTitle(DefaultCodespace, "track image cannot be blank")
	}

	if len(strings.TrimSpace(msg.Duration)) == 0 {
		return ErrInvalidTrackTitle(DefaultCodespace, "track duration cannot be blank")
	}

	if len(strings.TrimSpace(msg.Genre)) == 0 {
		return ErrInvalidTrackTitle(DefaultCodespace, "track genre cannot be blank")
	}

	/*if len(strings.TrimSpace(msg.Mood)) == 0 {
		return ErrInvalidTrackTitle(DefaultCodespace, "track mood cannot be blank")
	}*/

	if len(strings.TrimSpace(msg.Artists)) == 0 {
		return ErrInvalidTrackTitle(DefaultCodespace, "track artists cannot be blank")
	}

	/*if len(strings.TrimSpace(msg.Featuring)) == 0 {
		return ErrInvalidTrackTitle(DefaultCodespace, "track featuring cannot be blank")
	}

	if len(strings.TrimSpace(msg.Producers)) == 0 {
		return ErrInvalidTrackTitle(DefaultCodespace, "track producers cannot be blank")
	}

	if len(strings.TrimSpace(msg.Description)) == 0 {
		return ErrInvalidTrackTitle(DefaultCodespace, "track description cannot be blank")
	}*/

	if len(msg.Description) > MaxDescriptionLength {
		return ErrInvalidTrackTitle(DefaultCodespace, fmt.Sprintf("track description is longer than max length of %d", MaxDescriptionLength))
	}

	/*if len(strings.TrimSpace(msg.Description)) == 0 {
		return ErrInvalidTrackTitle(DefaultCodespace, "track description cannot be blank")
	}*/

	if len(msg.Copyright) > MaxCopyrightLength {
		return ErrInvalidTrackTitle(DefaultCodespace, fmt.Sprintf("track copyright is longer than max length of %d", MaxCopyrightLength))
	}

	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}

	return nil
}

// Implements Msg.
func (msg MsgCreateTrack) String() string {
	return fmt.Sprintf(`Create Track Message:
  Title: %s
  Address: %s
`, msg.Title, msg.Owner.String())
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateTrack) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateTrack) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

/****************************************
 * MsgPlay
 ****************************************/

var _ sdk.Msg = MsgPlay{}

// MsgPlay defines PlayTrack message
type MsgPlay struct {
	TrackID uint64         "json:`track_id` yaml:`track_id`"
	AccAddr sdk.AccAddress `json:"acc_addr"`
}

func NewMsgPlay(trackID uint64, accAddr sdk.AccAddress) MsgPlay {
	return MsgPlay{
		TrackID: trackID,
		AccAddr: accAddr,
	}
}

//nolint
func (msg MsgPlay) Route() string { return RouterKey }
func (msg MsgPlay) Type() string  { return TypeMsgPlayTrack }

// ValidateBasic
func (msg MsgPlay) ValidateBasic() sdk.Error {
	// TODO:
	// - improve check

	if msg.TrackID == 0 {
		return ErrUnknownTrack(DefaultCodespace, "album-id cannot be blank")
	}

	if msg.AccAddr.Empty() {
		return sdk.ErrInvalidAddress(msg.AccAddr.String())
	}

	return nil
}

// Implements Msg.
func (msg MsgPlay) String() string {
	return fmt.Sprintf(`Play Track Message:
  TrackID: %d
  AccAddr: %s
`, msg.TrackID, msg.AccAddr.String())
}

// GetSignBytes encodes the message for signing
func (msg MsgPlay) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgPlay) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.AccAddr}
}

/****************************************
 * MsgDeposit
 ****************************************/

var _ sdk.Msg = MsgDeposit{}

type MsgDeposit struct {
	TrackID   uint64         `json:"track_id" yaml:"track_id"`   // ID
	Depositor sdk.AccAddress `json:"depositor" yaml:"depositor"` // Address of the depositor
	Amount    sdk.Coins      `json:"amount" yaml:"amount"`       // Coins to add to the proposal's deposit
}

func NewMsgDeposit(depositor sdk.AccAddress, trackID uint64, amount sdk.Coins) MsgDeposit {
	return MsgDeposit{trackID, depositor, amount}
}

// Implements Msg.
// nolint
func (msg MsgDeposit) Route() string { return RouterKey }
func (msg MsgDeposit) Type() string  { return TypeMsgDeposit }

// Implements Msg.
func (msg MsgDeposit) ValidateBasic() sdk.Error {
	if msg.Depositor.Empty() {
		return sdk.ErrInvalidAddress(msg.Depositor.String())
	}
	if !msg.Amount.IsValid() {
		return sdk.ErrInvalidCoins(msg.Amount.String())
	}
	if msg.Amount.IsAnyNegative() {
		return sdk.ErrInvalidCoins(msg.Amount.String())
	}

	return nil
}

func (msg MsgDeposit) String() string {
	return fmt.Sprintf(`Deposit Message:
  Depositer:   %s
  Track ID: %d
  Amount:      %s
`, msg.Depositor, msg.TrackID, msg.Amount)
}

// Implements Msg.
func (msg MsgDeposit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// Implements Msg.
func (msg MsgDeposit) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Depositor}
}
