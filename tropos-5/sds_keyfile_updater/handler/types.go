package handler

import (
	"github.com/pborman/uuid"
)

const (
	AddressLength = 20
)

type EncryptedKeyJSONV3 struct {
	Address string     `json:"address"`
	Name    string     `json:"name"`
	Crypto  cryptoJSON `json:"crypto"`
	Id      string     `json:"id"`
	Version int        `json:"version"`
}

type cryptoJSON struct {
	Cipher       string                 `json:"cipher"`
	CipherText   string                 `json:"ciphertext"`
	CipherParams cipherparamsJSON       `json:"cipherparams"`
	KDF          string                 `json:"kdf"`
	KDFParams    map[string]interface{} `json:"kdfparams"`
	MAC          string                 `json:"mac"`
}

type cipherparamsJSON struct {
	IV string `json:"iv"`
}

type hdKeyBytes struct {
	HdPath     []byte
	Mnemonic   []byte
	Passphrase []byte
	PrivKey    []byte
}

type AccountKey struct {
	Id uuid.UUID // Version 4 "random" for unique id not derived from key data
	// to simplify lookups we also store the address
	Address Address
	// The HD path to use to derive this key
	HdPath string
	// The mnemonic for the underlying HD wallet
	Mnemonic string
	// a nickname
	Name string
	// The bip39 passphrase for the underlying HD wallet
	Passphrase string
	// we only store privkey as pubkey/address can be derived from it
	// privkey in this struct is always in plaintext
	PrivateKey []byte
}

type Address [AddressLength]byte

// SetBytes sets the address to the value of b.
// If b is larger than len(a) it will panic.
func (a *Address) SetBytes(b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-AddressLength:]
	}
	copy(a[AddressLength-len(b):], b)
}

func (a Address) Bytes() []byte { return a[:] }
