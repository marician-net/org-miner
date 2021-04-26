// We require the Hardhat Runtime Environment explicitly here. This is optional 
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
import { ethers } from "hardhat"
const hre = require("hardhat");
const ZapMasterJson = require("../artifacts/contracts/ZapMaster.sol/ZapMaster.json")

async function main() {
  let signers = await ethers.getSigners();
  let owner = signers[0]

  const SafeMath = await ethers.getContractFactory("SafeMath", signers[0]);
  const safeMath = await SafeMath.deploy();
  console.log("deployed SafeMath")

  const Utilities = await ethers.getContractFactory("Utilities", signers[0]);
  const utilities = await Utilities.deploy();
  console.log("deployed Utilities")

  const ZapStorage = await ethers.getContractFactory("ZapStorage", signers[0]);
  const zapStorage = await ZapStorage.deploy();
  console.log("deployed ZapStorage")

  const ZapGettersLibrary = await ethers.getContractFactory("ZapGettersLibrary", signers[0]);
  const zapGettersLibrary = await ZapGettersLibrary.deploy();
  console.log("deployed ZapGettersLibrary")

  const ZapTransfer = await ethers.getContractFactory("ZapTransfer", signers[0]);
  const zapTransfer = await ZapTransfer.deploy();
  console.log("deployed ZapTransfer")

  const ZapDispute = await ethers.getContractFactory("ZapDispute", {
    libraries: {
      ZapTransfer: zapTransfer.address,
    },
    signer: signers[0]
  });
  const zapDispute = await ZapDispute.deploy();
  console.log("deployed ZapDispute")

  const ZapStake = await ethers.getContractFactory("ZapStake", {
    libraries: {
      ZapTransfer: zapTransfer.address,
      ZapDispute: zapDispute.address
    },
    signer: signers[0]
  });
  const zapStake= await ZapStake.deploy();
  console.log("deployed ZapStake")

  const ZapLibrary = await ethers.getContractFactory("ZapLibrary", 
  {
    libraries: {
      ZapTransfer: zapTransfer.address,
    },
    signer: signers[0]
  });
  const zapLibrary = await ZapLibrary.deploy();
  console.log("deployed ZapLibrary")

  const Zap = await ethers.getContractFactory("Zap", 
  {
    libraries: {
      ZapStake: zapStake.address,
      ZapDispute: zapDispute.address,
      ZapLibrary: zapLibrary.address
    },
    signer: signers[0]
  });
  const zap = await Zap.deploy("0x5fbdb2315678afecb367f032d93f642f64180aa3");
  console.log("Zap address: " + zap.address)
  console.log("Zap Approve: " + await zap.approve(signers[1].address, 100) );
  console.log("Transfer: " + await zap.transfer("0x70997970C51812dc3A010C7d01b50e0d17dc79C8", 1));

  const ZapGetters = await ethers.getContractFactory("ZapGetters", 
  {
    libraries: {
      ZapTransfer: zapTransfer.address,
    },
    signer: signers[0]
  });
  const zapGetters = await ZapGetters.deploy("0x5fbdb2315678afecb367f032d93f642f64180aa3");
  console.log("deployed ZapGetters")

  // const ZapMaster = await ethers.getContractFactory("ZapMaster", 
  const ZapMaster = await ethers.getContractFactory(ZapMasterJson.abi, ZapMasterJson.bytecode, 
  // {
  //   libraries: {
  //     ZapTransfer: zapTransfer.address,
  //     ZapGettersLibrary: zapGettersLibrary.address,
  //     ZapLibrary: zapLibrary.address,
  //   },
  //   signer: signers[0]
  // }
  signers[0]
  );
  // const ZapMaster = await ethers.getContractAt("ZapMaster", zap.address, signers[0]);

  const zapMaster = await ZapMaster.deploy(zap.address);
  console.log("ZapMaster address: " + zapMaster.address)
  console.log("ZapMaster ABI: " + ZapMasterJson.abi)
  console.log("ZapMaster ByteCode: " + ZapMasterJson.bytecode)
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error);
    process.exit(1);
  });
