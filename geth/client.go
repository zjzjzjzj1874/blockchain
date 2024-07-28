package geth

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	Client *ethclient.Client
}

// NewClient creates a new client url => https://cloudflare-eth.com
func NewClient(url string) (*Client, error) {
	client, err := ethclient.Dial(url)
	return &Client{client}, err
}
