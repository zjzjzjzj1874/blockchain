package geth

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

// 初始化一个全局的client
var (
	client, err = NewClient("https://cloudflare-eth.com")
	ctx         = context.Background()
)

func TestClient_Balance(t *testing.T) {
	t.Run("#CLIENT", func(t *testing.T) {

		// 获取账户余额
		ba, err := client.BalanceByAccount("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("balance: ", ba.String())

		// 获取区块余额
		//fmt.Sprintf("区块余额: %v, err: %v", client.BalanceInBlock("0x71c7656ec7ab88b098defb751b7401b5f6d8976f", 5532993))

		// 获取待处理余额
		//fmt.Println("待处理余额： ", client.PendingBalanceAt("0x71c7656ec7ab88b098defb751b7401b5f6d8976f"))
	})
}

// 转账
func TestClient_Transfer(t *testing.T) {
	t.Run("#Transfer", func(t *testing.T) {
		// 获取账户余额
		err := client.Transfer(ctx, "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19", "0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
		if err != nil {
			log.Fatal(err)
		}
	})
}

func TestClient_Block(t *testing.T) {
	t.Run("#Block", func(t *testing.T) {
		// 获取区块头信息
		header, err := client.HeaderByNumber(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(header.Number.String())
		fmt.Println(string(header.Extra))
		fmt.Println(time.Unix(int64(header.Time), 0).Format(time.RFC3339))

		// 获取完整区块信息
		block, err := client.BlockByNumber(ctx, 5671744)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(block.Number().String())
		fmt.Println(string(block.Extra()))
		fmt.Println(time.Unix(int64(block.Time()), 0).Format(time.RFC3339))
		fmt.Println(block.Difficulty().String())
		fmt.Println(block.Hash().Hex())
		fmt.Println(block.Transactions().Len())
		fmt.Println(block.Transactions()[0].To().Hex()) // 发送到的账户 =>  0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

		count, err := client.TransactionCount(ctx, block.Hash())
		if err != nil {
			log.Fatal()
		}

		fmt.Println("Count =", count)
	})
	t.Run("#TX", func(t *testing.T) {
		err = client.TransactionIterateTx(ctx, 5671744)
		if err != nil {
			log.Fatal()
		}
	})
}

// 创建钱包
func TestClient_GenWallet(t *testing.T) {
	t.Run("#GenWallet", func(t *testing.T) {
		err := GenWallet()
		if err != nil {
			log.Fatal(err)
		}
	})
}

// 创建KeyStore
func TestClient_KeyStore(t *testing.T) {
	t.Run("#KeyStore", func(t *testing.T) {
		//err := CreateKs()
		file := "./tmp/UTC--2024-07-30T07-20-47.679277000Z--85c0a461989069b3e3c524dcfbb9cf84f1150bb1"
		pwd := "secret"
		err := ImportKs(file, pwd)
		if err != nil {
			log.Fatal(err)
		}
	})
}
