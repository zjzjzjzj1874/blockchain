package geth

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// 查询区块信息

// HeaderByNumber 获取区块头信息
func (c *Client) HeaderByNumber(ctx context.Context) (*types.Header, error) {
	header, err := c.Client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	return header, nil
}

// BlockByNumber 获取完整区块信息
func (c *Client) BlockByNumber(ctx context.Context, blockNumber int64) (*types.Block, error) {
	block, err := c.Client.BlockByNumber(ctx, big.NewInt(blockNumber))
	if err != nil {
		return nil, err
	}

	return block, nil
}

// TransactionCount 获取交易数量
func (c *Client) TransactionCount(ctx context.Context, blockHash common.Hash) (int64, error) {
	count, err := c.Client.TransactionCount(ctx, blockHash)
	if err != nil {
		return 0, err
	}

	return int64(count), nil
}
