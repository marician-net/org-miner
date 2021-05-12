package tracker

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/stretchr/testify/assert"
	"github.com/zapproject/zap-miner/common"

	"github.com/zapproject/zap-miner/db"
	"github.com/zapproject/zap-miner/rpc"
)

func TestDisputeString(t *testing.T) {

	// Gets the DisputeTracker ID
	tracker := &DisputeTracker{}

	// Converts the DisputeTracker ID to a string
	disputeTrackerStr := tracker.String()

	// Asserts disputeTrackerStr equals "DispuateTracker"
	assert.Equal(t, disputeTrackerStr, "DisputeTracker", "should return 'DisputeTracker' string")
}

func TestDisputeStatus(t *testing.T) {

	startBal := big.NewInt(356000)

	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}, DisputeStatus: big.NewInt(1)}

	client := rpc.NewMockClientWithValues(opts)

	fmt.Println(reflect.TypeOf(client))

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_disputeStatus"))

	// Path for test_disputeStatus directory
	testDisputePath := filepath.Join(os.TempDir(), "test_disputeStatus")

	tracker := &DisputeTracker{}

	ctx := context.WithValue(context.Background(), common.ClientContextKey, client)
	ctx = context.WithValue(ctx, common.DBContextKey, DB)
	trackerErr := tracker.Exec(ctx)

	getStatusKey, getStatusKeyErr := DB.Get(db.DisputeStatusKey)

	statusKey, statusKeyErr := hexutil.DecodeBig(string(getStatusKey))

	t.Logf("Dispute Status stored: %v\n", statusKey)

	testDisputeStatus := statusKey.Cmp(big.NewInt(1))

	// Asserts err has no value
	assert.Nil(t, err, err)

	// Asserts DB has a value
	assert.NotNil(t, DB)

	// Asserts testDisputePath exists
	assert.DirExists(t, testDisputePath)

	// Asserts trackerErr has no value
	assert.Nil(t, trackerErr, trackerErr)

	assert.Nil(t, getStatusKeyErr)

	assert.NotNil(t, getStatusKey)

	assert.Nil(t, statusKeyErr, statusKeyErr)

	assert.NotNil(t, statusKey)

	assert.Equal(t, testDisputeStatus, 0,

		"Dispute Status from client did not match what should have been stored in DB. %s != %s", statusKey, "one")

	DB.Close()

}

func TestDisputeStatusNegativeBalance(t *testing.T) {

	startBal := big.NewInt(356000)

	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}, DisputeStatus: big.NewInt(0)}
	client := rpc.NewMockClientWithValues(opts)

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_disputeStatus"))

	if err != nil {
		t.Fatal(err)
	}

	tracker := &DisputeTracker{}
	ctx := context.WithValue(context.Background(), common.ClientContextKey, client)
	ctx = context.WithValue(ctx, common.DBContextKey, DB)

	err = tracker.Exec(ctx)
	if err != nil {
		t.Fatal(err)
	}

	v, err := DB.Get(db.DisputeStatusKey)
	if err != nil {
		t.Fatal(err)
	}
	b, err := hexutil.DecodeBig(string(v))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Dispute Status stored: %v\n", string(v))
	if b.Cmp(big.NewInt(1)) != 0 {
		t.Fatalf("Dispute Status from client did not match what should have been stored in DB. %s != %s", b, "one")
	}
}
