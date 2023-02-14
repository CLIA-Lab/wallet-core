package bitcoin

import (
	"github.com/CLIA-Lab/wallet-core/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type addressCase struct {
	privateKey                     *PrivateKey
	compressedBase58CheckEncoded   string
	uncompressedBase58CheckEncoded string
}

func TestUncompressedBase58CheckEncode(t *testing.T) {
	for testCase := range getAddressTestCases() {
		actualAddress := GetAddress(GetPublicKey(testCase.privateKey))
		_, actualBase58CheckEncoded := actualAddress.ToBase58CheckEncode()

		assert.Equal(t, testCase.uncompressedBase58CheckEncoded, actualBase58CheckEncoded)
	}
}

func getAddressTestCases() <-chan *addressCase {
	channel := make(chan *addressCase)
	go func() {
		addressOfPrivateKey := map[string][2]string{
			"687dbdcee9bc5db1792c28d1de6331bbec9692ebd24eb0e5a00df5e0f2fc1887": {
				"1KyzGyXri468tZc4jDRz37hE3WToS5ZNQN",
				"1EKP4nHLC98PF8Vr1Fq9x7K1NBTkVdFkjG",
			},
			"0e276029b7ad649ba8a41cc7b520ff9b45aacc41764a559e8aa8979fdcbce1d4": {
				"19xTY6MHHhJvpKUJxntdF1ZyK5tefxTPUm",
				"1L7XCWoWFmqgQF5Gu1NTkhZLdZkyvbxBra",
			},
			"d1e624af7a0822e1099eb303450ae9fec2decc03ac2f51e7da9ed88fed519b97": {
				"1HukqZBsbeQaxtx5YnvK9qgzzj22mbBwyu",
				"1DrwsfjrfjeXGnzS16EvWxC4Pb77PUkCkS",
			},
			"6a665c1f7386857663fd0e6c4c35b2e30c6992a24f40c9fc9f46779c8daecb48": {
				"13so7udfCsn7t5WqN5UtHj39j5MJC3Bogp",
				"14SUR7GvqErQMEkWKbzLQQUcuCKfo2Dbd7",
			},
			"cfc9e4b64a163176f2c4da924150bbc14dc97bf50046483eea4bdc9b994b26b0": {
				"18ZxXNGLoUdLjxirnHYHGkKxZtYaR3LCnd",
				"1Q2qGacFL3yeUhgcS4928T6LqdeYF5YvAv",
			},
		}
		for privateKeyString, addresses := range addressOfPrivateKey {
			privateKeyBytes := utils.First32Bytes(utils.BytesFromString(privateKeyString))
			channel <- &addressCase{
				privateKey:                     &PrivateKey{Bytes: privateKeyBytes},
				compressedBase58CheckEncoded:   addresses[0],
				uncompressedBase58CheckEncoded: addresses[1],
			}
		}
		close(channel)
	}()
	return channel
}

func TestCompressedBase58CheckEncode(t *testing.T) {
	for testCase := range getAddressTestCases() {
		actualAddress := GetAddress(GetPublicKey(testCase.privateKey))
		actualBase58CheckEncoded, _ := actualAddress.ToBase58CheckEncode()

		assert.Equal(t, testCase.compressedBase58CheckEncoded, actualBase58CheckEncoded)
	}
}
