package ops

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	zapCommon "github.com/zapproject/zap-miner/common"
	"github.com/zapproject/zap-miner/config"
	"github.com/zapproject/zap-miner/contracts"
	"github.com/zapproject/zap-miner/contracts1"
	"github.com/zapproject/zap-miner/contracts2"
	"github.com/zapproject/zap-miner/db"
	"github.com/zapproject/zap-miner/rpc"
)

var ctx context.Context

var (
	requestID *big.Int
	amount    *big.Int
)

type testContract struct {
}

func (t testContract) AddTip(_requestID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	fmt.Printf("Contract simulation adding tip: %v, %v\n", _requestID, _amount)
	requestID = _requestID
	amount = _amount
	return nil, nil
}

func (t testContract) SubmitSolution(solution string, requestID *big.Int, value *big.Int) (*types.Transaction, error) {
	return nil, nil
}

func (t testContract) DidMine(challenge [32]byte) (bool, error) {
	return false, nil
}

type testSubmit struct {
	contract  *zapCommon.ContractInterface
	submitter *zapCommon.TransactionSubmitter
}

func (t testSubmit) PrepareTransaction(ctx context.Context, ctxName string, fn zapCommon.TransactionGeneratorFN) error {
	_, err := fn(ctx, *t.contract)
	return err
}

func TestRequestDataOps(t *testing.T) {
	exitCh := make(chan os.Signal)
	setup()
	cfg := config.GetConfig()

	submitter := NewSubmitter()
	DB, err := db.Open(cfg.DBFile)

	//delete any request id
	DB.Delete(db.RequestIdKey)

	if err != nil {
		log.Fatal(err)
	}

	proxy, err := db.OpenRemoteDB(DB)
	if err != nil {
		log.Fatal(err)
	}

	client, err := rpc.NewClient(cfg.NodeURL)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(cfg.ContractAddress)

	transactor1Instance, err := contracts1.NewZapTransactor(contractAddress, client)
	if err != nil {
		t.Fatalf("Problem with initializing the ZapTransactor: %v\n", err)
	}

	transactor2Instance, err := contracts2.NewZapTransactor(contractAddress, client)
	if err != nil {
		t.Fatalf("Problem with initializing the ZapTransactor: %v\n", err)
	}

	masterInstance, err := contracts.NewZapMaster(contractAddress, client)
	if err != nil {
		t.Fatalf("Problem creating zap master instance: %v\n", err)
	}

	ctx := context.WithValue(context.Background(), zapCommon.DBContextKey, DB)
	ctx = context.WithValue(ctx, zapCommon.DataProxyKey, proxy)
	proxy = ctx.Value(zapCommon.DataProxyKey).(db.DataServerProxy)
	ctx = context.WithValue(ctx, zapCommon.ClientContextKey, client)
	ctx = context.WithValue(ctx, zapCommon.TransactorContractContextKey, transactor1Instance)
	ctx = context.WithValue(ctx, zapCommon.NewTransactorContractContextKey, transactor2Instance)
	ctx = context.WithValue(ctx, zapCommon.MasterContractContextKey, masterInstance)
	reqData := CreateDataRequester(exitCh, submitter, 2, proxy)

	//it should not request data if not configured to do it
	cfg.RequestData = 0
	reqData.Start(ctx)
	time.Sleep(300 * time.Millisecond)
	if reqData.submittingRequests {
		t.Fatal("Should not be submitting requests without configured request id")
	}

	cfg.RequestData = 1
	DB.Put(db.RequestIdKey, []byte(hexutil.EncodeBig(big.NewInt(0))))
	reqData.Start(ctx)
	time.Sleep(2500 * time.Millisecond)
	reqIdBytes, err := DB.Get(db.RequestIdKey)
	if err != nil {
		t.Fatal("Error in getting Request ID")
	}
	reqIdInt, _ := strconv.Atoi(string(reqIdBytes))
	requestID = big.NewInt(int64(reqIdInt))
	log.Printf("Request ID: %d", requestID)
	if requestID == nil {
		t.Fatal("Should have requested data")
	}
	requestID = nil
	DB.Put(db.RequestIdKey, []byte(hexutil.EncodeBig(big.NewInt(1))))
	time.Sleep(2500 * time.Millisecond)
	if requestID != nil {
		t.Fatal("Should not have requested data when a challenge request is in progress")
	}

	exitCh <- os.Kill
	time.Sleep(300 * time.Millisecond)
	if reqData.submittingRequests {
		t.Fatal("Should not be submitting requests after exit sig")
	}
}
