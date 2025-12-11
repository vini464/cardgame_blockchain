// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface ICards {
  function balanceOf(address owner, uint id) external view returns (uint);
  function power(uint id) external view returns (uint);
  function giveBooster(address player) external;
}

contract Matches {
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

  bool public isWaiting;
  address public waitingPlayer;

  constructor(address cardsAddress) {
    cards = ICards(cardsAddress);
    isWaiting = false;
    nextMatchId = 0;
  }

  function enqueue() external {
    require(msg.sender != waitingPlayer, "You are already waiting");
    if (!isWaiting) {
      isWaiting = true;
      waitingPlayer = msg.sender;
    }
    else {
      matches[nextMatchId] = Match({
        playerA: waitingPlayer,
        cardA: 0,
        playerB: msg.sender,
        cardB: 0,
        finished: false,
        winner: address(0)
      });

      isWaiting = false;
      waitingPlayer = address(0);
      nextMatchId++;
    }
  }

  function playCard(uint matchId, uint cardId) external {
    require(matchId >= 0 && matchId < nextMatchId, "playCard: Invalid matchId");
    Match storage m = matches[matchId];
    require(msg.sender == m.playerA || msg.sender == m.playerB, "playCard: You arent a player in this match");
    require(cards.balanceOf(msg.sender, cardId) > 0, "playCard: You dont have that card bro");
    
    if (msg.sender == m.playerA) {
      m.cardA = cardId;
    }
    else {
      m.cardB = cardId;
    }

    if (m.cardA != 0 && m.cardB != 0) {
      m.finished = true;
      if (cards.power(m.cardA) > cards.power(m.cardB)) {
        m.winner = m.playerA;
        cards.giveBooster(m.playerA);
      }
      else if (cards.power(m.cardA) < cards.power(m.cardB)) {
        m.winner = m.playerB;
        cards.giveBooster(m.playerB);
      }
    }

  }

  function getMatch(uint matchId) external view returns (Match memory){
    require(matchId >= 0 && matchId < nextMatchId, "getMatch: invalid Id");
    return matches[matchId];
  }

}
