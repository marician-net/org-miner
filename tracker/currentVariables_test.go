package tracker

import (
	"context"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	zapCommon "github.com/zapproject/zap-miner/common"
	"github.com/zapproject/zap-miner/config"
	zap "github.com/zapproject/zap-miner/contracts"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/zapproject/zap-miner/db"
	"github.com/zapproject/zap-miner/rpc"
)

func TestCurrentVarableString(t *testing.T) {

	// Gets the CurrentVariablesTracker string
	// Type is *tracker.CurrentVariablesTracker
	tracker := &CurrentVariablesTracker{}

	// Converts the tracker data type to a string
	res := tracker.String()

	// Assert that the value of res is equal to "CurrentVariablesTracker"
	assert.Equal(t, res, "CurrentVariablesTracker", "Should return 'CurrentVariablesTracker' string")

}

func TestCurrentVariables(t *testing.T) {

	// Stores the bigInt balance 356000
	startBal := big.NewInt(356000)

	// Creates a hash
	hash := math.PaddedBigBytes(big.NewInt(256), 32)

	// Creates an array with the size of 32 bytes
	var b32 [32]byte

	// Iterates through hash
	for i, hashElement := range hash {

		b32[i] = hashElement

	}

	// Query String
	queryStr := "json(https://coinbase.com)"

	chal := &rpc.CurrentChallenge{ChallengeHash: b32,

		RequestID: big.NewInt(1),

		Difficulty: big.NewInt(500),

		QueryString: queryStr,

		Granularity: big.NewInt(1000),

		Tip: big.NewInt(0)}

	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1,

		GasPrice: big.NewInt(700000000),

		TokenBalance: big.NewInt(0),

		MiningStatus: true,

		Top50Requests: []*big.Int{},

		CurrentChallenge: chal}

	client := rpc.NewMockClientWithValues(opts)

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_currentVariables"))

	if err != nil {
		t.Fatal(err)
	}

	cfg := config.GetConfig()

	tracker := &CurrentVariablesTracker{}

	ctx := context.WithValue(context.Background(), zapCommon.ClientContextKey, client)

	ctx = context.WithValue(ctx, zapCommon.DBContextKey, DB)

	masterInstance := ctx.Value(zapCommon.MasterContractContextKey)

	if masterInstance == nil {

		contractAddress := common.HexToAddress(cfg.ContractAddress)

		// TODO create error state flag for mock client
		masterInstance, err = zap.NewZapMaster(contractAddress, client)

		if err != nil {
			runnerLog.Error("Problem creating zap master instance: %v\n", err)
			return
		}

		ctx = context.WithValue(ctx, zapCommon.MasterContractContextKey, masterInstance)
	}

	err = tracker.Exec(ctx)

	// Error handler for tracker.Exec
	if err != nil {
		t.Fatal(err)
	}

	// Gets the request ID as a bytes array
	getRequestId, err := DB.Get(db.RequestIdKey)

	// Error handler for getRequestId
	if err != nil {
		t.Fatal(err)
	}

	// Parses the requestId from a bytes array to a bigInt
	requestId, err := hexutil.DecodeBig(string(getRequestId))

	// Error handler for requestId
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("RequestID stored: %v\n", string(getRequestId))

	// If the requestId does not equal 1 log the error
	if requestId.Cmp(big.NewInt(1)) != 0 {
		t.Fatalf(
			"Current Request ID from client did not match what should have been stored in DB. %s != %s",
			requestId, string(rune(1)))
	}

	// Stores the bigInt comparison between requestId and a bigInt with the value of 1
	// Checking if requestId is greater than, equal to, or less than bigInt 1
	testQueryId := requestId.Cmp(big.NewInt(1))

	// Gets the QueryStringKey from the DB as a bytes array
	var getQueryStr, getQueryErr = DB.Get(db.QueryStringKey)

	// Error Handler for getQueryStr
	if getQueryErr != nil {
		t.Fatal(err)
	}

	// Converts getQueryStr from a bytes array to a string
	// Checks if its not equal to queryStr
	if string(getQueryStr) != queryStr {

		t.Fatalf("Expected query string to match test input: %s != %s\n",
			string(getQueryStr), queryStr)
	}

	// Asserts that the b32 challenge array length is 32
	assert.Len(t, b32, 32)

	// Assert that tracker is equal to "CurrentVariablesTracker"
	assert.Equal(t, tracker.String(), "CurrentVariablesTracker")

	// Assert that testQueryId equals 0(Equal to)
	assert.Equal(t, testQueryId, 0)

	// Asserts that returned QueryStringKey is equal to the queryString
	assert.Equal(t, string(getQueryStr), queryStr)

}
