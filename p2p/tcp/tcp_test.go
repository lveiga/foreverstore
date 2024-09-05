package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listAddress := ":3000"
	tr := NewTCPTransport(listAddress)
	assert.Equal(t, tr.listenAddress, listAddress)
	assert.Nil(t, tr.ListenAndAccept())
}
