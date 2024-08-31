package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

var TwitterOAuthConfig *oauth2.Config

type Postgresql struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Url      string `mapstructure:"url"`
}

func (p *Postgresql) String() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		p.User, p.Password, p.Host, p.Port, p.Database,
	)
}

type BTCRpcClient struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	CookiePath   string `mapstructure:"cookiepath"`
	Pass         string `mapstructure:"pass"`
	Network      string `mapstructure:"network"`
	StartBlock   uint64 `mapstructure:"startblock"`
	BatchSize    uint64 `mapstructure:"batchsize"`
	Publickey    string `mapstructure:"publickey"`
	PriceApi     string `mapstructure:"priceapi"`
	MemPoolApi   string `mapstructure:"mempoolapi"`
	QuickNodeApi string `mapstructure:"quicknodeapi"`
	TheGraphApi  string `mapstructure:"thegraphapi"`
	Enabled      bool   `mapstructure:"enabled"`
}

type RuneClient struct {
	OrdUrl          string `mapstructure:"ordurl"`
	Network         string `mapstructure:"network"`
	RyeRuneId       string `mapstructure:"ryeRuneId"`
	RyeIssuerPrvKey string `mapstructure:"ryeIssuerPrvKey"`
	RyeIssuerAddr   string `mapstructure:"ryeIssuerAddr"`
	TaskInterval    int64  `mapstructure:"taskInterval"` //in minutes
}

type EvmClient struct {
	RpcUrl                 string             `mapstructure:"rpcUrl"`
	NetworkId              uint64             `mapstructure:"network"`
	StartBlock             int64              `mapstructure:"startblock"`
	BatchSize              int64              `mapstructure:"batchsize"`
	InvoiceAddress         string             `mapstructure:"invoiceAddress"`
	ContractAddresses      *ContractAddresses `mapstructure:"contractAddresses"`
	Enabled                bool               `mapstructure:"enabled"`
	RyeContractAddress     string             `mapstructure:"ryeContractAddress"`
	HarvestContractAddress string             `mapstructure:"harvestContractAddress"`
	IndexerPrivateKey      string             `mapstructure:"indexerPrivateKey"`
	IndexerAddress         string             `mapstructure:"indexerAddress"`
}

type ContractAddresses struct {
	UsdcAddress string `mapstructure:"usdcAddress"`
	UsdtAddress string `mapstructure:"usdtAddress"`
}

type Config struct {
	Test                       string        `mapstructure:"test"`
	AppId                      string        `mapstructure:"appId"`
	Mode                       string        `mapstructure:"mode"`
	Bust                       int64         `mapstructure:"bust"`
	Env                        string        `mapstructure:"env"`
	Postgresql                 *Postgresql   `mapstructure:"postgresql"`
	BTCRpcClient               *BTCRpcClient `mapstructure:"btcRpcClient"`
	RuneClient                 *RuneClient   `mapstructure:"runeClient"`
	EvmClient                  *EvmClient    `mapstructure:"evmClient"`
	BTCInvoiceAddress          string        `mapstructure:"btcInvoiceAddress"`
	BTCAttestorPublicKey       string        `mapstructure:"btcAttestorAddress"`
	EvmDepositContractAddress  string        `mapstructure:"evmDepositContractAddress"`
	JwtKey                     string        `mapstructure:"jwtKey"`
	Port                       string        `mapstructure:"port"`
	TwitterApiKey              string        `mapstructure:"twitterApiKey"`
	TwitterApiKeySecret        string        `mapstructure:"twitterApiKeySecret"`
	TwitterCallBackUrl         string        `mapstructure:"twitterCallbackUrl"`
	Twitter135Key              string        `mapstructure:"twitter135Key"`
	TwitterOauth2Config        *oauth2.Config
	TwitterOfficialUserId      string `mapstructure:"twitterOfficialUserId"`
	TwitterOfficialTweetId     string `mapstructure:"twitterOfficialTweetId"`
	ProcessRyeOrderLoopEnabled bool   `mapstructure:"processRyeOrderLoopEnabled"`
	UpdateBTCPriceLoopEnabled  bool   `mapstructure:"updateBTCPriceLoopEnabled"`
}

func (c *Config) Valid() error {
	if c.Mode == "" {
		return errors.New("mode is empty")
	}

	// check port
	if c.Port == "" {
		return errors.New("port is not set")
	}

	return nil
}

func findGitRoot() (string, error) {
	// Start from the current directory.
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Walk up the directory tree until reaching the root.
	for {
		// Check if the .gitignore file exists in this directory.
		if _, err := os.Stat(filepath.Join(currentDir, ".gitignore")); err == nil {
			return currentDir, nil
		}

		// Move to the parent directory.
		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			// If the parent directory is the same as the current directory, we've reached the root.
			break
		}
		currentDir = parentDir
	}

	return "", errors.New("no .gitignore found in any parent directory")
}

// Singleton instance
var instance *Config
var once sync.Once

// New creates and returns a new Config instance
func New(mode string, absConfigPath string) *Config {
	once.Do(func() {
		if mode == "" {
			mode = os.Getenv("Mode")
			if mode == "" {
				mode = "local"
			}
		}

		instance = &Config{}
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")

		var configPath string
		baseDir, err := findGitRoot()

		if err != nil {
			fmt.Printf("Failed to find Git root: %v", err)
			os.Exit(1)
		}

		if absConfigPath == "" {
			switch mode {
			case "test":
				configPath = "config/test/"
			case "pre":
				configPath = "config/pre/"
			case "pro":
				configPath = "config/pro/"
			default:
				configPath = "config/local/"
			}
		} else {
			configPath = absConfigPath
		}

		fmt.Println("✅✅✅✅✅configPath: ", configPath)

		viper.AddConfigPath(filepath.Join(baseDir, configPath))
		viper.AddConfigPath(filepath.Join("/app/", configPath))

		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Error reading config file: %s\n", err)
			os.Exit(1)
		}

		viper.SetEnvPrefix("CONFIG_ENV")
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AutomaticEnv() // todo: check why this is not working

		viper.BindEnv("evmClient.indexerPrivateKey", "CONFIG_ENV_INDEXER_PRIVATE_KEY")
		viper.BindEnv("evmClient.indexerAddress", "CONFIG_ENV_INDEXER_ADDRESS")
		viper.BindEnv("evmClient.rpcUrl", "CONFIG_ENV_RPC")

		if err := viper.Unmarshal(&instance); err != nil {
			fmt.Printf("Error unmarshalling config: %s\n", err)
			os.Exit(1)
		}

		if err := instance.Valid(); err != nil {
			fmt.Printf("Validation error in config: %s\n", err)
			os.Exit(1)
		}

	})
	return instance
}
