package testutils

import (
	"github.com/RyeHarvestProtocol/programmable-layer/config"
)

func GetBTCConfig(network string) config.BTCRpcClient {
	if network == "testnet" {
		return config.BTCRpcClient{
			Host:        "http://localhost:18332",
			CookiePath:  "/Users/bcfh/Library/Application Support/Bitcoin/testnet3/.cookie",
			Publickey:   "03164e3563fe204cc8841b1e2b8ee4282ed9b8b14c032fc55fe8ee0fbc4471af7c",
			PriceApi:    "http://api.coinlayer.com/live?access_key=63d4572858e20a38c525db4e0569f1f7&target=usd",
			MemPoolApi:  "https://mempool.space/testnet/api/",
			TheGraphApi: "https://gateway-arbitrum.network.thegraph.com/api/772fd867acf86489cbb35877c04c5764/subgraphs/id/HUZDsRpEVP2AvzDCyzDHtdc64dyDxx8FQjzsmqSg4H3B",
		}
	} else {
		return config.BTCRpcClient{
			// Host:       "http://localhost:8332",
			// CookiePath: "/Volumes/MYDISK/btcnode/.cookie",
			Host:        "https://nd-948-043-849.p2pify.com",
			User:        "laughing-lichterman",
			Pass:        "storm-cabbie-stream-grunge-handed-stash",
			Publickey:   "03164e3563fe204cc8841b1e2b8ee4282ed9b8b14c032fc55fe8ee0fbc4471af7c",
			PriceApi:    "http://api.coinlayer.com/live?access_key=63d4572858e20a38c525db4e0569f1f7&target=usd",
			MemPoolApi:  "https://mempool.space/testnet/api/",
			TheGraphApi: "https://gateway-arbitrum.network.thegraph.com/api/772fd867acf86489cbb35877c04c5764/subgraphs/id/HUZDsRpEVP2AvzDCyzDHtdc64dyDxx8FQjzsmqSg4H3B",
		}
	}
}
