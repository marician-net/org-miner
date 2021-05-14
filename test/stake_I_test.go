package test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	zapCommon "github.com/zapproject/zap-miner/common"
	zap "github.com/zapproject/zap-miner/contracts"
	zap1 "github.com/zapproject/zap-miner/contracts1"
	"github.com/zapproject/zap-miner/ops"
)

func TestStake(t *testing.T) {
	setup()

	tmaster := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	publicAddress := ctx.Value(zapCommon.PublicAddress).(common.Address)
	status, _, _ := tmaster.GetStakerInfo(nil, publicAddress)

	assert.Equal(t, big.NewInt(1).Uint64(), status.Uint64(), "Staker status should be staked a inital miner 0 - status should be 1")
	RequestWithdraw(t)
	Withdraw(t)
	Deposit(t)
}

func Deposit(t *testing.T) {
	tmaster := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	publicAddress := ctx.Value(zapCommon.PublicAddress).(common.Address)
	balance, _ := tmaster.BalanceOf(nil, publicAddress)
	status, _, _ := tmaster.GetStakerInfo(nil, publicAddress)
	fmt.Println("Depositing: ", status)
	assert.Equal(t, big.NewInt(0).Uint64(), status.Uint64(), "Staker status does not match - status should be 0")

	dat := crypto.Keccak256([]byte("stakeAmount"))
	var dat32 [32]byte
	copy(dat32[:], dat)
	stakeAmt, _ := tmaster.GetUintVar(nil, dat32)
	assert.Greater(t, balance.Cmp(stakeAmt), 0, "Account 0 does not have enough Zap to stake")

	instance2 := ctx.Value(zapCommon.TransactorContractContextKey).(*zap1.ZapTransactor)
	auth, _ := ops.PrepareEthTransaction(ctx)
	instance2.DepositStake(auth)

	status, _, _ = tmaster.GetStakerInfo(nil, publicAddress)
	assert.Equal(t, big.NewInt(1).Uint64(), status.Uint64(), "Staker status does not match - status should be 1")
}

func Withdraw(t *testing.T) {
	tmaster := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	publicAddress := ctx.Value(zapCommon.PublicAddress).(common.Address)
	status, _, _ := tmaster.GetStakerInfo(nil, publicAddress)
	assert.Equal(t, big.NewInt(2).Uint64(), status.Uint64(), "Staker has not requested for withdrawal - status should be 2")

	auth, _ := ops.PrepareEthTransaction(ctx)

	instance2 := ctx.Value(zapCommon.TransactorContractContextKey).(*zap1.ZapTransactor)
	instance2.WithdrawStake(auth)
	status, _, _ = tmaster.GetStakerInfo(nil, publicAddress)
	fmt.Println("Withdrew: ", status)
	assert.Equal(t, big.NewInt(0).Uint64(), status.Uint64(), "Withdraw was not successful - status should be 0")
}

func RequestWithdraw(t *testing.T) {
	tmaster := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	publicAddress := ctx.Value(zapCommon.PublicAddress).(common.Address)
	status, _, _ := tmaster.GetStakerInfo(nil, publicAddress)
	assert.Equal(t, big.NewInt(1).Uint64(), status.Uint64(), "Account 0 is not staked - status should be 1")

	auth, _ := ops.PrepareEthTransaction(ctx)

	instance2 := ctx.Value(zapCommon.TransactorContractContextKey).(*zap1.ZapTransactor)
	instance2.RequestStakingWithdraw(auth)

	status, _, _ = tmaster.GetStakerInfo(nil, publicAddress)
	assert.Equal(t, big.NewInt(2).Uint64(), status.Uint64(), "Request withdraw was not successful - status should be 2")
}
