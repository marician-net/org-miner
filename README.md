# zap-miner

## BSC Testnet Contract Addresses
```
BSC TESTNET MINER CONTRACTS
ZapGettersLibary Address: 0x57b23103867b11F64dB2Ade885B7c3655F6be3c8
ZapTransfer Address: 0xAd6f4151ef0c7D49fb40CDd6e7d3Ed8977543080
ZapDispute Address: 0xF23937335c0794ea9920aC06c0f7947407919718
ZapStake Address: 0x00cf02aB915A9026C2328B07E9E45A854CCb28fd
ZapLibrary Address: 0x5896c4024de2DA768433123d442d8Ff2C1Ec694B
Zap Address: 0x4646939E9336139b16213B06EaCbA0a53b999f94
ZapMaster Address: 0xe623305CC15792f4e3E4Cd8880B5B9D665976df5
Vault Address: 0xc91fE0A6546f599b2887F342F5074E33D4fA37eE
```

## Installation

1) Install go-lang https://golang.org/doc/install
2) Run "npm i" in hardhat/ directory
    a) To install npm, linux: sudo apt install nodejs | wsl: https://stackoverflow.com/questions/49919063/installing-npm-node-on-bash-on-ubuntu-on-windows-wsl-what-architecture-does-u
 
## Node Setup
1) Open terminal #1 
2) Run "cd hardhat" change directory into hardhat/ directory
3) Run "start.sh" script in hardhat/ directory in terminal #1
 <!-- Run "go test" inside a package directory (i.e. /ops) in terminal #2 -->

## Execute (Single Miner/Thread)
(If running on local execute steps 1 - 3 from above)
1) Run "release_build.sh"
2) ./zap-miner [cmd]

## Execute (Remote Mining/Multi Miners)
(If running on local execute steps 1 - 3 from Node Setup)
1) Run "release_build.sh"
2) Run "./start_local.sh"
3) Check the logs/ for failing miners. Manually run "nohup ./zap-miner --config=local_cfgs/config{miner # 1-5}.json mine -r > logs/{miner # 1-5}.log &" for failing miners.

In order to run dispute commands: (inefficient)
4) Locate log with "TimeStamp: %!(EXTRA *big.Int=XXX)" and copy the big.Int value
5) Run "./zap-miner dispute new 1 {TimeStamp value copied from step 12} 4

## Subgraph

https://thegraph.com/explorer/subgraph/acemasterjb/zapminer

This subgraph indexes events on the BSC Testnet.

## Config.json
contractAddress (required) - address of Berry Contract
nodeURL (required) - When you connect to BSC Node.
publicAddress (required) - public address for your miner (note, no 0x)
privateKey - private key for miner node to sign transaction, without prefix 0x
ethClientTimeout (required) - timeout for making requests from your node
trackerCycle (required) - how often your database updates (in seconds)
trackers (required) - which pieces of the database you update
dbFile (required) - where you want to store your local database (if self-hosting)
serverHost (required) - location to host server
serverWhitelist (required) - whitelists which publicAddress can access the data server
fetchTimeout - timeout for requesting data from an API
requestData - sets wether your miner request data if the challenge is 0.  If yes,
    then you will addTip() to this number.  Enter a uint number representing request id to be requested
requestDataInterval - min frequency at which to request data at (in seconds, default 30)
gasMultiplier - Multiplies the submitted gasPrice
gasMax - a max for the gas price in gwei (note: this max comes BEFORE the gas multiplier.
    So a max gas cost of 10 gwei, can have gas prices up to 20 if gasMultiplier is 2)
heartbeat - an integer that controls how frequently the miner process should report the hashrate (larger is less frequent, try 1000000 to start)
numProcessors - an integer number of CPU cores/threads to use for mining.
   (cpu mining is disabled if there is a suitable GPU is found
disputeTimeDelta - how far back to store values for min/max range - default 5 (in minutes)
disputeThreshold - percentage of acceptable range outside min/max for dispute checking -
    default

## Testing
You have a few options when you want to test zap-miner
### Test Scripts
In the root of this project, run:

```bash
./runPkgTest.sh <PACKAGE_NAME>
```

This will run all tests, identified by `*_test.go`, in the given package. It will also give you coverage insights.
If instead you'd like to test a specific test in a package run:

```bash
./runTest.sh <TEST_NAME> <PACKAGE_NAME>
```
e.g. `./runTest.sh TestDataServer dataServer`

This will run the given `<TEST_NAME>` as long as it is inside a `*_test.go` file in the given `<PACKAGE_NAME>`.

### Go test from CLI
You could of course also `cd` into the package of choice and run the standard `go test [test flags...]` e.g.
```bash
cd pow
go test -v -cover  # Test with increased verbosity and give a coverage analysis
```

You can also run individual test just like when running the `./runTest.sh` script
```bash
cd zap-miner # while in the root of the project...
go test -v [test flags...] [PACKAGE_NAME] -run [TEST_NAME]
```

For more `go test` flags you can use, run this command in your terminal
```bash
go help testflag
```


## **Debug**

### Install
Delve is a debugger for the Go programming language. Follow the steps in this [repo](https://github.com/go-delve/delve) to install onto your machine.

*For debugging, we needed to take out the -s and -w flags in the ./release-build.sh script.*
1) Follow the steps above up until Execute - 1. Don't run the ```./zap-miner``` command.
2) Instead run ```dlv exec ./zap-miner [command]```. This allows the debugger to run the script.
3) Set breakpoint(s).
4) Enter ```continue``` command to allow the program to run till breakpoint.
5) Check variables using ```locals``` and/or ```print```.
6) Step In and Out of functions with ```step``` and ```stepout```.
7) Use ```restart``` to start over. Breakpoints will persist.
8) Use ```quit``` to exit debug mode.



### **Basic Debug Commands**
[continue](#continue) | Run until breakpoint or program termination.  
[break](#break) | Sets a breakpoint. (EX. ```break tracker/index.go:39``` That will set a breakpoint on line 39 in tracker.index.go file.)  
[breakpoints](#breakpoints) | Print out info for active breakpoints.  
[step](#step) | Single step through program.  
[stepout](#stepout) | Step out of the current function.  
[locals](#locals) | Print local variables.  
[print](#print) | Evaluate an expression.  
[restart](#restart) | Restart process.  


More commands here ```$GOPATH/src/github.com/go-delve/delve/tree/master/Documentation/cli/locspec.md``` or type ```help``` when in debug mode.

### **Docker**
``` docker build -t zapminer:branchname . ```
