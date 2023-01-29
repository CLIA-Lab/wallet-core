package bitcoin

import (
	"crypto/rand"
	"crypto/sha256"
	"github.com/tyler-smith/go-bip39"
	"math/big"
)

var ellipticCurveOrder = getEllipticCurveOrder()
var curveOrderMinusOne = new(big.Int).Sub(ellipticCurveOrder, big.NewInt(1))

func getEllipticCurveOrder() *big.Int {
	_1158 := big.NewInt(1158)
	_10Exp77 := new(big.Int).Exp(big.NewInt(10), big.NewInt(77), nil)
	return new(big.Int).Mul(_1158, _10Exp77)
}

func GeneratePrivateKeyBytesBigEndian() []byte {
	for {
		baseString := getRandomStringOfLength(25)
		privateKeyBytes := getSha256(baseString)
		privateKeyBigInt := toBigIntFromBigEndian(privateKeyBytes)

		if isInRange(privateKeyBigInt) {
			return privateKeyBytes
		}
	}
}

func getRandomStringOfLength(length int) string {
	randomBytes := getRandomBytes(length)
	return string(randomBytes)
}

func getRandomBytes(amount int) []byte {
	const characters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	randomBytes := make([]byte, amount)

	for i := 0; i < amount; i++ {
		randomBytes[i] = getRandomCharacter(characters)
	}
	return randomBytes
}

func getRandomCharacter(characters string) byte {
	randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
	if err != nil {
		panic(err)
	}
	return characters[randomIndex.Int64()]
}

func getSha256(baseString string) []byte {
	h := sha256.New()
	h.Write([]byte(baseString))
	privateKeyBytes := h.Sum(nil)
	return privateKeyBytes
}

func toBigIntFromBigEndian(bytes []byte) *big.Int {
	return new(big.Int).SetBytes(bytes)
}

func isInRange(privateKeyBigInt *big.Int) bool {
	return isLessThanCurveOrderMinusOne(privateKeyBigInt) &&
		isGreaterThan1(privateKeyBigInt)
}

func isLessThanCurveOrderMinusOne(privateKeyBigInt *big.Int) bool {
	return privateKeyBigInt.Cmp(curveOrderMinusOne) < 0
}

func isGreaterThan1(privateKeyBigInt *big.Int) bool {
	return privateKeyBigInt.Cmp(big.NewInt(1)) > 0
}

func GenerateMnemonicPhrase() string {
	entropy := getRandomBytes(32)
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		panic(err)
	}
	return mnemonic
}
