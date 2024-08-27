package attestor

type BTCBaseTransaction struct {
	Network   string
	Timestamp string
	TxHash    string
	BlockHash string
}

type BTCFundingTransaction struct {
	BTCBaseTransaction
	Sender   string
	Receiver string
	Amount   string // in sat
}

type BTCExitTransactionSig struct {
	Sender          string
	Receiver        string
	Amount          string
	SenderSignature string
}

type RuneExitTransactionSig struct {
	Sender           string
	Receiver         string
	RuneId           string
	Amount           string
	SenderSignature  string
	IndexerSignature string
}
