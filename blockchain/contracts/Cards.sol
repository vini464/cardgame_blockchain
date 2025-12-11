// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";


contract Cards is ERC1155 {

  mapping(uint => uint) public cardPower; // define uma carta com ID -> Poder

  constructor() ERC1155("") {
    // cartas comuns
    cardPower[1] = 7;
    cardPower[2] = 3;
    cardPower[3] = 5;
    cardPower[4] = 6;
    cardPower[5] = 2;
    // cartas raras
    cardPower[6] = 12;
    cardPower[7] = 11;
    cardPower[8] = 14;
    cardPower[9] = 13;
    // carta lendaria
    cardPower[10] = 20;
  }

  // da uma carta específica para um jogador
  function mintCard(address to, uint cardId, uint ammount) public {
    require(cardId > 0 && cardId <= 10, "mintCard: Invalid cardId");
    _mint(to, cardId, ammount, "");
  }

  // Abre um booster para um jogador
  function giveBooster(address player) public {
    uint rand = uint(keccak256(abi.encodePacked(block.prevrandao, block.timestamp, player, address(this))));
    for (uint i =0; i < 5; i++) {
      uint cardId = (rand%10) +1; 
      rand = uint(keccak256(abi.encode(rand)));
      mintCard(player, cardId, 1); // mint apenas uma cópia da carta
    }
  }
  function openBooster() external {
    giveBooster(msg.sender);
  }
  function power(uint cardId) external returns (uint) {
    return cardPower[cardId];
  }
}

