package attestor

import (
	"github.com/RyeHarvestProtocol/programmable-layer/config"
	"github.com/btcsuite/btcd/chaincfg"
)

type Attestor struct {
	Network             string
	NetworkParam        *chaincfg.Params
	AdminPublicKey      string // public key that's used for p2wsh address
	RegisteredContracts *RegisteredContracts
}

type RegisteredContracts struct {
	RyeMainContract        string // $rye mint & burn
	RyeStakingContract     string
	HarvestStakingContract string
}

func NewAttestor(config *config.Config) *Attestor {
	var networkParams *chaincfg.Params
	if config.BTCRpcClient.Network == "testnet" {
		networkParams = &chaincfg.TestNet3Params
	} else {
		networkParams = &chaincfg.MainNetParams
	}

	return &Attestor{
		Network:        config.BTCRpcClient.Network,
		NetworkParam:   networkParams,
		AdminPublicKey: config.BTCInvoiceAddress,
	}
}

// funding tx handler
// func FundingBtcTxHandler()

// exit tx

// stake tx

// transfer tx
