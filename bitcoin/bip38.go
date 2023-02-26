package bitcoin

import (
	"crypto/aes"
	"github.com/CLIA-Lab/wallet-core/utils"
	"golang.org/x/crypto/scrypt"
)

func encryptBip38(privateKey *PrivateKey, passphrase string) string {
	address := GetAddress(GetPublicKey(privateKey))
	_, addressBase58Check := address.ToBase58CheckEncode()
	addressHash := utils.ApplySha256(utils.ApplySha256([]byte(addressBase58Check)))[:4]
	derivedKey, err := scrypt.Key([]byte(passphrase), addressHash, 16384, 8, 8, 64)
	if err != nil {
		panic(err)
	}

	flag := byte(0xC0)

	derivedHalf1 := utils.First32Bytes(derivedKey[:32])
	derivedHalf2 := utils.First32Bytes(derivedKey[32:])
	data := encryptBip38Bytes(privateKey.Bytes, derivedHalf1, derivedHalf2)

	var allBytes [39]byte
	allBytes[0] = 0x01
	allBytes[1] = 0x42
	allBytes[2] = flag
	copy(allBytes[3:], addressHash)
	copy(allBytes[7:], data[:])

	return utils.Base58CheckEncode(allBytes[:])
}

func encryptBip38Bytes(privateKey, derivedHalf1, derivedHalf2 [32]byte) [32]byte {
	c, err := aes.NewCipher(derivedHalf2[:])
	if err != nil {
		panic(err)
	}
	for i := range derivedHalf1 {
		derivedHalf1[i] ^= privateKey[i]
	}
	dst := make([]byte, 48)
	c.Encrypt(dst, derivedHalf1[:16])
	c.Encrypt(dst[16:], derivedHalf1[16:])
	return utils.First32Bytes(dst[:32])
}

func decryptBip38(bip38, passphrase string) *PrivateKey {
	bip38Bytes := utils.Base58CheckDecode(bip38)
	addressHash := bip38Bytes[3:7]
	data := utils.First32Bytes(bip38Bytes[7:])
	derivedKey, err := scrypt.Key([]byte(passphrase), addressHash, 16384, 8, 8, 64)
	if err != nil {
		panic(err)
	}
	privateKey := new(PrivateKey)
	derivedHalf1 := utils.First32Bytes(derivedKey[:32])
	derivedHalf2 := utils.First32Bytes(derivedKey[32:])
	privateKey.Bytes = decryptBip38Bytes(data, derivedHalf1, derivedHalf2)
	return privateKey
}

func decryptBip38Bytes(src, derivedHalf1, derivedHalf2 [32]byte) [32]byte {
	c, err := aes.NewCipher(derivedHalf2[:])
	if err != nil {
		panic(err)
	}
	dst := make([]byte, 48)
	c.Decrypt(dst, src[:16])
	c.Decrypt(dst[16:], src[16:])
	dst = dst[:32]

	for i := range dst {
		dst[i] ^= derivedHalf1[i]
	}
	return utils.First32Bytes(dst)
}
