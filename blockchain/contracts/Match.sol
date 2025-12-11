// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

interface ICards {
  function balanceOf(address owner, uint id) external view returns (uint);
  function power(uint id) external view returns (uint);
}

contract Match {
  ICards public cards;

  struct Match {
    address playerA;
    uint cardA;
    address playerB;
    uint cardB;
    bool finished;
    address winner;
  }

  uint public nextMatchId;
  mapping(uint => Match) public matches;

  bool isWaiting;
  address waitingPlayer;

  constructor(address cardsAddress) {
    cards = ICards(cardsAddress);
    isWaiting = false;
    nextMatchId = 0;
  }

  function enqueue() external {
    if (!isWaiting) {
      isWaiting = true;
      waitingPlayer = msg.sender;
    } else {
      matches[nextMatchId] = Match({
        playerA: waitingPlayer,
        cardA: 0,
        playerB: msg.sender,
        cardB: 0,
        finished: false,
        winner: address(0)
      });

      isWaiting = false;
      nextMatchId++;
    }
  }

  function playCard(uint matchId, uint cardId) external {
    require(matchId > 0 && matchId < nextMatchId, "playCard: Invalid matchId");
    Match storage m = matches[matchId];
    require(msg.sender == m.playerA || msg.sender == m.playerB, "playCard: You arent a player in this match");
    require(cards.balanceOf(msg.sender, cardId), "playCard: You dont have that card bro");
    
    if (msg.sender == m.playerA) {
      m.cardA = cardId;
    }
    else {
      m.cardB = cardId;
    }

    if (m.cardA != 0 && m.cardB != 0) {
      m.finished = true;
      if (cards.power(cardA) > cards.power(cardB)) {
        m.winner = playerA;
        cards.connect(playerA).openBooster();
      }
      else if (cards.power(cardA) < cards.power(cardB)) {
        m.winner = playerB;
        cards.connect(playerB).openBooster();
      }
    }

  }

}
