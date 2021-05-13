package tracker

import (
	"context"
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

func TestDisputeString(t *testing.T) {

	tracker := &DisputeTracker{}

	disputeTrackerStr := tracker.String()

	assert.Equal(t, disputeTrackerStr, "DisputeTracker", "should return 'DisputeTracker' string")
}

func TestDisputeStatus(t *testing.T) {

	startBal := big.NewInt(356000)

	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}, DisputeStatus: big.NewInt(1)}

	client := rpc.NewMockClientWithValues(opts)

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_disputeStatus"))

	testDisputePath := filepath.Join(os.TempDir(), "test_disputeStatus")

	tracker := &DisputeTracker{}

	ctx := context.WithValue(context.Background(), common.ClientContextKey, client)
	ctx = context.WithValue(ctx, common.DBContextKey, DB)
	trackerErr := tracker.Exec(ctx)

	getStatusKey, getStatusKeyErr := DB.Get(db.DisputeStatusKey)

	statusKey, statusKeyErr := hexutil.DecodeBig(string(getStatusKey))

	t.Logf("Dispute Status stored: %v\n", statusKey)

	testDisputeStatus := statusKey.Cmp(big.NewInt(1))

	assert.Nil(t, err, err)

	assert.NotNil(t, DB)

	assert.DirExists(t, testDisputePath)

	assert.Nil(t, trackerErr, trackerErr)

	assert.Nil(t, getStatusKeyErr)

	assert.NotNil(t, getStatusKey)

	assert.Nil(t, statusKeyErr, statusKeyErr)

	assert.NotNil(t, statusKey)

	assert.Equal(t, testDisputeStatus, 0,

		"Dispute Status from client did not match what should have been stored in DB. %s != %s", statusKey, "one")

	DB.Close()

}
