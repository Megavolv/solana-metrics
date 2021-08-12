package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	logs := log.New()
	logs.SetLevel(log.InfoLevel)
	logs.SetOutput(os.Stdout)

	flags := NewFlags(logs)

	/*if err := flags.Check(); err != nil {
		logs.Error(err)
		os.Exit(1)
	}*/

	solana := NewSolana(flags, logs)
	ok, err := solana.Rpc.HealthCheck()
	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}
	logs.Info(ok)

	data, err := solana.Rpc.RpcVersion()
	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}

	logs.Info(data.FeatureSet)
	logs.Info(data.SolanaCore)

	solana.Rpc.RpcGetBalance()

}

/*
solana address Gbwg3HCbD9gzma2o6oTqYkDQnojeZ1z7Ygc95DPFA2me
solana block-height 72915100
solana block-production - таблица. Identity Leader Slots  Blocks Produced Skipped Slots Skipped Slot Percentage
solana block-time Block: 89179713, Date: 2021-08-07T13:30:05Z
solana catchup
solana cluster-version 1.7.10
solana epoch 219
solana epoch-info
	Block height: 72915392
	Slot: 89180003
	Epoch: 219
	Transaction Count: 31544139256
	Epoch Slot Range: [89084256..89516256)
	Epoch Completed Percent: 22.164%
	Epoch Completed Slots: 95747/432000 (336253 remaining)
	Epoch Completed Time: 15h 57m 31s/2days 22h 35m 59s (2days 6h 38m 28s remaining)

solana slot
solana stakes - очень много данных
solana validator-info get - много данных. Про каждого валидатора. Можно выборочно указать инфо ключ


solana validator-info get | grep Gbwg3HCbD9gzma2o6oTqYkDQnojeZ1z7Ygc95DPFA2me -A 4
Validator Identity: Gbwg3HCbD9gzma2o6oTqYkDQnojeZ1z7Ygc95DPFA2me
  Info Address: 6gKoy9jReuZYXYCeXeQaAR9T8Aoi2ZNTJd7St3HvKZww
  Keybase Username: ftxkiwi
  Name: Peters testnet validator

solana validator-info get 6gKoy9jReuZYXYCeXeQaAR9T8Aoi2ZNTJd7St3HvKZww
Validator Identity: Gbwg3HCbD9gzma2o6oTqYkDQnojeZ1z7Ygc95DPFA2me
  Info Address: 6gKoy9jReuZYXYCeXeQaAR9T8Aoi2ZNTJd7St3HvKZww
  Keybase Username: ftxkiwi
  Name: Peters testnet validator

solana validators Identity                                      Vote Account                            Commission  Last Vote       Root Slot    Skip Rate  Credits  Version  Active Stake



solana gossip IP Address      | Node identifier                              | Gossip | TPU   | RPC Address           | Version
solana leader-schedule | grep BuFKun14y2uagDQXKou6x4ArWuG46US7X8TEEJVvSMTq | wc -l
*/
