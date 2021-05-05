package tracker

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	zapCommon "github.com/zapproject/zap-miner/common"
	"github.com/zapproject/zap-miner/config"
	"github.com/zapproject/zap-miner/db"
	"github.com/zapproject/zap-miner/rpc"
	"github.com/zapproject/zap-miner/util"
)

func setup() {
	err := config.ParseConfig("../config.json")
	if err != nil {
		fmt.Errorf("Can't parse config for test.")
	}
	path := "../testConfig.json"
	err = util.ParseLoggingConfig(path)
	if err != nil {
		fmt.Errorf("Can't parse logging config for test.")
	}
}

func TestDisputeCheckerInRange(t *testing.T) {
	setup()
	opts := &rpc.MockOptions{ETHBalance: big.NewInt(300000), Nonce: 1, GasPrice: big.NewInt(7000000000),
		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}}
	disputeChecker := &disputeChecker{lastCheckedBlock: 500}
	DB, err := db.Open(filepath.Join(os.TempDir(), "disputeChecker_test"))
	if err != nil {
		t.Fatal(err)
	}
	client := rpc.NewMockClientWithValues(opts)
	ctx := context.WithValue(context.Background(), zapCommon.ClientContextKey, client)
	ctx = context.WithValue(ctx, zapCommon.DBContextKey, DB)
	BuildIndexTrackers()
	ethUSDPairs := indexes["ETH/USD"]
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	time.Sleep(2 * time.Second)
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	ctx = context.WithValue(ctx, zapCommon.ContractAddress, common.Address{0x0000000000000000000000000000000000000000})
	err = disputeChecker.Exec(ctx)
	if err != nil {
		t.Fatal(err)
	}
	DB.Close()
}

func TestDisputeCheckerOutOfRange(t *testing.T) {
	cfg := config.GetConfig()
	cfg.DisputeThreshold = 0.000000001
	opts := &rpc.MockOptions{ETHBalance: big.NewInt(300000), Nonce: 1, GasPrice: big.NewInt(7000000000),
		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}}
	disputeChecker := &disputeChecker{lastCheckedBlock: 500}
	DB, err := db.Open(filepath.Join(os.TempDir(), "disputeChecker_test"))
	if err != nil {
		t.Fatal(err)
	}
	client := rpc.NewMockClientWithValues(opts)
	ctx := context.WithValue(context.Background(), zapCommon.ClientContextKey, client)
	ctx = context.WithValue(ctx, zapCommon.DBContextKey, DB)
	BuildIndexTrackers()
	ethUSDPairs := indexes["ETH/USD"]
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	time.Sleep(2 * time.Second)
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	ctx = context.WithValue(ctx, zapCommon.ContractAddress, common.Address{0x0000000000000000000000000000000000000000})
	err = disputeChecker.Exec(ctx)
	if err != nil {
		t.Fatal(err)
	}
	DB.Close()
}

func execEthUsdPsrs(ctx context.Context, t *testing.T, psrs []*IndexTracker) {
	for _, psr := range psrs {
		//fmt.Print("\nIndex: ", psrIdx, psrs[psrIdx])
		err := psr.Exec(ctx)
		if err != nil {
			t.Fatalf("failed to execute psr: %v", err)
		}
	}
}
