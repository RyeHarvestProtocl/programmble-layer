package evmClient

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/RyeHarvestProtocol/programmable-layer/config"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EvmContractClient struct {
	Client                 *ethclient.Client
	HarvestContractAbi     *abi.ABI
	RyeContractAbi         *abi.ABI
	RyeContractAddress     string
	HarvestContractAddress string
	IndexerPrivateKey      string
	IndexerAddress         string
}

func NewEvmContractClient(config *config.Config) (*EvmContractClient, error) {
	fmt.Print("this is url")
	fmt.Println("this is url: ", config.EvmClient.RpcUrl)
	client, err := ethclient.Dial(config.EvmClient.RpcUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		return nil, err
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}
	fmt.Println("Current working directory:", cwd)

	// read minting abi
	harvestAbiJSON, err := os.ReadFile("/app/Minting.json")
	if err != nil {
		log.Fatalf("Failed to read ABI file: %v", err)
		return nil, err
	}

	// Check the format of the ABI JSON
	var harvestAbiStructure []interface{}
	if err := json.Unmarshal(harvestAbiJSON, &harvestAbiStructure); err != nil {
		log.Fatalf("Invalid ABI JSON format: %v", err)
		return nil, err
	}

	// Parse the ABI
	harvestParsedABI, err := abi.JSON(strings.NewReader(string(harvestAbiJSON)))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
		return nil, err
	}

	// below is rye abi
	ryeAbiJSON, err := os.ReadFile("./MockERC20.json")
	if err != nil {
		log.Fatalf("Failed to read ABI file: %v", err)
		return nil, err
	}

	// Check the format of the ABI JSON
	var ryeAbiStructure []interface{}
	if err := json.Unmarshal(ryeAbiJSON, &ryeAbiStructure); err != nil {
		log.Fatalf("Invalid ABI JSON format: %v", err)
		return nil, err
	}

	// Parse the ABI
	ryeParsedABI, err := abi.JSON(strings.NewReader(string(ryeAbiJSON)))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
		return nil, err
	}

	return &EvmContractClient{
		Client:                 client,
		HarvestContractAbi:     &harvestParsedABI,
		RyeContractAbi:         &ryeParsedABI,
		RyeContractAddress:     config.EvmClient.RyeContractAddress,
		HarvestContractAddress: config.EvmClient.HarvestContractAddress,
		IndexerPrivateKey:      config.EvmClient.IndexerPrivateKey,
		IndexerAddress:         config.EvmClient.IndexerAddress,
	}, nil
}

// mint rye
func (evmContractClient *EvmContractClient) SendMintRyeTransaction(userAddressStr string, indexerPrivateKeyStr string, amount uint64) (string, error) {
	privateKey, err := crypto.HexToECDSA(indexerPrivateKeyStr)
	if err != nil {
		log.Fatalf("Failed to convert private key: %v", err)
		return "", err
	}

	ryeContractAddress := common.HexToAddress(evmContractClient.RyeContractAddress)

	userAddress := common.HexToAddress(userAddressStr)
	amountBigInt := big.NewInt(1000 * 10 * 18)

	txData, err := evmContractClient.RyeContractAbi.Pack("mint", userAddress, amountBigInt)
	if err != nil {
		log.Fatalf("Failed to pack data for transaction: %v", err)
		return "", err
	}

	if err != nil {
		log.Fatalf("Failed to pack data for transaction: %v", err)
		return "", err
	}

	fromAddress := common.HexToAddress(evmContractClient.IndexerAddress)
	nonce, err := evmContractClient.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
		return "", err
	}

	senderAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// Check if the derived address matches the expected sender address
	if senderAddress != common.HexToAddress(evmContractClient.IndexerAddress) {
		log.Fatalf("Address derived from private key does not match the expected sender address")
		return "", fmt.Errorf("invalid sender address")
	}

	gasLimit := uint64(20000000)
	gasPrice, err := evmContractClient.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
		return "", err
	}

	tx := types.NewTransaction(nonce, ryeContractAddress, big.NewInt(0), gasLimit, gasPrice, txData)
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(big.NewInt(11155111)), privateKey)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
		return "", err
	}

	err = evmContractClient.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
		return "", err
	}

	fmt.Printf("Transaction sent! TX Hash: %s\n", signedTx.Hash().Hex())
	return signedTx.Hash().Hex(), nil
}

func (evmContractClient *EvmContractClient) SendMintHarvestTransaction(userAddress string, indexerPrivateKeyStr string, ryeAmount uint64) (string, error) {
	privateKey, err := crypto.HexToECDSA(indexerPrivateKeyStr)
	if err != nil {
		log.Fatalf("Failed to convert private key: %v", err)
		return "", err
	}
	// Derive the address from the private key
	senderAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// Check if the derived address matches the expected sender address
	if senderAddress != common.HexToAddress(evmContractClient.IndexerAddress) {
		log.Fatalf("Address derived from private key does not match the expected sender address")
		return "", fmt.Errorf("invalid sender address")
	}

	contractAddress := common.HexToAddress(evmContractClient.HarvestContractAddress)
	txData, err := evmContractClient.HarvestContractAbi.Pack("mintHYDR", 1, contractAddress, big.NewInt(int64(ryeAmount)), userAddress)
	if err != nil {
		log.Fatalf("Failed to pack data for transaction: %v", err)
		return "", err
	}

	fromAddress := common.HexToAddress(evmContractClient.IndexerAddress)
	nonce, err := evmContractClient.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
		return "", err
	}

	gasLimit := uint64(200000)
	gasPrice, err := evmContractClient.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
		return "", err
	}
	tx := types.NewTransaction(nonce, contractAddress, big.NewInt(0), gasLimit, gasPrice, txData)

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(big.NewInt(11155111)), privateKey)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
		return "", err
	}

	err = evmContractClient.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
		return "", err
	}

	fmt.Printf("Transaction sent! TX Hash: %s\n", signedTx.Hash().Hex())
	return signedTx.Hash().Hex(), nil
}

func (evmContractClient *EvmContractClient) SendClaimHarvestTransaction(userAddress string, indexerPrivateKeyStr string, roundId uint64) (string, error) {
	privateKey, err := crypto.HexToECDSA(indexerPrivateKeyStr)
	if err != nil {
		log.Fatalf("Failed to convert private key: %v", err)
		return "", err
	}

	contractAddress := common.HexToAddress(evmContractClient.HarvestContractAddress)
	txData, err := evmContractClient.HarvestContractAbi.Pack("claimReward", roundId, userAddress)
	if err != nil {
		log.Fatalf("Failed to pack data for transaction: %v", err)
		return "", err
	}

	fromAddress := common.HexToAddress(evmContractClient.IndexerAddress)
	nonce, err := evmContractClient.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
		return "", err
	}

	gasLimit := uint64(200000)
	gasPrice, err := evmContractClient.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
		return "", err
	}
	tx := types.NewTransaction(nonce, contractAddress, big.NewInt(0), gasLimit, gasPrice, txData)

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(big.NewInt(11155111)), privateKey)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
		return "", err
	}

	err = evmContractClient.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
		return "", err
	}

	fmt.Printf("Transaction sent! TX Hash: %s\n", signedTx.Hash().Hex())
	return signedTx.Hash().Hex(), nil
}
