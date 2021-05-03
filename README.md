# zap-miner

## Installation

```bash
```

## Testing

1) Open terminal #1 
2) Run "cd hardhat" change directory into hardhat/ directory
3) Run "start.sh" script in hardhat/ directory in terminal #1
 <!-- Run "go test" inside a package directory (i.e. /ops) in terminal #2 -->

## Execute
(If running on local execute steps 1 - 3 from above)
1) Run "release_build.sh"
2) ./zap-miner [cmd]


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