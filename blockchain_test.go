package main

import ( 
	"testing"
)

func TestGetDifficulty(t *testing.T){
	t.Error(getDifficulty())
}

func TestGetHash(t *testing.T){
	createGenesisBlock()
	index := len(chain)
	newBlock := newBlock(index, chain[index-1].generateHash(), nil, 0)

	if !addBlock(&newBlock) {
		t.Error("no se ha agregado el bloque")
	}

}