package btcservice

import (
	"fmt"
	"testing"

	"github.com/RyeHarvestProtocol/programmable-layer/testutils"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/keyfuse/tokucore/network"
	"github.com/keyfuse/tokucore/xcore"
	"github.com/keyfuse/tokucore/xcore/bip32"
)

func TestMultisig(t *testing.T) {
	net := network.TestNet
	networkParam := &chaincfg.TestNet3Params
	seed1 := []byte("doctor envelope garment program lunar manage zone oppose illegal muscle nice blind")
	hdkey1 := bip32.NewHDKey(seed1)
	path1 := fmt.Sprintf("m/86'/0'/0'/%d/%d", 0, 0)
	prvkey1, err := hdkey1.DeriveByPath(path1)
	testutils.AssertNil(err)
	pvstr1 := prvkey1.PrivateKey()
	testutils.AssertNil(err)
	address1 := xcore.NewPayToPubKeyHashAddress(pvstr1.PubKey().Hash160())
	t.Log("hd wallet address 1: ", address1.ToString(net))
	pubStr1 := fmt.Sprintf("%x", pvstr1.PubKey().SerializeCompressed()) // Convert to hex string
	t.Log("public key hex string 1: ", pubStr1)

	seed2 := []byte("salute kingdom share skirt measure net first inherit odor such person taxi")
	hdkey2 := bip32.NewHDKey(seed2)
	path2 := fmt.Sprintf("m/86'/0'/0'/%d/%d", 0, 0)
	prvkey2, err := hdkey2.DeriveByPath(path2)
	testutils.AssertNil(err)
	pvstr2 := prvkey2.PrivateKey()
	testutils.AssertNil(err)
	address2 := xcore.NewPayToPubKeyHashAddress(pvstr2.PubKey().Hash160())
	t.Log("hd wallet address 2: ", address2.ToString(net))
	pubStr2 := fmt.Sprintf("%x", pvstr2.PubKey().SerializeCompressed()) // Convert to hex string
	t.Log("public key hex string 2: ", pubStr2)
	// generate a 2 out of 2 p2wsh multisig address
	multisigAddress, err := GetUserMultiSigP2WSHAddress(pubStr1, pubStr2, networkParam)
	testutils.AssertNil(err)
	t.Log("2 out of 2 multi address: ", multisigAddress)

	// create a tx to spend the uxto of the multisigAddress

	// user 1 sign the tx

	// user2 sign the tx

	// verify with 2 signatures we can spend the utxo
}
