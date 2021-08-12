package main

import (
	"solana-metrics/solanarpc"

	log "github.com/sirupsen/logrus"
)

type Solana struct {
	flags  *Flags
	Rpc    *solanarpc.SolanaRpc
	logger *log.Logger
}

func NewSolana(flags *Flags, logger *log.Logger) *Solana {
	return &Solana{
		flags:  flags,
		Rpc:    solanarpc.NewSolanaRpc(flags.RpcAddress, logger),
		logger: logger,
	}
}
