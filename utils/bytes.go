package utils

import (
	bytes2 "bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/ripemd160"
	"math/big"
	"strings"
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

func ApplySha256(bytes []byte) []byte {
	sha256_h := sha256.New()
	sha256_h.Reset()
	sha256_h.Write(bytes)
	return sha256_h.Sum(nil)
}

func Base58CheckEncode(bytes []byte) string {
	hashedTwice := ApplySha256(ApplySha256(bytes))

	checkSum := hashedTwice[:4]
	bytesWithCheckSum := append(bytes, checkSum...)

	base58String := base58Encode(bytesWithCheckSum)
	return prependOneForEachLeadingZero(base58String, bytesWithCheckSum)
}

func base58Encode(b []byte) (s string) {
	/* See https://en.bitcoin.it/wiki/Base58Check_encoding */

	const BITCOIN_BASE58_TABLE = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	/* Convert big endian bytes to big int */
	x := new(big.Int).SetBytes(b)

	/* Initialize */
	r := new(big.Int)
	m := big.NewInt(58)
	zero := big.NewInt(0)
	s = ""

	/* Convert big int to string */
	for x.Cmp(zero) > 0 {
		/* x, r = (x / 58, x % 58) */
		x.QuoRem(x, m, r)
		/* Prepend ASCII character */
		s = string(BITCOIN_BASE58_TABLE[r.Int64()]) + s
	}

	return s
}

func prependOneForEachLeadingZero(target string, bytes []byte) string {
	for _, v := range bytes {
		if v != 0 {
			break
		}
		target = "1" + target
	}
	return target
}

func Base58CheckDecode(s string) []byte {
	bytesWithCheckSum := base58Decode(s)
	/* Add leading zero bytes */
	for i := 0; i < len(s); i++ {
		if s[i] != '1' {
			break
		}
		bytesWithCheckSum = append([]byte{0x00}, bytesWithCheckSum...)
	}
	if len(bytesWithCheckSum) < 5 {
		panic("Invalid base-58 check string: missing checksum.")
	}
	bytesWithoutCheckSum := bytesWithCheckSum[:len(bytesWithCheckSum)-4]
	expectedCheckSum := ApplySha256(ApplySha256(bytesWithoutCheckSum))[:4]
	actualCheckSum := bytesWithCheckSum[len(bytesWithCheckSum)-4:]

	if bytes2.Compare(expectedCheckSum, actualCheckSum) != 0 {
		panic("Invalid base-58 check string: invalid checksum.")
	}
	return bytesWithoutCheckSum
}

func base58Decode(s string) []byte {
	const BITCOIN_BASE58_TABLE = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	/* Initialize */
	x := big.NewInt(0)
	m := big.NewInt(58)

	/* Convert string to big int */
	for i := 0; i < len(s); i++ {
		b58index := strings.IndexByte(BITCOIN_BASE58_TABLE, s[i])
		if b58index == -1 {
			panic(fmt.Sprintf("Invalid base-58 character encountered: '%c', index %d.", s[i], i))
		}
		b58value := big.NewInt(int64(b58index))
		x.Mul(x, m)
		x.Add(x, b58value)
	}

	/* Convert big int to big endian bytes */
	return x.Bytes()
}

func ApplyRipemd160(bytes []byte) []byte {
	ripemd160_h := ripemd160.New()
	ripemd160_h.Reset()
	ripemd160_h.Write(bytes)
	return ripemd160_h.Sum(nil)
}
