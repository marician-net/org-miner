package test

import (
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

var configJSON1 = `{
    "zapTokenAddress": "0x5fbdb2315678afecb367f032d93f642f64180aa3",
    "contractAddress": "0xCD8a1C3ba11CF5ECfa6267617243239504a98d90",
    "nodeURL": "http://127.0.0.1:8545/",
    "publicAddress": "cd3B766CCDd6AE721141F452C550Ca635964ce71",
    "privateKey": "8166f546bab6da521a8369cab06c5d2b9e46670292d85c875ee9ec20e84ffb61",
    "serverHost": "localhost",
    "serverPort": 5001,
    "ethClientTimeout": 3000,
    "trackerCycle": 10,
    "requestData":1,
    "gasMultiplier": 1,
    "gasMax":30,
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
    "indexFolder": "indexes",
    "dbFile": "zapDB"
}`

var configJSON2 = `{
    "zapTokenAddress": "0x5fbdb2315678afecb367f032d93f642f64180aa3",
    "contractAddress": "0xCD8a1C3ba11CF5ECfa6267617243239504a98d90",
    "nodeURL": "http://127.0.0.1:8545/",
    "publicAddress": "2546BcD3c84621e976D8185a91A922aE77ECEc30",
    "privateKey": "ea6c44ac03bff858b476bba40716402b03e41b8e97e276d1baec7c37d42484a0",
    "serverHost": "localhost",
    "serverPort": 5001,
    "ethClientTimeout": 3000,
    "trackerCycle": 10,
    "requestData":1,
    "gasMultiplier": 1,
    "gasMax":30,
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
    "indexFolder": "indexes",
    "dbFile": "zapDB"
}`

var configJSON3 = `{
    "zapTokenAddress": "0x5fbdb2315678afecb367f032d93f642f64180aa3",
    "contractAddress": "0xCD8a1C3ba11CF5ECfa6267617243239504a98d90",
    "nodeURL": "http://127.0.0.1:8545/",
    "publicAddress": "bDA5747bFD65F08deb54cb465eB87D40e51B197E",
    "privateKey": "689af8efa8c651a91ad287602527f3af2fe9f6501a7ac4b061667b5a93e037fd",
    "serverHost": "localhost",
    "serverPort": 5001,
    "ethClientTimeout": 3000,
    "trackerCycle": 10,
    "requestData":1,
    "gasMultiplier": 1,
    "gasMax":30,
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
    "indexFolder": "indexes",
    "dbFile": "zapDB"
}`

var configJSON4 = `{
    "zapTokenAddress": "0x5fbdb2315678afecb367f032d93f642f64180aa3",
    "contractAddress": "0xCD8a1C3ba11CF5ECfa6267617243239504a98d90",
    "nodeURL": "http://127.0.0.1:8545/",
    "publicAddress": "dD2FD4581271e230360230F9337D5c0430Bf44C0",
    "privateKey": "de9be858da4a475276426320d5e9262ecfc3ba460bfac56360bfa6c4c28b4ee0",
    "serverHost": "localhost",
    "serverPort": 5001,
    "ethClientTimeout": 3000,
    "trackerCycle": 10,
    "requestData":1,
    "gasMultiplier": 1,
    "gasMax":30,
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
    "indexFolder": "indexes",
    "dbFile": "zapDB"
}`

var configJSON5 = `{
    "zapTokenAddress": "0x5fbdb2315678afecb367f032d93f642f64180aa3",
    "contractAddress": "0xCD8a1C3ba11CF5ECfa6267617243239504a98d90",
    "nodeURL": "http://127.0.0.1:8545/",
    "publicAddress": "8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199",
    "privateKey": "df57089febbacf7ba0bc227dafbffa9fc08a93fdc68e1e42411a14efcf23656e",
    "serverHost": "localhost",
    "serverPort": 5001,
    "ethClientTimeout": 3000,
    "trackerCycle": 10,
    "requestData":1,
    "gasMultiplier": 1,
    "gasMax":30,
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
    "indexFolder": "indexes",
    "dbFile": "zapDB"
}`

func setup() {
	config.ParseConfigBytes([]byte(configJSON))

	util.ParseLoggingConfig("../testConfig.json")

	buildContext()
}

func setupMiners() {
	config.ParseConfigBytes([]byte(configJSON1))
	util.ParseLoggingConfig("../testConfig.json")
	buildMinerContext(0)

	config.ParseConfigBytes([]byte(configJSON2))
	util.ParseLoggingConfig("../testConfig.json")
	buildMinerContext(1)

	config.ParseConfigBytes([]byte(configJSON3))
	util.ParseLoggingConfig("../testConfig.json")
	buildMinerContext(2)

	config.ParseConfigBytes([]byte(configJSON4))
	util.ParseLoggingConfig("../testConfig.json")
	buildMinerContext(3)

	config.ParseConfigBytes([]byte(configJSON5))
	util.ParseLoggingConfig("../testConfig.json")
	buildMinerContext(4)
}
