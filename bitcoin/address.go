package bitcoin

import (
	"github.com/CLIA-Lab/wallet-core/utils"
)

type Address struct {
	CompressedBytes   [20]byte
	UncompressedBytes [20]byte
}

func GetAddress(publicKey *PublicKey) *Address {
	compressedPubKey := publicKey.ToCompressed()
	uncompressedPubKey := publicKey.ToUncompressed()

	return &Address{
		CompressedBytes:   getAddressHash(compressedPubKey[:]),
		UncompressedBytes: getAddressHash(uncompressedPubKey[:]),
	}
}

func getAddressHash(publicKey []byte) [20]byte {
	publicKeySha256 := utils.ApplySha256(publicKey)
	addressBytes := utils.ApplyRipemd160(publicKeySha256)
	return utils.First20Bytes(addressBytes)
}

func (address *Address) ToBase58CheckEncode() (compressed, uncompressed string) {
	return base58CheckEncodeAddressBytes(address.CompressedBytes[:]),
		base58CheckEncodeAddressBytes(address.UncompressedBytes[:])
}

func base58CheckEncodeAddressBytes(bytes []byte) string {
	addressPrefix := []byte{0x00}
	return utils.Base58CheckEncode(append(addressPrefix, bytes...))
}
