package bitswap

import (
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/littlespeechless/boxo/bitswap/message"
	"github.com/littlespeechless/boxo/bitswap/tracer"
)

type sendOnlyTracer interface {
	MessageSent(peer.ID, message.BitSwapMessage)
}

var _ tracer.Tracer = nopReceiveTracer{}

// we need to only trace sends because we already trace receives in the polyfill object (to not get them traced twice)
type nopReceiveTracer struct {
	sendOnlyTracer
}

func (nopReceiveTracer) MessageReceived(peer.ID, message.BitSwapMessage) {}
