package main

import (
	"errors"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

type Flags struct {
	Port       int
	SolanaPath string
	RpcAddress string
	LogFilter  string
	logs       *log.Logger
}

func NewFlags(logs *log.Logger) *Flags {
	flags := &Flags{logs: logs}

	path := filepath.Join(
		os.Getenv("HOME"),
		".local/share/solana/install/active_release/bin/solana",
	)

	flag.IntVar(&flags.Port, "port", 8312, "Listen port")
	flag.StringVar(&flags.SolanaPath, "solana-path", path, "Solana path")
	flag.StringVar(&flags.RpcAddress, "rpc-address", "http://localhost:8899", "Solana rpc")
	flag.StringVar(&flags.LogFilter, "log-filter", "warn", "Log level. One of: warning, error, info")
	flag.Parse()

	switch flags.LogFilter {
	case "warning":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	default:
		log.SetLevel(log.WarnLevel)
	}

	return flags
}

func (f Flags) Check() error {
	if !Exists(f.SolanaPath) {
		return errors.New(f.SolanaPath + " - not found")
	}

	return nil
}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
