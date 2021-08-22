package app

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/mux"
	"github.com/jasonw/delp/server/blockchain"
	"github.com/jasonw/delp/server/data"
)

var privateKey = "c018ebe754eea79fc412fd1ec370d8f92d9b22e92487eafc9d7a13ae964df015"

type App struct {
	Router     *mux.Router
	Connection *blockchain.Connection
}

func (app *App) SetupApp() {
	getR := app.Router.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/reviews/{id:[0-9]+}", app.GetReviewsById)
	getR.HandleFunc("/reviews", app.GetReviews)

	postR := app.Router.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/reviews", app.PostReview)

}

func (app *App) GetReviewsById(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	id := getID(r)

	reply, err := app.Connection.Instance.GetReviewsById(&bind.CallOpts{}, id)
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(rw).Encode(reply)
}

func (app *App) GetReviews(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	count, err := app.Connection.Instance.ReviewCount(nil)
	if err != nil {
		panic(err)
	}
	var i uint8
	for i = 1; i <= count; i++ {
		reply, err := app.Connection.Instance.GetReviewsById(&bind.CallOpts{}, i)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(rw).Encode(reply)
	}

}

func (app *App) PostReview(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)

	var review data.Review
	err := decoder.Decode(&review)
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		panic(err)
	}
	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := app.Connection.Ethclient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	gasPrice, err := app.Connection.Ethclient.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	app.Connection.Auth.Nonce = big.NewInt(int64(nonce))
	app.Connection.Auth.GasPrice = gasPrice

	json.NewEncoder(rw).Encode(review)
	fmt.Println("gas: ", gasPrice)
	tx, err := app.Connection.Instance.CreateReview(app.Connection.Auth, review.Content, review.City)
	if err != nil {
		panic(err)
	}

	fmt.Println("tx sent: %s", tx.Hash().Hex())

}

func getID(r *http.Request) uint8 {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	idd := uint8(id)

	return idd
}
