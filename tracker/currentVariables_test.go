package tracker

import (
	"context"
	"fmt"
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

	if res != "CurrentVariablesTracker" {
		t.Fatalf("should return 'CurrentVariablesTracker' string")
	}

	// Assert that the value of res is equal to "CurrentVariablesTracker"
	assert.Equal(t, res, "CurrentVariablesTracker")

}

func TestCurrentVariables(t *testing.T) {

	startBal := big.NewInt(356000)

	// Creates a hash
	hash := math.PaddedBigBytes(big.NewInt(256), 32)

	// Creates an array with the size of 32 bytes
	var b32 [32]byte

	// Iterates through hash
	for i, v := range hash {
		b32[i] = v

	}

	// Query String
	queryStr := "json(https://coinbase.com)"

	chal := &rpc.CurrentChallenge{ChallengeHash: b32, RequestID: big.NewInt(1),

		// Sets the difficulty to 500
		Difficulty: big.NewInt(500),

		// Sets the QueryString to "json(https://coinbase.com)"
		QueryString: queryStr,

		Granularity: big.NewInt(1000),
		Tip:         big.NewInt(0)}

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

	fmt.Println("Working to Line 41")

	err = tracker.Exec(ctx)

	if err != nil {
		t.Fatal(err)
	}

	// Gets the request ID as a bytes array
	v, err := DB.Get(db.RequestIdKey)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Working to Line 51", v)

	b, err := hexutil.DecodeBig(string(v))

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("RequestID stored: %v\n", string(v))

	if b.Cmp(big.NewInt(1)) != 0 {
		t.Fatalf("Current Request ID from client did not match what should have been stored in DB. %s != %s", b, string(rune(1)))
	}

	// Gets the QueryStringKey from the DB as a bytes array
	v, err = DB.Get(db.QueryStringKey)

	if err != nil {
		t.Fatal(err)
	}

	// Converts the QueryStringKey from a bytes array to a string
	// Checks if the QueryStringKey returned from the DB is equivalent to the queryStr
	if string(v) != queryStr {
		t.Fatalf("Expected query string to match test input: %s != %s\n", string(v), queryStr)
	}

}
