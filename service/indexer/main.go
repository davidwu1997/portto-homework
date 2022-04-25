package main

import (
	"time"

	"go.uber.org/zap"

	"portto-homework/pkg/context"
	"portto-homework/pkg/db"
	"portto-homework/pkg/eth"
	"portto-homework/pkg/log"
)

var (
	rpcEndpoint      = "https://data-seed-prebsc-2-s3.binance.org:8545"
	startingBlockNum = uint64(100000)
	workerNum        = 10
)

func main() {
	pg, err := db.NewPostgres()
	if err != nil {
		panic(err)
	}

	eth := eth.New(rpcEndpoint, pg)

	numCh := make(chan uint64)
	for i := 0; i <= workerNum; i++ {
		go worker(numCh, eth)
	}

	for blockNum := startingBlockNum; ; blockNum++ {
		n, err := eth.GetCurrNum(context.Background())
		if err != nil {
			log.Global().Error("eth.GetBlockNum failed")
			continue
		}

		// wait 1 sec if no new blocks to curl
		if n == uint64(blockNum) {
			time.Sleep(time.Second)
			continue
		}

		numCh <- blockNum
	}
}

func worker(numCh <-chan uint64, eth eth.Eth) {
	for n := range numCh {
		ctx := context.Background()
		fetched, err := eth.FetchBlockAndSave(ctx, n)
		if err != nil {
			log.Global().With(zap.Uint64("blockNum", n)).Error("FetchBlockAndSave")
			continue
		}

		if fetched {
			log.Global().With(zap.Uint64("blockNum", n)).Info("fetched ans saved from RPC")
		}
	}
}
