package tracker

import (
	"fmt"
	"os"
	"testing"

	"github.com/zapproject/zap-miner/apiOracle"
	"github.com/zapproject/zap-miner/config"
	"github.com/zapproject/zap-miner/util"
)

// TODO: Set threshold low and test the  "out of range" failure
var configJSON = `{
	"publicAddress":"0000000000000000000000000000000000000000",
	"privateKey":"1111111111111111111111111111111111111111111111111111111111111111",
	"contractAddress":"0xCD8a1C3ba11CF5ECfa6267617243239504a98d90",
	"trackers": ["indexers", "balance", "currentVariables", "disputeStatus", "gas", "disputeChecker"],
	"IndexFolder": "..",
	"disputeThreshold": 1.0, 
	"disputeTimeDelta": "50s"
}
`

func TestMain(m *testing.M) {
	err := config.ParseConfigBytes([]byte(configJSON))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse mock config: %v\n", err)
		os.Exit(-1)
	}
	util.ParseLoggingConfig("")
	apiOracle.EnsureValueOracle()
	os.Exit(m.Run())
}
