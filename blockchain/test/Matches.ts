
import hre from "hardhat";
import { expect } from "chai";

const ZERO_ADDRESS = ethers.ZeroAddress;


describe("Match Contract", function () {
  let owner, playerA, playerB, cardsFactory, cards, cardsAddr, matchesFactory, matches;

  beforeEach(async () => {
    [owner, playerA, playerB] = await hre.ethers.getSigners();

    cardsFactory = await ethers.getContractFactory("Cards");
    cards = await cardsFactory.deploy();
    await cards.waitForDeployment(); // deploy trade contract
    cardsAddr = await cards.getAddress();

    matchesFactory = await ethers.getContractFactory("Matches");
    matches = await matchesFactory.deploy(cardsAddr);
    await matches.waitForDeployment(); // deploy trade contract

    // giving cards to players
    await cards.mintCard(playerA, 1, 1);
    await cards.mintCard(playerB, 2, 1);

  });

  it("PlayerA enters in queue", async () => {
    expect(await matches.isWaiting(), "isWaiting should be false").to.equal(false); // não há jogadores na fila
    await matches.connect(playerA).enqueue();
    expect(await matches.isWaiting(), "isWaiting should be true").to.equal(true); // o jogadorA está na fila
    expect(await matches.waitingPlayer(), "waitingPlayer should be playerA").to.equal(playerA);
  });

  it("PlayerA waits for playerB and began a match", async () => {
    expect(await matches.isWaiting(), "isWaiting should be false").to.equal(false); // não há jogadores na fila
    await matches.connect(playerA).enqueue();
    expect(await matches.isWaiting(), "isWaiting should be true").to.equal(true); // o jogadorA está na fila
    expect(await matches.waitingPlayer(), "waitingPlayer should be playerA").to.equal(playerA);

    await matches.connect(playerB).enqueue(); // jogadorB entra e pareia junto com o jogadorA;
    expect(await matches.isWaiting(), "isWaiting should be false").to.equal(false); // não há jogadores na fila
    expect(await matches.waitingPlayer(), "Waiting player should be address(0)").to.equal((ZERO_ADDRESS));
  });

  it("Full match betwen PlayerA and PlayerB, playerA wins", async () => {
    expect(await matches.isWaiting(), "isWaiting should be false").to.equal(false); // não há jogadores na fila
    await matches.connect(playerA).enqueue();
    expect(await matches.isWaiting(), "isWaiting should be true").to.equal(true); // o jogadorA está na fila
    expect(await matches.waitingPlayer(), "waitingPlayer should be playerA").to.equal(playerA);

    await matches.connect(playerB).enqueue(); // jogadorB entra e pareia junto com o jogadorA;
    expect(await matches.isWaiting(), "isWaiting should be false").to.equal(false); // não há jogadores na fila
    expect(await matches.waitingPlayer(), "Waiting player should be address(0)").to.equal((ZERO_ADDRESS));
    await matches.connect(playerA).playCard(0, 1); // matchId, cardId
    await matches.connect(playerB).playCard(0, 2); // matchId, cardId

    const m = await matches.matches(0);

    expect(m.winner, "winner should be playerA").to.equal(playerA);

  });

});
