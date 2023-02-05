package utils

import "encoding/hex"

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

func First32Bytes(bytes []byte) [32]byte {
	var ans [32]byte
	for i := 0; i < 32; i++ {
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
