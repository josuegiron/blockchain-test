package main

import ( 
	"testing"
)

func TestGetDifficulty(t *testing.T){
	t.Error("hola")
	t.Error(getDifficulty())
}

func TestGetHash(t *testing.T){
	genessisBlock := newBlock(0, "0", nil, 0)
	t.Error("hola")
	t.Error(proofOfWorck(genessisBlock))
}