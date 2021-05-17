package tracker

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zapproject/zap-miner/common"
	"github.com/zapproject/zap-miner/db"
)

func TestMeanAt(t *testing.T) {

	// Creates the test db for test_MeanAt
	db, dbErr := db.Open(filepath.Join(os.TempDir(), "test_MeanAt"))

	ctx := context.Background()
	ctx = context.WithValue(ctx, common.DBContextKey, db)

	// BuildIndexTrackers creates and initializes a new tracker instance
	BuildIndexTrackers()

	// Gets all the query strings related to ETH/USD pairs
	ethIndexes := indexes["ETH/USD"]

	// Tracks and validates each query string stored in ethIndexes
	execEthUsdPsrs(ctx, t, ethIndexes)

	// Gets the average price of ETH/USD the time the function is invoked
	MeanAt(ethIndexes, clck.Now())

	// Assert dbErr is not nil
	assert.Nil(t, dbErr)

	db.Close()
}
