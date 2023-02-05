package bitcoin

import "math/big"

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

func GetPublicKeyOfBigEndian(privateKeyBigEndian [32]byte) (x [32]byte, y [32]byte) {
	k := new(big.Int).SetBytes(privateKeyBigEndian[:])
	Q := secp256k1.ScalarBaseMult(k)
	Q.X.FillBytes(x[:])
	Q.Y.FillBytes(y[:])
	return
}
