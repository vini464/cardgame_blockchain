import hre from "hardhat";
import { expect } from "chai";

describe("Cards Contract", function () {
  let owner, player, cardsFactory, cards;

  beforeEach(async () => {
    [owner, player] = await hre.ethers.getSigners();

    cardsFactory = await ethers.getContractFactory("Cards");
    cards = await cardsFactory.deploy();
    await cards.waitForDeployment(); // deploy trade contract

  });

  it("Should let buy a booster", async () => {
    await cards.connect(player).openBooster();
    let total = 0n;
    for (let id = 1; id <= 10; id++) {
      const bal = await cards.balanceOf(player.address, id); // checa quantas cartas o jogador tem
      total += bal;
    }
    expect(total).to.equal(5n);
  });

});
