package bitcoin

import "github.com/CLIA-Lab/wallet-core/utils"

type PrivateKey struct {
	Bytes [32]byte
}

func NewPrivateKey() *PrivateKey {
	bytes := GeneratePrivateKeyBytesBigEndian()
	return &PrivateKey{Bytes: utils.First32Bytes(bytes)}
}

func NewPrivateKeyFromHex(hex string) *PrivateKey {
	bytes := utils.BytesFromString(hex)
	bytes32 := utils.First32Bytes(bytes)
	return &PrivateKey{Bytes: bytes32}
}

func (privateKey *PrivateKey) ToWif() string {
	wifPrefix := []byte{0x80}
	return utils.Base58CheckEncode(append(wifPrefix, privateKey.Bytes[:]...))
}
