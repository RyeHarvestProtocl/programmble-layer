package evmSigVerifier

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func VerifyEvmSingedMessage(message string, signatureHex string, addressHex string) (bool, error) {
	signatureHex = strings.TrimPrefix(signatureHex, "0x")

	// Convert the hex address to an Ethereum address type.
	address := common.HexToAddress(addressHex)

	// Hash the message to match the Ethereum signing format.
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(message))
	hash := crypto.Keccak256Hash([]byte(prefix + message))

	// Convert the hex signature to bytes.
	signatureBytes, err := hex.DecodeString(signatureHex)
	if err != nil {
		return false, fmt.Errorf("invalid signature format: %v", err)
	}

	// Ethereum signatures are [R || S || V] format, and V needs to be 0 or 1 for SigToPub
	// Correct the V value if necessary (Ethereum specific adjustment)
	if signatureBytes[64] >= 27 {
		signatureBytes[64] -= 27
	}

	// Recover the public key from the signature.
	publicKey, err := crypto.SigToPub(hash.Bytes(), signatureBytes)
	if err != nil {
		return false, fmt.Errorf("failed to recover public key: %v", err)
	}

	// Convert the recovered public key to an Ethereum address.
	recoveredAddr := crypto.PubkeyToAddress(*publicKey)

	// Check if the recovered address matches the given address.
	return address == recoveredAddr, nil
}
