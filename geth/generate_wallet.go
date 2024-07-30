package geth

import (
	"crypto/ecdsa"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func GenWallet() error {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return err
	}

	pkb := crypto.FromECDSA(privateKey)
	fmt.Println("privateKey Byte: ", hexutil.Encode(pkb))

	publicKey := privateKey.Public()
	pbe, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("assert type error")
	}
	pbk := crypto.FromECDSAPub(pbe)
	fmt.Println("Public Key Bytes: ", hexutil.Encode(pbk))

	address := crypto.PubkeyToAddress(*pbe).Hex()
	fmt.Println("address: ", address) // 0x78ced5d2D2Cc80AeB91fC9f279A06E73b41126F8

	hash := sha3.NewLegacyKeccak256()
	hash.Write(pbk[1:])

	fmt.Println("hash: ", hexutil.Encode(hash.Sum(nil)[12:])) // 0x78ced5d2d2cc80aeb91fc9f279a06e73b41126f8

	return nil
}
