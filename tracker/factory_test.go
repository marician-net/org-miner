package tracker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBalanceTracker(t *testing.T) {

	// Gets the BalanceTracker string
	// Type is *tracker.BalanceTracker
	balanceTracker, _ := createTracker("balance")

	// Converts the tracker data type to a string
	balanceTrackerStr := balanceTracker[0].String()

	// Asserts balanceTrackerStr has a value
	assert.NotNil(t, balanceTrackerStr)

	// Asserts balanceTrackerStr equals "BalanceTracker"
	assert.Equal(t, balanceTrackerStr, "BalanceTracker",

		"Expected BalanceTracker but got %s", balanceTrackerStr)

}

func TestCreateCurrentVariablesTracker(t *testing.T) {

	// Gets the CurrentVariablesTracker string
	// Type is *tracker.CurrentVariablesTracker
	currentVariablesTracker, _ := createTracker("currentVariables")

	// Converts the tracker data type to a string
	currentVariablesStr := currentVariablesTracker[0].String()

	// Asserts currentVariablesStr has a value
	assert.NotNil(t, currentVariablesStr)

	// Asserts currentVariablesStr is equal to "CurrentVariablesTracker"
	assert.Equal(t, currentVariablesStr, "CurrentVariablesTracker",

		"Expected CurrentVariablesTracker but got %s", currentVariablesStr)

}

func TestCreateDisputeStatusTracker(t *testing.T) {

	// Gets the DisputeStatusTracker string
	// Type is *tracker.DisputeTracker
	disputeStatusTracker, _ := createTracker("disputeStatus")

	// Converts the tracker data type to a string
	disputeStatusStr := disputeStatusTracker[0].String()

	// Asserts disputeStatusStr has a value
	assert.NotNil(t, disputeStatusStr)

	// Asserts disputeStatusStr is equal to DisputeTracker
	assert.Equal(t, disputeStatusStr, "DisputeTracker",

		"Expected DisputeTracker but got %s", disputeStatusStr)

}

func TestCreateGasTracker(t *testing.T) {

	gasTracker, _ := createTracker("gas")

	gasTrackerStr := gasTracker[0].String()

	if gasTrackerStr != "GasTracker" {
		t.Fatalf("Expected GasTracker but got %s", gasTrackerStr)
	}

	assert.Equal(t, gasTrackerStr, "GasTracker")

}

func TestCreateTokenBalanceTracker(t *testing.T) {

	tokenBalanceTracker, _ := createTracker("tokenBalance")

	tokenBalanceStr := tokenBalanceTracker[0].String()

	if tokenBalanceStr != "TokenTracker" {
		t.Fatalf("Expected TokenTracker but got %s", tokenBalanceStr)
	}

	assert.Equal(t, tokenBalanceStr, "TokenTracker")

}

func TestCreateIndexesTracker(t *testing.T) {

	indexersTracker, err := createTracker("indexers")

	if err != nil {
		t.Fatalf("Could not build IndexTracker")
	}

	if len(indexersTracker) == 0 {
		t.Fatalf("Could not build all IndexTrackers: only tracking %d indexes", len(indexersTracker))
	}

	assert.NotEqual(t, len(indexersTracker), 0)

}

func TestCreateDisputeCheckerTracker(t *testing.T) {

	disputeChecker, _ := createTracker("disputeChecker")

	disputeCheckerStr := disputeChecker[0].String()

	if disputeCheckerStr != "DisputeChecker" {
		t.Fatalf("Expected DisputeChecker but got %s", disputeCheckerStr)
	}

	assert.Equal(t, disputeCheckerStr, "DisputeChecker")

}

func TestCreateBadTracker(t *testing.T) {

	badTracker, err := createTracker("badTracker")

	if err == nil {
		t.Fatalf("expected error but instead received this tracker: %s", badTracker[0].String())
	}

	assert.Len(t, badTracker, 0)
}
