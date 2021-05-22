package rpc

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/stretchr/testify/assert"

	"github.com/zapproject/zap-miner/config"
	"github.com/zapproject/zap-miner/util"
)

func setup() {

	// Parses config.json
	err := config.ParseConfig("../config.json")

	// Checks if err is not nil
	// nil = No value
	if err != nil {
		fmt.Println("Can't parse config for test.")
	}

	// Stores the path
	path := "../testConfig.json"

	// Should be nil
	err = util.ParseLoggingConfig(path)

	// Error handler for the config
	if err != nil {
		fmt.Println("Can't parse logging config for test.")
	}
}

func TestABICodec(t *testing.T) {

	setup()

	codec, codecErr := BuildCodec()

	// Method Signature 0xe1eee6d6
	getRequestVarsSig := codec.methods["0xe1eee6d6"]

	// string, string, bytes32,  uint, uint, uint
	// Declare a bytes array with the length of 32
	var hash [32]byte

	copy([]byte("12345"), hash[:])

	// data - Returns "someQueryString" & "ETH/USD" as a bytes array
	data, dataErr := getRequestVarsSig.Outputs.Pack(
		"SomeQueryString",
		"ETH/USD",
		hash,
		big.NewInt(1000),
		big.NewInt(0),
		big.NewInt(0))

	for i := 0; i < len(data); i += 32 {

		// Encodes the bytes array to hex strings
		encode := hexutil.Encode(data[i : i+32])

		// Decodes the hex strings with a 0x prefix
		decode, decodeErr := hexutil.Decode(encode)

		fmt.Println(string(decode))

		// Asserts decodeErr has no value
		assert.Nil(t, decodeErr)
	}

	// Asserts codecErr has no value
	assert.Nil(t, codecErr, codecErr)

	// Asserts getRequestVarsSig has a value
	assert.NotNil(t, getRequestVarsSig, "Missing expected method matching test sig")

	// Asserts getRequestVarsSig.Name is equal to "getRequestVars"
	assert.Equal(t, getRequestVarsSig.Name,

		"getRequestVars", "Method name is unexpected. %s != getRequestVars", getRequestVarsSig.Name)

	// Asserts dataErr has no value
	assert.Nil(t, dataErr, dataErr)

}
