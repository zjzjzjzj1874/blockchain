package geth

import (
	"context"
	"fmt"
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

// TransactionIterateTx 遍历交易事务
func (c *Client) TransactionIterateTx(ctx context.Context, blockNumber int64) error {
	block, err := c.Client.BlockByNumber(ctx, big.NewInt(blockNumber))
	if err != nil {
		return err
	}

	for idx, tx := range block.Transactions() {
		if idx > 20 {
			break
		}
		fmt.Println("hash: ", tx.Hash().Hex())
		fmt.Println("value: ", tx.Value().String())
		fmt.Println("gas: ", tx.Gas())
		fmt.Println("gas price: ", tx.GasPrice().String())
		fmt.Println("Nonce: ", tx.Nonce())
		fmt.Println("data: ", string(tx.Data()))
		fmt.Println("to: ", tx.To().Hex())

		chainId, err := c.Client.NetworkID(ctx)
		if err != nil {
			return err
		}
		fmt.Println("chainId: ", chainId.String())
		receipt, err := c.Client.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			return err
		}

		fmt.Println("receipt status: ", receipt.Status) // 1
	}

	return nil
}
