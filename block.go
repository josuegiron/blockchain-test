package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

// Block doc ...
type Block struct {
	Index        int64       `json:"id"`
	Previus      string      `json:"previus"`
	Transactions interface{} `json:"transactions"`
	Nonce        int         `json:"nonce"`
}

func newBlock(index int64, previus string, transactions []interface{}, nonce int) Block {
	newBlock := Block{
		Index:        index,
		Previus:      previus,
		Transactions: transactions,
		Nonce:        nonce,
	}
	return newBlock
}

func (block *Block) generateHash() string {
	jsonBlock, _ := json.Marshal(block)
	return fmt.Sprintf("%x", sha256.Sum256(jsonBlock))
}
