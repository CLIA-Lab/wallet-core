package bitcoin

import (
	"bytes"
	"encoding/hex"
	"github.com/CLIA-Lab/wallet-core/utils"
	"math/big"
)

type PublicKey struct {
	X [32]byte
	Y [32]byte
}

func GetPublicKey(privateKey *PrivateKey) *PublicKey {
	x, y := getPublicKeyPointOfBigEndian(privateKey.Bytes)
	return &PublicKey{X: x, Y: y}
}

var secp256k1 EllipticCurve

func init() {
	secp256k1.P, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
	secp256k1.A, _ = new(big.Int).SetString("0000000000000000000000000000000000000000000000000000000000000000", 16)
	secp256k1.B, _ = new(big.Int).SetString("0000000000000000000000000000000000000000000000000000000000000007", 16)
	secp256k1.G.X, _ = new(big.Int).SetString("79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798", 16)
	secp256k1.G.Y, _ = new(big.Int).SetString("483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8", 16)
	secp256k1.N, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16)
	secp256k1.H, _ = new(big.Int).SetString("01", 16)
}

func getPublicKeyPointOfBigEndian(privateKeyBigEndian [32]byte) (x [32]byte, y [32]byte) {
	k := new(big.Int).SetBytes(privateKeyBigEndian[:])
	Q := secp256k1.ScalarBaseMult(k)
	Q.X.FillBytes(x[:])
	Q.Y.FillBytes(y[:])
	return
}

func (pubKey *PublicKey) ToUncompressedHex() string {
	uncompressed := pubKey.ToUncompressed()
	return hex.EncodeToString(uncompressed[:])
}

func (pubKey *PublicKey) ToUncompressed() [65]byte {
	x := pubKey.X[:]
	y := pubKey.Y[:]

	/* Pad X and Y coordinate bytes to 32-bytes */
	padded_x := append(bytes.Repeat([]byte{0x00}, 32-len(x)), x...)
	padded_y := append(bytes.Repeat([]byte{0x00}, 32-len(y)), y...)

	/* Add prefix 0x04 for uncompressed coordinates */
	return utils.First65Bytes(append([]byte{0x04}, append(padded_x, padded_y...)...))
}

func (pubKey *PublicKey) ToCompressedHex() string {
	compressed := pubKey.ToCompressed()
	return hex.EncodeToString(compressed[:])
}

func (pubKey *PublicKey) ToCompressed() [33]byte {
	x := pubKey.X[:]

	padded_x := append(bytes.Repeat([]byte{0x00}, 32-len(x)), x...)

	var compressedBytes []byte
	if utils.IsEven(pubKey.Y) {
		compressedBytes = append([]byte{0x02}, padded_x...)
	} else {
		compressedBytes = append([]byte{0x03}, padded_x...)
	}
	return utils.First33Bytes(compressedBytes)
}
