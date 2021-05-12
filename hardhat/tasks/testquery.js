const { task, taskArgs } = require("hardhat/config");
require("hardhat-deploy-ethers");
require('hardhat-deploy');


task("testquery", "Test")

       .setAction(async () => {

              // Test accounts
              const signers = await ethers.getSigners();

              const Zap = await ethers.getContractAt(
                     "Zap",
                     "0xb7278A61aa25c888815aFC32Ad3cC52fF24fE575",
                     signers[0]
              )

              const query = await Zap.requestData(
                     "json(https://api.gdax.com/products/BTC-USD/ticker).price",
                     "USD",
                     1000,
                     0
              )

              console.log(query)

       });