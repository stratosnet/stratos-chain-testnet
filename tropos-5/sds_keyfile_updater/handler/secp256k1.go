package handler

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	"github.com/stratosnet/stratos-chain/crypto/hd"
)

func PrivKeyToSdkPrivKey(privKey []byte) cryptotypes.PrivKey {
	return hd.EthSecp256k1.Generate()(privKey)
}

func PrivKeyToAddress(privKey []byte) Address {
	privKeyObject := PrivKeyToSdkPrivKey(privKey)
	return BytesToAddress(privKeyObject.PubKey().Address())
}

// BytesToAddress returns Address with value b.
// If b is larger than len(h), b will be cropped from the left.
func BytesToAddress(b []byte) Address {
	var a Address
	a.SetBytes(b)
	return a
}
