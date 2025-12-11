package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
		log.Fatal("Erro ao criar transactor:", err)
	}

	contractHex := input("Insira o endereço do seu contrato:\n")
	contractAddr := common.HexToAddress(contractHex)

	match, _ := NewMatchContract(common.HexToAddress("0xcontrato"), client)

	tx, _ := match.EnterQueue(auth)
	log.Println("tx enviada:", tx.Hash().Hex())
}

func input(text string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(text)
	scanner.Scan()
	return scanner.Text()
}
