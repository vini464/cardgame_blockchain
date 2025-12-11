import hre from "hardhat";
import { expect } from "chai";

describe("Trades Contract", function () {
  let owner,
    playerA,
    playerB,
    cardsFactory,
    cards,
    tradesFactory,
    trades,
    cardsAddr;

  beforeEach(async () => {
    [owner, playerA, playerB] = await hre.ethers.getSigners();

    cardsFactory = await ethers.getContractFactory("Cards");
    cards = await cardsFactory.deploy();
    await cards.waitForDeployment(); // deploy cards contract
    cardsAddr = await cards.getAddress();

    tradesFactory = await ethers.getContractFactory("Trades");
    trades = await tradesFactory.deploy(cardsAddr);
    await trades.waitForDeployment(); // deploy trade contract

    // giving cards to players
    await cards.mintCard(playerA, 1, 1);
    await cards.mintCard(playerB, 2, 1);

    // aproving contract to move cards:
    await cards.connect(playerA).setApprovalForAll(await trades.getAddress(), true);
    await cards.connect(playerB).setApprovalForAll(await trades.getAddress(), true);
  });

  it("PlayerA trade card 1 for playerB card 2", async () => {
    await trades.connect(playerA).createOffer(1, playerB.address, 2); // cria a oferta para o jogador B

    // ambos aceitam a transação
    await trades.connect(playerA).acceptOffer(0);
    await trades.connect(playerB).acceptOffer(0);

    expect(await cards.balanceOf(playerA.address, 2)).to.equal(1n);
    expect(await cards.balanceOf(playerB.address, 1)).to.equal(1n);
  });
});
