package test

import (
	"fmt"
	"os"

	"github.com/zapproject/zap-miner/config"
	"github.com/zapproject/zap-miner/util"
)

var configJSON = `{
	"zapTokenAddress": "0x5fbdb2315678afecb367f032d93f642f64180aa3",
    "contractAddress": "0xCD8a1C3ba11CF5ECfa6267617243239504a98d90",
    "nodeURL": "http://127.0.0.1:8545/",
    "publicAddress": "f39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
    "privateKey": "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
    "serverHost": "localhost",
    "serverPort": 5001,
    "ethClientTimeout": 3000,
    "trackerCycle": 10,
    "requestData":1,
    "gasMultiplier": 1,
    "gasMax":30,
    "serverWhitelist": ["0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266", "0xcd3B766CCDd6AE721141F452C550Ca635964ce71", "0x2546BcD3c84621e976D8185a91A922aE77ECEc30", "0xbDA5747bFD65F08deb54cb465eB87D40e51B197E", "0xdD2FD4581271e230360230F9337D5c0430Bf44C0", "0x8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199"],
    "useGPU":false,
    "trackers": [
        "balance",
        "disputeStatus",
        "gas",
        "tokenBalance",
        "indexers",
        "newCurrentVariables",
        "currentVariables"
    ],
    "indexFolder": "..",
    "dbFile": "zapDB"
}
`

func setup() {
	err := config.ParseConfigBytes([]byte(configJSON))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse mock config: %v\n", err)
		os.Exit(-1)
	}

	err = util.ParseLoggingConfig("../testConfig.json")

	buildContext()
}
