import { HardhatUserConfig } from "hardhat/types";
import "hardhat-typechain";


const config: HardhatUserConfig = {

  solidity: {
    compilers: [{ version: "0.4.24", settings: {} }, { version: "0.5.1", settings: {} }],
  },
  networks: {
    localhost: {
      url: "http://127.0.0.1:8545/",

    },
    hardhat: {

    },
    coverage: {
      url: "http://127.0.0.1:8555", // Coverage launches its own ganache-cli client
    },
  },

};

export default config