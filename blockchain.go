package main 

import (
	"strings"
	"go.mnc.gt/log"
)

// BlockChain doc ...
type BlockChain struct {
	Chain []Block `json:"blocks"`
}

const (
	difficulty = 2
) 

func newBlockChain() BlockChain {
	newBlockChain := BlockChain{}
	newBlockChain.createGenesisBlock()
	return newBlockChain
}

func (blockChain *BlockChain)createGenesisBlock(){
	
	genessisBlock := newBlock(0, "genesis", nil, 0)

	blockChain.Chain = append(blockChain.Chain, genessisBlock) 
	
}

func (blockChain *BlockChain)getDifficulty() (string) {
	str := ""
	for i := 0; i < difficulty; i++ {
		str += "0"
	}

	return str
}

func (blockChain *BlockChain)proofOfWorck(block *Block) string {
	block.Nonce = 0
	computedHash := block.generateHash()
	for !strings.HasPrefix(computedHash, blockChain.getDifficulty()) {
		block.Nonce++
		computedHash = block.generateHash()
	}

	return computedHash
}

func (blockChain *BlockChain)addBlock(block Block) bool {
	previus :=  blockChain.Chain[block.Index-1].generateHash()

	if block.Previus != previus {
		return false
	}
	
	log.Info(blockChain.proofOfWorck(&block))

	blockChain.Chain = append(blockChain.Chain, block)

	return true
}

func (blockChain *BlockChain)mine(transactions interface{}) int64 {
	index:= int64(len(blockChain.Chain))
	previus := blockChain.Chain[index - 1].generateHash()
	newBlock := Block{
		Index: index,
		Transactions: transactions,
		Previus: previus,
		Nonce: 0,
	}

	if !blockChain.addBlock(newBlock){
		return 0
	}

	return newBlock.Index
}