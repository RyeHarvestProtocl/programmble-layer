package btcservice

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
)

func GetUserMultiSigP2WSHAddress(userPubKeyStr string, indexerPubKeyStr string, networkParam *chaincfg.Params) (string, error) {
	pk1, err := hex.DecodeString(indexerPubKeyStr) // user's public key
	if err != nil {
		return "", err
	}

	pk2, err := hex.DecodeString(userPubKeyStr) // user's public key
	if err != nil {
		return "", err
	}

	// Create redeem script for 2 of 2 multi-sig
	builder := txscript.NewScriptBuilder()
	// Add the minimum number of needed signatures
	builder.AddOp(txscript.OP_2)
	// Add the 3 public keys
	builder.AddData(pk1).AddData(pk2)
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

	addr, err := btcutil.NewAddressWitnessScriptHash(witnessScriptHash[:], networkParam)
	if err != nil {
		return "", err
	}

	return addr.EncodeAddress(), nil
}
