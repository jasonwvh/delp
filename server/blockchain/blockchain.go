package blockchain

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jasonw/delp/server/reviews"
)

type Connection struct {
	Ethclient *ethclient.Client
	Instance  *reviews.Reviews
	Auth      *bind.TransactOpts
}

func NewConnection(_contractAddress string, _privateKey string, _ethNetwork string) (*Connection, error) {
	client, err := ethclient.Dial(_ethNetwork)
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.HexToECDSA(_privateKey)
	if err != nil {
		panic(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}

	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)

	instance, err := reviews.NewReviews(common.HexToAddress(_contractAddress), client)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to blockchain")

	return &Connection{
		Ethclient: client,
		Instance:  instance,
		Auth:      auth,
	}, nil
}
