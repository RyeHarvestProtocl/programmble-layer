package btcSigVerifier

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/mrtnetwork/bitcoin/address"
	"github.com/mrtnetwork/bitcoin/keypair"
)

func VerifyBTCMessage(message, signatureBase64 string, userAddress string, network string) (bool, error) {
	// Decode the base64-encoded compact signature.
	sigBytes, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		fmt.Println("err when Decoding signature: ", err)
		return false, fmt.Errorf("error decoding signature: %v", err)
	}

	// Compute the hash of the message with the magic prefix.
	messageHash := magicHash(message)

	// Extract the public key from the signature and message hash.
	recoveredPubKey, _, err := ecdsa.RecoverCompact(sigBytes, messageHash)
	if err != nil {
		fmt.Println("err when Extracting the public key : ", err)
		return false, fmt.Errorf("error recovering public key from signature: %v", err)
	}

	compressedPubKey := recoveredPubKey.SerializeCompressed()
	compressedHex := hex.EncodeToString(compressedPubKey)

	// compute address from public key
	publicKey, err := keypair.NewECPPublicFromHex(compressedHex)
	if err != nil {
		fmt.Println("error computing address from public key : ", err)
		return false, fmt.Errorf("error computing address from public key: %v", err)
	}

	var p2trAddr string
	if network == "testnet" {
		p2trAddr = publicKey.ToTaprootAddress().Show(address.TestnetNetwork)
	} else {
		p2trAddr = publicKey.ToTaprootAddress().Show(address.MainnetNetwork)
	}
	// fmt.Println("this is p2trAddr: ", p2trAddr)

	// Verify that the address matches the provided one.
	if userAddress == p2trAddr {
		return true, nil
	} else {
		fmt.Printf("public key does not match the one recovered from the signature. userAddress: %s, recoveredAdress: %s \n", userAddress, p2trAddr)
		return false, fmt.Errorf("public key does not match the one recovered from the signature")
	}
}

func PublicKeyToTaprootAddress(publicKeyStr string, network string) (string, error) {
	// compute address from public key
	publicKey, err := keypair.NewECPPublicFromHex(publicKeyStr)
	if err != nil {
		fmt.Println("error computing address from public key : ", err)
		return "", fmt.Errorf("error computing address from public key: %v", err)
	}

	var p2trAddr string
	if network == "testnet" {
		p2trAddr = publicKey.ToTaprootAddress().Show(address.TestnetNetwork)
	} else {
		p2trAddr = publicKey.ToTaprootAddress().Show(address.MainnetNetwork)
	}

	return p2trAddr, nil
}

func PublicKeyToSegwitAddress(publicKeyStr string, network string) (string, error) {
	// compute address from public key
	publicKey, err := keypair.NewECPPublicFromHex(publicKeyStr)
	if err != nil {
		fmt.Println("error computing address from public key : ", err)
		return "", fmt.Errorf("error computing address from public key: %v", err)
	}

	var p2trAddr string
	if network == "testnet" {
		p2trAddr = publicKey.ToSegwitAddress().Show(address.TestnetNetwork)
	} else {
		p2trAddr = publicKey.ToSegwitAddress().Show(address.MainnetNetwork)
	}

	return p2trAddr, nil
}

// varint encodes an integer as a variable-length quantity. For simplicity, this implementation
// covers only the cases needed for the example: numbers less than 253.
func varint(n int) []byte {
	if n < 0xfd {
		return []byte{byte(n)}
	}
	// Implementations for larger numbers can follow Bitcoin's varint encoding rules,
	// but are omitted for brevity since message lengths are unlikely to exceed this.
	panic("varint for number >= 253 not implemented")
}

// doubleSHA256 performs SHA256(SHA256(data)) and returns the result.
func doubleSHA256(data []byte) []byte {
	first := sha256.Sum256(data)
	second := sha256.Sum256(first[:])
	return second[:]
}

// magicHash creates a hash of the message, prefixed with "Bitcoin Signed Message:\n" and length prefixes,
// using double SHA256, mimicking Bitcoin's message signing format.
func magicHash(message string) []byte {
	magicPrefix := "Bitcoin Signed Message:\n"
	prefix1 := varint(len(magicPrefix))
	messageBuffer := []byte(message)
	prefix2 := varint(len(messageBuffer))
	var buf bytes.Buffer
	buf.Write(prefix1)
	buf.WriteString(magicPrefix)
	buf.Write(prefix2)
	buf.Write(messageBuffer)
	hash := doubleSHA256(buf.Bytes())
	return hash
}

// computeMagicHash creates a double SHA256 hash of the message, prefixed by the Bitcoin magic prefix.
func computeMagicHash(message string) []byte {
	magicPrefix := "Bitcoin Signed Message:\n"
	var buf bytes.Buffer
	buf.WriteByte(byte(len(magicPrefix)))
	buf.WriteString(magicPrefix)
	buf.WriteByte(byte(len(message)))
	buf.WriteString(message)
	hash := chainhash.DoubleHashB(buf.Bytes())
	return hash
}
