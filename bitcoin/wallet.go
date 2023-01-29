package bitcoin

import (
	"crypto/rand"
	"crypto/sha256"
	//"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"math/big"
)

var ellipticCurveOrder = getEllipticCurveOrder()

func generateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func generatePrivateKeyBytes() []byte { // @TODO include "BigEndian" in the name
	seed, err := generateRandomString(25)
	if err != nil {
		panic(err)
	}
	ellipticCurveOrder := getEllipticCurveOrder()
	for {
		h := sha256.New()
		h.Write([]byte(seed))
		privateKeyBytes := h.Sum(nil)
		privateKeyBigInt := toBigIntFromBigEndian(privateKeyBytes)

		if privateKeyBigInt.Cmp(big.NewInt(0)) > 0 && privateKeyBigInt.Cmp(ellipticCurveOrder) < 0 {
			return privateKeyBytes
		}
	}
}

func first32Bytes(bytes []byte) [32]byte {
	ans := [32]byte{}
	for i := 0; i < 32; i++ {
		ans[i] = bytes[i]
	}
	return ans
}

func getEllipticCurveOrder() *big.Int {
	ans := big.NewInt(1158)
	tenPow77 := new(big.Int)
	tenPow77.Exp(big.NewInt(10), big.NewInt(77), nil)
	return ans.Mul(ans, tenPow77)
}

func toBigIntFromBigEndian(bytes []byte) *big.Int {
	return new(big.Int).SetBytes(bytes)
}
