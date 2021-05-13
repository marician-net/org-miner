package test

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	ZapCommon "github.com/zapproject/zap-miner/common"
	"github.com/zapproject/zap-miner/config"
	"github.com/zapproject/zap-miner/contracts"
	"github.com/zapproject/zap-miner/contracts1"
	"github.com/zapproject/zap-miner/contracts2"
	"github.com/zapproject/zap-miner/rpc"
	"github.com/zapproject/zap-miner/token"
)

var ctx context.Context

func buildContext() error {
	cfg := config.GetConfig()

	if !cfg.EnablePoolWorker {
		//create an rpc client
		client, err := rpc.NewClient(cfg.NodeURL)
		if err != nil {
			log.Fatal(err)
		}

		//create an instance of the Zap master contract for on-chain interactions
		tokenAddress := common.HexToAddress(cfg.TokenAddress)
		contractAddress := common.HexToAddress(cfg.ContractAddress)
		masterInstance, err := contracts.NewZapMaster(contractAddress, client)
		transactorInstance, err := contracts1.NewZapTransactor(contractAddress, client)
		newZapInstance, err := contracts2.NewZap(contractAddress, client)
		newTransactorInstance, err := contracts2.NewZapTransactor(contractAddress, client)
		tokenInstance, err := token.NewZapTokenTransactor(tokenAddress, client)
		tokenListener, err := token.NewZapTokenFilterer(tokenAddress, client)
		if err != nil {
			log.Fatal(err)
		}

		ctx = context.WithValue(context.Background(), ZapCommon.ClientContextKey, client)
		ctx = context.WithValue(ctx, ZapCommon.ContractAddress, contractAddress)
		ctx = context.WithValue(ctx, ZapCommon.MasterContractContextKey, masterInstance)
		ctx = context.WithValue(ctx, ZapCommon.TransactorContractContextKey, transactorInstance)
		ctx = context.WithValue(ctx, ZapCommon.TokenTransactorContractContextKey, tokenInstance)
		ctx = context.WithValue(ctx, ZapCommon.NewZapContractContextKey, newZapInstance)
		ctx = context.WithValue(ctx, ZapCommon.NewTransactorContractContextKey, newTransactorInstance)
		ctx = context.WithValue(ctx, ZapCommon.TokenFilterContextKey, tokenListener)

		privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
		if err != nil {
			return fmt.Errorf("problem getting private key: %s", err.Error())
		}
		ctx = context.WithValue(ctx, ZapCommon.PrivateKey, privateKey)

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return fmt.Errorf("error casting public key to ECDSA")
		}

		publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		ctx = context.WithValue(ctx, ZapCommon.PublicAddress, publicAddress)

		//Issue #55, halt if client is still syncing with Ethereum network
		s, err := client.IsSyncing(ctx)
		if err != nil {
			return fmt.Errorf("could not determine if Ethereum client is syncing: %v\n", err)
		}
		if s {
			return fmt.Errorf("ethereum node is still sycning with the network")
		}
	}
	return nil
}
