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

func TestCreateTracker(t *testing.T) {

	currentVariablesTracker, _ := createTracker("currentVariables")
	if currentVariablesTracker[0].String() != "CurrentVariablesTracker" {
		t.Fatalf("Expected CurrentVariablesTracker but got %s", currentVariablesTracker[0].String())
	}

	disputeStatusTracker, _ := createTracker("disputeStatus")
	if disputeStatusTracker[0].String() != "DisputeTracker" {
		t.Fatalf("Expected DisputeTracker but got %s", disputeStatusTracker[0].String())
	}

	gasTracker, _ := createTracker("gas")
	if gasTracker[0].String() != "GasTracker" {
		t.Fatalf("Expected GasTracker but got %s", gasTracker[0].String())
	}

	tokenBalanceTracker, _ := createTracker("tokenBalance")
	if tokenBalanceTracker[0].String() != "TokenTracker" {
		t.Fatalf("Expected TokenTracker but got %s", tokenBalanceTracker[0].String())
	}

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
