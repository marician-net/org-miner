package test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	zapCommon "github.com/zapproject/zap-miner/common"
	zap "github.com/zapproject/zap-miner/contracts"
	"github.com/zapproject/zap-miner/util"
)

func TestBalance(t *testing.T) {
	setup()

	addr := ctx.Value(zapCommon.PublicAddress).(common.Address)
	instance := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	zapBalance, _ := instance.BalanceOf(nil, addr)
	assert.Equal(t, "1500000000000.00", util.FormatERC20Balance(zapBalance), "Incorrect Balance")
}
