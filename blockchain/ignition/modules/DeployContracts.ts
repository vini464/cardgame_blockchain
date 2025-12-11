import { ethers } from "hardhat";

async function main() {
  console.log("\n====Deploying Cards=====\n");
  const cards = await ethers.deployContract("Cards");
  await cards.waitForDeployment();
  const cardsAddress = await cards.getAddress();
  console.log("Cards contract deployed at:", cardsAddress);


  console.log("\n====Deploying Trades=====\n");
  const trades = await ethers.deployContract("Trades", [cardsAddress]);
  await trades.waitForDeployment();
  const tradesAddress = await trades.getAddress();
  console.log("Trades contract deployed at:", tradesAddress);

  console.log("\n====Deploying Match=====\n");
  const matches = await ethers.deployContract("Matches", [cardsAddress]);
  await matches.waitForDeployment();
  const matchesAddress = await matches.getAddress();
  console.log("Matches contract deployed at:", matchesAddress);
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
