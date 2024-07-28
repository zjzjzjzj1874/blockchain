package geth

import (
	"context"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// BalanceByAccount 使用Account获取账户余额
func (c *Client) BalanceByAccount(account string) (*big.Float, error) {
	ac := common.HexToAddress(account)
	balance, err := c.Client.BalanceAt(context.Background(), ac, nil)
	if err != nil {
		return nil, err
	}

	// 以太坊中的数字是使用尽可能小的单位来处理的，因为它们是定点精度，在ETH中它是wei。要读取ETH值，您必须做计算wei/10^18
	fb := new(big.Float)
	fb.SetString(balance.String())
	ethVal := new(big.Float).Quo(fb, big.NewFloat(math.Pow10(18)))

	return ethVal, nil
}

// BalanceInBlock 区块余额
func (c *Client) BalanceInBlock(account string, blockNumber int64) (*big.Float, error) {
	balance, err := c.Client.BalanceAt(context.Background(), common.HexToAddress(account), big.NewInt(blockNumber))
	if err != nil {
		return nil, err
	}

	// 以太坊中的数字是使用尽可能小的单位来处理的，因为它们是定点精度，在ETH中它是wei。要读取ETH值，您必须做计算wei/10^18
	fb := new(big.Float)
	fb.SetString(balance.String())
	ethVal := new(big.Float).Quo(fb, big.NewFloat(math.Pow10(18)))

	return ethVal, nil
}

// PendingBalanceAt 返回待处理余额
func (c *Client) PendingBalanceAt(account string) (*big.Float, error) {
	ba, err := c.Client.PendingBalanceAt(context.Background(), common.HexToAddress(account))
	if err != nil {
		return nil, err
	}

	fb := new(big.Float)
	fb.SetString(ba.String())

	return new(big.Float).Quo(fb, big.NewFloat(math.Pow10(18))), nil
}
