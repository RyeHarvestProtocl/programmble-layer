package btcSigVerifier

import (
	"testing"
)

func TestGetTaprootAddrFromPublicKey(t *testing.T) {
	// publicKey, err := keypair.NewECPPublicFromHex("03c99fc32c7b476d95d57a3bc6e38e7363f65278dd8bf98ba96b1d7dacd416d9e7")
	// if err != nil {
	// 	t.Error("error when create public key fro hex: ", publicKey)
	// }

	// p2tr := publicKey.ToTaprootAddress().Show()
	// t.Log("==p2tr: ", p2tr)
	// assert.Equal(t, p2tr, "bc1pcccgun7lmwrec5q2ajhchnt4jyc50w9v4yml7zyr95ddjxr2wr6qtdkwqr")

	success, err := VerifyBTCMessage("RyeHarvest: Bind to UserId: 2", "IB+XC/4+K5ewCWj06AgfDVpRbrAU+g0sNt5JDjAc4ZG3IyZAtW5B94rcpwK7WlGrAqXpzcCTq4sko9HbUicXqFM=", "tb1p7wpf7te2wdqpnnn6jdudfmu8722l25mvzadhk3ny5l9uxp8vsucsceydsr", "")
	t.Log("==success: ", success)
	if err != nil {
		t.Errorf("failed to verify: %s", err)
	}
}
