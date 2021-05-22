import { HardhatUserConfig } from "hardhat/types";
import "hardhat-typechain";
import "hardhat-deploy";
import "@nomiclabs/hardhat-ethers";
import './tasks/faucet';
import './tasks/checkbalance';
import './tasks/checkbalances';
import './tasks/buyzap';
import './tasks/initProvider';
import './tasks/initProviderCurve';
import './tasks/setEndpointParams';
import './tasks/bond';
import './tasks/dispatch';
import './tasks/dispatchCoinGecko';
import './tasks/dispatchCGPriceClient';
import './tasks/dispatchBittrex';
import './tasks/checkClient';
import './tasks/dispatchFun';


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