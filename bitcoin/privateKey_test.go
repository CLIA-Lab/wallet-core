package bitcoin

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type privateKeyCase struct {
	privateKey  *PrivateKey
	expectedWif string
}

func ExampleNewPrivateKey() {
	privateKey := NewPrivateKey()
	fmt.Println(len(privateKey.Bytes))
	// Output:
	// 32
}

func ExamplePrivateKey_ToHex() {
	privateKey := NewPrivateKeyFromHex("905e986484cd97da5fc593d061e3610684147f7f1133d509b8334e13b052ded9")
	fmt.Println(privateKey.ToHex())
	// Output:
	// 905e986484cd97da5fc593d061e3610684147f7f1133d509b8334e13b052ded9
}

func ExamplePrivateKey_ToWif() {
	privateKey := NewPrivateKeyFromHex("905e986484cd97da5fc593d061e3610684147f7f1133d509b8334e13b052ded9")
	fmt.Println(privateKey.ToWif())
	// Output:
	// 5JusJkqHvemE1KWEoCZP7tp2DdZgUXVoDdkPvPk1VqCNEi586pi
}

func TestHas256Bits(t *testing.T) {
	privateKey := NewPrivateKey()
	assert.Len(t, privateKey.Bytes, 32)
}

func TestToWif(t *testing.T) {
	for testCase := range getPrivateKeyTestCases() {
		assert.Equal(
			t,
			testCase.expectedWif,
			testCase.privateKey.ToWif(),
		)
	}
}

func getPrivateKeyTestCases() <-chan *privateKeyCase {
	channel := make(chan *privateKeyCase)
	go func() {
		wifOf := map[string]string{
			"0bc297b6eb9528d5d5f4e098fe3338ea39add0ac047df29be807ce0128de8bbc": "5HuTygrhJ9GXuw5JNAWZmccpz7cV8fjEoNZnRABSuYvcjCx2drq",
			"2e7452fcef19d8da95ecb54d5088d1443f14c1aef2312c25fde0ef50da5bd222": "5JAkCMa3kG4PvCYftWTDeVYNu4dPaX6wu596uq6VYk3WHZ1iXgZ",
			"18f142b3f3d546802b44a4f97e8e643c66537506e4fa089771752071ce9ebbbb": "5J1Ghge8Qb1QHYbvrr7X7Bycr7Q6nhPgsFAXFKybYmbeJuAC6oA",
			"5bb46bddde0ab9c6c7e3ccfc16a21e715891b79296af62ed7207d02265a79e5e": "5JWg4FLJKvh6XUshauskc4ScAxWecPETcxKdsz2Pvvw972joDJ1",
			"8fd210d6124b1fd8a680c998b01282f528722c101c92ad8aaa693270a56f75f0": "5JudHUkKnWxcPKfbo3DB6uRo4Bby5wmaShSvPoV6ooPr6SRqyw9",
		}
		for privateKeyHex, wif := range wifOf {
			channel <- &privateKeyCase{
				privateKey:  NewPrivateKeyFromHex(privateKeyHex),
				expectedWif: wif,
			}
		}
		close(channel)
	}()
	return channel
}
