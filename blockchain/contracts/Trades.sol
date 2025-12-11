// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC1155/IERC1155Receiver.sol"; // para receber as cartas durante as trocas, funciona como um intermediário
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";

interface ICards {
  function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes calldata data) external;
}

contract Trades is IERC1155Receiver, ERC165{
  // troca de cartas entre jogadores
  struct Trade {
    address playerA;
    uint cardA; 
    bool acceptedA;

    address playerB;
    uint cardB;
    bool acceptedB;

    bool complete;
  }

  ICards public cards;

  uint nextTradeId;
  mapping(uint => Trade) public trades; // map com as trocas por Id

  constructor(address cardsContract) {
    nextTradeId = 0;
    cards = ICards(cardsContract);
  }

  function createOffer(uint cardA, address playerB, uint cardB) external {
    trades[nextTradeId] = Trade({
      playerA: msg.sender,
      cardA: cardA,
      playerB: playerB,
      cardB: cardB,
      acceptedA: false,
      acceptedB: false,
      complete: false
    });
    nextTradeId++;
  }

  function cancelOffer(uint tradeId) external {
    require(tradeId >= 0 && tradeId < nextTradeId, "cancelOffer: Invalid tradeId");
    Trade storage t = trades[tradeId];
    require(!t.complete, "cancelOffer: this trade is already done");
    require((msg.sender == t.playerA && t.acceptedA) || (msg.sender == t.playerB && t.acceptedB), "cancelOffer: You didn't accepted this trade yet");

    if (msg.sender == t.playerB) {
      cards.safeTransferFrom(address(this), msg.sender, t.cardB, 1, ""); // give card back to player
      t.acceptedB = false;
    } 
    else {
      cards.safeTransferFrom(address(this), msg.sender, t.cardA, 1, ""); // give card back to player
      t.acceptedA = false;
    }
  }

  // aceita uma troca proposta e envia a carta para o intermediário 
  function acceptOffer(uint tradeId) external {
    require(tradeId >= 0 && tradeId < nextTradeId, "acceptOffer: Invalid tradeId");
    Trade storage t = trades[tradeId];
    require(msg.sender == t.playerA || msg.sender == t.playerB, "acceptOffer: You aren't in the trade get out!");
    require((msg.sender == t.playerA && !t.acceptedA) || (msg.sender == t.playerB && !t.acceptedB), "acceptOffer: Your had already accepted this offer");
    uint card = t.cardA;
    if (msg.sender == t.playerB) {
      card = t.cardB;
    }
    cards.safeTransferFrom(msg.sender, address(this), card, 1, "");
    
    if (msg.sender == t.playerA) {
      t.acceptedA = true;
    } else {
      t.acceptedB = true;
    }
    _tryComplete(tradeId);
  }
  
  // tenta completar a troca, verifica se ambos aceitaram e manda as cartas para cada um
  function _tryComplete(uint tradeId) internal {
    Trade storage t = trades[tradeId];
    if (t.acceptedA && t.acceptedB && !t.complete) {
      cards.safeTransferFrom(address(this), t.playerA, t.cardB, 1, "");
      cards.safeTransferFrom(address(this), t.playerB, t.cardA, 1, "");
      t.complete = true;
    }
  }

  // functions to handle with receiving cards from players:
      function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC165, IERC165)
        returns (bool)
    {
        return interfaceId == type(IERC1155Receiver).interfaceId
            || super.supportsInterface(interfaceId);
    }

    function onERC1155Received(
        address operator,
        address from,
        uint256 id,
        uint256 value,
        bytes calldata data
    ) external pure override returns (bytes4) {
        return this.onERC1155Received.selector;
    }

    function onERC1155BatchReceived(
        address operator,
        address from,
        uint256[] calldata ids,
        uint256[] calldata values,
        bytes calldata data
    ) external pure override returns (bytes4) {
        return this.onERC1155BatchReceived.selector;
    }

}

