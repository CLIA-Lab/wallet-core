package utils

import (
	"crypto/rand"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestIsEven(t *testing.T) {
	for i := 0; i < 30; i++ {
		x, err := rand.Int(rand.Reader, new(big.Int).Exp(big.NewInt(2), big.NewInt(32), nil))
		if err != nil {
			panic(err)
		}
		xPadded := Pad32(x.Bytes())
		assert.Equal(t, x.Bit(0) == 0, IsEven(xPadded))
	}
}
