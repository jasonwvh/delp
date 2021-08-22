package handlers

import (
	"github.com/jasonw/delp/server/blockchain"
)

type Reviews struct {
	c *blockchain.Connection
}

func NewReviews(c *blockchain.Connection) *Reviews {
	return &Reviews{c}
}
