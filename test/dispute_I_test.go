package test

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	zapCommon "github.com/zapproject/zap-miner/common"
	"github.com/zapproject/zap-miner/config"
	"github.com/zapproject/zap-miner/contracts"
	zap "github.com/zapproject/zap-miner/contracts"
	zap1 "github.com/zapproject/zap-miner/contracts1"
	"github.com/zapproject/zap-miner/ops"
	"github.com/zapproject/zap-miner/rpc"
	token "github.com/zapproject/zap-miner/token"
	"github.com/zapproject/zap-miner/tracker"
	"github.com/zapproject/zap-miner/util"
)

func TestDispute(t *testing.T) {
	setupMiners()
	setup()

	// [show] check if there are existing disputes - should be 0
	tokenAbi, _ := abi.JSON(strings.NewReader(zap1.ZapDisputeABI))
	contractAddress := ctx.Value(zapCommon.ContractAddress).(common.Address)
	client := ctx.Value(zapCommon.ClientContextKey).(rpc.ETHClient)
	header, _ := client.HeaderByNumber(ctx, nil)
	startBlock := big.NewInt(54) //big.NewInt(10e3 * 14)
	startBlock.Sub(header.Number, startBlock)
	newDisputeID := tokenAbi.Events["NewDispute"].ID()
	query := ethereum.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   nil,
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{newDisputeID}},
	}
	logs, _ := client.FilterLogs(ctx, query)
	assert.Equal(t, 0, len(logs), "There should be 0 disputes")

	// start miners
	timestamp := StartMiners(t)

	// [new] dispute last miner as first miner
	MakeNew(t, timestamp)
}

func MakeNew(t *testing.T, timestamp *big.Int) {
	auth, _ := ops.PrepareEthTransaction(ctx)
	instance2 := ctx.Value(zapCommon.TransactorContractContextKey).(*zap1.ZapTransactor)
	addr1 := common.HexToAddress("0xCD8a1C3ba11CF5ECfa6267617243239504a98d90")
	addr2 := common.HexToAddress("0xb7278a61aa25c888815afc32ad3cc52ff24fe575")
	amt1 := big.NewInt(10000)
	instance := ctx.Value(zapCommon.TokenTransactorContractContextKey).(*token.ZapTokenTransactor)
	instance.Approve(auth, addr1, amt1)
	auth, _ = ops.PrepareEthTransaction(ctx)
	instance.Approve(auth, addr2, amt1)

	array := [32]byte{}
	data := []byte("disputeFee")
	data = crypto.Keccak256(data)
	for i := 0; i < 32; i++ {
		array[i] = data[i]
	}
	master := ctx.Value(zapCommon.MasterContractContextKey).(*contracts.ZapMaster)
	blocknum, _ := master.GetUintVar(nil, array)
	fmt.Println("Dispute Fee: ", blocknum)

	auth, _ = ops.PrepareEthTransaction(ctx)
	instance2.BeginDispute(auth, big.NewInt(1), timestamp, big.NewInt(4))
	// fmt.Printf("dispute started with txn: %s\n", tx)
}

func Show(t *testing.T) {

	cfg := config.GetConfig()

	tokenAbi, _ := abi.JSON(strings.NewReader(zap1.ZapDisputeABI))
	contractAddress := ctx.Value(zapCommon.ContractAddress).(common.Address)
	client := ctx.Value(zapCommon.ClientContextKey).(rpc.ETHClient)

	////just use nil for most of the variables, only using this object to call UnpackLog which only uses the abi
	bar := bind.NewBoundContract(contractAddress, tokenAbi, nil, nil, nil)

	header, _ := client.HeaderByNumber(ctx, nil)
	startBlock := big.NewInt(54) //big.NewInt(10e3 * 14)
	startBlock.Sub(header.Number, startBlock)
	newDisputeID := tokenAbi.Events["NewDispute"].ID()
	query := ethereum.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   nil,
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{newDisputeID}},
	}

	logs, _ := client.FilterLogs(ctx, query)
	instance := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	assert.Equal(t, 0, len(logs), "There should be 0 disputes")

	fmt.Printf("There are currently %d open disputes\n", len(logs))
	fmt.Printf("-------------------------------------\n")
	for _, rawDispute := range logs {
		dispute := zap1.ZapDisputeNewDispute{}
		bar.UnpackLog(&dispute, "NewDispute", rawDispute)
		_, executed, votePassed, _, reportedAddr, reportingMiner, _, uintVars, currTally, _ := instance.GetAllDisputeVars(nil, dispute.DisputeId)

		votingEnds := time.Unix(uintVars[3].Int64(), 0)
		createdTime := votingEnds.Add(-7 * 24 * time.Hour)

		var descString string
		if executed {
			descString = "complete, "
			if votePassed {
				descString += "successful"
			} else {
				descString += "rejected"
			}
		} else {
			descString = "in progress"
		}

		fmt.Printf("Dispute %s (%s):\n", dispute.DisputeId.String(), descString)
		fmt.Printf("    Accused Party: %s\n", reportedAddr.Hex())
		fmt.Printf("    Disputed by: %s\n", reportingMiner.Hex())
		fmt.Printf("    Created on:  %s\n", createdTime.Format("3:04 PM January 02, 2006 MST"))
		fmt.Printf("    Fee: %s ZAP\n", util.FormatERC20Balance(uintVars[8]))
		fmt.Printf("    \n")
		fmt.Printf("    Value disputed for requestID %d:\n", dispute.RequestId.Uint64())

		allSubmitted, _ := getNonceSubmissions(ctx, uintVars[5], &dispute)
		disputedValTime := allSubmitted[uintVars[6].Uint64()].Created

		for i := len(allSubmitted) - 1; i >= 0; i-- {
			sub := allSubmitted[i]
			valStr := fmt.Sprintf("%f\n", sub.Price)
			var pointerStr string
			if i == int(uintVars[6].Uint64()) {
				pointerStr = " <--disputed"
			}

			fmt.Printf("      %s @ %s%s\n", valStr, sub.Created.Format("3:04:05 PM"), pointerStr)
		}
		fmt.Printf("    \n")

		tmp := new(big.Float)
		tmp.SetInt(currTally)
		currTallyFloat, _ := tmp.Float64()
		tmp.SetInt(uintVars[7])
		currQuorum, _ := tmp.Float64()
		currTallyFloat += currQuorum
		currTallyRatio := currTallyFloat / 2 * currQuorum
		fmt.Printf("    Currently %.0f%% of %s ZAP support this dispute (%s votes)\n", currTallyRatio*100, util.FormatERC20Balance(uintVars[7]), uintVars[4])

		result := tracker.CheckValueAtTime(dispute.RequestId.Uint64(), uintVars[2], disputedValTime)
		if result == nil || len(result.Datapoints) < 0 {
			fmt.Printf("      No data available for recommendation\n")
			continue
		}
		fmt.Printf("      Recommendation:\n")
		fmt.Printf("      Vote %t\n", !result.WithinRange)
		fmt.Printf("      Submitted value %s, expected range %.0f to %0.f\n", uintVars[2].String(), result.Low, result.High)
		numToShow := 3
		if numToShow > len(result.Datapoints) {
			numToShow = len(result.Datapoints)
		}
		fmt.Printf("      Based on %d locally saved datapoints within %.0f minutes (showing closest %d)\n",
			len(result.Datapoints), cfg.DisputeTimeDelta.Duration.Minutes(), numToShow)
		minTotalDelta := time.Duration(math.MaxInt64)
		index := 0
		for i := 0; i < len(result.Datapoints)-numToShow; i++ {
			totalDelta := time.Duration(0)
			for j := 0; j < numToShow; j++ {
				delta := result.Times[i+j].Sub(disputedValTime)
				if delta < 0 {
					delta = -delta
				}
				totalDelta += delta
			}
			if totalDelta < minTotalDelta {
				minTotalDelta = totalDelta
				index = i
			}
		}
		for i := 0; i < numToShow; i++ {
			dp := result.Datapoints[index+i]
			t := result.Times[index+i]
			fmt.Printf("        %f, ", dp)
			delta := disputedValTime.Sub(t)
			if delta > 0 {
				fmt.Printf("%.0fs before\n", delta.Seconds())
			} else {
				fmt.Printf("%.0fs after\n", (-delta).Seconds())
			}
		}
	}
}

func StartMiners(t *testing.T) *big.Int {
	// make requestData
	auth, _ := ops.PrepareEthTransaction(ctx)
	instance1 := ctx.Value(zapCommon.TransactorContractContextKey).(*zap1.ZapTransactor)
	instance1.RequestData(auth,
		"json(https://api.binance.com/api/v1/klines?symbol=BTCUSDT&interval=1d&limit=1).0.4", "BTC/USD",
		new(big.Int).SetInt64(10000), new(big.Int).SetInt64(0))

	// submit solution for 5 miners
	for i := 0; i < 5; i++ {
		value := int64(100) + int64(i)
		authx, _ := ops.PrepareEthTransaction(minerCtx[i])
		instancex := minerCtx[i].Value(zapCommon.TransactorContractContextKey).(*zap1.ZapTransactor)
		instancex.SubmitMiningSolution(authx, strconv.Itoa(i+1), big.NewInt(1), big.NewInt(value))
	}

	array := [32]byte{}
	data := []byte("timeOfLastNewValue")
	data = crypto.Keccak256(data)
	for i := 0; i < 32; i++ {
		array[i] = data[i]
	}

	instance := ctx.Value(zapCommon.MasterContractContextKey).(*contracts.ZapMaster)
	uvar, _ := instance.GetUintVar(nil, array)
	return uvar
}
