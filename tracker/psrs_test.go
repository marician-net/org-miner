package tracker

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zapproject/zap-miner/common"
	"github.com/zapproject/zap-miner/db"
)

func TestGetLatest(t *testing.T) {

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

	// Gets thprice/volumes for ETH/USD
	getLatest, test := getLatest(ethIndexes, clck.Now())

	latestLen := len(getLatest)

	ethIndexesLen := len(ethIndexes)

	fmt.Println(test)

	// Asserts dbErr has no value
	assert.Nil(t, dbErr)

	// Asserts the array length equal the amount of ETH/USD query strings
	assert.Equal(t, latestLen, ethIndexesLen)
}

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

	// Gets the average price & average volume of ETH/USD the time the function is invoked
	getMeanAt, confidence := MeanAt(ethIndexes, clck.Now())

	// Assert dbErr is not nil
	assert.Nil(t, dbErr)

	// Asserts the mean price is positive
	assert.Positive(t, getMeanAt.Price)

	// Asserts the mean volume is negative
	assert.Positive(t, getMeanAt.Volume)

	// Assert confidence is not 0
	// 0 = No Value
	assert.NotEqual(t, confidence, 0, "Confidence:", confidence, "has no value")

	db.Close()
}
