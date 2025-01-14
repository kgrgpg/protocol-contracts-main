import "@nomiclabs/hardhat-waffle";
import "@nomicfoundation/hardhat-verify";
import "@typechain/hardhat";
import "tsconfig-paths/register";
import "hardhat-abi-exporter";
import "uniswap-v2-deploy-plugin";
import "solidity-coverage";
import "hardhat-gas-reporter";
import "./tasks/addresses";
import "@openzeppelin/hardhat-upgrades";

import { getHardhatConfigNetworks } from "@zetachain/networks";
import * as dotenv from "dotenv";
import type { HardhatUserConfig } from "hardhat/types";

dotenv.config();

const config: HardhatUserConfig = {
  //@ts-ignore
  etherscan: {
    apiKey: {
      amoy_testnet: process.env.POLYGONSCAN_API_KEY || "",
      base_mainnet: process.env.BASESCAN_API_KEY || "",
      base_testnet: process.env.BASESCAN_API_KEY || "",
      // BSC
      bsc: process.env.BSCSCAN_API_KEY || "",

      bscTestnet: process.env.BSCSCAN_API_KEY || "",
      // ETH
      goerli: process.env.ETHERSCAN_API_KEY || "",
      mainnet: process.env.ETHERSCAN_API_KEY || "",
    },
    //@ts-ignore
    customChains: [
      {
        chainId: 80002,
        network: "amoy_testnet",
        urls: {
          apiURL: "https://api-amoy.polygonscan.com/api",
          browserURL: "https://amoy.polygonscan.com",
        },
      },
      {
        chainId: 84532,
        network: "base_testnet",
        urls: {
          apiURL: "https://api-sepolia.basescan.org/api",
          browserURL: "https://sepolia.basescan.org/",
        },
      },
      {
        chainId: 8453,
        network: "base_mainnet",
        urls: {
          apiURL: "https://api.basescan.org/api",
          browserURL: "https://basescan.org/",
        },
      },
    ],
  },
  networks: {
    ...getHardhatConfigNetworks(),
    hardhat: {},
  },
  solidity: {
    compilers: [
      { version: "0.4.19" /** For zevm/wzeta.sol */ },
      { version: "0.5.10" /** For create2 factory */ },
      {
        settings: {
          optimizer: {
            enabled: true,
            runs: 999999,
          },
        },
        version: "0.5.16",
      },
      {
        settings: {
          optimizer: {
            enabled: true,
            runs: 999999,
          },
        },
        version: "0.6.6" /** For uniswap v2 */,
      },
      { version: "0.7.6" /** For uniswap v3 */ },
      { version: "0.8.7" },
    ],
  },
};
export default config;
