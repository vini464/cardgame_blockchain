package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/vini464/cardgame_blockchain/cards"
	"github.com/vini464/cardgame_blockchain/matches"
	"github.com/vini464/cardgame_blockchain/trades"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	key := input("Insira sua chave: \n")

	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatal(err)
	}

	chainID := big.NewInt(31337)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	contractHex := input("Insira o endereço do seu contrato Match:\n")
	matchAddr := common.HexToAddress(contractHex)
	contractHex = input("Insira o endereço do seu contrato Trade:\n")
	tradeAddr := common.HexToAddress(contractHex)
	contractHex = input("Insira o endereço do seu contrato Card:\n")
	cardAddr := common.HexToAddress(contractHex)

	match, err := matches.NewMatches(matchAddr, client)
	if err != nil {
		log.Fatal("Erro ao criar match:", err)
	}

	trade, err := trades.NewTrades(tradeAddr, client)
	if err != nil {
		log.Fatal("Erro ao criar trade:", err)
	}

	card, err := cards.NewCards(cardAddr, client)
	if err != nil {
		log.Fatal("Erro ao criar card:", err)
	}

	tx, _ := card.OpenBooster(auth)
	log.Println("tx enviada:", tx.Hash().Hex())

	playerHex := input("Insira o endereço do playerB:\n")

	playerAddr := common.HexToAddress(playerHex)

	tx, _ = trade.CreateOffer(auth, big.NewInt(1), playerAddr, big.NewInt(2))
	log.Println("tx enviada:", tx.Hash().Hex())

	tx, _ = match.Enqueue(auth)
	log.Println("tx enviada:", tx.Hash().Hex())
}

func input(text string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(text)
	scanner.Scan()
	return scanner.Text()
}
