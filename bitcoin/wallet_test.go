package bitcoin

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestPrivateKeyHas256Bits(t *testing.T) {
	privateKeyBytes := GeneratePrivateKeyBytesBigEndian()
	assert.Len(t, privateKeyBytes, 32)
}

func TestPrivateKeysDifferEachOther(t *testing.T) {
	const NumberOfKeys = 30
	privateKeys := [NumberOfKeys][]byte{}

	for i := 0; i < NumberOfKeys; i++ {
		privateKeys[i] = GeneratePrivateKeyBytesBigEndian()
		for j := 0; j < i; j++ {
			assert.True(t, notEqualBytes(privateKeys[i], privateKeys[j]))
		}
	}
}

func notEqualBytes(expected []byte, actual []byte) bool {
	for idx, expectedByte := range expected {
		if expectedByte != actual[idx] {
			return true
		}
	}
	return false
}

func TestPrivateKeysAreLessThanCurveOrderMinus1(t *testing.T) {
	for i := 0; i < 30; i++ {
		privateKey := toBigIntFromBigEndian(GeneratePrivateKeyBytesBigEndian())
		assert.True(t, privateKey.Cmp(curveOrderMinusOne) < 0)
	}
}

func TestPrivateKeysAreGreaterThan1(t *testing.T) {
	for i := 0; i < 30; i++ {
		privateKey := toBigIntFromBigEndian(GeneratePrivateKeyBytesBigEndian())
		assert.True(t, privateKey.Cmp(big.NewInt(1)) > 0)
	}
}

func TestToBigInt(t *testing.T) {
	numbersToBigEndianBytes := map[uint64][]byte{
		25:      {25},
		78:      {78},
		133434:  {2, 9, 58},
		1600000: {24, 106, 0},
		4:       {4},
		0:       {0},
		1:       {1},
	}
	for x, xInBytes := range numbersToBigEndianBytes {
		assertBytesIsBigEndianOfNumber(t, xInBytes, x)
	}
}

func assertBytesIsBigEndianOfNumber(t *testing.T, bytes []byte, x uint64) {
	assert.Equal(t, int64(x), toBigIntFromBigEndian(bytes).Int64())
}
