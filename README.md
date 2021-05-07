# zap-miner

## Installation

1) Install go-lang https://golang.org/doc/install
2) Run "npm i" in hardhat/ directory
    a) To install npm, linux: sudo apt install nodejs | wsl: https://stackoverflow.com/questions/49919063/installing-npm-node-on-bash-on-ubuntu-on-windows-wsl-what-architecture-does-u

## Testing


## Node Setup
1) Open terminal #1 
2) Run "cd hardhat" change directory into hardhat/ directory
3) Run "start.sh" script in hardhat/ directory in terminal #1
 <!-- Run "go test" inside a package directory (i.e. /ops) in terminal #2 -->

## Execute
(If running on local execute steps 1 - 3 from above)
1) Run "release_build.sh"
2) ./zap-miner [cmd]

In order to run dispute commands: (inefficient)
1) Run "./zap-miner --config local_cfgs/config1.json mine"
2) CTRL+C when "Successfully submitted solution" is logged
3) Run "./zap-miner --config local_cfgs/config2.json mine"
4) CTRL+C when "Successfully submitted solution" is logged
5) Run "./zap-miner --config local_cfgs/config3.json mine"
6) CTRL+C when "Successfully submitted solution" is logged
7) Run "./zap-miner --config local_cfgs/config4.json mine"
8) CTRL+C when "Successfully submitted solution" is logged
9) Run "./zap-miner --config local_cfgs/config5.json mine"
10) CTRL+C when "Successfully submitted solution" is logged
11) Run "./zap-miner dataserver"
12) Locate log with "TimeStamp: %!(EXTRA *big.Int=XXX)" and copy the big.Int value
13) Run "./zap-miner approve 10000 0xCD8a1C3ba11CF5ECfa6267617243239504a98d90"
14) Run "./zap-miner approve 10000 0xb7278a61aa25c888815afc32ad3cc52ff24fe575"
15) Run "./zap-miner dispute new 1 {TimeStamp value copied from step 12} 4
- This command disputes miner 5 


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
