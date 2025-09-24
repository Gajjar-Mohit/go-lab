package p2p

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T){
	listertAddr := ":4000"
	tr := NewTCPTransport(listertAddr)
	assert.Equal(t, tr.ListenAddress, listertAddr)
	
	//Server
	// tr.Start()
}