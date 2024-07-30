package geth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// Transfer privateKey: e788e315941db0d191c1ecc31ef7cbc7aab42b47ac06688c49978fe7b9ece3d4
// to: 0x71c7656ec7ab88b098defb751b7401b5f6d8976f
func (c *Client) Transfer(ctx context.Context, privateKeyBytes, toAddr string) error {
	privateKey, err := crypto.HexToECDSA(privateKeyBytes)
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	pbe, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("cannot assert type")
	}
	fromAddr := crypto.PubkeyToAddress(*pbe)
	nonce, err := c.Client.PendingNonceAt(ctx, fromAddr)
	if err != nil {
		return err
	}
	fmt.Println("nonce: ", nonce)
	value := big.NewInt(1000000000000000000) // 1eth = 10^18 wei
	gasLimit := uint64(21000)
	gasPrice, err := c.Client.SuggestGasPrice(ctx)
	if err != nil {
		return err
	}

	fmt.Println("gasPrice: ", gasPrice.String())
	toAddress := common.HexToAddress(toAddr)
	var data []byte

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainId, err := c.Client.NetworkID(ctx)
	if err != nil {
		return err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		return err
	}

	fmt.Println("signedTx: ", signedTx.Hash().Hex())
	err = c.Client.SendTransaction(ctx, signedTx)
	if err != nil {
		return err
	}

	fmt.Println("tx send: ", signedTx.Hash().Hex())

	return nil
}
