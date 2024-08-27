package contracts

import "github.com/RyeHarvestProtocol/programmable-layer/models"

type TransferRuneRequest struct {
	models.BaseModel
	models.RyeHarvestTxRequestBase
	RuneId     string        `json:"rune_id"`
	RuneAmount string        `json:"rune_amount"`
	Status     models.Status `json:"status"`
	TxHash     *string       `json:"tx_hash"`
	Retries    uint64        `json:"retries"`
	LastError  string        `json:"last_error"`
}

// claim rewards
type ClaimHarvestRequest struct {
	models.BaseModel
	models.RyeHarvestTxRequestBase
	RewardReceiver string `json:"reward_receiver"`
}

// mint harvest
type MintHarvestRequest struct {
	models.BaseModel
	RyeAmount string `json:"rye_amount"`
}
