package server

import (
	"github.com/ipfs/go-bitswap/server/internal/decision"
)

type (
	Receipt                = decision.Receipt
	PeerBlockRequestFilter = decision.PeerBlockRequestFilter
	TaskComparator         = decision.TaskComparator
	ScoreLedger            = decision.ScoreLedger
	ScorePeerFunc          = decision.ScorePeerFunc
)
