package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jasonw/delp/server/reviews"
)

func main() {
	client, err := ethclient.Dial("http://192.168.1.14:8545")
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.HexToECDSA("c018ebe754eea79fc412fd1ec370d8f92d9b22e92487eafc9d7a13ae964df015")
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
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
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(1000000)

	address, tx, instance, err := reviews.DeployReviews(auth, client)
	if err != nil {
		panic(err)
	}

	fmt.Println(address.Hex())

	_, _ = instance, tx
}
