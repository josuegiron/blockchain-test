package main

import (
	"crypto/sha256"
)

// Block doc...
type Block struct{
	Index []byte
	Transactions []byte
	Timestap string
	PreviusBlock []byte
}

func newBlock(transactions []byte) Block{
	newBlock := Block{
		Transactions: transactions,
		} 
	return newBlock
}


func (block *Block) encript(){
	block.Index = sha256.Sum256(block.Transactions)


}
