package structs

import "github.com/matrixbotio/exchange-gates-lib/internal/structs"

// APICredentialsTypeKeypair - public and private key pair
var APICredentialsTypeKeypair APICredentialsType = "keypair"

// APICredentialsType - API credentials type ^ↀᴥↀ^
type APICredentialsType string

// APICredentials - data for authorization to the exchange API
type APICredentials struct {
	Type APICredentialsType `json:"type"`

	Keypair  APIKeypair          `json:"keypair"`
	Password structs.APIPassword `json:"password"`
	Email    structs.APIEmail    `json:"email"`
}

// WorkerChannels - channels container to control the worker
type WorkerChannels struct {
	WsDone chan struct{}
	WsStop chan struct{}
}

// CheckOrdersResponse - data on checked and restored orders
type CheckOrdersResponse struct {
	ExecutedOrders  []structs.OrderData
	CancelledOrders []structs.OrderData
}

// APIKeypair - data for authorization via public and private keys
type APIKeypair struct {
	Public string `json:"public"`
	Secret string `json:"secret"`
}
