package dataServer

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	zapCommon "github.com/zapproject/zap-miner/common"
	"github.com/zapproject/zap-miner/config"
	"github.com/zapproject/zap-miner/contracts"
	"github.com/zapproject/zap-miner/contracts2"
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

var badConfigJSON = `{
	"publicAddress":"0000000000000000000000000000000000000000",
	"privateKey":"1111111111111111111111111111111111111111111111111111111111111111",
	"contractAddress":"0xCD8a1C3ba11CF5ECfa6267617243239504a98d90",
	"trackers": ["badGas"],
	"IndexFolder": "../indexes",
	"disputeThreshold": 1.0,
	"disputeTimeDelta": "50s",
	"nodeURL": "http://127.0.0.1:8545/",
	"serverHost": "localhost",
    "serverPort": 5001,
    "requestData":1,
    "useGPU":false,
    "dbFile": "badZapDB"
}
`

func TestDataServer(t *testing.T) {
	exitCh := make(chan int)
	setup()
	cfg := config.GetConfig()
	if len(cfg.DBFile) == 0 {
		t.Fatal("Missing dbFile config setting")
	}

	DB, err := db.Open(cfg.DBFile)
	if err != nil {
		t.Fatalf("Error in creating a local db: %v", err)
	}
	client, err := rpc.NewClient(cfg.NodeURL)
	if err != nil {
		t.Fatalf("Error in creating a rpc server: %v", err)
	}
	contractAddress := common.HexToAddress(cfg.ContractAddress)
	masterInstance, err := contracts.NewZapMaster(contractAddress, client)
	if err != nil {
		t.Fatalf("Problem creating zap master instance: %v\n", err)
	}

	proxy, err := db.OpenRemoteDB(DB)
	if err != nil {
		t.Fatalf("Problem creating remote db: %v\n", err)
	}

	instance, err := contracts2.NewZap(contractAddress, client)
	if err != nil {
		t.Fatalf("Problem with initializing contracts2: %v\n", err)
	}

	//get address from config
	_fromAddress := cfg.PublicAddress

	//convert to address
	fromAddress := common.HexToAddress(_fromAddress)

	ctx := context.WithValue(context.Background(), zapCommon.DBContextKey, DB)
	ctx = context.WithValue(ctx, zapCommon.ClientContextKey, client)
	ctx = context.WithValue(ctx, zapCommon.MasterContractContextKey, masterInstance)
	ctx = context.WithValue(ctx, zapCommon.DataProxyKey, proxy)
	ctx = context.WithValue(ctx, zapCommon.NewZapContractContextKey, instance)
	balance, err := client.BalanceAt(ctx, fromAddress, nil)
	if err != nil {
		t.Fatalf("Error getting the balance for the address: %s, %v", fromAddress, err)
	}
	ds, err := CreateServer(ctx)
	if err != nil {
		t.Fatalf("Error creating server: %v", err)
	}

	ds.Start(ctx, exitCh)
	// Put request for the remote proxy server
	proxy.Put(db.BalanceKey, []byte(hexutil.EncodeBig(balance)))

	time.Sleep(5000 * time.Millisecond)

	isReady := <-ds.Ready()

	if !isReady {
		t.Fatal("Data Server is not ready")
	}

	time.Sleep(2 * time.Second)
	resp, err := proxy.Get(db.BalanceKey)
	if err != nil {
		t.Fatalf("Error getting the balance: %v", err)
	}
	// defer resp.Body.Close()
	resp_int, _ := hexutil.DecodeBig(string(resp))
	fmt.Printf("Finished, balance: %v\n", resp_int)
	exitCh <- 1
	time.Sleep(1 * time.Second)
	if !ds.Stopped {
		t.Fatal("Did not stop server")
	}
}
