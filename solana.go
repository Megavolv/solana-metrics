package main

import (
	log "github.com/sirupsen/logrus"
)

type Solana struct {
	flags      *Flags
	Connection *Connection
	logs       *log.Logger
}

func NewSolana(flags *Flags, logs *log.Logger) *Solana {
	return &Solana{
		flags:      flags,
		Connection: NewConnection(flags.RpcAddress),
		logs:       logs,
	}
}
