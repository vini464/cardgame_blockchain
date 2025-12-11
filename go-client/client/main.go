package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

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

	// Menu inicial:
FOR:
	for {
		op := input("1 - Jogar;\n2 - Pegar Booter;\n3 - Trocar Carta;\n4 - Ver Coleção; 0 - sair")
		switch op {
		case "1":
			id, _ := match.NextMatchId(&bind.CallOpts{})
			tx, _ := match.Enqueue(auth)
			log.Println("tx enviada:", tx.Hash().Hex())
			for waiting, _ := match.IsWaiting(&bind.CallOpts{}); waiting; {
				fmt.Println("Waiting for a opponent")
				time.Sleep(time.Second * 2)
			}
			fmt.Println("You are playing in match: ", id)
			for id := 1; id < 11; id++ {
				bal, _ := card.BalanceOf(&bind.CallOpts{}, auth.From, big.NewInt(int64(id)))
				fmt.Println("Card: ", id, " Quantidade: ", bal)
			}
			card := intInput("Escolha sua carta:\n> ")
			match.PlayCard(auth, id, big.NewInt(int64(card)))

		case "2":
			tx, _ := card.OpenBooster(auth)
			log.Println("tx enviada:", tx.Hash().Hex())
		case "3":
			opt := input("1 - Criar Troca;\n2 - Aceitar uma troca\n3 - Cancelar uma Troca\n(any) - Voltar\n> ")
			switch opt {
			case "1":
				playerHex := input("Insira o endereço do playerB:\n")
				playerAddr := common.HexToAddress(playerHex)
				tx, _ := trade.CreateOffer(auth, big.NewInt(1), playerAddr, big.NewInt(2))
				log.Println("tx enviada:", tx.Hash().Hex())
			case "2":
				tradeId := intInput("Digite o Id da troca: ")
				tx, _ := trade.AcceptOffer(auth, big.NewInt(int64(tradeId)))
				log.Println("tx enviada:", tx.Hash().Hex())
			case "3":
			default:
				fmt.Println("Uknown Command - voltando para o menu inicial")
				tradeId := intInput("Digite o Id da troca: ")
				tx, _ := trade.CancelOffer(auth, big.NewInt(int64(tradeId)))
				log.Println("tx enviada:", tx.Hash().Hex())
			}
		case "4":
			for id := 1; id < 11; id++ {
				bal, _ := card.BalanceOf(&bind.CallOpts{}, auth.From, big.NewInt(int64(id)))
				fmt.Println("Card: ", id, " Quantidade: ", bal)
			}
		case "0":
			break FOR
		default:
			fmt.Println("Uknown Command")

		}
	}

}

func input(text string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(text)
	scanner.Scan()
	return scanner.Text()
}

func intInput(text string) int {
	var err error
	var num int
	for {
		t := input(text)
		num, err = strconv.Atoi(t)
		if err == nil {
			break
		}
	}
	return num
}
