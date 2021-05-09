package tracker

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"github.com/zapproject/zap-miner/common"
	"github.com/zapproject/zap-miner/db"
	"github.com/zapproject/zap-miner/rpc"
)

func TestStringId(t *testing.T) {

	// Stores the value "BalanceTracker" with the type *tracker.BalanceTracker
	tracker := &BalanceTracker{}

	// Converts the type of tracker to a string
	res := tracker.String()

	// If res does not equal "BalanceTracker" log
	if res != "BalanceTracker" {
		t.Fatalf("should return 'BalanceTracker' string")
	}

	// Assert that the value of res is equal to "BalanceTracker"
	assert.Equal(t, res, "BalanceTracker")

}
func TestPositiveBalance(t *testing.T) {

	// +1 = Greater than
	// 0 = Equal To
	// -1 = Less than

	// Stores the bigInt balance 356000
	startBal := big.NewInt(356000)

	// Stores the bigInt value 0
	zeroBal := big.NewInt(0)

	// Stores the bigInt comparison between startBal(356000) and zeroBal(0)
	// Checking if startBal is greater than, equal to, or less than zeroBal
	testPositiveBal := startBal.Cmp(zeroBal)

	// Assert that testPositiveBal equals +1(Greater than)
	assert.Equal(t, testPositiveBal, +1)

	dbBalanceTest(startBal, t)

}

func TestZeroBalance(t *testing.T) {

	//  +1 = Greater than
	//	0 = Equal To
	//  -1 = Less than

	// Stores the bigInt balance 0
	startBal := big.NewInt(0)

	// Stores the bigInt value 0
	zeroBal := big.NewInt(0)

	// Stores the bigInt comparison between startBal(0) and zeroBal(0)
	// Checking if startBal is greater than, equal to, or less than zeroBal
	testZeroBal := startBal.Cmp(zeroBal)

	// Assert that testZeroBal equals 0(Equal to)
	assert.Equal(t, testZeroBal, 0)

	dbBalanceTest(startBal, t)
}

func TestNegativeBalance(t *testing.T) {

	//  +1 = Greater than
	//	0 = Equal To
	//  -1 = Less than

	// Stores the bigInt balance -753
	startBal := big.NewInt(-753)

	// Stores the bigInt balance 0
	zeroBal := big.NewInt(0)

	// Config options for the mock configs
	opts := &rpc.MockOptions{
		ETHBalance:    startBal,
		Nonce:         1,
		GasPrice:      big.NewInt(700000000),
		TokenBalance:  big.NewInt(0),
		Top50Requests: []*big.Int{}}

	// NewMockClientWithValues creates a mock client with default values to return for calls
	client := rpc.NewMockClientWithValues(opts)

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_balance"))

	fmt.Println(DB)

	// Error Handler
	if err != nil {
		t.Fatal(err)
	}

	// Prints the BalanceTracker id "BalanceTracker"
	tracker := &BalanceTracker{}

	ctx := context.WithValue(context.Background(), common.ClientContextKey, client)
	ctx = context.WithValue(ctx, common.DBContextKey, DB)
	err = tracker.Exec(ctx)

	if err == nil {
		t.Fatal(err)

	}

	// Stores the bigInt comparison between startBal(0) and zeroBal(0)
	// Checking if startBal is greater than, equal to, or less than zeroBal
	testNegBal := startBal.Cmp(zeroBal)

	// Assert that testZeroBal equals -1(Less than)
	assert.Equal(t, testNegBal, -1)
}

func dbBalanceTest(startBal *big.Int, t *testing.T) {

	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}}
	client := rpc.NewMockClientWithValues(opts)

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_balance"))

	if err != nil {
		t.Fatal(err)
	}

	tracker := &BalanceTracker{}

	ctx := context.WithValue(context.Background(), common.ClientContextKey, client)
	ctx = context.WithValue(ctx, common.DBContextKey, DB)
	err = tracker.Exec(ctx)

	if err != nil {
		t.Fatal(err)
	}

	// This value is only read during the TestPositiveBalance and TestZeroBalance tests
	// If TestPositiveBalance, v = 35600
	// If TestNegativeBalance, v = 0
	v, err := DB.Get(db.BalanceKey)

	if err != nil {

		t.Fatal(err)
	}

	b, err := hexutil.DecodeBig(string(v))

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Balance stored: %v\n", string(v))

	if b.Cmp(startBal) != 0 {
		t.Fatalf("Balance from client did not match what should have been stored in DB. %s != %s", b, startBal)
	}
	DB.Close()
}
