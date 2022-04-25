package eth

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"

	"portto-homework/pkg/context"
)

var (
	ErrNotFound = ethereum.NotFound
)

type Eth interface {
	GetCurrNum(ctx context.Context) (uint64, error)
	GetBlocks(ctx context.Context, n uint64) ([]*Block, error)
	GetBlock(ctx context.Context, hash common.Hash) (*Block, error)
	GetTransaction(ctx context.Context, txHash common.Hash) (*Transaction, error)

	// FetchBlockAndSave fetch block from RPC and save if block not exist in db.
	// Return true if block is fetched from RPC, and false if block is saved already.
	FetchBlockAndSave(ctx context.Context, blockNum uint64) (bool, error)
}
