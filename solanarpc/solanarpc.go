package solanarpc

import (
	"bytes"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type SolanaRpc struct {
	address string
	logger  *log.Logger
	caller  func(*bytes.Buffer) (*http.Response, error)
}

func NewSolanaRpc(address string, logger *log.Logger) *SolanaRpc {
	c := &SolanaRpc{
		address: address,
		logger:  logger,
	}
	c.caller = c.rpcCall
	return c
}

func (r SolanaRpc) rpcCall(data *bytes.Buffer) (*http.Response, error) {
	return http.Post(r.address, "application/json", data)
}
