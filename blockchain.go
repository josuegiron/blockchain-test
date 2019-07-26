package main 

import (
	"strings"
	"go.mnc.gt/log"
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
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

	blockChain.announceNewBlock(newBlock) // Announce new block in blockChain net

	return newBlock.Index
}

func (blockChain *BlockChain)consensus(){
	 longestChain := blockChain
	 
	 for _, node := range net.Nodes{
		 
		currentLen := len(longestChain.Chain)

		 url := fmt.Sprintf("http://%v%v/chain", node.URL, port)
		 response, err := http.Get(url)
		if err != nil {
			log.Error(err)
		}
		var newChain BlockChain
		decode := json.NewDecoder(response.Body)
		err = decode.Decode(&newChain)
		if err != nil {
			log.Error(err)
			continue
		}
		
		length:= len(newChain.Chain)
		 if validateNewChain(newChain) && currentLen < length{
			 longestChain = &newChain
		 }
	}

	blockChain = longestChain
}

func validateNewChain(blockChain BlockChain) bool {
	for i, block := range blockChain.Chain{
		if i == 0 {
			continue
		}
		if block.Previus != blockChain.Chain[i-1].generateHash(){
			return false
		}
	}
	return true
}

func (blockChain *BlockChain)validateBlock(block Block) bool{
	if block.Previus != blockChain.Chain[block.Index-1].generateHash(){
		log.Info("El bloque anterior no coincide con el bloque anterior")
		return false
	}

	if !strings.HasPrefix(block.generateHash(), blockChain.getDifficulty()){
		log.Info("El bloque no genera un hash válido...")
		return false
	}
	return true 
}


func (blockChain *BlockChain)announceNewBlock(block Block){
	response, err := json.Marshal(block)
	if err != nil {
		log.Error(err)
		return
	}
	
	body := bytes.NewReader(response)

	for _, node := range net.Nodes {
		url := fmt.Sprintf("http://%v%v/chain/block", node.URL, port)
		_, err := http.Post(url, "application/json", body)
		if err != nil {
			log.Error(err)
			continue
		}
	}	
	return
}