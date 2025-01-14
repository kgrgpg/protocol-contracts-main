import { ethers } from "hardhat";

async function main() {
  // 1. Get the contract factory
  const RateLimitedFactory = await ethers.getContractFactory("RateLimitedZRC20");

  // 2. Deploy it
  const contract = await RateLimitedFactory.deploy(
    "Rate Limited Token",
    "RLT",
    18,            // decimals
    7001,          // chainid_ (example)
    2,             // coinType_ (e.g. Gas=1, ERC20=2, your choice)
    200000,        // gasLimit_
    "0x735b14BB79463307AAcBED86DAf3322B1e6226aB" // systemContractAddress_ (FUNGIBLE_MODULE_ADDRESS for dev?)
  );
  await contract.deployed();

  console.log("Deployed RateLimitedZRC20 at:", contract.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});