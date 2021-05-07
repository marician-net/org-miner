package tracker

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/zapproject/zap-miner/common"
	"github.com/zapproject/zap-miner/db"
	"github.com/zapproject/zap-miner/rpc"
)

func TestStringId(t *testing.T) {

	tracker := &BalanceTracker{}

	// Converts the value of tracker to a string
	res := tracker.String()

	// If res does not equal "BalanceTracker" log
	if res != "BalanceTracker" {
		t.Fatalf("should return 'BalanceTracker' string")
	}

	// If res equals "BalanceTracker" print "BalanceTracker ID:", "BalanceTracker"
	if res == "BalanceTracker" {
		fmt.Println("BalanceTracker ID:", res)
	}

	assert.Equal(t, 2, 1)
}
func TestPositiveBalance(t *testing.T) {

	// Stores the bigInt balance 356000
	startBal := big.NewInt(356000)

	dbBalanceTest(startBal, t)
}

func TestZeroBalance(t *testing.T) {

	// Stores the bigInt balance 0
	startBal := big.NewInt(0)

	dbBalanceTest(startBal, t)
}

func TestNegativeBalance(t *testing.T) {

	// Stores the bigInt balance -753
	startBal := big.NewInt(-753)

	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}}
	client := rpc.NewMockClientWithValues(opts)

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_balance"))

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
}

func dbBalanceTest(startBal *big.Int, t *testing.T) {

	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}}
	client := rpc.NewMockClientWithValues(opts)

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_balance"))

	if err != nil {
		t.Fatal(err)
	}

	// Assigned the string value "BalanceTracker"
	tracker := &BalanceTracker{}

	ctx := context.WithValue(context.Background(), common.ClientContextKey, client)
	ctx = context.WithValue(ctx, common.DBContextKey, DB)
	err = tracker.Exec(ctx)

	if err != nil {
		t.Fatal(err)
	}

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
