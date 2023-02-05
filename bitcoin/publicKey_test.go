package bitcoin

import (
	"github.com/CLIA-Lab/wallet-core/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type keysCase struct {
	privateKey [32]byte
	publicKeyX [32]byte
	publicKeyY [32]byte
}

func TestGetPublicKeyOfBigEndian(t *testing.T) {
	for testCase := range getTestCases() {
		actualX, actualY := GetPublicKeyOfBigEndian(testCase.privateKey)
		expectedX := testCase.publicKeyX
		expectedY := testCase.publicKeyY

		assert.True(t, utils.EqualBytes(expectedX[:], actualX[:]))
		assert.True(t, utils.EqualBytes(expectedY[:], actualY[:]))
	}
}

func getTestCases() <-chan *keysCase {
	channel := make(chan *keysCase)
	go func() {
		publicKeyOf := map[string][2]string{
			"905e986484cd97da5fc593d061e3610684147f7f1133d509b8334e13b052ded9": {
				"ca167032d15483f557426b662dd06c54511ea7616ed9e4020647d49644962e86",
				"7a1be4b4be273838766153a2bce2d7eb746be27591fd38cd145a7fe72ad3f425",
			},
			"86642cd13ead811301cbd9c1981b429b5c0148c697b1bd2bd34dd9a2aecc30c0": {
				"c09eec7e40578183aa73b8e186c2a8a37997ac04d3b25917ea95a4c89c7c1716",
				"070a027d124b70f677227e0fc90c9dd7828078bc77d660374de4e8a15f14365f",
			},
			"3dbc701825ed7b530297c68927544ec8cb894c0753b72be6a0cf2c67fe94f412": {
				"c5fce944e455bc53f83370b7c41ff93b2f8bc8cfdda24f19a537c273fdfb0b93",
				"3a524665b8b00a9263e97b3a5e5ccf35aa198b3556c8a6d078a5ecc73f1ef7af",
			},
			"40a77aa3a80462386a0d507ac3745b170c4f917c406318cf7ceb8d4abc23710a": {
				"c323ac364aa5097abf8d6f03561c9a81cf87a3c2b4bf42103288c33a4af583fc",
				"8fd3fc3e7adc1dd43f51e5a64410539e00a95e10bfd23b7e9789022185e42efa",
			},
			"dbbfbf0c2b3703c4a99770b303fc1561f0cf5cf434ca53f510d034986593e490": {
				"0a3852d6541478ecedfd2f606126178b2339811a31e9a3fc28017c86eb4f72e4",
				"4b8c44007d05bc201d5c6d2cc11f61c228ca924e411fdb11217466ca2957167c",
			},
		}
		for privateKeyString, publicKeyStrings := range publicKeyOf {
			testCase := &keysCase{
				privateKey: utils.First32Bytes(utils.BytesFromString(privateKeyString)),
				publicKeyX: utils.First32Bytes(utils.BytesFromString(publicKeyStrings[0])),
				publicKeyY: utils.First32Bytes(utils.BytesFromString(publicKeyStrings[1])),
			}
			channel <- testCase
		}
		close(channel)
	}()
	return channel
}

func TestPublicKeyCoordinatesHave32BytesEach(t *testing.T) {
	firstTestCase := <-getTestCases()
	privateKey := firstTestCase.privateKey
	x, y := GetPublicKeyOfBigEndian(privateKey)

	assert.Len(t, x, 32)
	assert.Len(t, y, 32)
}
