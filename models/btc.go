package models

type Status int

const (
	Pending Status = iota
	Processing
	Error
	Failed
	Success
)

type TransactionRequestStatusBase struct {
	BaseModel
	TxHash                        *string
	BlockNumber                   *uint64
	BlockTimestamp                *uint64
	Retries                       uint64
	ProgrammableTransactionStatus string // success, failed
	ProgrammableResult            string
	L1TransactionStatus           Status // pending, processing, success, failed.
	LastError                     *string
}

type BTCTransactionBase struct {
	Network     string `json:"network"`
	Sender      string `json:"sender"`
	Receiver    string `json:"receiver"`
	Amount      string `json:"amount"`
	Timestamp   uint64 `json:"timestamp"`
	BlockHeight uint64 `json:"block_height"`
	TxHash      string `json:"tx_hash" gorm:"uniqueIndex"`
	TxIndex     uint64 `json:"tx_index"`
}

type RyeHarvestTxRequestBase struct {
	ControllerAddress      string  `json:"controller_address"`
	ControllerPublicKey    string  `josn:"controller_public_key"`
	Nonce                  uint64  `json:"nonce"`
	ExpiredAt              uint64  `json:"expired_at"`
	RawTransferTransaction *string `json:"raw_transfer_transaction"` // should contain btc/rye token transfer
	RawCallData            *string `json:"raw_call_data"`
	TransferSig            *string `json:"transfer_sig"`
	CalldataSig            *string `json:"calldata_sig"`
}

type BTCNetworkInfo struct {
	BaseModel
	Network     string `json:"network" gorm:"uniqueIndex"`
	BlockNumber uint64 `json:"block_number"`
}

// user multisig address
type MultiSigAddress struct {
	BaseModel
	Network         string `json:"network"`
	MultisigAddress string `json:"multisig_address"`
	UserPublicKey   string `json:"user_public_key"`
}

// multisig btc funding tx
type MultisigBTCFunding struct {
	BaseModel
	BTCTransactionBase
	Status      bool   `json:"status"` // pending, confirmed, timeout
	ConfirmedAt string `json:"confirmed_at"`
}

// multisig runes funding tx
type MultisigRuneFunding struct {
	BaseModel
	BTCTransactionBase
	RuneTokenID     string `json:"rune_token_id"`
	RuneTokenAmount string `json:"rune_token_amount"`
	Status          bool   `json:"status"` // pending, confirmed, timeout
	ConfirmedAt     string `json:"confirmed_at"`
}
