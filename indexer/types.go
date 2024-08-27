package indexer

import (
	"encoding/json"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/rpcclient"
)

type BTCClient struct {
	url            string
	headers        map[string]string // Use headers to store either cookie or basic auth credentials
	Network        string
	publickeyBytes []byte
	PriceApi       string
	MemPoolApi     string
	QuickNodeApi   string
	RpcClient      *rpcclient.Client
	GraphQLClient  *GraphQLClient
}

type JSONRPCRequest struct {
	ID      int           `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params,omitempty"`
	JSONRPC string        `json:"jsonrpc"`
}

type JSONRPCResponse struct {
	Result json.RawMessage `json:"result"`
	Error  *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	ID int `json:"id"`
}

type BlockHeightResponse struct {
	Height int `json:"height"`
}

type CoinlayerResponse struct {
	Success bool   `json:"success"`
	Target  string `json:"target"`
	Rates   struct {
		BTC float64 `json:"BTC"`
	} `json:"rates"`
}

type MempoolUtxoResponse struct {
	TxId  string `json:"txid"`
	Vout  uint32 `json:"vout"`
	Value uint64 `json:"value"`
}

type QuicknodeUtxoResponse struct {
	TxId  string `json:"txid"`
	Vout  uint32 `json:"vout"`
	Value string `json:"value"`
}

type QuicknodeUtxo struct {
	TxId  string `json:"txid"`
	Vout  uint32 `json:"vout"`
	Value uint64 `json:"value"`
}

type quicknodeUtxoRequest struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

type quicknodeUtxoResult struct {
	ID      *int                    `json:"id"`
	Result  []QuicknodeUtxoResponse `json:"result"`
	Jsonrpc string                  `json:"jsonrpc"`
}

type BTCParsedTx struct {
	TxHash                     string
	Confirmations              uint64
	Vins                       []btcjson.Vin
	Vouts                      []btcjson.Vout
	VinTxId                    string //record VinTxId & VoutIndex here, so that we can get sender using VinTxId in batch
	VoutIndex                  uint32
	Recipients                 []string
	RecipientsAndAmounts       map[string]uint64
	RecipientsAmountsVoutIndex map[string]uint32
	Network                    string
	BlockHash                  string
	Time                       int64
}
