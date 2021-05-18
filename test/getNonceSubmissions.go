package test

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zapproject/zap-miner/apiOracle"
	zapCommon "github.com/zapproject/zap-miner/common"
	zap "github.com/zapproject/zap-miner/contracts"
	zap1 "github.com/zapproject/zap-miner/contracts1"
	"github.com/zapproject/zap-miner/rpc"
)

func getNonceSubmissions(ctx context.Context, valueBlock *big.Int, dispute *zap1.ZapDisputeNewDispute) ([]*apiOracle.PriceStamp, error) {
	instance := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	tokenAbi, err := abi.JSON(strings.NewReader(zap1.ZapLibraryABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse abi: %v", err)
	}
	contractAddress := ctx.Value(zapCommon.ContractAddress).(common.Address)
	client := ctx.Value(zapCommon.ClientContextKey).(rpc.ETHClient)

	//just use nil for most of the variables, only using this object to call UnpackLog which only uses the abi
	bar := bind.NewBoundContract(contractAddress, tokenAbi, nil, nil, nil)

	allVals, err := instance.GetSubmissionsByTimestamp(nil, dispute.RequestId, dispute.Timestamp)
	if err != nil {
		return nil, fmt.Errorf("failed to get other submitted values for dispute: %v", err)
	}

	allAddrs, err := instance.GetMinersByRequestIdAndTimestamp(nil, dispute.RequestId, dispute.Timestamp)
	if err != nil {
		return nil, fmt.Errorf("failed to get miner addresses for dispute: %v", err)
	}

	const blockStep = 67
	high := int64(valueBlock.Uint64())
	low := high - blockStep
	nonceSubmitID := tokenAbi.Events["NonceSubmitted"].ID()
	var timedValues [5]*apiOracle.PriceStamp
	// timedValues := make([]*apiOracle.PriceStamp, 5)
	found := 0
	for found < 5 {
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(low),
			ToBlock:   big.NewInt(high),
			Addresses: []common.Address{contractAddress},
			Topics:    [][]common.Hash{{nonceSubmitID}},
		}

		logs, err := client.FilterLogs(ctx, query)
		if err != nil {
			return nil, fmt.Errorf("failed to get nonce logs: %v", err)
		}

		for _, l := range logs {
			nonceSubmit := zap1.ZapLibraryNonceSubmitted{}
			err := bar.UnpackLog(&nonceSubmit, "NonceSubmitted", l)
			if err != nil {
				return nil, fmt.Errorf("failed to unpack into object: %v", err)
			}
			header, err := client.HeaderByNumber(ctx, big.NewInt(int64(l.BlockNumber)))
			if err != nil {
				return nil, fmt.Errorf("failed to get nonce block header: %v", err)
			}
			for i := 0; i < 5; i++ {
				if nonceSubmit.Miner == allAddrs[i] {
					valTime := time.Unix(int64(header.Time), 0)

					bigF := new(big.Float)
					bigF.SetInt(allVals[i])
					f, _ := bigF.Float64()

					timedValues[i] = &apiOracle.PriceStamp{
						Created:   valTime,
						PriceInfo: apiOracle.PriceInfo{Price: f},
					}
					found++
					break
				}
			}
		}
		high -= blockStep
		low = high - blockStep
	}
	return timedValues[:], nil
}
