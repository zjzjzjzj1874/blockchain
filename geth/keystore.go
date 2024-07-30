package geth

import (
	"fmt"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func CreateKs(pwd string) error {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)

	account, err := ks.NewAccount(pwd)
	if err != nil {
		return err
	}

	fmt.Println("Account Address: ", account.Address.Hex())
	fmt.Println("Account URL: ", account.URL.String())

	return nil
}

func ImportKs(path, pwd string) error {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jb, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	account, err := ks.Import(jb, pwd, pwd)
	if err != nil {
		return err
	}

	fmt.Println("Account Address: ", account.Address.Hex())
	fmt.Println("Account URL: ", account.URL.String())

	return nil
}
