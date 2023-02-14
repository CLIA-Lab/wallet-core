package utils

import (
	bytes2 "bytes"
	"encoding/hex"
)

func NotEqualBytes(expected []byte, actual []byte) bool {
	for idx, expectedByte := range expected {
		if expectedByte != actual[idx] {
			return true
		}
	}
	return false
}

func EqualBytes(expected []byte, actual []byte) bool {
	return !NotEqualBytes(expected, actual)
}

func First20Bytes(bytes []byte) [20]byte {
	const AMOUNT = 20
	var ans [AMOUNT]byte
	for i := 0; i < AMOUNT; i++ {
		ans[i] = bytes[i]
	}
	return ans
}

func First32Bytes(bytes []byte) [32]byte {
	const AMOUNT = 32
	var ans [AMOUNT]byte
	for i := 0; i < AMOUNT; i++ {
		ans[i] = bytes[i]
	}
	return ans
}

func First33Bytes(bytes []byte) [33]byte {
	const AMOUNT = 33
	var ans [AMOUNT]byte
	for i := 0; i < AMOUNT; i++ {
		ans[i] = bytes[i]
	}
	return ans
}

func First65Bytes(bytes []byte) [65]byte {
	const AMOUNT = 65
	var ans [AMOUNT]byte
	for i := 0; i < AMOUNT; i++ {
		ans[i] = bytes[i]
	}
	return ans
}

func BytesFromString(s string) []byte {
	ans, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return ans
}

func IsEven(bigEndianBytes [32]byte) bool {
	return bigEndianBytes[31]%2 == 0
}

func Pad32(bytes []byte) [32]byte {
	zeroes := bytes2.Repeat([]byte{0}, 32-len(bytes))
	return First32Bytes(append(zeroes, bytes...))
}
