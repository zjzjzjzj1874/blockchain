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
		//0x71c7656ec7ab88b098defb751b7401b5f6d8976f => 自己的小狐狸钱包账号
		//0x88b4B3cF8bb8EE4defB5535e2De512c52e7BA5F0 => 官方Demo测试用账号

		// 获取账户余额
		//client.BalanceByAccount("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
		ba, err := client.BalanceByAccount("0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(ba.String())

		// 获取区块余额
		fmt.Println(client.BalanceInBlock("0x71c7656ec7ab88b098defb751b7401b5f6d8976f", 5532993))

		// 获取待处理余额
		fmt.Println(client.PendingBalanceAt("0x71c7656ec7ab88b098defb751b7401b5f6d8976f"))
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
}
