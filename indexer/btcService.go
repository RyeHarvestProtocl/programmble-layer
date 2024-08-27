package indexer

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/RyeHarvestProtocol/programmable-layer/config"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/keyfuse/tokucore/network"
	"github.com/keyfuse/tokucore/xcore"
	"github.com/keyfuse/tokucore/xcrypto"
)

func New(config *config.Config) (*BTCClient, error) {
	publickeyInBytes, err := hex.DecodeString(config.BTCRpcClient.Publickey)
	if err != nil {
		return nil, fmt.Errorf("error reading public key: %w", err)
	}

	// create graph ql client
	graphqlClient := NewGraphQLClient(config.BTCRpcClient.TheGraphApi, 3)
	// call the graph ql api to make sure it works

	client := &BTCClient{
		url:            config.BTCRpcClient.Host,
		headers:        make(map[string]string),
		Network:        config.BTCRpcClient.Network,
		publickeyBytes: publickeyInBytes,
		PriceApi:       config.BTCRpcClient.PriceApi,
		MemPoolApi:     config.BTCRpcClient.MemPoolApi,
		QuickNodeApi:   config.BTCRpcClient.QuickNodeApi,
		GraphQLClient:  graphqlClient,
	}

	if config.BTCRpcClient.CookiePath != "" {
		cookieData, err := os.ReadFile(config.BTCRpcClient.CookiePath)
		if err != nil {
			return nil, fmt.Errorf("error reading cookie file: %w", err)
		}

		// Assuming the format "__cookie__:some_random_value"
		cookieParts := strings.SplitN(string(cookieData), ":", 2)
		if len(cookieParts) == 2 {
			client.headers["Authorization"] = "Basic " + basicAuth("__cookie__", cookieParts[1])
		} else {
			return nil, fmt.Errorf("invalid cookie file format")
		}
	} else {
		client.headers["Authorization"] = "Basic " + basicAuth(config.BTCRpcClient.User, config.BTCRpcClient.Pass)
	}

	connCfg := &rpcclient.ConnConfig{
		Host:         "nd-948-043-849.p2pify.com",
		User:         config.BTCRpcClient.User,
		Pass:         config.BTCRpcClient.Pass,
		DisableTLS:   true,
		HTTPPostMode: true,
	}
	rpcClient, err := rpcclient.New(connCfg, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating rpc client: %w", err)
	}

	client.RpcClient = rpcClient

	return client, nil
}

func (c *BTCClient) BatchGetBlockVerboseTx(blockNumbers []uint64) ([]btcjson.GetBlockVerboseTxResult, error) {
	// Step 1: Batch call to getblockhash
	hashRequests := make([]JSONRPCRequest, len(blockNumbers))
	for i, blockNumber := range blockNumbers {
		hashRequests[i] = JSONRPCRequest{
			ID:      i,
			Method:  "getblockhash",
			Params:  []interface{}{blockNumber},
			JSONRPC: "2.0",
		}
	}

	hashResultTemplates := make([]interface{}, len(hashRequests))
	for i := range hashResultTemplates {
		hashResultTemplates[i] = "" // Expecting a string result for the hash
	}
	hashesInterface, err := c.callBatchWithRetry(hashRequests, hashResultTemplates)
	if err != nil {
		fmt.Println("error when get hashes: ", err)
		return nil, err
	}

	hashes := make([]string, len(hashesInterface))
	for i, v := range hashesInterface {
		hashes[i] = v.(string) // Type assertion, assuming the result is a string
	}
	// fmt.Println("this is hashes: ", hashes)

	// Step 2: Use the hashes to batch call getblock with verbose=2
	blockRequests := make([]JSONRPCRequest, len(hashes))
	for i, hash := range hashes {
		blockRequests[i] = JSONRPCRequest{
			ID:      i,
			Method:  "getblock",
			Params:  []interface{}{hash, 2},
			JSONRPC: "2.0",
		}
	}

	blockResultTemplates := make([]interface{}, len(blockRequests))
	for i := range blockResultTemplates {
		blockResultTemplates[i] = btcjson.GetBlockVerboseTxResult{}
	}
	blocksInterface, err := c.callBatchWithRetry(blockRequests, blockResultTemplates)
	if err != nil {
		return nil, err
	}

	blocks := make([]btcjson.GetBlockVerboseTxResult, len(blocksInterface))
	for i, v := range blocksInterface {
		blocks[i] = v.(btcjson.GetBlockVerboseTxResult) // Type assertion
	}

	return blocks, nil
}

// get a summary of addressAmounts
func (c *BTCClient) GetTransferDetailsByTxId(txDetails btcjson.TxRawResult, blockHash string) (*BTCParsedTx, error) {
	addressAmounts := make(map[string]uint64)
	addressAmountsVoutIndex := make(map[string]uint32)

	// parse vout to get to address and amount
	for _, vout := range txDetails.Vout {
		scriptPubKey := vout.ScriptPubKey
		script, err := hex.DecodeString(scriptPubKey.Hex)
		if err != nil {
			fmt.Println(err)
			continue
		}
		var chainParams chaincfg.Params
		if c.Network == "mainnet" {
			chainParams = chaincfg.MainNetParams
		} else {
			chainParams = chaincfg.TestNet3Params
		}

		// extract an address lists here
		_, addrs, _, err := txscript.ExtractPkScriptAddrs(script, &chainParams)
		if err != nil {
			fmt.Println("err when ExtractPkScriptAddrs: ", err)
			continue
		}

		if len(addrs) > 0 {
			for _, addr := range addrs {
				received := vout.Value
				// format the amount
				receivedBTCAmount, err := btcutil.NewAmount(received)
				if err != nil {
					return nil, fmt.Errorf("error converting amount to BTC: %w", err)
				}
				receivedSatAmount := receivedBTCAmount.ToUnit(btcutil.AmountSatoshi)
				// fmt.Println("✅receivedSatAmount: ", receivedSatAmount)

				addressAmounts[addr.EncodeAddress()] = uint64(receivedSatAmount)

				addressAmountsVoutIndex[addr.EncodeAddress()] = vout.N
			}
		}
	}
	// fmt.Println("===addressAmounts: ", addressAmounts)
	// fmt.Println("===txDetails.Vin[0].Txid: ", txDetails.Vin[0].Txid)

	return &BTCParsedTx{
		TxHash:                     txDetails.Txid,
		Confirmations:              txDetails.Confirmations,
		Vins:                       txDetails.Vin,
		Vouts:                      txDetails.Vout,
		VinTxId:                    txDetails.Vin[0].Txid, // todo: we only check the first vin here. but if user build a specific tx whose input includes some other utxos it might be a problem
		VoutIndex:                  txDetails.Vin[0].Vout,
		RecipientsAndAmounts:       addressAmounts,
		RecipientsAmountsVoutIndex: addressAmountsVoutIndex,
		Network:                    c.Network,
		BlockHash:                  txDetails.BlockHash,
	}, nil
}

// ParseOpReturnData parses OP_RETURN data from a transaction
func (c *BTCClient) ParseOpReturnData(tx *BTCParsedTx) []string {
	var opReturnData []string

	for _, vout := range tx.Vouts {
		if vout.ScriptPubKey.Type == "nulldata" {
			// Extract the OP_RETURN data. The hex field contains the entire script in hex format.
			dataHex := vout.ScriptPubKey.Hex
			data, err := hex.DecodeString(dataHex)
			if err != nil {
				fmt.Printf("Error decoding OP_RETURN data: %v\n", err)
				continue
			}
			// Strip OP_RETURN (0x6a) and possible PUSHDATA (0x4c) opcode(s)
			if len(data) > 0 && data[0] == 0x6a {
				data = data[1:]
				// Handle possible PUSHDATA1 (0x4c) opcode
				if len(data) > 0 && data[0] == 0x4c {
					data = data[2:] // Skip PUSHDATA1 length byte
				}
			}
			opReturnData = append(opReturnData, string(data))
		}
	}

	return opReturnData
}

// GetSenderByVinTxId:
// Parse the first vin
// If there are more than 1 user sender address in the tx, the same invoice tx will be recorded as valid deposit for more than 1 time.
// To avoid this, we simply assume the first vin to as sender
func (c *BTCClient) GetSenderByVinTxId(vinTxId string, voutIndex uint32) (string, error) {
	if vinTxId == "" {
		return "0xCoinbase", nil
	} else {
		vinTxDetails, err := c.GetTxDetailsByTxIdAndBlockHash(vinTxId, "")
		if err != nil {
			return "", err
		}
		// get sender address from vin scriptPubKey
		if int(voutIndex) < len(vinTxDetails.Vout) {
			refVout := vinTxDetails.Vout[voutIndex]
			if len(refVout.ScriptPubKey.Address) > 0 {
				senderAddress := refVout.ScriptPubKey.Address
				return senderAddress, nil
			} else {
				return "", errors.New("no addresses found in scriptPubKey")
			}
		} else {
			fmt.Println("Vout index out of range")
			return "", errors.New("vout index out of range")
		}
	}
}

func (c *BTCClient) GetSenderByVinTxDetailsAndVoutIndex(vinTxId string, voutIndex uint32) (string, error) {
	if vinTxId == "" {
		return "0xCoinbase", nil
	} else {
		vinTxDetails, err := c.GetTxDetailsByTxIdAndBlockHash(vinTxId, "")
		if err != nil {
			return "", err
		}
		// get sender address from vin scriptPubKey
		if int(voutIndex) < len(vinTxDetails.Vout) {
			refVout := vinTxDetails.Vout[voutIndex]
			if len(refVout.ScriptPubKey.Address) > 0 {
				senderAddress := refVout.ScriptPubKey.Address
				return senderAddress, nil
			} else {
				return "", errors.New("no addresses found in scriptPubKey")
			}
		} else {
			fmt.Println("Vout index out of range")
			return "", errors.New("vout index out of range")
		}
	}
}

func (c *BTCClient) GetTxDetailsByTxIdAndBlockHash(txId string, blockHash string) (*btcjson.TxRawResult, error) {
	var params []interface{}
	if blockHash != "" {
		params = []interface{}{txId, true, blockHash}
	} else {
		params = []interface{}{txId, true}
	}

	request := JSONRPCRequest{
		ID:     1, // Arbitrary ID, but should be unique per request
		Method: "getrawtransaction",
		Params: params,
	}

	var txDetails *btcjson.TxRawResult

	// Execute the RPC call
	_, err := c.callWithRetry(request, &txDetails)
	if err != nil {
		return nil, fmt.Errorf("error getting transaction details for txId %s and blockHash %s: %w", txId, blockHash, err)
	}

	return txDetails, nil
}

// returns a map: txId => TxRawResult
func (c *BTCClient) BatchGetTxDetailsByTxIdsAndBlockHashes(txIds []string, blockHashes []string) (map[string]*btcjson.TxRawResult, error) {
	if len(txIds) != len(blockHashes) {
		return nil, fmt.Errorf("length of txIds and blockHashes must be equal")
	}

	requests := make([]JSONRPCRequest, len(txIds))
	for i := range txIds {
		var params []interface{}
		if blockHashes[i] != "" {
			params = []interface{}{txIds[i], true, blockHashes[i]}
		} else {
			params = []interface{}{txIds[i], true}
		}

		requests[i] = JSONRPCRequest{
			ID:      i,
			Method:  "getrawtransaction",
			Params:  params,
			JSONRPC: "2.0",
		}
	}

	resultTemplates := make([]interface{}, len(requests))
	for i := range resultTemplates {
		resultTemplates[i] = &btcjson.TxRawResult{}
	}

	resultsInterface, err := c.callBatchWithRetry(requests, resultTemplates)
	if err != nil {
		return nil, fmt.Errorf("error getting transaction details: %w", err)
	}

	results := make([]*btcjson.TxRawResult, len(resultsInterface))
	for i, v := range resultsInterface {
		results[i] = v.(*btcjson.TxRawResult) // Type assertion
	}

	if len(txIds) != len(results) {
		return nil, fmt.Errorf("length of txIds and results must be equal")
	}

	var resultMap = make(map[string]*btcjson.TxRawResult)
	for _, result := range results {
		resultMap[result.Txid] = result
	}

	return resultMap, nil
}

func (c *BTCClient) GetTxInsSumAmount(vins []btcjson.Vin) (uint64, error) {
	// for each vin, get the tx details
	sumVinValue := uint64(0)
	for _, vin := range vins {
		vinTxId := vin.Txid
		voutIndex := vin.Vout
		vinTxDetails, err := c.GetTxDetailsByTxIdAndBlockHash(vinTxId, "")
		if err != nil {
			fmt.Println("error when get vin tx details: ", err)
			return 0, err
		}
		if int(voutIndex) < len(vinTxDetails.Vout) {
			refVout := vinTxDetails.Vout[voutIndex]
			amount, err := btcutil.NewAmount(refVout.Value)
			if err != nil {
				return 0, fmt.Errorf("error converting amount to BTC: %w", err)
			}

			sumVinValue += uint64(amount.ToUnit(btcutil.AmountSatoshi))
		}
	}
	return sumVinValue, nil
}

// below are p2wsh functions

// GenerateP2wshAddress generates multisig wallet using P2WSH:
// 1/2 p2msh
// 1. user's pub key 2. our pub key
// Param: network: mainnet, testnet
func (c *BTCClient) GenerateP2WSHAddress(userPubkey string, networkString string) string {
	redeemScript := xcore.NewPayToMultiSigScript(1, []byte("03164e3563fe204cc8841b1e2b8ee4282ed9b8b14c032fc55fe8ee0fbc4471af7c"), []byte("03edfdd728085f48fc6a2a8b14505cda36eeb3727c455b0832d1a77d322e53c4f3"))
	redeem, _ := redeemScript.GetLockingScriptBytes()
	// fmt.Println("==== this is redeem script: ", string(redeem))

	multi := xcore.NewPayToWitnessV0ScriptHashAddress(xcrypto.Sha256(redeem))
	if networkString == "mainnet" {
		return multi.ToString(network.MainNet)
	} else {
		return multi.ToString(network.TestNet)
	}
}

func (c *BTCClient) BuildMultiSigP2WSHAddrUsingPublicKey(pubKeyStr string, networkString string) (string, error) {
	pk2, err := hex.DecodeString(pubKeyStr) // user's public key
	if err != nil {
		return "", err
	}

	// Create redeem script for 2 of 3 multi-sig
	builder := txscript.NewScriptBuilder()
	// Add the minimum number of needed signatures
	builder.AddOp(txscript.OP_1)
	// Add the 3 public keys
	builder.AddData(c.publickeyBytes).AddData(pk2)
	// Add the total number of public keys in the multi-sig script
	builder.AddOp(txscript.OP_2)
	// Add the check-multi-sig op-code
	builder.AddOp(txscript.OP_CHECKMULTISIG)
	// Redeem script is the script program in the format of []byte
	redeemScript, err := builder.Script()
	if err != nil {
		return "", err
	}

	// Calculate the SHA256 of the redeem script for P2WSH
	witnessScriptHash := sha256.Sum256(redeemScript)

	// Create a P2WSH address from the witness script hash
	// If using Bitcoin main net, then pass &chaincfg.MainNetParams as second argument
	var network *chaincfg.Params
	if networkString == "mainnet" {
		network = &chaincfg.MainNetParams
	} else {
		network = &chaincfg.TestNet3Params
	}

	addr, err := btcutil.NewAddressWitnessScriptHash(witnessScriptHash[:], network)
	if err != nil {
		return "", err
	}

	return addr.EncodeAddress(), nil
}

// only returns UTXOs for addresses that are part of the wallet managed by the node
func (c *BTCClient) GetWalletUTXOs(addresses []string, minConf, maxConf int, minAmount float64) ([]*btcjson.ListUnspentResult, error) {
	// Define the JSON RPC request to fetch UTXOs
	request := JSONRPCRequest{
		ID:     2, // Arbitrary ID, different from the one used in GetLatestBlockHeight
		Method: "listunspent",
		Params: []interface{}{
			minConf,
			maxConf,
			addresses,
			true, // include_unsafe: Include outputs that are not safe to spend
			map[string]interface{}{
				"minimumAmount": "0", // Minimum amount of UTXOs to include
			},
		},
	}

	var utxos []*btcjson.ListUnspentResult

	// Execute the RPC call
	_, err := c.callWithRetry(request, &utxos)
	if err != nil {
		return nil, fmt.Errorf("error getting UTXOs: %w", err)
	}

	return utxos, nil
}

func (c *BTCClient) CheckIfConfirmed(txId string) (bool, error) {
	txDetails, err := c.GetTxDetailsByTxIdAndBlockHash(txId, "")
	if err != nil {
		return false, err
	}

	if txDetails.Confirmations > 0 {
		return true, nil
	}

	return false, nil
}

func (c *BTCClient) GetBTCPrice() (float64, error) {
	// Make the HTTP GET request
	var result CoinlayerResponse
	if err := c.callEndpoint(c.PriceApi, "GET", nil, &result); err != nil {
		fmt.Printf("Error: %v\n", err)
		return 0, err
	} else {
		fmt.Printf("Bitcoin Price: %f\n", result.Rates.BTC)
		return result.Rates.BTC, nil
	}
}

func (c *BTCClient) GetBTCPriceByTheGraph() (float64, error) {
	// Make the HTTP GET request
	var result CoinlayerResponse
	if err := c.callEndpoint(c.PriceApi, "GET", nil, &result); err != nil {
		fmt.Printf("Error: %v\n", err)
		return 0, err
	} else {
		fmt.Printf("Bitcoin Price: %f\n", result.Rates.BTC)
		return result.Rates.BTC, nil
	}
}

func (c *BTCClient) GetUtxo(address string) ([]MempoolUtxoResponse, error) {
	// Make the HTTP GET request
	var result []MempoolUtxoResponse
	if err := c.callEndpoint(fmt.Sprintf("%s/address/%s/utxo", c.MemPoolApi, address), "GET", nil, &result); err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	} else {
		return result, nil
	}
}

func (c *BTCClient) GetUtxoQuickNode(address string) ([]QuicknodeUtxo, error) {
	// Check if QuickNodeApi is properly set
	if c.QuickNodeApi == "" {
		return nil, fmt.Errorf("QuickNodeApi is not set")
	}

	// Prepare the request payload
	payload := quicknodeUtxoRequest{
		Method: "bb_getutxos",
		Params: []interface{}{address, map[string]bool{"confirmed": false}},
	}

	// Marshal the payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Error marshalling payload: %v\n", err)
		return nil, err
	}

	// Create the POST request
	req, err := http.NewRequest("POST", c.QuickNodeApi, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making HTTP request: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read and parse the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return nil, err
	}

	var result quicknodeUtxoResult
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return nil, err
	}

	// Convert value from string to uint64 and populate the result slice
	var utxos []QuicknodeUtxo
	for _, utxo := range result.Result {
		value, err := strconv.ParseUint(utxo.Value, 10, 64)
		if err != nil {
			fmt.Printf("Error converting value to uint64: %v\n", err)
			return nil, err
		}
		utxos = append(utxos, QuicknodeUtxo{
			TxId:  utxo.TxId,
			Vout:  utxo.Vout,
			Value: value,
		})
	}

	return utxos, nil
}

func (c *BTCClient) GetLatestBlockHeight() (int, error) {
	request := JSONRPCRequest{
		ID:     1, // Arbitrary ID, but should be unique per request
		Method: "getblockcount",
		Params: []interface{}{},
	}

	var blockHeight int

	// Execute the RPC call
	_, err := c.callWithRetry(request, &blockHeight)
	if err != nil {
		return 0, fmt.Errorf("error getting latest block height: %w", err)
	}

	return blockHeight, nil
}

func (c *BTCClient) SendRawTransaction(tx *wire.MsgTx, allowHighFees bool) (*chainhash.Hash, error) {
	return c.RpcClient.SendRawTransaction(tx, allowHighFees)
}

func (c *BTCClient) callWithRetry(request JSONRPCRequest, resultTemplate interface{}) (interface{}, error) {
	var err error
	maxRetries := 5
	waitTime := 500 * time.Millisecond // Initial wait time of 500ms

	for attempt := 1; attempt <= maxRetries; attempt++ {
		var result interface{}
		result, err = c.call(request, resultTemplate)
		if err == nil {
			return result, nil
		}

		fmt.Printf("Attempt %d failed, error: %v, request json rpc: %s, request params: %s \n", attempt, err, request.JSONRPC, request.Params)
		if attempt < maxRetries {
			time.Sleep(waitTime)
			waitTime *= 2 // Double the wait time for the next attempt
		}
	}

	return nil, fmt.Errorf("after %d attempts, last error: %w", maxRetries, err)
}

func (c *BTCClient) callBatchWithRetry(requests []JSONRPCRequest, resultTemplate []interface{}) ([]interface{}, error) {
	var results []interface{}
	var err error
	maxRetries := 5
	waitTime := 500 * time.Millisecond // Initial wait time of 500ms

	for attempt := 1; attempt <= maxRetries; attempt++ {
		results, err = c.batchCall(requests, resultTemplate)
		if err == nil {
			return results, nil
		}

		fmt.Printf("Batch attempt %d failed, error: %v\n", attempt, err)
		if attempt < maxRetries {
			time.Sleep(waitTime)
			waitTime *= 2 // Double the wait time for the next attempt
		}
	}

	return nil, fmt.Errorf("after %d batch attempts, last error: %w", maxRetries, err)
}

func (c *BTCClient) call(request JSONRPCRequest, resultTemplate interface{}) (interface{}, error) {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	req, err := http.NewRequest("POST", c.url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("====received non-200 response status: %d, error: %s", resp.StatusCode, resp.Body)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error when get response data: %v , error: %v\n", responseBody, err)
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	// fmt.Printf("Raw response body: %s\n", responseBody) // Log the raw response body

	var response JSONRPCResponse
	err = json.Unmarshal(responseBody, &response)

	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	if response.Error != nil {
		return nil, fmt.Errorf("error in response: %s", response.Error.Message)
	}

	// Directly unmarshal the result into the provided resultTemplate pointer.
	err = json.Unmarshal(response.Result, resultTemplate)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling result: %w", err)
	}

	return resultTemplate, nil
}

func (c *BTCClient) batchCall(requests []JSONRPCRequest, resultTemplate []interface{}) ([]interface{}, error) {
	requestBody, err := json.Marshal(requests)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	req, err := http.NewRequest("POST", c.url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var responses []JSONRPCResponse
	err = json.Unmarshal(responseBody, &responses)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	results := make([]interface{}, len(responses))
	for i, response := range responses {
		if response.Error != nil {
			return nil, fmt.Errorf("error in response: %s", response.Error.Message)
		}

		resultType := reflect.TypeOf(resultTemplate[i])
		resultPtr := reflect.New(resultType).Interface()

		err := json.Unmarshal(response.Result, &resultPtr)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling result for request %d: %w", i, err)
		}

		results[i] = reflect.ValueOf(resultPtr).Elem().Interface()
	}

	return results, nil
}

func (client *BTCClient) callEndpoint(apiEndpoint string, method string, payload interface{}, result interface{}) error {
	// Convert the payload to a JSON byte slice
	var body []byte
	var err error
	if payload != nil {
		body, err = json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("error marshaling payload: %w", err)
		}
	}

	// Maximum number of retries
	const maxRetries = 3
	// Initial backoff interval in seconds
	const backoffInterval = 2

	// Attempt to make the request with retries
	for i := 0; i <= maxRetries; i++ {
		req, err := http.NewRequest(method, apiEndpoint, bytes.NewReader(body))
		if err != nil {
			return fmt.Errorf("error creating request: %w", err)
		}

		// Perform the request
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			if i < maxRetries {
				time.Sleep(time.Duration(backoffInterval) * time.Second)
				continue // Retry the request
			}
			return fmt.Errorf("error sending request: %w", err)
		}

		defer resp.Body.Close()

		// Check if the response status code is OK
		if resp.StatusCode != http.StatusOK {
			if i < maxRetries {
				time.Sleep(time.Duration(backoffInterval) * time.Second)
				continue // Retry the request
			}
			bodyBytes, _ := io.ReadAll(resp.Body)
			return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, bodyBytes)
		}

		// Decode the response body into the result interface
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			time.Sleep(3 * time.Second)
			if i < maxRetries {
				time.Sleep(time.Duration(backoffInterval) * time.Second)
				continue // Retry the request
			}
			return fmt.Errorf("error decoding response: %w", err)
		}

		return nil // Success, exit the loop
	}

	return fmt.Errorf("exceeded maximum retries")
}

// func getSenderAddresses(txId string) ([]string, error) {
// 	var senderAddresses []string
// 	for _, vin := range txDetails.Vin {
// 			refTxDetails, err := queryTxDetails(vin.Txid)
// 			if err != nil {
// 					return nil, err
// 			}
// 			senderAddress := refTxDetails.Vout[vin.VoutIndex].ScriptPubKey.Addresses[0]
// 			senderAddresses = append(senderAddresses, senderAddress)
// 	}
// 	return senderAddresses, nil
// }

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
