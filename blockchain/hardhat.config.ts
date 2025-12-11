import type { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";

const config: HardhatUserConfig = {
  solidity: "0.8.20",
  networks: {
    geth: {
      url: "http://ip_no1:8545",
      accounts: [PRIVATE_KEY],
    },
  },
};

export default config;
