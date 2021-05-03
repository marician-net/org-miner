const { task } = require("hardhat/config");
require("hardhat-deploy-ethers");
require("hardhat-deploy");

task("query", "Test Query")

    .setAction(async () => {

        const signers = await ethers.getSigners();

        const zap = await ethers.getContractAt(
            "Zap",
            '0xb7278A61aa25c888815aFC32Ad3cC52fF24fE575',
            signers[0]
        );


        await zap.requestData(
            "json(https://api.gdax.com/products/BTC-USD/ticker).price",
            "USD",
            1000,
            0
        )

    })
