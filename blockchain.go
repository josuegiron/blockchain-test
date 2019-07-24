package main 

import (
	"strings"
	"log"
)

// Chain doc ...
var chain []Block

const (
	difficulty = 1
) 

func createGenesisBlock(){
	
	genessisBlock := newBlock(0, "0", nil, 0)

	chain = append(chain, genessisBlock) 
	
}

func getDifficulty() (string) {
	str := ""
	for i := 0; i < difficulty; i++ {
		str += "0"
	}

	return str
}

func proofOfWorck(block Block) string {
	
	block.Nonce = 0
	computedHash := block.generateHash()
	for !strings.HasPrefix(computedHash, getDifficulty()) {
		block.Nonce++
		computedHash = block.generateHash()
		log.Println(computedHash)
	}

	return computedHash
}