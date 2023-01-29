package bitcoin

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestPrivateKeyHas256Bits(t *testing.T) {
	privateKeyBytes := generatePrivateKeyBytes()
	assert.Len(t, privateKeyBytes, 32)
}

func TestPrivateKeysDiffer(t *testing.T) {
	privateKeys := [30][]byte{}
	for i := 0; i < 30; i++ {
		privateKeys[i] = generatePrivateKeyBytes()
		for j := 0; j < i; j++ {
			assert.NotEqual(t, privateKeys[i], privateKeys[j])
		}
	}
}

func TestPrivateKeysAreLessThanCurveOrder(t *testing.T) {
	for i := 0; i < 30; i++ {
		privateKey := toBigIntFromBigEndian(generatePrivateKeyBytes())
		assert.True(t, privateKey.Cmp(ellipticCurveOrder) < 0)
	}
}

func TestPrivateKeysAreGreaterThanZero(t *testing.T) {
	for i := 0; i < 30; i++ {
		privateKey := toBigIntFromBigEndian(generatePrivateKeyBytes())
		assert.True(t, privateKey.Cmp(big.NewInt(0)) > 0)
	}
}

func TestToBigInt(t *testing.T) {
	numbers := []uint64{25, 78, 133434, 1600000, 4, 0, 1}
	inBinary := [][]byte{[]byte{25}, []byte{78}, []byte{2, 9, 58}, []byte{24, 106, 0}, []byte{4}, []byte{0}, []byte{1}}
	for i, x := range numbers {
		xInBytes := inBinary[i]
		fmt.Println(xInBytes)
		assert.Equal(t, int64(x), toBigIntFromBigEndian(xInBytes).Int64())
	}
}

func TestFirst32Bytes(t *testing.T) {
	bs1 := make([]byte, 32)
	bs1[0] = 1
	bs1[1] = 2
	bs1[2] = 3
	expectedBs1 := [32]byte{}
	expectedBs1[0] = 1
	expectedBs1[1] = 2
	expectedBs1[2] = 3
	assert.Equal(t, expectedBs1, first32Bytes(bs1))
}
