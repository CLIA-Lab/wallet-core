package bitcoin

import (
	"encoding/hex"
	"github.com/CLIA-Lab/wallet-core/utils"
)

type PrivateKey struct {
	Bytes [32]byte
}

func NewPrivateKey() *PrivateKey {
	bytes := generatePrivateKeyBytesBigEndian()
	return &PrivateKey{Bytes: utils.First32Bytes(bytes)}
}

func NewPrivateKeyFromHex(hex string) *PrivateKey {
	if len(hex) != 64 {
		panic("Length of hexadecimal string must equal 64.")
	}
	bytes := utils.BytesFromString(hex)
	bytes32 := utils.First32Bytes(bytes)
	return &PrivateKey{Bytes: bytes32}
}

func NewPrivateKeyFromBip38Encrypted(bip38, passphrase string) *PrivateKey {
	return decryptBip38(bip38, passphrase)
}

func (privateKey *PrivateKey) ToWif() string {
	wifPrefix := []byte{0x80}
	return utils.Base58CheckEncode(append(wifPrefix, privateKey.Bytes[:]...))
}

func (privateKey *PrivateKey) ToHex() string {
	return hex.EncodeToString(privateKey.Bytes[:])
}

func (privateKey *PrivateKey) ToBip38Encrypted(passphrase string) string {
	return encryptBip38(privateKey, passphrase)
}
