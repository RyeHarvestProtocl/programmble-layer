package services

import (
	"fmt"

	"github.com/RyeHarvestProtocol/programmable-layer/config"
	"github.com/RyeHarvestProtocol/programmable-layer/models"
	"github.com/btcsuite/btcd/chaincfg"
	"gorm.io/gorm"
)

type BTCIndexerService struct {
	Network          string
	NetworkParam     *chaincfg.Params
	BTCBlockInfoRepo *BaseService[models.BTCNetworkInfo]
}

func NewBTCIndexerService(db *gorm.DB, config *config.Config) *BTCIndexerService {
	var networkParams *chaincfg.Params
	if config.BTCRpcClient.Network == "testnet" {
		networkParams = &chaincfg.TestNet3Params
	} else {
		networkParams = &chaincfg.MainNetParams
	}

	return &BTCIndexerService{
		Network:          config.BTCRpcClient.Network,
		NetworkParam:     networkParams,
		BTCBlockInfoRepo: NewBaseService[models.BTCNetworkInfo](db),
	}
}

func (service *BTCIndexerService) UpdateBlockHeight(tx *gorm.DB, blockHeight uint64, network string) error {
	fmt.Printf("updating network %s 's blockHeight to %d \n", network, blockHeight)

	result := tx.Model(&models.BTCNetworkInfo{}).Where("network = ?", network).Update("block_number", blockHeight)
	if result.Error != nil {
		return fmt.Errorf("failed to update block height: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected, check if the network is correct: %s", network)
	}
	return nil
}
