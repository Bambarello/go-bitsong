package cli

import (
	"fmt"
	"github.com/bitsongofficial/go-bitsong/x/track/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	defaultPage  = 1
	defaultLimit = 30 // should be consistent with tendermint/tendermint/rpc/core/pipe.go:19
)

func QueryDepositsByTxQuery(cliCtx context.CLIContext, params types.QueryTrackParams) ([]byte, error) {
	events := []string{
		fmt.Sprintf("%s.%s='%s'", sdk.EventTypeMessage, sdk.AttributeKeyAction, types.TypeMsgDeposit),
		fmt.Sprintf("%s.%s='%s'", types.EventTypeDepositTrack, types.AttributeKeyTrackID, []byte(fmt.Sprintf("%d", params.TrackID))),
	}

	// NOTE: SearchTxs is used to facilitate the txs query which does not currently
	// support configurable pagination.
	searchResult, err := utils.QueryTxsByEvents(cliCtx, events, defaultPage, defaultLimit)
	if err != nil {
		return nil, err
	}

	var deposits []types.Deposit

	for _, info := range searchResult.Txs {
		for _, msg := range info.Tx.GetMsgs() {
			if msg.Type() == types.TypeMsgDeposit {
				depMsg := msg.(types.MsgDeposit)

				deposits = append(deposits, types.Deposit{
					Depositor: depMsg.Depositor,
					TrackID:   params.TrackID,
					Amount:    depMsg.Amount,
				})
			}
		}
	}

	if cliCtx.Indent {
		return cliCtx.Codec.MarshalJSONIndent(deposits, "", "  ")
	}

	return cliCtx.Codec.MarshalJSON(deposits)
}
