package bitcoin

import (
	"crypto/sha256"
	"github.com/CLIA-Lab/wallet-core/utils"
	"golang.org/x/crypto/ripemd160"
	"math/big"
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
	publicKeySha256 := applySha256(publicKey)
	addressBytes := applyRipemd160(publicKeySha256)
	return utils.First20Bytes(addressBytes)
}

func applySha256(bytes []byte) []byte {
	sha256_h := sha256.New()
	sha256_h.Reset()
	sha256_h.Write(bytes)
	return sha256_h.Sum(nil)
}

func applyRipemd160(bytes []byte) []byte {
	ripemd160_h := ripemd160.New()
	ripemd160_h.Reset()
	ripemd160_h.Write(bytes)
	return ripemd160_h.Sum(nil)
}

func (address *Address) ToBase58CheckEncode() (compressed, uncompressed string) {
	return toBase58CheckEncode(address.CompressedBytes[:]),
		toBase58CheckEncode(address.UncompressedBytes[:])
}

func toBase58CheckEncode(bytes []byte) string {
	versionByte := []byte{0x00}
	bytesWithVersion := append(versionByte, bytes...)

	hashedTwice := applySha256(applySha256(bytesWithVersion))

	checkSum := hashedTwice[0:4]
	bytesWithVersion = append(bytesWithVersion, checkSum...)

	/* Encode base58 string */
	base58String := base58Encode(bytesWithVersion)
	return prependOneForEachLeadingZero(base58String, bytesWithVersion)
}

func base58Encode(b []byte) (s string) {
	/* See https://en.bitcoin.it/wiki/Base58Check_encoding */

	const BITCOIN_BASE58_TABLE = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	/* Convert big endian bytes to big int */
	x := new(big.Int).SetBytes(b)

	/* Initialize */
	r := new(big.Int)
	m := big.NewInt(58)
	zero := big.NewInt(0)
	s = ""

	/* Convert big int to string */
	for x.Cmp(zero) > 0 {
		/* x, r = (x / 58, x % 58) */
		x.QuoRem(x, m, r)
		/* Prepend ASCII character */
		s = string(BITCOIN_BASE58_TABLE[r.Int64()]) + s
	}

	return s
}

func prependOneForEachLeadingZero(target string, bytes []byte) string {
	for _, v := range bytes {
		if v != 0 {
			break
		}
		target = "1" + target
	}
	return target
}
