package bitcoin

import "github.com/CLIA-Lab/wallet-core/utils"

type PrivateKey struct {
	Bytes [32]byte
}

func NewPrivateKey() *PrivateKey {
	bytes := GeneratePrivateKeyBytesBigEndian()
	return &PrivateKey{Bytes: utils.First32Bytes(bytes)}
}
