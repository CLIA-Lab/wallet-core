package bitcoin

import (
	"github.com/CLIA-Lab/wallet-core/utils"
	"github.com/stretchr/testify/assert"
	"math/big"
	"strings"
	"testing"
)

func TestPrivateKeysDifferEachOther(t *testing.T) {
	const NumberOfKeys = 30
	privateKeys := [NumberOfKeys][]byte{}

	for i := 0; i < NumberOfKeys; i++ {
		privateKeys[i] = generatePrivateKeyBytesBigEndian()
		for j := 0; j < i; j++ {
			assert.True(t, utils.NotEqualBytes(privateKeys[i], privateKeys[j]))
		}
	}
}

func TestPrivateKeysAreLessThanCurveOrderMinus1(t *testing.T) {
	for i := 0; i < 30; i++ {
		privateKey := toBigIntFromBigEndian(generatePrivateKeyBytesBigEndian())
		assert.True(t, privateKey.Cmp(curveOrderMinusOne) < 0)
	}
}

func TestPrivateKeysAreGreaterThan0(t *testing.T) {
	for i := 0; i < 30; i++ {
		privateKey := toBigIntFromBigEndian(generatePrivateKeyBytesBigEndian())
		assert.True(t, privateKey.Cmp(big.NewInt(0)) > 0)
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

func TestMnemonicPhraseHas24Words(t *testing.T) {
	mnemonicPhrase := GenerateMnemonicPhrase()
	words := strings.Split(mnemonicPhrase, " ")
	assert.Len(t, words, 24)
}

func TestMnemonicPhraseIsDifferentEachTime(t *testing.T) {
	const NUMBER_OF_PHRASES = 30
	mnemonics := make([]string, NUMBER_OF_PHRASES)

	for i := 0; i < NUMBER_OF_PHRASES; i++ {
		mnemonics[i] = GenerateMnemonicPhrase()
		assertIsNotEqualToAllBefore(t, i, mnemonics)
	}
}

func assertIsNotEqualToAllBefore(t *testing.T, index int, mnemonics []string) {
	for j := 0; j < index; j++ {
		assert.NotEqual(t, mnemonics[index], mnemonics[j])
	}
}
