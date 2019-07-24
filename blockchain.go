package main 

import (
	"strings"
	"log"
)

// Chain doc ...
var chain []Block

const (
	difficulty = 2
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

func proofOfWorck(block *Block) string {
	block.Nonce = 0
	computedHash := block.generateHash()
	for !strings.HasPrefix(computedHash, getDifficulty()) {
		block.Nonce++
		computedHash = block.generateHash()
	}

	return computedHash
}

func addBlock(block *Block) bool {
	previus :=  chain[block.Index-1].generateHash()

	if block.Previus != previus {
		return false
	}
	
	log.Println(proofOfWorck(block))

	return true
}