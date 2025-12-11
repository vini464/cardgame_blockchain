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

	card.SetApprovalForAll(auth, tradeAddr, true);

	// Menu inicial:
FOR:
	for {
		op := input("1 - Jogar;\n2 - Pegar Booter;\n3 - Trocar Carta;\n4 - Ver Coleção;\n5 - Ver Partidas\n6 - Ver Trocas\n 0 - sair")
		switch op {
		case "1":
			id, _ := match.NextMatchId(&bind.CallOpts{})
			tx, _ := match.Enqueue(auth)
			log.Println("tx enviada:", tx.Hash().Hex())
			for {
				waiting, _ := match.IsWaiting(&bind.CallOpts{}) 
				fmt.Println("Waiting for a opponent")
				time.Sleep(time.Second * 2)
				if !waiting {
					break
				}
			}

			maxId, _ := match.NextMatchId(&bind.CallOpts{})

			fmt.Println("Voce está jogando nas seguintes partidas: ", id)
			for id := 0; id < int(maxId.Int64()); id++ {
				currentMatch, err := match.GetMatch(&bind.CallOpts{}, big.NewInt(int64(id)))
				if err == nil && !currentMatch.Finished && (currentMatch.PlayerA == auth.From || currentMatch.PlayerB == auth.From) {
					fmt.Println("-> ", id)
				}
			}
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
				tx, err := trade.CreateOffer(auth, big.NewInt(1), playerAddr, big.NewInt(2))
				if err != nil {
					log.Println(err)
				} else {
					log.Println("tx enviada:", tx.Hash().Hex())
				}
			case "2":
				tradeId := intInput("Digite o Id da troca: ")
				tx, err := trade.AcceptOffer(auth, big.NewInt(int64(tradeId)))
				if err != nil {
					log.Println(err)
				} else {
					log.Println("tx enviada:", tx.Hash().Hex())
				}
			case "3":
				tradeId := intInput("Digite o Id da troca: ")
				tx, err := trade.CancelOffer(auth, big.NewInt(int64(tradeId)))
				if err != nil {
					log.Println(err)
				} else {
					log.Println("tx enviada:", tx.Hash().Hex())
				}
			default:
				fmt.Println("Uknown Command - voltando para o menu inicial")
			}
		case "4":
			for id := 1; id < 11; id++ {
				bal, _ := card.BalanceOf(&bind.CallOpts{}, auth.From, big.NewInt(int64(id)))
				fmt.Println("Card: ", id, " Quantidade: ", bal)
			}
		case "5":
			maxId, _ := match.NextMatchId(&bind.CallOpts{})
			for id := 0; id < int(maxId.Int64()); id++ {
				currentMatch, err := match.GetMatch(&bind.CallOpts{}, big.NewInt(int64(id)))
				if err == nil {
					fmt.Println("Match #", id)
					fmt.Println("playerA  -> ", currentMatch.PlayerA.Hex())
					fmt.Println("playerB  -> ", currentMatch.PlayerB.Hex())
					fmt.Println("CardA    -> ", currentMatch.CardA.Int64())
					fmt.Println("CardB    -> ", currentMatch.CardB.Int64())
					fmt.Println("Finished -> ", currentMatch.Finished)
					fmt.Println("Winner   -> ", currentMatch.Winner.Hex())
				}
			}
		case "6":
			maxId, _ := trade.NextTradeId(&bind.CallOpts{})
			for id := 0; id < int(maxId.Int64()); id++ {
				currentTrade, err := trade.GetTrade(&bind.CallOpts{}, big.NewInt(int64(id)))
				if err == nil {
					fmt.Println("Trade #", id)
					fmt.Println("PlayerA   -> ", currentTrade.PlayerA.Hex())
					fmt.Println("PlayerB   -> ", currentTrade.PlayerB.Hex())
					fmt.Println("CardA     -> ", currentTrade.CardA.Int64())
					fmt.Println("CardB     -> ", currentTrade.CardB.Int64())
					fmt.Println("AcceptedA -> ", currentTrade.AcceptedA)
					fmt.Println("AcceptedB -> ", currentTrade.AcceptedB)
					fmt.Println("Complete  -> ", currentTrade.Complete)

				}
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
