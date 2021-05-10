package tracker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBalanceTracker(t *testing.T) {

	balanceTracker, _ := createTracker("balance")

	balanceTrackerStr := balanceTracker[0].String()

	if balanceTrackerStr != "BalanceTracker" {
		t.Fatalf("Expected BalanceTracker but got %s", balanceTrackerStr)
	}

	assert.Equal(t, balanceTrackerStr, "BalanceTracker")

}

func TestCreateCurrentVariablesTracker(t *testing.T) {

	currentVariablesTracker, _ := createTracker("currentVariables")

	currentVariablesStr := currentVariablesTracker[0].String()

	if currentVariablesStr != "CurrentVariablesTracker" {
		t.Fatalf("Expected CurrentVariablesTracker but got %s", currentVariablesStr)
	}

	assert.Equal(t, currentVariablesStr, "CurrentVariablesTracker")

}

func TestCreateDisputeStatusTracker(t *testing.T) {

	disputeStatusTracker, _ := createTracker("disputeStatus")

	disputeStatusStr := disputeStatusTracker[0].String()

	if disputeStatusStr != "DisputeTracker" {
		t.Fatalf("Expected DisputeTracker but got %s", disputeStatusStr)
	}

	assert.Equal(t, disputeStatusStr, "DisputeTracker")

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

func TestCreateTracker(t *testing.T) {

	indexersTracker, err := createTracker("indexers")
	if err != nil {
		t.Fatalf("Could not build IndexTracker")
	}
	if len(indexersTracker) == 0 {
		t.Fatalf("Could not build all IndexTrackers: only tracking %d indexes", len(indexersTracker))
	}

	disputeChecker, _ := createTracker("disputeChecker")
	if disputeChecker[0].String() != "DisputeChecker" {
		t.Fatalf("Expected DisputeChecker but got %s", disputeChecker[0].String())
	}

	badTracker, err := createTracker("badTracker")
	if err == nil {
		t.Fatalf("expected error but instead received this tracker: %s", badTracker[0].String())
	}

}
